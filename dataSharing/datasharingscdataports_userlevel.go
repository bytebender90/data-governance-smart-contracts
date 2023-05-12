package main

import (
	"encoding/json"
	"fmt"
	"strings"

	"github.com/hyperledger/fabric/core/chaincode/shim"
	pb "github.com/hyperledger/fabric/protos/peer"
)

func (c *ShareData) requestAccess(stub shim.ChaincodeStubInterface, args []string) pb.Response {
	var err error
	var organizationofrequester string
	var jsonResp string
	if len(args) != 5 {
		return shim.Error("Invalid Argument Number for function requestAccess. Expecting 5")
	}
	requestid := args[0]
	datasetname := args[1]
	permission := args[2]
	datetime := args[3]
	customAccessRights := strings.Split(args[4], ",")
	objectType := "RequestAccess"
	userRole, err := c.getRole(stub)
	if err != nil {
		return shim.Error(err.Error())
	}
	/*
	//Check in case is not authorized and creates request only for custom access rights
	if (permission != "read" && permission != "modify" && permission != "persist") { 
		queryResponseAccessRights := c.querySpecificDatasetPermissionByConsumer(stub, []string{datasetname})

		// Check if the response is successful
		if queryResponseAccessRights.Status != shim.OK {
			return shim.Error("Failed to query specific dataset permission by consumer")
		}

		// Check if the response is not empty
		if len(queryResponseAccessRights.Payload) <= 1 {
			return shim.Error("Query result about access rights is empty")
		}
	}*/

	usernameofrequester, err := c.getUsername(stub)
	if err != nil {
		return shim.Error(err.Error())
	}

	keyUserDatasetPermissions := datasetname + "/" + usernameofrequester
	//Check if is already authorized
	checkExistedAuthorizedUser, err := stub.GetState(keyUserDatasetPermissions)
	if err != nil {
		jsonResp = "{\"Error\":\"Failed to check authorization for the dataset: " + datasetname + "\"}"
		return shim.Error(jsonResp)
	} else if (checkExistedAuthorizedUser == nil) && (permission == "" || permission == " "){
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
	
	//get name,surname,organization & username of the attributes values of the user
	nameofrequester, err := c.getName(stub)
	if err != nil {
		return shim.Error(err.Error())
	}

	surnameofrequester, err := c.getSurname(stub)
	if err != nil {
		return shim.Error(err.Error())
	}

	isExternalUser, err := c.getIsExternalUser(stub)
	if err != nil {
		return shim.Error(err.Error())
	}

	if isExternalUser == "false" {
		organizationofrequester, err = c.getOrganization(stub)
		if err != nil {
			return shim.Error(err.Error())
		}

	} else if isExternalUser == "true" {
		organizationofrequester, err = c.getExternalOrg(stub)
		if err != nil {
			return shim.Error(err.Error())
		}

	} else {
		return shim.Error("Unknown int/ext user")
	}


	emailofrequester, err := c.getEmail(stub)
	if err != nil {
		return shim.Error(err.Error())
	}




	/*
	//Invoke Metadata in order to get username Provider
	getMetadataProviderUsername, err := getMetadataProvider(stub, []string{datasetname})
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
	}*/

	requestAccess := RequestAccess{objectType, requestid, datasetname, metadataResponse.UsernameOfProvider, metadataResponse.OrgOfProvider, metadataResponse.DatasetDescription, permission,
		customAccessRights, nameofrequester, surnameofrequester, organizationofrequester, usernameofrequester, emailofrequester, userRole, datetime}
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

func (c *ShareData) requestRevokeAccess(stub shim.ChaincodeStubInterface, args []string) pb.Response {
	var err error
	var organizationofrequester string
	var jsonResp string

	if len(args) != 5 {
		return shim.Error("Invalid Argument Number for function requestRevokeAccess. Expecting 5")
	}
	requestid := args[0]
	datasetname := args[1]
	permission := args[2]
	datetime := args[3]
	customAccessRights := strings.Split(args[4], ",")
	objectType := "Request4RevokedAccess"


	usernameofrequester, err := c.getUsername(stub)
	if err != nil {
		return shim.Error(err.Error())
	}

	datasetKey := datasetname + "/" + usernameofrequester
	authorizeUserAsbytes, err := stub.GetState(datasetKey) //get the permission
	if err != nil {
		jsonResp = "{\"Error\":\"Failed to check authorization for the dataset: " + datasetKey + "\"}"
		return shim.Error(jsonResp)
	} else if authorizeUserAsbytes != nil {
		authorizedUsers := AuthorizedUsers{}
		err = json.Unmarshal(authorizeUserAsbytes, &authorizedUsers) //unmarshal it aka JSON.parse()
		if err != nil {
			return shim.Error(err.Error())
		}
		if checkCustomAccessRights(customAccessRights,authorizedUsers.Users.CustomAccessRights) {
			fmt.Println("Appropriate custom access rights")
		} else {
			jsonResp = "{\"Error\":\"Wrong custom access rights\"}"
			return shim.Error(jsonResp)
		}
	}

	userRole, err := c.getRole(stub)
	if err != nil {
		return shim.Error(err.Error())
	}

	//get name,surname,organization & username of the attributes values of the user
	nameofrequester, err := c.getName(stub)
	if err != nil {
		return shim.Error(err.Error())
	}

	surnameofrequester, err := c.getSurname(stub)
	if err != nil {
		return shim.Error(err.Error())
	}

	isExternalUser, err := c.getIsExternalUser(stub)
	if err != nil {
		return shim.Error(err.Error())
	}

	if isExternalUser == "false" {
		organizationofrequester, err = c.getOrganization(stub)
		if err != nil {
			return shim.Error(err.Error())
		}

	} else if isExternalUser == "true" {
		organizationofrequester, err = c.getExternalOrg(stub)
		if err != nil {
			return shim.Error(err.Error())
		}

	} else {
		return shim.Error("Unknown int/ext user")
	}

	emailofrequester, err := c.getEmail(stub)
	if err != nil {
		return shim.Error(err.Error())
	}



	//Invoke Metadata in order to get username Provider
	getMetadataProviderUsername, err := getMetadataProvider(stub, []string{datasetname})
	if err != nil {
		return shim.Error(err.Error())
	}

	requestRevokeAccess := RequestRevokeAccessByUser{objectType, requestid, datasetname, getMetadataProviderUsername, permission,
		customAccessRights, nameofrequester, surnameofrequester, organizationofrequester, usernameofrequester, emailofrequester, userRole, datetime}
	requestRevokeAccessJSONasBytes, err := json.Marshal(requestRevokeAccess)
	if err != nil {
		return shim.Error(err.Error())
	}

	// === Save to state ===
	err = stub.PutState(requestid, requestRevokeAccessJSONasBytes)
	if err != nil {
		return shim.Error(err.Error())
	}

	messageofdata := "Your revoke request to dataset: " + datasetname + " has been submitted"
	bytemessage := []byte(messageofdata)
	return shim.Success(bytemessage)
}

func (c *ShareData) authorizeUsers(stub shim.ChaincodeStubInterface, args []string) pb.Response {
	var jsonResp string
	if len(args) != 2 {
		return shim.Error("Incorrect number of arguments. Expecting 2. RequestID & timestamp")
	}

	requestID := args[0]
	timestamp := args[1]

	usernameofMetadataProvider, err := c.getUsername(stub)
	if err != nil {
		return shim.Error(err.Error())
	}


	objectType := "AuthorizedUsers"
	usersJSON := RequestAccess{}
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
	name := usersJSON.Name
	surname := usersJSON.Surname
	organization := usersJSON.Organization
	username := usersJSON.Username
	email := usersJSON.Email
	role := usersJSON.Role
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
	datasetKey := datasetname + "/" + username

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

	authorizeUserAsbytes, err := stub.GetState(datasetKey) //get the permission
	if err != nil {
		jsonResp = "{\"Error\":\"Failed to check authorization for the dataset: " + datasetKey + "\"}"
		return shim.Error(jsonResp)
	} else if authorizeUserAsbytes == nil && (permission == "" || permission == " "){
		jsonResp = "{\"Error\":\"The permission is not defined. Failed to find init public dataset: " + datasetname + "\"}"
		return shim.Error(jsonResp)
	}else if authorizeUserAsbytes != nil {
		authorizedUsers := AuthorizedUsers{}
		err = json.Unmarshal(authorizeUserAsbytes, &authorizedUsers) //unmarshal it aka JSON.parse()
		if err != nil {
			return shim.Error(err.Error())
		}

		if permission == "" || permission == " " {
			newPermission = authorizedUsers.Users.Permission
		}

		fmt.Println("Append custom access rights")
		if len(customAccessRights) > 0 {
			customAccessRights = appendCustomAccessRights(customAccessRights,authorizedUsers.Users.CustomAccessRights)
		}

	}
	/*{
		authorizedUsers := AuthorizedUsers{}
		err = json.Unmarshal(authorizeUserAsbytes, &authorizedUsers) //unmarshal it aka JSON.parse()
		if err != nil {
			return shim.Error(err.Error())
		}
		if checkCustomAccessRights(customAccessRights,authorizedUsers.Users.CustomAccessRights) {
			fmt.Println("Appropriate custom access rights")
		} else {
			jsonResp = "{\"Error\":\"Wrong custom access rights\"}"
			return shim.Error(jsonResp)
		}
	}*/


	datasetUser := &CredentialsAuthorizedUsers{newPermission, customAccessRights, name, surname, organization, username, email, role}
	// ==== Create user object for permission and marshal to JSON ====
	datasetInfo := &AuthorizedUsers{objectType, datasetname, datasetProvider, datasetProviderOrg , description, datasetUser}
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
		return shim.Error("Failed to update request access:" + err.Error())
	}

	//Invoke logger
	sendInfoTologger, err := setAccessInfotologger(stub, []string{datasetname, usernameofMetadataProvider, username, organization, requestID, "Accept", timestamp, permission, strings.Join(customAccessRights, ","), datasetProviderOrg, description})
	if err != nil {
		return shim.Error(err.Error())
	}
	fmt.Println(sendInfoTologger)

	messageofdata := "The user " + name + " " + surname + " has been authorized " + "to dataset " + datasetname
	bytemessage := []byte(messageofdata)

	return shim.Success(bytemessage)
}

func (c *ShareData) revokeAccess(stub shim.ChaincodeStubInterface, args []string) pb.Response {
	var jsonResp string
	var err error
	if len(args) != 5 {
		return shim.Error("Incorrect number of arguments for function revokeAccess. Expecting 5")
	}

	datasetname := args[0]
	permission := args[1]
	username := args[2]
	timestamp := args[3]
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
	keydataset := datasetname + "/" + username
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
			}else if permission == "persist" {
				newPermission = []string{"read", "modify"}
			} 

			authorizedUsers.Users.Permission = newPermission

			finalCustomAccessRightsResult := removeElements(authorizedUsers.Users.CustomAccessRights, customAccessRights)
			authorizedUsers.Users.CustomAccessRights = finalCustomAccessRightsResult

			datasetInfoJSONasBytes, err := json.Marshal(authorizedUsers)
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
			}else if permission == "persist" {
				newPermission = []string{"read", "modify"}
			} 

			authorizedUsers.Users.Permission = newPermission

			datasetInfoJSONasBytes, err := json.Marshal(authorizedUsers)
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

		finalCustomAccessRightsResult := removeElements(authorizedUsers.Users.CustomAccessRights, customAccessRights)
		authorizedUsers.Users.CustomAccessRights = finalCustomAccessRightsResult

		datasetInfoJSONasBytes, err := json.Marshal(authorizedUsers)
		if err != nil {
			return shim.Error(err.Error())
		}

		// === Save to state ===
		err = stub.PutState(keydataset, datasetInfoJSONasBytes)
		if err != nil {
			return shim.Error(err.Error())
		}

	}

	//Invoke logger 4 user revoked access
	sendInfoTologger4RevokedAccess, err := setRevokedAccessInfotologger(stub, []string{datasetname, usernameofMetadataProvider, username, authorizedUsers.Users.Organization, permission, timestamp, strings.Join(customAccessRights, ","), authorizedUsers.DatasetProviderOrg, authorizedUsers.Description})
	if err != nil {
		return shim.Error(err.Error())
	}
	fmt.Println(sendInfoTologger4RevokedAccess)

	messageofdata := "The user has been revoked for dataset " + datasetname
	bytemessage := []byte(messageofdata)
	return shim.Success(bytemessage)

}

