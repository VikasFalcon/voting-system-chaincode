package main

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/hyperledger/fabric-contract-api-go/contractapi"
)

type SmartContract struct {
	contractapi.Contract
}

type Voter struct {
	VoterID           string `json:"voterID"`
	Name              string `json:"name"`
	IsEligibleForVote bool   `json:"isEligibleForVote"`
}

type Election struct {
	ElectionID string         `json:"electionID"`
	Name       string         `json:"Name"`
	Candidates []string       `json:"Candidates"`
	Votes      map[string]int `json:"votes"`
	StartTime  time.Time      `json:"startTime"`
	EndTime    time.Time      `json:"endTime"`
}

// RegisterVoter - is use to register new unique voter
func (s *SmartContract) RegisterVoter(ctx contractapi.TransactionContextInterface, voterID, name string) error {
	isVoterExists, err := s.isStateExists(ctx, voterID)
	if err != nil {
		return fmt.Errorf("RegisterVoter(): %s", err.Error())
	}

	if isVoterExists {
		return fmt.Errorf("RegisterVoter(): voter is already exists with voterID: %s", voterID)
	}

	voter := Voter{
		VoterID:           voterID,
		Name:              name,
		IsEligibleForVote: true,
	}

	voterJSON, err := json.Marshal(voter)
	if err != nil {
		return fmt.Errorf("RegisterVoter(): error while marshalling voter with voterID: %s, error:%s", voterID, err.Error())
	}

	return ctx.GetStub().PutState(voterID, voterJSON)
}

// isStateExists - is use to check if data already exists in state
func (s *SmartContract) isStateExists(ctx contractapi.TransactionContextInterface, key string) (bool, error) {
	state, err := ctx.GetStub().GetState(key)
	if err != nil {
		return false, fmt.Errorf("isStateExists(): error while getting state. error:%s", err.Error())
	}

	return state != nil, nil
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
