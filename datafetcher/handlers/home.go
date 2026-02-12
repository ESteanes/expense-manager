package handlers

import (
	"log"
	"net/http"

	"github.com/esteanes/up-bank-go/datafetcher/models"
	"github.com/esteanes/up-bank-go/datafetcher/providers/monzo"
	"github.com/esteanes/up-bank-go/datafetcher/templates"
)

type HomeHandler struct {
	Uri             string
	Log             *log.Logger
	IsUpBankEnabled bool
	MonzoOAuth      *monzo.OAuthManager
}

func NewHomeHandler(log *log.Logger, upBankEnabled bool, monzoOAuth *monzo.OAuthManager) *HomeHandler {
	return &HomeHandler{
		Uri:             "/",
		Log:             log,
		IsUpBankEnabled: upBankEnabled,
		MonzoOAuth:      monzoOAuth,
	}
}

func (h *HomeHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	h.Log.Printf("%s %s params: %s", r.Method, r.RequestURI, r.URL.Query())

	// "/" pattern matches all unmatched paths in ServeMux, so return 404 for non-root paths
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}

	switch r.Method {
	case http.MethodGet:
		h.Get(w, r)
	case http.MethodPost:
		h.Post(w, r)
	default:
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}

// Post handles POST requests (placeholder)
func (h *HomeHandler) Post(w http.ResponseWriter, r *http.Request) {}

// GetAccounts returns a channel of accounts (for use by other handlers)
func (h *HomeHandler) Get(w http.ResponseWriter, r *http.Request) {
	providers := []models.ProviderStatus{
		{
			Name:        "Up Bank",
			Description: "Australian digital bank",
			IconLabel:   "UP",
			IsUpBank:    true,
			Status:      getUpBankStatus(h.IsUpBankEnabled),
			IsConnected: h.IsUpBankEnabled,
			ShowAction:  false,
		},
		{
			Name:        "Monzo",
			Description: "UK digital bank",
			IconLabel:   "M",
			IsUpBank:    false,
			Status:      getMonzoStatus(h.MonzoOAuth),
			IsConnected: h.MonzoOAuth != nil && h.MonzoOAuth.IsAuthorized(),
			ActionURL:   "/auth/monzo",
			ActionLabel: getMonzoActionLabel(h.MonzoOAuth),
			ShowAction:  h.MonzoOAuth != nil,
		},
	}

	templates.Home("Bank Aggregator", providers).Render(r.Context(), w)
}

func getUpBankStatus(enabled bool) string {
	if enabled {
		return "Connected"
	}
	return "Not Configured"
}

func getMonzoStatus(oauth *monzo.OAuthManager) string {
	if oauth == nil {
		return "Not Configured"
	}
	if oauth.IsAuthorized() {
		return "Connected"
	}
	return "Not Connected"
}

func getMonzoActionLabel(oauth *monzo.OAuthManager) string {
	if oauth != nil && oauth.IsAuthorized() {
		return "Manage"
	}
	return "Connect"
}
