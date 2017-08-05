package main

import (
	"context"
	"log"

	"github.com/alanfran/steampipe/protocol"
	"google.golang.org/grpc"
)

// SteamPipe is a client for the Steam Pipe grpc service.
type SteamPipe struct {
	server string
}

// NewSteamPipe returns a reference to an initialized SteamPipe.
func NewSteamPipe(address string) *SteamPipe {
	return &SteamPipe{
		server: address,
	}
}

// Query sends a server query to the Steam Pipe grpc service and
// returns a response or an error.
func (s *SteamPipe) Query(address string) (*protocol.Response, error) {
	conn, err := grpc.Dial(s.server, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Error connecting to Steam Pipe: %v", err)
	}
	defer conn.Close()

	client := protocol.NewSteamPipeClient(conn)

	resp, err := client.Query(context.Background(), &protocol.Address{Addr: address})
	if err != nil {
		log.Fatalf("Error querying... %v", err)
	}

	return resp, err
}
