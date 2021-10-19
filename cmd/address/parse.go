package address

import (
	"github.com/labstack/echo/v4"
	log "github.com/sirupsen/logrus"
	"github.com/thepieterdc/gopos/pkg/logging"
	"net/http"

	postal "github.com/openvenues/gopostal/parser"
	"github.com/thepieterdc/gopos/pkg/address/parse"
	"github.com/thepieterdc/gopos/pkg/web"
)

// Initialise the logging fields.
var logger = log.WithFields(logging.RunningStage()).WithFields(logging.AddressComponent())

// ParseHandler handles the /address/parse route.
func ParseHandler(ctx echo.Context) error {
	// Parse the arguments.
	input := new(parse.RequestQuery)
	if err := web.ParseAndValidate(&ctx, input); err != nil {
		return err
	}

	// Build the options.
	options := postal.ParserOptions{Country: "", Language: ""}
	if len(input.Country) > 0 {
		options.Country = input.Country
	}

	// Parse the input address.
	parsed := postal.ParseAddressOptions(input.Query, options)

	// Build the response.
	logger.Infof("Parsed input \"%s\" into %d fields.", input, len(parsed))
	response := make(map[string]interface{})
	for _, entry := range parsed {
		response[entry.Label] = entry.Value
	}

	// Send the response.
	return ctx.JSON(http.StatusOK, response)
}
