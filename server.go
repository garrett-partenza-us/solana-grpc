package main

import (
	"log"
	"context"
	"net"
	"net/http"

	solgrpc "github.com/garrett-partenza-us/solana-grpc/grpc"
	"github.com/garrett-partenza-us/solana-grpc/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {

	// Create a new TCP listener
	listener, err := net.Listen("tcp", ":8898")
	if err != nil {
		log.Fatal("Failed to create gRPC listener: ", err)
	}

	// Create new gRPC server
	grpcServer := grpc.NewServer()

	// Register the SolanaService server with the gRPC server
	proto.RegisterSolanaServiceServer(grpcServer, &solgrpc.Server{})

	// Register for reflection for client understanding
	reflection.Register(grpcServer)

	go func() {
		log.Println("gRPC server listening on port 8898...")
		if err := grpcServer.Serve(listener); err != nil {
			log.Fatal("Failed to serve: ", err)
		}
	}()

	ctx := context.Background()
	mux := runtime.NewServeMux()

	err = proto.RegisterSolanaServiceHandlerFromEndpoint(ctx, mux, "localhost:8898", []grpc.DialOption{grpc.WithTransportCredentials(insecure.NewCredentials())})
	if err != nil {
		log.Fatal("Failed to register gRPC gateway handler: ", err)
	}

	log.Println("gRPC Gatway listening on port 8080")
	if err := http.ListenAndServe(":8080", mux); err != nil {
		log.Fatal("Failed to serve HTTP: ", err)
	}

}
