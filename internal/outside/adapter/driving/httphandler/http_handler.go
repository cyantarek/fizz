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
	router.HandleFunc("/api/v1/send", h.send).Methods(http.MethodPost)
}

func New(emailService incoming.EmailService) *HTTPHandler {
	return &HTTPHandler{emailService: emailService}
}
