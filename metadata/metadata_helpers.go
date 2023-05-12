package main

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"

	"github.com/hyperledger/fabric/core/chaincode/shim"
	"github.com/hyperledger/fabric/core/chaincode/shim/ext/cid"
	pb "github.com/hyperledger/fabric/protos/peer"
)

//get the external user flag of the attribute values
func (c *Metadata) getIsExternalUser(stub shim.ChaincodeStubInterface) (string, error) {
	isExternalUser, ok, err := cid.GetAttributeValue(stub, "isExternalUser")

	if err != nil {
		return "", err
	}

	if !ok {
		return "", errors.New("isExternalUser attribute is missing")
	}

	return isExternalUser, nil
}

//get the external user org of the attribute values
func (c *Metadata) getExternalOrg(stub shim.ChaincodeStubInterface) (string, error) {
	externalOrg, ok, err := cid.GetAttributeValue(stub, "externalOrg")

	if err != nil {
		return "", err
	}

	if !ok {
		return "", errors.New("externalOrg attribute is missing")
	}

	return externalOrg, nil
}

func (c *Metadata) getRole(stub shim.ChaincodeStubInterface) (string, error) {
	organization, ok, err := cid.GetAttributeValue(stub, "dataportsRole")

	if err != nil {
		return "", err
	}

	if !ok {
		return "", errors.New("role attribute is missing")
	}

	return organization, nil
}

func (c *Metadata) getUsername(stub shim.ChaincodeStubInterface) (string, error) {
	username, ok, err := cid.GetAttributeValue(stub, "username")

	if err != nil {
		return "", err
	}

	if !ok {
		return "", errors.New("username attribute is missing")
	}

	return username, nil
}

func (c *Metadata) changeProviderEmail(stub shim.ChaincodeStubInterface, args []string) pb.Response {
	var err error
	emailOfProvider := args[0]
	datasetName := args[1]

	// ==== Check if name already exists ====
	datasetMetadataAsBytes, err := stub.GetState(datasetName)
	if err != nil {
		return shim.Error("Failed to check if the name of metadata exists: " + err.Error())
	} else if datasetMetadataAsBytes == nil {
		return shim.Error("Does not exists metadata with name: " + datasetName)
	}
	datasetMetadata := DatasetMetadata{}
	err = json.Unmarshal(datasetMetadataAsBytes, &datasetMetadata) //unmarshal it aka JSON.parse()
	if err != nil {
		return shim.Error(err.Error())
	}

	usernameOfProviderCID, err := c.getUsername(stub)
	if err != nil {
		return shim.Error(err.Error())
	}

	if datasetMetadata.UsernameOfProvider != usernameOfProviderCID {
		return shim.Error("Only the data provider can update the dataset")
	}

	datasetMetadata.EmailOfProvider = emailOfProvider

	MetadataAsBytes, _ := json.Marshal(datasetMetadata)
	err = stub.PutState(datasetName, MetadataAsBytes)
	if err != nil {
		return shim.Error(err.Error())
	}

	return shim.Success(nil)

}

func getQueryResultForQueryString(stub shim.ChaincodeStubInterface, queryString string) ([]byte, error) {

	fmt.Printf("- getQueryResultForQueryString queryString:\n%s\n", queryString)

	resultsIterator, err := stub.GetQueryResult(queryString)
	if err != nil {
		return nil, err
	}
	defer resultsIterator.Close()

	buffer, err := constructQueryResponseFromIterator(resultsIterator)
	if err != nil {
		return nil, err
	}

	fmt.Printf("- getQueryResultForQueryString queryResult:\n%s\n", buffer.String())

	return buffer.Bytes(), nil
}

func constructQueryResponseFromIterator(resultsIterator shim.StateQueryIteratorInterface) (*bytes.Buffer, error) {
	// buffer is a JSON array containing QueryResults
	var buffer bytes.Buffer
	buffer.WriteString("[")

	bArrayMemberAlreadyWritten := false
	for resultsIterator.HasNext() {
		queryResponse, err := resultsIterator.Next()
		if err != nil {
			return nil, err
		}
		// Add a comma before array members, suppress it for the first array member
		if bArrayMemberAlreadyWritten {
			buffer.WriteString(",")
		}
		buffer.WriteString("{\"Key\":")
		buffer.WriteString("\"")
		buffer.WriteString(queryResponse.Key)
		buffer.WriteString("\"")

		buffer.WriteString(", \"Record\":")
		// Record is a JSON object, so we write as-is
		buffer.WriteString(string(queryResponse.Value))
		buffer.WriteString("}")
		bArrayMemberAlreadyWritten = true
	}
	buffer.WriteString("]")

	return &buffer, nil
}
