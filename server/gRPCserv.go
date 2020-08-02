package main

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net"
	"os"
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
	pb.RegisterMetricServiceServer(grpcServer, &Server{})
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

func (s *Server) UploadAsset(stream pb.MetricService_UploadAssetServer) error {
	log.Println("Starting serving the asset upload!")
	fo, err := os.Create("./output.txt")
	if err != nil {
		return err
	}
	defer fo.Close()

	var ass *pb.AssetUpload
	for {
		ass, err = stream.Recv()
		if err == io.EOF {
			// end of file
			err = stream.SendAndClose(&pb.AssetUploadReply{
				Message: "File uploaded successfully!",
			})
			log.Println("Finished reading the file!")
			return nil
		}
		if _, err := fo.Write(ass.GetContent()); err != nil {
			return err
		}
	}
}
