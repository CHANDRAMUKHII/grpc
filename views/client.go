package main

import (
	"context"
	"fmt"
	"log"

	"google.golang.org/grpc"

	pb "bulkapi/pb"
)

func main() {

	conn, err := grpc.Dial("localhost:5000", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()

	client := pb.NewMongoDBServiceClient(conn)

	parameters := []string{"P003",
		"P004",
		"P003"}
	req := &pb.FetchRequest{
		Parameters: parameters,
	}

	resp, err := client.FetchDataFromMongoDB(context.Background(), req)
	if err != nil {
		log.Fatalf("request failed: %v", err)
	}

	for i, data := range resp.FetchedData {
		fmt.Printf("MedicalHistory of %s: %s\n", parameters[i], data)
	}
}
