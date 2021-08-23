package echohttp

import (
	"context"
	"fizz/config"
	"fizz/internal/pkg/logger"
	"fmt"
	"github.com/labstack/echo-contrib/prometheus"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type EchoHTTP struct {
	Router *echo.Echo
}

func New() *EchoHTTP {
	e := echo.New()

	e.Server.Addr = fmt.Sprintf(":%s", config.Cfg.HttpPort)

	e.Use(middleware.Recover())
	e.Use(middleware.Logger())

	prom := prometheus.NewPrometheus("fizz", nil)
	prom.Use(e)

	return &EchoHTTP{Router: e}
}

func (eh *EchoHTTP) Run() {
	logger.Log.Println("HTTP REST server started on", eh.Router.Server.Addr)
	//h.printRoutes()

	go func() {
		logger.Log.Fatal(eh.Router.Start(eh.Router.Server.Addr))
	}()
}

func (eh *EchoHTTP) Shutdown(ctx context.Context) {
	logger.Log.Println("emergency termination call. terminating HTTP REST Web server")

	err := eh.Router.Shutdown(ctx)
	if err != nil {
		logger.Log.Println(err)
	}

	logger.Log.Println("HTTP REST server terminated")
}
