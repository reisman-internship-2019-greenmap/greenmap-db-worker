package main

import (
	"context"
	"log"

	//"gopkg.in/mgo.v2/bson"

	"github.com/eliotscott/greenmap-db-worker/cmd/dap"
)

func main() {
	log.Println("spinning up greenmap-db-worker...")
	ctx := context.Background()

	dap, err := dap.NewDap(ctx, "listingsAndReviews")
	if err != nil {
		log.Fatalf("connection to collection: failed:: [%v]", err)
	}
	log.Printf("connected to collection [%v]", dap.Conn.Name())
}
