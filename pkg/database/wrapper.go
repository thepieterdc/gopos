package database

import "go.mongodb.org/mongo-driver/mongo"

// Database wrapper around the database to hide internal details.
type Database struct {
	*mongo.Database
}
