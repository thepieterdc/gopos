package database

import (
	"context"
	"fmt"
	log "github.com/sirupsen/logrus"
	"github.com/thepieterdc/gopos/internal/pkg/configuration"
	"github.com/thepieterdc/gopos/internal/pkg/logging"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"net/url"
	"time"
)

// Connect initialises a database connection.
func Connect(config *configuration.Configuration) (*Database, error) {
	// Initialise the logging fields.
	logger := log.WithFields(logging.BootStage()).WithFields(logging.DatabaseComponent())

	// Validate whether the configuration is valid.
	if len(config.MongoUri()) == 0 {
		logger.Info("No connection string was configured. Skipping.")
		return nil, nil
	}

	logger.Info("Attempting to connect.")

	// Create a client.
	client, err := mongo.NewClient(options.Client().ApplyURI(config.MongoUri()))
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
	database, err := url.Parse(config.MongoUri())
	if err != nil {
		return nil, err
	}
	databaseName := database.Path[1:]
	if len(databaseName) == 0 {
		return nil, fmt.Errorf("could not extract database name: %s", database)
	}

	logger.Info("Connected successfully.")

	// Build a new database instance.
	return &Database{client: client, db: client.Database(databaseName)}, nil
}
