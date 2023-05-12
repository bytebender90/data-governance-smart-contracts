package main

import (
	"fmt"

	"github.com/hyperledger/fabric/core/chaincode/shim"
	pb "github.com/hyperledger/fabric/protos/peer"
)

type Metadata struct{}

func (c *Metadata) Init(stub shim.ChaincodeStubInterface) pb.Response {
	return shim.Success(nil)
}

// Implement Invoke
func (c *Metadata) Invoke(stub shim.ChaincodeStubInterface) pb.Response {

	function, args := stub.GetFunctionAndParameters()

	switch function {
	case "uploadMetadata":
		return c.uploadMetadata(stub, args)
	case "updateMetadata":
		return c.updateMetadata(stub, args)
	case "createDataSourceMetadata":
		return c.createDataSourceMetadata(stub, args)
	case "updateDataSourceMetadata":
		return c.updateDataSourceMetadata(stub, args)
	case "queryMetadataStruct":
		return c.queryMetadataStruct(stub)
	case "queryMetadataByName":
		return c.queryMetadataByName(stub, args)
	case "queryMetadataByDataSourceID":
		return c.queryMetadataByDataSourceID(stub, args)
	case "queryMetadataByOwner":
		return c.queryMetadataByOwner(stub)
	case "queryMetadataByOwnerAndDatasetName":
		return c.queryMetadataByOwnerAndDatasetName(stub, args)
	case "queryMetadataByProvider":
		return c.queryMetadataByProvider(stub)
	case "queryMetadataByProviderAndDatasetName":
		return c.queryMetadataByProviderAndDatasetName(stub, args)
	case "queryMetadataByOrganizationOwner":
		return c.queryMetadataByOrganizationOwner(stub)
	case "queryMetadataByOrganizationOwnerAndDatasetName":
		return c.queryMetadataByOrganizationOwnerAndDatasetName(stub, args)
	case "queryMetadataByQueryString":
		return c.queryMetadataByQueryString(stub, args)
	case "getMetadataProvider":
		return c.getMetadataProvider(stub, args)
	case "getMetadataKeywordTag":
		return c.getMetadataKeywordTag(stub, args)
	case "getMetadataEndpoint":
		return c.getMetadataEndpoint(stub, args)
	case "changeProviderEmail":
		return c.changeProviderEmail(stub, args)
	case "queryMetadataByIdentifier":
		return c.queryMetadataByIdentifier(stub, args)
	}

	fmt.Println("invoke did not find func: " + function) //error
	return shim.Error("Received unknown function invocation")
}

func main() {
	err := shim.Start(new(Metadata))
	if err != nil {
		fmt.Printf("Error starting chaincode sample: %s", err)
	}
}

