package main

import (
	"golang.org/x/net/context"

	"testing"

	"fmt"

	linq "github.com/ahmetb/go-linq"

	"google.golang.org/grpc"
	"google.golang.org/grpc/grpclog"

	pb "github.com/drivr/go/spiri/pb"
	"github.com/stretchr/testify/assert"
)

func TestSearchPlaces(t *testing.T) {
	stops := make(chan struct{}, 1)
	stopc := make(chan struct{}, 1)
	run(10000, stops)
	spiriClient := startClient(stopc)

	places, err := spiriClient.SearchPlaces(context.Background(), &pb.SearchPlacesRequest{
		Query: "Jagtvej 111",
		Lat:   55.696454,
		Lng:   12.550954,
	})

	if err != nil {
		t.Error(err)
	}

	assert.True(t, len(places.Locations) > 0, fmt.Sprintf("Locations count returned by places search should be greater than 0, but was %v.", len(places.Locations)))
	assert.True(t, linq.From(places.Locations).All(func(place interface{}) bool {
		return place.(*pb.Location).Lat > 0 && place.(*pb.Location).Lng > 0
	}), fmt.Sprintf("All locations must have Lat and Lng. Locations: %+v", places.Locations))

	stopc <- struct{}{}
	stops <- struct{}{}
}

func startClient(stop chan struct{}) pb.SpiriClient {
	var opts []grpc.DialOption
	opts = append(opts, grpc.WithInsecure())
	conn, err := grpc.Dial("127.0.0.1:10000", opts...)
	if err != nil {
		grpclog.Fatalf("fail to dial: %v", err)
	}

	client := pb.NewSpiriClient(conn)

	go func() {
		<-stop
		conn.Close()
	}()

	return client
}
