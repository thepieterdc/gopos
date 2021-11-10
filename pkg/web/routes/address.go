package routes

import (
	"github.com/labstack/echo/v4"
	"github.com/thepieterdc/gopos/pkg/web/routes/address"
)

func registerAddressRoutes(srv *echo.Echo) {
	// Build the group.
	g := srv.Group("/address")

	// Register the routes.
	g.GET("/normalise", address.NormaliseHandler)
	g.GET("/resolve", address.ResolveHandler)
}
