package main

import (
	"context"
	"fmt"
	"io"
	"log"

	"github.com/itsksaurabh/udemy/grpc/stream/countdown/countdownpb"

	"google.golang.org/grpc"
)

func main() {
	conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("failed to connect: %v", err)
	}

	defer conn.Close()

	c := countdownpb.NewCountDownClient(conn)

	timer := int32(10)

	stream, err := c.Start(context.Background(), &countdownpb.CountdownRequest{Timer: timer})
	if err != nil {
		log.Fatalf("failed to start timer: %v", err)
	}

	for {
		msg, err := stream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatalf("stream read failed: %v", err)
		}

		fmt.Println(msg)
	}
}
