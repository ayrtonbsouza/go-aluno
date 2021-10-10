package main

import (
	"log"
	"net"

	"google.golang.org/grpc"
)

func main() {
	listener, error := net.Listen("tcp", "localhost:50051")

	if error != nil {
		log.Fatalf("Could not connect: %v", error)
	}

	grpcServer := grpc.NewServer()

	if error := grpcServer.Serve(listener); error != nil {
		log.Fatalf("Could not serve: %v", error)
	}
}
