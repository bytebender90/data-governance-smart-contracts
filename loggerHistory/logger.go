package main

import (
	"fmt"

	"github.com/hyperledger/fabric/core/chaincode/shim"
	pb "github.com/hyperledger/fabric/protos/peer"
)

type Logger struct{}

func (t *Logger) Init(stub shim.ChaincodeStubInterface) pb.Response {
	fmt.Println("Dataports Transactions log & History")
	return shim.Success(nil)
}

func (t *Logger) Invoke(stub shim.ChaincodeStubInterface) pb.Response {

	function, args := stub.GetFunctionAndParameters()

	switch function {
	case "monitorAccessRequest":
		return t.monitorAccessRequest(stub, args)
	case "queryLoggerAccessProvider":
		return t.queryLoggerAccessProvider(stub)
	case "queryLoggerAccessConsumer":
		return t.queryLoggerAccessConsumer(stub)
	case "queryLoggerAccessProviderByDatasetName":
		return t.queryLoggerAccessProviderByDatasetName(stub, args)
	case "queryLoggerAccessConsumerByDatasetName":
		return t.queryLoggerAccessConsumerByDatasetName(stub, args)
	case "monitorAccessRequestOrgs":
		return t.monitorAccessRequestOrgs(stub, args)
	case "queryLoggerAccessOrgProvider":
		return t.queryLoggerAccessOrgProvider(stub)
	case "queryLoggerAccessOrgConsumer":
		return t.queryLoggerAccessOrgConsumer(stub)
	case "queryLoggerAccessOrgProviderByDatasetName":
		return t.queryLoggerAccessOrgProviderByDatasetName(stub, args)
	case "queryLoggerAccessOrgConsumerByDatasetName":
		return t.queryLoggerAccessOrgConsumerByDatasetName(stub, args)
	case "monitorUploadMetadata":
		return t.monitorUploadMetadata(stub, args)
	case "queryLoggerMetadata":
		return t.queryLoggerMetadata(stub)
	case "monitorRevokedAccessUser":
		return t.monitorRevokedAccessUser(stub, args)
	case "monitorRevokedAccessOrg":
		return t.monitorRevokedAccessOrg(stub, args)
	case "queryLoggerRevokedAccessProvider":
		return t.queryLoggerRevokedAccessProvider(stub)
	case "queryLoggerRevokedAccessConsumer":
		return t.queryLoggerRevokedAccessConsumer(stub)
	case "queryLoggerRevokedAccessProviderByDatasetName":
		return t.queryLoggerRevokedAccessProviderByDatasetName(stub, args)
	case "queryLoggerRevokedAccessConsumerByDatasetName":
		return t.queryLoggerRevokedAccessConsumerByDatasetName(stub, args)
	case "queryLoggerRevokedAccessOrgProvider":
		return t.queryLoggerRevokedAccessOrgProvider(stub)
	case "queryLoggerRevokedAccessOrgConsumer":
		return t.queryLoggerRevokedAccessOrgConsumer(stub)
	case "queryLoggerRevokedAccessOrgProviderByDatasetName":
		return t.queryLoggerRevokedAccessOrgProviderByDatasetName(stub, args)
	case "queryLoggerRevokedAccessOrgConsumerByDatasetName":
		return t.queryLoggerRevokedAccessOrgConsumerByDatasetName(stub, args)
	case "monitorRevokedAccessPublic":
		return t.monitorRevokedAccessPublic(stub, args)
	case "monitorAccessPublic":
		return t.monitorAccessPublic(stub, args)
	case "queryLoggerRevokedAccessPublicProvider":
		return t.queryLoggerRevokedAccessPublicProvider(stub)
	case "queryLoggerRevokedAccessPublicProviderByDatasetName":
		return t.queryLoggerRevokedAccessPublicProviderByDatasetName(stub, args)
	case "queryLoggerAllRevokedAccessPublic":
		return t.queryLoggerAllRevokedAccessPublic(stub)
	case "queryLoggerAllRevokedAccessPublicByDatasetName":
		return t.queryLoggerAllRevokedAccessPublicByDatasetName(stub, args)
	case "queryLoggerAllAccessPublic":
		return t.queryLoggerAllAccessPublic(stub)
	case "queryLoggerAllAccessPublicByDatasetName":
		return t.queryLoggerAllAccessPublicByDatasetName(stub, args)
	case "queryLoggerAllAccessPublicProvider":
		return t.queryLoggerAllAccessPublicProvider(stub)
	case "queryLoggerAllAccessPublicProviderByDatasetName":
		return t.queryLoggerAllAccessPublicProviderByDatasetName(stub, args)
	case "monitorDatasourceMetadata":
		return t.monitorDatasourceMetadata(stub, args)
	case "queryLoggerAllDatasourceMetadata":
		return t.queryLoggerAllDatasourceMetadata(stub)
	case "queryLoggerDatasourceMetadataByName":
		return t.queryLoggerDatasourceMetadataByName(stub, args)
	}

	return shim.Error("Invalid Command")

}

func main() {
	err := shim.Start(new(Logger))
	if err != nil {
		fmt.Printf("Error starting chaincode sample: %s", err)
	}
}
