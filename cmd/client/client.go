package main

import (
	"context"
	"fmt"
	"io"
	"log"
	"time"

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
	// AddUserVerbose(client)
	AddUsers(client)
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

func AddUsers(client pb.UserServiceClient) {
	requests := []*pb.User{
		{
			Id:    "a1",
			Name:  "John Doe",
			Email: "johndoe@mail.com",
		},
		{
			Id:    "a2",
			Name:  "Jane Doe",
			Email: "janedoe@mail.com",
		},
		{
			Id:    "a3",
			Name:  "John Due",
			Email: "johndue@mail.com",
		},
		{
			Id:    "a4",
			Name:  "Jane Due",
			Email: "janedue@mail.com",
		},
		{
			Id:    "a5",
			Name:  "Jenny Doe",
			Email: "jennydoe@mail.com",
		},
	}

	stream, error := client.AddUsers(context.Background())

	if error != nil {
		log.Fatalf("Error creating request: %v", error)
	}

	for _, request := range requests {
		stream.Send(request)
		time.Sleep(time.Second * 3)
		fmt.Println("Sending request: ", request)
	}

	response, error := stream.CloseAndRecv()

	if error != nil {
		log.Fatalf("Error receiving response: %v", error)
	}

	fmt.Println(response)
}
