package main

import (
	"encoding/json"
	"fmt"
	"strconv"
	"strings"
	"reflect"

	"github.com/hyperledger/fabric/core/chaincode/shim"
	pb "github.com/hyperledger/fabric/protos/peer"
)

func (c *Metadata) uploadMetadata(stub shim.ChaincodeStubInterface, args []string) pb.Response {
	var err error
	var changes []string
	var orgOfProvider string

	objectType := "Metadata"
	usernameOfProvider := args[0]

	isExternalUser, err := c.getIsExternalUser(stub)
	if err != nil {
		return shim.Error(err.Error())
	}
	if isExternalUser == "false" {
		orgOfProvider = args[1]

	} else if isExternalUser == "true" {
		orgOfProvider, err = c.getExternalOrg(stub)
		if err != nil {
			return shim.Error(err.Error())
		}
	} else {
		return shim.Error("Unknown int/ext user")
	}

	emailOfProvider := args[2]
	usernameOfOwner := args[3]
	orgOfOwner := args[4]
	datasetName := args[5]
	identifier := args[6]
	datasetDescription := args[7]
	releaseDateTime := args[8]
	modificationDateTime := args[9]
	themeCategory := args[10]
	keywordTag := args[11]
	language := args[12]
	distribution := args[13]
	dataVelocity, err := strconv.Atoi(args[14])
	if err != nil {
		fmt.Println(err)
		fmt.Printf("%T \n %v", dataVelocity, dataVelocity)
	}
	spatialGeographicCoverage := args[15]
	temporalCoverageStart := args[16]
	temporalCoverageEnd := args[17]
	industryDomain := args[18]
	dataVolume, err := strconv.Atoi(args[19])
	if err != nil {
		fmt.Println(err)
		fmt.Printf("%T \n %v", dataVolume, dataVolume)
	}
	comments := args[20]
	accessRights := strings.Split(args[21], ",")
	contractAgreementHash := args[22]
	contractAgreementURL := args[23]
	termsConditionsHash := args[24]
	termsConditionsURL := args[25]
	copyrightHash := args[26]
	copyrightURL := args[27]
	customAccessRights := strings.Split(args[28], ",")
	endpoint := args[29]
	blockchain := args[30]
	hasAgent := args[31]
	version := 1

	//DataSourceMetadata
	datasourceID := args[32]
	dataSourceType := args[33]
	dataModelsType := args[34]
	dataModelsValue := args[35]
	dataModelsMetadata := args[36]
	dataProvidedType := args[37]
	dataProvidedValue := args[38]
	dataProvidedMetadata := args[39]
	attributes := args[40]
	service := args[41]
	servicePath := args[42]
	mapping := args[43]
	dataportsDataModelandFormat, err := strconv.ParseBool(args[44])
	if err != nil {
		fmt.Printf("%s: %t\n", args[44], dataportsDataModelandFormat)
	}

	// ==== Check if dataset name already exists ====
	checkNamePartAsBytes, err := stub.GetState(datasetName)
	if err != nil {
		return shim.Error("Failed to check if the name of metadata exists: " + err.Error())
	} else if checkNamePartAsBytes != nil {
		return shim.Error("This name already exists: " + datasetName)
	}

	//Datasource Metadata
	var dataModelsStruct = &DataModelsStruct{
		dataModelsType,
		dataModelsValue,
		dataModelsMetadata,
	}
	var dataProvidedStruct = &DataProvidedStruct{
		dataProvidedType,
		dataProvidedValue,
		dataProvidedMetadata,
	}

	var dataSourceMetadata = &DataSourceMetadata{
		datasourceID,
		dataSourceType,
		dataModelsStruct,
		dataProvidedStruct,
		attributes,
		service,
		servicePath,
		mapping,
		dataportsDataModelandFormat,
	}

	//Data Governance Metadata
	var datasetMetadata = &DatasetMetadata{objectType,
		usernameOfProvider,
		orgOfProvider,
		emailOfProvider,
		usernameOfOwner,
		orgOfOwner,
		datasetName,
		identifier,
		datasetDescription,
		releaseDateTime,
		modificationDateTime,
		themeCategory,
		keywordTag,
		language,
		distribution,
		dataVelocity,
		spatialGeographicCoverage,
		temporalCoverageStart,
		temporalCoverageEnd,
		industryDomain,
		dataVolume,
		comments,
		accessRights,
		contractAgreementHash,
		contractAgreementURL,
		termsConditionsHash,
		termsConditionsURL,
		copyrightHash,
		copyrightURL,
		customAccessRights,
		endpoint,
		blockchain,
		version,
		hasAgent,
		dataSourceMetadata,
	}

	MetadataAsBytes, _ := json.Marshal(datasetMetadata)
	err = stub.PutState(datasetName, MetadataAsBytes)
	if err != nil {
		return shim.Error(err.Error())
	}

	changes = append(changes, "created")
	//InvokeLogger
	sendInfoToLogger, err := setMetadataInfotoLogger(stub, []string{datasetName, usernameOfProvider, orgOfProvider, strings.Join(changes, ","), strconv.Itoa(version), comments, releaseDateTime})
	if err != nil {
		return shim.Error(err.Error())
	}
	fmt.Println(sendInfoToLogger)

	return shim.Success(nil)

}

