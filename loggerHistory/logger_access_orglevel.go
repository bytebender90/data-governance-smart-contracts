package main

import (
	"encoding/json"
	"fmt"
	"strconv"
	"strings"

	"github.com/hyperledger/fabric/core/chaincode/shim"
	pb "github.com/hyperledger/fabric/protos/peer"
)

func (t *Logger) monitorAccessRequestOrgs(stub shim.ChaincodeStubInterface, args []string) pb.Response {
	var err error
	var jsonResp string
	if len(args) != 14 {
		return shim.Error("Incorrect arguments. Function monitorAccessRequestOrgs expecting 14 values")
	}
	dataset := args[0]
	usernameOfProvider := args[1]
	organizationConsumer := args[2]
	requestID := args[3]
	status := args[4]
	timestamp := args[5]
	usernameConsumer := args[6]
	nameConsumer := args[7]
	surnameConsumer := args[8]
	roleConsumer := args[9]
	permission := args[10]
	customAccessRights := strings.Split(args[11], ",")
	orgOfProvider := args[12]
	description := args[13]
	objectType := "LoggerAccessOrgs"
	objectTypeKey := "LoggerAccessOrgsKey"

	if len(dataset) <= 0 {
		return shim.Error("Dataset must be a non-empty string")
	}

	if len(usernameOfProvider) <= 0 {
		return shim.Error("Username of provider must be a non-empty string")
	}

	if len(organizationConsumer) <= 0 {
		return shim.Error("Organization must be a non-empty string")
	}
	if len(status) <= 0 {
		return shim.Error("Status must be a non-empty string")
	}

	monitorAccessKeyCounterOrg := "accessOrgKey" + dataset + "/" + usernameOfProvider + "/" + organizationConsumer
	KeyLoggerAccessOrgJSON := LoggerAccessOrgsCounter{}
	count := -1
	loggerAccesOrgsAsbytes, err := stub.GetState(monitorAccessKeyCounterOrg) //get the access from chaincode state
	if err != nil {
		jsonResp = "{\"Error\":\"Failed to get access history for " + monitorAccessKeyCounterOrg + "\"}"
		return shim.Error(jsonResp)
	} else if loggerAccesOrgsAsbytes == nil {
		count = 0
		loggerAccessOrg := &LoggerRevokedAccessCounterOrg{objectTypeKey, dataset, count}
		loggerAccesOrgsAsbytes, err = json.Marshal(loggerAccessOrg)
		if err != nil {
			return shim.Error(err.Error())
		}

	} else {
		err = json.Unmarshal(loggerAccesOrgsAsbytes, &KeyLoggerAccessOrgJSON) //unmarshal it aka JSON.parse()
		if err != nil {
			return shim.Error(err.Error())
		}
		count = KeyLoggerAccessOrgJSON.Count + 1
		KeyLoggerAccessOrgJSON.Count = count
		loggerAccesOrgsAsbytes, _ = json.Marshal(KeyLoggerAccessOrgJSON)
	}

	// === Save key for revoked access to state ===
	err = stub.PutState(monitorAccessKeyCounterOrg, loggerAccesOrgsAsbytes)
	if err != nil {
		return shim.Error(err.Error())
	}

	LoggerAccess := &LoggerAccessOrgs{objectType, dataset, usernameOfProvider, orgOfProvider, description, organizationConsumer, timestamp, permission, customAccessRights, status, requestID, count,
		usernameConsumer, nameConsumer, surnameConsumer, roleConsumer}
	LoggerAccessJSONasBytes, err := json.Marshal(LoggerAccess)
	if err != nil {
		return shim.Error(err.Error())
	}

	monitorAccessKey := "accessOrg" + dataset + "/" + usernameOfProvider + organizationConsumer + "/" + strconv.Itoa(count)
	// === Save to state ===
	err = stub.PutState(monitorAccessKey, LoggerAccessJSONasBytes)
	if err != nil {
		return shim.Error(err.Error())
	}

	messageofdata := "Clearing House access for dataset: " + dataset
	bytemessage := []byte(messageofdata)
	return shim.Success(bytemessage)

}

