package handler

import (
	"context"
	"log"

	pb "app.ir/proto"
)

func (s *Server) StoreUserSegmantation(ctx context.Context, in *pb.UserSegmantRequest) (*pb.UserSegmantResponse, error) {
	log.Printf("New user added: %v", in.GetUser())
	return &pb.UserSegmantResponse{
		Message: "New user added",
		Success: "true",
	}, nil
}

func (s *Server) ShowUserInSegmant(ctx context.Context, in *pb.SegmantRequest) (*pb.SegmantResponse, error) {
	log.Printf("We should query on Segment: %d", in.GetSegmant())
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
