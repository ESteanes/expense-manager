package datafetcher

import (
	"context"
	"fmt"
	"log"
	"os/exec"
	"runtime"
	"time"

	"github.com/esteanes/up-bank-go/datafetcher/providers/monzo"
)

const (
	oauthTimeout = 5 * time.Minute
)

// OAuthStartupFlow handles the OAuth authorization flow on application startup
type OAuthStartupFlow struct {
	Log          *log.Logger
	OAuthManager *monzo.OAuthManager
}

// NewOAuthStartupFlow creates a new OAuth startup flow handler
func NewOAuthStartupFlow(log *log.Logger, oauthManager *monzo.OAuthManager) *OAuthStartupFlow {
	return &OAuthStartupFlow{
		Log:          log,
		OAuthManager: oauthManager,
	}
}

// Run executes the OAuth flow if authorization is required
func (f *OAuthStartupFlow) Run() {
	if f.OAuthManager == nil || f.OAuthManager.IsAuthorized() {
		return
	}

	authURL := f.OAuthManager.GetAuthorizationURL()

	f.printAuthBanner(authURL)

	if err := openBrowser(authURL); err != nil {
		f.Log.Printf("Could not open browser automatically: %v", err)
		fmt.Println("Please open the URL above manually in your browser.")
	}

	f.waitForAuthorization()
}

// printAuthBanner displays the authorization banner with the auth URL
func (f *OAuthStartupFlow) printAuthBanner(authURL string) {
	fmt.Println()
	fmt.Println("╔════════════════════════════════════════════════════════════════╗")
	fmt.Println("║                    MONZO AUTHORIZATION REQUIRED                 ║")
	fmt.Println("╠════════════════════════════════════════════════════════════════╣")
	fmt.Println("║                                                                  ║")
	fmt.Println("║  Opening browser for Monzo authorization...                     ║")
	fmt.Println("║                                                                  ║")
	fmt.Println("║  If the browser doesn't open, visit:                            ║")
	fmt.Printf("║  %s\n", authURL)
	fmt.Println("║                                                                  ║")
	fmt.Println("╚════════════════════════════════════════════════════════════════╝")
	fmt.Println()
}

// waitForAuthorization waits for the user to complete authorization
func (f *OAuthStartupFlow) waitForAuthorization() {
	ctx, cancel := context.WithTimeout(context.Background(), oauthTimeout)
	defer cancel()

	f.Log.Println("Waiting for Monzo authorization...")
	if err := f.OAuthManager.WaitForAuthorization(ctx); err != nil {
		f.Log.Printf("Monzo authorization failed or timed out: %v", err)
		fmt.Println("\nMonzo authorization was not completed. You can still authorize later at /auth/monzo")
	} else {
		fmt.Println()
		fmt.Println("✓ Monzo authorization successful!")
		fmt.Println()
	}
}

// openBrowser tries to open the URL in the default browser
func openBrowser(url string) error {
	var cmd string
	var args []string

	switch runtime.GOOS {
	case "windows":
		cmd = "cmd"
		args = []string{"/c", "start", url}
	case "darwin":
		cmd = "open"
		args = []string{url}
	default: // Linux and others
		cmd = "xdg-open"
		args = []string{url}
	}

	return exec.Command(cmd, args...).Start()
}
