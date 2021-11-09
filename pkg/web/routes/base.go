package routes

import "github.com/labstack/echo/v4"

func Register(srv *echo.Echo) {
	registerAddressRoutes(srv)
	srv.GET("/health", HealthHandler)
	registerGoogleRoutes(srv)
	srv.GET("/timezone", TimezoneHandler)
}
