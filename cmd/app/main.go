package main

import (
	"context"
	"fizz/internal/pkg/logger"
	"fizz/internal/pkg/transports/http"
	"fizz/internal/pkg/transports/middlewares"
	"log"
	"os"
	"os/signal"
	"syscall"
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

	// =========================================================================
	// Pkg Layer
	// =========================================================================

	// =========================================================================
	// Adapter Layer
	// =========================================================================

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

	// =========================================================================
	// Ports Layer
	// =========================================================================

	// =========================================================================
	// Transport Runner Layer
	// =========================================================================

	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM, syscall.SIGKILL, syscall.SIGSTOP)

	// run http transport
	httpTransport.Run()

	// =========================================================================
	// Transport Halting Layer
	// =========================================================================

	shutdownCtx := context.Background()

	<-sigChan
	httpTransport.Shutdown(shutdownCtx)
}
