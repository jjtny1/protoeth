all:

proto-go:
	go run go/cmd/main.go
	protoc -I ./protos --go_out=. ./protos/solidity_types.proto  ./protos/solidity_contract.proto 