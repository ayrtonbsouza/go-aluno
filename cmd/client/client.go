package main

import (
	"context"
	"fmt"
	"log"

	"github.com/codeedu/fc2-grpc/pb"
	"google.golang.org/grpc"
)

func main() {
	connection, error := grpc.Dial("localhost:50051", grpc.WithInsecure())

	if error != nil {
		log.Fatalf("Could not connect to the gRPC Server: %v", error)
	}

	defer connection.Close()

	client := pb.NewUserServiceClient(connection)

	AddUser(client)
}

func AddUser(client pb.UserServiceClient) {
	request := &pb.User{
		Id:    "0",
		Name:  "John Doe",
		Email: "johndoe@mail.com",
	}

	response, error := client.AddUser(context.Background(), request)

	if error != nil {
		log.Fatalf("Could not make gRPC request: %v", error)
	}

	fmt.Println(response)
}
