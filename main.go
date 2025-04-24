package main

import (
	"fmt"

	"github.com/VikasFalcon/voting-system-chaincode/contracts"
	"github.com/hyperledger/fabric-contract-api-go/contractapi"
)

func main() {
	chaincode, err := contractapi.NewChaincode(new(contracts.VotingContract))
	if err != nil {
		fmt.Printf("error while chaincode initialising. error:%s", err.Error())
		return
	}

	if err := chaincode.Start(); err != nil {
		fmt.Printf("error while chaincode start. error:%s", err.Error())
	}
}
