package main

import (
	"context"
	"fmt"
	prescription "github.com/johnwoz123/pharmacy-micro-project/prescription-service/proto/prescription"
	"github.com/micro/go-micro"
)

const (
	PORT = ":50050"
)

type repository interface {
	Create(*prescription.Prescription) (*prescription.Prescription, error)
	GetAll() []*prescription.Prescription
}

type Repository struct {
	prescriptions []*prescription.Prescription
}

type service struct {
	repo repository
}

func (repo *Repository) Create(prescription *prescription.Prescription) (*prescription.Prescription, error) {
	updated := append(repo.prescriptions, prescription)
	repo.prescriptions = updated
	return prescription, nil
}

func (repo *Repository) GetAll() []*prescription.Prescription {
	return repo.prescriptions
}

func (s *service) CreatePrescription(ctx context.Context, req *prescription.Prescription, res *prescription.Response) error {
	// save
	prescript, err := s.repo.Create(req)
	if err != nil {
		return err
	}
	res.DrugCreated = true
	res.Prescription = prescript
	return nil
}

func (s *service) GetPrescriptions(ctx context.Context, req *prescription.GetRequest, res *prescription.GetResponse) error {
	allPrescriptions := s.repo.GetAll()
	res.Prescriptions = allPrescriptions
	return nil
}

func main() {
	prescriptionRepo := &Repository{}

	serv := micro.NewService(micro.Name("pharmacy-micro-project.prescription.service"), )

	serv.Init()

	// register the service
	prescription.RegisterPrescriptionServiceHandler(serv.Server(), &service{prescriptionRepo})

	// Run the server
	if err := serv.Run(); err != nil {
		fmt.Println(err)
	}

}
