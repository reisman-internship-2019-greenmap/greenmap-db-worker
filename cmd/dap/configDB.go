package dap

import (
	"context"
	"fmt"
	"os"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// DAP (data access portal) exportable member
type DAP struct {
	DAP *mongo.Collection
}

// setEnv sets environment variables
// TODO: deprecate before shipping!
func setEnv() {
	os.Setenv("MONGO_HOST", "mongodb+srv://foo:bar@cluster0-zabfy.gcp.mongodb.net/test?retryWrites=true&w=majority")
	os.Setenv("MONGO_DATABASE", "greenmap-products")
}

// NewDap creates a mongoDB collection DAP
func NewDap(ctx context.Context, collectionName string) (*DAP, error) {
	setEnv() // TODO: deprecate before shipping!
	d := &DAP{}
	uri := os.Getenv("MONGO_HOST")

	// Create client
	client, err := mongo.NewClient(options.Client().ApplyURI(uri))
	if err != nil {
		return nil, fmt.Errorf("couldn't connect to mongo: %v", err)
	}

	// Connect to database
	err = client.Connect(ctx)
	if err != nil {
		return nil, fmt.Errorf("mongo client couldn't connect with background context: %v", err)
	}

	// Connect to collection
	d.DAP = client.Database(os.Getenv("MONGO_DATABASE")).Collection(collectionName)

	return d, nil
}
