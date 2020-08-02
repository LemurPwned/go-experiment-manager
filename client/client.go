package main

import (
	"bufio"
	"context"
	"io"
	"log"
	"os"
	"path/filepath"
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
	c := pb.NewMetricServiceClient(conn)
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

	sendAsset("./client.go", c)
}

func sendAsset(fn string, client pb.MetricServiceClient) {
	fo, err := os.Open(fn)
	if err != nil {
		log.Fatalln(err)
	}
	defer fo.Close()

	ctx := context.Background()
	stream, err := client.UploadAsset(ctx)
	if err != nil {
		log.Fatalln(err)
	}
	defer stream.CloseSend()

	stat, err := fo.Stat()
	if err != nil {
		log.Fatalln(err)
	}
	log.Printf("Size is %d", stat.Size())
	req := &pb.AssetUpload{
		Data: &pb.AssetUpload_Info{
			Info: &pb.AssetInfo{
				AssetName: "test",
				AssetType: filepath.Ext(fn),
			},
		},
	}
	err = stream.Send(req)
	if err != nil {
		log.Fatalln(err)
	}

	maxBufferSize := 1024
	buffer := make([]byte, maxBufferSize)
	reader := bufio.NewReader(fo)
	for {
		n, err := reader.Read(buffer)
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatalf("Failed to read: %s\n", err)
		}

		req := &pb.AssetUpload{
			Data: &pb.AssetUpload_Content{
				Content: buffer[:n],
			},
		}
		err = stream.Send(req)
		if err != nil {
			log.Fatalf("Failed to send: %s\n", err)
		}
	}
	res, err := stream.CloseAndRecv()
	if err != nil {
		log.Fatalf("Failed to close and receive reply %s\n", err)
	}
	log.Printf("Res: %s", res.GetMessage())
	log.Println("File uploaded!")
}
