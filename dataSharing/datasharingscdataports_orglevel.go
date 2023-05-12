package main

import (
	"encoding/json"
	"fmt"
	"strings"

	"github.com/hyperledger/fabric/core/chaincode/shim"
	pb "github.com/hyperledger/fabric/protos/peer"
)

func (c *ShareData) requestAccessByOrg(stub shim.ChaincodeStubInterface, args []string) pb.Response {
	var err error
	var jsonResp string
	if len(args) != 5 {
		return shim.Error("Invalid Argument Number for function requestAccessByOrg. Expecting 5")
	}
	requestid := args[0]
	datasetname := args[1]
	permission := args[2]
	datetime := args[3]
	customAccessRights := strings.Split(args[4], ",")
	objectType := "RequestAccessByOrg"

	isExternalUser, err := c.getIsExternalUser(stub)
	if err != nil {
		return shim.Error(err.Error())
	}
	if isExternalUser == "true" {
		return shim.Error("External users cannot request organizational access")
	}

	userRole, err := c.getRole(stub)
	if err != nil {
		return shim.Error(err.Error())
	}
	if userRole != "admin" {
		return shim.Error("Only org admin can create organizational access request")
	}

	organizationofrequester, err := c.getOrganization(stub)
	if err != nil {
		return shim.Error(err.Error())
	}

	keyOrgDatasetPermissions := datasetname + "/" + organizationofrequester
	//Check if is already authorized
	checkExistedAuthorizedOrg, err := stub.GetState(keyOrgDatasetPermissions)
	if err != nil {
		jsonResp = "{\"Error\":\"Failed to check authorization for the dataset: " + datasetname + "\"}"
		return shim.Error(jsonResp)
	} else if (checkExistedAuthorizedOrg == nil) && (permission == "" || permission == " "){
		jsonResp = "{\"Error\":\"The permission is not defined. Failed to find init access request for dataset: " + datasetname + "\"}"
		return shim.Error(jsonResp)
	} 

	metadataString, err := getMetadataByDatasetName(stub, []string{datasetname})
	if err != nil {
		return shim.Error(err.Error())
	}

	metadataResponse := &DatasetMetadataResponse{}
	err = json.Unmarshal([]byte(metadataString), metadataResponse)
	if err != nil {
		jsonResp = "{\"Error\":\"Failed to decode JSON" + datasetname + "\"}"
		return shim.Error(jsonResp)
	}

	//check if the custom access rights exists on the dataset metadata
	if checkCustomAccessRights(customAccessRights, metadataResponse.CustomAccessRights) {
		fmt.Println("Appropriate custom access rights")
	} else {
		jsonResp = "{\"Error\":\"Wrong custom access rights\"}"
		return shim.Error(jsonResp)
	}



	nameofrequester, err := c.getName(stub)
	if err != nil {
		return shim.Error(err.Error())
	}

	surnameofrequester, err := c.getSurname(stub)
	if err != nil {
		return shim.Error(err.Error())
	}

	usernameofrequester, err := c.getUsername(stub)
	if err != nil {
		return shim.Error(err.Error())
	}


	//Invoke Metadata in order to get username Provider
	sendInfoToMetadata, err := getMetadataProvider(stub, []string{datasetname})
	if err != nil {
		return shim.Error(err.Error())
	}

	//Invoke Metadata in order to get endpoint
	getEndpointFromMetadata, err := getMetadataEndpoint(stub, []string{datasetname})
	if err != nil {
		return shim.Error(err.Error())
	}

	//Invoke Metadata in order to get KeywordTag
	getKeywordTagFromMetadata, err := getMetadataKeywordTag(stub, []string{datasetname})
	if err != nil {
		return shim.Error(err.Error())
	}

	requestAccess := RequestAccessByOrg{objectType, requestid, datasetname, sendInfoToMetadata, getKeywordTagFromMetadata,
		getEndpointFromMetadata, permission, customAccessRights, organizationofrequester, datetime, usernameofrequester, nameofrequester, surnameofrequester, userRole}
	requestAccessJSONasBytes, err := json.Marshal(requestAccess)
	if err != nil {
		return shim.Error(err.Error())
	}

	// === Save to state ===
	err = stub.PutState(requestid, requestAccessJSONasBytes)
	if err != nil {
		return shim.Error(err.Error())
	}

	messageofdata := "Your request to dataset: " + datasetname + " has been submitted"
	bytemessage := []byte(messageofdata)
	return shim.Success(bytemessage)
}

