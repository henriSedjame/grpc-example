package main

import (
	"context"
	"flag"
	blockchain "github.com/hsedjame/grpc-blockchain/proto"
	"google.golang.org/grpc"
	"log"
	"time"
)

func main() {
	addFlag := flag.Bool("add", false, "add block")
	getFlag := flag.Bool("get", false, "get blockchain")
	flag.Parse()

	if conn, err := grpc.Dial("localhost:8080", grpc.WithInsecure()); err != nil {
		log.Fatalf("Cannot dial server : %v", err)
	} else {
		client := blockchain.NewBlockchainClient(conn)

		if *addFlag {
			if h, err := client.AddBlock(context.Background(), &blockchain.AddBlockRequest{
				Data: time.Now().String(),
			} ); err != nil {
				log.Fatalf("Add Block failed : %v", err)
			} else {
				log.Printf("Block added successfully with hash %v", h)
			}
		}


		if *getFlag {
			if chain, err := client.GetBlockchain(context.Background(), &blockchain.GetBlockchainRequest{}); err != nil {
				log.Fatalf("Get Blockchain failed : %v", err)
 			} else {
 				for _,b := range chain.Blocks{
 					log.Printf("{\"Hash\": %s, \"Data\" : %s}", b.Hash, b.Data)
				}
			}
		}
	}
}
