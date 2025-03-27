package grpc

import (
	"io"
	"context"

	"github.com/garrett-partenza-us/solana-grpc/methods"
	"github.com/garrett-partenza-us/solana-grpc/proto"
)

// Implements the SolanaService
type Server struct {
	proto.UnimplementedSolanaServiceServer
}

func (s *Server) GetLatestBlockHash(ctx context.Context, in *proto.GetLatestBlockHashRequest) (*proto.GetLatestBlockHashResponse, error) {
	jsonResponse := methods.GetLatestBlockHashJSONRPC()
	return &proto.GetLatestBlockHashResponse{
		Blockhash:            jsonResponse.Result.Value.Blockhash,
		LastValidBlockHeight: jsonResponse.Result.Value.LastValidBlockHeight,
	}, nil
}

func (s *Server) GetAccountBalance(ctx context.Context, in *proto.GetAccountBalanceRequest) (*proto.GetAccountBalanceResponse, error) {
	jsonResponse := methods.GetAccountBalanceJSONRPC(in.Account)
	return &proto.GetAccountBalanceResponse{
		Slot:  jsonResponse.Result.Context.Slot,
		Value: jsonResponse.Result.Value,
	}, nil

}

func (s *Server) GetSlotLeader(in *proto.GetSlotLeaderRequest, stream proto.SolanaService_GetSlotLeaderServer) error {
	for {
		jsonResponse := methods.GetSlotLeaderJSONRPC()
		response := &proto.GetSlotLeaderResponse {
			Pubkey: jsonResponse.Result,
		}

		if err := stream.Send(response); err != nil {
			if err == io.EOF {
				break
			}
			return err
		}
	}
	return nil
}
