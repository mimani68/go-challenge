package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"net"

	pb "app.ir/proto"

	"google.golang.org/grpc"
)

var (
	port = flag.Int("port", 50051, "The server port")
)

type server struct {
	pb.EstimateServer
}

func (s *server) StoreUserSegmantation(ctx context.Context, in *pb.UserSegmantRequest) (*pb.UserSegmantResponse, error) {
	log.Printf("New user added: %v", in.GetUser())
	return &pb.UserSegmantResponse{
		Message: "New user added",
		Success: "true",
	}, nil
}

func (s *server) ShowUserInSegmant(ctx context.Context, in *pb.SegmantRequest) (*pb.SegmantResponse, error) {
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

func main() {
	flag.Parse()
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", *port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterEstimateServer(s, &server{})
	log.Printf("server listening at %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
