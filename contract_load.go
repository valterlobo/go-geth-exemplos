package main

import (
	"fmt"
	"log"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"

	store "go-eth-tutorial/store"
)

func main() {
	client, err := ethclient.Dial("wss://polygon-mumbai.g.alchemy.com/v2/{KEY}")
	if err != nil {
		log.Fatal(err)
	}

	address := common.HexToAddress("0xd1503A3eeB11C169dEA8704FFA3c57817721Fd2e")
	instance, err := store.NewStore(address, client)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("contract is loaded")
	_ = instance
}
