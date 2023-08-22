package main

import (
	"context"
	"encoding/hex"
	"fmt"
	"log"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
)

func main() {
	client, err := ethclient.Dial("wss://polygon-mumbai.g.alchemy.com/v2/{KEY}")
	if err != nil {
		log.Fatal(err)
	}

	contractAddress := common.HexToAddress("0xd1503A3eeB11C169dEA8704FFA3c57817721Fd2e")
	bytecode, err := client.CodeAt(context.Background(), contractAddress, nil) // nil is latest block
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(hex.EncodeToString(bytecode)) // 60806...10029

}
