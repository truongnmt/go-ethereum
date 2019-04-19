# smart-contracts-with-go
A simple example of how to deploy and interact with ETH smart contracts using Golang.

## Prerequisites
- [solc](http://solidity.readthedocs.io/en/develop/installing-solidity.html)
- geth (go-ethereum)
- Ethereum node (Parity, Geth, Infura ...)

```
go get github.com/ethereum/go-ethereum
cd $GOPATH/src/github.com/ethereum/go-ethereum/
make
make devtools
```

## Generating inbox.go
```
abigen --sol=Inbox.sol --pkg=main --out=inbox.go
```

## Running
```
go build . && ./smart-contracts
```
or
```
go run main.go inbox.go
go run get_message.go inbox.go
go run set_message.go inbox.go    
```

Ref: 
- [A Step By Step Guide To Testing and Deploying Ethereum Smart Contracts in Go](https://hackernoon.com/a-step-by-step-guide-to-testing-and-deploying-ethereum-smart-contracts-in-go-9fc34b178d78)
- [Smart Contract Compilation & ABI](https://goethereumbook.org/en/smart-contract-compile/)