package main

import (
	"context"
	"log"
	"net"

	"github.com/itsksaurabh/udemy/grpc/unary/sum/sumpb"

	"google.golang.org/grpc"
)

type server struct{}

func main() {
	lis, err := net.Listen("tcp", "localhost:50051")
	if err != nil {
		log.Fatalf("failed to liesten: %v", err)
	}

	s := grpc.NewServer()
	sumpb.RegisterSumServer(s, &server{})

	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to start server %v", err)
	}

}

func (*server) Add(ctx context.Context, req *sumpb.SumRequest) (*sumpb.SumResponse, error) {
	a, b := req.GetNumbers().GetA(), req.GetNumbers().GetB()
	sum := a + b
	return &sumpb.SumResponse{Result: sum}, nil
}
