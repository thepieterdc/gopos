package database

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo"
)

// Database wrapper around the database to hide internal details.
type Database struct {
	client *mongo.Client
	db     *mongo.Database
}

// Disconnect closes the connection.
func (db Database) Disconnect() error {
	return db.client.Disconnect(context.Background())
}
