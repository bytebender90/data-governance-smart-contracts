package main

import (
	"encoding/json"
	"fmt"
	"strconv"
	"strings"

	"github.com/hyperledger/fabric/core/chaincode/shim"
	pb "github.com/hyperledger/fabric/protos/peer"
)

func (t *Logger) monitorRevokedAccessUser(stub shim.ChaincodeStubInterface, args []string) pb.Response {
	var jsonResp string
	var err error
	if len(args) != 9 {
		return shim.Error("Incorrect arguments. Logger SC for monitoring revoked access expecting 9 values")
	}

	dataset := args[0]
	usernameOfProvider := args[1]
	usernameOfConsumer := args[2]
	organization := args[3]
	revokedPermission := args[4]
	timestamp := args[5]
	customAccessRights := strings.Split(args[6], ",")
	orgOfProvider := args[7]
	description := args[8]
	objectType := "LoggerRevokedAccess"
	objectTypeKey := "LoggerRevokedAccessKey"

	if len(dataset) <= 0 {
		return shim.Error("Dataset must be a non-empty string")
	}

	if len(organization) <= 0 {
		return shim.Error("Organization must be a non-empty string")
	}

	if len(usernameOfProvider) <= 0 {
		return shim.Error("Username of provider must be a non-empty string")
	}

	if len(usernameOfConsumer) <= 0 {
		return shim.Error("Username of consumer must be a non-empty string")
	}

	monitorRevokedAccessKeyCounter := "revokedUserKEY" + dataset + "/" + usernameOfProvider + "/" + usernameOfConsumer
	KeyLoggerRevokedAccessJSON := LoggerRevokedAccessCounter{}
	count := -1
	loggerAccessAsbytes, err := stub.GetState(monitorRevokedAccessKeyCounter) //get the revoked access from chaincode state
	if err != nil {
		jsonResp = "{\"Error\":\"Failed to get revoked access history for " + monitorRevokedAccessKeyCounter + "\"}"
		return shim.Error(jsonResp)
	} else if loggerAccessAsbytes == nil {
		count = 0
		loggerAccess := &LoggerRevokedAccessCounter{objectTypeKey, dataset, count}
		loggerAccessAsbytes, err = json.Marshal(loggerAccess)
		if err != nil {
			return shim.Error(err.Error())
		}

	} else {
		err = json.Unmarshal(loggerAccessAsbytes, &KeyLoggerRevokedAccessJSON) //unmarshal it aka JSON.parse()
		if err != nil {
			return shim.Error(err.Error())
		}
		count = KeyLoggerRevokedAccessJSON.Count + 1
		KeyLoggerRevokedAccessJSON.Count = count
		loggerAccessAsbytes, _ = json.Marshal(KeyLoggerRevokedAccessJSON)
	}

	// === Save key for revoked access to state ===
	err = stub.PutState(monitorRevokedAccessKeyCounter, loggerAccessAsbytes)
	if err != nil {
		return shim.Error(err.Error())
	}

	revokedAccess := &LoggerRevokedAccess{objectType, dataset, usernameOfProvider, orgOfProvider, description, usernameOfConsumer, organization, timestamp, revokedPermission, customAccessRights, count}
	revokedAccessasBytes, err := json.Marshal(revokedAccess)
	if err != nil {
		return shim.Error(err.Error())
	}
	monitorAccessKey := "revokedUser" + dataset + "/" + usernameOfProvider + "/" + usernameOfConsumer + "/" + strconv.Itoa(count)
	// === Save revoked access to state ===
	err = stub.PutState(monitorAccessKey, revokedAccessasBytes)
	if err != nil {
		return shim.Error(err.Error())
	}

	messageofRevokedAccess := "History of revoked access for dataset: " + dataset + "provider: " + usernameOfProvider + "consumer: " + usernameOfProvider + "count: " + strconv.Itoa(count)
	bytemessage := []byte(messageofRevokedAccess)
	return shim.Success(bytemessage)

}

func (t *Logger) queryLoggerRevokedAccessProvider(stub shim.ChaincodeStubInterface) pb.Response {
	usernameOfProvider, err := t.getUsername(stub)
	if err != nil {
		return shim.Error(err.Error())
	}
	//create index key
	queryString := fmt.Sprintf("{\"selector\":{\"docType\":\"LoggerRevokedAccess\",\"usernameOfProvider\":\"%s\"}}", usernameOfProvider)
	queryResults, err := getQueryResultForQueryString(stub, queryString)
	if err != nil {
		return shim.Error(err.Error())
	}
	return shim.Success(queryResults)
}