func (t *Logger) monitorRevokedAccessOrg(stub shim.ChaincodeStubInterface, args []string) pb.Response {
	var jsonResp string
	var err error
	if len(args) != 12 {
		return shim.Error("Incorrect arguments. Function monitorRevokedAccessOrg expecting 12 values")
	}
	dataset := args[0]
	usernameOfProvider := args[1]
	organizationConsumer := args[2]
	revokedPermission := args[3]
	timestamp := args[4]
	usernameConsumer := args[5]
	nameConsumer := args[6]
	surnameConsumer := args[7]
	roleConsumer := args[8]
	customAccessRights := strings.Split(args[9], ",")
	orgOfProvider := args[10]
	description := args[11]
	objectType := "LoggerRevokedAccessOrgs"
	objectTypeKey := "LoggerRevokedAccessKeyOrgs"

	if len(dataset) <= 0 {
		return shim.Error("Dataset must be a non-empty string")
	}

	if len(organizationConsumer) <= 0 {
		return shim.Error("Organization must be a non-empty string")
	}

	if len(usernameOfProvider) <= 0 {
		return shim.Error("Username of provider must be a non-empty string")
	}

	monitorRevokedAccessKeyCounterOrg := "revokedOrgKey" + dataset + "/" + usernameOfProvider + "/" + organizationConsumer
	KeyLoggerRevokedAccessOrgJSON := LoggerRevokedAccessCounterOrg{}
	count := -1
	loggerAccesOrgsAsbytes, err := stub.GetState(monitorRevokedAccessKeyCounterOrg) //get the revoked access from chaincode state
	if err != nil {
		jsonResp = "{\"Error\":\"Failed to get revoked access history for " + monitorRevokedAccessKeyCounterOrg + "\"}"
		return shim.Error(jsonResp)
	} else if loggerAccesOrgsAsbytes == nil {
		count = 0
		loggerAccessOrg := &LoggerRevokedAccessCounterOrg{objectTypeKey, dataset, count}
		loggerAccesOrgsAsbytes, err = json.Marshal(loggerAccessOrg)
		if err != nil {
			return shim.Error(err.Error())
		}

	} else {
		err = json.Unmarshal(loggerAccesOrgsAsbytes, &KeyLoggerRevokedAccessOrgJSON) //unmarshal it aka JSON.parse()
		if err != nil {
			return shim.Error(err.Error())
		}
		count = KeyLoggerRevokedAccessOrgJSON.Count + 1
		KeyLoggerRevokedAccessOrgJSON.Count = count
		loggerAccesOrgsAsbytes, _ = json.Marshal(KeyLoggerRevokedAccessOrgJSON)
	}

	// === Save key for revoked access to state ===
	err = stub.PutState(monitorRevokedAccessKeyCounterOrg, loggerAccesOrgsAsbytes)
	if err != nil {
		return shim.Error(err.Error())
	}

	revokedAccess := &LoggerRevokedAccessOrgs{objectType, dataset, usernameOfProvider, orgOfProvider, description, organizationConsumer, timestamp, revokedPermission, customAccessRights, count,
		usernameConsumer, nameConsumer, surnameConsumer, roleConsumer}
	revokedAccessasBytes, err := json.Marshal(revokedAccess)
	if err != nil {
		return shim.Error(err.Error())
	}
	monitorAccessKey := "revokedOrg" + dataset + "/" + usernameOfProvider + "/" + organizationConsumer + "/" + strconv.Itoa(count)
	// === Save revoked access to state ===
	err = stub.PutState(monitorAccessKey, revokedAccessasBytes)
	if err != nil {
		return shim.Error(err.Error())
	}

	messageofRevokedAccess := "History of revoked access for dataset: " + dataset + "provider: " + usernameOfProvider + "consumer: " + usernameOfProvider + "count: " + strconv.Itoa(count)
	bytemessage := []byte(messageofRevokedAccess)
	return shim.Success(bytemessage)

}

func (t *Logger) queryLoggerRevokedAccessOrgProvider(stub shim.ChaincodeStubInterface) pb.Response {
	usernameOfProvider, err := t.getUsername(stub)
	if err != nil {
		return shim.Error(err.Error())
	}
	//create index key
	queryString := fmt.Sprintf("{\"selector\":{\"docType\":\"LoggerRevokedAccessOrgs\",\"usernameOfProvider\":\"%s\"}}", usernameOfProvider)
	queryResults, err := getQueryResultForQueryString(stub, queryString)
	if err != nil {
		return shim.Error(err.Error())
	}
	return shim.Success(queryResults)
}

func (t *Logger) queryLoggerRevokedAccessOrgConsumer(stub shim.ChaincodeStubInterface) pb.Response {
	organizationOfConsumer, err := t.getOrganization(stub)
	if err != nil {
		return shim.Error(err.Error())
	}

	//create index key
	queryString := fmt.Sprintf("{\"selector\":{\"docType\":\"LoggerRevokedAccessOrgs\",\"organizationConsumer\":\"%s\"}}", organizationOfConsumer)
	queryResults, err := getQueryResultForQueryString(stub, queryString)
	if err != nil {
		return shim.Error(err.Error())
	}
	return shim.Success(queryResults)
}