func (c *Metadata) updateMetadata(stub shim.ChaincodeStubInterface, args []string) pb.Response {
	var err error
	var changes []string
	var orgOfProvider string

	usernameOfProvider := args[0]

	isExternalUser, err := c.getIsExternalUser(stub)
	if err != nil {
		return shim.Error(err.Error())
	}

	if isExternalUser == "false" {
		orgOfProvider = args[1]

	} else if isExternalUser == "true" {
		orgOfProvider, err = c.getExternalOrg(stub)
		if err != nil {
			return shim.Error(err.Error())
		}

	} else {
		return shim.Error("Unknown int/ext user")
	}
	emailOfProvider := args[2]
	usernameOfOwner := args[3]
	orgOfOwner := args[4]
	datasetName := args[5]
	datasetDescription := args[6]
	modificationDateTime := args[7]
	themeCategory := args[8]
	keywordTag := args[9]
	language := args[10]
	distribution := args[11]
	dataVelocity, err := strconv.Atoi(args[12])
	if err != nil {
		fmt.Println(err)
		fmt.Printf("%T \n %v", dataVelocity, dataVelocity)
	}
	spatialGeographicCoverage := args[13]
	temporalCoverageStart := args[14]
	temporalCoverageEnd := args[15]
	industryDomain := args[16]
	dataVolume, err := strconv.Atoi(args[17])
	if err != nil {
		fmt.Println(err)
		fmt.Printf("%T \n %v", dataVolume, dataVolume)
	}
	comments := args[18]
	accessRights := strings.Split(args[19], ",")
	contractAgreementHash := args[20]
	contractAgreementURL := args[21]
	termsConditionsHash := args[22]
	termsConditionsURL := args[23]
	copyrightHash := args[24]
	copyrightURL := args[25]
	customAccessRights := strings.Split(args[26], ",")
	endpoint := args[27]
	blockchain := args[28]
	hasAgent := args[29]

	dataSourceID := args[30]

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
		return shim.Error("Only the data provider" + datasetMetadata.UsernameOfProvider + "can update the dataset! Not the invoking user: " + usernameOfProviderCID)
	}

	if datasetMetadata.UsernameOfOwner != usernameOfOwner {
		changes = append(changes, "username of owner")
	}

	if datasetMetadata.OrgOfOwner != orgOfOwner {
		changes = append(changes, "organization of owner")
	}

	if datasetMetadata.DatasetDescription != datasetDescription {
		changes = append(changes, "description")
	}

	if datasetMetadata.ThemeCategory != themeCategory {
		changes = append(changes, "themeCategory")
	}

	if datasetMetadata.KeywordTag != keywordTag {
		changes = append(changes, "keywordTag")
	}

	if datasetMetadata.Language != language {
		changes = append(changes, "language")
	}

	if datasetMetadata.Distribution != distribution {
		changes = append(changes, "distribution")
	}

	if datasetMetadata.DataVelocity != dataVelocity {
		changes = append(changes, "data velocity")
	}

	if datasetMetadata.SpatialGeographicCoverage != spatialGeographicCoverage {
		changes = append(changes, "spatialGeographicCoverage")
	}

	if datasetMetadata.TemporalCoverageStart != temporalCoverageStart {
		changes = append(changes, "temporalCoverageStart")
	}

	if datasetMetadata.TemporalCoverageEnd != temporalCoverageEnd {
		changes = append(changes, "temporalCoverageEnd")
	}

	if datasetMetadata.IndustryDomain != industryDomain {
		changes = append(changes, "industryDomain")
	}

	if datasetMetadata.DataVolume != dataVolume {
		changes = append(changes, "dataVolume")
	}

	if datasetMetadata.Comments != comments {
		changes = append(changes, "comments")
	}

	if len(datasetMetadata.AccessRights) != len(accessRights) {
		changes = append(changes, "access rights")
	}

	if datasetMetadata.ContractAgreementHash != contractAgreementHash {
		changes = append(changes, "contract agreement hash")
	}

	if datasetMetadata.TermsConditionsHash != termsConditionsHash {
		changes = append(changes, "terms & conditions")
	}

	if datasetMetadata.ContractAgreementURL != contractAgreementURL {
		changes = append(changes, "contract agreements URL")
	}

	if datasetMetadata.TermsConditionsURL != termsConditionsURL {
		changes = append(changes, "terms & conditions URL")
	}

	if datasetMetadata.CopyrightHash != copyrightHash {
		changes = append(changes, "copyrightHash")
	}

	if datasetMetadata.CopyrightURL != copyrightURL {
		changes = append(changes, "copyrightURL")
	}
	checkCustomAccessRights := reflect.DeepEqual(datasetMetadata.CustomAccessRights, customAccessRights)
	if checkCustomAccessRights == false {
		changes = append(changes, "custom access rights")
	}

	if datasetMetadata.Endpoint != endpoint {
		changes = append(changes, "endpoint")
	}

	if datasetMetadata.Blockchain != blockchain {
		changes = append(changes, "blockchain")
	}

	if datasetMetadata.HasAgent != hasAgent {
		changes = append(changes, "agent")
	}

	if datasetMetadata.DataSourceMetadata.DataSourceID != dataSourceID {
		changes = append(changes, "dataSourceID")
	}

	datasetMetadata.UsernameOfProvider = usernameOfProvider
	datasetMetadata.OrgOfProvider = orgOfProvider
	datasetMetadata.EmailOfProvider = emailOfProvider
	datasetMetadata.UsernameOfOwner = usernameOfOwner
	datasetMetadata.OrgOfOwner = orgOfOwner
	datasetMetadata.DatasetName = datasetName
	datasetMetadata.DatasetDescription = datasetDescription
	datasetMetadata.ModificationDateTime = modificationDateTime

	datasetMetadata.ThemeCategory = themeCategory
	datasetMetadata.KeywordTag = keywordTag
	datasetMetadata.Language = language
	datasetMetadata.Distribution = distribution
	datasetMetadata.DataVelocity = dataVelocity
	datasetMetadata.SpatialGeographicCoverage = spatialGeographicCoverage
	datasetMetadata.TemporalCoverageStart = temporalCoverageStart
	datasetMetadata.TemporalCoverageEnd = temporalCoverageEnd
	datasetMetadata.IndustryDomain = industryDomain
	datasetMetadata.DataVolume = dataVolume
	datasetMetadata.Comments = comments

	datasetMetadata.AccessRights = accessRights
	datasetMetadata.ContractAgreementHash = contractAgreementHash
	datasetMetadata.ContractAgreementURL = contractAgreementURL
	datasetMetadata.TermsConditionsHash = termsConditionsHash
	datasetMetadata.TermsConditionsURL = termsConditionsURL
	datasetMetadata.CopyrightHash = copyrightHash
	datasetMetadata.CopyrightURL = copyrightURL
	datasetMetadata.CustomAccessRights = customAccessRights

	datasetMetadata.Endpoint = endpoint
	datasetMetadata.Blockchain = blockchain
	datasetMetadata.Version = datasetMetadata.Version + 1
	datasetMetadata.DataSourceMetadata.DataSourceID = dataSourceID
	datasetMetadata.HasAgent = hasAgent

	MetadataAsBytes, _ := json.Marshal(datasetMetadata)
	err = stub.PutState(datasetName, MetadataAsBytes)
	if err != nil {
		return shim.Error(err.Error())
	}

	//InvokeLogger
	sendInfoToLogger, err := setMetadataInfotoLogger(stub, []string{datasetName, usernameOfProvider, orgOfProvider, strings.Join(changes, ","), strconv.Itoa(datasetMetadata.Version), datasetMetadata.Comments, modificationDateTime})
	if err != nil {
		return shim.Error(err.Error())
	}
	fmt.Println(sendInfoToLogger)
	return shim.Success(nil)

}

