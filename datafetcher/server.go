package datafetcher

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/esteanes/up-bank-go/datafetcher/handlers"
	authHandlers "github.com/esteanes/up-bank-go/datafetcher/handlers/auth"
	unifiedHandlers "github.com/esteanes/up-bank-go/datafetcher/handlers/unified"
	upbankHandlers "github.com/esteanes/up-bank-go/datafetcher/handlers/upbank"
	"github.com/esteanes/up-bank-go/datafetcher/providers"
	"github.com/esteanes/up-bank-go/datafetcher/providers/monzo"
	"github.com/esteanes/up-bank-go/datafetcher/providers/upbank"

	"github.com/alexedwards/scs/v2"
)

// MonzoConfig holds Monzo OAuth configuration
type MonzoConfig struct {
	ClientID     string
	ClientSecret string
	RedirectURL  string
}

// TLSConfig holds optional TLS configuration
type TLSConfig struct {
	CertFile string
	KeyFile  string
}

var sessionManager *scs.SessionManager

// HandleRequests function to define the routes and start the server
func HandleRequests(upBankToken string, monzoCfg MonzoConfig, tlsCfg TLSConfig, log *log.Logger) {
	sessionManager = scs.New()
	sessionManager.Lifetime = 24 * time.Hour

	// Create providers
	var bankProviders []providers.BankProvider

	// Initialize Up Bank provider
	if upBankToken != "" {
		log.Println("Initializing Up Bank provider")
		bankProviders = append(bankProviders, upbank.NewProvider(upBankToken, log))
	}

	// Initialize Monzo OAuth manager and provider
	var monzoOAuthManager *monzo.OAuthManager
	var monzoProvider *monzo.Provider

	if monzoCfg.ClientID != "" && monzoCfg.ClientSecret != "" {
		log.Println("Initializing Monzo OAuth manager")
		var err error
		monzoOAuthManager, err = monzo.NewOAuthManager(monzo.OAuthConfig{
			ClientID:     monzoCfg.ClientID,
			ClientSecret: monzoCfg.ClientSecret,
			RedirectURL:  monzoCfg.RedirectURL,
		}, log)
		if err != nil {
			log.Printf("Failed to initialize Monzo OAuth: %v", err)
		} else {
			monzoProvider = monzo.NewProvider(monzoOAuthManager, log)
			bankProviders = append(bankProviders, monzoProvider)
		}
	}

	// Create aggregator with all enabled providers
	aggregator := providers.NewAggregator(log, bankProviders...)

	// Creating unified handlers
	accountHandler := unifiedHandlers.NewAccountHandler(log, aggregator)
	transactionsHandler := unifiedHandlers.NewTransactionsHandler(log, aggregator, accountHandler)
	transactionsCsvHandler := unifiedHandlers.NewTransactionsCsvHandler(log, aggregator)
	staticFileHandler := upbankHandlers.NewStaticFileHandler(log)
	homeHandler := handlers.NewHomeHandler(log, upBankToken != "", monzoOAuthManager)

	mux := http.NewServeMux()

	// Bank data routes
	mux.HandleFunc(accountHandler.Uri, accountHandler.ServeHTTP)
	mux.HandleFunc(transactionsHandler.Uri, transactionsHandler.ServeHTTP)
	mux.HandleFunc(transactionsCsvHandler.Uri, transactionsCsvHandler.ServeHTTP)
	mux.HandleFunc(staticFileHandler.Uri, staticFileHandler.ServeHTTP)
	mux.HandleFunc(homeHandler.Uri, homeHandler.ServeHTTP)

	// Monzo OAuth routes
	if monzoOAuthManager != nil {
		monzoAuthHandler := authHandlers.NewMonzoAuthHandler(monzoOAuthManager, log)
		mux.HandleFunc(monzoAuthHandler.AuthUri, monzoAuthHandler.ServeAuth)
		mux.HandleFunc(monzoAuthHandler.CallbackUri, monzoAuthHandler.ServeCallback)
		mux.HandleFunc(monzoAuthHandler.LogoutUri, monzoAuthHandler.ServeLogout)
	}

	muxWithSessionMiddleware := sessionManager.LoadAndSave(mux)

	// Start the server in a goroutine so we can handle OAuth flow
	server := &http.Server{
		Addr:    "127.0.0.1:8080",
		Handler: muxWithSessionMiddleware,
	}

	go func() {
		log.Println("Server starting on https://localhost:8080")
		if err := server.ListenAndServeTLS(tlsCfg.CertFile, tlsCfg.KeyFile); err != nil && err != http.ErrServerClosed {
			log.Fatalf("Error starting server: %v", err)
		}
	}()

	// Give server a moment to start
	time.Sleep(100 * time.Millisecond)

	// Handle Monzo OAuth flow on startup if needed
	oauthFlow := NewOAuthStartupFlow(log, monzoOAuthManager)
	oauthFlow.Run()

	// Print final status
	fmt.Println()
	log.Printf("Initialized %d bank provider(s)", aggregator.ProviderCount())
	if monzoOAuthManager != nil && monzoOAuthManager.IsAuthorized() {
		log.Println("Monzo: Connected")
	}
	if upBankToken != "" {
		log.Println("Up Bank: Connected")
	}
	fmt.Println()
	log.Println("Server ready at https://localhost:8080")
	fmt.Println()

	// Block forever (server is running in goroutine)
	select {}
}
