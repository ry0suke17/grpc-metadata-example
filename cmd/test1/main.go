package main

import (
	"context"
	"fmt"
	grpcmw "github.com/grpc-ecosystem/go-grpc-middleware"
	"github.com/ry0suke17/grpc-metadata-example/proto/gen"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
	"log"
	"net"
	"os"
	"os/signal"
)

func main() {
	port := 8080
	listener, err := net.Listen("tcp", fmt.Sprintf(":%d", port))
	if err != nil {
		panic(err)
	}

	s := grpc.NewServer(grpc.UnaryInterceptor(grpcmw.ChainUnaryServer(
		func(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
			trailer := metadata.Pairs("test-key", "test value")
			grpc.SetTrailer(ctx, trailer)
			return handler(ctx, req)
		},
	)))
	go func() {
		log.Printf("start gRPC server port: %v", port)
		s.Serve(listener)
	}()
	gen.RegisterTest1Server(s, test1{})

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt)
	<-quit
	log.Println("stopping gRPC server...")
	s.GracefulStop()
}

var _ gen.Test1Server = (*test1)(nil)

type test1 struct{}

func (t test1) Test1(_ context.Context, _ *gen.Test1Request) (*gen.Test1Response, error) {
	return &gen.Test1Response{
		Text: "test1",
	}, nil
}
