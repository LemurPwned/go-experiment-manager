package main

import (
	"context"
	"log"
	"time"

	pb "github.com/LemurPwned/goman/proto"

	"google.golang.org/grpc"
)

const (
	address     = "localhost:9000"
	defaultName = "world"
)

func main() {
	// Set up a connection to the server.
	conn, err := grpc.Dial(address, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := pb.NewMetricSericeClient(conn)

	// Contact the server and print out its response.

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	s := `{ "id": "3132sd", "name": "pymtj", "date": "2012-02-03" }`

	r, err := c.SendMetrics(ctx, &pb.Metric{ExperimentID: "ewq",
		CreatedAt: time.Now().Unix(), MetricBody: s})
	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}
	log.Printf("Greeting: %s", r.GetMessage())
}
