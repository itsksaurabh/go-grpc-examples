package main

import (
	"log"
	"net"
	"time"

	"github.com/itsksaurabh/udemy/grpc/stream/server-streaming/countdown/countdownpb"
	"google.golang.org/grpc"
)

type server struct{}

func main() {
	lis, err := net.Listen("tcp", "localhost:50051")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	s := grpc.NewServer()
	countdownpb.RegisterCountDownServer(s, &server{})

	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}

// Start starts countdown
func (*server) Start(req *countdownpb.CountdownRequest, stream countdownpb.CountDown_StartServer) error {
	t := req.GetTimer()
	for t > 0 {
		res := countdownpb.CountDownResponse{Count: t}
		stream.Send(&res)
		t--
		time.Sleep(time.Second)
	}
	return nil
}
