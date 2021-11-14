package cmd

import (
	"github.com/go-playground/validator"
	"github.com/labstack/echo/v4"
	log "github.com/sirupsen/logrus"
	"github.com/thepieterdc/gopos/internal/pkg/configuration"
	"github.com/thepieterdc/gopos/internal/pkg/logging"
	"github.com/thepieterdc/gopos/pkg/database"
	"github.com/thepieterdc/gopos/pkg/web"
	"github.com/thepieterdc/gopos/pkg/web/routes"
)

// Web command.
func Web() {
	// Initialise logging.
	log.SetFormatter(&log.JSONFormatter{})

	// Load the settings.
	config := configuration.Configure()

	// Attempt to connect to the database.
	db, err := database.Connect(config)
	if err != nil {
		log.WithFields(logging.BootStage()).WithFields(logging.DatabaseComponent()).Fatal(err)
	}

	// Cleanup the database connection.
	defer func() {
		if db != nil {
			if err := db.Disconnect(); err != nil {
				log.WithFields(logging.ShutdownStage()).WithFields(logging.DatabaseComponent()).Fatal(err)
			}
		}
	}()

	// Build the webserver and register all the routes.
	srv := echo.New()
	routes.Register(srv)

	// Register the middlewares.
	srv.Use(web.ContextMiddleware(db))
	srv.Use(web.VersionHeaderMiddleware)
	web.PrometheusMiddleware(srv)

	// Register data validator.
	srv.Validator = &web.Validator{Validator: validator.New()}

	// Start the server.
	srv.Logger.Fatal(srv.Start("0.0.0.0:8000"))
}