func (c *ShareData) deleteRequest(stub shim.ChaincodeStubInterface, args []string) pb.Response {
	var jsonResp string
	if len(args) != 2 {
		return shim.Error("Incorrect number of arguments for function deleteRequest. Expecting 2")
	}
	requestID := args[0]
	timestamp := args[1]

	usernameofProvider, err := c.getUsername(stub)
	if err != nil {
		return shim.Error(err.Error())
	}
	usersJSON := RequestAccess{}
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
	err = stub.DelState(requestID) //remove the request based on the requestid from chaincode state
	if err != nil {
		return shim.Error("Failed to delete request access:" + err.Error())
	}
	//if usersJSON.ObjectType != "Request4RevokedAccess" {
	//Invoke logger
	sendInfoTologger, err := setAccessInfotologger(stub, []string{usersJSON.Dataset_Name, usernameofProvider, usersJSON.Username, usersJSON.Organization, requestID, "Deny", timestamp, usersJSON.Permission, strings.Join(usersJSON.CustomAccessRights, ","), usersJSON.DatasetProviderOrg, usersJSON.Description})
	if err != nil {
		return shim.Error(err.Error())
	}
	fmt.Println(sendInfoTologger)
	//}

	messageofdata := "Request: " + usersJSON.RequestID + " from the user " + usersJSON.Name + " " + usersJSON.Surname + " has been deleted " + " for the dataset " + usersJSON.Dataset_Name
	bytemessage := []byte(messageofdata)
	return shim.Success(bytemessage)
}

