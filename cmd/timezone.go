package cmd

import (
	"github.com/bradfitz/latlong"
	"github.com/labstack/echo/v4"
	"github.com/thepieterdc/gopos/pkg/request"
	"github.com/thepieterdc/gopos/pkg/timezone"
	"net/http"
)

// TimezoneHandler handles the /timezone route.
func TimezoneHandler(ctx echo.Context) error {
	// Parse the arguments.
	input := new(timezone.RequestQuery)
	if err := request.ParseAndValidate(&ctx, input); err != nil {
		return err
	}

	// Find the timezone for this coordinate pair.
	tz := latlong.LookupZoneName(input.Latitude, input.Longitude)

	// Build the response.
	response := &timezone.Response{
		Latitude:  input.Latitude,
		Longitude: input.Longitude,
		Timezone:  tz,
	}

	// Send the response.
	return ctx.JSON(http.StatusOK, response)
}
