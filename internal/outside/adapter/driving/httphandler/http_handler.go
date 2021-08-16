package httphandler

import (
	"fizz/internal/core/port/incoming"
	"net/http"

	"github.com/gorilla/mux"
)

type HTTPHandler struct {
	emailService incoming.EmailService
}

func (h HTTPHandler) Wire(router *mux.Router) {
	router.HandleFunc("/ping", nil).Methods(http.MethodGet)
}

func New(emailService incoming.EmailService) *HTTPHandler {
	return &HTTPHandler{emailService: emailService}
}
