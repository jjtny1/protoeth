package main

import (
  "encoding/json"
  "os"

  protoGenerator "github.com/jjtny1/protoeth/go/proto_generator"
)

const template = `
[
  {
    "inputs": [],
    "name": "retrieve",
    "outputs": [
      {
        "internalType": "uint256",
        "name": "",
        "type": "uint256"
      }
    ],
    "stateMutability": "view",
    "type": "function"
  },
  {
    "inputs": [
      {
        "internalType": "uint256",
        "name": "num",
        "type": "uint256"
      },
      {
        "internalType": "uint8[15]",
        "name": "stuff",
        "type": "uint8[15]"
      },
      {
        "internalType": "address",
        "name": "f",
        "type": "address"
      }
    ],
    "name": "store",
    "outputs": [],
    "stateMutability": "nonpayable",
    "type": "function"
  }
]
`

func main() {
  var abi protoGenerator.ABI
  err := json.Unmarshal([]byte(template), &abi)
  if err != nil {
    panic("unable to unmarshal abi: " + err.Error())
  }

  protoString, err := protoGenerator.GenerateSolidityContractProtoFileFromABIFile(abi)
  if err != nil {
    panic("unable to generate protobuf file: " + err.Error())
  }
  err = os.WriteFile("protos/solidity_contract.proto", []byte(protoString), 0644)
  if err != nil {
    panic("unable to write generated protobuf file: " + err.Error())
  }
}
