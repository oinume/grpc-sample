package client

import (
	"context"
	"fmt"

	"github.com/oinume/grpc-sample/proto-gen/go/proto/api/v1"
	"google.golang.org/grpc"
)

type Client struct {
	address     string
	usersClient api_v1.UsersClient
}

func New(address string) *Client {
	return &Client{
		address: address,
	}
}

func (c *Client) Connect(ctx context.Context) error {
	//interceptors := grpc_middleware.ChainUnaryClient(
	//	grpcdd.UnaryClientInterceptor(grpcdd.WithServiceName(c.opts.serviceName)),
	//)
	//
	var dialOpts []grpc.DialOption
	dialOpts = append(dialOpts, grpc.WithUnaryInterceptor(interceptors))
	if c.opts.insecure {
		dialOpts = append(dialOpts, grpc.WithInsecure())
	}
	conn, err := grpc.DialContext(ctx, c.address, dialOpts...)
	if err != nil {
		return fmt.Errorf("grpc.DialContext failed: %v", err)
	}
	c.conn = conn

	c.cli = api_v1.NewUsersClient(conn)
	return nil

}

func (c *Client) ListUsers() {
}
