package main

import (
	"fmt"

	"github.com/hyperledger/fabric/core/chaincode/shim"
)

//QueryMetadata
func getMetadataByDatasetName(stub shim.ChaincodeStubInterface, args []string) (string, error) {
	if len(args) != 1 {
		return "", fmt.Errorf("Incorrect arguments. Expecting dataset name")
	}
	dataset := args[0]

	params := []string{"queryMetadataByName", dataset}
	queryArgs := make([][]byte, len(params))
	for i, arg := range params {
		queryArgs[i] = []byte(arg)
	}

	response := stub.InvokeChaincode("metadatasc", queryArgs, "global")
	if response.Status != shim.OK {
		return "", fmt.Errorf("Failed to query chaincode Metadata for Metadata info. Got error: %s", response.Message)
	}

	return string(response.Payload), nil
}

//QueryMetadataForDataProvider
func getMetadataProvider(stub shim.ChaincodeStubInterface, args []string) (string, error) {
	if len(args) != 1 {
		return "", fmt.Errorf("Incorrect arguments. In order to invoke Metadata expecting dataset")
	}
	dataset := args[0]

	params := []string{"getMetadataProvider", dataset}
	queryArgs := make([][]byte, len(params))
	for i, arg := range params {
		queryArgs[i] = []byte(arg)
	}

	response := stub.InvokeChaincode("metadatasc", queryArgs, "global")
	if response.Status != shim.OK {
		return "", fmt.Errorf("Failed to query chaincode Metadata for Metadata provider. Got error: %s", response.Message)
	}

	return string(response.Payload), nil
}

//QueryMetadataForKeywordTag
func getMetadataKeywordTag(stub shim.ChaincodeStubInterface, args []string) (string, error) {
	if len(args) != 1 {
		return "", fmt.Errorf("Incorrect arguments. In order to invoke Metadata expecting dataset")
	}
	dataset := args[0]

	params := []string{"getMetadataKeywordTag", dataset}
	queryArgs := make([][]byte, len(params))
	for i, arg := range params {
		queryArgs[i] = []byte(arg)
	}

	response := stub.InvokeChaincode("metadatasc", queryArgs, "global")
	if response.Status != shim.OK {
		return "", fmt.Errorf("Failed to query chaincode Metadata for KeywordTag. Got error: %s", response.Message)
	}

	return string(response.Payload), nil
}

//QueryMetadataForEndpoint
func getMetadataEndpoint(stub shim.ChaincodeStubInterface, args []string) (string, error) {
	if len(args) != 1 {
		return "", fmt.Errorf("Incorrect arguments. In order to invoke Metadata expecting dataset")
	}
	dataset := args[0]

	params := []string{"getMetadataEndpoint", dataset}
	queryArgs := make([][]byte, len(params))
	for i, arg := range params {
		queryArgs[i] = []byte(arg)
	}

	response := stub.InvokeChaincode("metadatasc", queryArgs, "global")
	if response.Status != shim.OK {
		return "", fmt.Errorf("Failed to query chaincode Metadata for endpoint. Got error: %s", response.Message)
	}

	return string(response.Payload), nil
}

// Invoke logger about user requests
func setAccessInfotologger(stub shim.ChaincodeStubInterface, args []string) (string, error) {
	if len(args) != 11 {
		return "", fmt.Errorf("Incorrect arguments. In order to invoke logger expecting 11 values")
	}
	dataset := args[0]
	usernameOfProvider := args[1]
	usernameOfConsumer := args[2]
	organization := args[3]
	requestID := args[4]
	status := args[5]
	timestamp := args[6]
	permission := args[7]
	customAccessRights := args[8]
	orgOfProvider := args[9]
	description := args[10]

	params := []string{"monitorAccessRequest", dataset, usernameOfProvider, usernameOfConsumer, organization, requestID, status, timestamp, permission, customAccessRights, orgOfProvider, description}
	invokeArgs := make([][]byte, len(params))
	for i, arg := range params {
		invokeArgs[i] = []byte(arg)
	}

	response := stub.InvokeChaincode("logger", invokeArgs, "global")
	if response.Status != shim.OK {
		return "", fmt.Errorf("Failed to invoke chaincode logger. Got error: %s", response.Payload)
	}
	return "invoke starts", nil
}

// Invoke logger about user revoked access
func setRevokedAccessInfotologger(stub shim.ChaincodeStubInterface, args []string) (string, error) {
	if len(args) != 9 {
		return "", fmt.Errorf("Incorrect arguments. In order to invoke logger for revoked access expecting 9 values")
	}

	dataset := args[0]
	usernameOfProvider := args[1]
	usernameOfConsumer := args[2]
	organization := args[3]
	revokedPermission := args[4]
	timestamp := args[5]
	customAccessRights := args[6]
	orgOfProvider := args[7]
	description := args[8]

	params := []string{"monitorRevokedAccessUser", dataset, usernameOfProvider, usernameOfConsumer, organization, revokedPermission, timestamp, customAccessRights, orgOfProvider, description}
	invokeArgs := make([][]byte, len(params))
	for i, arg := range params {
		invokeArgs[i] = []byte(arg)
	}

	response := stub.InvokeChaincode("logger", invokeArgs, "global")
	if response.Status != shim.OK {
		return "", fmt.Errorf("Failed to invoke chaincode logger. Got error: %s", response.Payload)
	}
	return "invoke starts", nil
}

