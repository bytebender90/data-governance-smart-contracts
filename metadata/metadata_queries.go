package main

import (
	"fmt"

	"github.com/hyperledger/fabric/core/chaincode/shim"
	pb "github.com/hyperledger/fabric/protos/peer"
)

func (c *Metadata) queryMetadataStruct(stub shim.ChaincodeStubInterface) pb.Response {
	queryString := ("{\"selector\":{\"docType\":\"Metadata\"}}")

	queryResults, err := getQueryResultForQueryString(stub, queryString)
	if err != nil {
		return shim.Error(err.Error())
	}
	return shim.Success(queryResults)

}

func (c *Metadata) queryMetadataByName(stub shim.ChaincodeStubInterface, args []string) pb.Response {
	var jsonResp string
	var err error

	if len(args) != 1 {
		return shim.Error("Incorrect number of arguments. Expecting name of metadata to query")
	}

	datasetName := args[0]
	valAsbytes, err := stub.GetState(datasetName)
	if err != nil {
		jsonResp = "{\"Error\":\"Failed to get state for " + datasetName + "\"}"
		return shim.Error(jsonResp)
	} else if valAsbytes == nil {
		jsonResp = "{\"Error\":\"Metadata does not exist: " + datasetName + "\"}"
		return shim.Error(jsonResp)
	}

	return shim.Success(valAsbytes)

}

func (c *Metadata) queryMetadataByDataSourceID(stub shim.ChaincodeStubInterface, args []string) pb.Response {
	if len(args) != 1 {
		return shim.Error("Incorrect arguments. Expecting dataSourceID")
	}
	dataSourceID := args[0]

	queryString := fmt.Sprintf("{\"selector\":{\"docType\":\"Metadata\", \"dataSourceMetadata.dataSourceID\":\"%s\"}}", dataSourceID)

	queryResults, err := getQueryResultForQueryString(stub, queryString)
	if err != nil {
		return shim.Error(err.Error())
	}
	return shim.Success(queryResults)

}

func (c *Metadata) queryMetadataByProvider(stub shim.ChaincodeStubInterface) pb.Response {
	usernameOfProvider, err := c.getUsername(stub)
	if err != nil {
		return shim.Error(err.Error())
	}

	queryString := fmt.Sprintf("{\"selector\":{\"docType\":\"Metadata\", \"usernameOfProvider\":\"%s\"}}", usernameOfProvider)

	queryResults, err := getQueryResultForQueryString(stub, queryString)
	if err != nil {
		return shim.Error(err.Error())
	}
	return shim.Success(queryResults)

}

func (c *Metadata) queryMetadataByProviderAndDatasetName(stub shim.ChaincodeStubInterface, args []string) pb.Response {
	if len(args) != 1 {
		return shim.Error("Incorrect arguments. Expecting dataset name")
	}
	dataset := args[0]

	usernameOfProvider, err := c.getUsername(stub)
	if err != nil {
		return shim.Error(err.Error())
	}

	queryString := fmt.Sprintf("{\"selector\":{\"docType\":\"Metadata\", \"usernameOfProvider\":\"%s\", \"datasetName\":\"%s\"}}", usernameOfProvider, dataset)

	queryResults, err := getQueryResultForQueryString(stub, queryString)
	if err != nil {
		return shim.Error(err.Error())
	}
	return shim.Success(queryResults)

}

func (c *Metadata) queryMetadataByOwner(stub shim.ChaincodeStubInterface) pb.Response {

	usernameOfOwner, err := c.getUsername(stub)
	if err != nil {
		return shim.Error(err.Error())
	}

	queryString := fmt.Sprintf("{\"selector\":{\"docType\":\"Metadata\", \"usernameOfOwner\":\"%s\"}}", usernameOfOwner)

	queryResults, err := getQueryResultForQueryString(stub, queryString)
	if err != nil {
		return shim.Error(err.Error())
	}
	return shim.Success(queryResults)

}

