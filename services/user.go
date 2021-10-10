package services

import (
	"context"
	"fmt"

	"github.com/codeedu/fc2-grpc/pb"
)

// type UserServiceServer interface {
// 	AddUser(context.Context, *User) (*User, error)
// 	mustEmbedUnimplementedUserServiceServer()
// }

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
