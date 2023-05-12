package main

import (
	"encoding/json"
	"fmt"
	"strconv"

	"github.com/hyperledger/fabric/core/chaincode/shim"
	pb "github.com/hyperledger/fabric/protos/peer"
)

func (t *Logger) monitorDatasourceMetadata(stub shim.ChaincodeStubInterface, args []string) pb.Response {
	var jsonResp string
	var err error
	if len(args) != 4 {
		return shim.Error("Incorrect arguments. Logger SC for monitoring data source metadata expecting 4 values")
	}

	dataSourceMetadata := args[0]
	status := args[1]
	changes := []string{args[2]}
	timestamp := args[3]
	objectType := "DataSourceMetadata"
	objectTypeKey := "DataSourceMetadataKey"

	if len(dataSourceMetadata) <= 0 {
		return shim.Error("DataSourceMetadata must be a non-empty string")
	}

	if len(status) <= 0 {
		return shim.Error("Status must be a non-empty string")
	}

	if len(timestamp) <= 0 {
		return shim.Error("Timestamp must be a non-empty string")
	}

	monitorDataSourceKeyCounter := "dataSourceMetadataKEY" + dataSourceMetadata
	KeyLoggerMonitorDataSourceJSON := DataSourceMetadataCounter{}
	count := -1
	dataSourceMetadataAsbytes, err := stub.GetState(monitorDataSourceKeyCounter)
	if err != nil {
		jsonResp = "{\"Error\":\"Failed to get data source metadata history for " + monitorDataSourceKeyCounter + "\"}"
		return shim.Error(jsonResp)
	} else if dataSourceMetadataAsbytes == nil {
		count = 0
		dataSourceMetadata := &LoggerRevokedAccessCounter{objectTypeKey, dataSourceMetadata, count}
		dataSourceMetadataAsbytes, err = json.Marshal(dataSourceMetadata)
		if err != nil {
			return shim.Error(err.Error())
		}

	} else {
		err = json.Unmarshal(dataSourceMetadataAsbytes, &KeyLoggerMonitorDataSourceJSON) //unmarshal it aka JSON.parse()
		if err != nil {
			return shim.Error(err.Error())
		}
		count = KeyLoggerMonitorDataSourceJSON.Count + 1
		KeyLoggerMonitorDataSourceJSON.Count = count
		dataSourceMetadataAsbytes, _ = json.Marshal(KeyLoggerMonitorDataSourceJSON)
	}

	// === Save key for data source metadata to state ===
	err = stub.PutState(monitorDataSourceKeyCounter, dataSourceMetadataAsbytes)
	if err != nil {
		return shim.Error(err.Error())
	}

	revokedAccess := &DataSourceMetadata{objectType, dataSourceMetadata, status, changes, timestamp}
	revokedAccessasBytes, err := json.Marshal(revokedAccess)
	if err != nil {
		return shim.Error(err.Error())
	}
	monitorAccessKey := "dataSourceMetadata" + "/" + dataSourceMetadata + "/" + strconv.Itoa(count)
	// === Save revoked access to state ===
	err = stub.PutState(monitorAccessKey, revokedAccessasBytes)
	if err != nil {
		return shim.Error(err.Error())
	}

	return shim.Success(nil)

}

func (t *Logger) queryLoggerAllDatasourceMetadata(stub shim.ChaincodeStubInterface) pb.Response {

	queryString := fmt.Sprintf("{\"selector\":{\"docType\":\"DataSourceMetadata\"}}")
	queryResults, err := getQueryResultForQueryString(stub, queryString)
	if err != nil {
		return shim.Error(err.Error())
	}
	return shim.Success(queryResults)
}

func (t *Logger) queryLoggerDatasourceMetadataByName(stub shim.ChaincodeStubInterface, args []string) pb.Response {
	if len(args) != 1 {
		return shim.Error("Invalid Argument Number. Expecting datasource metadata name")
	}
	dataSourceMetadata := args[0]

	//create index key
	queryString := fmt.Sprintf("{\"selector\":{\"docType\":\"DataSourceMetadata\",\"dataSourceMetadata\":\"%s\"}}", dataSourceMetadata)
	queryResults, err := getQueryResultForQueryString(stub, queryString)
	if err != nil {
		return shim.Error(err.Error())
	}
	return shim.Success(queryResults)
}
