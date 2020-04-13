package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/itsksaurabh/udemy/grpc/stream/client-streaming/sumAll/sumAllpb"

	"google.golang.org/grpc"
)

func main() {
	conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("could not connect: %v", err)
	}
	defer conn.Close()

	c := sumallpb.NewSumAllServiceClient(conn)

	stream, err := c.SumAll(context.Background())
	if err != nil {
		log.Fatalf("request failed: %v", err)
	}

	// send 0 to 10 numbers to the stream
	for i := 0; i <= 10; i++ {
		fmt.Printf("sending %v into the stream\n", i)
		stream.Send(&sumallpb.NumberRequest{N: int32(i)})
		time.Sleep(100 * time.Millisecond)
	}

	// close the stream and recieve result
	res, err := stream.CloseAndRecv()
	if err != nil {
		log.Fatalf("failed to recieve response: %v", err)
	}

	fmt.Println("Sum of numbers: ", res.Result)
}
