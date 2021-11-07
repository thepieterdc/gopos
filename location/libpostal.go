package location

import (
	postal "github.com/openvenues/gopostal/parser"
)

// Representation of the lib postal model, conforms to the LocationFromSource interface
type LibPostal struct {
	Country  string
	Language string
}

// Parse returns the output of libpostal parsed to the internal Location model
func (*LibPostal) Parse(query interface{}) (Location, error) {
	// Parse the input address.
	parsed := postal.ParseAddressOptions(query, {Country: LibPostal.Country, Language: LibPostal.Language})

	// Convert to the location struct

	return Location{}, nil
}

// Confidence parses the input string and returns the confidence of the result
func (*LibPostal) Confidence() (float32, error) {

	return 1.0, nil
}
