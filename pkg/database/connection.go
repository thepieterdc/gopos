package database

import (
	"context"
	"fmt"
	"github.com/thepieterdc/gopos/pkg/configuration"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"net/url"
	"time"
)

// Connect initialises a database connection.
func Connect(config *configuration.Configuration) (*Database, error) {
	// Validate whether the configuration is valid.
	if len(config.MongoUri()) == 0 {
		return nil, nil
	}

	log.Println("[DB] Connecting to the database.")

	// Create a client.
	client, err := mongo.NewClient(options.Client().ApplyURI(config.MongoUri()))
	if err != nil {
		return nil, fmt.Errorf("[DB] %w", err)
	}

	// Connect to the database.
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	err = client.Connect(ctx)
	if err != nil {
		return nil, fmt.Errorf("[DB] %w", err)
	}

	// Test the connection.
	err = client.Ping(ctx, nil)
	if err != nil {
		return nil, fmt.Errorf("[DB] %w", err)
	}

	// Fetch the database name.
	database, err := url.Parse(config.MongoUri())
	if err != nil {
		return nil, fmt.Errorf("[DB] %w", err)
	}
	databaseName := database.Path[1:]
	if len(databaseName) == 0 {
		return nil, fmt.Errorf("[DB] could not extract database name: %s", database)
	}

	log.Println("[DB] Connected to the database.")

	// Build a new database instance.
	return &Database{client: client, db: client.Database(databaseName)}, nil
}
