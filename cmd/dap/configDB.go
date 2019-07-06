package dap

import (
	"context"
	"fmt"
	"os"

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

// setEnv sets environment variables
// TODO: deprecate before shipping!
func setEnv() {
	os.Setenv("MONGO_HOST", "mongodb+srv://cluster0-zabfy.gcp.mongodb.net/test")
	os.Setenv("MONGO_USERNAME", "foo")
	os.Setenv("MONGO_PASSWORD", "bar")
	os.Setenv("MONGO_DATABASE", "sample_airbnb")
}

// NewDap creates a mongoDB collection DAP
func NewDap(ctx context.Context, collectionName string) (*DAP, error) {
	setEnv() // TODO: deprecate before shipping!
	dap := &DAP{}

	ctx = context.WithValue(ctx, hostKey, os.Getenv("MONGO_HOST"))
	ctx = context.WithValue(ctx, usernameKey, os.Getenv("MONGO_USERNAME"))
	ctx = context.WithValue(ctx, passwordKey, os.Getenv("MONGO_PASSWORD"))
	ctx = context.WithValue(ctx, databaseKey, os.Getenv("MONGO_DATABASE"))

	uri := fmt.Sprintf(`mongodb://%s:%s@%s/%s`,
		ctx.Value(usernameKey).(string),
		ctx.Value(passwordKey).(string),
		ctx.Value(hostKey).(string),
		ctx.Value(databaseKey).(string),
	)

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
