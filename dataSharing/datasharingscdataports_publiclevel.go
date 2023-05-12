package main

import (
	"encoding/json"
	"fmt"
	"strings"

	"github.com/hyperledger/fabric/core/chaincode/shim"
	pb "github.com/hyperledger/fabric/protos/peer"
)

func (c *ShareData) setDatasetPublic(stub shim.ChaincodeStubInterface, args []string) pb.Response {
	var jsonResp string
	var newPermission []string
	if len(args) != 4 {
		return shim.Error("Incorrect number of arguments. Expecting 4. Dataset name, Permission, Custom Access Rights & timestamp")
	}

	datasetname := args[0]
	permission := args[1]
	timestamp := args[2]
	customAccessRights := strings.Split(args[3], ",")

	metadataString, err := getMetadataByDatasetName(stub, []string{datasetname})
	if err != nil {
		return shim.Error(err.Error())
	}


	//check if the custom access rights exists on the dataset metadata
	metadataResponse := &DatasetMetadataResponse{}
	err = json.Unmarshal([]byte(metadataString), metadataResponse)
	if err != nil {
		jsonResp = "{\"Error\":\"Failed to decode JSON" + datasetname + "\"}"
		return shim.Error(jsonResp)
	}

	if checkCustomAccessRights(customAccessRights, metadataResponse.CustomAccessRights) {
		fmt.Println("Appropriate custom access rights")
	} else {
		jsonResp = "{\"Error\":\"Wrong custom access rights\"}"
		return shim.Error(jsonResp)
	}
	

	//Check if it already set with public access
	checkExistedPublicDataset, err := stub.GetState(datasetname)
	if err != nil {
		jsonResp = "{\"Error\":\"Failed to check authorization for the dataset: " + datasetname + "\"}"
		return shim.Error(jsonResp)
	} else if (checkExistedPublicDataset == nil) && (permission == "" || permission == " "){
		jsonResp = "{\"Error\":\"The permission is not defined. Failed to find init public dataset: " + datasetname + "\"}"
		return shim.Error(jsonResp)
	} else if checkExistedPublicDataset != nil{
		publicDataset := PublicDataset{}
		err = json.Unmarshal(checkExistedPublicDataset, &publicDataset) //unmarshal it aka JSON.parse()
		if err != nil {
			return shim.Error(err.Error())
		}
		if permission == "" || permission == " " {
			newPermission = publicDataset.Permission
		}
		fmt.Println("Append custom access rights")
		if len(customAccessRights) > 0 {
			customAccessRights = appendCustomAccessRights(customAccessRights, publicDataset.CustomAccessRights)
		}
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
	
	if permission == "read" {
		newPermission = []string{"read"}
	} else if permission == "modify" {
		newPermission = []string{"read", "modify"}
	} else if permission == "persist" {
		newPermission = []string{"read", "modify", "persist"}
	} 
	
	//Invoke Metadata in order to get endpoint
	objectType := "PublicDatasets"
	datasetInfo := &PublicDataset{objectType, datasetname, usernameofMetadataProvider, newPermission, customAccessRights}
	datasetInfoJSONasBytes, err := json.Marshal(datasetInfo)
	if err != nil {
		return shim.Error(err.Error())
	}
	datasetKey := datasetname
	// === Save to state ===
	err = stub.PutState(datasetKey, datasetInfoJSONasBytes)
	if err != nil {
		return shim.Error(err.Error())
	}

	//Invoke logger
	sendInfoTologger, err := setAccessInfotologgerPublic(stub, []string{datasetname, usernameofMetadataProvider, permission, timestamp, strings.Join(customAccessRights, ",")})
	if err != nil {
		return shim.Error(err.Error())
	}
	fmt.Println(sendInfoTologger)

	messageofdata := "The dataset: " + datasetname + " is public with permission" + permission
	bytemessage := []byte(messageofdata)

	return shim.Success(bytemessage)
}

func (c *ShareData) revokeDatasetPublic(stub shim.ChaincodeStubInterface, args []string) pb.Response {
	var jsonResp string
	var err error
	if len(args) != 4 {
		return shim.Error("Incorrect number of arguments. Expecting 4")
	}

	datasetname := args[0]
	permission := args[1]
	timestamp := args[2]
	customAccessRights := strings.Split(args[3], ",")

	usernameofMetadataProvider, err := c.getUsername(stub)
	if err != nil {
		return shim.Error(err.Error())
	}

	authorizeUserAsbytes, err := stub.GetState(datasetname) //get the permission
	if err != nil {
		jsonResp = "{\"Error\":\"Failed to check authorization for the dataset: " + datasetname + "\"}"
		return shim.Error(jsonResp)
	} else if authorizeUserAsbytes == nil {
		jsonResp = "{\"Error\":\"Failed to find dataset: " + datasetname + "\"}"
		return shim.Error(jsonResp)
	}
	publicDataset := PublicDataset{}
	err = json.Unmarshal(authorizeUserAsbytes, &publicDataset) //unmarshal it aka JSON.parse()
	if err != nil {
		return shim.Error(err.Error())
	}
	if publicDataset.Dataset_Provider != usernameofMetadataProvider {
		return shim.Error("The invoking user is not the provider")
	}


	if checkCustomAccessRights(customAccessRights, publicDataset.CustomAccessRights) {
		fmt.Println("Appropriate custom access rights")
	} else {
		jsonResp = "{\"Error\":\"Wrong custom access rights\"}"
		return shim.Error(jsonResp)
	}



	if (permission == "read") || (permission == "modify") || (permission == "persist") && ((customAccessRights[0] != "") || (customAccessRights[0] != " ")){
		var newPermission []string
		if permission == "read" {
			err = stub.DelState(datasetname) //remove from chaincode state
			if err != nil {
				return shim.Error("Failed to delete state:" + err.Error())
			}
		}else{

		 	if permission == "modify" {
				newPermission = []string{"read"}
			} 	else if permission == "persist" {
				newPermission = []string{"read", "modify"}
			} 

			publicDataset.Permission = newPermission

			finalCustomAccessRightsResult := removeElements(publicDataset.CustomAccessRights, customAccessRights)
			publicDataset.CustomAccessRights = finalCustomAccessRightsResult

			datasetInfoJSONasBytes, err := json.Marshal(publicDataset)
			if err != nil {
				return shim.Error(err.Error())
			}
			// === Save to state ===
			err = stub.PutState(datasetname, datasetInfoJSONasBytes)
			if err != nil {
				return shim.Error(err.Error())
			}
		}
	} else if (permission == "read") || (permission == "modify") || (permission == "persist") {
		var newPermission []string
		if permission == "read" {
			err = stub.DelState(datasetname) //remove from chaincode state
			if err != nil {
				return shim.Error("Failed to delete state:" + err.Error())
			}
		}else{
			if permission == "modify" {
				newPermission = []string{"read"}
			} 	else if permission == "persist" {
				newPermission = []string{"read", "modify"}
			} 

			publicDataset.Permission = newPermission

			datasetInfoJSONasBytes, err := json.Marshal(publicDataset)
			if err != nil {
				return shim.Error(err.Error())
			}
			// === Save to state ===
			err = stub.PutState(datasetname, datasetInfoJSONasBytes)
			if err != nil {
				return shim.Error(err.Error())
			}
		}
	} else if (customAccessRights[0] != "") || (customAccessRights[0] != " "){
		
		finalCustomAccessRightsResult := removeElements(publicDataset.CustomAccessRights, customAccessRights)
		publicDataset.CustomAccessRights = finalCustomAccessRightsResult

		datasetInfoJSONasBytes, err := json.Marshal(publicDataset)
		if err != nil {
			return shim.Error(err.Error())
		}
		// === Save to state ===
		err = stub.PutState(datasetname, datasetInfoJSONasBytes)
		if err != nil {
			return shim.Error(err.Error())
		}

	}

	//Invoke logger
	sendInfoTologger, err := setRevokedAccessPublicInfotologger(stub, []string{datasetname, usernameofMetadataProvider, permission, timestamp, strings.Join(customAccessRights, ",")})
	if err != nil {
		return shim.Error(err.Error())
	}
	fmt.Println(sendInfoTologger)

	messageofdata := "Have been changed access rights for dataset: " + datasetname
	bytemessage := []byte(messageofdata)
	return shim.Success(bytemessage)
}

func (c *ShareData) queryPublicDatasets(stub shim.ChaincodeStubInterface) pb.Response {

	queryString := ("{\"selector\":{\"docType\":\"PublicDatasets\"}}")
	queryResults, err := getQueryResultForQueryString(stub, queryString)
	if err != nil {
		return shim.Error(err.Error())
	}
	return shim.Success(queryResults)
}

func (c *ShareData) queryPublicDatasetsByDatasetName(stub shim.ChaincodeStubInterface, args []string) pb.Response {
	if len(args) != 1 {
		return shim.Error("Invalid Argument Number! Expecting the dataset name")
	}
	datasetName := args[0]

	queryString := fmt.Sprintf("{\"selector\":{\"docType\":\"PublicDatasets\",\"dataset_name\": \"%s\"}}", datasetName)
	queryResults, err := getQueryResultForQueryString(stub, queryString)
	if err != nil {
		return shim.Error(err.Error())
	}
	return shim.Success(queryResults)
}

func (c *ShareData) queryPublicDatasetsByProvider(stub shim.ChaincodeStubInterface) pb.Response {
	usernameofAuthorizedUser, err := c.getUsername(stub)
	if err != nil {
		return shim.Error(err.Error())
	}

	queryString := fmt.Sprintf("{\"selector\":{ \"docType\":\"PublicDatasets\",\"dataset_provider\": \"%s\"}}", usernameofAuthorizedUser)

	queryResults, err := getQueryResultForQueryString(stub, queryString)
	if err != nil {
		return shim.Error(err.Error())
	}

	return shim.Success(queryResults)
}
