package main

import (
	blockchain "github.com/hsedjame/grpc-blockchain/proto"
	"github.com/hsedjame/grpc-blockchain/src/main/server/model"
	"github.com/hsedjame/grpc-blockchain/src/main/server/services"
	"google.golang.org/grpc"
	"log"
	"net"
)

func main() {
	if listener, err := net.Listen("tcp", ":8080"); err != nil {
		log.Fatalf("Unable to listen port 8080 : %v", err)
	} else {
		grpcServer := grpc.NewServer()

		blockchain.RegisterBlockchainServer(grpcServer, &services.BlockchainServer{
			Blockchain: model.NewBlockChain(),
		})

		if err := grpcServer.Serve(listener); err != nil{
			log.Fatalf("Grpc server fail to run : %v", err)
		}
	}

}
