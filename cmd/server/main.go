package main

import (
	"github.com/bhubhik/habit-tracker-api/internal/server/hello"
	pb "github.com/bhubhik/habit-tracker-api/pb/hello/v1"
	"google.golang.org/grpc"
	"log"
	"net"
)

func main() {
	lis, err := net.Listen("tcp", ":5001")
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	s := grpc.NewServer()
	pb.RegisterHelloServiceServer(s, &hello.HelloServer{})

	log.Println("gRPC server listening on :5001")
	if err := s.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}
