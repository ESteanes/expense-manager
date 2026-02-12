package main

import (
	"fmt"
	"log"
	"os"

	"github.com/esteanes/up-bank-go/datafetcher"
)

func main() {
	fmt.Println("Starting server on port 8080...")
	logger := log.New(os.Stdout, "datafetcher:", log.Default().Flags())

	// Up Bank configuration
	upBankToken := os.Getenv("UP_BANK_TOKEN")

	// Monzo OAuth configuration
	monzoClientID := os.Getenv("MONZO_CLIENT_ID")
	monzoClientSecret := os.Getenv("MONZO_CLIENT_SECRET")
	monzoRedirectURL := os.Getenv("MONZO_REDIRECT_URL") // Optional, defaults to https://localhost:8080/auth/monzo/callback

	// TLS configuration (defaults to ./localhost.crt and ./localhost.key)
	tlsCertFile := os.Getenv("TLS_CERT_FILE")
	if tlsCertFile == "" {
		tlsCertFile = "./localhost.crt"
	}
	tlsKeyFile := os.Getenv("TLS_KEY_FILE")
	if tlsKeyFile == "" {
		tlsKeyFile = "./localhost.key"
	}

	// Validate that at least one bank is configured
	upBankConfigured := upBankToken != ""
	monzoConfigured := monzoClientID != "" && monzoClientSecret != ""

	if !upBankConfigured && !monzoConfigured {
		logger.Fatalln("No bank configured. Set UP_BANK_TOKEN for Up Bank and/or MONZO_CLIENT_ID + MONZO_CLIENT_SECRET for Monzo")
		os.Exit(1)
	}

	if upBankConfigured {
		logger.Println("Up Bank: Configured")
	} else {
		logger.Println("Up Bank: Not configured (set UP_BANK_TOKEN to enable)")
	}

	if monzoConfigured {
		logger.Println("Monzo: OAuth configured (visit /auth/monzo to connect)")
	} else {
		logger.Println("Monzo: Not configured (set MONZO_CLIENT_ID and MONZO_CLIENT_SECRET to enable)")
	}

	monzoCfg := datafetcher.MonzoConfig{
		ClientID:     monzoClientID,
		ClientSecret: monzoClientSecret,
		RedirectURL:  monzoRedirectURL,
	}

	tlsCfg := datafetcher.TLSConfig{
		CertFile: tlsCertFile,
		KeyFile:  tlsKeyFile,
	}

	datafetcher.HandleRequests(upBankToken, monzoCfg, tlsCfg, logger)
}
