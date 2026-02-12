package monzo

import (
	"context"
	"crypto/rand"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"sync"
	"time"

	"golang.org/x/oauth2"
)

const (
	monzoAuthURL  = "https://auth.monzo.com"
	monzoTokenURL = "https://api.monzo.com/oauth2/token"
)

// OAuthConfig holds the OAuth2 configuration for Monzo
type OAuthConfig struct {
	ClientID     string
	ClientSecret string
	RedirectURL  string
	TokenFile    string // Path to store tokens for persistence
}

// OAuthManager handles Monzo OAuth2 authentication
type OAuthManager struct {
	config       *oauth2.Config
	token        *oauth2.Token
	tokenMu      sync.RWMutex
	state        string
	log          *log.Logger
	tokenFile    string
	authDone     chan struct{}
	authErr      error
	isAuthorized bool
}

// NewOAuthManager creates a new OAuth manager for Monzo
func NewOAuthManager(cfg OAuthConfig, log *log.Logger) (*OAuthManager, error) {
	if cfg.ClientID == "" || cfg.ClientSecret == "" {
		return nil, fmt.Errorf("MONZO_CLIENT_ID and MONZO_CLIENT_SECRET are required")
	}

	if cfg.RedirectURL == "" {
		cfg.RedirectURL = "https://localhost:8080/auth/monzo/callback"
	}

	if cfg.TokenFile == "" {
		homeDir, err := os.UserHomeDir()
		if err != nil {
			homeDir = "."
		}
		cfg.TokenFile = filepath.Join(homeDir, ".monzo-tokens.json")
	}

	oauthConfig := &oauth2.Config{
		ClientID:     cfg.ClientID,
		ClientSecret: cfg.ClientSecret,
		RedirectURL:  cfg.RedirectURL,
		Scopes:       []string{}, // Monzo doesn't use scopes
		Endpoint: oauth2.Endpoint{
			AuthURL:  monzoAuthURL,
			TokenURL: monzoTokenURL,
		},
	}

	manager := &OAuthManager{
		config:    oauthConfig,
		log:       log,
		tokenFile: cfg.TokenFile,
		authDone:  make(chan struct{}),
	}

	// Try to load existing tokens
	if err := manager.loadTokens(); err != nil {
		log.Printf("No existing Monzo tokens found: %v", err)
	} else if manager.token != nil {
		manager.isAuthorized = true
		log.Println("Loaded existing Monzo tokens")
	}

	return manager, nil
}

// IsAuthorized returns whether we have valid tokens
func (m *OAuthManager) IsAuthorized() bool {
	m.tokenMu.RLock()
	defer m.tokenMu.RUnlock()
	return m.token != nil
}

// GetAuthorizationURL returns the URL the user should visit to authorize
func (m *OAuthManager) GetAuthorizationURL() string {
	// Generate a random state for CSRF protection
	b := make([]byte, 16)
	rand.Read(b)
	m.state = base64.URLEncoding.EncodeToString(b)

	return m.config.AuthCodeURL(m.state, oauth2.AccessTypeOffline)
}

// ProcessCallback processes the OAuth callback and returns an error if it fails
func (m *OAuthManager) ProcessCallback(r *http.Request) error {
	// Verify state
	if r.URL.Query().Get("state") != m.state {
		m.authErr = fmt.Errorf("invalid state parameter")
		return m.authErr
	}

	// Get authorization code
	code := r.URL.Query().Get("code")
	if code == "" {
		errMsg := r.URL.Query().Get("error")
		errDesc := r.URL.Query().Get("error_description")
		m.authErr = fmt.Errorf("authorization failed: %s - %s", errMsg, errDesc)
		return m.authErr
	}

	// Exchange code for tokens
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	token, err := m.config.Exchange(ctx, code)
	if err != nil {
		m.authErr = fmt.Errorf("token exchange failed: %w", err)
		return m.authErr
	}

	m.tokenMu.Lock()
	m.token = token
	m.isAuthorized = true
	m.tokenMu.Unlock()

	// Save tokens to file
	if err := m.saveTokens(); err != nil {
		m.log.Printf("Warning: failed to save tokens: %v", err)
	}

	m.log.Println("Monzo authorization successful!")

	// Signal that auth is complete
	select {
	case <-m.authDone:
		// Already closed
	default:
		close(m.authDone)
	}

	return nil
}

// WaitForAuthorization blocks until authorization is complete or context is cancelled
func (m *OAuthManager) WaitForAuthorization(ctx context.Context) error {
	if m.IsAuthorized() {
		return nil
	}

	select {
	case <-m.authDone:
		return m.authErr
	case <-ctx.Done():
		return ctx.Err()
	}
}

// GetValidToken returns a valid access token, refreshing if necessary
func (m *OAuthManager) GetValidToken(ctx context.Context) (string, error) {
	m.tokenMu.Lock()
	defer m.tokenMu.Unlock()

	if m.token == nil {
		return "", fmt.Errorf("not authorized - visit /auth/monzo to authorize")
	}

	// Check if token needs refresh (with 1 minute buffer)
	if m.token.Expiry.Before(time.Now().Add(time.Minute)) {
		m.log.Println("Refreshing Monzo access token...")

		tokenSource := m.config.TokenSource(ctx, m.token)
		newToken, err := tokenSource.Token()
		if err != nil {
			return "", fmt.Errorf("failed to refresh token: %w", err)
		}

		m.token = newToken

		// Save refreshed tokens
		if err := m.saveTokens(); err != nil {
			m.log.Printf("Warning: failed to save refreshed tokens: %v", err)
		}

		m.log.Println("Token refreshed successfully")
	}

	return m.token.AccessToken, nil
}

// GetHTTPClient returns an HTTP client that automatically handles token refresh
func (m *OAuthManager) GetHTTPClient(ctx context.Context) *http.Client {
	m.tokenMu.RLock()
	token := m.token
	m.tokenMu.RUnlock()

	if token == nil {
		return http.DefaultClient
	}

	return m.config.Client(ctx, token)
}

// saveTokens persists tokens to file
func (m *OAuthManager) saveTokens() error {
	if m.tokenFile == "" || m.token == nil {
		return nil
	}

	data, err := json.MarshalIndent(m.token, "", "  ")
	if err != nil {
		return fmt.Errorf("failed to marshal tokens: %w", err)
	}

	if err := os.WriteFile(m.tokenFile, data, 0600); err != nil {
		return fmt.Errorf("failed to write token file: %w", err)
	}

	m.log.Printf("Tokens saved to %s", m.tokenFile)
	return nil
}

// loadTokens loads tokens from file
func (m *OAuthManager) loadTokens() error {
	if m.tokenFile == "" {
		return fmt.Errorf("no token file configured")
	}

	data, err := os.ReadFile(m.tokenFile)
	if err != nil {
		return fmt.Errorf("failed to read token file: %w", err)
	}

	var token oauth2.Token
	if err := json.Unmarshal(data, &token); err != nil {
		return fmt.Errorf("failed to unmarshal tokens: %w", err)
	}

	m.token = &token
	return nil
}

// Logout clears stored tokens
func (m *OAuthManager) Logout() error {
	m.tokenMu.Lock()
	m.token = nil
	m.isAuthorized = false
	m.authDone = make(chan struct{}) // Reset auth channel
	m.tokenMu.Unlock()

	if m.tokenFile != "" {
		if err := os.Remove(m.tokenFile); err != nil && !os.IsNotExist(err) {
			return fmt.Errorf("failed to remove token file: %w", err)
		}
	}

	m.log.Println("Monzo tokens cleared")
	return nil
}
