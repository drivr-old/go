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
	spiriGrpc "github.com/drivr/go/spiri/external/grpc"
	"github.com/drivr/go/spiri/external/log"
	"github.com/drivr/go/spiri/services/place"
	pb "github.com/drivr/go/spiri/spiri"
)

var (
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
	pb.RegisterSpiriServer(grpcServer, createSpiriServer())
	go func() {
		go grpcServer.Serve(lis)
		<-stop
		grpcServer.GracefulStop()
	}()
}

func createSpiriServer() *spiriGrpc.SpiriServer {
	client := &http.Client{Timeout: 30 * time.Second}
	logger := log.New()

	placeRepository := papi.NewPlaceRepository(client, logger)

	return spiriGrpc.New(
		place.NewService(placeRepository, logger),
	)
}
