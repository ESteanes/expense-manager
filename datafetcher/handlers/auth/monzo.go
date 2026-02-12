package auth

import (
	"log"
	"net/http"

	"github.com/esteanes/up-bank-go/datafetcher/providers/monzo"
	"github.com/esteanes/up-bank-go/datafetcher/templates"
)

// MonzoAuthHandler handles Monzo OAuth2 authentication routes
type MonzoAuthHandler struct {
	AuthUri      string
	CallbackUri  string
	LogoutUri    string
	OAuthManager *monzo.OAuthManager
	Log          *log.Logger
}

// NewMonzoAuthHandler creates a new Monzo auth handler
func NewMonzoAuthHandler(oauthManager *monzo.OAuthManager, log *log.Logger) *MonzoAuthHandler {
	return &MonzoAuthHandler{
		AuthUri:      "/auth/monzo",
		CallbackUri:  "/auth/monzo/callback",
		LogoutUri:    "/auth/monzo/logout",
		OAuthManager: oauthManager,
		Log:          log,
	}
}

// ServeAuth handles requests to the main auth page
func (h *MonzoAuthHandler) ServeAuth(w http.ResponseWriter, r *http.Request) {
	h.Log.Printf("%s %s", r.Method, r.RequestURI)

	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	if h.OAuthManager == nil {
		http.Error(w, "Monzo OAuth not configured", http.StatusServiceUnavailable)
		return
	}

	if h.OAuthManager.IsAuthorized() {
		templates.MonzoAuthorized().Render(r.Context(), w)
	} else {
		authURL := h.OAuthManager.GetAuthorizationURL()
		templates.MonzoConnect(authURL).Render(r.Context(), w)
	}
}

// ServeCallback processes the OAuth callback from Monzo
func (h *MonzoAuthHandler) ServeCallback(w http.ResponseWriter, r *http.Request) {
	h.Log.Printf("%s %s", r.Method, r.RequestURI)

	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	if h.OAuthManager == nil {
		http.Error(w, "Monzo OAuth not configured", http.StatusServiceUnavailable)
		return
	}

	if err := h.OAuthManager.ProcessCallback(r); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	templates.MonzoCallbackSuccess().Render(r.Context(), w)
}

// ServeLogout disconnects the Monzo account
func (h *MonzoAuthHandler) ServeLogout(w http.ResponseWriter, r *http.Request) {
	h.Log.Printf("%s %s", r.Method, r.RequestURI)

	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	if h.OAuthManager == nil {
		http.Error(w, "Monzo OAuth not configured", http.StatusServiceUnavailable)
		return
	}

	if err := h.OAuthManager.Logout(); err != nil {
		h.Log.Printf("Error during logout: %v", err)
	}

	http.Redirect(w, r, h.AuthUri, http.StatusFound)
}
