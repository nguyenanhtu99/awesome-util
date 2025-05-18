// client/main.go
package main

import (
	"awesome-util/go/grpc/proto/client_proto"
	"context"
	"log"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	conn, err := grpc.NewClient("localhost:9191", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("could not connect: %v", err)
	}
	defer conn.Close()

	c := client_proto.NewHelloServiceClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Minute)
	defer cancel()

	res, err := c.SayHello(ctx, &client_proto.HelloRequest{Name: "Tuna"})
	if err != nil {
		log.Fatalf("error: %v", err)
	}

	log.Println("📨 Response:", res)
}
