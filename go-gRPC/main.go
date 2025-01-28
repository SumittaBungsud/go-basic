package main

import (
	"context"
	pb "go-grpc/grpc-hello-world/proto"
	"log"
	"net"

	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

const (
	port = ":50051"
)

type server struct{
	pb.UnimplementedGreeterServer
}

func (s *server) SayHello(ctx context.Context, in *pb.HelloRequest) (*pb.HelloReply, error) {
	log.Printf("Receive %v", in.GetName())

	if len(in.GetName()) > 8 {
		return nil, status.Error(codes.InvalidArgument, "Name must not more than 8 characters")
	}
	return &pb.HelloReply{Messages: "Hello " + in.GetName()}, nil
}

func main(){
	lis, err := net.Listen("tcp", port)
	if err != nil{
		log.Fatalf("failed to listen: %v", err)
	}

	grpcServer := grpc.NewServer()
	pb.RegisterGreeterServer(grpcServer, &server{})

	log.Printf("Server is listening on port %v", port)

	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}