func (c *ShareData) requestRevokeAccessByOrg(stub shim.ChaincodeStubInterface, args []string) pb.Response {
	var jsonResp string
	if len(args) != 5 {
		return shim.Error("Invalid Argument Number for function requestRevokeAccessByOrg. Expecting 5")
	}
	requestid := args[0]
	datasetname := args[1]
	permission := args[2]
	datetime := args[3]
	customAccessRights := strings.Split(args[4], ",")
	objectType := "Request4RevokedAccessByOrg"

	role, err := c.getRole(stub)
	if err != nil {
		return shim.Error(err.Error())
	}
	if role != "admin" {
		return shim.Error("Only org admin can request revoked access")
	}

	organization, err := c.getOrganization(stub)
	if err != nil {
		return shim.Error(err.Error())
	}

	keyOrgDatasetPermissions := datasetname + "/" + organization
	//Check if is already authorized
	checkExistedAuthorizedOrg, err := stub.GetState(keyOrgDatasetPermissions)
	if err != nil {
		jsonResp = "{\"Error\":\"Failed to check authorization for the dataset: " + datasetname + "\"}"
		return shim.Error(jsonResp)
	} else if checkExistedAuthorizedOrg != nil {
		authorizedOrgs := AuthorizedOrgs{}
		err = json.Unmarshal(checkExistedAuthorizedOrg, &authorizedOrgs) //unmarshal it aka JSON.parse()
		if err != nil {
			return shim.Error(err.Error())
		}
		if checkCustomAccessRights(customAccessRights,authorizedOrgs.CustomAccessRights) {
			fmt.Println("Appropriate custom access rights")
		} else {
			jsonResp = "{\"Error\":\"Wrong custom access rights\"}"
			return shim.Error(jsonResp)
		}
	}

	nameofrequester, err := c.getName(stub)
	if err != nil {
		return shim.Error(err.Error())
	}

	surnameofrequester, err := c.getSurname(stub)
	if err != nil {
		return shim.Error(err.Error())
	}

	usernameofrequester, err := c.getUsername(stub)
	if err != nil {
		return shim.Error(err.Error())
	}


	authOrgsJSON := AuthorizedOrgs{}
	datasetKey := datasetname + "/" + organization
	authOrgAsBytes, err := stub.GetState(datasetKey)
	if err != nil {
		return shim.Error("Failed to get organization access:" + err.Error())
	} else if authOrgAsBytes == nil {
		return shim.Error("The organization is not granted")
	}

	err = json.Unmarshal(authOrgAsBytes, &authOrgsJSON) //unmarshal it aka JSON.parse()
	if err != nil {
		return shim.Error(err.Error())
	}

	if organization != authOrgsJSON.Organization {
		return shim.Error("Unauthorized organization")
	}

	requestRevokeAccessByOrg := &RequestRevokeAccessByOrg{objectType, requestid, datasetname, authOrgsJSON.Dataset_Provider, permission, customAccessRights, organization, datetime, usernameofrequester, nameofrequester, surnameofrequester, role}
	requestRevokeAccessByOrgJSONasBytes, err := json.Marshal(requestRevokeAccessByOrg)
	if err != nil {
		return shim.Error(err.Error())
	}

	// === Save revoke request to state ===
	err = stub.PutState(requestid, requestRevokeAccessByOrgJSONasBytes)
	if err != nil {
		return shim.Error(err.Error())
	}
	messageofdata := "Your revoke request with id: " + requestid + "to dataset: " + datasetname + " has been submitted"
	bytemessage := []byte(messageofdata)
	return shim.Success(bytemessage)

}

