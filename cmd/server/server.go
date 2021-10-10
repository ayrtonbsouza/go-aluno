package main

import (
	"log"
	"net"

	"github.com/codeedu/fc2-grpc/pb"
	"github.com/codeedu/fc2-grpc/services"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func main() {
	listener, error := net.Listen("tcp", "localhost:50051")

	if error != nil {
		log.Fatalf("Could not connect: %v", error)
	}

	grpcServer := grpc.NewServer()
	pb.RegisterUserServiceServer(grpcServer, services.NewUserService())
	reflection.Register(grpcServer)

	if error := grpcServer.Serve(listener); error != nil {
		log.Fatalf("Could not serve: %v", error)
	}
}
