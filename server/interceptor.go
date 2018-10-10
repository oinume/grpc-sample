package server

import (
	"context"

	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
)

func SampleUnaryServerInterceptor() grpc.UnaryServerInterceptor {
	return func(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
		md, ok := metadata.FromIncomingContext(ctx)
		if !ok {
			return handler(ctx, req)
		}
		_, ok = md["sample"]
		if !ok {
			return handler(ctx, req)
		}
		return handler(ctx, req)
	}
}
