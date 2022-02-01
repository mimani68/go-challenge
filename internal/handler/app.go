package handler

import (
	"log"
	"net"

	pb "app.ir/proto"
	"google.golang.org/grpc"
)

type Server struct {
	pb.EstimateServer
}

func RunHandler(serverListerner net.Listener) {
	s := grpc.NewServer()
	pb.RegisterEstimateServer(s, &Server{})
	if err := s.Serve(serverListerner); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
