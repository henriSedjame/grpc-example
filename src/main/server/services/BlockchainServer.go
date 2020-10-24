package services

import (
	"context"
	blockchain "github.com/hsedjame/grpc-blockchain/proto"
	"github.com/hsedjame/grpc-blockchain/src/main/server/model"
)

type BlockchainServer struct {
	Blockchain *model.Blockchain
}

func (b *BlockchainServer) AddBlock(ctx context.Context, request *blockchain.AddBlockRequest) (*blockchain.AddBlockResponse, error) {
	block := b.Blockchain.AddBlock(request.Data)
	return &blockchain.AddBlockResponse{
		Hash: block.Hash,
	}, nil
}

func (b *BlockchainServer) GetBlockchain(ctx context.Context, request *blockchain.GetBlockchainRequest) (*blockchain.GetBlockchainResponse, error) {

	response := new(blockchain.GetBlockchainResponse)
	for _, block := range b.Blockchain.Blocks {
		response.Blocks = append(response.Blocks, &blockchain.Block{
			Hash: block.Hash,
			PreviousHash: block.PrevBlockHash,
			Data: block.Data,
		})
	}
	return response, nil
}

