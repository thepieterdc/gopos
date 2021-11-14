package routes

import (
	"github.com/labstack/echo/v4"
	"github.com/thepieterdc/gopos/internal/pkg/version"
	"net/http"
)

// healthResponse result of the /health route.
type healthResponse struct {
	Status  bool   `json:"status"`
	Version string `json:"version"`
}

// HealthHandler handles the /health route.
func HealthHandler(ctx echo.Context) error {
	response := &healthResponse{
		Status:  true,
		Version: version.VERSION,
	}

	// Send the response.
	return ctx.JSON(http.StatusOK, response)
}
