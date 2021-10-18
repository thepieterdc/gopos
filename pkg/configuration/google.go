package configuration

// GoogleApiBase gets the base domain for contacting the Google Places API.
func (c Configuration) GoogleApiBase() string {
	return c.config.GetString(varGoogleApiBase)
}

// GoogleApiKey gets the API key for contacting the Google Places API.
func (c Configuration) GoogleApiKey() string {
	return c.config.GetString(varGoogleApiKey)
}
