package server

import (
	_ "fmt"
	"context"

	"github.com/oinume/grpc-sample/proto-gen/go/proto/api/v1"
)

type UsersServer struct {}

func (s *UsersServer) GetUser(ctx context.Context, in *api_v1.GetUserRequest) (*api_v1.GetUserResponse, error) {
	return &api_v1.GetUserResponse{
		Id: in.GetId(),
		Name: "hello",
	}, nil
}

func (s *UsersServer) CreateUser(ctx context.Context, in *api_v1.CreateUserRequest) (*api_v1.CreateUserResponse, error) {
	return &api_v1.CreateUserResponse{
		Id: "12345",
		Name: in.GetName(),
	}, nil
}

func (s *UsersServer) UpdateUser(ctx context.Context, in *api_v1.UpdateUserRequest) (*api_v1.UpdateUserResponse, error) {
	return &api_v1.UpdateUserResponse{}, nil
}

func (s *UsersServer) DeleteUser(ctx context.Context, in *api_v1.DeleteUserRequest) (*api_v1.DeleteUserResponse, error) {
	return &api_v1.DeleteUserResponse{}, nil
}
