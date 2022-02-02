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
	return &pb.SegmantResponse{
		Users: []*pb.User{
			{
				User:    "12",
				Segmant: "Football",
			},
		},
		Success: "true",
	}, nil
}
