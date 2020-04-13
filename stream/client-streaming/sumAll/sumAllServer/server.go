package main

import (
	"io"
	"log"
	"net"

	sumallpb "github.com/itsksaurabh/udemy/grpc/stream/client-streaming/sumAll/sumAllpb"
	"google.golang.org/grpc"
)

type server struct{}

func main() {
	lis, err := net.Listen("tcp", "localhost:50051")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	s := grpc.NewServer()
	sumallpb.RegisterSumAllServiceServer(s, &server{})

	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}

// SumAll sums up all the numbers recieved from the client stream
func (*server) SumAll(stream sumallpb.SumAllService_SumAllServer) error {
	var sum int32

	for {
		msg, err := stream.Recv()
		if err == io.EOF {
			return stream.SendAndClose(&sumallpb.SumResponse{Result: sum})
		}

		if err != nil {
			log.Fatalf("could not recieve stream: %v", err)
		}
		sum = sum + msg.GetN()
	}
}
