package httphandler

import (
	"encoding/json"
	"fizz/internal/core/application/dto"
	"fizz/internal/pkg/logger"
	"net/http"
)

type sendRequest struct {
	From    string   `json:"from"`
	Subject string   `json:"subject"`
	Body    string   `json:"body"`
	To      []string `json:"to"`
	Cc      []string `json:"cc"`
}

type sendResponse struct {
	Message string `json:"message"`
}

type lookupResponse struct {
	ID           string   `json:"id,omitempty"`
	From         string   `json:"from,omitempty"`
	To           []string `json:"to,omitempty"`
	EmailBackend string   `json:"email_backend,omitempty"`
	Status       string   `json:"status,omitempty"`
}

const (
	MAILGUN = "MAILGUN"
)

func (h HTTPHandler) lookupStatus(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		h.errorJSON(w, sendResponse{
			Message: err.Error(),
		}, http.StatusBadRequest)

		return
	}

	emailID := r.FormValue("id")

	emailLookup, err := h.emailService.LookupStatus(r.Context(), emailID)
	if err != nil {
		h.errorJSON(w, sendResponse{
			Message: err.Error(),
		}, http.StatusBadRequest)

		return
	}

	h.successJSON(w, emailLookup)
}

func (h HTTPHandler) send(w http.ResponseWriter, r *http.Request) {
	var in sendRequest

	err := json.NewDecoder(r.Body).Decode(&in)
	if err != nil {
		h.errorJSON(w, sendResponse{
			Message: err.Error(),
		}, http.StatusBadRequest)

		return
	}

	logger.Log.Println("request received")

	err = h.emailService.Send(r.Context(), dto.SendEmail{
		From:    in.From,
		To:      in.To,
		Cc:      in.Cc,
		Subject: in.Subject,
		Body:    in.Body,
	}, MAILGUN)
	if err != nil {
		h.errorJSON(w, sendResponse{
			Message: err.Error(),
		}, http.StatusInternalServerError)

		return
	}

	h.successJSON(w, sendResponse{
		Message: "email queued successfully",
	})
}
