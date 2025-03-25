package main

import (
	"context"
	"github.com/garrett-partenza-us/solana-grpc/proto"
	"google.golang.org/grpc"
	"log"
)

func main() {

	conn, err := grpc.Dial("localhost:8898", grpc.WithInsecure())
	if err != nil {
		log.Fatal("Error connecting to gRPC server: ", err)
	}
	defer conn.Close()

	client := proto.NewSolanaServiceClient(conn)

	runGetLatestBlockHash(client)
	runGetAccountBalance(client)

}

func runGetLatestBlockHash(client proto.SolanaServiceClient) {
	req := &proto.GetLatestBlockHashRequest{}
	res, err := client.GetLatestBlockHash(context.Background(), req)
	if err != nil {
		log.Fatal("Could not get block hash: ", err)
	}
	log.Println(res)
}

func runGetAccountBalance(client proto.SolanaServiceClient) {
	req := &proto.GetAccountBalanceRequest{
		Account: "6mqzvkBvdfSVhmCts4iLyBCxySigk8RMD3Uv9HVyPqQ8",
	}
	res, err := client.GetAccountBalance(context.Background(), req)
	if err != nil {
		log.Fatal("Could not get account balance: ", err)
	}
	log.Println(res)
}
