package server

import (
	"context"

	ptypes_empty "github.com/golang/protobuf/ptypes/empty"
	api_v1 "github.com/oinume/grpc-sample/proto-gen/go/proto/api/v1"
)

type UsersServer struct{}

func (s *UsersServer) ListUsers(ctx context.Context, in *api_v1.ListUsersRequest) (*api_v1.ListUsersResponse, error) {
	return &api_v1.ListUsersResponse{
		Users: []*api_v1.User{
			{Id: "1", Name: "oinume", RealName: "kazuhiro oinuma"},
			{Id: "2", Name: "akuwano", RealName: "Akihiro Kuwano"},
			{Id: "3", Name: "oranie", RealName: "Takashi Narita"},
		},
	}, nil
}

func (s *UsersServer) GetUser(ctx context.Context, in *api_v1.GetUserRequest) (*api_v1.User, error) {
	return &api_v1.User{
		Id:       in.GetId(),
		Name:     "oinume",
		RealName: "Kazuhiro Oinuma",
	}, nil
}

func (s *UsersServer) CreateUser(ctx context.Context, in *api_v1.CreateUserRequest) (*api_v1.User, error) {
	return &api_v1.User{
		Id:       "12345",
		Name:     in.GetName(),
		RealName: in.GetRealName(),
	}, nil
}

func (s *UsersServer) UpdateUser(ctx context.Context, in *api_v1.UpdateUserRequest) (*api_v1.UpdateUserResponse, error) {
	return &api_v1.UpdateUserResponse{}, nil
}

func (s *UsersServer) DeleteUser(ctx context.Context, in *api_v1.DeleteUserRequest) (*ptypes_empty.Empty, error) {
	return &ptypes_empty.Empty{}, nil
}
