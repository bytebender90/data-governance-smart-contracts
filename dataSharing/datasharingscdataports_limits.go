package main

import (
	"encoding/json"
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/hyperledger/fabric/core/chaincode/shim"
	pb "github.com/hyperledger/fabric/protos/peer"
)

func (c *ShareData) setDatasetLimits(stub shim.ChaincodeStubInterface, args []string) pb.Response {
	var jsonResp string
	if len(args) != 3 {
		return shim.Error("Invalid Argument Number. Expected 1) dataset name 2) Date 3) Fee")
	}
	objectType := "DatasetLimits"
	datasetname := args[0]

	if len(datasetname) <= 0 {
		return shim.Error("Dataset name must be a non empty string")
	}

	date := args[1]

	if len(date) <= 0 {
		return shim.Error("Date name must on the format 2006/01/02")
	}
	fee := args[2]
	if len(fee) <= 0 {
		fee = "0"
	}

	feeInt, err := strconv.Atoi(fee)
	if err != nil {
		fmt.Println(err)

	}

	usernameofMetadataProvider, err := c.getUsername(stub)
	if err != nil {
		return shim.Error(err.Error())
	}

	//Invoke Metadata in order to check Provider
	sendInfoToMetadata, err := getMetadataProvider(stub, []string{datasetname})
	if err != nil {
		return shim.Error(err.Error())
	}
	if usernameofMetadataProvider != sendInfoToMetadata {
		jsonResp = "{\"Error\":\"Not authorized user:  " + usernameofMetadataProvider + "for dataset: " + datasetname + "\"}"
		return shim.Error(jsonResp)
	}

	key := datasetname + "limits"
	datasetLimits := DatasetLimits{objectType, datasetname, usernameofMetadataProvider, date, feeInt}
	datasetLimitsJSONasBytes, err := json.Marshal(datasetLimits)
	if err != nil {
		return shim.Error(err.Error())
	}

	// === Save to state ===
	err = stub.PutState(key, datasetLimitsJSONasBytes)
	if err != nil {
		return shim.Error(err.Error())
	}

	messageofdata := "Have been set limits for dataset: " + datasetname
	bytemessage := []byte(messageofdata)
	return shim.Success(bytemessage)

}

func (c *ShareData) queryAllDatasetLimits(stub shim.ChaincodeStubInterface) pb.Response {
	queryString := ("{\"selector\":{\"docType\":\"DatasetLimits\"}}")

	queryResults, err := getQueryResultForQueryString(stub, queryString)
	if err != nil {
		return shim.Error(err.Error())
	}
	return shim.Success(queryResults)
}

func (c *ShareData) queryDatasetLimitsByProvider(stub shim.ChaincodeStubInterface) pb.Response {
	usernameofProvider, err := c.getUsername(stub)
	if err != nil {
		return shim.Error(err.Error())
	}

	queryString := fmt.Sprintf("{\"selector\":{\"docType\":\"DatasetLimits\",\"dataset_provider\":\"%s\"}}", usernameofProvider)

	queryResults, err := getQueryResultForQueryString(stub, queryString)
	if err != nil {
		return shim.Error(err.Error())
	}
	return shim.Success(queryResults)
}

func (c *ShareData) queryDatasetLimitByName(stub shim.ChaincodeStubInterface, args []string) pb.Response {
	var name, jsonResp string
	var err error

	if len(args) != 1 {
		return shim.Error("Incorrect number of arguments. Expecting dataset name to query")
	}

	name = args[0]
	key := name + "limits"
	valAsbytes, err := stub.GetState(key)
	if err != nil {
		jsonResp = "{\"Error\":\"Failed to get state for " + key + "\"}"
		return shim.Error(jsonResp)
	} else if valAsbytes == nil {
		jsonResp = "{\"Error\":\"Dataset limits does not exist: " + key + "\"}"
		return shim.Error(jsonResp)
	}

	return shim.Success(valAsbytes)
}

func (c *ShareData) checkLimits(stub shim.ChaincodeStubInterface, args []string) pb.Response {
	var jsonResp string
	if len(args) < 1 {
		return shim.Error("Incorrect number of arguments. Expecting 1. Dataset name in order to check for limits.")
	}

	datasetname := args[0]
	keylimits := datasetname + "limits"

	datasetLimitsJson := DatasetLimits{}
	limitsAsbytes, err := stub.GetState(keylimits) //get limits for the dataset
	if err != nil {
		jsonResp = "{\"Error\":\"Failed to get get limits for dataset: " + datasetname + "\"}"
		return shim.Error(jsonResp)
	} else if limitsAsbytes == nil {
		return shim.Success(nil)
	}

	err = json.Unmarshal(limitsAsbytes, &datasetLimitsJson)
	if err != nil {
		jsonResp = "{\"Error\":\"Failed to decode JSON" + keylimits + "\"}"
		return shim.Error(jsonResp)
	}
	date := datasetLimitsJson.DateLimit
	dateString := strings.Split(date, "/")
	year, err := strconv.Atoi(dateString[0])
	if err != nil {
		fmt.Println(err)

	}
	day, err := strconv.Atoi(dateString[1])
	if err != nil {
		fmt.Println(err)

	}
	month, err := strconv.Atoi(dateString[2])
	if err != nil {
		fmt.Println(err)

	}

	//get current date
	current_date := time.Now()
	finalcurrent_date := current_date.Format("2006/01/02")
	dateStringCurrent := strings.Split(finalcurrent_date, "/")
	yearCurrent, err := strconv.Atoi(dateStringCurrent[0])
	if err != nil {
		fmt.Println(err)

	}
	dayCurrent, err := strconv.Atoi(dateStringCurrent[1])
	if err != nil {
		fmt.Println(err)

	}
	monthCurrent, err := strconv.Atoi(dateStringCurrent[2])
	if err != nil {
		fmt.Println(err)

	}
	limitdate := Date(year, day, month)

	finalCurrentDate := Date(yearCurrent, dayCurrent, monthCurrent)
	compareDates := finalCurrentDate.Sub(limitdate).Hours() / 24
	if compareDates <= 0 {
		jsonResp = "{\"Error\":\"The dataset" + datasetname + "is not available yet" + "\"}"
		return shim.Error(jsonResp)
	}
	return shim.Success(nil)
}

func Date(year, month, day int) time.Time {
	return time.Date(year, time.Month(month), day, 0, 0, 0, 0, time.UTC)

}
