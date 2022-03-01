// Package main implements a client for Greeter service.
package main

import (
	"context"
	"flag"
	"log"
	"time"

	"google.golang.org/grpc"
	codec "google.golang.org/grpc/examples/bsonCodec/codec"
	"go.mongodb.org/mongo-driver/bson"
)

const (
	defaultName = "world"
)

var (
	addr = flag.String("addr", "localhost:50051", "the address to connect to")
	name = flag.String("name", defaultName, "Name to greet")
)

func main() {
	flag.Parse()
	// Set up a connection to the server.
	conn, err := grpc.Dial(*addr, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()

	// Contact the server and print out its response.
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	var out bson.D
	in := bson.D{{"hello", "world"}}
	err = conn.Invoke(ctx, "/BSONCodec.Test/Send", in, &out, grpc.ForceCodec(codec.BSONCodec{}))
	if err != nil {
		log.Fatalf("could not send: %v", err)
	}
	log.Printf("Response: %s", out)
}