//only dataset Provider can query permissions
func (c *ShareData) queryDatasetPermissionByProvider(stub shim.ChaincodeStubInterface) pb.Response {
	//create index key
	usernameofProvider, err := c.getUsername(stub)
	if err != nil {
		return shim.Error(err.Error())
	}
	queryString := fmt.Sprintf("{\"selector\":{\"docType\":\"AuthorizedUsers\",\"dataset_provider\":\"%s\"}}", usernameofProvider)
	queryResults, err := getQueryResultForQueryString(stub, queryString)
	if err != nil {
		return shim.Error(err.Error())
	}
	return shim.Success(queryResults)
}

//org admin query permissions for all org users
func (c *ShareData) queryDatasetPermissionForOrgUserByOrgAdmin(stub shim.ChaincodeStubInterface) pb.Response {
	userRole, err := c.getRole(stub)
	if err != nil {
		return shim.Error(err.Error())
	}
	if userRole != "admin" {
		return shim.Error("Only org admin can query permissions for org users")
	}

	organizationofadmin, err := c.getOrganization(stub)
	if err != nil {
		return shim.Error(err.Error())
	}

	queryString := fmt.Sprintf("{\"selector\":{\"docType\":\"AuthorizedUsers\",\"users.organization\":\"%s\"}}", organizationofadmin)
	queryResults, err := getQueryResultForQueryString(stub, queryString)
	if err != nil {
		return shim.Error(err.Error())
	}
	return shim.Success(queryResults)
}

