package place

import (
	"github.com/drivr/go/spiri/domain"
	"github.com/drivr/go/spiri/domain/place"
)

// Service - interface for place service
type Service interface {
	Search(req *place.Query) ([]place.Place, error)
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

func (p *service) Search(req *place.Query) ([]place.Place, error) {
	// TODO: validation

	places, err := p.repository.Query(req)

	if err != nil {
		p.logger.Log(err)
		return nil, err
	}

	return places, nil
}