// Invoke logger about orgs revoked access
func setRevokedAccessOrgsInfotologger(stub shim.ChaincodeStubInterface, args []string) (string, error) {
	if len(args) != 12 {
		return "", fmt.Errorf("Incorrect arguments for function setRevokedAccessOrgsInfotologger. Expecting 12")
	}

	dataset := args[0]
	usernameOfProvider := args[1]
	organizationConsumer := args[2]
	revokedPermission := args[3]
	datetime := args[4]
	consumerUsername := args[5]
	consumerName := args[6]
	consumerSurname := args[7]
	consumerRole := args[8]
	customAccessRights := args[9]
	orgOfProvider := args[10]
	description := args[11]

	params := []string{"monitorRevokedAccessOrg", dataset, usernameOfProvider, organizationConsumer, revokedPermission, datetime, consumerUsername, consumerName, consumerSurname, consumerRole, customAccessRights, orgOfProvider, description}
	invokeArgs := make([][]byte, len(params))
	for i, arg := range params {
		invokeArgs[i] = []byte(arg)
	}

	response := stub.InvokeChaincode("logger", invokeArgs, "global")
	if response.Status != shim.OK {
		return "", fmt.Errorf("Failed to invoke chaincode logger. Got error: %s", response.Payload)
	}
	return "invoke starts", nil
}

// Invoke logger about orgs requests
func setAccessInfotologgerOrgs(stub shim.ChaincodeStubInterface, args []string) (string, error) {
	if len(args) != 14 {
		return "", fmt.Errorf("Incorrect arguments. In order to invoke logger expecting 14 values")
	}
	dataset := args[0]
	usernameOfProvider := args[1]
	organizationOfConsumer := args[2]
	requestID := args[3]
	status := args[4]
	datetime := args[5]
	consumerUsername := args[6]
	consumerName := args[7]
	consumerSurname := args[8]
	consumerRole := args[9]
	permission := args[10]
	customAccessRights := args[11]
	orgOfProvider := args[12]
	description := args[13]

	params := []string{"monitorAccessRequestOrgs", dataset, usernameOfProvider, organizationOfConsumer, requestID, status, datetime, consumerUsername, consumerName, consumerSurname, consumerRole, permission, customAccessRights, orgOfProvider, description}
	invokeArgs := make([][]byte, len(params))
	for i, arg := range params {
		invokeArgs[i] = []byte(arg)
	}

	response := stub.InvokeChaincode("logger", invokeArgs, "global")
	if response.Status != shim.OK {
		return "", fmt.Errorf("Failed to invoke chaincode logger for org access. Got error: %s", response.Message)
	}
	return "invoke starts", nil
}

// Invoke logger about public access
func setAccessInfotologgerPublic(stub shim.ChaincodeStubInterface, args []string) (string, error) {
	if len(args) != 5 {
		return "", fmt.Errorf("Incorrect arguments. In order to invoke logger for public access expecting dataset name, username of provider, permission, custom access rights & timestamp")
	}
	dataset := args[0]
	usernameOfProvider := args[1]
	permission := args[2]
	timestamp := args[3]
	customAccessRights := args[4]

	params := []string{"monitorAccessPublic", dataset, usernameOfProvider, permission, timestamp, customAccessRights}
	invokeArgs := make([][]byte, len(params))
	for i, arg := range params {
		invokeArgs[i] = []byte(arg)
	}

	response := stub.InvokeChaincode("logger", invokeArgs, "global")
	if response.Status != shim.OK {
		return "", fmt.Errorf("Failed to invoke chaincode logger for public access. Got error: %s", response.Message)
	}
	return "invoke starts", nil
}

// Invoke logger about public revoked access
func setRevokedAccessPublicInfotologger(stub shim.ChaincodeStubInterface, args []string) (string, error) {
	if len(args) != 5 {
		return "", fmt.Errorf("Incorrect arguments. In order to invoke logger for revoked public access expecting dataset name, username of provider, revoked permission, monitorRevokedAccessPublic & timestamp")
	}

	dataset := args[0]
	usernameOfProvider := args[1]
	revokedPermission := args[2]
	timestamp := args[3]
	customAccessRights := args[4]

	params := []string{"monitorRevokedAccessPublic", dataset, usernameOfProvider, revokedPermission, timestamp, customAccessRights}
	invokeArgs := make([][]byte, len(params))
	for i, arg := range params {
		invokeArgs[i] = []byte(arg)
	}

	response := stub.InvokeChaincode("logger", invokeArgs, "global")
	if response.Status != shim.OK {
		return "", fmt.Errorf("Failed to invoke chaincode logger. Got error: %s", response.Payload)
	}
	return "invoke starts", nil
}
