package models

import (
	"context"
	// "encoding/json"
	"fmt"
	"math/rand"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type Record struct {
	ContactNumber   string `json:"contactNumber"`
	MedicalHistory  string `json:"medicalHistory"`
	DateOfDischarge string `json:"dateOfDischarge"`
	PatientID       string `json:"patientID"`
	FirstName       string `json:"firstName"`
	LastName        string `json:"lastName"`
	DateOfBirth     string `json:"dateOfBirth"`
	Gender          string `json:"gender"`
}

func InsertData(client *mongo.Client) {
	rand.Seed(time.Now().UnixNano())

	count := 100000
	collection := client.Database("crud").Collection("patients")

	for i := 0; i < count; i++ {
		record := generateRandomRecord()
		// jsonData, err := json.Marshal(record)
		// if err != nil {
		// 	fmt.Println("Error marshaling JSON:", err)
		// 	return
		// }
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()
		_, _ = collection.InsertOne(ctx, bson.M{
			"patientid":       record.PatientID,
			"firstname":       record.FirstName,
			"lastname":        record.LastName,
			"dob":             record.DateOfBirth,
			"gender":          record.Gender,
			"contactnumber":   record.ContactNumber,
			"medicalhistory":  record.MedicalHistory,
			"dateofdischarge": record.DateOfDischarge,
		})
		fmt.Println(i)
	}
}

func generateRandomRecord() Record {

	contactNumber := generateRandomNumber(10)
	medicalHistory := generateRandomString(50)
	dateOfDischarge := generateRandomDate()
	patientID := generateRandomStringWithMinLength(4)
	firstName := generateRandomString(10)
	lastName := generateRandomString(10)
	dateOfBirth := generateRandomDateOfBirth()
	gender := generateRandomGender()

	return Record{
		ContactNumber:   contactNumber,
		MedicalHistory:  medicalHistory,
		DateOfDischarge: dateOfDischarge,
		PatientID:       patientID,
		FirstName:       firstName,
		LastName:        lastName,
		DateOfBirth:     dateOfBirth,
		Gender:          gender,
	}
}

func generateRandomNumber(length int) string {
	digits := "0123456789"
	number := ""
	for i := 0; i < length; i++ {
		number += string(digits[rand.Intn(len(digits))])
	}
	return number
}

func generateRandomString(length int) string {
	chars := "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	str := ""
	for i := 0; i < length; i++ {
		str += string(chars[rand.Intn(len(chars))])
	}
	return str
}

func generateRandomStringWithMinLength(minLength int) string {
	length := minLength + rand.Intn(10)
	return generateRandomString(length)
}

func generateRandomDate() string {
	startDate := time.Date(2020, 1, 1, 0, 0, 0, 0, time.UTC)
	endDate := time.Now()
	randomDate := startDate.Add(time.Duration(rand.Int63n(endDate.Unix()-startDate.Unix())) * time.Second)
	return randomDate.Format("2006-01-02")
}

func generateRandomDateOfBirth() string {
	startDate := time.Date(1920, 1, 1, 0, 0, 0, 0, time.UTC)
	endDate := time.Date(2005, 1, 1, 0, 0, 0, 0, time.UTC)
	randomDate := startDate.Add(time.Duration(rand.Int63n(endDate.Unix()-startDate.Unix())) * time.Second)
	return randomDate.Format("2006-01-02")
}

func generateRandomGender() string {
	genders := []string{"male", "female"}
	return genders[rand.Intn(len(genders))]
}
