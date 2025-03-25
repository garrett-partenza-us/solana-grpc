package main

import (
	"log"
	"context"
	"google.golang.org/grpc"
	"github.com/garrett-partenza-us/solana-grpc/proto"
)

func main() {

	conn, err := grpc.Dial("localhost:8898", grpc.WithInsecure())
	if err != nil {
		log.Fatal("Error connecting to gRPC server: ", err)
	}
	defer conn.Close()

	client := proto.NewSolanaServiceClient(conn)

	req := &proto.GetLatestBlockHashRequest{}
	res, err := client.GetLatestBlockHash(context.Background(), req)
	if err != nil {
		log.Fatal("Could not get block hash: ", err)
	}

	log.Println(res)

}
