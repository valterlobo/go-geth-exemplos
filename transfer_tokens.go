package main

import (
	"context"
	"fmt"
	"log"
	"math/big"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/metachris/eth-go-bindings/erc20"
)

func main() {
	client, err := ethclient.Dial("https://polygon-mumbai.g.alchemy.com/v2/{key}") // API KEY needed from infura
	if err != nil {
		log.Fatal(err)
	}
	address := common.HexToAddress("0xf57107A130a7170fb1dE16424046B553f2701c23") // contract of token
	token, err := erc20.NewErc20(address, client)                                // see metachris/eth-go-bindings for go bindings
	if err != nil {
		log.Fatal(err)
	}
	name, err := token.Name(nil)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Name of token:", name) // optional
	//key, err := ioutil.ReadFile("./keystores/{{KEYSTORE_FILE}}") // keystore of wallet that holds ERC-20 tokens
	chainId, err := client.NetworkID(context.Background())
	if err != nil {
		log.Fatal(err)
	}

	privateKey, err := crypto.HexToECDSA("xxxxxxxxx")
	if err != nil {
		log.Fatal(err)
	}

	publicKey := privateKey.Public()
	fmt.Println(publicKey)
	auth, err := bind.NewKeyedTransactorWithChainID(privateKey, chainId)
	if err != nil {
		fmt.Println("AQUI")
		log.Fatal(err)
	}
	amount := new(big.Int)
	amount.SetString("1000000000000000000000", 10)
	value := amount                                                                                           //big.NewInt(1000000000000000000000)                                                               // 10 of given token
	tx, err := token.Transfer(auth, common.HexToAddress("0x279609F03D9624E2Efa18ac5bFEc03be9D56EE9a"), value) // recipient address
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("transfer pending: 0x%x\n", tx.Hash())
}
