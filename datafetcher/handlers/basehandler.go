package handlers

import (
	"context"
	"log"
	"net/http"

	"github.com/esteanes/expense-manager/datafetcher/upclient"
)

type Handler interface {
	Post(w http.ResponseWriter, r *http.Request)
	Get(w http.ResponseWriter, r *http.Request)
}

type BaseHandler struct {
	Uri      string
	Log      *log.Logger
	UpClient *upclient.APIClient
	UpAuth   context.Context
	Handler  Handler // Embed the Handler interface
}

func (h *BaseHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodPost:
		h.Handler.Post(w, r)
	case http.MethodGet:
		h.Handler.Get(w, r)
	default:
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}
