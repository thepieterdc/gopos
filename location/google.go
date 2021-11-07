package location

// Base Google Place url.
const apiUrl = "%s/maps/api/place/details/json?key=%s&place_id=%s&fields=address_component,adr_address,business_status,formatted_address,geometry,icon,name,photo,place_id,plus_code,type,url,utc_offset,vicinity"

// Representation of the lib postal model, conforms to the LocationFromSource interface
type GooglePlaces struct {
	apiKey string
}

// Parse returns the output of the google places call parsed to the internal Location model
func (*GooglePlaces) Parse(id interface{}) (Location, error) {
	// Make the call to fetchDetails() and convert to Location

	return Location{}, nil
}

// Confidence parses the input string and returns the confidence of the result
func (*GooglePlaces) Confidence() (float32, error) {
	return 1.0, nil
}

func (*GooglePlaces) fetchDetails() (GooglePlaceDetailsResponse, error) {
	return GooglePlaceDetailsResponse{}, nil
}
