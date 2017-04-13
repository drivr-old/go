package place

import (
	"testing"

	"github.com/drivr/go/spiri/domain/mock_domain"
	"github.com/drivr/go/spiri/domain/place"
	"github.com/drivr/go/spiri/domain/place/mock_place"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

func TestSearch(t *testing.T) {
	req := &place.Query{
		Query: "Query",
		Lat:   55.66,
		Lng:   66.55,
	}
	mockCtrl := gomock.NewController(t)
	repo := mock_place.NewMockRepository(mockCtrl)
	logger := mock_domain.NewMockLogger(mockCtrl)
	svc := NewService(repo, logger)
	places := []place.Place{
		place.Place{
			Lat:  55.666,
			Lng:  66.555,
			Name: "Jagtvej 111",
			Ref:  "qweqwe123",
			Type: "google",
		},
	}
	repo.EXPECT().Query(req).Return(places, nil)

	resp, err := svc.Search(req)

	assert.Nil(t, err)
	assert.Equal(t, places, resp)
}
