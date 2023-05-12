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

type DatasetMetadataResponse struct {
	ObjectType string `json:"docType"`
	//Access control properties
	UsernameOfProvider string `json:"usernameOfProvider"`
	OrgOfProvider      string `json:"orgOfProvider"`
	EmailOfProvider    string `json:"emailOfProvider"`
	UsernameOfOwner    string `json:"usernameOfOwner"`
	OrgOfOwner         string `json:"orgOfOwner"`
	//Informative properties
	DatasetName          string `json:"datasetName"`
	Identifier           string `json:"identifier"`
	DatasetDescription   string `json:"datasetDescription"`
	ReleaseDateTime      string `json:"releaseDateTime"`
	ModificationDateTime string `json:"modificationDateTime"`
	//Descriptive properties
	ThemeCategory             string `json:"themeCategory"`
	KeywordTag                string `json:"keywordTag"`
	Language                  string `json:"language"`
	Distribution              string `json:"distribution"`
	DataVelocity              int    `json:"dataVelocity"`
	SpatialGeographicCoverage string `json:"spatialGeographicCoverage"`
	TemporalCoverageStart     string `json:"temporalCoverageStart"`
	TemporalCoverageEnd       string `json:"temporalCoverageEnd"`
	IndustryDomain            string `json:"industryDomain"`
	DataVolume                int    `json:"dataVolume"`
	Comments                  string `json:"comments"`
	//Policy-related properties
	AccessRights          []string `json:"accessRights"`
	ContractAgreementHash string   `json:"contractAgreementHash"`
	ContractAgreementURL  string   `json:"contractAgreementURL"`
	TermsConditionsHash   string   `json:"termsConditionsHash"`
	TermsConditionsURL    string   `json:"termsConditionsURL"`
	CopyrightHash         string   `json:"copyrightHash"`
	CopyrightURL          string   `json:"copyrightsURL"`
	CustomAccessRights    []string `json:"customAccessRights"`
	//Extra Data Governance metadata
	Endpoint   string `json:"endpoint"`
	Blockchain string `json:"blockchain"`
	Version    int    `json:"version"`
	HasAgent   string `json:"hasAgent"`
	//Data source metadata
	DataSourceMetadata *DataSourceMetadata `json:"dataSourceMetadata"`
}

type DataSourceMetadata struct {
	DataSourceID                string              `json:"dataSourceID"`
	DataSourceType              string              `json:"dataSourceType"`
	DataModels                  *DataModelsStruct   `json:"dataModels"`
	DataProvided                *DataProvidedStruct `json:"dataProvided"`
	Attributes                  string              `json:"attributes"`
	Service                     string              `json:"service"`
	ServicePath                 string              `json:"servicePath"`
	Mapping                     string              `json:"mapping"`
	DataportsDataModelandFormat bool                `json:"isDataPortsCompatible"`
}

type DataModelsStruct struct {
	Type     string `json:"type"`
	Value    string `json:"value"`
	Metadata string `json:"metadata"`
}

type DataProvidedStruct struct {
	Type     string `json:"type"`
	Value    string `json:"value"`
	Metadata string `json:"metadata"`
}

/*
func appendCustomAccessRights(requestedRights []string, existedRights []string) []string {
    // Create a map to hold all unique values.
    // The keys will be the values and the values will always be true.
    mergedMap := make(map[string]bool)

    // Add all values from the first input slice to the map.
    for _, value := range requestedRights {
        mergedMap[value] = true
    }

    // Add all values from the second input slice to the map.
    // This will overwrite any existing keys with the same value.
    for _, value := range existedRights {
        mergedMap[value] = true
    }

    // Create a new slice to hold the unique values.
    // The capacity is set to the size of the map to avoid unnecessary allocations.
    mergedSlice := make([]string, 0, len(mergedMap))

    // Add all keys from the map to the slice.
    for key := range mergedMap {
        mergedSlice = append(mergedSlice, key)
    }

    // Return the final slice.
    return mergedSlice
}
*/
func appendCustomAccessRights(requestedRights []string, existedRights []string) []string {
    // Create a map to hold all unique non-empty values.
    // The keys will be the values and the values will always be true.
    mergedMap := make(map[string]bool)

    // Add all non-empty values from the first input slice to the map.
    for _, value := range requestedRights {
        if value != "" && value != " " {
            mergedMap[value] = true
        }
    }

    // Add all non-empty values from the second input slice to the map.
    // This will overwrite any existing keys with the same value.
    for _, value := range existedRights {
        if value != "" && value != " " {
            mergedMap[value] = true
        }
    }

    // Create a new slice to hold the unique non-empty values.
    // The capacity is set to the size of the map to avoid unnecessary allocations.
    mergedSlice := make([]string, 0, len(mergedMap))

    // Add all keys from the map to the slice.
    for key := range mergedMap {
        mergedSlice = append(mergedSlice, key)
    }

    // Return the final slice.
    return mergedSlice
}