func (c *Metadata) createDataSourceMetadata(stub shim.ChaincodeStubInterface, args []string) pb.Response {
	var changes []string
	dateTime := args[0]
	datasetName := args[1]
	datasourceID := args[2]
	dataSourceType := args[3]
	dataModelsType := args[4]
	dataModelsValue := args[5]
	dataModelsMetadata := args[6]
	dataProvidedType := args[7]
	dataProvidedValue := args[8]
	dataProvidedMetadata := args[9]
	attributes := args[10]
	service := args[11]
	servicePath := args[12]
	mapping := args[13]
	dataportsDataModelandFormat, err := strconv.ParseBool(args[14])
	if err != nil {
		fmt.Printf("%s: %t\n", args[14], dataportsDataModelandFormat)
	}

	fmt.Println("- start adding data source metadata for dataset: ", datasetName)
	// ==== Check if dataset name already exists ====
	dataSourceMetadataAsBytes, err := stub.GetState(datasetName)
	if err != nil {
		return shim.Error("Failed to check if the name of metadata exists: " + err.Error())
	} else if dataSourceMetadataAsBytes == nil {
		return shim.Error("The data governance metadata for: " + datasetName + " does not exist")
	}

	dataSourceMetadata := DatasetMetadata{}
	err = json.Unmarshal(dataSourceMetadataAsBytes, &dataSourceMetadata) //unmarshal it aka JSON.parse()
	if err != nil {
		return shim.Error(err.Error())
	}
	dataSourceMetadata.DataSourceMetadata.DataSourceID = datasourceID
	dataSourceMetadata.DataSourceMetadata.DataSourceType = dataSourceType
	dataSourceMetadata.DataSourceMetadata.DataModels.Type = dataModelsType
	dataSourceMetadata.DataSourceMetadata.DataModels.Value = dataModelsValue
	dataSourceMetadata.DataSourceMetadata.DataModels.Metadata = dataModelsMetadata
	dataSourceMetadata.DataSourceMetadata.DataProvided.Type = dataProvidedType
	dataSourceMetadata.DataSourceMetadata.DataProvided.Value = dataProvidedValue
	dataSourceMetadata.DataSourceMetadata.DataProvided.Metadata = dataProvidedMetadata
	dataSourceMetadata.DataSourceMetadata.Attributes = attributes
	dataSourceMetadata.DataSourceMetadata.Service = service
	dataSourceMetadata.DataSourceMetadata.ServicePath = servicePath
	dataSourceMetadata.DataSourceMetadata.Mapping = mapping
	dataSourceMetadata.DataSourceMetadata.DataportsDataModelandFormat = dataportsDataModelandFormat

	status := "created"
	changes = append(changes, "data source metadata added")

	fmt.Println("Invoke logger chaincode")
	//InvokeLogger
	sendInfoToLogger, err := setDataSourceMetadataInfotoLogger(stub, []string{datasetName, status, strings.Join(changes, ","), dateTime})
	if err != nil {
		return shim.Error(err.Error())
	}
	fmt.Println(sendInfoToLogger)

	createDataSourceMetadataAsBytes, _ := json.Marshal(dataSourceMetadata)
	err = stub.PutState(datasetName, createDataSourceMetadataAsBytes)
	if err != nil {
		return shim.Error(err.Error())
	}

	return shim.Success(nil)

}

