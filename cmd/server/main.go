package main

import (
	"github.com/bhubhik/habit-tracker-api/internal/server/habit"
	"github.com/bhubhik/habit-tracker-api/internal/server/hello"
	habitpb "github.com/bhubhik/habit-tracker-api/pb/habit/v1"
	hellopb "github.com/bhubhik/habit-tracker-api/pb/hello/v1"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"log"
	"net"
)

func main() {
	lis, err := net.Listen("tcp", ":5001")
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	s := grpc.NewServer()
	habitpb.RegisterHabitServiceServer(s, &habit.HabitServiceServer{})
	hellopb.RegisterHelloServiceServer(s, &hello.HelloServer{})

	reflection.Register(s)

	log.Println("gRPC server listening on :5001")
	if err := s.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}
