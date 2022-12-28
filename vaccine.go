package main

import (
	"encoding/json"
	"fmt"
	"log"
	"strconv"
	"github.com/hyperledger/fabric-contract-api-go/contractapi"
)

type VaccineSmartContract struct {
	contractapi.Contract
}

type Vaccine struct {
	IdVaccine    int    `json:"IdVaccine"`
	Gtin         int    `json:"gtin"`
	Name         string `json:"name"`
	Version      string `json:"version"`
	Country      string `json:"country"`
	MinTemp      int    `json:"minTemp"`
	MaxTemp      int    `json:"maxTemp"`
	ExpirityDays int    `json:"expirityDays"`
	Laboratory   string `json:"laboratory"`
	MinDose      int    `json:"minDose"`
	MaxDose      int    `json:"maxdose"`
	DoseInterval int    `json:"doseInterval"`
}

func (s *VaccineSmartContract) InitLedger(ctx contractapi.TransactionContextInterface) error {
	vaccines := []Vaccine{
		{IdVaccine: 54546465, Gtin: 12345678, Name: "Coronavac",                    Version: "1",         Country: "Brazil", MinTemp: -5,   MaxTemp: 50,  ExpirityDays: 180, Laboratory: "Butantanta",         MinDose: 2, MaxDose: 2, DoseInterval: 30},
		{IdVaccine: 87535904, Gtin: 76548952, Name: "Pfizer/BioNTech",              Version: "1",         Country: "USA",    MinTemp: -200, MaxTemp: -70, ExpirityDays: 120, Laboratory: "Pfizer",             MinDose: 1, MaxDose: 1, DoseInterval: 0},
		{IdVaccine: 57379532, Gtin: 84356757, Name: "Oxford/AstraZenica/Fiocruz",   Version: "1",         Country: "USA",    MinTemp: -20,  MaxTemp: 50,  ExpirityDays: 120, Laboratory: "Orford",             MinDose: 1, MaxDose: 1, DoseInterval: 0},
		{IdVaccine: 84656347, Gtin: 67545456, Name: "Moderna",                      Version: "1 Beta",    Country: "USA",    MinTemp: -10,  MaxTemp: 60,  ExpirityDays: 120, Laboratory: "Moderna",            MinDose: 1, MaxDose: 1, DoseInterval: 0},
		{IdVaccine: 13955467, Gtin: 23454357, Name: "Sputnik V/Instituto Gamaleya", Version: "0.1 Alpha", Country: "Russia", MinTemp: -20,  MaxTemp: 100, ExpirityDays: 120, Laboratory: "Instituto Gamaleya", MinDose: 1, MaxDose: 2, DoseInterval: 40},
		{IdVaccine: 24879236, Gtin: 79653432, Name: "Janssen",                      Version: "1 Alpha",   Country: "USA",    MinTemp: -10,  MaxTemp: 70,  ExpirityDays: 120, Laboratory: "Johnson-Johnson",    MinDose: 1, MaxDose: 1, DoseInterval: 0},
	}
	for _, vaccine := range vaccines {
		vaccineJSON, err := json.Marshal(vaccine)
		if err != nil {
			return err
		}

		err = ctx.GetStub().PutState(strconv.Itoa(vaccine.IdVaccine), vaccineJSON)
		if err != nil {
			return fmt.Errorf("failed to put vaccine to world state. %v", err)
		}
	}
	return nil
}

func (s *VaccineSmartContract) CreateVaccine(ctx contractapi.TransactionContextInterface, id int, gt int, na string, ve string, co string, minT int, maxT int, expD int, lab string, minD int, maxD int, doseI int,) error {
	exists, err := s.VaccineExists(ctx, id)
	
	if err != nil {
		return err
	}
	if exists {
		return fmt.Errorf("the Vaccine %d already exists", id)
	}
	vaccine := Vaccine{
		IdVaccine:    id,
		Gtin:         gt,
		Name:         na,
		Version:      ve,
		Country:      co,
		MinTemp:      minT,
		MaxTemp:      maxT,
		ExpirityDays: expD,
		Laboratory:   lab,
		MinDose:      minD,
		MaxDose:      maxD,
		DoseInterval: doseI,
	}

	vaccineJSON, err := json.Marshal(vaccine)
	if err != nil {
		return err
	}

	return ctx.GetStub().PutState(strconv.Itoa(id), vaccineJSON)
}

