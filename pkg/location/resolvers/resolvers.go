package resolvers

import (
	"fmt"
	"github.com/thepieterdc/gopos/pkg/location"
)

// ResolvingOptions optional flags that aid in resolving.
type ResolvingOptions struct {
	Country string
}

type ResolvingResult struct {
	AddressInfo    location.AddressInfo
	DisplayAddress string
}

// LocationResolver resolves an input query to a location.
type LocationResolver interface {
	Name() string
	Resolve(input string, options ResolvingOptions) (*ResolvingResult, error)
}

// GetByName gets the location resolver with the given name.
func GetByName(name string) (*LocationResolver, error) {
	// Build the list of resolvers.
	allResolvers := []LocationResolver{
		&GoogleResolver{},
		&LibPostalResolver{},
	}

	// Find the requested resolver.
	for _, resolver := range allResolvers {
		if resolver.Name() == name {
			return &resolver, nil
		}
	}

	// Resolver was not found.
	return nil, fmt.Errorf("resolver with name %s was not found", name)
}