func (c *ShareData) authorizeOrgs(stub shim.ChaincodeStubInterface, args []string) pb.Response {
	var jsonResp string
	if len(args) != 2 {
		return shim.Error("Incorrect number of arguments. Expecting 2. RequestID & datetime")
	}

	requestID := args[0]
	datetime := args[1]

	usernameofMetadataProvider, err := c.getUsername(stub)
	if err != nil {
		return shim.Error(err.Error())
	}

	objectType := "AuthorizedOrgs"
	usersJSON := RequestAccessByOrg{}
	valAsbytes, err := stub.GetState(requestID) //get the request from chaincode state
	if err != nil {
		jsonResp = "{\"Error\":\"Failed to get request access for " + requestID + "\"}"
		return shim.Error(jsonResp)
	} else if valAsbytes == nil {
		jsonResp = "{\"Error\":\"Does not exists request for dataset with requestID" + requestID + "\"}"
		return shim.Error(jsonResp)
	}

	err = json.Unmarshal(valAsbytes, &usersJSON)
	if err != nil {
		jsonResp = "{\"Error\":\"Failed to decode JSON" + requestID + "\"}"
		return shim.Error(jsonResp)
	}
	datasetname := usersJSON.Dataset_Name
	datasetProvider := usersJSON.Dataset_Provider
	permission := usersJSON.Permission
	customAccessRights := usersJSON.CustomAccessRights
	organization := usersJSON.Organization
	consumerUsername := usersJSON.Username
	consumerName := usersJSON.Name
	consumerSurname := usersJSON.Surname
	consumerRole := usersJSON.Role
	description := usersJSON.Description
	datasetProviderOrg := usersJSON.DatasetProviderOrg

	//Invoke Metadata in order to check Provider
	sendInfoToMetadata, err := getMetadataProvider(stub, []string{datasetname})
	if err != nil {
		return shim.Error(err.Error())
	}
	if usernameofMetadataProvider != sendInfoToMetadata {
		jsonResp = "{\"Error\":\"Not authorized user:  " + usernameofMetadataProvider + "for dataset: " + datasetname + "\"}"
		return shim.Error(jsonResp)
	}
	datasetKey := datasetname + "/" + organization

	metadataString, err := getMetadataByDatasetName(stub, []string{datasetname})
	if err != nil {
		return shim.Error(err.Error())
	}

	metadataResponse := &DatasetMetadataResponse{}
	err = json.Unmarshal([]byte(metadataString), metadataResponse)
	if err != nil {
		jsonResp = "{\"Error\":\"Failed to decode JSON" + datasetname + "\"}"
		return shim.Error(jsonResp)
	}

	//check if the custom access rights exists on the dataset metadata
	if checkCustomAccessRights(customAccessRights, metadataResponse.CustomAccessRights) {
		fmt.Println("Appropriate custom access rights")
	} else {
		jsonResp = "{\"Error\":\"Wrong custom access rights\"}"
		return shim.Error(jsonResp)
	}


	var newPermission []string
	if permission == "read" {
		newPermission = []string{"read"}
	} else if permission == "modify" {
		newPermission = []string{"read", "modify"}
	} else if permission == "persist" {
		newPermission = []string{"read", "modify", "persist"}
	} else {
		fmt.Println("Not the usual permission")
	}

	authorizeOrgAsbytes, err := stub.GetState(datasetKey) //get the permission
	if err != nil {
		jsonResp = "{\"Error\":\"Failed to check authorization for the dataset: " + datasetKey + "\"}"
		return shim.Error(jsonResp)
	} else if authorizeOrgAsbytes == nil && (permission == "" || permission == " "){
		jsonResp = "{\"Error\":\"The permission is not defined. Failed to find init public dataset: " + datasetname + "\"}"
		return shim.Error(jsonResp)
	}else if authorizeOrgAsbytes != nil {
		authorizedOrgs := AuthorizedOrgs{}
		err = json.Unmarshal(authorizeOrgAsbytes, &authorizedOrgs) //unmarshal it aka JSON.parse()
		if err != nil {
			return shim.Error(err.Error())
		}

		if permission == "" || permission == " " {
			newPermission = authorizedOrgs.Permission
		}

		fmt.Println("Append custom access rights")
		if len(customAccessRights) > 0 {
			customAccessRights = appendCustomAccessRights(customAccessRights,authorizedOrgs.CustomAccessRights)
		}

	}

	datasetInfo := &AuthorizedOrgs{objectType, datasetname, datasetProvider, datasetProviderOrg, description, newPermission, customAccessRights, organization,
		consumerUsername, consumerName, consumerSurname, consumerRole}
	datasetInfoJSONasBytes, err := json.Marshal(datasetInfo)
	if err != nil {
		return shim.Error(err.Error())
	}

	// === Save to state ===
	err = stub.PutState(datasetKey, datasetInfoJSONasBytes)
	if err != nil {
		return shim.Error(err.Error())
	}

	err = stub.DelState(requestID) //remove the request based on the requestid from chaincode state
	if err != nil {
		return shim.Error("Failed to delete organizational access request:" + err.Error())
	}

	//Invoke logger
	sendInfoTologger, err := setAccessInfotologgerOrgs(stub, []string{datasetname, usernameofMetadataProvider, organization, requestID, "Accept", datetime, consumerUsername, consumerName, consumerSurname, consumerRole, permission, strings.Join(customAccessRights, ","), datasetProviderOrg, description})
	if err != nil {
		return shim.Error(err.Error())
	}
	fmt.Println(sendInfoTologger)

	indexName := "organization~dataset_name"
	organizationDatasetIndexKey, err := stub.CreateCompositeKey(indexName, []string{organization, datasetname})
	if err != nil {
		return shim.Error(err.Error())
	}
	value := []byte{0x00}
	err = stub.PutState(organizationDatasetIndexKey, value)
	if err != nil {
		return shim.Error(err.Error())
	}

	messageofdata := "Organization " + organization + " has been authorized " + "to dataset " + datasetname
	bytemessage := []byte(messageofdata)

	return shim.Success(bytemessage)
}

