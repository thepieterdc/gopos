package main

import (
	"github.com/go-playground/validator"
	"github.com/labstack/echo-contrib/prometheus"
	"github.com/labstack/echo/v4"
	log "github.com/sirupsen/logrus"
	"github.com/thepieterdc/gopos/cmd"
	"github.com/thepieterdc/gopos/pkg/configuration"
	"github.com/thepieterdc/gopos/pkg/database"
	"github.com/thepieterdc/gopos/pkg/logging"
	"github.com/thepieterdc/gopos/pkg/web"
)

func main() {
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
	cmd.RegisterAddressRoutes(srv)
	srv.GET("/health", cmd.HealthHandler)
	cmd.RegisterGoogleRoutes(srv)
	srv.GET("/timezone", cmd.TimezoneHandler)

	// Register the custom context.
	srv.Use(web.ContextMiddleware(db))

	// Register the prometheus middleware.
	prom := prometheus.NewPrometheus("gopos", nil)
	prom.Use(srv)

	// Register data validator.
	srv.Validator = &web.Validator{Validator: validator.New()}

	// Start the server.
	srv.Logger.Fatal(srv.Start("0.0.0.0:8000"))
}
