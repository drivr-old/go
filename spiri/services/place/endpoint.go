package place

import (
	"context"

	domainplace "github.com/drivr/go/spiri/domain/place"
	"github.com/go-kit/kit/endpoint"
)

func MakeSearchEndpoint(s Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(*domainplace.Query)
		resp, err := s.Search(req)
		return resp, err
	}
}
