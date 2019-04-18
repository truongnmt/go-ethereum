## Prerequisites
- [solc](http://solidity.readthedocs.io/en/develop/installing-solidity.html)
- geth (go-ethereum)

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
