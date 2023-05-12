package main

import (
	"encoding/json"
	"fmt"
	"strconv"
	"strings"

	"github.com/hyperledger/fabric/core/chaincode/shim"
	pb "github.com/hyperledger/fabric/protos/peer"
)

func (t *Logger) monitorAccessPublic(stub shim.ChaincodeStubInterface, args []string) pb.Response {
	var err error
	var jsonResp string
	if len(args) != 5 {
		return shim.Error("Incorrect arguments. Clearing House for Public level access history expecting 5 values")
	}
	dataset := args[0]
	usernameOfProvider := args[1]
	permission := args[2]
	timestamp := args[3]
	customAccessRights := strings.Split(args[4], ",")
	objectType := "LoggerAccessPublic"
	objectTypeKey := "LoggerAccessPublicKey"

	if len(dataset) <= 0 {
		return shim.Error("Dataset must be a non-empty string")
	}

	if len(usernameOfProvider) <= 0 {
		return shim.Error("Username of provider must be a non-empty string")
	}

	monitorAccessPublicKey := "accessKey" + "/" + dataset + "/" + usernameOfProvider
	loggerAccessPublicKeyJSON := LoggerAccessPublicCounter{}
	count := -1
	loggerAccesPublicKeyAsbytes, err := stub.GetState(monitorAccessPublicKey) //get the public access from chaincode state
	if err != nil {
		jsonResp = "{\"Error\":\"Failed to get access history for " + monitorAccessPublicKey + "\"}"
		return shim.Error(jsonResp)
	} else if loggerAccesPublicKeyAsbytes == nil {
		count = 0
		loggerAccessPublicKey := &LoggerRevokedAccessCounterPublic{objectTypeKey, dataset, count}
		loggerAccesPublicKeyAsbytes, err = json.Marshal(loggerAccessPublicKey)
		if err != nil {
			return shim.Error(err.Error())
		}

	} else {
		err = json.Unmarshal(loggerAccesPublicKeyAsbytes, &loggerAccessPublicKeyJSON) //unmarshal it aka JSON.parse()
		if err != nil {
			return shim.Error(err.Error())
		}
		count = loggerAccessPublicKeyJSON.Count + 1
		loggerAccessPublicKeyJSON.Count = count
		loggerAccesPublicKeyAsbytes, _ = json.Marshal(loggerAccessPublicKeyJSON)
	}
	// === Save  key for public access to state ===
	err = stub.PutState(monitorAccessPublicKey, loggerAccesPublicKeyAsbytes)
	if err != nil {
		return shim.Error(err.Error())
	}

	loggerAccessPublic := &LoggerAccessPublic{objectType, dataset, usernameOfProvider, permission, timestamp, customAccessRights, count}
	loggerAccessPublicJSONasBytes, err := json.Marshal(loggerAccessPublic)
	if err != nil {
		return shim.Error(err.Error())
	}

	monitorAccessPublic := "access" + "/" + dataset + "/" + usernameOfProvider + "/" + strconv.Itoa(count)
	// === Save to state ===
	err = stub.PutState(monitorAccessPublic, loggerAccessPublicJSONasBytes)
	if err != nil {
		return shim.Error(err.Error())
	}

	messageofdata := "Clearing House public access for dataset: " + dataset + "provider: " + usernameOfProvider + "permission: " + permission
	bytemessage := []byte(messageofdata)
	return shim.Success(bytemessage)

}

