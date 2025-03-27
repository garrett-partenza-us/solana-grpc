package main

import (
	"io"
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
	runGetSlotLeaderStream(client);

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

func runGetSlotLeaderStream(client proto.SolanaServiceClient) {
	req := &proto.GetSlotLeaderRequest{}
	stream, err := client.GetSlotLeaderStream(context.Background(), req)
	if err != nil {
		log.Fatal("Error starting slot leader stream: ", err)
	}
	for {
		response, err := stream.Recv()
		if err != nil {
			if err == io.EOF {
				log.Println("Server has closed the connection")
			}
		}
		log.Println("Recieved slot leader: ", response.Pubkey)
	}
}
