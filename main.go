package main

import (
	"github.com/go-playground/validator"
	"github.com/labstack/echo-contrib/prometheus"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/thepieterdc/gopos/cmd"
	"github.com/thepieterdc/gopos/pkg/database"
	"github.com/thepieterdc/gopos/pkg/environment"
	"github.com/thepieterdc/gopos/pkg/web"
	"log"
)

func main() {
	// Holder for the database connection.
	var db *database.Database

	// Connect to the database if configured.
	if len(environment.MongoUri) != 0 {
		log.Println("Connecting to the database...")

		var err error
		db, err = database.Connect()
		if err != nil {
			log.Fatal(err)
		}

		log.Println("Connected to the database.")
	}

	// Build the webserver and register all the routes.
	srv := echo.New()
	cmd.RegisterAddressRoutes(srv)
	srv.GET("/health", cmd.HealthHandler)
	cmd.RegisterGoogleRoutes(srv)
	srv.GET("/timezone", cmd.TimezoneHandler)

	// Register the custom context.
	srv.Use(web.ContextMiddleware(db))

	// Register the logging middleware.
	srv.Use(middleware.Logger())

	// Register the prometheus middleware.
	prom := prometheus.NewPrometheus("gopos", nil)
	prom.Use(srv)

	// Register data validator.
	srv.Validator = &web.Validator{Validator: validator.New()}

	// Start the server.
	srv.Logger.Fatal(srv.Start("0.0.0.0:8000"))
}
