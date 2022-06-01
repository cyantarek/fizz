package handlers

import (
	"net/http"

	"github.com/labstack/echo/v4"

	"github.com/cyantarek/fizz/internal/pkg/logger"
	"github.com/cyantarek/fizz/internal/services/emails"
)

const (
	MAILGUN = "MAILGUN"
)

func (h Handlers) lookupStatus(e echo.Context) error {
	emailID := e.Param("id")

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

func (h Handlers) send(e echo.Context) error {
	var in sendRequest

	err := e.Bind(&in)
	if err != nil {
		return e.JSON(http.StatusBadRequest, err.Error())
	}

	logger.Log.Println("request received")

	id, err := h.emailService.Send(e.Request().Context(), emails.SendEmail{
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
		ID:      id,
	})
}

func (h Handlers) getStats(e echo.Context) error {
	stats, err := h.marketingService.GetCompleteStats(e.Request().Context())
	if err != nil {
		return err
	}

	return e.JSON(200, stats)
}

type sendRequest struct {
	From    string   `json:"from"`
	Subject string   `json:"subject"`
	Body    string   `json:"body"`
	To      []string `json:"to"`
	Cc      []string `json:"cc"`
}

type sendResponse struct {
	Message string `json:"message"`
	ID      string `json:"id"`
}

type lookupResponse struct {
	ID           string   `json:"id,omitempty"`
	From         string   `json:"from,omitempty"`
	To           []string `json:"to,omitempty"`
	EmailBackend string   `json:"email_backend,omitempty"`
	Status       string   `json:"status,omitempty"`
}
