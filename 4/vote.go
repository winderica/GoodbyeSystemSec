package main

import (
	"bytes"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"math/rand"

	"github.com/hyperledger/fabric/core/chaincode/shim"
	sc "github.com/hyperledger/fabric/protos/peer"
)

type SmartContract struct {
}

type Candidate struct {
	Id        string `json:"id"`
	Name      string `json:"name"`
	Delegates uint   `json:"delegates"`
}

func getId() string {
	b := make([]byte, 8)
	rand.Read(b)
	return hex.EncodeToString(b)
}

func (s *SmartContract) Init(APIstub shim.ChaincodeStubInterface) sc.Response {
	return shim.Success(nil)
}

func (s *SmartContract) Invoke(APIstub shim.ChaincodeStubInterface) sc.Response {
	function, args := APIstub.GetFunctionAndParameters()
	if function == "queryCandidate" {
		return s.queryCandidate(APIstub, args)
	} else if function == "initLedger" {
		return s.initLedger(APIstub)
	} else if function == "addCandidate" {
		return s.addCandidate(APIstub, args)
	} else if function == "queryAllCandidates" {
		return s.queryAllCandidates(APIstub)
	} else if function == "voteForCandidate" {
		return s.voteForCandidate(APIstub, args)
	}
	return shim.Error("Invalid Smart Contract function name.")
}

func (s *SmartContract) queryCandidate(APIstub shim.ChaincodeStubInterface, args []string) sc.Response {
	if len(args) != 1 {
		return shim.Error("Incorrect number of arguments. Expecting 1")
	}
	candidateAsBytes, _ := APIstub.GetState(args[0])
	return shim.Success(candidateAsBytes)
}

func (s *SmartContract) initLedger(APIstub shim.ChaincodeStubInterface) sc.Response {
	candidates := []Candidate{
		Candidate{Id: getId(), Name: "Joe Biden", Delegates: 0},
		Candidate{Id: getId(), Name: "Bernie Sanders", Delegates: 0},
		Candidate{Id: getId(), Name: "Elizabeth Warren", Delegates: 0},
		Candidate{Id: getId(), Name: "Michael Bloomberg", Delegates: 0},
		Candidate{Id: getId(), Name: "Pete Buttigieg", Delegates: 0},
		Candidate{Id: getId(), Name: "Amy Klobuchar", Delegates: 0},
		Candidate{Id: getId(), Name: "Tulsi Gabbard", Delegates: 0},
		Candidate{Id: getId(), Name: "Andrew Yang", Delegates: 0},
		Candidate{Id: getId(), Name: "Deval Patrick", Delegates: 0},
		Candidate{Id: getId(), Name: "John Delaney", Delegates: 0},
		Candidate{Id: getId(), Name: "Michael Bennet", Delegates: 0},
		Candidate{Id: getId(), Name: "Tom Steyer", Delegates: 0},
	}
	i := 0
	for i < len(candidates) {
		candidateAsBytes, _ := json.Marshal(candidates[i])
		APIstub.PutState(candidates[i].Id, candidateAsBytes)
		i = i + 1
	}
	return shim.Success(nil)
}

func (s *SmartContract) addCandidate(APIstub shim.ChaincodeStubInterface, args []string) sc.Response {
	if len(args) != 1 {
		return shim.Error("Incorrect number of arguments. Expecting 1")
	}
	var candidate = Candidate{Id: getId(), Name: args[0], Delegates: 0}

	candidateAsBytes, _ := json.Marshal(candidate)
	APIstub.PutState(candidate.Id, candidateAsBytes)

	return shim.Success(nil)
}

func (s *SmartContract) queryAllCandidates(APIstub shim.ChaincodeStubInterface) sc.Response {
	resultsIterator, err := APIstub.GetStateByRange("", "")
	if err != nil {
		return shim.Error(err.Error())
	}
	defer resultsIterator.Close()

	var buffer bytes.Buffer
	buffer.WriteString("[")

	bArrayMemberAlreadyWritten := false
	for resultsIterator.HasNext() {
		queryResponse, err := resultsIterator.Next()
		if err != nil {
			return shim.Error(err.Error())
		}
		if bArrayMemberAlreadyWritten == true {
			buffer.WriteString(",")
		}
		buffer.WriteString("{\"Key\":")
		buffer.WriteString("\"")
		buffer.WriteString(queryResponse.Key)
		buffer.WriteString("\"")
		buffer.WriteString(", \"Record\":")
		buffer.WriteString(string(queryResponse.Value))
		buffer.WriteString("}")
		bArrayMemberAlreadyWritten = true
	}
	buffer.WriteString("]")

	return shim.Success(buffer.Bytes())
}

func (s *SmartContract) voteForCandidate(APIstub shim.ChaincodeStubInterface, args []string) sc.Response {
	if len(args) != 1 {
		return shim.Error("Incorrect number of arguments. Expecting 1")
	}
	candidateAsBytes, _ := APIstub.GetState(args[0])
	candidate := Candidate{}

	json.Unmarshal(candidateAsBytes, &candidate)
	candidate.Delegates++

	candidateAsBytes, _ = json.Marshal(candidate)
	APIstub.PutState(args[0], candidateAsBytes)

	return shim.Success(nil)
}

func main() {
	err := shim.Start(new(SmartContract))
	if err != nil {
		fmt.Printf("Error creating new Smart Contract: %s", err)
	}
}
