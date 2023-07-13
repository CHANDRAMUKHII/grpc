package main

import (
	"bulkapi/controllers"
	"bulkapi/models"
	"bulkapi/pb"
	"fmt"
	"log"
	"net"

	"google.golang.org/grpc"
)

func main() {
	client, err := models.ConnectDB()
	defer models.DisconnectDB(client)

	if err != nil {
		fmt.Print("Error : ", err)
	}

	model := &models.MongoDBModel{Client: client}
	server := &controllers.Server{Model: model}

	lis, err := net.Listen("tcp", ":5002")
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	fmt.Print("Server listening on port 5002")

	s := grpc.NewServer()

	pb.RegisterMongoDBServiceServer(s, server)

	if err := s.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}
