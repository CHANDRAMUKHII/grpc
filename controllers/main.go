package main

import (
	model "bulkapi/models"
	pb "bulkapi/pb"
	"context"
	"fmt"
	"log"
	"net"

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
	fetchedData := make([]string, len(req.Parameters))
	for i, parameter := range req.Parameters {
		data := model.Fetchdata(parameter, client)
		fetchedData[i] = data
	}
	return &pb.FetchResponse{
		FetchedData: fetchedData,
	}, nil
}

func main() {

	//MONGODB CONNECTION//

	var err error
	serverAPI := options.ServerAPI(options.ServerAPIVersion1)
	opts := options.Client().ApplyURI(uri).SetServerAPIOptions(serverAPI)

	client, err = mongo.Connect(context.TODO(), opts)

	if err != nil {
		panic(err)
	}
	defer func() {
		if err = client.Disconnect(context.TODO()); err != nil {
			panic(err)
		}
	}()

	var result bson.M
	if err := client.Database("admin").RunCommand(context.TODO(), bson.D{{Key: "ping", Value: 1}}).Decode(&result); err != nil {
		fmt.Print("Error!")
		panic(err)
	}
	fmt.Println("Pinged your deployment. You successfully connected to MongoDB!")

	lis, err := net.Listen("tcp", ":5000")
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	fmt.Print("Server listening on port 5000")

	s := grpc.NewServer()

	pb.RegisterMongoDBServiceServer(s, &server{})

	if err := s.Serve(lis); err != nil {
		log.Fatalf("Failed to serve : %v", err)
	}
}
