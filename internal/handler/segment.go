package handler

import (
	"context"
	"log"

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
	log.Printf("New user added: %s", in.GetUser())
	return s.srv.StoreUserSegmantation(ctx, in)
}

func (s *segmentHandler) ShowUserInSegmant(ctx context.Context, in *pb.SegmantRequest) (*pb.SegmantResponse, error) {
	log.Printf("We should query on Segment: %s", in.GetSegmant())
	return s.srv.ShowUserInSegmant(ctx, in)
}