func (t *Logger) queryLoggerRevokedAccessConsumer(stub shim.ChaincodeStubInterface) pb.Response {
	usernameOfConsumer, err := t.getUsername(stub)
	if err != nil {
		return shim.Error(err.Error())
	}
	//create index key
	queryString := fmt.Sprintf("{\"selector\":{\"docType\":\"LoggerRevokedAccess\",\"usernameOfConsumer\":\"%s\"}}", usernameOfConsumer)
	queryResults, err := getQueryResultForQueryString(stub, queryString)
	if err != nil {
		return shim.Error(err.Error())
	}
	return shim.Success(queryResults)
}

func (t *Logger) queryLoggerRevokedAccessProviderByDatasetName(stub shim.ChaincodeStubInterface, args []string) pb.Response {
	if len(args) != 1 {
		return shim.Error("Invalid Argument Number. Expecting dataset name")
	}
	datasetname := args[0]
	usernameOfProvider, err := t.getUsername(stub)
	if err != nil {
		return shim.Error(err.Error())
	}
	//create index key
	queryString := fmt.Sprintf("{\"selector\":{\"docType\":\"LoggerRevokedAccess\",\"usernameOfProvider\":\"%s\",\"dataset\":\"%s\"}}", usernameOfProvider, datasetname)
	queryResults, err := getQueryResultForQueryString(stub, queryString)
	if err != nil {
		return shim.Error(err.Error())
	}
	return shim.Success(queryResults)
}

func (t *Logger) queryLoggerRevokedAccessConsumerByDatasetName(stub shim.ChaincodeStubInterface, args []string) pb.Response {
	if len(args) != 1 {
		return shim.Error("Invalid Argument Number. Expecting dataset name")
	}
	datasetname := args[0]
	usernameOfConsumer, err := t.getUsername(stub)
	if err != nil {
		return shim.Error(err.Error())
	}
	//create index key
	queryString := fmt.Sprintf("{\"selector\":{\"docType\":\"LoggerRevokedAccess\",\"usernameOfConsumer\":\"%s\",\"dataset\":\"%s\"}}", usernameOfConsumer, datasetname)
	queryResults, err := getQueryResultForQueryString(stub, queryString)
	if err != nil {
		return shim.Error(err.Error())
	}
	return shim.Success(queryResults)
}

func (t *Logger) monitorAccessRequest(stub shim.ChaincodeStubInterface, args []string) pb.Response {
	var err error
	var jsonResp string
	if len(args) != 11 {
		return shim.Error("Incorrect arguments. Logger SC for monitoring access requests expecting 11 values")
	}
	dataset := args[0]
	usernameOfProvider := args[1]
	usernameOfConsumer := args[2]
	organization := args[3]
	requestID := args[4]
	status := args[5]
	timestamp := args[6]
	permission := args[7]
	customAccessRights := strings.Split(args[8], ",")
	orgOfProvider := args[9]
	description := args[10]
	objectType := "LoggerAccess"
	objectTypeKey := "LoggerAccessKey"

	if len(dataset) <= 0 {
		return shim.Error("Dataset must be a non-empty string")
	}

	if len(organization) <= 0 {
		return shim.Error("Organization must be a non-empty string")
	}

	if len(usernameOfProvider) <= 0 {
		return shim.Error("Username of provider must be a non-empty string")
	}

	if len(usernameOfConsumer) <= 0 {
		return shim.Error("Username of consumer must be a non-empty string")
	}

	if len(status) <= 0 {
		return shim.Error("Status must be a non-empty string")
	}

	monitorAccessKeyCounter := "accessUserKEY" + dataset + "/" + usernameOfProvider + "/" + usernameOfConsumer
	KeyLoggerAccessJSON := LoggerAccessCounter{}
	count := -1
	loggerAccessKeyAsbytes, err := stub.GetState(monitorAccessKeyCounter) //get the revoked access from chaincode state
	if err != nil {
		jsonResp = "{\"Error\":\"Failed to get access history for " + monitorAccessKeyCounter + "\"}"
		return shim.Error(jsonResp)
	} else if loggerAccessKeyAsbytes == nil {
		count = 0
		loggerAccess := &LoggerAccessCounter{objectTypeKey, dataset, count}
		loggerAccessKeyAsbytes, err = json.Marshal(loggerAccess)
		if err != nil {
			return shim.Error(err.Error())
		}

	} else {
		err = json.Unmarshal(loggerAccessKeyAsbytes, &KeyLoggerAccessJSON) //unmarshal it aka JSON.parse()
		if err != nil {
			return shim.Error(err.Error())
		}
		count = KeyLoggerAccessJSON.Count + 1
		KeyLoggerAccessJSON.Count = count
		loggerAccessKeyAsbytes, _ = json.Marshal(KeyLoggerAccessJSON)
	}

	// === Save key for access to state ===
	err = stub.PutState(monitorAccessKeyCounter, loggerAccessKeyAsbytes)
	if err != nil {
		return shim.Error(err.Error())
	}

	LoggerAccess := &LoggerAccess{objectType, dataset, usernameOfProvider, orgOfProvider, description, usernameOfConsumer, organization, timestamp, permission, customAccessRights, status, requestID, count}
	LoggerAccessJSONasBytes, err := json.Marshal(LoggerAccess)
	if err != nil {
		return shim.Error(err.Error())
	}

	monitorAccessKey := "accessUser" + dataset + "/" + usernameOfProvider + "/" + usernameOfConsumer + strconv.Itoa(count)
	// === Save to state ===
	err = stub.PutState(monitorAccessKey, LoggerAccessJSONasBytes)
	if err != nil {
		return shim.Error(err.Error())
	}

	messageofdata := "Clearing House user access for dataset: " + dataset
	bytemessage := []byte(messageofdata)
	return shim.Success(bytemessage)

}