func (c *ShareData) revokeAccessOrg(stub shim.ChaincodeStubInterface, args []string) pb.Response {
	var jsonResp string
	var err error
	if len(args) != 5 {
		return shim.Error("Incorrect number of arguments for function revokeAccessOrg. Expecting 5")
	}

	datasetname := args[0]
	permission := args[1]
	organization := args[2]
	datetime := args[3]
	customAccessRights := strings.Split(args[4], ",")

	usernameofMetadataProvider, err := c.getUsername(stub)
	if err != nil {
		return shim.Error(err.Error())
	}

	metadataString, err := getMetadataByDatasetName(stub, []string{datasetname})
	if err != nil {
		return shim.Error(err.Error())
	}

	metadataResponse := &DatasetMetadataResponse{}
	err = json.Unmarshal([]byte(metadataString), metadataResponse)
	if err != nil {
		jsonResp = "{\"Error\":\"Failed to decode JSON" + datasetname + "\"}"
		return shim.Error(jsonResp)
	}

	//check if the custom access rights exists on the dataset metadata
	if checkCustomAccessRights(customAccessRights, metadataResponse.CustomAccessRights) {
		fmt.Println("Appropriate custom access rights")
	} else {
		jsonResp = "{\"Error\":\"Wrong custom access rights\"}"
		return shim.Error(jsonResp)
	}

	//Invoke Metadata in order to check Provider
	sendInfoToMetadata, err := getMetadataProvider(stub, []string{datasetname})
	if err != nil {
		return shim.Error(err.Error())
	}
	if sendInfoToMetadata != usernameofMetadataProvider {
		return shim.Error("Not authorized to revoke access")
	}
	keydataset := datasetname + "/" + organization
	authorizeUserAsbytes, err := stub.GetState(keydataset) //get the permission
	if err != nil {
		jsonResp = "{\"Error\":\"Failed to check authorization for the dataset: " + keydataset + "\"}"
		return shim.Error(jsonResp)
	} else if authorizeUserAsbytes == nil {
		jsonResp = "{\"Error\":\"Dataset does not exist: " + keydataset + "\"}"
		return shim.Error(jsonResp)
	}
	authorizedOrgs := AuthorizedOrgs{}
	err = json.Unmarshal(authorizeUserAsbytes, &authorizedOrgs) //unmarshal it aka JSON.parse()
	if err != nil {
		return shim.Error(err.Error())
	}

	if (permission == "read") || (permission == "modify") || (permission == "persist") && ((customAccessRights[0] != "") || (customAccessRights[0] != " ")) {
		var newPermission []string
		if permission == "read" {
			err = stub.DelState(keydataset) //remove from chaincode state
			if err != nil {
				return shim.Error("Failed to delete state:" + err.Error())
			}
		}else{ 
			if permission == "modify" {
				newPermission = []string{"read"}
			} else if permission == "persist" {
				newPermission = []string{"read", "modify"}
			}

			authorizedOrgs.Permission = newPermission

			finalCustomAccessRightsResult := removeElements(authorizedOrgs.CustomAccessRights, customAccessRights)
			authorizedOrgs.CustomAccessRights = finalCustomAccessRightsResult

			datasetInfoJSONasBytes, err := json.Marshal(authorizedOrgs)
			if err != nil {
				return shim.Error(err.Error())
			}

			// === Save to state ===
			err = stub.PutState(keydataset, datasetInfoJSONasBytes)
			if err != nil {
				return shim.Error(err.Error())
			}
		}
	} else if (permission == "read") || (permission == "modify") || (permission == "persist") {
		var newPermission []string
		if permission == "read" {
			err = stub.DelState(keydataset) //remove from chaincode state
			if err != nil {
				return shim.Error("Failed to delete state:" + err.Error())
			}
		}else{
			if permission == "modify" {
				newPermission = []string{"read"}
			} else if permission == "persist" {
				newPermission = []string{"read", "modify"}
			}

			authorizedOrgs.Permission = newPermission

			datasetInfoJSONasBytes, err := json.Marshal(authorizedOrgs)
			if err != nil {
				return shim.Error(err.Error())
			}

			// === Save to state ===
			err = stub.PutState(keydataset, datasetInfoJSONasBytes)
			if err != nil {
				return shim.Error(err.Error())
			}
		}

	} else if (customAccessRights[0] != "") || (customAccessRights[0] != " ") {

		finalCustomAccessRightsResult := removeElements(authorizedOrgs.CustomAccessRights, customAccessRights)
		authorizedOrgs.CustomAccessRights = finalCustomAccessRightsResult

		datasetInfoJSONasBytes, err := json.Marshal(authorizedOrgs)
		if err != nil {
			return shim.Error(err.Error())
		}

		// === Save to state ===
		err = stub.PutState(keydataset, datasetInfoJSONasBytes)
		if err != nil {
			return shim.Error(err.Error())
		}
	}

	//Invoke logger 4 org revoked access
	sendInfoTologger4RevokedAccessOrg, err := setRevokedAccessOrgsInfotologger(stub, []string{datasetname, usernameofMetadataProvider, organization, permission, datetime,
		authorizedOrgs.Username, authorizedOrgs.Name, authorizedOrgs.Surname, authorizedOrgs.Role, strings.Join(customAccessRights, ","), authorizedOrgs.DatasetProviderOrg, authorizedOrgs.Description})
	if err != nil {
		return shim.Error(err.Error())
	}
	fmt.Println(sendInfoTologger4RevokedAccessOrg)

	messageofdata := "The organization has been revoked for dataset " + datasetname
	bytemessage := []byte(messageofdata)
	return shim.Success(bytemessage)

}

