package public

import (
	"bytes"
	"io/ioutil"
	"net/http"
	"testing"

	"github.com/drivr/go/spiri/domain"
	dp "github.com/drivr/go/spiri/domain/place"
	"github.com/drivr/go/spiri/external/api/public/mock_public"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

func TestQuery(t *testing.T) {
	testCtrl := gomock.NewController(t)
	defer testCtrl.Finish()

	// Setup
	httpClient := mock_public.NewMockhttpClient(testCtrl)
	logger := domain.NewMockLogger(testCtrl)
	repo := NewPlaceRepository(httpClient, logger)
	var searchRequest, placeRequest *http.Request
	httpClient.EXPECT().Do(gomock.Any()).Do(func(req *http.Request) { searchRequest = req }).Return(getSearchResponse(), nil)
	httpClient.EXPECT().Do(gomock.Any()).Do(func(req *http.Request) { placeRequest = req }).Return(getPlaceResponse(), nil)

	// Target
	resp, err := repo.Query(&dp.Query{
		Lat:   55.666,
		Lng:   66.555,
		Query: "Jagtvej 111",
	})

	// Assert
	assert.Equal(t, "http://api.drivr.local/places/search?query=Jagtvej+111&latlng=55.666000,66.555000", searchRequest.URL.String())
	assert.Equal(t, "http://api.drivr.local/places/ChIJVT2cO1NSUkYRLSl-47GTbWw?type=google", placeRequest.URL.String())
	assert.Nil(t, err)
	assert.NotNil(t, resp)
	assert.Equal(t, 1, len(resp))
	assert.Equal(t, 56.247399, resp[0].Lat)
	assert.Equal(t, 9.870089, resp[0].Lng)
	assert.Equal(t, "Jagtvej 111, København N, Denmark", resp[0].Name)
	assert.Equal(t, "ChIJVT2cO1NSUkYRLSl-47GTbWw", resp[0].Ref)
	assert.Equal(t, "google", resp[0].Type)
}

func getSearchResponse() *http.Response {
	json := `
		{
			"places": [
				{
				"ref": "ChIJVT2cO1NSUkYRLSl-47GTbWw",
				"name": "Jagtvej 111, København N, Denmark",
				"categories": [
					{
					"type": "street_address"
					},
					{
					"type": "geocode"
					}
				],
				"type": "google"
				}
			]
		}
	`

	return &http.Response{
		StatusCode: 200,
		Body:       ioutil.NopCloser(bytes.NewBuffer([]byte(json))),
	}
}

func getPlaceResponse() *http.Response {
	json := `
		{
			"name": "Jagtvej 111",
			"location": {
				"placeRef": "Jagtvej 111",
				"streetName": "Jagtvej",
				"houseNumber": "111",
				"zipCode": "8450",
				"city": "Favrskov Municipality",
				"country": "DK",
				"lat": 56.247399,
				"lng": 9.870089,
				"formattedAddress": "Jagtvej 111 \n8450 Favrskov Municipality Denmark",
				"singleLineFormattedAddress": "Jagtvej 111 , 8450 Favrskov Municipality Denmark",
				"addressString": "Jagtvej 111, 8450 Hammel, Denmark"
			},
			"type": "google"
		}
	`

	return &http.Response{
		StatusCode: 200,
		Body:       ioutil.NopCloser(bytes.NewBuffer([]byte(json))),
	}
}
