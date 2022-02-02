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

func (u *segmentService) StoreUserSegmantation(ctx context.Context, pb *pb.UserSegmantRequest) (*pb.UserSegmantResponse, error) {
	// a := u.
	return &pb.UserSegmantResponse{
		Message: "New user added",
		Success: "true",
	}, nil
}

func (u *segmentService) ShowUserInSegmant(ctx context.Context, pb *pb.UserSegmantRequest) (*pb.SegmantResponse, error) {
	// return u.Repository.GetSegments(ctx, skip, limit)
	fmt.Println("Need implemention")
}

func (u *segmentService) mustEmbedUnimplementedEstimateServer() {
	// return u.Repository.GetSegmentById(ctx, id)
	fmt.Println("Need implemention")
}

// type SegmentService interface {
// 	HelloService(context.Context) (string, error)
// }

// type segmentService struct {
// 	repo repository.SegmentRepository
// }

// func NewSegmentService(repo repository.SegmentRepository) SegmentService {
// 	return &segmentService{repo}
// }

// func (u *segmentService) HelloService(ctx context.Context) (string, error) {
// 	return "salam", nil
// }
