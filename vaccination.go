package main

import (
	"encoding/json"
	"fmt"
	"log"
	"strconv"
	"github.com/hyperledger/fabric-contract-api-go/contractapi"
)

type VaccinationSmartContract struct {
	contractapi.Contract
}

type Vaccination struct {
	IdVaccination    int    `json:"IdVaccination"`
	IdPerson         int    `json:"idPerson"`
	IdVaccine        int    `json:"idVaccine"`
	IdQuestion       int 	`json:"idQuestion"`
	IdAnswer         int 	`json:"idAnswer"`
	Applicator       string `json:"applicator"`
	MinTemp      	 int    `json:"minTemp"`
	MaxTemp      	 int    `json:"maxTemp"`
	ExpirityInDays   int    `json:"expirityInDays"`
	Facility         string `json:"facility"`
	Dose      		 int    `json:"dose"`
	Local      		 string `json:"local"`
	Lot      		 int    `json:"lot"`
	ExpirationDate   string `json:"expirationDate"`
}

func (s *VaccinationSmartContract) InitLedger(ctx contractapi.TransactionContextInterface) error {
	vaccinations := []Vaccination{
		{IdVaccination: 58973498, IdPerson: 25679788, IdVaccine: 84656347, IdQuestion: 35267896, IdAnswer: 65678765, Applicator: "facility A", MinTemp: -100, MaxTemp: -30, ExpirityInDays: 120, Facility: "hospital A", Dose: 1, Local: "nyc", Lot: 64565, ExpirationDate: "2021-07-25"},
		{IdVaccination: 43578459, IdPerson: 27253727, IdVaccine: 87535904, IdQuestion: 27864589, IdAnswer: 17895642, Applicator: "facility B", MinTemp: -200, MaxTemp: -30, ExpirityInDays: 120, Facility: "hospital B", Dose: 1, Local: "nyc", Lot: 34537, ExpirationDate: "2021-09-19"},
		{IdVaccination: 78392877, IdPerson: 53227256, IdVaccine: 84356757, IdQuestion: 35267896, IdAnswer: 65678765, Applicator: "facility A", MinTemp: -200, MaxTemp: -30, ExpirityInDays: 120, Facility: "hospital A", Dose: 1, Local: "nyc", Lot: 64565, ExpirationDate: "2021-07-25"},
		{IdVaccination: 76638393, IdPerson: 52926438, IdVaccine: 67545456, IdQuestion: 27864589, IdAnswer: 17895642, Applicator: "facility B", MinTemp: -200, MaxTemp: -70, ExpirityInDays: 120, Facility: "hospital B", Dose: 1, Local: "nyc", Lot: 34537, ExpirationDate: "2021-09-19"},
		{IdVaccination: 97827828, IdPerson: 82624721, IdVaccine: 23454357, IdQuestion: 35267896, IdAnswer: 65678765, Applicator: "facility A", MinTemp: -200, MaxTemp: -30, ExpirityInDays: 120, Facility: "hospital A", Dose: 1, Local: "nyc", Lot: 64565, ExpirationDate: "2021-07-25"},
		{IdVaccination: 98872672, IdPerson: 72925279, IdVaccine: 79653432, IdQuestion: 27864589, IdAnswer: 17895642, Applicator: "facility B", MinTemp: -100, MaxTemp: -20, ExpirityInDays: 120, Facility: "hospital B", Dose: 1, Local: "nyc", Lot: 34537, ExpirationDate: "2021-09-19"},
		{IdVaccination: 28326729, IdPerson: 96553445, IdVaccine: 67545456, IdQuestion: 35267896, IdAnswer: 65678765, Applicator: "facility A", MinTemp: -100, MaxTemp: -30, ExpirityInDays: 120, Facility: "hospital A", Dose: 1, Local: "nyc", Lot: 64565, ExpirationDate: "2021-07-25"},
		{IdVaccination: 78302624, IdPerson: 64546073, IdVaccine: 84356757, IdQuestion: 27864589, IdAnswer: 17895642, Applicator: "facility B", MinTemp: -200, MaxTemp: -20, ExpirityInDays: 120, Facility: "hospital B", Dose: 1, Local: "nyc", Lot: 34537, ExpirationDate: "2021-09-19"},
		{IdVaccination: 10287367, IdPerson: 57854378, IdVaccine: 87535904, IdQuestion: 35267896, IdAnswer: 65678765, Applicator: "facility A", MinTemp: -100, MaxTemp: -20, ExpirityInDays: 120, Facility: "hospital A", Dose: 1, Local: "nyc", Lot: 64565, ExpirationDate: "2021-07-25"},
		{IdVaccination: 87352892, IdPerson: 49557886, IdVaccine: 84656347, IdQuestion: 27864589, IdAnswer: 17895642, Applicator: "facility B", MinTemp: -200, MaxTemp: -70, ExpirityInDays: 120, Facility: "hospital B", Dose: 1, Local: "nyc", Lot: 34537, ExpirationDate: "2021-09-19"},
		
	}
	for _, vaccination := range vaccinations {
		vaccinationJSON, err := json.Marshal(vaccination)
		if err != nil {
			return err
		}

		err = ctx.GetStub().PutState(strconv.Itoa(vaccination.IdVaccination), vaccinationJSON)
		if err != nil {
			return fmt.Errorf("failed to put vaccination to world state. %v", err)
		}
	}
	return nil
}

