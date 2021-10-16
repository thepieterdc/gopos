package cmd

import (
	"github.com/labstack/echo/v4"
	"github.com/thepieterdc/gopos/pkg/health"
	"net/http"
)

// HealthHandler handles the /health route.
func HealthHandler(ctx echo.Context) error {
	response := &health.Response{
		Status: true,
	}

	// Send the response.
	return ctx.JSON(http.StatusOK, response)
}
