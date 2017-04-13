package place

import (
	"context"

	linq "github.com/ahmetb/go-linq"
	domainplace "github.com/drivr/go/spiri/domain/place"
	pb "github.com/drivr/go/spiri/pb"
)

func DecodeGRPCSearchRequest(_ context.Context, grpcReq interface{}) (interface{}, error) {
	req := grpcReq.(*pb.SearchPlacesRequest)

	return &domainplace.Query{
		Query: req.Query,
		Lat:   req.Lat,
		Lng:   req.Lng,
	}, nil
}

func EncodeGRPCSearchResponse(_ context.Context, response interface{}) (interface{}, error) {
	resp := response.([]domainplace.Place)
	result := &pb.Places{}

	linq.From(resp).SelectT(mapPlaceToLocation).ToSlice(&result.Locations)

	return result, nil
}

func mapPlaceToLocation(place domainplace.Place) *pb.Location {
	return &pb.Location{
		AddressString: place.Name,
		Lat:           place.Lat,
		Lng:           place.Lng,
	}
}
