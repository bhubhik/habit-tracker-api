package habit

import (
	"context"
	"time"

	pb "github.com/bhubhik/habit-tracker-api/pb/habit/v1"
	"google.golang.org/protobuf/types/known/emptypb"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type HabitServiceServer struct {
	pb.UnimplementedHabitServiceServer
}

func (s *HabitServiceServer) CreateHabit(ctx context.Context, req *pb.CreateHabitRequest) (*pb.Habit, error) {
	// TODO: Add storage later
	return &pb.Habit{
		Id:         "1",
		Name:       req.GetName(),
		CreatedAt:  timestamppb.Now(),
		Recurrence: req.GetRecurrence(),
	}, nil
}

func (s *HabitServiceServer) GetHabits(ctx context.Context, req *pb.GetHabitsRequest) (*pb.GetHabitResponse, error) {
	//Returning dummy habits for now
	habits := []*pb.Habit{
		{
			Id:         "1",
			Name:       "Read",
			CreatedAt:  timestamppb.New(time.Date(2024, 5, 1, 0, 0, 0, 0, time.UTC)),
			Recurrence: []*pb.Recurrence{},
		},
	}
	return &pb.GetHabitResponse{Habits: habits}, nil
}

func (s *HabitServiceServer) UpdateHabit(ctx context.Context, req *pb.UpdateHabitRequest) (*pb.Habit, error) {
	return &pb.Habit{
		Id:         req.GetId(),
		Name:       req.GetName(),
		CreatedAt:  timestamppb.New(time.Date(2025, 5, 4, 0, 0, 0, 0, time.UTC)),
		Recurrence: req.GetRecurrence(),
	}, nil
}

func (s *HabitServiceServer) DeleteHabit(ctx context.Context, req *pb.DeleteHabitRequest) (*emptypb.Empty, error) {
	return &emptypb.Empty{}, nil
}
