package main

import (
	"fmt"
	"io"
	"log"
	"net"

	"github.com/itsksaurabh/udemy/grpc/stream/bi-directional-streaming/feeds/feedpb"

	"google.golang.org/grpc"
)

type server struct{}

func main() {
	lis, err := net.Listen("tcp", "localhost:50051")
	if err != nil {
		log.Fatalf("could not listen: %v", err)
	}

	s := grpc.NewServer()
	feedpb.RegisterFeedsServer(s, &server{})

	if err := s.Serve(lis); err != nil {
		log.Fatalf("could not start the server: %v", err)
	}
}

// Broadcast reads client stream and broadcasts recieved feeds
func (*server) Broadcast(stream feedpb.Feeds_BroadcastServer) error {
	for {
		msg, err := stream.Recv()
		if err == io.EOF {
			return nil
		}
		if err != nil {
			log.Fatalf("could not recieve from stream : %v", err)
			return err
		}

		feed := "New Feed Recieved: " + msg.GetFeed()
		fmt.Println("sending new feed...")
		stream.Send(&feedpb.FeedResponse{Feed: feed})
	}
}
