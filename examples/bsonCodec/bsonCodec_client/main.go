// Package main implements a client for Greeter service.
package main

import (
	"context"
	"flag"
	"io"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"google.golang.org/grpc"
	codec "google.golang.org/grpc/examples/bsonCodec/codec"
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

	stream, err := conn.NewStream(ctx, &grpc.StreamDesc{
		StreamName:    "Send",
		ServerStreams: true,
	}, "/BSONCodec.Test/Send", grpc.ForceCodec(codec.BSONCodec{}))
	if err != nil {
		log.Fatalf("could not create stream: %v", err)
	}
	log.Printf("Created client stream %v\n", stream)
	in := bson.D{{"hello", "world"}}
	if err := stream.SendMsg(&in); err != nil {
		log.Fatalf("could not SendMsg: %v", err)
	}
	if err := stream.CloseSend(); err != nil {
		log.Fatalf("could not CloseSend: %v", err)
	}
	log.Printf("Sent message %v\n", in)
	for {
		var out bson.D
		err := stream.RecvMsg(&out)
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatalf("could not RecvMsg: %v", err)
		}
		log.Printf("Response: %s", out)
	}
}
