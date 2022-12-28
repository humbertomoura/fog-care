package main

import (
	"encoding/json"
	"fmt"
	"log"
	"strconv"
	"github.com/hyperledger/fabric-contract-api-go/contractapi"
)

type PersonSmartContract struct {
	contractapi.Contract
}

type Person struct {

	IdPerson    int    `json:"IdPerson"`
	Name        string `json:"name"`
	Gender      string `json:"gender"`
	Birthdate   string `json:"birthdate"`
	Mother    	string `json:"mother"`
	Father    	string `json:"father"`
	Address    	string `json:"address"`
	City    	string `json:"city"`
	State    	string `json:"state"`
	Country    	string `json:"country"`
	Zip    		string `json:"zip"`
	Cid10_01    string `json:"cid10_01"`
	Cid10_02    string `json:"cid10_02"`
	Cid10_03    string `json:"cid10_03"`
	Cid10_04    string `json:"cid10_04"`
	Cid10_05    string `json:"cid10_05"`

	
}

func (s *PersonSmartContract) InitLedger(ctx contractapi.TransactionContextInterface) error {
	persons := []Person{
		{IdPerson: 57854378, Name: "John wash", Gender: "male", Birthdate: "1990-10-11", Mother: "Mary wash", Father: "Will Wash", Address: "120 St, 235", City: "Los Angeles", State: "CA", Country: "USA", Zip: "90038", Cid10_01: "E11", Cid10_02: "",Cid10_03: "",Cid10_04: "",Cid10_05: ""},
		{IdPerson: 49557886, Name: "Erik Climber", Gender: "male", Birthdate: "1994-02-12", Mother: "Jane Climber", Father: "Paul Climber", Address: "14 St, 235", City: "Los Angeles", State: "CA", Country: "USA", Zip: "90030", Cid10_01: "E11", Cid10_02: "",Cid10_03: "",Cid10_04: "",Cid10_05: ""},
		{IdPerson: 64546073, Name: "Jane Litle", Gender: "female", Birthdate: "1993-03-10", Mother: "Kat Litle", Father: "Donald Litle", Address: "34 St, 235", City: "Los Angeles", State: "CA", Country: "USA", Zip: "90018", Cid10_01: "E11", Cid10_02: "",Cid10_03: "",Cid10_04: "",Cid10_05: ""},
		{IdPerson: 25679788, Name: "Ann Lane", Gender: "female", Birthdate: "1992-07-10", Mother: "Zoe Lane", Father: "Charlie Lane", Address: "23 St, 235", City: "Los Angeles", State: "CA", Country: "USA", Zip: "90028", Cid10_01: "E11", Cid10_02: "",Cid10_03: "",Cid10_04: "",Cid10_05: ""},
		{IdPerson: 96553445, Name: "Julia Nickson", Gender: "female", Birthdate: "1984-05-09", Mother: "Joe Nickson", Father: "Jack Nickson", Address: "15 St, 275", City: "New York", State: "NY", Country: "USA", Zip: "80047", Cid10_01: "J45", Cid10_02: "",Cid10_03: "",Cid10_04: "",Cid10_05: ""},
		{IdPerson: 72925279, Name: "Mary Lee", Gender: "female", Birthdate: "1991-06-10", Mother: "Cloe wash", Father: "Patrick Wash", Address: "50 St, 235", City: "Los Angeles", State: "CA", Country: "USA", Zip: "90090", Cid10_01: "E11", Cid10_02: "",Cid10_03: "",Cid10_04: "",Cid10_05: ""},
		{IdPerson: 82624721, Name: "Leonard Nickson", Gender: "male", Birthdate: "1989-08-10", Mother: "Doroth  Nickson", Father: "Robert Nickson", Address: "40 St, 235", City: "Los Angeles", State: "CA", Country: "USA", Zip: "90070", Cid10_01: "E11", Cid10_02: "",Cid10_03: "",Cid10_04: "",Cid10_05: ""},
		{IdPerson: 52926438, Name: "Neil Garden", Gender: "male", Birthdate: "1991-08-10", Mother: "Rebeca Garden", Father: "Jhon Garden", Address: "35 St, 235", City: "Los Angeles", State: "CA", Country: "USA", Zip: "90071", Cid10_01: "E11", Cid10_02: "",Cid10_03: "",Cid10_04: "",Cid10_05: ""},
		{IdPerson: 53227256, Name: "Kent Lean", Gender: "male", Birthdate: "1993-11-10", Mother: "Mary Lean", Father: "Will Lean", Address: "32 St, 235", City: "Los Angeles", State: "CA", Country: "USA", Zip: "90078", Cid10_01: "E11", Cid10_02: "",Cid10_03: "",Cid10_04: "",Cid10_05: ""},
		{IdPerson: 27253727, Name: "Jane Timberlake", Gender: "female", Birthdate: "1984-12-08", Mother: "Vanessa Timberlake", Father: "John Timberlake", Address: "24 St, 205", City: "New York", State: "NY", Country: "USA", Zip: "70166", Cid10_01: "J45", Cid10_02: "",Cid10_03: "",Cid10_04: "",Cid10_05: ""},
	}
	for _, person := range persons {
		personJSON, err := json.Marshal(person)
		if err != nil {
			return err
		}

		err = ctx.GetStub().PutState(strconv.Itoa(person.IdPerson), personJSON)
		if err != nil {
			return fmt.Errorf("failed to put person to world state. %v", err)
		}
	}
	return nil
}

