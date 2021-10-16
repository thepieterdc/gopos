package web

import (
	"github.com/labstack/echo/v4"
	"github.com/thepieterdc/gopos/pkg/database"
)

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
