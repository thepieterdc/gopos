package database

import (
	"context"
	"fmt"
	"github.com/thepieterdc/gopos/pkg/environment"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"net/url"
	"time"
)

// Connect initialises a database connection.
func Connect() (*Database, error) {
	// Create a client.
	client, err := mongo.NewClient(options.Client().ApplyURI(environment.MongoUri))
	if err != nil {
		return nil, err
	}

	// Connect to the database.
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	err = client.Connect(ctx)
	if err != nil {
		return nil, err
	}

	// Test the connection.
	err = client.Ping(ctx, nil)
	if err != nil {
		return nil, err
	}

	// Fetch the database name.
	database, err := url.Parse(environment.MongoUri)
	if err != nil {
		return nil, err
	}
	databaseName := database.Path[1:]
	if len(databaseName) == 0 {
		return nil, fmt.Errorf("could not extract database name: %s", database)
	}

	return &Database{client.Database(databaseName)}, nil
}
