// Package main implements a server for Greeter service.
package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"net"

	"google.golang.org/grpc"
	codec "google.golang.org/grpc/examples/bsonCodec/codec"
	"go.mongodb.org/mongo-driver/bson"
)

var (
	port = flag.Int("port", 50051, "The server port")
)

// Handler of: type methodHandler func(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor UnaryServerInterceptor) (interface{}, error)
// dec unmarshals data with codec and handles internal stats and logging
func _BSON_TEST_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(bson.D)
	if err := dec(in); err != nil {
		return nil, err
	}
	// No interceptor is defined so we will always enter this block
	if interceptor == nil {
		log.Printf("Received no interceptor: %v", in)
		return in, nil
	}

	// If we were to use interceptors...
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/BSONCodec.Test/Send",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		log.Printf("Received via interceptor: %v", in)
		return in, nil
	}
	return interceptor(ctx, in, info, handler)
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
		Methods: []grpc.MethodDesc{
			{
				MethodName: "Send",
				Handler:    _BSON_TEST_Handler,
			},
		},
		Streams:  []grpc.StreamDesc{},
	}, nil)

	log.Printf("server listening at %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
