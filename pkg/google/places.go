package google

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/thepieterdc/gopos/pkg/configuration"
	"log"
	"net/http"
)

// Base Google Place url.
const apiUrl = "%s/maps/api/place/details/json?key=%s&place_id=%s&fields=address_component,adr_address,business_status,formatted_address,geometry,icon,name,photo,place_id,plus_code,type,url,utc_offset,vicinity"

// Get the configuration.
var config = configuration.Configure()

// GetPlaceDetailsById sends a lookup request to the Google Maps API with the
// given place ID.
func GetPlaceDetailsById(id string) (*GooglePlaceDetails, error) {
	// Prepare the request for the Google API.
	requestUrl := fmt.Sprintf(apiUrl, config.GoogleApiBase(), config.GoogleApiKey(), id)
	request, err := http.NewRequest("GET", requestUrl, bytes.NewBuffer(nil))
	if err != nil {
		return nil, err
	}

	// Send the request.
	log.Printf("Querying Google API for place id: %s.\n", id)
	client := &http.Client{}
	response, err := client.Do(request)
	if err != nil {
		return nil, err
	}

	// Parse the response.
	defer response.Body.Close()
	if response.StatusCode == http.StatusOK {
		// Read the response.
		var parsedResponse GooglePlaceDetailsResponse
		err = json.NewDecoder(response.Body).Decode(&parsedResponse)
		if err != nil {
			return nil, err
		}

		return &parsedResponse.Result, nil
	}

	return nil, fmt.Errorf("received status code: %d", response.StatusCode)
}
