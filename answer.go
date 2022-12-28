package main

import (
	"encoding/json"
	"fmt"
	"log"
	"strconv"
	"github.com/hyperledger/fabric-contract-api-go/contractapi"
)

type AnswerSmartContract struct {
	contractapi.Contract
}

type Answer struct {
	IdAnswer     int    `json:"IdAnswer"`
	idPerson     int 	`json:"idPerson"`
	date         string `json:"date"`
	answer01     string `json:"answer01"`
	answer02     string `json:"answer02"`
	answer03     string `json:"answer03"`
	answer04     string `json:"answer04"`
	answer05     string `json:"answer05"`
	answer06     string `json:"answer06"`
	answer07     string `json:"answer07"`
	answer08     string `json:"answer08"`
	answer09     string `json:"answer09"`
	answer10     string `json:"answer10"`
	cid10_03     string `json:"cid10_03"`
	cid10_04     string `json:"cid10_04"`
	cid10_05     string `json:"cid10_05"`
}

func (s *AnswerSmartContract) InitLedger(ctx contractapi.TransactionContextInterface) error {
	answers := []Answer{
		{IdAnswer: 65678765, idPerson: 12345678, date: "2020-20-09", answer01: "A", answer02: "A", answer03: "B", answer04: "A", answer05: "C", answer06: "D", answer07: "A", answer08: "C", answer09: "A", answer10: "A"},
		{IdAnswer: 17895642, idPerson: 12345678, date: "2020-20-10", answer01: "C", answer02: "B", answer03: "B", answer04: "C", answer05: "A", answer06: "C", answer07: "B", answer08: "B", answer09: "C", answer10: "B"},
		
	}
	for _, answer := range answers {
		answerJSON, err := json.Marshal(answer)
		if err != nil {
			return err
		}

		err = ctx.GetStub().PutState(strconv.Itoa(answer.IdAnswer), answerJSON)
		if err != nil {
			return fmt.Errorf("failed to put to world state. %v", err)
		}
	}
	return nil
}

func (s *AnswerSmartContract) CreateAnswer(ctx contractapi.TransactionContextInterface, id int, ip int, da string, a01 string, a02 string, a03 string, a04 string, a05 string, a06 string, a07 string, a08 string, a09 string, a10 string) error {
	exists, err := s.AnswerExists(ctx, id)
	if err != nil {
		return err
	}
	if exists {
		return fmt.Errorf("the Answer %s already exists", id)
	}
	answer := Answer{
		IdAnswer:     id,
		idPerson:	  ip,
		date:         da,
		answer01:     a01, 
		answer02:	  a02,
		answer03:	  a03,
		answer04:	  a04,
		answer05:     a05,
		answer06:     a06,
		answer07:     a07,
		answer08:     a08,
		answer09:     a09,
		answer10:     a10,
	}

	answerJSON, err := json.Marshal(answer)
	if err != nil {
		return err
	}

	return ctx.GetStub().PutState(strconv.Itoa(id), answerJSON)
}

func (s *AnswerSmartContract) ReadAnswer(ctx contractapi.TransactionContextInterface, id int) (*Answer, error) {

	answerJSON, err := ctx.GetStub().GetState(strconv.Itoa(id))
	if err != nil {
		return nil, fmt.Errorf("failed to read from world state: %v", err)
	}
	if answerJSON == nil {
		return nil, fmt.Errorf("the answer %s does not exist", id)
	}

	var answer Answer
	err = json.Unmarshal(answerJSON, &answer)
	if err != nil {
		return nil, err
	}

	return &answer, nil
}

func (s *AnswerSmartContract) UpdateAnswer(ctx contractapi.TransactionContextInterface, id int, ip int, da string, a01 string, a02 string, a03 string, a04 string, a05 string, a06 string, a07 string, a08 string, a09 string, a10 string) error {
	exists, err := s.AnswerExists(ctx, id)

	if err != nil {
		return err
	}
	if !exists {
		return fmt.Errorf("the answer %s does not exist", id)
	}

	answer := Answer{
		IdAnswer:     id,
		idPerson:	  ip,
		date:         da,
		answer01:     a01, 
		answer02:	  a02,
		answer03:	  a03,
		answer04:	  a04,
		answer05:     a05,
		answer06:     a06,
		answer07:     a07,
		answer08:     a08,
		answer09:     a09,
		answer10:     a10,
	}
	answerJSON, err := json.Marshal(answer)
	if err != nil {
		return err
	}

	return ctx.GetStub().PutState(strconv.Itoa(id), answerJSON)
}

func (s *AnswerSmartContract) DeleteAnswer(ctx contractapi.TransactionContextInterface, id int) error {
	exists, err := s.AnswerExists(ctx, id)
	if err != nil {
		return err
	}
	if !exists {
		return fmt.Errorf("the answer %s does not exist", id)
	}

	return ctx.GetStub().DelState(strconv.Itoa(id))
}

func (s *AnswerSmartContract) AnswerExists(ctx contractapi.TransactionContextInterface, id int) (bool, error) {
	answerJSON, err := ctx.GetStub().GetState(strconv.Itoa(id))
	if err != nil {
		return false, fmt.Errorf("failed to read from world state: %v", err)
	}

	return answerJSON != nil, nil
}


func (s *AnswerSmartContract) GetAllAnswers(ctx contractapi.TransactionContextInterface) ([]*Answer, error) {
	resultsIterator, err := ctx.GetStub().GetStateByRange("", "")
	if err != nil {
		return nil, err
	}
	defer resultsIterator.Close()

	var answers []*Answer
	for resultsIterator.HasNext() {
		queryResponse, err := resultsIterator.Next()
		if err != nil {
			return nil, err
		}

		var answer Answer
		err = json.Unmarshal(queryResponse.Value, &answer)
		if err != nil {
			return nil, err
		}
		answers = append(answers, &answer)
	}

	return answers, nil
}

func main() {

	assetChaincode, err := contractapi.NewChaincode(&AnswerSmartContract{})
	if err != nil {
		log.Panicf("Error creating asset-transfer-basic chaincode: %v", err)
	}

	if err := assetChaincode.Start(); err != nil {
		log.Panicf("Error starting asset-transfer-basic chaincode: %v", err)
	}
}
