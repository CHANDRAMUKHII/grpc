package main

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"time"

	"google.golang.org/grpc"

	pb "bulkapi/pb"
)

type PatientID struct {
	ID string `json:"patientid"`
}

func main() {
	conn, err := grpc.Dial("localhost:5002", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()

	client := pb.NewMongoDBServiceClient(conn)

	jsonData, err := ioutil.ReadFile("ids_lakh.json")
	if err != nil {
		log.Fatalf("failed to read JSON file: %v", err)
	}

	var patientIDs []PatientID
	err = json.Unmarshal(jsonData, &patientIDs)
	if err != nil {
		log.Fatalf("failed to unmarshal JSON data: %v", err)
	}
	var average time.Duration
	for i := 0; i < 3; i++ {
		startTime := time.Now()

		batchSize := 100
		for i := 0; i < len(patientIDs); i += batchSize {
			end := i + batchSize
			if end > len(patientIDs) {
				end = len(patientIDs)
			}

			batchRequest := &pb.BatchFetchRequest{}
			for _, patientID := range patientIDs[i:end] {
				batchRequest.PatientIds = append(batchRequest.PatientIds, patientID.ID)
			}

			resp, err := client.FetchDataBatchFromMongoDB(context.Background(), batchRequest)
			if err != nil {
				log.Fatalf("request failed: %v", err)
			}

			for _, fetchedData := range resp.FetchedData {
				fmt.Println(fetchedData)
			}

		}

		elapsedTime := time.Since(startTime)
		average += elapsedTime

		fmt.Println("Total time taken:", elapsedTime)
	}
	average /= 3
	fmt.Print("Average time taken : ", average)
}
