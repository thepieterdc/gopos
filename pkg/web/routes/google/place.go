package google

import (
	"github.com/labstack/echo/v4"
	log "github.com/sirupsen/logrus"
	"github.com/thepieterdc/gopos/internal/pkg/google"
	"github.com/thepieterdc/gopos/internal/pkg/logging"
	"github.com/thepieterdc/gopos/pkg/web"
	"net/http"
)

// Initialise the logger.
var logger = log.WithFields(logging.RunningStage()).WithFields(logging.GoogleComponent())

// PlaceHandler handles the /google/place route.
//
// @Summary Get the details of a Google Place ID.
// @Description Returns 503 when no Google Maps API key is configured.
// @Param id path string true "Google Place ID."
// @Produce json
// @Success 200 {object} google.GooglePlaceDetails
// @Failure 400 "Missing place ID."
// @Failure 503 "Google Maps API key is not configured."
// @Router /google/place/{id} [get]
func PlaceHandler(c echo.Context) error {
	// Cast the context.
	ctx := c.(*web.GoposContext)

	// Validate the request.
	id := ctx.Param("id")
	if len(id) == 0 {
		// TODO: Proper error handling.
		return ctx.NoContent(http.StatusBadRequest)
	}

	// Attempt to get the place details from the database cache.
	if ctx.DB != nil {
		placeDetails, err := ctx.DB.FindPlaceDetailsById(id)
		if err != nil {
			return err
		}

		// If the place details were found, return them to the client.
		if placeDetails != nil {
			logger.WithField("source", "db").Infof("Fetched place details for %s from cache.", id)
			return ctx.JSON(http.StatusOK, placeDetails)
		}
	}

	// Fetch the place details from Google.
	placeDetails, err := google.GetPlaceDetailsById(id)
	if err != nil {
		return err
	}

	logger.WithField("source", "api").Infof("Fetched place details for %s from Google API.", id)

	// Store the place details in the database.
	if ctx.DB != nil {
		defer ctx.DB.SavePlaceDetails(placeDetails)
	}

	// Send the response back to the client.
	return ctx.JSON(http.StatusOK, placeDetails)
}
