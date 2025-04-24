package main

import (
	"fmt"

	"github.com/hyperledger/fabric-contract-api-go/contractapi"
)

type SmartContract struct {
	contractapi.Contract
}

func main() {
	chaincode, err := contractapi.NewChaincode(new(SmartContract))
	if err != nil {
		fmt.Printf("error while chaincode initialising. error:%s", err.Error())
		return
	}

	if err := chaincode.Start(); err != nil {
		fmt.Printf("error while chaincode start. error:%s", err.Error())
	}
}
