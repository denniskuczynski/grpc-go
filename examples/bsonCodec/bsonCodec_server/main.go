// Package main implements a server for Greeter service.
package main

import (
	"flag"
	"fmt"
	"log"
	"net"

	"go.mongodb.org/mongo-driver/bson"
	"google.golang.org/grpc"
	codec "google.golang.org/grpc/examples/bsonCodec/codec"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/peer"
)

var (
	port = flag.Int("port", 50051, "The server port")
)

func _BSON_TEST_Handler(srv interface{}, stream grpc.ServerStream) error {
	peer, ok := peer.FromContext(stream.Context())
	if ok {
		fmt.Printf("Peer Addr: %v AuthInfo: %v\n", peer.Addr, peer.AuthInfo)
	}
	md, ok := metadata.FromIncomingContext(stream.Context())
	if ok {
		fmt.Printf("MetaData: %v\n", md)
	}

	m := new(bson.D)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}

	// Send once
	fmt.Printf("Sending %v\n", m)
	if err := stream.SendMsg(m); err != nil {
		fmt.Printf("send error %v\n", err)
	}

	// Send twice
	fmt.Printf("Sending again %v\n", m)
	if err := stream.SendMsg(m); err != nil {
		fmt.Printf("send error %v\n", err)
	}

	return nil
}

func main() {
	flag.Parse()
	lis, err := net.Listen("tcp", fmt.Sprintf("localhost:%d", *port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer(grpc.ForceServerCodec(codec.BSONCodec{}))

	s.RegisterService(&grpc.ServiceDesc{
		ServiceName: "BSONCodec.Test",
		HandlerType: nil,
		Methods:     []grpc.MethodDesc{},
		Streams: []grpc.StreamDesc{
			{
				StreamName:    "Send",
				Handler:       _BSON_TEST_Handler,
				ServerStreams: true,
			},
		},
	}, nil)

	log.Printf("server listening at %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
