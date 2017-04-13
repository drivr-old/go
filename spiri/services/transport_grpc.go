package services

import (
	oldcontext "golang.org/x/net/context"

	pb "github.com/drivr/go/spiri/pb"
	"github.com/drivr/go/spiri/services/place"
	"github.com/go-kit/kit/endpoint"
	grpctransport "github.com/go-kit/kit/transport/grpc"
)

func MakeGRPCServer(endpoints Endpoints) pb.SpiriServer {
	options := []grpctransport.ServerOption{}
	return &spiriServer{
		placesSearch: grpctransport.NewServer(
			endpoints.PlacesSearchEndpoint,
			place.DecodeGRPCSearchRequest,
			place.EncodeGRPCSearchResponse,
			options...,
		),
	}
}

type Endpoints struct {
	PlacesSearchEndpoint endpoint.Endpoint
}

type spiriServer struct {
	placesSearch grpctransport.Handler
}

func (s *spiriServer) ReverseGeocode(ctx oldcontext.Context, req *pb.ReverseGeocodeRequest) (*pb.Location, error) {
	// TODO: implement
	return &pb.Location{}, nil
}

func (s *spiriServer) SearchPlaces(ctx oldcontext.Context, req *pb.SearchPlacesRequest) (*pb.Places, error) {
	_, rep, err := s.placesSearch.ServeGRPC(ctx, req)
	if err != nil {
		return nil, err
	}
	return rep.(*pb.Places), nil
}
