package habit

import (
	"context"

	pb "github.com/bhubhik/habit-tracker-api/pb/habit/v1"
	"google.golang.org/protobuf/types/known/emptypb"
)

type HabitServiceServer struct {
	pb.UnimplementedHabitServiceServer
}

func (s *HabitServiceServer) CreateHabit(ctx context.Context, req *pb.CreateHabitRequest) (*pb.Habit, error) {
	// TODO: Add storage later
	return &pb.Habit{
		Id:        "1",
		Name:      req.GetName(),
		CreatedAt: "2025-05-03",
		Frequency: req.GetFrequency(),
	}, nil
}

func (s *HabitServiceServer) GetHabits(ctx context.Context, req *pb.GetHabitsRequest) (*pb.GetHabitResponse, error) {
	//Returning dummy habits for now
	habits := []*pb.Habit{
		{
			Id:        "1",
			Name:      "Read",
			CreatedAt: "2025-05-01",
			Frequency: []string{"monday", "wednesday"},
		},
	}
	return &pb.GetHabitResponse{Habits: habits}, nil
}

func (s *HabitServiceServer) UpdateHabit(ctx context.Context, req *pb.UpdateHabitRequest) (*pb.Habit, error) {
	return &pb.Habit{
		Id:        req.GetId(),
		Name:      req.GetName(),
		CreatedAt: "2025-05-01",
		Frequency: req.GetFrequency(),
	}, nil
}

func (s *HabitServiceServer) DeleteHabit(ctx context.Context, req *pb.DeleteHabitRequest) (*emptypb.Empty, error) {
	return &emptypb.Empty{}, nil
}
