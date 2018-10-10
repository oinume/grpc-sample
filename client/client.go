package client

import (
	"context"
	"fmt"

	"github.com/oinume/grpc-sample/proto-gen/go/proto/api/v1"
	"google.golang.org/grpc"
)

type Client struct {
	address     string
	options     *Options
	conn        *grpc.ClientConn
	usersClient api_v1.UsersClient
}

type Options struct {
	Insecure bool
}

func New(address string, options *Options) *Client {
	return &Client{
		address: address,
		options: options,
	}
}

func (c *Client) Connect(ctx context.Context) error {
	// https://medium.com/@shijuvar/writing-grpc-interceptors-in-go-bf3e7671fe48
	dialOptions := []grpc.DialOption{
		withClientUnaryInterceptor(),
	}
	if c.options.Insecure {
		dialOptions = append(dialOptions, grpc.WithInsecure())
	}
	conn, err := grpc.DialContext(ctx, c.address, dialOptions...)
	if err != nil {
		return fmt.Errorf("grpc.DialContext failed: %v", err)
	}
	c.conn = conn
	c.usersClient = api_v1.NewUsersClient(conn)
	return nil
}

func (c *Client) ListUsers(ctx context.Context) ([]*api_v1.User, error) {
	resp, err := c.usersClient.ListUsers(ctx, &api_v1.ListUsersRequest{})
	if err != nil {
		return nil, err
	}
	return resp.GetUsers(), nil
}
