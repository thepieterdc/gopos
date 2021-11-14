package resolvers

import (
	expand "github.com/openvenues/gopostal/expand"
	parser "github.com/openvenues/gopostal/parser"
	"github.com/thepieterdc/gopos/pkg/location"
)

// LibPostalResolverName name of the resolver.
const LibPostalResolverName = "libpostal"

// Mapping labels.
const labelCity = "city"
const labelCountry = "country"
const labelPostalCode = "postcode"
const labelStateOrProvince = "state"
const labelStreet = "road"
const labelStreetNumber = "house_number"
const labelUnit = "unit"

// LibPostalResolver a location resolver that uses libpostal.
type LibPostalResolver struct {
}

func (r LibPostalResolver) Name() string {
	return LibPostalResolverName
}

func (r LibPostalResolver) Resolve(input string, options ResolvingOptions) (*ResolvingResult, error) {
	// Build the options.
	parserOptions := parser.ParserOptions{Country: options.Country}

	// Normalise the address.
	normalised := expand.ExpandAddress(input)[0]

	// Resolve the address.
	resolvedFields := make(map[string]string)
	resolved := parser.ParseAddressOptions(normalised, parserOptions)
	for _, entry := range resolved {
		resolvedFields[entry.Label] = entry.Value
	}

	// Override the country field if this was passed.
	if len(parserOptions.Country) > 0 {
		resolvedFields[labelCountry] = parserOptions.Country
	}

	// Format the address as a location.
	return &ResolvingResult{
		AddressInfo: location.AddressInfo{
			City:            resolvedFields[labelCity],
			Country:         resolvedFields[labelCountry],
			PostalCode:      resolvedFields[labelPostalCode],
			StateOrProvince: resolvedFields[labelStateOrProvince],
			Street:          resolvedFields[labelStreet],
			StreetNumber:    resolvedFields[labelStreetNumber],
			Unit:            resolvedFields[labelUnit],
		},
		DisplayAddress: normalised,
	}, nil
}
