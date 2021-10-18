package database

import (
	"context"
	"github.com/thepieterdc/gopos/pkg/google"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

// Google Place details database collection.
const googlePlaceDetailsCollection = "google_place_details"

// FindPlaceDetailsById queries the database to find the given Google Place
// details.
func (db Database) FindPlaceDetailsById(id string) (*google.GooglePlaceDetails, error) {
	// Try to fetch the place id from the database.
	collection := db.db.Collection(googlePlaceDetailsCollection)
	var placeDetails google.GooglePlaceDetails
	err := collection.FindOne(context.Background(), bson.M{"place_id": id}).Decode(&placeDetails)
	if err != nil {
		// Something went wrong.
		if err != mongo.ErrNoDocuments {
			return nil, err
		}

		// The document was not found.
		return nil, nil
	}

	// Return the document.
	return &placeDetails, nil
}

// SavePlaceDetails saves the given place details into the database.
func (db Database) SavePlaceDetails(details *google.GooglePlaceDetails) error {
	// Insert the details into the database.
	collection := db.db.Collection(googlePlaceDetailsCollection)
	_, err := collection.InsertOne(context.Background(), details)
	return err
}