func (s *VaccinationSmartContract) CreateVaccination(ctx contractapi.TransactionContextInterface, id int, idpe int, idvc int, idqu int, idan int, ap string, minT int, maxT int, expD int, fa string, do int, loc string, lot int, expDa string) error {
	exists, err := s.VaccinationExists(ctx, id)
	if err != nil {
		return err
	}
	if exists {
		return fmt.Errorf("the Vaccination %d already exists", id)
	}
	vaccination := Vaccination{
		IdVaccination:    id,
		IdPerson:		  idpe,		
		IdVaccine:		  idvc,
		IdQuestion:       idqu,
		IdAnswer:         idan,
		Applicator:       ap,
		MinTemp:          minT,
		MaxTemp:          maxT,
		ExpirityInDays:   expD,
		Facility:         fa,
		Dose:             do,
		Local:            loc,
		Lot:              lot,
		ExpirationDate:   expDa,
	}

	vaccinationJSON, err := json.Marshal(vaccination)
	if err != nil {
		return err
	}

	return ctx.GetStub().PutState(strconv.Itoa(id), vaccinationJSON)
}

func (s *VaccinationSmartContract) ReadVaccination(ctx contractapi.TransactionContextInterface, id int) (*Vaccination, error) {

	vaccinationJSON, err := ctx.GetStub().GetState(strconv.Itoa(id))
	if err != nil {
		return nil, fmt.Errorf("failed to read vaccination from world state: %v", err)
	}
	if vaccinationJSON == nil {
		return nil, fmt.Errorf("the vaccination %d does not exist", id)
	}

	var vaccination Vaccination
	err = json.Unmarshal(vaccinationJSON, &vaccination)
	if err != nil {
		return nil, err
	}

	return &vaccination, nil
}

func (s *VaccinationSmartContract) UpdateVaccination(ctx contractapi.TransactionContextInterface, id int, idpe int, idvc int, idqu int, idan int, ap string, minT int, maxT int, expD int, fa string, do int, loc string, lot int, expDa string) error {
	exists, err := s.VaccinationExists(ctx, id)

	if err != nil {
		return err
	}
	if !exists {
		return fmt.Errorf("the vaccination %d does not exist", id)
	}

	vaccination := Vaccination{
		IdVaccination:    id,
		IdPerson:		  idpe,		
		IdVaccine:		  idvc,
		IdQuestion:       idqu,
		IdAnswer:         idan,
		Applicator:       ap,
		MinTemp:          minT,
		MaxTemp:          maxT,
		ExpirityInDays:   expD,
		Facility:         fa,
		Dose:             do,
		Local:            loc,
		Lot:              lot,
		ExpirationDate:   expDa,
	}
	vaccinationJSON, err := json.Marshal(vaccination)
	if err != nil {
		return err
	}

	return ctx.GetStub().PutState(strconv.Itoa(id), vaccinationJSON)
}

func (s *VaccinationSmartContract) DeleteVaccination(ctx contractapi.TransactionContextInterface, id int) error {
	exists, err := s.VaccinationExists(ctx, id)
	if err != nil {
		return err
	}
	if !exists {
		return fmt.Errorf("the vaccination %d does not exist", id)
	}

	return ctx.GetStub().DelState(strconv.Itoa(id))
}

func (s *VaccinationSmartContract) VaccinationExists(ctx contractapi.TransactionContextInterface, id int) (bool, error) {
	vaccinationJSON, err := ctx.GetStub().GetState(strconv.Itoa(id))
	if err != nil {
		return false, fmt.Errorf("failed to read vaccination from world state: %v", err)
	}

	return vaccinationJSON != nil, nil
}

// GetAllVaccinations returns all cars found in world state
func (s *VaccinationSmartContract) GetAllVaccinations(ctx contractapi.TransactionContextInterface) ([]*Vaccination, error) {
	resultsIterator, err := ctx.GetStub().GetStateByRange("", "")
	if err != nil {
		return nil, err
	}
	defer resultsIterator.Close()

	var vaccinations []*Vaccination
	for resultsIterator.HasNext() {
		queryResponse, err := resultsIterator.Next()
		if err != nil {
			return nil, err
		}

		var vaccination Vaccination
		err = json.Unmarshal(queryResponse.Value, &vaccination)
		if err != nil {
			return nil, err
		}
		vaccinations = append(vaccinations, &vaccination)
	}

	return vaccinations, nil
}

func main() {

	assetChaincode, err := contractapi.NewChaincode(&VaccinationSmartContract{})
	if err != nil {
		log.Panicf("Error creating vaccination chaincode: %v", err)
	}

	if err := assetChaincode.Start(); err != nil {
		log.Panicf("Error starting vaccination chaincode: %v", err)
	}
}
