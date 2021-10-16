package cmd

import (
	"github.com/labstack/echo/v4"
	"github.com/thepieterdc/gopos/cmd/address"
)

func RegisterAddressRoutes(srv *echo.Echo) {
	// Build the group.
	g := srv.Group("/address")

	// Register the routes.
	g.GET("/parse", address.ParseHandler)
}
