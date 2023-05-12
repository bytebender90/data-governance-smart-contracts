package main

import (
	"encoding/json"

	"github.com/hyperledger/fabric/core/chaincode/shim"
	pb "github.com/hyperledger/fabric/protos/peer"
)

func (t *Logger) monitorUploadMetadata(stub shim.ChaincodeStubInterface, args []string) pb.Response {
	if len(args) != 7 {
		return shim.Error("Incorrect arguments. Logger expecting 7 arguments")
	}
	dataset := args[0]
	username := args[1]
	organization := args[2]
	changes := []string{args[3]}
	version := args[4]
	comments := args[5]
	datetime := args[6]
	objectType := "LoggerMetadata"

	if len(dataset) <= 0 {
		return shim.Error("Dataset must be a non-empty string")
	}

	if len(organization) <= 0 {
		return shim.Error("Organization must be a non-empty string")
	}

	if len(username) <= 0 {
		return shim.Error("Username must be a non-empty string")
	}

	LoggerMetadata := &LoggerMetadata{objectType, dataset, username, organization, changes, comments, datetime}
	LoggerMetadataJSONasBytes, err := json.Marshal(LoggerMetadata)
	if err != nil {
		return shim.Error(err.Error())
	}

	keyMetadataChanges := dataset + version

	// === Save to state ===
	err = stub.PutState(keyMetadataChanges, LoggerMetadataJSONasBytes)
	if err != nil {
		return shim.Error(err.Error())
	}

	messageofdata := "Clearing House metadata for dataset: " + dataset
	bytemessage := []byte(messageofdata)
	return shim.Success(bytemessage)

}

func (t *Logger) queryLoggerMetadata(stub shim.ChaincodeStubInterface) pb.Response {
	//create index key
	queryString := ("{\"selector\":{\"docType\":\"LoggerMetadata\"}}")
	queryResults, err := getQueryResultForQueryString(stub, queryString)
	if err != nil {
		return shim.Error(err.Error())
	}
	return shim.Success(queryResults)
}
