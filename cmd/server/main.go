package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"net"
	"net/http"

	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	"github.com/oinume/grpc-sample/proto-gen/go/proto/api/v1"
	"github.com/oinume/grpc-sample/server"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

var (
	httpPort = flag.Int("http-port", 5000, "HTTP server port. default is 5000")
	grpcPort = flag.Int("grpc-port", 5001, "gRPC server port. default is 5001")
)

func main() {
	flag.Parse()

	if *httpPort == *grpcPort {
		log.Fatalf("Can't specify same httpPort for a server.")
	}

	errors := make(chan error)
	go func() {
		errors <- startGRPCServer(*grpcPort)
	}()

	go func() {
		errors <- startHTTPServer(*grpcPort, *httpPort)
	}()

	for err := range errors {
		log.Fatal(err)
	}
}

func startGRPCServer(port int) error {
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	var opts []grpc.ServerOption
	//opts = append(opts, interceptor.WithUnaryServerInterceptors())
	s := grpc.NewServer(opts...)

	// Register reflection service on gRPCs.
	usersServer := &server.UsersServer{}
	api_v1.RegisterUsersServer(s, usersServer)

	reflection.Register(s)
	fmt.Printf("Starting gRPC server on %d\n", port)
	return s.Serve(lis)
}

func startHTTPServer(grpcPort, httpPort int) error {
	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	muxOptions := runtime.WithMarshalerOption(runtime.MIMEWildcard, &runtime.JSONPb{
		OrigName:     true,
		EmitDefaults: true,
	})
	gatewayMux := runtime.NewServeMux(muxOptions)
	opts := []grpc.DialOption{grpc.WithInsecure()}
	endpoint := fmt.Sprintf("127.0.0.1:%d", grpcPort)
	if err := api_v1.RegisterUsersHandlerFromEndpoint(ctx, gatewayMux, endpoint, opts); err != nil {
		return err
	}
	fmt.Printf("Starting HTTP server on %v\n", httpPort)
	return http.ListenAndServe(fmt.Sprintf(":%d", httpPort), gatewayMux)
}
