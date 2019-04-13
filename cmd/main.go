package main

import (
	"context"
	"fmt"
	"log"
	"net"
	"net/http"
	"os"
	"strconv"

	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	api_v1 "github.com/oinume/grpc-sample/proto-gen/go/proto/api/v1"
	"github.com/oinume/grpc-sample/server"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func main() {
	port, err := getEnvInt("PORT", 5000)
	if err != nil {
		log.Fatalln("PORT must be integer value")
	}
	grpcPort, err := getEnvInt("GRPC_PORT", 5001)
	if err != nil {
		log.Fatalln("GRPC_PORT must be integer value")
	}
	if port == grpcPort {
		log.Fatalln("Can't specify same port for a server")
	}

	errors := make(chan error)
	go func() {
		errors <- startGRPCServer(grpcPort)
	}()

	go func() {
		errors <- startHTTPServer(grpcPort, port)
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
	fmt.Printf("Listening on %v\n", httpPort)
	return http.ListenAndServe(fmt.Sprintf(":%d", httpPort), gatewayMux)
}

func getEnvInt(key string, dvalue int) (int, error) {
	if v := os.Getenv(key); v == "" {
		return dvalue, nil
	} else {
		value, err := strconv.ParseInt(v, 10, 32)
		if err != nil {
			return 0, err
		}
		return int(value), nil
	}
}
