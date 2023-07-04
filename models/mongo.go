package models

import (
	"context"

	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type Patient struct {
	PatientID       string    `json:"PatientID"`
	FirstName       string    `json:"FirstName"`
	LastName        string    `json:"LastName"`
	DOB             string    `json:"DOB"`
	Gender          string    `json:"Gender"`
	ContactNumber   string    `json:"ContactNumber"`
	MedicalHistory  string    `json:"MedicalHistory"`
	DateOfDischarge time.Time `json:"DateOfDischarge"`
}

func Fetchdata(parameter string, client *mongo.Client) string {
	id := parameter
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	collection := client.Database("crud").Collection("patients")
	var patient Patient
	err := collection.FindOne(ctx, bson.M{"patientid": id}).Decode(&patient)
	if err != nil {

		return ""
	}

	if err != nil {

		return ""
	}
	data := patient.MedicalHistory

	return data

}
