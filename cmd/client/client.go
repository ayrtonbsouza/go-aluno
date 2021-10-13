package main

import (
	"context"
	"fmt"
	"io"
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

	// AddUser(client)
	AddUserVerbose(client)
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

func AddUserVerbose(client pb.UserServiceClient) {
	request := &pb.User{
		Id:    "0",
		Name:  "John Doe",
		Email: "johndoe@mail.com",
	}

	responseStream, error := client.AddUserVerbose(context.Background(), request)

	if error != nil {
		log.Fatalf("Could not make gRPC request: %v", error)
	}

	for {
		stream, error := responseStream.Recv()
		if error == io.EOF {
			break
		}
		if error != nil {
			log.Fatalf("Could not receive the message: %v", error)
		}
		fmt.Println("Status:", stream.Status, " - ", stream.GetUser())
	}
}
