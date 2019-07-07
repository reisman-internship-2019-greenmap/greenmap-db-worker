package main

import (
	"context"
	"log"

	"gopkg.in/mgo.v2/bson"

	"github.com/eliotscott/greenmap-db-worker/cmd/dap"
)

// Product struct with marshalling to bson.D type for mongodb
type Product struct {
	ID           bson.ObjectId `bson:"_id,omitempty"`
	Name         string        `bson:"listing_url"`
	Manufacturer string        `bson:"manufacturer"`
	Category     string        `bson:"category"`
	ESGScore     string        `bson:"esg_score"`
}

func main() {
	log.Println("spinning up greenmap-db-worker...")
	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	// connect to staging DAP
	staging, err := dap.NewDap(ctx, "staging")
	if err != nil {
		log.Fatalf("connection to collection: failed:: [%v]", err)
	}
	log.Printf("connected to collection [%v]", staging.DAP.Name())

	// connect to production DAP
	production, err := dap.NewDap(ctx, "production")
	if err != nil {
		log.Fatalf("connection to collection: failed:: [%v]", err)
	}
	log.Printf("connected to collection [%v]", production.DAP.Name())

	// go routine to parse database goes here
}
