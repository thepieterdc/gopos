package address

import (
	"github.com/labstack/echo/v4"
	"github.com/thepieterdc/gopos/pkg/location"
	"github.com/thepieterdc/gopos/pkg/location/resolvers"
	"net/http"

	"github.com/thepieterdc/gopos/pkg/web"
)

// resolveRequestQuery query parameters of the /address/resolve route.
type resolveRequestQuery struct {
	Country  string `query:"country"`
	Query    string `query:"query" validate:"required"`
	Resolver string `query:"resolver"`
}

// resolveResponse result of the /address/resolve route.
type resolveResponse struct {
	AddressInfo    location.AddressInfo `json:"address_info"`
	DisplayAddress string               `json:"display_address"`
	Query          string               `json:"query"`
	Resolver       string               `json:"resolver"`
}

// ResolveHandler handles the /address/resolve route.
//
// @Summary Resolve an address query into structured address information.
// @Param query query string true "Address to resolve."
// @Param country query string false "Country of the address, improves accuracy if known."
// @Param resolver query string false "Resolver to use." Enums(libpostal, google) default(libpostal)
// @Produce json
// @Success 200 {object} resolveResponse
// @Failure 400 "Missing or invalid query parameters."
// @Router /address/resolve [get]
func ResolveHandler(ctx echo.Context) error {
	// Parse the arguments.
	input := new(resolveRequestQuery)
	if err := web.ParseAndValidate(&ctx, input); err != nil {
		return err
	}

	// Find the resolver or set the default.
	resolverName := input.Resolver
	if len(resolverName) == 0 {
		resolverName = resolvers.LibPostalResolverName
	}
	resolver, err := resolvers.GetByName(resolverName)
	if err != nil {
		return err
	}

	// Build the options.
	options := resolvers.ResolvingOptions{
		Country: input.Country,
	}

	// Resolve the input query.
	resolved, err := (*resolver).Resolve(input.Query, options)
	if err != nil {
		return err
	}

	// Build the response.
	response := resolveResponse{
		AddressInfo:    resolved.AddressInfo,
		DisplayAddress: resolved.DisplayAddress,
		Query:          input.Query,
		Resolver:       (*resolver).Name(),
	}
	return ctx.JSON(http.StatusOK, response)
}
