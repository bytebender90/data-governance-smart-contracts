package main

import (
	"encoding/json"
	"errors"
	"fmt"

	"github.com/hyperledger/fabric/core/chaincode/shim"
	"github.com/hyperledger/fabric/core/chaincode/shim/ext/cid"
	pb "github.com/hyperledger/fabric/protos/peer"
)

// InvokeLogger
func setMetadataInfotoLogger(stub shim.ChaincodeStubInterface, args []string) (string, error) {
	if len(args) != 7 {
		return "", fmt.Errorf("Incorrect arguments. In order to invoke Logger expecting 7 values")
	}
	dataset := args[0]
	username := args[1]
	organization := args[2]
	changes := args[3]
	version := args[4]
	comments := args[5]
	dateTime := args[6]

	params := []string{"monitorUploadMetadata", dataset, username, organization, changes, version, comments, dateTime}
	invokeArgs := make([][]byte, len(params))
	for i, arg := range params {
		invokeArgs[i] = []byte(arg)
	}

	response := stub.InvokeChaincode("logger", invokeArgs, "global")
	if response.Status != shim.OK {
		return "", fmt.Errorf("Failed to invoke chaincode Logger. Got error: %s", response.Payload)
	}
	return "invoke starts", nil
}

func setDataSourceMetadataInfotoLogger(stub shim.ChaincodeStubInterface, args []string) (string, error) {
	if len(args) != 4 {
		return "", fmt.Errorf("Incorrect arguments. In order to invoke Logger expecting 4 values")
	}
	dataSource := args[0]
	status := args[1]
	changes := args[2]
	dateTime := args[3]

	params := []string{"monitorDatasourceMetadata", dataSource, status, changes, dateTime}
	invokeArgs := make([][]byte, len(params))
	for i, arg := range params {
		invokeArgs[i] = []byte(arg)
	}

	response := stub.InvokeChaincode("logger", invokeArgs, "global")
	if response.Status != shim.OK {
		return "", fmt.Errorf("Failed to invoke chaincode Logger. Got error: %s", response.Payload)
	}
	return "invoke starts", nil
}

func (c *Metadata) getMetadataProvider(stub shim.ChaincodeStubInterface, args []string) pb.Response {
	var err error
	if len(args) != 1 {
		return shim.Error("Incorrect number of arguments. Expecting name of metadata")
	}

	title := args[0]

	value, err := stub.GetState(title)
	if err != nil {
		return shim.Error("Failed to get dataset on IDS broker")
	}

	datasetMetadata := DatasetMetadata{}
	err = json.Unmarshal(value, &datasetMetadata) //unmarshal it aka JSON.parse()
	if err != nil {
		return shim.Error(err.Error())
	}

	return shim.Success([]byte(datasetMetadata.UsernameOfProvider))
}

func (c *Metadata) getMetadataEndpoint(stub shim.ChaincodeStubInterface, args []string) pb.Response {
	var err error
	if len(args) != 1 {
		return shim.Error("Incorrect number of arguments. Expecting name of metadata")
	}

	title := args[0]

	value, err := stub.GetState(title)
	if err != nil {
		return shim.Error("Failed to get dataset on IDS broker")
	}

	datasetMetadata := DatasetMetadata{}
	err = json.Unmarshal(value, &datasetMetadata) //unmarshal it aka JSON.parse()
	if err != nil {
		return shim.Error(err.Error())
	}
	return shim.Success([]byte(datasetMetadata.Endpoint))
}

func (c *Metadata) getMetadataKeywordTag(stub shim.ChaincodeStubInterface, args []string) pb.Response {
	var err error
	if len(args) != 1 {
		return shim.Error("Incorrect number of arguments. Expecting name of metadata")
	}

	title := args[0]

	value, err := stub.GetState(title)
	if err != nil {
		return shim.Error("Failed to get dataset on IDS broker")
	}
	
	datasetMetadata := DatasetMetadata{}
	err = json.Unmarshal(value, &datasetMetadata) //unmarshal it aka JSON.parse()
	if err != nil {
		return shim.Error(err.Error())
	}
	return shim.Success([]byte(datasetMetadata.KeywordTag))
}

//get the organization of the attribute values
func (c *Metadata) getOrganization(stub shim.ChaincodeStubInterface) (string, error) {
	organization, ok, err := cid.GetAttributeValue(stub, "organization")

	if err != nil {
		return "", err
	}

	if !ok {
		return "", errors.New("organization attribute is missing")
	}

	return organization, nil
}
