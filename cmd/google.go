package cmd

import (
	"github.com/labstack/echo/v4"
	log "github.com/sirupsen/logrus"
	"github.com/thepieterdc/gopos/cmd/google"
	"github.com/thepieterdc/gopos/pkg/configuration"
	"github.com/thepieterdc/gopos/pkg/logging"
	"net/http"
)

// Get the configuration.
var config = configuration.Configure()

func RegisterGoogleRoutes(srv *echo.Echo) {
	// Build the group.
	g := srv.Group("/google")

	// Check whether the API key is filled in.
	if len(config.GoogleApiKey()) == 0 {
		g.Any("/*", func(ctx echo.Context) error {
			log.WithFields(logging.RunningStage()).WithFields(logging.GoogleComponent()).Warn("Google Maps API key is not set.")
			return ctx.NoContent(http.StatusServiceUnavailable)
		})
		return
	}

	// Register the routes.
	g.GET("/place/:id", google.PlaceHandler)
}