func (c *ShareData) deleteRequestOrgs(stub shim.ChaincodeStubInterface, args []string) pb.Response {
	var jsonResp string
	if len(args) != 2 {
		return shim.Error("Incorrect number of arguments. Expecting request id & datetime")
	}
	requestID := args[0]
	datetime := args[1]

	usernameofProvider, err := c.getUsername(stub)
	if err != nil {
		return shim.Error(err.Error())
	}
	usersJSON := RequestAccessByOrg{}
	valAsbytes, err := stub.GetState(requestID) //get the request from chaincode state
	if err != nil {
		jsonResp = "{\"Error\":\"Failed to get request access for request ID " + requestID + "\"}"
		return shim.Error(jsonResp)
	} else if valAsbytes == nil {
		jsonResp = "{\"Error\":\"Does not exists request " + requestID + "\"}"
		return shim.Error(jsonResp)

	}

	err = json.Unmarshal(valAsbytes, &usersJSON)
	if err != nil {
		jsonResp = "{\"Error\":\"Failed to decode JSON" + requestID + "\"}"
		return shim.Error(jsonResp)
	}

	if usernameofProvider != usersJSON.Dataset_Provider {
		return shim.Error("Not authorized to delete request")

	}
	datasetname := usersJSON.Dataset_Name
	organization := usersJSON.Organization
	consumerUsername := usersJSON.Username
	consumerName := usersJSON.Name
	consumerSurname := usersJSON.Surname
	consumerRole := usersJSON.Role

	if usersJSON.ObjectType != "Request4RevokedAccessByOrg" {
		//Invoke logger
		sendInfoTologger, err := setAccessInfotologgerOrgs(stub, []string{datasetname, usernameofProvider, organization, requestID, "Deny", datetime, consumerUsername, consumerName, consumerSurname, consumerRole, usersJSON.Permission, strings.Join(usersJSON.CustomAccessRights, ","), usersJSON.DatasetProviderOrg, usersJSON.Description})
		if err != nil {
			return shim.Error(err.Error())
		}
		fmt.Println(sendInfoTologger)

	}

	err = stub.DelState(requestID) //remove the request based on the requestid from chaincode state
	if err != nil {
		return shim.Error("Failed to delete request access:" + err.Error())
	}

	messageofdata := "Request: " + usersJSON.RequestID + " for the dataset " + datasetname + " from the organization " + organization + " has been deleted "
	bytemessage := []byte(messageofdata)
	return shim.Success(bytemessage)
}

