package models

import (
	"context"
	"fmt"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Patient struct {
	PatientID      string `json:"patientid" bson:"patientid"`
	FirstName      string `json:"firstname" bson:"firstname"`
	LastName       string `json:"lastname" bson:"lastname"`
	DOB            string `json:"dob" bson:"dob"`
	Gender         string `json:"gender" bson:"gender"`
	ContactNumber  string `json:"contactnumber" bson:"contactnumber"`
	MedicalHistory string `json:"medicalhistory" bson:"medicalhistory"`
}

type MongoDBModel struct {
	Client *mongo.Client
}

func (m *MongoDBModel) FetchData(ctx context.Context, patientID string) string {
	collection := m.Client.Database("crud").Collection("patients")
	var patient Patient

	err := collection.FindOne(ctx, bson.M{"patientid": patientID}).Decode(&patient)
	if err != nil {
		log.Printf("Failed to fetch data for patient ID %s: %v", patientID, err)
		return "Patient not found!!"
	}

	return patient.MedicalHistory
}

func ConnectDB() (*mongo.Client, error) {
	const uri = "mongodb+srv://new-user:new-user@cluster0.grve526.mongodb.net/"

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	client, err := mongo.Connect(ctx, options.Client().ApplyURI(uri))
	if err != nil {
		return nil, err
	}

	// Ping the MongoDB
	err = client.Ping(ctx, nil)
	if err != nil {
		client.Disconnect(ctx)
		return nil, err
	}

	fmt.Println("Connected to MongoDB successfully!")
	return client, nil
}

func DisconnectDB(client *mongo.Client) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	client.Disconnect(ctx)
}
