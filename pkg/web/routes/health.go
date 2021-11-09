package routes

import (
	"github.com/labstack/echo/v4"
	"github.com/thepieterdc/gopos/internal/pkg/version"
	"github.com/thepieterdc/gopos/pkg/health"
	"net/http"
)

// HealthHandler handles the /health route.
func HealthHandler(ctx echo.Context) error {
	response := &health.Response{
		Status:  true,
		Version: version.VERSION,
	}

	// Send the response.
	return ctx.JSON(http.StatusOK, response)
}