func checkCustomAccessRights(requestedRights []string, existedRights []string) bool {
	fmt.Println(len(requestedRights))
	if len(requestedRights) == 0 || requestedRights[0] == "" || requestedRights[0] == " " {
		return true
	}

	for _, val1 := range requestedRights {
		exists := false
		for _, val2 := range existedRights {
			if val1 == val2 {
				exists = true
				break
			}
		}
		if !exists {
			fmt.Println("Value", val1, "in array requestedRights does not exist in existedRights")
			return false
		}
	}

	return true
}

func removeElements(s1 []string, s2 []string) []string {
	// Create a map to store the elements of the second slice
	m := make(map[string]bool)
	for _, elem := range s2 {
		m[elem] = true
	}

	// Create a new slice to store the elements of the first slice that are not present in the second slice
	result := make([]string, 0)
	for _, elem := range s1 {
		if _, exists := m[elem]; !exists {
			result = append(result, elem)
		}
	}

	return result
}

//get the external user flag of the attribute values
func (c *ShareData) getIsExternalUser(stub shim.ChaincodeStubInterface) (string, error) {
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
func (c *ShareData) getExternalOrg(stub shim.ChaincodeStubInterface) (string, error) {
	externalOrg, ok, err := cid.GetAttributeValue(stub, "externalOrg")

	if err != nil {
		return "", err
	}

	if !ok {
		return "", errors.New("externalOrg attribute is missing")
	}

	return externalOrg, nil
}

//get the name of the attribute values
func (c *ShareData) getName(stub shim.ChaincodeStubInterface) (string, error) {
	name, ok, err := cid.GetAttributeValue(stub, "name")

	if err != nil {
		return "", err
	}

	if !ok {
		return "", errors.New("name attribute is missing")
	}

	return name, nil
}

//get the surname of the attribute values
func (c *ShareData) getSurname(stub shim.ChaincodeStubInterface) (string, error) {
	surname, ok, err := cid.GetAttributeValue(stub, "surname")

	if err != nil {
		return "", err
	}

	if !ok {
		return "", errors.New("surname attribute is missing")
	}

	return surname, nil
}

//get the organization of the attribute values
func (c *ShareData) getOrganization(stub shim.ChaincodeStubInterface) (string, error) {
	organization, ok, err := cid.GetAttributeValue(stub, "organization")

	if err != nil {
		return "", err
	}

	if !ok {
		return "", errors.New("organization attribute is missing")
	}

	return organization, nil
}

//get the username of the attribute values
func (c *ShareData) getUsername(stub shim.ChaincodeStubInterface) (string, error) {
	username, ok, err := cid.GetAttributeValue(stub, "username")

	if err != nil {
		return "", err
	}

	if !ok {
		return "", errors.New("username attribute is missing")
	}

	return username, nil
}

func (c *ShareData) getRole(stub shim.ChaincodeStubInterface) (string, error) {
	role, ok, err := cid.GetAttributeValue(stub, "dataportsRole")

	if err != nil {
		return "", err
	}

	if !ok {
		return "", errors.New("role attribute is missing")
	}

	return role, nil
}

func (c *ShareData) getEmail(stub shim.ChaincodeStubInterface) (string, error) {
	role, ok, err := cid.GetAttributeValue(stub, "email")

	if err != nil {
		return "", err
	}

	if !ok {
		return "", errors.New("email attribute is missing")
	}

	return role, nil
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

func (c *ShareData) getRequestByID(stub shim.ChaincodeStubInterface, args []string) pb.Response {
	var requestID, jsonResp string
	var err error

	if len(args) != 1 {
		return shim.Error("Incorrect number of arguments. Expecting requestID to query")
	}

	requestID = args[0]
	valAsbytes, err := stub.GetState(requestID) //get the request from chaincode state
	if err != nil {
		jsonResp = "{\"Error\":\"Failed to get state for " + requestID + "\"}"
		return shim.Error(jsonResp)
	}

	return shim.Success(valAsbytes)
}

func (c *ShareData) updateInfoAuthorizeUsers(stub shim.ChaincodeStubInterface, args []string) pb.Response {
	var jsonResp string
	var err error
	if len(args) != 3 {
		return shim.Error("Incorrect number of arguments. Expecting 3")
	}

	keydataset := args[0]
	name := args[1]
	surname := args[2]

	authorizeUserAsbytes, err := stub.GetState(keydataset) //get the permission
	if err != nil {
		jsonResp = "{\"Error\":\"Failed to check authorization for the dataset: " + keydataset + "\"}"
		return shim.Error(jsonResp)
	} else if authorizeUserAsbytes == nil {
		jsonResp = "{\"Error\":\"The user is not authorized for the dataset: " + keydataset + "\"}"
		return shim.Error(jsonResp)
	}
	authorizedUsers := AuthorizedUsers{}
	err = json.Unmarshal(authorizeUserAsbytes, &authorizedUsers) //unmarshal it aka JSON.parse()
	if err != nil {
		return shim.Error(err.Error())
	}

	authorizedUsers.Users.Name = name
	authorizedUsers.Users.Surname = surname
	datasetInfoJSONasBytes, err := json.Marshal(authorizedUsers)
	if err != nil {
		return shim.Error(err.Error())
	}

	// === Save to state ===
	err = stub.PutState(keydataset, datasetInfoJSONasBytes)
	if err != nil {
		return shim.Error(err.Error())
	}

	return shim.Success(nil)
}
