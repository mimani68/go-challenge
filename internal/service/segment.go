package service

import (
	"context"

	"app.ir/internal/data/repository"
	pb "app.ir/proto"
)

var db = []*pb.User{
	{
		User:    "1",
		Segmant: "Football",
	},
	{
		User:    "2",
		Segmant: "Football",
	},
	{
		User:    "3",
		Segmant: "Vollybal",
	},
}

type segmentService struct {
	repo repository.SegmentRepository
}

func NewSegmentService(repo repository.SegmentRepository) pb.SegmentServiceServer {
	return &segmentService{repo}
}

func (u *segmentService) StoreUserSegmantation(ctx context.Context, pb_request *pb.UserSegmantRequest) (*pb.UserSegmantResponse, error) {
	db = append(db, &pb.User{
		User:    pb_request.User,
		Segmant: pb_request.Segmant,
	})
	return &pb.UserSegmantResponse{
		Message: "New user added",
		Success: "true",
	}, nil
}

func (u *segmentService) ShowUserInSegmant(ctx context.Context, pb_request *pb.SegmantRequest) (*pb.SegmantResponse, error) {
	result := []*pb.User{}
	for _, value := range db {
		if value.Segmant == pb_request.Segmant {
			result = append(result, value)
		}
	}
	return &pb.SegmantResponse{
		Users:   result,
		Success: "true",
	}, nil
}
