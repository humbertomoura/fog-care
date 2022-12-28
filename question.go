package main

import ( 
	"encoding/json"
	"fmt"
	"log"
	"strconv"
	"github.com/hyperledger/fabric-contract-api-go/contractapi"
)

type QuestionSmartContract struct {
	contractapi.Contract
}

type Question struct {
	IdQuestion    int     `json:"IdQuestion"`
	IdVaccine     int     `json:"idVaccine"`
	Version       string  `json:"version"`
	Date          string  `json:"date"`
	Entity        string  `json:"entity"`
	Question01    string  `json:"question01"`
	Question02    string  `json:"question02"`
	Question03    string  `json:"question03"`
	Question04    string  `json:"question04"`
	Question05    string  `json:"question05"`
	Question06    string  `json:"question06"`
	Question07    string  `json:"question07"`
	Question08    string  `json:"question08"`
	Question09    string  `json:"question09"`
	Question10    string  `json:"question10"`
	
}

func (s *QuestionSmartContract) InitLedger(ctx contractapi.TransactionContextInterface) error {
	questions := []Question{
		{IdQuestion: 35267896, IdVaccine: 54546465, Version: "1", Date: "2020-11-20", Entity: "Gov USA", Question01: "Q1", Question02: "Q2", Question03: "Q3",Question04: "Q4", Question05: "Q5",Question06: "Q6", Question07: "Q7",Question08: "Q8",Question09: "Q9",Question10: "Q10"},
		{IdQuestion: 27864589, IdVaccine: 87535904, Version: "1", Date: "2020-08-07", Entity: "Gov USA", Question01: "Q1", Question02: "Q2", Question03: "Q3",Question04: "Q4", Question05: "Q5",Question06: "Q6", Question07: "Q7",Question08: "Q8",Question09: "Q9",Question10: "Q10"},
	
	}
	for _, question := range questions {
		questionJSON, err := json.Marshal(question)
		if err != nil {
			return err
		}

		err = ctx.GetStub().PutState(strconv.Itoa(question.IdQuestion), questionJSON)
		if err != nil {
			return fmt.Errorf("failed to put question to world state. %v", err)
		}
	}
	return nil
}

func (s *QuestionSmartContract) CreateQuestion(ctx contractapi.TransactionContextInterface, id int, iv int, ve string, da string, en string, q01 string, q02 string, q03 string, q04 string, q05 string, q06 string, q07 string, q08 string, q09 string, q10 string) error {
	exists, err := s.QuestionExists(ctx, id)
	if err != nil {
		return err
	}
	if exists {
		return fmt.Errorf("The Question %d already exists", id)
	}
	question := Question{
		IdQuestion:   id,
		IdVaccine:    iv,
		Version:	  ve,
		Date:		  da,
		Entity:       en,
		Question01:   q01, 
		Question02:	  q02,
		Question03:	  q03,
		Question04:	  q04,
		Question05:   q05,
		Question06:   q06,
		Question07:   q07,
		Question08:   q08,
		Question09:   q09,
		Question10:   q10,
	}

	questionJSON, err := json.Marshal(question)
	if err != nil {
		return err
	}

	return ctx.GetStub().PutState(strconv.Itoa(id), questionJSON)
}

func (s *QuestionSmartContract) ReadQuestion(ctx contractapi.TransactionContextInterface, id int) (*Question, error) {

	questionJSON, err := ctx.GetStub().GetState(strconv.Itoa(id))
	if err != nil {
		return nil, fmt.Errorf("failed to read from world state: %v", err)
	}
	if questionJSON == nil {
		return nil, fmt.Errorf("the question %d does not exist", id)
	}

	var question Question
	err = json.Unmarshal(questionJSON, &question)
	if err != nil {
		return nil, err
	}

	return &question, nil
}

func (s *QuestionSmartContract) UpdateQuestion(ctx contractapi.TransactionContextInterface, id int, iv int, ve string, da string, en string, q01 string, q02 string, q03 string, q04 string, q05 string, q06 string, q07 string, q08 string, q09 string, q10 string) error {
	exists, err := s.QuestionExists(ctx, id)

	if err != nil {
		return err
	}
	if !exists {
		return fmt.Errorf("The question %d does not exist", id)
	}

	question := Question{
		IdQuestion:   id,
		IdVaccine:    iv,
		Version:	  ve,
		Date:		  da,
		Entity:       en,
		Question01:   q01, 
		Question02:	  q02,
		Question03:	  q03,
		Question04:	  q04,
		Question05:   q05,
		Question06:   q06,
		Question07:   q07,
		Question08:   q08,
		Question09:   q09,
		Question10:   q10,
	}
	questionJSON, err := json.Marshal(question)
	if err != nil {
		return err
	}

	return ctx.GetStub().PutState(strconv.Itoa(id), questionJSON)
}

func (s *QuestionSmartContract) DeleteQuestion(ctx contractapi.TransactionContextInterface, id int) error {
	exists, err := s.QuestionExists(ctx, id)
	if err != nil {
		return err
	}
	if !exists {
		return fmt.Errorf("The question %d does not exist", id)
	}

	return ctx.GetStub().DelState(strconv.Itoa(id))
}

func (s *QuestionSmartContract) QuestionExists(ctx contractapi.TransactionContextInterface, id int) (bool, error) {
	questionJSON, err := ctx.GetStub().GetState(strconv.Itoa(id))
	if err != nil {
		return false, fmt.Errorf("failed to read from world state: %v", err)
	}

	return questionJSON != nil, nil
}


func (s *QuestionSmartContract) GetAllQuestions(ctx contractapi.TransactionContextInterface) ([]*Question, error) {
	resultsIterator, err := ctx.GetStub().GetStateByRange("", "")
	if err != nil {
		return nil, err
	}
	defer resultsIterator.Close()

	var questions []*Question
	for resultsIterator.HasNext() {
		queryResponse, err := resultsIterator.Next()
		if err != nil {
			return nil, err
		}

		var question Question
		err = json.Unmarshal(queryResponse.Value, &question)
		if err != nil {
			return nil, err
		}
		questions = append(questions, &question)
	}

	return questions, nil
}

func main() {

	assetChaincode, err := contractapi.NewChaincode(&QuestionSmartContract{})
	if err != nil {
		log.Panicf("Error creating question chaincode: %v", err)
	}

	if err := assetChaincode.Start(); err != nil {
		log.Panicf("Error starting question chaincode: %v", err)
	}
}
