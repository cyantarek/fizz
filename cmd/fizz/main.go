package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	mailgungo "github.com/mailgun/mailgun-go"

	"github.com/cyantarek/fizz/config"
	"github.com/cyantarek/fizz/internal/fizz/handlers"
	"github.com/cyantarek/fizz/internal/pkg/common/rest"
	"github.com/cyantarek/fizz/internal/pkg/logger"
	"github.com/cyantarek/fizz/internal/services/emails"
	"github.com/cyantarek/fizz/internal/services/emails/mailgun"
	"github.com/cyantarek/fizz/internal/services/emails/postgres"
	"github.com/cyantarek/fizz/internal/services/marketing"
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
	// Infrastructures Layer
	// =========================================================================

	// mailgun
	mailgunClient := mailgungo.NewMailgun(config.Cfg.MailgunDomain, config.Cfg.MailgunAPIKey)
	mailgunEmail := mailgun.New(mailgunClient)

	// postgres
	sqlClient, err := sqlx.Open("postgres", fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=%s", config.Cfg.DBUsername, config.Cfg.DBPassword, config.Cfg.DBHost, config.Cfg.DBPort, config.Cfg.DBName, "disable"))
	if err != nil {
		log.Fatal(err)
	}

	err = sqlClient.Ping()
	if err != nil {
		log.Fatal(err)
	}

	postgresEmailRepository := postgres.New(sqlClient)

	// =========================================================================
	// Services Layer
	// =========================================================================

	emailService := emails.NewEmailService(mailgunEmail, postgresEmailRepository)
	marketingService := marketing.NewMarketingService(mailgunEmail)

	// =========================================================================
	// Handlers Layer
	// =========================================================================

	api := rest.New()
	serverHandler := handlers.New(api, emailService, marketingService)

	go func() {
		if err := serverHandler.Start(":9000"); err != nil {
			log.Fatal(err)
		}
	}()

	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)

	receiveChan := <-sigChan
	logger.Log.Println(receiveChan.String())
	logger.Log.Println(receiveChan.Signal)

	shutdownCtx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()

	logger.Log.Println("emergency termination call. terminating HTTP REST Web server")

	err = api.Shutdown(shutdownCtx)
	if err != nil {
		logger.Log.Println(err)
	}

	logger.Log.Println("HTTP REST server terminated")
}
