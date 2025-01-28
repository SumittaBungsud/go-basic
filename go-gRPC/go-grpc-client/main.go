package main

import (
	"context"
	"log"
	"time"

	pb "go-grpc-client/grpc-hello-world/proto"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

const (
  address     = "localhost:50051"
  defaultName = "Sumitta naja eiei"
)

func main() {
  // Set up a connection to the server.
  ctx, cancel := context.WithTimeout(context.Background(), time.Second)
  defer cancel()

  creds := insecure.NewCredentials()

  // Create a new gRPC client connection (ต่อไปยังที่เดียวกับ gRPC server)
  conn, err := grpc.NewClient(address, grpc.WithTransportCredentials(creds))
  if err != nil {
    log.Fatalf("Failed to create client: %v", err)
  }
  defer conn.Close()
  c := pb.NewGreeterClient(conn)

  // Contact the server and print out its response.
  r, err := c.SayHello(ctx, &pb.HelloRequest{Name: defaultName})
  if err != nil {
    log.Fatalf("could not greet: %v", err)
  }
  log.Printf("Greeting: %s", r.GetMessages())
}