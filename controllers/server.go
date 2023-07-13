package controllers

import (
	"context"
	"sync"

	"bulkapi/models"
	pb "bulkapi/pb"
)

type Server struct {
	pb.UnimplementedMongoDBServiceServer
	Model *models.MongoDBModel
}

func (s *Server) FetchDataFromMongoDB(ctx context.Context, req *pb.FetchRequest) (*pb.FetchResponse, error) {
	fetchedData := s.Model.FetchData(ctx, req.Parameters)
	return &pb.FetchResponse{
		FetchedData: fetchedData,
	}, nil
}

func (s *Server) FetchDataBatchFromMongoDB(ctx context.Context, req *pb.BatchFetchRequest) (*pb.BatchFetchResponse, error) {
	fetchedData := make([]string, 0, len(req.PatientIds))
	responseChan := make(chan string, len(req.PatientIds))
	done := make(chan bool)

	var wg sync.WaitGroup
	wg.Add(len(req.PatientIds))

	for _, patientID := range req.PatientIds {
		go func(id string) {
			defer wg.Done()
			data := s.Model.FetchData(ctx, id)

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
