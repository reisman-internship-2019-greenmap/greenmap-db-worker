package main

import (
	"context"
	"log"

	"gopkg.in/mgo.v2/bson"

	"github.com/eliotscott/greenmap-db-worker/cmd/dap"
)

// Example of json to bson marshaling
// Note, there are many more fields in the listingsAndReview data sample
type Listing struct {
	ID         bson.ObjectId `bson:"_id,omitempty"`
	ListingURL string        `bson:"listing_url"`
}

func main() {
	log.Println("spinning up greenmap-db-worker...")
	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	dap, err := dap.NewDap(ctx, "listingsAndReviews")
	if err != nil {
		log.Fatalf("connection to collection: failed:: [%v]", err)
	}
	log.Printf("connected to collection [%v]", dap.Conn.Name())
}
