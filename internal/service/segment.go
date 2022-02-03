package service

import (
	"context"
	"fmt"

	"app.ir/internal/data/repository"
	"app.ir/internal/domain"
	pb "app.ir/proto"
)

var db = []*pb.User{}

type segmentService struct {
	repo repository.SegmentRepository
}

func NewSegmentService(repo repository.SegmentRepository) pb.SegmentServiceServer {
	return &segmentService{repo}
}

func (u *segmentService) StoreUserSegmentation(ctx context.Context, pb_request *pb.UserSegmentRequest) (*pb.UserSegmentResponse, error) {
	// db = append(db, &pb.User{
	// 	User:    pb_request.User,
	// 	Segment: pb_request.Segment,
	// })
	userId, err := u.repo.CreateSegment(ctx, domain.Segment{
		UNAME:   pb_request.User,
		Segment: pb_request.Segment,
	})
	if err != nil {
		return &pb.UserSegmentResponse{
			Message: "We enconted to storeing new record into DB.",
			Success: false,
		}, err
	}
	return &pb.UserSegmentResponse{
		Message: fmt.Sprintf("New user added %s", userId),
		Success: true,
	}, nil
}

func (u *segmentService) ShowUserInSegment(ctx context.Context, pb_request *pb.SegmentRequest) (*pb.SegmentResponse, error) {
	result := []*pb.User{}
	for _, value := range db {
		if value.Segment == pb_request.Segment {
			result = append(result, value)
		}
	}
	return &pb.SegmentResponse{
		Users:   result,
		Success: true,
	}, nil
}

func (u *segmentService) Estimate(ctx context.Context, pb_request *pb.SegmentRequest) (*pb.EstimateResponse, error) {
	estimateNumber, err := u.repo.CountUserInSegment(ctx, pb_request.Segment)
	if err != nil {
		return &pb.EstimateResponse{
			Success: true,
			Count:   0,
		}, err
	}
	return &pb.EstimateResponse{
		Success: true,
		Count:   estimateNumber,
	}, nil
}
