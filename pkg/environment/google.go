package environment

import "os"

// googleApiBase Base domain to use for contacting the Google Places API. This
// can be changed to allow mocking while testing.
var googleApiBase = os.Getenv("GOOGLE_API_BASE")

// GoogleApiBase getter for the similar named environment variable.
func GoogleApiBase() string {
	if len(googleApiBase) > 0 {
		return googleApiBase
	}

	// Default value.
	return "https://maps.googleapis.com"
}

// GoogleApiKey API key to use for contacting the Google Places API.
var GoogleApiKey = os.Getenv("GOOGLE_API_KEY")
