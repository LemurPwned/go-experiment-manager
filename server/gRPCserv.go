package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net"
	"time"

	pb "github.com/LemurPwned/goman/proto"
	"google.golang.org/grpc"
)

type Server struct {
}

func initServer() {
	connectMongo()
	log.Println("Attempting to launch the server")
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", 9000))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	grpcServer := grpc.NewServer()
	pb.RegisterMetricSericeServer(grpcServer, &Server{})
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %s", err)
	}
}

func parseExperimentFromString(incomingStringBody string) Experiment {
	var exp Experiment
	err := json.Unmarshal([]byte(incomingStringBody), &exp)
	if err != nil {
		log.Printf("Failed to parse the sent experiment! %s\n", err)
	}
	return exp
}

// SendMetrics sends the metrics to the server
func (s *Server) SendMetrics(ctx context.Context, in *pb.Metric) (*pb.MetricsReply, error) {
	tm := time.Unix(in.GetCreatedAt(), 0)
	bd := in.GetMetricBody()
	log.Printf("Received: %v, %v", bd, tm)
	exp := parseExperimentFromString(bd)
	log.Println(exp)
	err := addExperimentToMongo(exp)
	if err != nil {
		return nil, err
	}
	return &pb.MetricsReply{Message: "hello from the serve"}, nil
}
