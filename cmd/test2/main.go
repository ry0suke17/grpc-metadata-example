package main

import (
	"context"
	"errors"
	"fmt"
	grpcmw "github.com/grpc-ecosystem/go-grpc-middleware"
	"github.com/ry0suke17/grpc-metadata-example/proto/gen"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/metadata"
	"log"
	"net"
	"os"
	"os/signal"
)

var errSomething = errors.New("err something")

func main() {
	port := 8081
	listener, err := net.Listen("tcp", fmt.Sprintf(":%d", port))
	if err != nil {
		panic(err)
	}

	conn, err := grpc.Dial(
		"test1-svc:8080",
		grpc.WithTransportCredentials(insecure.NewCredentials()),
		grpc.WithUnaryInterceptor(func(ctx context.Context, method string, req, reply interface{}, cc *grpc.ClientConn, invoker grpc.UnaryInvoker, opts ...grpc.CallOption) error {
			var trailer metadata.MD
			err := invoker(ctx, method, req, reply, cc, append(opts, grpc.Trailer(&trailer))...)
			log.Printf("trailer: %+v", trailer)
			if t := trailer.Get("test-key"); len(t) != 0 {
				return errSomething
			}
			return err
		}),
	)
	if err != nil {
		panic(err)
	}
	defer conn.Close()
	test1Client := gen.NewTest1Client(conn)

	s := grpc.NewServer(grpc.UnaryInterceptor(grpcmw.ChainUnaryServer(
		func(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
			resp, err := handler(ctx, req)
			if errors.Is(err, errSomething) {
				log.Printf("err: %v", err)
			}
			return resp, err
		},
	)))
	go func() {
		log.Printf("start gRPC server port: %v", port)
		s.Serve(listener)
	}()
	gen.RegisterTest2Server(s, &test2{test1Client: test1Client})

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt)
	<-quit
	log.Println("stopping gRPC server...")
	s.GracefulStop()
}

var _ gen.Test2Server = (*test2)(nil)

type test2 struct {
	test1Client gen.Test1Client
}

func (t *test2) Test2(ctx context.Context, _ *gen.Test2Request) (*gen.Test2Response, error) {
	var trailer metadata.MD
	resp, err := t.test1Client.Test1(ctx, &gen.Test1Request{}, grpc.Trailer(&trailer))
	if err != nil {
		return nil, err
	}
	context.WithValue(ctx, "test-ctx-key", trailer)
	return &gen.Test2Response{
		Text: fmt.Sprintf("%s: %s", "test2", resp.Text),
	}, nil
}