func (c *ShareData) queryAllAccessRequestsByOrgProvider(stub shim.ChaincodeStubInterface) pb.Response {
	usernameofProvider, err := c.getUsername(stub)
	if err != nil {
		return shim.Error(err.Error())
	}

	//create index key
	queryString := fmt.Sprintf("{\"selector\":{\"docType\":\"RequestAccessByOrg\",\"dataset_provider\":\"%s\"}}", usernameofProvider)
	queryResults, err := getQueryResultForQueryString(stub, queryString)
	if err != nil {
		return shim.Error(err.Error())
	}

	return shim.Success(queryResults)
}

func (c *ShareData) queryAccessRequestsOrgProviderByDatasetName(stub shim.ChaincodeStubInterface, args []string) pb.Response {
	if len(args) != 1 {
		return shim.Error("Invalid Argument Number. Expecting dataset name")
	}
	datasetname := args[0]
	usernameofProvider, err := c.getUsername(stub)
	if err != nil {
		return shim.Error(err.Error())
	}

	//create index key
	queryString := fmt.Sprintf("{\"selector\":{\"docType\":\"RequestAccessByOrg\",\"dataset_provider\":\"%s\",\"dataset_name\":\"%s\"}}", usernameofProvider, datasetname)
	queryResults, err := getQueryResultForQueryString(stub, queryString)
	if err != nil {
		return shim.Error(err.Error())
	}

	return shim.Success(queryResults)
}

