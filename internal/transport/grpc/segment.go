package grpc

import (
	"context"
	"log"

	pb "app.ir/proto"
)

func (s *Server) StoreUserSegmantation(ctx context.Context, in *pb.UserSegmantRequest) (*pb.UserSegmantResponse, error) {
	log.Printf("New user added: %s", in.GetUser())
	//
	// FIXME: Store in Mongodb
	//
	return &pb.UserSegmantResponse{
		Message: "New user added",
		Success: "true",
	}, nil
}

func (s *Server) ShowUserInSegmant(ctx context.Context, in *pb.SegmantRequest) (*pb.SegmantResponse, error) {
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
