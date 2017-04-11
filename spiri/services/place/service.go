package place

import (
	linq "github.com/ahmetb/go-linq"
	"github.com/drivr/go/spiri/domain"
	"github.com/drivr/go/spiri/domain/place"
	pb "github.com/drivr/go/spiri/spiri"
)

// Service - interface for place service
type Service interface {
	Search(req *place.Query) (*pb.Places, error)
}

type service struct {
	repository place.Repository
	logger     domain.Logger
}

// NewService - creates a new Place service
func NewService(repository place.Repository, logger domain.Logger) Service {
	return &service{
		repository: repository,
		logger:     logger,
	}
}

func (p *service) Search(req *place.Query) (*pb.Places, error) {
	result := &pb.Places{}
	places, err := p.repository.Query(req)

	if err != nil {
		p.logger.Log(err)
		return result, err
	}

	if places != nil {
		linq.From(places).SelectT(mapPlaceToLocation).ToSlice(&result.Locations)
	}

	return result, err
}

func mapPlaceToLocation(place place.Place) *pb.Location {
	return &pb.Location{
		AddressString: place.Name,
		Lat:           place.Lat,
		Lng:           place.Lng,
	}
}