func (c *Metadata) updateDataSourceMetadata(stub shim.ChaincodeStubInterface, args []string) pb.Response {
	var changes []string
	datasetName := args[0]
	dataModelsType := args[1]
	dataModelsValue := args[2]
	dataModelsMetadata := args[3]
	dataProvidedType := args[4]
	dataProvidedValue := args[5]
	dataProvidedMetadata := args[6]
	attributes := args[7]
	service := args[8]
	servicePath := args[9]
	mapping := args[10]
	dataportsDataModelandFormat, err := strconv.ParseBool(args[11])
	if err != nil {
		fmt.Printf("%s: %t\n", args[11], dataportsDataModelandFormat)
	}
	dateTime := args[12]

	fmt.Println("- start updating data source metadata for dataset: ", datasetName)
	// ==== Check if metadata already exists ====
	dataSourceMetadataAsBytes, err := stub.GetState(datasetName)
	if err != nil {
		return shim.Error("Failed to check if the metadata exist: " + err.Error())
	} else if dataSourceMetadataAsBytes == nil {
		return shim.Error("Does not exist metadata with name: " + datasetName)
	}
	dataSourceMetadata := DatasetMetadata{}
	err = json.Unmarshal(dataSourceMetadataAsBytes, &dataSourceMetadata) //unmarshal it aka JSON.parse()
	if err != nil {
		return shim.Error(err.Error())
	}

	if len(dataModelsType) > 0 && dataSourceMetadata.DataSourceMetadata.DataModels.Type != dataModelsType {
		dataSourceMetadata.DataSourceMetadata.DataModels.Type = dataModelsType
		changes = append(changes, "data Models type")
	}
	if len(dataModelsValue) > 0 && dataSourceMetadata.DataSourceMetadata.DataModels.Value != dataModelsValue {
		dataSourceMetadata.DataSourceMetadata.DataModels.Value = dataModelsValue
		changes = append(changes, "data Models value")
	}
	if len(dataModelsMetadata) > 0 && dataSourceMetadata.DataSourceMetadata.DataModels.Metadata != dataModelsMetadata {
		dataSourceMetadata.DataSourceMetadata.DataModels.Metadata = dataModelsMetadata
		changes = append(changes, "data Models metadata")
	}
	if len(dataProvidedType) > 0 && dataSourceMetadata.DataSourceMetadata.DataProvided.Type != dataProvidedType {
		dataSourceMetadata.DataSourceMetadata.DataProvided.Type = dataProvidedType
		changes = append(changes, "data Provided type")
	}
	if len(dataProvidedValue) > 0 && dataSourceMetadata.DataSourceMetadata.DataProvided.Value != dataProvidedValue {
		dataSourceMetadata.DataSourceMetadata.DataProvided.Value = dataProvidedValue
		changes = append(changes, "data Provided value")
	}
	if len(dataProvidedMetadata) > 0 && dataSourceMetadata.DataSourceMetadata.DataProvided.Metadata != dataProvidedMetadata {
		dataSourceMetadata.DataSourceMetadata.DataProvided.Metadata = dataProvidedMetadata
		changes = append(changes, "data Provided metadata")
	}
	if len(attributes) > 0 && dataSourceMetadata.DataSourceMetadata.Attributes != attributes {
		dataSourceMetadata.DataSourceMetadata.Attributes = attributes
		changes = append(changes, "attributes")
	}
	if len(service) > 0 && dataSourceMetadata.DataSourceMetadata.Service != service {
		dataSourceMetadata.DataSourceMetadata.Service = service
		changes = append(changes, "service")
	}
	if len(servicePath) > 0 && dataSourceMetadata.DataSourceMetadata.ServicePath != servicePath {
		dataSourceMetadata.DataSourceMetadata.ServicePath = servicePath
		changes = append(changes, "service path")
	}
	if len(mapping) > 0 && dataSourceMetadata.DataSourceMetadata.Mapping != mapping {
		dataSourceMetadata.DataSourceMetadata.Mapping = mapping
		changes = append(changes, "mapping")
	}
	if dataSourceMetadata.DataSourceMetadata.DataportsDataModelandFormat != dataportsDataModelandFormat {
		dataSourceMetadata.DataSourceMetadata.DataportsDataModelandFormat = dataportsDataModelandFormat
		changes = append(changes, "compatible Dataports format")
	}

	status := "updated"
	//InvokeLogger
	sendInfoToLogger, err := setDataSourceMetadataInfotoLogger(stub, []string{datasetName, status, strings.Join(changes, ","), dateTime})
	if err != nil {
		return shim.Error(err.Error())
	}
	fmt.Println(sendInfoToLogger)

	updatedMetadataAsBytes, _ := json.Marshal(dataSourceMetadata)
	err = stub.PutState(datasetName, updatedMetadataAsBytes)
	if err != nil {
		return shim.Error(err.Error())
	}

	return shim.Success(nil)

}
