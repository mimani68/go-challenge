package service

import (
	"context"
	"fmt"

	"app.ir/internal/data/repository"
	pb "app.ir/proto"
)

type SegmentService interface {
	StoreUserSegmantation(context.Context, *pb.UserSegmantRequest) (*pb.UserSegmantResponse, error)
	ShowUserInSegmant(context.Context, *pb.SegmantRequest) (*pb.SegmantResponse, error)
	mustEmbedUnimplementedEstimateServer()
}

type segmentService struct {
	repo repository.SegmentRepository
}

func NewSegmentService(repo repository.SegmentRepository) SegmentService {
	return &segmentService{repo}
}

func (u *segmentService) StoreUserSegmantation(ctx context.Context, pb_request *pb.UserSegmantRequest) (*pb.UserSegmantResponse, error) {
	// a := u.
	return &pb.UserSegmantResponse{
		Message: "New user added",
		Success: "true",
	}, nil
}

func (u *segmentService) ShowUserInSegmant(ctx context.Context, pb_request *pb.SegmantRequest) (*pb.SegmantResponse, error) {
	// return u.Repository.GetSegments(ctx, skip, limit)
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

func (u *segmentService) mustEmbedUnimplementedEstimateServer() {
	fmt.Println("Need implemention")
}
