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

func (h HTTPHandler) send(w http.ResponseWriter, r *http.Request) {
	var in sendRequest

	err := json.NewDecoder(r.Body).Decode(&in)
	if err != nil {
		h.errorJSON(w, sendResponse{
			Message: err.Error(),
		})

		return
	}

	logger.Log.Println("request received")

	err = h.emailService.Send(r.Context(), dto.SendEmail{
		From:    in.From,
		To:      in.To,
		Cc:      in.Cc,
		Subject: in.Subject,
		Body:    in.Body,
	})
	if err != nil {
		h.errorJSON(w, sendResponse{
			Message: err.Error(),
		})

		return
	}

	h.successJSON(w, sendResponse{
		Message: "email queued successfully",
	})
}