//only dataset Provider can query permissions
func (c *ShareData) querySpecificDatasetPermissionByProvider(stub shim.ChaincodeStubInterface, args []string) pb.Response {
	if len(args) != 1 {
		return shim.Error("Incorrect number of arguments. Expecting 1")
	}
	datasetname := args[0]

	//create index key
	usernameofProvider, err := c.getUsername(stub)
	if err != nil {
		return shim.Error(err.Error())
	}
	queryString := fmt.Sprintf("{\"selector\":{\"docType\":\"AuthorizedUsers\",\"dataset_name\":\"%s\",\"dataset_provider\":\"%s\"}}", datasetname, usernameofProvider)
	queryResults, err := getQueryResultForQueryString(stub, queryString)
	if err != nil {
		return shim.Error(err.Error())
	}
	return shim.Success(queryResults)
}

func (c *ShareData) querySpecificDatasetPermissionByConsumer(stub shim.ChaincodeStubInterface, args []string) pb.Response {
	if len(args) != 1 {
		return shim.Error("Incorrect number of arguments. Expecting 1")
	}
	datasetname := args[0]

	//create index key
	usernameofConsumer, err := c.getUsername(stub)
	if err != nil {
		return shim.Error(err.Error())
	}
	queryString := fmt.Sprintf("{\"selector\":{\"docType\":\"AuthorizedUsers\",\"dataset_name\":\"%s\",\"users.username\":\"%s\"}}", datasetname, usernameofConsumer)
	queryResults, err := getQueryResultForQueryString(stub, queryString)
	if err != nil {
		return shim.Error(err.Error())
	}
	return shim.Success(queryResults)
}

func (c *ShareData) queryDatasetPermissionByAuthorizedUsers(stub shim.ChaincodeStubInterface) pb.Response {
	usernameofAuthorizedUser, err := c.getUsername(stub)
	if err != nil {
		return shim.Error(err.Error())
	}

	queryString := fmt.Sprintf("{\"selector\":{ \"docType\":\"AuthorizedUsers\",\"users.username\": \"%s\"}};", usernameofAuthorizedUser)

	queryResults, err := getQueryResultForQueryString(stub, queryString)
	if err != nil {
		return shim.Error(err.Error())
	}

	return shim.Success(queryResults)
}

