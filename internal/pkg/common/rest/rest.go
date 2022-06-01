package rest

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func New() *echo.Echo {
	r := echo.New()
	r.GET("/helthz", healthz)

	return r
}

func healthz(c echo.Context) error {
	return c.JSON(http.StatusOK, echo.Map{"status": "I'm alive, don't worry"})
}
