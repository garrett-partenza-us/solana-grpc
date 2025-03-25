package main

import (
	"log"
	"net"

	solgrpc "github.com/garrett-partenza-us/solana-grpc/grpc"
	"github.com/garrett-partenza-us/solana-grpc/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func main() {

	// Create a new TCP listener
	listener, err := net.Listen("tcp", ":8898")
	if err != nil {
		log.Fatal("Failed to create gRPC listener: ", err)
	}

	// Create new gRPC server
	s := grpc.NewServer()

	// Register the SolanaService server with the gRPC server
	proto.RegisterSolanaServiceServer(s, &solgrpc.Server{})

	// Register for reflection for client understanding
	reflection.Register(s)

	log.Println("gRPC server listening on port 8898...")
	if err := s.Serve(listener); err != nil {
		log.Fatal("Failed to serve: ", err)
	}

}
