package services

import (
	"context"
	"fmt"
	"io"
	"log"
	"time"

	"github.com/codeedu/fc2-grpc/pb"
)

type UserService struct {
	pb.UnimplementedUserServiceServer
}

func NewUserService() *UserService {
	return &UserService{}
}

func (*UserService) AddUser(context context.Context, request *pb.User) (*pb.User, error) {
	// Insert in pseudodatabase
	fmt.Println(request.Name)

	return &pb.User{
		Id:    "1",
		Name:  request.GetName(),
		Email: request.GetEmail(),
	}, nil
}

func (*UserService) AddUserVerbose(request *pb.User, stream pb.UserService_AddUserVerboseServer) error {
	stream.Send(&pb.UserResultStream{
		Status: "Init",
		User:   &pb.User{},
	})
	time.Sleep(time.Second * 3)

	stream.Send(&pb.UserResultStream{
		Status: "Inserting",
		User:   &pb.User{},
	})
	time.Sleep(time.Second * 3)

	stream.Send(&pb.UserResultStream{
		Status: "User has been inserted",
		User: &pb.User{
			Id:    "123",
			Name:  request.GetName(),
			Email: request.GetEmail(),
		},
	})
	time.Sleep(time.Second * 3)

	stream.Send(&pb.UserResultStream{
		Status: "Completed",
		User: &pb.User{
			Id:    "123",
			Name:  request.GetName(),
			Email: request.GetEmail(),
		},
	})
	time.Sleep(time.Second * 3)

	return nil
}

func (*UserService) AddUsers(stream pb.UserService_AddUsersServer) error {
	users := []*pb.User{}

	for {
		request, error := stream.Recv()
		if error == io.EOF {
			return stream.SendAndClose(&pb.Users{
				User: users,
			})
		}

		if error != nil {
			log.Fatalf("Error receiving stream: %v", error)
		}

		users = append(users, &pb.User{
			Id:    request.GetId(),
			Name:  request.GetName(),
			Email: request.GetEmail(),
		})
		fmt.Println("Adding", request.GetName())
	}
}