func (c *Metadata) queryMetadataByOrganizationOwner(stub shim.ChaincodeStubInterface) pb.Response {
	var organizationOfOwner string

	dataportsRole, err := c.getRole(stub)
	if err != nil {
		return shim.Error(err.Error())
	}
	if dataportsRole != "admin" {
		return shim.Error("Only organization admin can query organizational owner metadata")
	}
	//check for external user
	isExternalUser, err := c.getIsExternalUser(stub)
	if err != nil {
		return shim.Error(err.Error())
	}
	if isExternalUser == "false" {
		organizationOfOwner, err = c.getOrganization(stub)
		if err != nil {
			return shim.Error(err.Error())
		}

	} else if isExternalUser == "true" {
		organizationOfOwner, err = c.getExternalOrg(stub)
		if err != nil {
			return shim.Error(err.Error())
		}
	} else {
		return shim.Error("Unknown int/ext user")
	}

	queryString := fmt.Sprintf("{\"selector\":{\"docType\":\"Metadata\", \"orgOfOwner\":\"%s\"}}", organizationOfOwner)

	queryResults, err := getQueryResultForQueryString(stub, queryString)
	if err != nil {
		return shim.Error(err.Error())
	}
	return shim.Success(queryResults)

}

func (c *Metadata) queryMetadataByOrganizationOwnerAndDatasetName(stub shim.ChaincodeStubInterface, args []string) pb.Response {

	dataset := args[0]
	organizationOfOwner, err := c.getOrganization(stub)
	if err != nil {
		return shim.Error(err.Error())
	}

	queryString := fmt.Sprintf("{\"selector\":{\"docType\":\"Metadata\", \"orgOfOwner\":\"%s\", \"datasetName\":\"%s\"}}", organizationOfOwner, dataset)

	queryResults, err := getQueryResultForQueryString(stub, queryString)
	if err != nil {
		return shim.Error(err.Error())
	}
	return shim.Success(queryResults)

}

func (c *Metadata) queryMetadataByOwnerAndDatasetName(stub shim.ChaincodeStubInterface, args []string) pb.Response {
	if len(args) != 1 {
		return shim.Error("Incorrect arguments. Expecting dataset name")
	}
	dataset := args[0]
	organizationOfOwner, err := c.getOrganization(stub)
	if err != nil {
		return shim.Error(err.Error())
	}

	usernameOfOwner, err := c.getUsername(stub)
	if err != nil {
		return shim.Error(err.Error())
	}

	queryString := fmt.Sprintf("{\"selector\":{\"docType\":\"Metadata\", \"usernameOfOwner\":\"%s\", \"orgOfOwner\":\"%s\", \"datasetName\":\"%s\"}}", usernameOfOwner, organizationOfOwner, dataset)

	queryResults, err := getQueryResultForQueryString(stub, queryString)
	if err != nil {
		return shim.Error(err.Error())
	}
	return shim.Success(queryResults)

}

func (c *Metadata) queryMetadataByIdentifier(stub shim.ChaincodeStubInterface, args []string) pb.Response {
	if len(args) != 1 {
		return shim.Error("Incorrect arguments. Expecting identifier")
	}
	identifier := args[0]

	queryString := fmt.Sprintf("{\"selector\":{\"docType\":\"Metadata\", \"identifier\":\"%s\"}}", identifier)

	queryResults, err := getQueryResultForQueryString(stub, queryString)
	if err != nil {
		return shim.Error(err.Error())
	}
	return shim.Success(queryResults)

}

func (c *Metadata) queryMetadataByQueryString(stub shim.ChaincodeStubInterface, args []string) pb.Response {
	if len(args) != 1 {
		return shim.Error("Incorrect arguments. Expecting query string")
	}
	queryString := args[0]
	queryResults, err := getQueryResultForQueryString(stub, queryString)
	if err != nil {
		return shim.Error(err.Error())
	}
	return shim.Success(queryResults)

}
