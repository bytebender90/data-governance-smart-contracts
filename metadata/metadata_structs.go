package main

type DatasetMetadata struct {
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
	CustomAccessRights    []string   `json:"customAccessRights"`
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
