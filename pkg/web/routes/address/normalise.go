package address

import (
	"github.com/labstack/echo/v4"
	"net/http"

	postal "github.com/openvenues/gopostal/expand"
	"github.com/thepieterdc/gopos/pkg/web"
)

// requestQuery query parameters of the /address/normalise route.
type requestQuery struct {
	Query string `query:"query" validate:"required"`
}

// response result of the /address/normalise route.
type response struct {
	Normalised []string `json:"normalised"`
	Query      string   `json:"query"`
}

// NormaliseHandler handles the /address/normalise route.
func NormaliseHandler(ctx echo.Context) error {
	// Parse the arguments.
	input := new(requestQuery)
	if err := web.ParseAndValidate(&ctx, input); err != nil {
		return err
	}

	// Normalise the input address.
	normalised := postal.ExpandAddress(input.Query)

	// Build the response.
	ret := response{
		Normalised: normalised,
		Query:      input.Query,
	}

	// Send the response.
	return ctx.JSON(http.StatusOK, ret)
}
