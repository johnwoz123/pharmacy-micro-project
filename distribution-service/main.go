package main

import (
	"context"
	"errors"
	"fmt"
	proto "github.com/johnwoz123/pharmacy-micro-project/distribution-service/proto/distribution"
	"github.com/micro/go-micro"
)

type DistroRepo interface {
	FindAvailableCarrier(*proto.Requirements) (*proto.Distributor, error)
}

type DistributionRepository struct {
	distributors []*proto.Distributor
}

func (repo *DistributionRepository) FindAvailableCarrier(requirments *proto.Requirements) (*proto.Distributor, error) {
	for _, reqs := range repo.distributors {
		if requirments.Capacity <= reqs.Capacity && requirments.MaxLoad <= reqs.MaxLoad {
			return reqs, nil
		}
	}
	return nil, errors.New("there are not any distributors available for transport with these requirements")
}

//grpc handler
type service struct {
	repo DistroRepo
}

func (s *service) FindAvailableCarrier(ctx context.Context, req *proto.Requirements, res *proto.Response) error {
	// Find the next available vessel
	distributor, err := s.repo.FindAvailableCarrier(req)
	if err != nil {
		return err
	}

	// Set the vessel as part of the response message type
	res.Distributor = distributor
	return nil
}

func main() {
	distributors := []*proto.Distributor{
		&proto.Distributor{Id: "cardinal1", Capacity: 10000, MaxLoad: 20000, CarMake: "Ford", CarModel: "Caravan", CarrierId: "cardinal-car-1", Name: "Cardinal", AvailableForTransport: true,},
	}

	repo :=&DistributionRepository{distributors}

	srv := micro.NewService(micro.Name("pharmacy-micro-project.distribution.service"))

	srv.Init()

	proto.RegisterDistributionServiceHandler(srv.Server(), &service{repo})

	if err := srv.Run(); err != nil {
		fmt.Println(err)
	}
}
