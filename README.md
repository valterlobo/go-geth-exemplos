## Como desenvolver em Go com a lib go-ethereum:

### O que é Go Ethereum (Geth) 
Go Ethereum (Geth) é o principal cliente Ethereum de código aberto para a plataforma de blockchain Ethereum.É escrito em Go e está disponível para Windows, macOS e Linux.

### Como instalar o Geth (Linux)

 ```
 sudo apt-get install geth 
 ```
 
Para criar o modulo e obter a lib geth, pode usar os seguintes comandos:

```
go mod init [NOME-DO-SEU-PACKAGE]
go get -d github.com/ethereum/go-ethereum/
```


 Exemplo de código que você pode usar para conectar-se ao nó Ethereum e obter o bloco mais recente:

```
package main

import (
    "fmt"
    "github.com/ethereum/go-ethereum/ethclient"
)

func main() {
    // Crie um cliente Ethereum.
    client, err := ethclient.Dial("localhost:8545")
    if err != nil {
        fmt.Println(err)
        return
    }

    // Obtenha o bloco mais recente.
    block, err := client.BlockByNumber(big.NewInt(1))
    if err != nil {
        fmt.Println(err)
        return
    }

    // Imprima o hash do bloco.
    fmt.Println(block.Hash())
}
```


Para compilar e executar seu aplicativo Go, você pode usar os seguintes comandos:

```
go build
go run main.go
```

## Referências

### Ethereum Development with Go
https://goethereumbook.org/ethereum-development-with-go.pdf

### Installing Geth ( install tools)
https://geth.ethereum.org/docs/getting-started/installing-geth

### go-ethereum - Official Go implementation of the Ethereum protocol
https://geth.ethereum.org/

### Create an API to interact with Ethereum Blockchain using Golang PART 1
https://hackernoon.com/create-an-api-to-interact-with-ethereum-blockchain-using-golang-part-1-sqf3z7z


### A Step By Step Guide To Testing and Deploying Ethereum Smart Contracts in Go
https://hackernoon.com/a-step-by-step-guide-to-testing-and-deploying-ethereum-smart-contracts-in-go-9fc34b178d78


