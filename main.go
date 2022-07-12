package main

import (
	"fmt"

	"github.com/ethereum/go-ethereum/common"
	"github.com/flashbots/go-boost-utils/bls"
	"github.com/flashbots/go-boost-utils/types"
)

func main() {
	newPrivateKey, _, err := bls.GenerateNewKeypair()
	if err != nil {
		fmt.Println("could not generate secret key", "test")
	}
	boostSecretKey := *newPrivateKey
	pkBytes := bls.PublicKeyFromSecretKey(&boostSecretKey).Compress()
	pubKey := types.PublicKey{}
	pubKey.FromSlice(pkBytes)
	fmt.Println("private key ", common.BytesToHash(boostSecretKey.ToBEndian()).Hex())
	fmt.Println("public key ", pubKey.String())
}
