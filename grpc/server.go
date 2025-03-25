package grpc

import (
	"log"
	"context"

	"github.com/garrett-partenza-us/solana-grpc/proto"
	"github.com/garrett-partenza-us/solana-grpc/methods"
)

// Implements the SolanaService
type Server struct {
    proto.UnimplementedSolanaServiceServer
}

func (s *Server) GetLatestBlockHash(ctx context.Context, in *proto.GetLatestBlockHashRequest) (*proto.GetLatestBlockHashResponse, error) {
	jsonResponse := methods.GetLatestBlockHashJSONRPC()
	log.Println(jsonResponse)
	return &proto.GetLatestBlockHashResponse{
		Blockhash: jsonResponse.Result.Value.Blockhash,
		LastValidBlockHeight: jsonResponse.Result.Value.LastValidBlockHeight,
	}, nil
}
