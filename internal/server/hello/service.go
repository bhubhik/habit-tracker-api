package hello

import (
	"context"
	pb "github.com/bhubhik/habit-tracker-api/pb/hello/v1"
)

type HelloServer struct {
	pb.UnimplementedHelloServiceServer
}

func (s *HelloServer) SayHello(ctx context.Context, req *pb.SayHelloRequest) (*pb.SayHelloResponse, error) {
	message := "Hello, " + req.GetName()
	return &pb.SayHelloResponse{Message: message}, nil
}
