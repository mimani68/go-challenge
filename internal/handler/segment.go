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

func (s *segmentHandler) StoreUserSegmantation(ctx context.Context, in *pb.UserSegmantRequest) (*pb.UserSegmantResponse, error) {
	return s.srv.StoreUserSegmantation(ctx, in)
}

func (s *segmentHandler) ShowUserInSegmant(ctx context.Context, in *pb.SegmantRequest) (*pb.SegmantResponse, error) {
	return s.srv.ShowUserInSegmant(ctx, in)
}

func (s *segmentHandler) Estimate(ctx context.Context, in *pb.SegmantRequest) (*pb.EstimateResponse, error) {
	return s.srv.Estimate(ctx, in)
}
