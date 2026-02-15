package monzo

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"path/filepath"

	"github.com/zalando/go-keyring"
	"golang.org/x/oauth2"
)

const (
	keychainService = "up-bank-go-monzo"
	keychainAccount = "oauth-token"
)

// TokenStore abstracts token persistence.
type TokenStore interface {
	SaveToken(token *oauth2.Token) error
	LoadToken() (*oauth2.Token, error)
	ClearToken() error
}

// KeychainTokenStore stores tokens in the OS credential store via go-keyring.
type KeychainTokenStore struct{}

func (k *KeychainTokenStore) SaveToken(token *oauth2.Token) error {
	data, err := json.Marshal(token)
	if err != nil {
		return fmt.Errorf("failed to marshal token: %w", err)
	}
	if err := keyring.Set(keychainService, keychainAccount, string(data)); err != nil {
		return fmt.Errorf("failed to store token in keychain: %w", err)
	}
	return nil
}

func (k *KeychainTokenStore) LoadToken() (*oauth2.Token, error) {
	data, err := keyring.Get(keychainService, keychainAccount)
	if err != nil {
		return nil, fmt.Errorf("failed to load token from keychain: %w", err)
	}
	var token oauth2.Token
	if err := json.Unmarshal([]byte(data), &token); err != nil {
		return nil, fmt.Errorf("failed to unmarshal token: %w", err)
	}
	return &token, nil
}

func (k *KeychainTokenStore) ClearToken() error {
	err := keyring.Delete(keychainService, keychainAccount)
	if err != nil {
		return fmt.Errorf("failed to delete token from keychain: %w", err)
	}
	return nil
}

// FileTokenStore stores tokens as JSON in a file with atomic writes.
type FileTokenStore struct {
	path string
}

func (f *FileTokenStore) SaveToken(token *oauth2.Token) error {
	data, err := json.MarshalIndent(token, "", "  ")
	if err != nil {
		return fmt.Errorf("failed to marshal token: %w", err)
	}

	dir := filepath.Dir(f.path)
	tmp, err := os.CreateTemp(dir, ".monzo-tokens-*.tmp")
	if err != nil {
		return fmt.Errorf("failed to create temp file: %w", err)
	}
	tmpName := tmp.Name()

	if _, err := tmp.Write(data); err != nil {
		tmp.Close()
		os.Remove(tmpName)
		return fmt.Errorf("failed to write temp file: %w", err)
	}
	if err := tmp.Close(); err != nil {
		os.Remove(tmpName)
		return fmt.Errorf("failed to close temp file: %w", err)
	}
	if err := os.Chmod(tmpName, 0600); err != nil {
		os.Remove(tmpName)
		return fmt.Errorf("failed to set file permissions: %w", err)
	}
	if err := os.Rename(tmpName, f.path); err != nil {
		os.Remove(tmpName)
		return fmt.Errorf("failed to rename temp file: %w", err)
	}
	return nil
}

func (f *FileTokenStore) LoadToken() (*oauth2.Token, error) {
	data, err := os.ReadFile(f.path)
	if err != nil {
		return nil, fmt.Errorf("failed to read token file: %w", err)
	}
	var token oauth2.Token
	if err := json.Unmarshal(data, &token); err != nil {
		return nil, fmt.Errorf("failed to unmarshal token: %w", err)
	}
	return &token, nil
}

func (f *FileTokenStore) ClearToken() error {
	if err := os.Remove(f.path); err != nil && !os.IsNotExist(err) {
		return fmt.Errorf("failed to remove token file: %w", err)
	}
	return nil
}

// NewTokenStore returns a TokenStore, preferring the OS keychain and falling
// back to file-based storage if the keychain is unavailable.
func NewTokenStore(tokenFilePath string, logger *log.Logger) TokenStore {
	// Test if keychain is available by doing a write/read/delete cycle
	const testKey = "oauth-token-probe"
	if err := keyring.Set(keychainService, testKey, "test"); err == nil {
		if _, err := keyring.Get(keychainService, testKey); err == nil {
			_ = keyring.Delete(keychainService, testKey)
			logger.Println("Using OS keychain for Monzo token storage")
			return &KeychainTokenStore{}
		}
		_ = keyring.Delete(keychainService, testKey)
	}

	logger.Printf("Keychain unavailable, using file-based token storage at %s", tokenFilePath)
	return &FileTokenStore{path: tokenFilePath}
}