func (t *Logger) queryLoggerAccessProvider(stub shim.ChaincodeStubInterface) pb.Response {
	usernameOfProvider, err := t.getUsername(stub)
	if err != nil {
		return shim.Error(err.Error())
	}
	//create index key
	queryString := fmt.Sprintf("{\"selector\":{\"docType\":\"LoggerAccess\",\"usernameOfProvider\":\"%s\"}}", usernameOfProvider)
	queryResults, err := getQueryResultForQueryString(stub, queryString)
	if err != nil {
		return shim.Error(err.Error())
	}
	return shim.Success(queryResults)
}

func (t *Logger) queryLoggerAccessConsumer(stub shim.ChaincodeStubInterface) pb.Response {
	usernameOfConsumer, err := t.getUsername(stub)
	if err != nil {
		return shim.Error(err.Error())
	}
	//create index key
	queryString := fmt.Sprintf("{\"selector\":{\"docType\":\"LoggerAccess\",\"usernameOfConsumer\":\"%s\"}}", usernameOfConsumer)
	queryResults, err := getQueryResultForQueryString(stub, queryString)
	if err != nil {
		return shim.Error(err.Error())
	}
	return shim.Success(queryResults)
}

func (t *Logger) queryLoggerAccessProviderByDatasetName(stub shim.ChaincodeStubInterface, args []string) pb.Response {
	if len(args) != 1 {
		return shim.Error("Invalid Argument Number. Expecting dataset name")
	}
	datasetname := args[0]
	usernameOfProvider, err := t.getUsername(stub)
	if err != nil {
		return shim.Error(err.Error())
	}
	//create index key
	queryString := fmt.Sprintf("{\"selector\":{\"docType\":\"LoggerAccess\",\"usernameOfProvider\":\"%s\",\"dataset\":\"%s\"}}", usernameOfProvider, datasetname)
	queryResults, err := getQueryResultForQueryString(stub, queryString)
	if err != nil {
		return shim.Error(err.Error())
	}
	return shim.Success(queryResults)
}

func (t *Logger) queryLoggerAccessConsumerByDatasetName(stub shim.ChaincodeStubInterface, args []string) pb.Response {
	if len(args) != 1 {
		return shim.Error("Invalid Argument Number. Expecting dataset name")
	}
	datasetname := args[0]
	usernameOfConsumer, err := t.getUsername(stub)
	if err != nil {
		return shim.Error(err.Error())
	}
	//create index key
	queryString := fmt.Sprintf("{\"selector\":{\"docType\":\"LoggerAccess\",\"usernameOfConsumer\":\"%s\",\"dataset\":\"%s\"}}", usernameOfConsumer, datasetname)
	queryResults, err := getQueryResultForQueryString(stub, queryString)
	if err != nil {
		return shim.Error(err.Error())
	}
	return shim.Success(queryResults)
}
