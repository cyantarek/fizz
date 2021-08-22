package httphandler

import (
	"github.com/labstack/echo/v4"
)

func (h HTTPHandler) Wire(router *echo.Echo) {
	router.POST("/api/v1/send", h.send)
	router.GET("/api/v1/lookup/:id", h.lookupStatus)

	router.GET("/health", func(ctx echo.Context) error {
		return ctx.String(200, "healthy")
	})
}
