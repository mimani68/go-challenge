package handler

import (
	"context"
	"log"

	pb "app.ir/proto"
)

type SegmentHandler struct {
	pb.SegmentServiceServer
}

func NewSegmentHandler(a pb.SegmentServiceServer) pb.SegmentServiceServer {
	return &SegmentHandler{a}
}

func (s *SegmentHandler) StoreUserSegmantation(ctx context.Context, in *pb.UserSegmantRequest) (*pb.UserSegmantResponse, error) {
	log.Printf("New user added: %s", in.GetUser())
	//
	// FIXME: Store in Mongodb
	//
	return &pb.UserSegmantResponse{
		Message: "New user added",
		Success: "true",
	}, nil
	// return s.StoreUserSegmantation(ctx, in)
}

func (s *SegmentHandler) ShowUserInSegmant(ctx context.Context, in *pb.SegmantRequest) (*pb.SegmantResponse, error) {
	log.Printf("We should query on Segment: %s", in.GetSegmant())
	//
	// FIXME: Store in Mongodb
	//
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

func (s *SegmentHandler) mustEmbedUnimplementedSegmentServiceServer() {
	s.mustEmbedUnimplementedSegmentServiceServer()
}
