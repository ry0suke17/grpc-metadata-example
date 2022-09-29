package main

import (
	"context"
	"fmt"
	"github.com/ry0suke17/grpc-metadata-example/proto/gen"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"time"
)

func main() {
	time.Sleep(10 * time.Second)

	conn, err := grpc.Dial("test2-svc:8081", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		panic(err)
	}
	defer conn.Close()
	test2Client := gen.NewTest2Client(conn)

	resp, err := test2Client.Test2(context.Background(), &gen.Test2Request{})
	if err != nil {
		panic(err)
	}
	fmt.Printf("%+v", resp)
}