func (t *Logger) monitorRevokedAccessPublic(stub shim.ChaincodeStubInterface, args []string) pb.Response {
	var jsonResp string
	var err error
	if len(args) != 5 {
		return shim.Error("Incorrect arguments. Logger SC for monitoring revoked access for public access expecting 5 values")
	}
	dataset := args[0]
	usernameOfProvider := args[1]
	revokedPermission := args[2]
	timestamp := args[3]
	customAccessRights := strings.Split(args[4], ",")
	objectType := "LoggerRevokedAccessPublic"
	objectTypeKey := "LoggerRevokedAccessPublicKey"

	if len(dataset) <= 0 {
		return shim.Error("Dataset must be a non-empty string")
	}

	if len(usernameOfProvider) <= 0 {
		return shim.Error("Username of provider must be a non-empty string")
	}

	monitorRevokedAccessCounterPublic := "revokedKey" + "/" + dataset + "/" + usernameOfProvider
	loggerRevokedAccessPublicKeyJSON := LoggerRevokedAccessCounterPublic{}
	count := -1
	loggerAccesPublicsAsbytes, err := stub.GetState(monitorRevokedAccessCounterPublic) //get the revoked access from chaincode state
	if err != nil {
		jsonResp = "{\"Error\":\"Failed to get revoked access history for " + monitorRevokedAccessCounterPublic + "\"}"
		return shim.Error(jsonResp)
	} else if loggerAccesPublicsAsbytes == nil {
		count = 0
		loggerAccessPublic := &LoggerRevokedAccessCounterPublic{objectTypeKey, dataset, count}
		loggerAccesPublicsAsbytes, err = json.Marshal(loggerAccessPublic)
		if err != nil {
			return shim.Error(err.Error())
		}

	} else {
		err = json.Unmarshal(loggerAccesPublicsAsbytes, &loggerRevokedAccessPublicKeyJSON) //unmarshal it aka JSON.parse()
		if err != nil {
			return shim.Error(err.Error())
		}
		count = loggerRevokedAccessPublicKeyJSON.Count + 1
		loggerRevokedAccessPublicKeyJSON.Count = count
		loggerAccesPublicsAsbytes, _ = json.Marshal(loggerRevokedAccessPublicKeyJSON)
	}

	// === Save  for revoked access to state ===
	err = stub.PutState(monitorRevokedAccessCounterPublic, loggerAccesPublicsAsbytes)
	if err != nil {
		return shim.Error(err.Error())
	}

	revokedAccess := &LoggerRevokedAccessPublic{objectType, dataset, usernameOfProvider, timestamp, revokedPermission, customAccessRights, count}
	revokedAccessasBytes, err := json.Marshal(revokedAccess)
	if err != nil {
		return shim.Error(err.Error())
	}
	monitorAccess := "revoked" + "/" + dataset + "/" + usernameOfProvider + "/" + strconv.Itoa(count)
	// === Save revoked access to state ===
	err = stub.PutState(monitorAccess, revokedAccessasBytes)
	if err != nil {
		return shim.Error(err.Error())
	}

	messageofRevokedAccess := "History of revoked public access for dataset: " + dataset + "provider: " + usernameOfProvider + "count: " + strconv.Itoa(count)
	bytemessage := []byte(messageofRevokedAccess)
	return shim.Success(bytemessage)

}

func (t *Logger) queryLoggerRevokedAccessPublicProvider(stub shim.ChaincodeStubInterface) pb.Response {
	usernameOfProvider, err := t.getUsername(stub)
	if err != nil {
		return shim.Error(err.Error())
	}
	//create index key
	queryString := fmt.Sprintf("{\"selector\":{\"docType\":\"LoggerRevokedAccessPublic\",\"usernameOfProvider\":\"%s\"}}", usernameOfProvider)
	queryResults, err := getQueryResultForQueryString(stub, queryString)
	if err != nil {
		return shim.Error(err.Error())
	}
	return shim.Success(queryResults)
}

func (t *Logger) queryLoggerRevokedAccessPublicProviderByDatasetName(stub shim.ChaincodeStubInterface, args []string) pb.Response {
	if len(args) != 1 {
		return shim.Error("Invalid Argument Number. Expecting dataset name")
	}
	datasetname := args[0]
	usernameOfProvider, err := t.getUsername(stub)
	if err != nil {
		return shim.Error(err.Error())
	}
	//create index key
	queryString := fmt.Sprintf("{\"selector\":{\"docType\":\"LoggerRevokedAccessPublic\",\"usernameOfProvider\":\"%s\",\"dataset\":\"%s\"}}", usernameOfProvider, datasetname)
	queryResults, err := getQueryResultForQueryString(stub, queryString)
	if err != nil {
		return shim.Error(err.Error())
	}
	return shim.Success(queryResults)
}

