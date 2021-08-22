package main

import (
	"context"
	"fizz/config"
	"fizz/internal/core/application"
	"fizz/internal/outside/adapter/driven"
	"fizz/internal/pkg/transports/echohttp"
	"fmt"

	"log"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"github.com/mailgun/mailgun-go/v3"

	"fizz/internal/outside/adapter/driving/httphandler"
	"fizz/internal/pkg/logger"
)

var GitCommit string

func main() {
	logger.New()
	logger.Log.Println("Git tag ver:", GitCommit)

	_ = log.New(os.Stdout, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile)

	if config.Cfg.EmailBackend != "" && config.Cfg.EmailBackend != "mailgun" {
		logger.Log.Fatal("invalid email backend")
	}

	if config.Cfg.EmailBackend != "" {
		logger.Log.Info("no email backend set. Defaulting to mailgun")
	}

	// =========================================================================
	// 3rd Party Services Layer
	// =========================================================================

	// mailgun
	mailgunClient := mailgun.NewMailgun(config.Cfg.MailgunDomain, config.Cfg.MailgunAPIKey)

	sqlClient, err := sqlx.Open("postgres", fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=%s", config.Cfg.DBUsername, config.Cfg.DBPassword, config.Cfg.DBHost, config.Cfg.DBPort, config.Cfg.DBName, "disable"))
	if err != nil {
		log.Fatal(err)
	}

	err = sqlClient.Ping()
	if err != nil {
		log.Fatal(err)
	}

	// =========================================================================
	// Pkg Layer
	// =========================================================================

	// =========================================================================
	// Adapter Layer
	// =========================================================================

	mailgunEmail := driven.NewMailgunEmail(mailgunClient)
	postgresEmailRepository := driven.NewEmailPostgres(sqlClient)

	// =========================================================================
	// Middleware Layer
	// =========================================================================

	// =========================================================================
	// Transport Layer
	// =========================================================================

	// http transport for web: echo
	httpTransport := echohttp.New()

	// =========================================================================
	// Service Layer
	// =========================================================================

	emailService := application.NewEmailService(mailgunEmail, postgresEmailRepository)

	// =========================================================================
	// Ports Layer
	// =========================================================================

	httpHandlers := httphandler.New(emailService)
	httpHandlers.Wire(httpTransport.Router)

	// =========================================================================
	// Transport Runner Layer
	// =========================================================================

	// run http transport
	httpTransport.Run()

	// =========================================================================
	// Transport Halting Layer
	// =========================================================================

	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)

	recv := <-sigChan
	logger.Log.Println(recv.String())
	logger.Log.Println(recv.Signal)

	shutdownCtx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()
	httpTransport.Shutdown(shutdownCtx)
}
