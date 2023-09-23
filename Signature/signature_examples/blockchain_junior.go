package main

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha256"
	"fmt"
	// "os"
	"time"
	// "encoding/json"
)

type Block struct {
	Data         map[string]string
    Hash         string
    PrevHash     string
    Timestamp    time.Time
    Signature    []byte
}

type Blockchain struct {
	Blocks []*Block
}

func NewGenesisBlock(creator string, privateKey *rsa.PrivateKey) *Block {
	data := map[string]string{"creator": creator}
	return NewBlock(data, "", privateKey)
}

func NewBlock(data map[string]string, prevHash string, privateKey *rsa.PrivateKey) *Block {
	block := &Block{
		Data:         data,
		PrevHash:     prevHash,
		Timestamp:    time.Now(),
	}
    block.Hash = block.calculateHash()
    block.Signature = block.signBlock(privateKey)
	return block
}

func (b *Block) calculateHash() string {
	blockData := fmt.Sprintf("%v-%v-%v", b.Data, b.PrevHash, b.Timestamp)
	hash := sha256.Sum256([]byte(blockData))
    return fmt.Sprintf("%x", hash)
}

func (b *Block) signBlock(privateKey *rsa.PrivateKey) []byte {
	blockHash := b.calculateHash()
	signature, _ := rsa.SignPKCS1v15(rand.Reader, privateKey, crypto.SHA256, []byte(blockHash))
	return signature
}

func main() {
    // Generate private & public keys
	privateKey, _ := rsa.GenerateKey(rand.Reader, 1024)

	// Create Genesis Block
	genesisBlock := NewGenesisBlock("Alice", privateKey)

	// Initialize the Blockchain with the Genesis Block
	blockchain := &Blockchain{
		Blocks: []*Block{genesisBlock},
	}

	fmt.Println("Blockchain Initialized: ", blockchain)
}
