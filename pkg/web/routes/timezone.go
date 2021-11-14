package routes

import (
	"github.com/bradfitz/latlong"
	"github.com/labstack/echo/v4"
	"github.com/thepieterdc/gopos/pkg/web"
	"net/http"
)

// timezoneRequestQuery query parameters of the /timezone route.
type timezoneRequestQuery struct {
	Latitude  float64 `query:"latitude" validate:"required"`
	Longitude float64 `query:"longitude" validate:"required"`
}

// timezoneResponse result of the /timezone route.
type timezoneResponse struct {
	Latitude  float64 `json:"latitude"`
	Longitude float64 `json:"longitude"`
	Timezone  string  `json:"timezone"`
}

// TimezoneHandler handles the /timezone route.
func TimezoneHandler(ctx echo.Context) error {
	// Parse the arguments.
	input := new(timezoneRequestQuery)
	if err := web.ParseAndValidate(&ctx, input); err != nil {
		return err
	}

	// Find the timezone for this coordinate pair.
	tz := latlong.LookupZoneName(input.Latitude, input.Longitude)

	// Build the response.
	response := &timezoneResponse{
		Latitude:  input.Latitude,
		Longitude: input.Longitude,
		Timezone:  tz,
	}

	// Send the response.
	return ctx.JSON(http.StatusOK, response)
}