func (c *ShareData) queryAllAccessRequestsByOrgConsumer(stub shim.ChaincodeStubInterface) pb.Response {
	organizationofconsumer, err := c.getOrganization(stub)
	if err != nil {
		return shim.Error(err.Error())
	}

	//create index key
	queryString := fmt.Sprintf("{\"selector\":{\"docType\":\"RequestAccessByOrg\",\"organization\":\"%s\"}}", organizationofconsumer)
	queryResults, err := getQueryResultForQueryString(stub, queryString)
	if err != nil {
		return shim.Error(err.Error())
	}

	return shim.Success(queryResults)
}

func (c *ShareData) queryAccessRequestsOrgConsumerByDatasetName(stub shim.ChaincodeStubInterface, args []string) pb.Response {
	if len(args) != 1 {
		return shim.Error("Invalid Argument Number. Expecting dataset name")
	}
	datasetname := args[0]
	organizationofconsumer, err := c.getOrganization(stub)
	if err != nil {
		return shim.Error(err.Error())
	}
	//create index key
	queryString := fmt.Sprintf("{\"selector\":{\"docType\":\"RequestAccessByOrg\",\"organization\":\"%s\",\"dataset_name\":\"%s\"}}", organizationofconsumer, datasetname)
	queryResults, err := getQueryResultForQueryString(stub, queryString)
	if err != nil {
		return shim.Error(err.Error())
	}

	return shim.Success(queryResults)
}

func (c *ShareData) queryAllAccessByOrgConsumer(stub shim.ChaincodeStubInterface) pb.Response {
	organizationofconsumer, err := c.getOrganization(stub)
	if err != nil {
		return shim.Error(err.Error())
	}

	//create index key
	queryString := fmt.Sprintf("{\"selector\":{\"docType\":\"AuthorizedOrgs\",\"organization\":\"%s\"}}", organizationofconsumer)
	queryResults, err := getQueryResultForQueryString(stub, queryString)
	if err != nil {
		return shim.Error(err.Error())
	}

	return shim.Success(queryResults)
}

func (c *ShareData) queryAccessOrgConsumerByDatasetName(stub shim.ChaincodeStubInterface, args []string) pb.Response {
	if len(args) != 1 {
		return shim.Error("Invalid Argument Number. Expecting dataset name")
	}
	datasetname := args[0]
	organizationofconsumer, err := c.getOrganization(stub)
	if err != nil {
		return shim.Error(err.Error())
	}
	//create index key
	queryString := fmt.Sprintf("{\"selector\":{\"docType\":\"AuthorizedOrgs\",\"organization\":\"%s\",\"dataset_name\":\"%s\"}}", organizationofconsumer, datasetname)
	queryResults, err := getQueryResultForQueryString(stub, queryString)
	if err != nil {
		return shim.Error(err.Error())
	}

	return shim.Success(queryResults)
}

func (c *ShareData) queryAllAccessByOrgProvider(stub shim.ChaincodeStubInterface) pb.Response {
	usernameofProvider, err := c.getUsername(stub)
	if err != nil {
		return shim.Error(err.Error())
	}

	//create index key
	queryString := fmt.Sprintf("{\"selector\":{\"docType\":\"AuthorizedOrgs\",\"dataset_provider\":\"%s\"}}", usernameofProvider)
	queryResults, err := getQueryResultForQueryString(stub, queryString)
	if err != nil {
		return shim.Error(err.Error())
	}

	return shim.Success(queryResults)
}

