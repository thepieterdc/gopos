package main

import (
	"encoding/json"
	"github.com/bradfitz/latlong"
	"log"
	"net/http"
	"strconv"
)

// InvalidArgumentError an invalid (or missing) argument value was found.
type InvalidArgumentError struct {
	Error    string `json:"error"`
	Argument string `json:"argument"`
}

// TimezoneResponse response of the /timezone call.
type TimezoneResponse struct {
	Latitude  float64 `json:"latitude"`
	Longitude float64 `json:"longitude"`
	Timezone  string  `json:"timezone"`
}

/**
 * Handles invalid arguments.
 */
func invalidArgument(w http.ResponseWriter, argument string) {
	jsonResponse(w, http.StatusBadRequest, InvalidArgumentError{
		Error:    "invalid_argument",
		Argument: argument,
	})
}

func jsonResponse(w http.ResponseWriter, statusCode int, body interface{}) {
	// Convert the body to JSON.
	jsonBody, err := json.Marshal(body)
	if err != nil {
		log.Fatal(err)
		return
	}

	// Send the response.
	w.WriteHeader(statusCode)
	w.Header().Set("Content-Type", "application/json")
	_, err = w.Write(jsonBody)
	if err != nil {
		log.Fatal(err)
	}
}

/**
 * Handles the /timezone route.
 */
func timezoneHandler(w http.ResponseWriter, r *http.Request) {
	// Extract the query string arguments.
	queryString := r.URL.Query()
	latitude, err := strconv.ParseFloat(queryString.Get("latitude"), 64)
	if err != nil {
		invalidArgument(w, "latitude")
		return
	}

	longitude, err := strconv.ParseFloat(queryString.Get("longitude"), 64)
	if err != nil {
		invalidArgument(w, "longitude")
		return
	}

	// Find the address for this coordinate pair.
	timezone := latlong.LookupZoneName(latitude, longitude)

	// Send the response.
	jsonResponse(w, http.StatusOK, TimezoneResponse{
		Latitude:  latitude,
		Longitude: longitude,
		Timezone:  timezone,
	})
}

func main() {
	// Register the timezone route.
	http.HandleFunc("/timezone", timezoneHandler)

	// Start the server.
	log.Fatal(http.ListenAndServe("0.0.0.0:8000", nil))
}
