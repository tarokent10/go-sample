package main

import (
	"context"
	"fmt"
	"log"
	"net"

	pb "github.com/tarokent10/go-sample/grpcsample/grpc"
	"google.golang.org/grpc"
)

const (
	port = ":50051"
)

// server is used to implement helloworld.GreeterServer.
type server struct {
	pb.UnimplementedHelloServiceServer
}

// SayHello implements helloworld.GreeterServer
func (s *server) SayHello(ctx context.Context, in *pb.User) (*pb.HelloResult, error) {
	log.Printf("Received: %v", fmt.Sprintf("こんにちは、%sさん（%d歳）", in.Name, in.Age))
	return &pb.HelloResult{Result: 0}, nil
}

func main() {
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	println("listen port at" + port)
	s := grpc.NewServer()
	pb.RegisterHelloServiceServer(s, &server{})
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
