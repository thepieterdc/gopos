package cmd

import (
	"github.com/labstack/echo/v4"
	"github.com/thepieterdc/gopos/cmd/google"
	"github.com/thepieterdc/gopos/pkg/environment"
	"net/http"
)

func RegisterGoogleRoutes(srv *echo.Echo) {
	// Build the group.
	g := srv.Group("/google")

	// Check whether the API key is filled in.
	if len(environment.GoogleApiKey) == 0 {
		g.Any("/*", func(ctx echo.Context) error {
			srv.Logger.Error("Google API key is not set.")
			return ctx.NoContent(http.StatusServiceUnavailable)
		})
		return
	}

	// Register the routes.
	g.GET("/place/:id", google.PlaceHandler)
}
