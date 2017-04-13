package main

import (
	"flag"
	"fmt"
	"net"
	"net/http"

	"google.golang.org/grpc"
	"google.golang.org/grpc/grpclog"

	"time"

	papi "github.com/drivr/go/spiri/external/api/public"
	"github.com/drivr/go/spiri/external/log"
	pb "github.com/drivr/go/spiri/pb"
	"github.com/drivr/go/spiri/services"
	"github.com/drivr/go/spiri/services/place"
)

var (
	// TODO: extract to config
	port = flag.Int("port", 10000, "The server port")
)

func main() {
	flag.Parse()

	stop := make(chan struct{}, 1)

	run(*port, stop)
}

func run(port int, stop chan struct{}) {
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", port))
	if err != nil {
		grpclog.Fatalf("failed to listen: %v", err)
	}
	var opts []grpc.ServerOption
	grpcServer := grpc.NewServer(opts...)

	placeService := createPlaceService()

	endpoints := services.Endpoints{
		PlacesSearchEndpoint: place.MakeSearchEndpoint(placeService),
	}
	spiriServer := services.MakeGRPCServer(endpoints)
	pb.RegisterSpiriServer(grpcServer, spiriServer)

	go func() {
		go grpcServer.Serve(lis)
		<-stop
		grpcServer.GracefulStop()
	}()
}

func createPlaceService() place.Service {
	client := &http.Client{Timeout: 30 * time.Second}
	logger := log.New()
	placeRepository := papi.NewPlaceRepository(client, logger)

	return place.NewService(placeRepository, logger)
}