func (s *VaccineSmartContract) ReadVaccine(ctx contractapi.TransactionContextInterface, id int) (*Vaccine, error) {

	vaccineJSON, err := ctx.GetStub().GetState(strconv.Itoa(id))
	if err != nil {
		return nil, fmt.Errorf("failed to read vaccine from world state: %v", err)
	}
	if vaccineJSON == nil {
		return nil, fmt.Errorf("the vaccine %d does not exist", id)
	}

	var vaccine Vaccine
	err = json.Unmarshal(vaccineJSON, &vaccine)
	if err != nil {
		return nil, err
	}

	return &vaccine, nil
}

func (s *VaccineSmartContract) UpdateVaccine(ctx contractapi.TransactionContextInterface, id int, gt int, na string, ve string, co string, minT int, maxT int, expD int, lab string, minD int, maxD int, doseI int,) error {
	exists, err := s.VaccineExists(ctx, id)

	if err != nil {
		return err
	}
	if !exists {
		return fmt.Errorf("the vaccine %d does not exist", id)
	}

	vaccine := Vaccine{
		IdVaccine:    id,
		Gtin:         gt,
		Name:         na,
		Version:      ve,
		Country:      co,
		MinTemp:      minT,
		MaxTemp:      maxT,
		ExpirityDays: expD,
		Laboratory:   lab,
		MinDose:      minD,
		MaxDose:      maxD,
		DoseInterval: doseI,
	}
	vaccineJSON, err := json.Marshal(vaccine)
	if err != nil {
		return err
	}

	return ctx.GetStub().PutState(strconv.Itoa(id), vaccineJSON)
}

func (s *VaccineSmartContract) DeleteVaccine(ctx contractapi.TransactionContextInterface, id int) error {
	exists, err := s.VaccineExists(ctx, id)
	if err != nil {
		return err
	}
	if !exists {
		return fmt.Errorf("the vaccine %d does not exist", id)
	}

	return ctx.GetStub().DelState(strconv.Itoa(id))
}

func (s *VaccineSmartContract) VaccineExists(ctx contractapi.TransactionContextInterface, id int) (bool, error) {
	vaccineJSON, err := ctx.GetStub().GetState(strconv.Itoa(id))
	if err != nil {
		return false, fmt.Errorf("failed to read vaccine from world state: %v", err)
	}

	return vaccineJSON != nil, nil
}

// GetAllVaccines returns all cars found in world state
func (s *VaccineSmartContract) GetAllVaccines(ctx contractapi.TransactionContextInterface) ([]*Vaccine, error) {
	resultsIterator, err := ctx.GetStub().GetStateByRange("", "")
	if err != nil {
		return nil, err
	}
	defer resultsIterator.Close()

	var vaccines []*Vaccine
	for resultsIterator.HasNext() {
		queryResponse, err := resultsIterator.Next()
		if err != nil {
			return nil, err
		}

		var vaccine Vaccine
		err = json.Unmarshal(queryResponse.Value, &vaccine)
		if err != nil {
			return nil, err
		}
		vaccines = append(vaccines, &vaccine)
	}

	return vaccines, nil
}

func main() {

	assetChaincode, err := contractapi.NewChaincode(&VaccineSmartContract{})
	if err != nil {
		log.Panicf("Error creating asset-transfer-basic chaincode: %v", err)
	}

	if err := assetChaincode.Start(); err != nil {
		log.Panicf("Error starting asset-transfer-basic chaincode: %v", err)
	}
}
