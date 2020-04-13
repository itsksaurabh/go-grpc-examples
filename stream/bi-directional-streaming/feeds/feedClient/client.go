package main

import (
	"context"
	"fmt"
	"io"
	"log"
	"strconv"
	"time"

	"github.com/itsksaurabh/udemy/grpc/stream/bi-directional-streaming/feeds/feedpb"

	"google.golang.org/grpc"
)

func main() {
	conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("could not connect: %v", err)
	}

	c := feedpb.NewFeedsClient(conn)
	//  get client stream
	stream, err := c.Broadcast(context.Background())
	if err != nil {
		log.Fatalf("failed to call Broadcast: %v", err)
	}

	// make blocking channel
	waitc := make(chan struct{})

	// send feeds to the stream ( go routine )
	go func() {
		for i := 1; i <= 5; i++ {
			feed := "This is feed number " + strconv.Itoa(i)
			if err := stream.Send(&feedpb.FeedRequest{Feed: feed}); err != nil {
				log.Fatalf("error while sending feed: %v", err)
			}
			time.Sleep(time.Second)
		}
		if err := stream.CloseSend(); err != nil {
			log.Fatalf("failed to close stream: %v", err)
		}
	}()

	// recieve feeds frrom the stream ( go routine )
	go func() {
		for {
			msg, err := stream.Recv()
			if err == io.EOF {
				close(waitc)
				return
			}
			if err != nil {
				log.Fatalf("failed to recieve: %v", err)
				close(waitc)
				return
			}

			fmt.Println("New feed recieved : ", msg.GetFeed())
		}

	}()

	<-waitc
}
