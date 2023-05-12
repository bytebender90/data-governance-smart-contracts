package main

import (
	"fmt"

	"github.com/hyperledger/fabric/core/chaincode/shim"
	pb "github.com/hyperledger/fabric/protos/peer"
)

type ShareData struct{}

func (c *ShareData) Init(stub shim.ChaincodeStubInterface) pb.Response {
	return shim.Success(nil)
}

// Implement Invoke
func (c *ShareData) Invoke(stub shim.ChaincodeStubInterface) pb.Response {

	function, args := stub.GetFunctionAndParameters()

	switch function {
	//user level functions
	case "requestAccess": 
		return c.requestAccess(stub, args)
	case "authorizeUsers": 
		return c.authorizeUsers(stub, args)
	case "revokeAccess": 
		return c.revokeAccess(stub, args)
	case "deleteRequest": 
		return c.deleteRequest(stub, args)
	case "requestRevokeAccess": 
		return c.requestRevokeAccess(stub, args)
	case "queryAllAccessRevokeByUserConsumer": 
		return c.queryAllAccessRevokeByUserConsumer(stub)
	case "queryAllRevokeAccessByUserProvider": 
		return c.queryAllRevokeAccessByUserProvider(stub)
	case "queryAccessRevokeUserConsumerByDatasetName": 
		return c.queryAccessRevokeUserConsumerByDatasetName(stub, args)
	case "queryAccessRevokeUserProviderByDatasetName": 
		return c.queryAccessRevokeUserProviderByDatasetName(stub, args)
	case "queryAllAccessRequestsProvider": 
		return c.queryAllAccessRequestsProvider(stub)
	case "queryAccessRequestsProviderByName": 
		return c.queryAccessRequestsProviderByName(stub, args)
	case "queryAllAccessRequestsConsumer": 
		return c.queryAllAccessRequestsConsumer(stub)
	case "querySpecificDatasetPermissionByProvider": 
		return c.querySpecificDatasetPermissionByProvider(stub, args)
	case "querySpecificDatasetPermissionByConsumer": 
		return c.querySpecificDatasetPermissionByConsumer(stub, args)
	case "queryDatasetPermissionByProvider": 
		return c.queryDatasetPermissionByProvider(stub)
	case "queryDatasetPermissionByAuthorizedUsers": 
		return c.queryDatasetPermissionByAuthorizedUsers(stub)
	case "queryDatasetPermissionForOrgUserByOrgAdmin": 
		return c.queryDatasetPermissionForOrgUserByOrgAdmin(stub)
		//organization level functions
	case "queryAccessRequestsOrgConsumerByDatasetName": 
		return c.queryAccessRequestsOrgConsumerByDatasetName(stub, args)
	case "queryAllAccessRequestsByOrgConsumer": 
		return c.queryAllAccessRequestsByOrgConsumer(stub)
	case "queryAccessRequestsOrgProviderByDatasetName": 
		return c.queryAccessRequestsOrgProviderByDatasetName(stub, args)
	case "queryAllAccessRequestsByOrgProvider": 
		return c.queryAllAccessRequestsByOrgProvider(stub)
	case "queryAllAccessByOrgConsumer":
		return c.queryAllAccessByOrgConsumer(stub)
	case "queryAccessOrgConsumerByDatasetName": 
		return c.queryAccessOrgConsumerByDatasetName(stub, args)
	case "queryAllAccessByOrgProvider": 
		return c.queryAllAccessByOrgProvider(stub)
	case "queryAccessOrgProviderByDatasetName": 
		return c.queryAccessOrgProviderByDatasetName(stub, args)
	case "requestAccessByOrg": 
		return c.requestAccessByOrg(stub, args)
	case "authorizeOrgs": 
		return c.authorizeOrgs(stub, args)
	case "deleteRequestOrgs":
		return c.deleteRequestOrgs(stub, args)
	case "revokeAccessOrg":
		return c.revokeAccessOrg(stub, args)
	case "requestRevokeAccessByOrg": 
		return c.requestRevokeAccessByOrg(stub, args)
	case "queryAllAccessRevokeByOrgConsumer": 
		return c.queryAllAccessRevokeByOrgConsumer(stub)
	case "queryAccessRevokeOrgConsumerByDatasetName": 
		return c.queryAccessRevokeOrgConsumerByDatasetName(stub, args)
	case "queryAllRevokeAccessByOrgProvider": 
		return c.queryAllRevokeAccessByOrgProvider(stub)
	case "queryAccessRevokeOrgProviderByDatasetName": 
		return c.queryAccessRevokeOrgProviderByDatasetName(stub, args)
	//public level functions
	case "setDatasetPublic":
		return c.setDatasetPublic(stub, args)
	case "revokeDatasetPublic":
		return c.revokeDatasetPublic(stub, args)
	case "queryPublicDatasets":
		return c.queryPublicDatasets(stub)
	case "queryPublicDatasetsByProvider":
		return c.queryPublicDatasetsByProvider(stub)
	case "queryPublicDatasetsByDatasetName":
		return c.queryPublicDatasetsByDatasetName(stub, args)
	//helper
	case "getRequestByID":
		return c.getRequestByID(stub, args)
	case "updateInfoAuthorizeUsers":
		return c.updateInfoAuthorizeUsers(stub, args)
	}
	return shim.Error("Invalid Command")
}

func main() {
	err := shim.Start(new(ShareData))
	if err != nil {
		fmt.Printf("Error starting chaincode sample: %s", err)
	}
}

