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
	ElectionID string          `json:"electionID"`
	Name       string          `json:"Name"`
	Candidates []string        `json:"Candidates"`
	Votes      map[string]int  `json:"votes"`
	HasVoted   map[string]bool `json:"hasVoted"`
	StartTime  time.Time       `json:"startTime"`
	EndTime    time.Time       `json:"endTime"`
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

// RegisterElection - is use to register unique election details
func (s *SmartContract) RegisterElection(ctx contractapi.TransactionContextInterface, electionID, name string, candidates []string, startTime, endTime string) error {
	isElectionExists, err := s.isStateExists(ctx, electionID)
	if err != nil {
		return fmt.Errorf("RegisterElection(): %s", err.Error())
	}

	if isElectionExists {
		return fmt.Errorf("RegisterElection(): election is already exists with electionID: %s", electionID)
	}

	_startTime, err := time.Parse(time.RFC3339, startTime)
	if err != nil {
		return fmt.Errorf("RegisterElection(): error while converting startTime in expected format for electionID:: %s, error: %s", electionID, err.Error())
	}

	_endTime, err := time.Parse(time.RFC3339, endTime)
	if err != nil {
		return fmt.Errorf("RegisterElection(): error while converting endTime in expected format for electionID:: %s, error: %s", electionID, err.Error())
	}

	var votes = make(map[string]int)
	if len(candidates) < 2 {
		fmt.Errorf("RegisterElection(): at least 2 candidate require for election")
	}

	for _, c := range candidates {
		votes[c] = 0
	}

	var hasVoted = make(map[string]bool)

	election := Election{
		ElectionID: electionID,
		Name:       name,
		Candidates: candidates,
		Votes:      votes,
		HasVoted:   hasVoted,
		StartTime:  _startTime,
		EndTime:    _endTime,
	}

	electionJSON, err := json.Marshal(election)
	if err != nil {
		return fmt.Errorf("RegisterElection(): error while marshalling voter with electionID: %s, error:%s", electionID, err.Error())
	}

	return ctx.GetStub().PutState(electionID, electionJSON)
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
