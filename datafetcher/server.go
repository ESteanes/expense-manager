package datafetcher

import (
	"context"
	"log"
	"net/http"
	"time"

	"github.com/esteanes/expense-manager/datafetcher/handlers"
	"github.com/esteanes/expense-manager/datafetcher/templates"
	"github.com/esteanes/expense-manager/datafetcher/upclient"

	"github.com/alexedwards/scs/v2"
)

func getInfo(w http.ResponseWriter, r *http.Request) {

}

func NewNowHandler(now func() time.Time) NowHandler {
	return NowHandler{Now: now}
}

type NowHandler struct {
	Now func() time.Time
}

func (nh NowHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	templates.TimeComponent(nh.Now()).Render(r.Context(), w)
}

type GlobalState struct {
	Count int
}

var global GlobalState
var sessionManager *scs.SessionManager

func getHandler(w http.ResponseWriter, r *http.Request) {
	userCount := sessionManager.GetInt(r.Context(), "count")
	component := templates.Page(global.Count, userCount)
	component.Render(r.Context(), w)
}

func postHandler(w http.ResponseWriter, r *http.Request) {
	// Update state.
	r.ParseForm()

	// Check to see if the global button was pressed.
	if r.Form.Has("global") {
		global.Count++
	}
	//TODO: Update session.
	if r.Form.Has("user") {
		currentCount := sessionManager.GetInt(r.Context(), "count")
		sessionManager.Put(r.Context(), "count", currentCount+1)
	}
	// Display the form.
	getHandler(w, r)
}

func handleInfo(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		postHandler(w, r)
		return
	}
	getHandler(w, r)
}

// HandleRequests function to define the routes and start the server
func HandleRequests(upBankToken string, log *log.Logger) {
	sessionManager = scs.New()
	sessionManager.Lifetime = 24 * time.Hour

	// Registering client
	auth := context.WithValue(context.Background(), upclient.ContextAccessToken, upBankToken)
	configuration := upclient.NewConfiguration()
	apiClient := upclient.NewAPIClient(configuration)

	// Creating individual handlers
	accountHandler := handlers.NewAccountHandler(log, apiClient, auth)
	transactionsHandler := handlers.NewTransactionHandler(log, apiClient, auth, accountHandler)
	transactionsCsvHandler := handlers.NewTransactionCsvHandler(log, apiClient, auth)
	staticFileHandler := handlers.NewStaticFileHandler(log)
	mux := http.NewServeMux()
	mux.HandleFunc(accountHandler.Uri, accountHandler.ServeHTTP)
	mux.HandleFunc(transactionsHandler.Uri, transactionsHandler.ServeHTTP)
	mux.HandleFunc(transactionsCsvHandler.Uri, transactionsCsvHandler.ServeHTTP)
	mux.HandleFunc(staticFileHandler.Uri, staticFileHandler.ServeHTTP)
	mux.HandleFunc("/info", getInfo)
	mux.Handle("/time", NewNowHandler(time.Now))
	mux.HandleFunc("/counter", handleInfo)
	log.Println("Serving request at localhost:8080")
	muxWithSessionMiddleware := sessionManager.LoadAndSave(mux)
	if err := http.ListenAndServe("0.0.0.0:8080", muxWithSessionMiddleware); err != nil {
		log.Printf("error listening: %v", err)
	}
}
