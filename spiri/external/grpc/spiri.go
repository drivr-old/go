package grpc

import (
	"golang.org/x/net/context"

	domainplace "github.com/drivr/go/spiri/domain/place"
	"github.com/drivr/go/spiri/services/place"
	pb "github.com/drivr/go/spiri/spiri"
)

type SpiriServer struct {
	placeService place.Service
}

func New(placeService place.Service) *SpiriServer {
	return &SpiriServer{
		placeService: placeService,
	}
}

func (s *SpiriServer) ReverseGeocode(ctx context.Context, req *pb.ReverseGeocodeRequest) (*pb.Location, error) {
	// TODO: implement
	return &pb.Location{}, nil
}

func (s *SpiriServer) SearchPlaces(ctx context.Context, req *pb.SearchPlacesRequest) (*pb.Places, error) {
	return s.placeService.Search(&domainplace.Query{
		Query: req.Query,
		Lat:   req.Lat,
		Lng:   req.Lng,
	})
}