func (c *ShareData) queryAccessOrgProviderByDatasetName(stub shim.ChaincodeStubInterface, args []string) pb.Response {
	if len(args) != 1 {
		return shim.Error("Invalid Argument Number. Expecting dataset name")
	}
	datasetname := args[0]
	usernameofProvider, err := c.getUsername(stub)
	if err != nil {
		return shim.Error(err.Error())
	}
	//create index key
	queryString := fmt.Sprintf("{\"selector\":{\"docType\":\"AuthorizedOrgs\",\"dataset_provider\":\"%s\",\"dataset_name\":\"%s\"}}", usernameofProvider, datasetname)
	queryResults, err := getQueryResultForQueryString(stub, queryString)
	if err != nil {
		return shim.Error(err.Error())
	}

	return shim.Success(queryResults)
}

//Query requests for revoked access

func (c *ShareData) queryAllAccessRevokeByOrgConsumer(stub shim.ChaincodeStubInterface) pb.Response {
	organizationofconsumer, err := c.getOrganization(stub)
	if err != nil {
		return shim.Error(err.Error())
	}

	//create index key
	queryString := fmt.Sprintf("{\"selector\":{\"docType\":\"Request4RevokedAccessByOrg\",\"organization\":\"%s\"}}", organizationofconsumer)
	queryResults, err := getQueryResultForQueryString(stub, queryString)
	if err != nil {
		return shim.Error(err.Error())
	}

	return shim.Success(queryResults)
}

func (c *ShareData) queryAccessRevokeOrgConsumerByDatasetName(stub shim.ChaincodeStubInterface, args []string) pb.Response {
	if len(args) != 1 {
		return shim.Error("Invalid Argument Number. Expecting dataset name")
	}
	datasetname := args[0]
	organizationofconsumer, err := c.getOrganization(stub)
	if err != nil {
		return shim.Error(err.Error())
	}
	//create index key
	queryString := fmt.Sprintf("{\"selector\":{\"docType\":\"Request4RevokedAccessByOrg\",\"organization\":\"%s\",\"dataset_name\":\"%s\"}}", organizationofconsumer, datasetname)
	queryResults, err := getQueryResultForQueryString(stub, queryString)
	if err != nil {
		return shim.Error(err.Error())
	}

	return shim.Success(queryResults)
}

func (c *ShareData) queryAllRevokeAccessByOrgProvider(stub shim.ChaincodeStubInterface) pb.Response {
	usernameofProvider, err := c.getUsername(stub)
	if err != nil {
		return shim.Error(err.Error())
	}

	//create index key
	queryString := fmt.Sprintf("{\"selector\":{\"docType\":\"Request4RevokedAccessByOrg\",\"dataset_provider\":\"%s\"}}", usernameofProvider)
	queryResults, err := getQueryResultForQueryString(stub, queryString)
	if err != nil {
		return shim.Error(err.Error())
	}

	return shim.Success(queryResults)
}

func (c *ShareData) queryAccessRevokeOrgProviderByDatasetName(stub shim.ChaincodeStubInterface, args []string) pb.Response {
	if len(args) != 1 {
		return shim.Error("Invalid Argument Number. Expecting dataset name")
	}
	datasetname := args[0]
	usernameofProvider, err := c.getUsername(stub)
	if err != nil {
		return shim.Error(err.Error())
	}
	//create index key
	queryString := fmt.Sprintf("{\"selector\":{\"docType\":\"Request4RevokedAccessByOrg\",\"dataset_provider\":\"%s\",\"dataset_name\":\"%s\"}}", usernameofProvider, datasetname)
	queryResults, err := getQueryResultForQueryString(stub, queryString)
	if err != nil {
		return shim.Error(err.Error())
	}

	return shim.Success(queryResults)
}
