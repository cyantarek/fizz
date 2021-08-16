package httphandler

import (
	"fizz/internal/core/application/dto"

	"net/http"
)

func (h HTTPHandler) send(w http.ResponseWriter, r *http.Request) {
	err := h.emailService.Send(r.Context(), dto.SendEmail{
		From: "",
		To:   nil,
		Body: "",
	})

	if err != nil {
		return
	}
}