//Provider can query request access for a dataset
func (c *ShareData) queryAccessRequestsProviderByName(stub shim.ChaincodeStubInterface, args []string) pb.Response {
	if len(args) != 1 {
		return shim.Error("Invalid Argument Number")
	}
	datasetname := args[0]
	usernameofProvider, err := c.getUsername(stub)
	if err != nil {
		return shim.Error(err.Error())
	}

	//create index key
	queryString := fmt.Sprintf("{\"selector\":{\"docType\":\"RequestAccess\",\"dataset_name\":\"%s\",\"dataset_provider\":\"%s\"}}", datasetname, usernameofProvider)
	queryResults, err := getQueryResultForQueryString(stub, queryString)
	if err != nil {
		return shim.Error(err.Error())
	}

	return shim.Success(queryResults)
}

//Provider can query all request access for his/her dataset
func (c *ShareData) queryAllAccessRequestsProvider(stub shim.ChaincodeStubInterface) pb.Response {
	usernameofProvider, err := c.getUsername(stub)
	if err != nil {
		return shim.Error(err.Error())
	}

	//create index key
	queryString := fmt.Sprintf("{\"selector\":{\"docType\":\"RequestAccess\",\"dataset_provider\":\"%s\"}}", usernameofProvider)
	queryResults, err := getQueryResultForQueryString(stub, queryString)
	if err != nil {
		return shim.Error(err.Error())
	}

	return shim.Success(queryResults)
}

func (c *ShareData) queryAllAccessRequestsConsumer(stub shim.ChaincodeStubInterface) pb.Response {
	usernameofconsumer, err := c.getUsername(stub)
	if err != nil {
		return shim.Error(err.Error())
	}

	//create index key
	queryString := fmt.Sprintf("{\"selector\":{\"docType\":\"RequestAccess\",\"username\":\"%s\"}}", usernameofconsumer)
	queryResults, err := getQueryResultForQueryString(stub, queryString)
	if err != nil {
		return shim.Error(err.Error())
	}

	return shim.Success(queryResults)
}

//Query requests for revoked access

func (c *ShareData) queryAllAccessRevokeByUserConsumer(stub shim.ChaincodeStubInterface) pb.Response {
	usernameofConsumer, err := c.getUsername(stub)
	if err != nil {
		return shim.Error(err.Error())
	}
	//create index key
	queryString := fmt.Sprintf("{\"selector\":{\"docType\":\"Request4RevokedAccess\",\"username\":\"%s\"}}", usernameofConsumer)
	queryResults, err := getQueryResultForQueryString(stub, queryString)
	if err != nil {
		return shim.Error(err.Error())
	}

	return shim.Success(queryResults)
}

func (c *ShareData) queryAccessRevokeUserConsumerByDatasetName(stub shim.ChaincodeStubInterface, args []string) pb.Response {
	if len(args) != 1 {
		return shim.Error("Invalid Argument Number. Expecting dataset name")
	}
	datasetname := args[0]
	usernameofConsumer, err := c.getUsername(stub)
	if err != nil {
		return shim.Error(err.Error())
	}
	//create index key
	queryString := fmt.Sprintf("{\"selector\":{\"docType\":\"Request4RevokedAccess\",\"username\":\"%s\",\"dataset_name\":\"%s\"}}", usernameofConsumer, datasetname)
	queryResults, err := getQueryResultForQueryString(stub, queryString)
	if err != nil {
		return shim.Error(err.Error())
	}

	return shim.Success(queryResults)
}

func (c *ShareData) queryAllRevokeAccessByUserProvider(stub shim.ChaincodeStubInterface) pb.Response {
	usernameofProvider, err := c.getUsername(stub)
	if err != nil {
		return shim.Error(err.Error())
	}

	//create index key
	queryString := fmt.Sprintf("{\"selector\":{\"docType\":\"Request4RevokedAccess\",\"dataset_provider\":\"%s\"}}", usernameofProvider)
	queryResults, err := getQueryResultForQueryString(stub, queryString)
	if err != nil {
		return shim.Error(err.Error())
	}

	return shim.Success(queryResults)
}

func (c *ShareData) queryAccessRevokeUserProviderByDatasetName(stub shim.ChaincodeStubInterface, args []string) pb.Response {
	if len(args) != 1 {
		return shim.Error("Invalid Argument Number. Expecting dataset name")
	}
	datasetname := args[0]
	usernameofProvider, err := c.getUsername(stub)
	if err != nil {
		return shim.Error(err.Error())
	}
	//create index key
	queryString := fmt.Sprintf("{\"selector\":{\"docType\":\"Request4RevokedAccess\",\"dataset_provider\":\"%s\",\"dataset_name\":\"%s\"}}", usernameofProvider, datasetname)
	queryResults, err := getQueryResultForQueryString(stub, queryString)
	if err != nil {
		return shim.Error(err.Error())
	}

	return shim.Success(queryResults)
}