func (t *Logger) queryLoggerRevokedAccessOrgProviderByDatasetName(stub shim.ChaincodeStubInterface, args []string) pb.Response {
	if len(args) != 1 {
		return shim.Error("Invalid Argument Number. Expecting dataset name")
	}
	datasetname := args[0]
	usernameOfProvider, err := t.getUsername(stub)
	if err != nil {
		return shim.Error(err.Error())
	}
	//create index key
	queryString := fmt.Sprintf("{\"selector\":{\"docType\":\"LoggerRevokedAccessOrgs\",\"usernameOfProvider\":\"%s\",\"dataset\":\"%s\"}}", usernameOfProvider, datasetname)
	queryResults, err := getQueryResultForQueryString(stub, queryString)
	if err != nil {
		return shim.Error(err.Error())
	}
	return shim.Success(queryResults)
}

func (t *Logger) queryLoggerRevokedAccessOrgConsumerByDatasetName(stub shim.ChaincodeStubInterface, args []string) pb.Response {
	if len(args) != 1 {
		return shim.Error("Invalid Argument Number. Expecting dataset name")
	}
	datasetname := args[0]
	organizationOfConsumer, err := t.getOrganization(stub)
	if err != nil {
		return shim.Error(err.Error())
	}
	//create index key
	queryString := fmt.Sprintf("{\"selector\":{\"docType\":\"LoggerRevokedAccessOrgs\",\"organizationConsumer\":\"%s\",\"dataset\":\"%s\"}}", organizationOfConsumer, datasetname)
	queryResults, err := getQueryResultForQueryString(stub, queryString)
	if err != nil {
		return shim.Error(err.Error())
	}
	return shim.Success(queryResults)
}

func (t *Logger) queryLoggerAccessOrgProvider(stub shim.ChaincodeStubInterface) pb.Response {
	usernameOfProvider, err := t.getUsername(stub)
	if err != nil {
		return shim.Error(err.Error())
	}
	//create index key
	queryString := fmt.Sprintf("{\"selector\":{\"docType\":\"LoggerAccessOrgs\",\"usernameOfProvider\":\"%s\"}}", usernameOfProvider)
	queryResults, err := getQueryResultForQueryString(stub, queryString)
	if err != nil {
		return shim.Error(err.Error())
	}
	return shim.Success(queryResults)
}

func (t *Logger) queryLoggerAccessOrgConsumer(stub shim.ChaincodeStubInterface) pb.Response {
	organizationOfConsumer, err := t.getOrganization(stub)
	if err != nil {
		return shim.Error(err.Error())
	}
	//create index key
	queryString := fmt.Sprintf("{\"selector\":{\"docType\":\"LoggerAccessOrgs\",\"organizationConsumer\":\"%s\"}}", organizationOfConsumer)
	queryResults, err := getQueryResultForQueryString(stub, queryString)
	if err != nil {
		return shim.Error(err.Error())
	}
	return shim.Success(queryResults)
}

func (t *Logger) queryLoggerAccessOrgProviderByDatasetName(stub shim.ChaincodeStubInterface, args []string) pb.Response {
	if len(args) != 1 {
		return shim.Error("Invalid Argument Number. Expecting dataset name")
	}
	datasetname := args[0]
	usernameOfProvider, err := t.getUsername(stub)
	if err != nil {
		return shim.Error(err.Error())
	}
	//create index key
	queryString := fmt.Sprintf("{\"selector\":{\"docType\":\"LoggerAccessOrgs\",\"usernameOfProvider\":\"%s\",\"dataset\":\"%s\"}}", usernameOfProvider, datasetname)
	queryResults, err := getQueryResultForQueryString(stub, queryString)
	if err != nil {
		return shim.Error(err.Error())
	}
	return shim.Success(queryResults)
}

func (t *Logger) queryLoggerAccessOrgConsumerByDatasetName(stub shim.ChaincodeStubInterface, args []string) pb.Response {
	if len(args) != 1 {
		return shim.Error("Invalid Argument Number. Expecting dataset name")
	}
	datasetname := args[0]
	organizationOfConsumer, err := t.getOrganization(stub)
	if err != nil {
		return shim.Error(err.Error())
	}
	//create index key
	queryString := fmt.Sprintf("{\"selector\":{\"docType\":\"LoggerAccessOrgs\",\"organizationConsumer\":\"%s\",\"dataset\":\"%s\"}}", organizationOfConsumer, datasetname)
	queryResults, err := getQueryResultForQueryString(stub, queryString)
	if err != nil {
		return shim.Error(err.Error())
	}
	return shim.Success(queryResults)
}
