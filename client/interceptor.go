package client

import (
	"context"
	"fmt"
	"time"

	"google.golang.org/grpc"
)

//type UnaryClientInterceptor func(
//	ctx context.Context,
//	method string,
//	req, reply interface{},
//	cc *ClientConn,
//	invoker UnaryInvoker,
//	opts ...CallOption
//) error

func clientInterceptor(
	ctx context.Context,
	method string,
	req interface{},
	reply interface{},
	cc *grpc.ClientConn,
	invoker grpc.UnaryInvoker,
	opts ...grpc.CallOption,
) error {
	// Logic before invoking the invoker
	start := time.Now()
	// Calls the invoker to execute RPC
	err := invoker(ctx, method, req, reply, cc, opts...)
	// Logic after invoking the invoker
	fmt.Printf("Invoked RPC method=%s; Duration=%s; Error=%v\n", method, time.Since(start), err)
	return err
}

func withClientUnaryInterceptor() grpc.DialOption {
	return grpc.WithUnaryInterceptor(clientInterceptor)
}
