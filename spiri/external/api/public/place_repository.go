package public

import (
	"fmt"
	"net/http"
	"net/url"

	linq "github.com/ahmetb/go-linq"
	domainplace "github.com/drivr/go/spiri/domain/place"

	"github.com/drivr/go/spiri/domain"
)

// TODO: configurable url
const baseURL = "http://api.drivr.local/places"

type httpClient interface {
	Do(req *http.Request) (*http.Response, error)
}

// PlaceRepository - repository for places
type PlaceRepository struct {
	httpClient httpClient
	logger     domain.Logger
}

// NewPlaceRepository - creates a new instance
func NewPlaceRepository(httpClient httpClient, logger domain.Logger) domainplace.Repository {
	return &PlaceRepository{
		httpClient: httpClient,
		logger:     logger,
	}
}

// Query - search places by query
func (r *PlaceRepository) Query(query *domainplace.Query) ([]domainplace.Place, error) {
	latStr := url.QueryEscape(fmt.Sprintf("%f", query.Lat))
	lngStr := url.QueryEscape(fmt.Sprintf("%f", query.Lng))
	url := fmt.Sprintf("%v/search?query=%v&latlng=%v,%v", baseURL, url.QueryEscape(query.Query), latStr, lngStr)
	var result []domainplace.Place

	req, err := getRequest(url)
	if err != nil {
		return result, err
	}

	resp, err := r.httpClient.Do(req)
	if err != nil {
		r.logger.Log(fmt.Sprintf("Error searching for places. %v\n", err))
		return result, err
	}

	if resp.StatusCode != 200 {
		r.logger.Log(fmt.Sprintf("Error searching for places. %v\n", resp.StatusCode))
		return result, nil
	}

	var placesResponse placesResponse
	decode(resp, &placesResponse)

	r.enrichPlaces(&placesResponse)

	linq.From(placesResponse.Places).SelectT(mapPlace).ToSlice(&result)
	return result, nil
}

func (r *PlaceRepository) enrichPlaces(resp *placesResponse) {
	donec, errc := make(chan bool), make(chan error)
	reqCount := 0
	for idx := range resp.Places {
		if resp.Places[idx].Type == "google" {
			reqCount++
			go r.enrichPlace(&resp.Places[idx], donec, errc)
		}
	}

	for i := 0; i < reqCount; i++ {
		select {
		case <-donec:
		case err := <-errc:
			r.logger.Log(err)
		}
	}
}

func (r *PlaceRepository) enrichPlace(pl *place, donec chan bool, errc chan error) {
	url := fmt.Sprintf("%v/%v?type=google", baseURL, pl.Ref)
	req, err := getRequest(url)
	if err != nil {
		errc <- err
		return
	}

	resp, err := r.httpClient.Do(req)
	if err != nil {
		errc <- err
		return
	}

	if resp.StatusCode != 200 {
		errc <- fmt.Errorf("Error enriching place. %v", resp.StatusCode)
		return
	}

	var presp placeResponse
	if err := decode(resp, &presp); err != nil {
		errc <- err
		return
	}

	pl.Lat = presp.Location.Lat
	pl.Lng = presp.Location.Lng

	donec <- true
}

func mapPlace(pl interface{}) domainplace.Place {
	return domainplace.Place{
		Lat:  pl.(place).Lat,
		Lng:  pl.(place).Lng,
		Name: pl.(place).Name,
		Ref:  pl.(place).Ref,
		Type: pl.(place).Type,
	}
}
