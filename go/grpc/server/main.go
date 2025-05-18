// server/main.go
package main

import (
	"awesome-util/go/grpc/proto/server_proto"
	"context"
	"log"
	"net"

	"google.golang.org/grpc"
)

type server struct {
	server_proto.UnimplementedHelloServiceServer
}

func (s *server) SayHello(ctx context.Context, req *server_proto.HelloRequest) (*server_proto.HelloResponse, error) {
	return &server_proto.HelloResponse{Message: "Xin chào, " + req.Name + "!", Num: 4}, nil
}

func main() {
	lis, err := net.Listen("tcp", ":9191")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	s := grpc.NewServer()
	server_proto.RegisterHelloServiceServer(s, &server{})

	log.Println("🚀 gRPC Server listening on :9191")
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