func (s *PersonSmartContract) CreatePerson(ctx contractapi.TransactionContextInterface, id int, na string, ge string, bd string, mo string, fa string, ad string, ci string, st string, co string, zi string, c1 string, c2 string, c3 string, c4 string, c5 string,) error {
	exists, err := s.PersonExists(ctx, id)
	if err != nil {
		return err
	}
	if exists {
		return fmt.Errorf("The Person %d already exists", id)
	}
	person := Person{
		IdPerson:    id,
		Name:        na, 
		Gender:      ge,
		Birthdate:   bd,
		Mother:    	 mo,
		Father:    	 fa,
		Address:     ad,	
		City:    	 ci,
		State:    	 st,
		Country:     co,	
		Zip:    	 zi, 	
		Cid10_01:    c1,
		Cid10_02:    c2,
		Cid10_03:    c3,
		Cid10_04:    c4,
		Cid10_05:    c5,
	}

	personJSON, err := json.Marshal(person)
	if err != nil {
		return err
	}

	return ctx.GetStub().PutState(strconv.Itoa(id), personJSON)
}

func (s *PersonSmartContract) ReadPerson(ctx contractapi.TransactionContextInterface, id int) (*Person, error) {

	personJSON, err := ctx.GetStub().GetState(strconv.Itoa(id))
	if err != nil {
		return nil, fmt.Errorf("failed to read person from world state: %v", err)
	}
	if personJSON == nil {
		return nil, fmt.Errorf("The person %d does not exist", id)
	}

	var person Person
	err = json.Unmarshal(personJSON, &person)
	if err != nil {
		return nil, err
	}

	return &person, nil
}

func (s *PersonSmartContract) UpdatePerson(ctx contractapi.TransactionContextInterface, id int, na string, ge string, bd string, mo string, fa string, ad string, ci string, st string, co string, zi string, c1 string, c2 string, c3 string, c4 string, c5 string,) error {
	exists, err := s.PersonExists(ctx, id)

	if err != nil {
		return err
	}
	if !exists {
		return fmt.Errorf("The person %d does not exist", id)
	}

	person := Person{
		IdPerson:    id,
		Name:        na, 
		Gender:      ge,
		Birthdate:   bd,
		Mother:    	 mo,
		Father:    	 fa,
		Address:     ad,	
		City:    	 ci,
		State:    	 st,
		Country:     co,	
		Zip:    	 zi, 	
		Cid10_01:    c1,
		Cid10_02:    c2,
		Cid10_03:    c3,
		Cid10_04:    c4,
		Cid10_05:    c5,
	}
	personJSON, err := json.Marshal(person)
	if err != nil {
		return err
	}

	return ctx.GetStub().PutState(strconv.Itoa(id), personJSON)
}

func (s *PersonSmartContract) DeletePerson(ctx contractapi.TransactionContextInterface, id int) error {
	exists, err := s.PersonExists(ctx, id)
	if err != nil {
		return err
	}
	if !exists {
		return fmt.Errorf("The person %d does not exist", id)
	}

	return ctx.GetStub().DelState(strconv.Itoa(id))
}

func (s *PersonSmartContract) PersonExists(ctx contractapi.TransactionContextInterface, id int) (bool, error) {
	personJSON, err := ctx.GetStub().GetState(strconv.Itoa(id))
	if err != nil {
		return false, fmt.Errorf("failed to read from world state: %v", err)
	}

	return personJSON != nil, nil
}


func (s *PersonSmartContract) GetAllPersons(ctx contractapi.TransactionContextInterface) ([]*Person, error) {
	resultsIterator, err := ctx.GetStub().GetStateByRange("", "")
	if err != nil {
		return nil, err
	}
	defer resultsIterator.Close()

	var persons []*Person
	for resultsIterator.HasNext() {
		queryResponse, err := resultsIterator.Next()
		if err != nil {
			return nil, err
		}

		var person Person
		err = json.Unmarshal(queryResponse.Value, &person)
		if err != nil {
			return nil, err
		}
		persons = append(persons, &person)
	}

	return persons, nil
}

func main() {

	assetChaincode, err := contractapi.NewChaincode(&PersonSmartContract{})
	if err != nil {
		log.Panicf("Error creating person chaincode: %v", err)
	}

	if err := assetChaincode.Start(); err != nil {
		log.Panicf("Error person chaincode: %v", err)
	}
}
