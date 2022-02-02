package service

import (
	"context"

	"app.ir/internal/data/repository"
	pb "app.ir/proto"
)

type segmentService struct {
	repo repository.SegmentRepository
}

func NewSegmentService(repo repository.SegmentRepository) pb.SegmentServiceServer {
	return &segmentService{repo}
}

func (u *segmentService) StoreUserSegmantation(ctx context.Context, pb_request *pb.UserSegmantRequest) (*pb.UserSegmantResponse, error) {
	return &pb.UserSegmantResponse{
		Message: "New user added",
		Success: "true",
	}, nil
}

func (u *segmentService) ShowUserInSegmant(ctx context.Context, pb_request *pb.SegmantRequest) (*pb.SegmantResponse, error) {
	result := []*pb.User{}
	db := []*pb.User{
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
