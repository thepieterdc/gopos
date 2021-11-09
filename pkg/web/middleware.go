package web

import (
	"github.com/labstack/echo-contrib/prometheus"
	"github.com/labstack/echo/v4"
	"github.com/thepieterdc/gopos/internal/pkg/version"
	"github.com/thepieterdc/gopos/pkg/database"
)

// HeaderVersion HTTP Header that includes the current version.
const HeaderVersion = "X-Gopos-Version"

// GoposContext custom context that exposes the database instance to routes.
type GoposContext struct {
	DB *database.Database

	echo.Context
}

// ContextMiddleware provides the GoposContext as a middleware function.
func ContextMiddleware(db *database.Database) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(ctx echo.Context) error {
			cc := &GoposContext{DB: db, Context: ctx}
			return next(cc)
		}
	}
}

// PrometheusMiddleware registers the Prometheus middleware on the server.
func PrometheusMiddleware(srv *echo.Echo) {
	prom := prometheus.NewPrometheus("gopos", nil)
	prom.Use(srv)
}

// VersionHeaderMiddleware adds the current Gopos version to every HTTP
// response.
func VersionHeaderMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(ctx echo.Context) error {
		ctx.Response().Header().Set(HeaderVersion, version.VERSION)
		return next(ctx)
	}
}
