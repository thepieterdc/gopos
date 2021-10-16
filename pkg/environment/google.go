package environment

import "os"

// GoogleApiKey API key to use for contacting the Google Places API.
var GoogleApiKey = os.Getenv("GOOGLE_API_KEY")
