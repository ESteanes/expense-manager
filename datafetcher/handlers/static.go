package handlers

import (
	"log"
	"net/http"
	"path/filepath"

	"github.com/esteanes/up-bank-go/datafetcher/functions"
)

type StaticFileHandler struct {
	*functions.BaseHandler
}

func NewStaticFileHandler(log *log.Logger) *StaticFileHandler {
	handler := &StaticFileHandler{}
	handler.BaseHandler = &functions.BaseHandler{
		Uri:      "/static/",
		Log:      log,
		UpClient: nil,
		UpAuth:   nil,
		Handler:  handler}
	return handler
}
func (s *StaticFileHandler) Get(w http.ResponseWriter, r *http.Request) {
	if r.RequestURI != "/static/css/output.css" {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
	filePath := r.URL.Path[len("/static/"):]
	fullPath := filepath.Join(".", "static", filePath)
	s.Log.Printf("Attempting to serve path: %s", fullPath)
	http.ServeFile(w, r, fullPath)
}

func (s *StaticFileHandler) Post(w http.ResponseWriter, r *http.Request) {
	s.Log.Println("Unsupported HTTP Method")
	http.Error(w, "Internal Server Error", http.StatusInternalServerError)
}
