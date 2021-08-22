package httphandler

import (
	"fizz/internal/core/application/dto"
	"fizz/internal/pkg/logger"
	"github.com/labstack/echo/v4"
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

func (h HTTPHandler) lookupStatus(e echo.Context) error {
	emailID := e.FormValue("id")

	emailLookup, err := h.emailService.LookupStatus(e.Request().Context(), emailID)
	if err != nil {
		return e.JSON(http.StatusInternalServerError, err.Error())
	}

	return e.JSON(200, lookupResponse{
		ID:           emailLookup.ID,
		From:         emailLookup.From,
		To:           emailLookup.To,
		EmailBackend: emailLookup.EmailBackend,
		Status:       emailLookup.Status,
	})
}

func (h HTTPHandler) send(e echo.Context) error {
	var in sendRequest

	err := e.Bind(&in)
	if err != nil {
		return e.JSON(http.StatusBadRequest, err.Error())
	}

	logger.Log.Println("request received")

	err = h.emailService.Send(e.Request().Context(), dto.SendEmail{
		From:    in.From,
		To:      in.To,
		Cc:      in.Cc,
		Subject: in.Subject,
		Body:    in.Body,
	}, MAILGUN)
	if err != nil {
		return e.JSON(http.StatusInternalServerError, err.Error())
	}

	return e.JSON(200, sendResponse{
		Message: "email queued successfully",
	})
}
