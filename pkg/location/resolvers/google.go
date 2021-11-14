package resolvers

import (
	"errors"
)

// GoogleResolverName name of the resolver.
const GoogleResolverName = "google"

// GoogleResolver a location resolver that uses the Google Maps API.
type GoogleResolver struct {
}

func (r GoogleResolver) Name() string {
	return GoogleResolverName
}

func (r GoogleResolver) Resolve(input string, options ResolvingOptions) (*ResolvingResult, error) {
	return nil, errors.New("not implemented")
}
