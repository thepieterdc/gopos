package routes

import (
	"github.com/labstack/echo/v4"
	echoSwagger "github.com/swaggo/echo-swagger"
	// Generated API documentation.
	_ "github.com/thepieterdc/gopos/docs"
)

func Register(srv *echo.Echo) {
	registerAddressRoutes(srv)
	srv.GET("/docs/*", echoSwagger.WrapHandler)
	srv.GET("/health", HealthHandler)
	registerGoogleRoutes(srv)
	srv.GET("/timezone", TimezoneHandler)
}
