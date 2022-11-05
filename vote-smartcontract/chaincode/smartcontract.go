package chaincode

import (
	"encoding/json"
	"fmt"
	"github.com/hyperledger/fabric-contract-api-go/contractapi"
)

// SmartContract provides functions for managing an Asset
type SmartContract struct {
	contractapi.Contract
}

var NextId int

type Vote struct {
	Id       int    `json:"id"`
	Username string `json:"username"`
	Votes    int    `json:"votes"`
}

func (s *SmartContract) InitLedger(ctx contractapi.TransactionContextInterface) error {
	NextId = 0
	return nil
}

func (s *SmartContract) VoteUser(ctx contractapi.TransactionContextInterface, username string) (*Vote, error) {
	voteBytes, err := ctx.GetStub().GetState(username)
	if err != nil {
		return nil, fmt.Errorf("failed to read from world state: %v", err)
	}

	user := Vote{}

	if voteBytes == nil {
		user = Vote{
			Id:       NextId,
			Username: username,
			Votes:    1,
		}
		NextId++
	} else {
		err = json.Unmarshal(voteBytes, &user)
		if err != nil {
			return nil, err
		}
		user.Votes = user.Votes+1
	}

	userJson, err := json.Marshal(user)
	if err != nil {
		return nil, err
	}
	err = ctx.GetStub().PutState(username, userJson)
	if err != nil {
		return nil, err
	}
	return &user,nil
}

func (s *SmartContract) GetUserVote(ctx contractapi.TransactionContextInterface, username string) (*Vote, error) {
	voteBytes, err := ctx.GetStub().GetState(username)
	if err != nil {
		return nil, fmt.Errorf("failed to read from world state: %v", err)
	}
	if voteBytes == nil {
		return nil, fmt.Errorf("the UserVote %v does not exist", username)
	}
	user := Vote{}
	err = json.Unmarshal(voteBytes, &user)
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (s *SmartContract) GetAllVotes(ctx contractapi.TransactionContextInterface) ([]*Vote,error) {
	iter, err := ctx.GetStub().GetStateByRange("","")
	if err != nil {
		return nil, err
	}

	votes := make([]*Vote,0)

	for iter.HasNext() {
		kv, err := iter.Next()
		if err != nil {
			return nil, err
		}
		user := Vote{}
		value := kv.GetValue()
		err = json.Unmarshal(value, &user)
		if err != nil {
			return nil,err
		}
		votes = append(votes, &user)
	}

	return votes,nil
}