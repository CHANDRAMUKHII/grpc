package main

import (
	"bulkapi/models"
	pb "bulkapi/pb"
	"context"
	"fmt"
	"log"
	"net"
	"sync"

	"google.golang.org/grpc"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const uri = "mongodb+srv://new-user:new-user@cluster0.grve526.mongodb.net/"

var client *mongo.Client

type server struct {
	pb.MongoDBServiceServer
}

func (s *server) FetchDataFromMongoDB(ctx context.Context, req *pb.FetchRequest) (*pb.FetchResponse, error) {
	fetchedData := models.Fetchdata(ctx, req.Parameters, client)
	return &pb.FetchResponse{
		FetchedData: fetchedData,
	}, nil
}

func (s *server) FetchDataBatchFromMongoDB(ctx context.Context, req *pb.BatchFetchRequest) (*pb.BatchFetchResponse, error) {

	fetchedData := make([]string, 0, len(req.PatientIds))
	responseChan := make(chan string, len(req.PatientIds))
	done := make(chan bool)

	var wg sync.WaitGroup
	wg.Add(len(req.PatientIds))

	for _, patientID := range req.PatientIds {
		go func(id string) {
			defer wg.Done()
			data := models.Fetchdata(ctx, id, client)

			responseChan <- data
		}(patientID)
	}

	go func() {
		wg.Wait()
		done <- true
	}()

	for i := 0; i < len(req.PatientIds); i++ {
		data := <-responseChan
		fetchedData = append(fetchedData, data)
	}

	close(responseChan)
	<-done

	return &pb.BatchFetchResponse{
		FetchedData: fetchedData,
	}, nil
}

func main() {
	var err error

	serverAPI := options.ServerAPI(options.ServerAPIVersion1)
	opts := options.Client().ApplyURI(uri).SetServerAPIOptions(serverAPI).SetMaxPoolSize(10000)

	client, err = mongo.Connect(context.Background(), opts)

	if err != nil {
		panic(err)
	}
	defer func() {
		if err = client.Disconnect(context.Background()); err != nil {
			panic(err)
		}
	}()

	var result bson.M
	if err := client.Database("admin").RunCommand(context.Background(), bson.D{{Key: "ping", Value: 1}}).Decode(&result); err != nil {
		fmt.Print("Error!")
		panic(err)
	}
	fmt.Println("Pinged your deployment. You successfully connected to MongoDB!")

	lis, err := net.Listen("tcp", ":5002")
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	fmt.Print("Server listening on port 5002")

	s := grpc.NewServer()

	pb.RegisterMongoDBServiceServer(s, &server{})

	if err := s.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}
