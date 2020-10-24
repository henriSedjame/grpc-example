package model

import (
	"crypto/sha256"
	"encoding/hex"
)

type Block struct {
	Hash          string
	PrevBlockHash string
	Data          string
}

type Blockchain struct {
	Blocks []*Block
}

func newHash( data string, prevHash string) string {
	bytes := sha256.Sum256([]byte(prevHash + data))
	return hex.EncodeToString(bytes[:])
}

func newBlock(data string, prevBlockHash string) *Block  {
	return & Block{
		Hash: newHash(data, prevBlockHash),
		Data: data,
		PrevBlockHash: prevBlockHash,
	}
}

func (chain *Blockchain) AddBlock(data string) *Block {
	prevBlock := chain.Blocks[len(chain.Blocks) -1]
	block := newBlock(data, prevBlock.Hash)
	chain.Blocks = append(chain.Blocks, block)
	return block
}

func NewBlockChain() *Blockchain {
	return &Blockchain{
		Blocks: []*Block{
			newBlock("Genesys block", ""),
		},
	}
}
