package handler

import (
	"context"

	pb "app.ir/proto"
)

type segmentHandler struct {
	srv pb.SegmentServiceServer
}

func NewSegmentHandler(a pb.SegmentServiceServer) pb.SegmentServiceServer {
	return &segmentHandler{
		srv: a,
	}
}

func (s *segmentHandler) StoreUserSegmentation(ctx context.Context, in *pb.UserSegmentRequest) (*pb.UserSegmentResponse, error) {
	return s.srv.StoreUserSegmentation(ctx, in)
}

func (s *segmentHandler) ShowUserInSegment(ctx context.Context, in *pb.SegmentRequest) (*pb.SegmentResponse, error) {
	return s.srv.ShowUserInSegment(ctx, in)
}

func (s *segmentHandler) Estimate(ctx context.Context, in *pb.SegmentRequest) (*pb.EstimateResponse, error) {
	return s.srv.Estimate(ctx, in)
}
