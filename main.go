package main

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"github.com/go-playground/validator"
	"github.com/julienschmidt/httprouter"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/thepieterdc/gopos/pkg/request"

	//postal "github.com/openvenues/gopostal/parser"
	"github.com/thepieterdc/gopos/cmd"
	"github.com/thepieterdc/gopos/src"
	"go.mongodb.org/mongo-driver/bson"
	"net/url"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"net/http"
	"os"
	"time"
)

// Environment variables.
var (
	GOOGLE_API_KEY = os.Getenv("GOOGLE_API_KEY")
	MONGO_URI      = os.Getenv("MONGO_URI")
)

// Database connection.
var database *mongo.Database

// InternalError something went wrong on the server.
type InternalError struct {
	Error   string `json:"error"`
	Message string `json:"message"`
}

// InvalidArgumentError an invalid (or missing) argument value was found.
type InvalidArgumentError struct {
	Error    string `json:"error"`
	Argument string `json:"argument"`
}

/**
 * Initialises the database connection.
 */
func initialiseDatabase() *mongo.Database {
	// Create a client.
	client, err := mongo.NewClient(options.Client().ApplyURI(MONGO_URI))
	if err != nil {
		log.Fatalf("could not instantiate Mongo client: %v\n", err)
	}

	// Connect to the database.
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	err = client.Connect(ctx)
	if err != nil {
		log.Fatalf("could not connect to database: %v\n", err)
	}

	// Test the connection.
	err = client.Ping(ctx, nil)
	if err != nil {
		log.Fatalf("could not ping database: %v\n", err)
	}

	// Fetch the database name.
	database, err := url.Parse(MONGO_URI)
	if err != nil {
		log.Fatalf("could not extract database name: %v\n", err)
	}
	databaseName := database.Path[1:]
	if len(databaseName) == 0 {
		log.Fatalf("could not extract database name: %s\n", database)
	}

	return client.Database(databaseName)
}

/**
 * Handles internal errors.
 */
func internalError(w http.ResponseWriter, error string) {
	log.Println(fmt.Errorf(error))
	jsonResponse(w, http.StatusInternalServerError, InternalError{
		Error:   "internal_error",
		Message: error,
	})
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
		log.Println(fmt.Errorf("error sending response: %v", err))
	}
}

///**
// * Handles the /format route.
// */
//func formatHandler(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
//	// Extract the query string arguments.
//	queryString := r.URL.Query()
//	input := queryString.Get("input")
//	if len(input) == 0 {
//		invalidArgument(w, "input")
//		return
//	}
//
//	// Format the address.
//	response := make(map[string]interface{})
//	for _, entry := range postal.ParseAddress(input) {
//		response[entry.Label] = entry.Value
//	}
//
//	// Send the response.
//	jsonResponse(w, http.StatusOK, response)
//}

/**
 * Handles the /google/place/:id route.
 */
func googlePlaceHandler(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	// Extract the requestUrl parameters.
	id := params.ByName("id")
	if len(id) == 0 {
		invalidArgument(w, "id")
		return
	}

	// Try to fetch the place id from the database.
	collection := database.Collection(src.GooglePlaceIdCollection)
	var placeDetails src.GooglePlaceDetails
	err := collection.FindOne(r.Context(), bson.M{"place_id": id}).Decode(&placeDetails)
	if err != nil && err != mongo.ErrNoDocuments {
		internalError(w, fmt.Sprintf("%v", err))
		return
	}

	// If the place details exist, send them to the client.
	if err == nil {
		jsonResponse(w, http.StatusOK, placeDetails)
		return
	}

	// Prepare the request for the Google API.
	requestUrl := fmt.Sprintf(src.GoogleApiUrl, GOOGLE_API_KEY, id)
	req, err := http.NewRequest("GET", requestUrl, bytes.NewBuffer(nil))
	log.Printf("Querying Google API for place id: %s.\n", id)
	if err != nil {
		internalError(w, fmt.Sprintf("%v", err))
		return
	}

	// Send the request.
	req.Header.Set("Content-Type", "application/json")
	client := &http.Client{}
	response, err := client.Do(req)
	if err != nil {
		internalError(w, fmt.Sprintf("%v", err))
		return
	}

	// Parse the response.
	defer response.Body.Close()
	if response.StatusCode == http.StatusOK {
		// Read the response from Google.
		var parsedResponse src.GooglePlaceDetailsResponse
		err = json.NewDecoder(response.Body).Decode(&parsedResponse)
		if err != nil {
			internalError(w, fmt.Sprintf("%v", err))
			return
		}

		// Send the response.
		jsonResponse(w, http.StatusOK, parsedResponse.Result)

		// Save the response in the database.
		_, err = collection.InsertOne(r.Context(), parsedResponse.Result)
		if err != nil {
			log.Println(fmt.Errorf(fmt.Sprintf("%v", err)))
		}
	} else {
		internalError(w, fmt.Sprintf("status code is %d.", response.StatusCode))
	}
}

func main() {
	//// Validate the settings.
	//if len(GOOGLE_API_KEY) == 0 {
	//	log.Fatal("GOOGLE_API_KEY is missing.")
	//}
	//if len(MONGO_URI) == 0 {
	//	log.Fatal("MONGO_URI is missing.")
	//}
	//
	//// Connect to the database.
	//log.Println("Connecting to the database...")
	//database = initialiseDatabase()
	//log.Println("Connection OK.")

	// Build the router and register all the routes.
	//router := httprouter.New()
	//router.GET("/format", formatHandler)
	//router.GET("/google/place/:id", googlePlaceHandler)

	// Build the webserver and register all the routes.
	srv := echo.New()
	srv.GET("/timezone", cmd.TimezoneHandler)

	// Register middleware.
	srv.Use(middleware.Logger())

	// Register data validator.
	srv.Validator = &request.Validator{Validator: validator.New()}

	// Start the server.
	srv.Logger.Fatal(srv.Start("0.0.0.0:8000"))
}
