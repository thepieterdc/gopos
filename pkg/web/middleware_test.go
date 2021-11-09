package web

import (
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
	"github.com/thepieterdc/gopos/internal/pkg/version"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestVersionHeaderMiddleware(t *testing.T) {
	// Build a webserver.
	e := echo.New()
	e.Use(VersionHeaderMiddleware)
	e.GET("/", func(ctx echo.Context) error {
		return ctx.JSON(http.StatusNoContent, nil)
	})

	// Send the request.
	req := httptest.NewRequest(http.MethodGet, "/", nil)
	res := httptest.NewRecorder()
	e.ServeHTTP(res, req)

	// Assertions.
	assert.Equal(t, http.StatusNoContent, res.Code)
	assert.Equal(t, version.VERSION, res.Header().Get(HeaderVersion))
}