func (t *Logger) queryLoggerAllRevokedAccessPublic(stub shim.ChaincodeStubInterface) pb.Response {
	//create index key
	queryString := "{\"selector\":{\"docType\":\"LoggerRevokedAccessPublic\"}}"
	queryResults, err := getQueryResultForQueryString(stub, queryString)
	if err != nil {
		return shim.Error(err.Error())
	}
	return shim.Success(queryResults)
}

func (t *Logger) queryLoggerAllRevokedAccessPublicByDatasetName(stub shim.ChaincodeStubInterface, args []string) pb.Response {
	if len(args) != 1 {
		return shim.Error("Invalid Argument Number. Expecting dataset name")
	}
	datasetname := args[0]
	//create index key
	queryString := fmt.Sprintf("{\"selector\":{\"docType\":\"LoggerRevokedAccessPublic\",\"dataset\":\"%s\"}}", datasetname)
	queryResults, err := getQueryResultForQueryString(stub, queryString)
	if err != nil {
		return shim.Error(err.Error())
	}
	return shim.Success(queryResults)
}

func (t *Logger) queryLoggerAllAccessPublic(stub shim.ChaincodeStubInterface) pb.Response {
	//create index key
	queryString := "{\"selector\":{\"docType\":\"LoggerAccessPublic\"}}"
	queryResults, err := getQueryResultForQueryString(stub, queryString)
	if err != nil {
		return shim.Error(err.Error())
	}
	return shim.Success(queryResults)
}

func (t *Logger) queryLoggerAllAccessPublicByDatasetName(stub shim.ChaincodeStubInterface, args []string) pb.Response {
	if len(args) != 1 {
		return shim.Error("Invalid Argument Number. Expecting dataset name")
	}
	datasetname := args[0]
	//create index key
	queryString := fmt.Sprintf("{\"selector\":{\"docType\":\"LoggerAccessPublic\",\"dataset\":\"%s\"}}", datasetname)
	queryResults, err := getQueryResultForQueryString(stub, queryString)
	if err != nil {
		return shim.Error(err.Error())
	}
	return shim.Success(queryResults)
}

func (t *Logger) queryLoggerAllAccessPublicProvider(stub shim.ChaincodeStubInterface) pb.Response {
	usernameOfProvider, err := t.getUsername(stub)
	if err != nil {
		return shim.Error(err.Error())
	}
	//create index key
	queryString := fmt.Sprintf("{\"selector\":{\"docType\":\"LoggerAccessPublic\",\"usernameOfProvider\":\"%s\"}}", usernameOfProvider)
	queryResults, err := getQueryResultForQueryString(stub, queryString)
	if err != nil {
		return shim.Error(err.Error())
	}
	return shim.Success(queryResults)
}

func (t *Logger) queryLoggerAllAccessPublicProviderByDatasetName(stub shim.ChaincodeStubInterface, args []string) pb.Response {
	if len(args) != 1 {
		return shim.Error("Invalid Argument Number. Expecting dataset name")
	}
	datasetname := args[0]
	usernameOfProvider, err := t.getUsername(stub)
	if err != nil {
		return shim.Error(err.Error())
	}
	//create index key
	queryString := fmt.Sprintf("{\"selector\":{\"docType\":\"LoggerAccessPublic\",\"usernameOfProvider\":\"%s\",\"dataset\":\"%s\"}}", usernameOfProvider, datasetname)
	queryResults, err := getQueryResultForQueryString(stub, queryString)
	if err != nil {
		return shim.Error(err.Error())
	}
	return shim.Success(queryResults)
}
