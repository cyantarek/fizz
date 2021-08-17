package httphandler

import (
	"github.com/gorilla/mux"
	"net/http"
)

func (h HTTPHandler) Wire(router *mux.Router) {
	router.HandleFunc("/api/v1/send", h.send).Methods(http.MethodPost)
	router.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("healthy\n"))
	}).Methods(http.MethodGet)
}
