package main

import (
	"context"
	"encoding/json"
	prescription "github.com/johnwoz123/pharmacy-micro-project/prescription-service/proto/prescription"
	micro "github.com/micro/go-micro"
	"io/ioutil"
	"log"
	"os"
)

const (
	defaultFilename = "prescription.json"
)

func parseFile(file string) (*prescription.Prescription, error) {
	var prescript *prescription.Prescription
	data, err := ioutil.ReadFile(file)
	if err != nil {
		return nil, err
	}
	json.Unmarshal(data, &prescript)
	return prescript, err
}

func main() {

	// Set up a connection to the server.
	service := micro.NewService(micro.Name("pharmacy-micro-project.prescription.cli"))
	service.Init()
	client := prescription.NewPrescriptionService("pharmacy-micro-project.prescription.service",service.Client())

	// Contact the server and print out its response.
	file := defaultFilename
	if len(os.Args) > 1 {
		file = os.Args[1]
	}

	pre, err := parseFile(file)

	if err != nil {
		log.Fatalf("Could not parse file: %v", err)
	}

	r, err := client.CreatePrescription(context.Background(), pre)
	if err != nil {
		log.Fatalf("Could not create prescription: %v", err)
	}
	log.Printf("Created: %t", r.DrugCreated)

	// Test GetAll
	allDrugs, err := client.GetPrescriptions(context.Background(), &prescription.GetRequest{})
	if err != nil {
		log.Fatalf("Could not get all prescription: %v", err)
	}

	for _, v := range allDrugs.Prescriptions {
		log.Println(v)
	}

}
