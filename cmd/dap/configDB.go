package dap

import (
	"context"
	"fmt"
	"os"

	"gopkg.in/mgo.v2/bson"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// Config exports a configured data access portal
type DAP struct {
	Conn *mongo.Collection
}

type key string

const (
	hostKey     = key("hostKey")
	usernameKey = key("usernameKey")
	passwordKey = key("passwordKey")
	databaseKey = key("databaseKey")
)

// Example of json to bson marshaling
// Note, there are many more fields in the listingsAndReview data sample
type Listing struct {
	ID         bson.ObjectId `bson:"_id,omitempty"`
	ListingURL string        `bson:"listing_url"`
}

// setEnv sets environment variables
// TODO: deprecate before shipping!
func setEnv() {
	os.Setenv("MONGO_HOST", "mongodb+srv://foo:bar@cluster0-zabfy.gcp.mongodb.net/test?retryWrites=true&w=majority")
	os.Setenv("MONGO_DATABASE", "sample_airbnb")
}

// NewDap creates a mongoDB collection DAP
func NewDap(ctx context.Context, collectionName string) (*DAP, error) {
	setEnv() // TODO: deprecate before shipping!
	dap := &DAP{}
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
	dap.Conn = client.Database(os.Getenv("MONGO_DATABASE")).Collection(collectionName)

	return dap, nil
}
