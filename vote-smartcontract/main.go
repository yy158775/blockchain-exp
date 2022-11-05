package main

import (
	"github.com/hyperledger/fabric-contract-api-go/contractapi"
	"log"
	"vote/chaincode"
)

func main() {
	voteChaincode, err := contractapi.NewChaincode(&chaincode.SmartContract{})
	if err != nil {
		log.Panicf("Error creating asset-tansfer-basic chaincode: %v",err)
	}

	if err := voteChaincode.Start();err != nil {
		log.Panicf("Error starting asset-transfer ")
	}
}