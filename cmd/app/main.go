package main

import (
	"context"
	"fizz/config"
	"fizz/internal/core/application"
	"fizz/internal/outside/adapter/driven"
	"github.com/mailgun/mailgun-go/v3"
	"log"
	"os"
	"os/signal"
	"syscall"

	"fizz/internal/outside/adapter/driving/httphandler"
	"fizz/internal/pkg/logger"
	"fizz/internal/pkg/transports/http"
	"fizz/internal/pkg/transports/middlewares"
)

var GitCommit string

func main() {
	logger.New()
	logger.Log.Println("Git tag ver:", GitCommit)

	_ = log.New(os.Stdout, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile)

	// =========================================================================
	// 3rd Party Services Layer
	// =========================================================================

	// mailgun
	mailgunClient := mailgun.NewMailgun(config.Cfg.MailgunDomain, config.Cfg.MailgunAPIKey)

	// =========================================================================
	// Pkg Layer
	// =========================================================================

	// =========================================================================
	// Adapter Layer
	// =========================================================================

	mailgunEmail := driven.NewMailgunEmail(mailgunClient)

	// =========================================================================
	// Middleware Layer
	// =========================================================================

	mws := middlewares.NewMiddleware()

	// =========================================================================
	// Transport Layer
	// =========================================================================

	// http transport for web
	httpTransport := http.New(mws)

	// =========================================================================
	// Service Layer
	// =========================================================================

	emailService := application.NewEmailService(mailgunEmail)

	// =========================================================================
	// Ports Layer
	// =========================================================================

	httpHandlers := httphandler.New(emailService)
	httpHandlers.Wire(httpTransport.Router)

	// =========================================================================
	// Transport Runner Layer
	// =========================================================================

	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM, syscall.SIGSTOP)

	// run http transport
	httpTransport.Run()

	// =========================================================================
	// Transport Halting Layer
	// =========================================================================

	shutdownCtx := context.Background()

	<-sigChan
	httpTransport.Shutdown(shutdownCtx)
}
