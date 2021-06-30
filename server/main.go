package main

import (
	"context"
	pb "github.com/akazwz/go-grpc-demo/proto"
	"google.golang.org/grpc"
	"log"
	"net"
)

const PORT = ":9001"

type server struct {
	pb.UnimplementedGreeterServer
}

func (s *server) SayHello(_ context.Context, in *pb.HelloRequest) (*pb.HelloReply, error) {
	log.Printf("Received: %v", in.GetName())
	return &pb.HelloReply{Message: "Hello" + in.GetName()}, nil
}

func main() {
	lis, err := net.Listen("tcp", PORT)
	if err != nil {
		log.Fatalf("faild to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterGreeterServer(s, &server{})
	log.Printf("server listenning at %v", lis.Addr())
	err = s.Serve(lis)
	if err != nil {
		log.Fatalf("faild to serve: %v", err)
	}
}
