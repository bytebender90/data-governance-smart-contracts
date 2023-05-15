/*
Copyright 2023 Centre for Research and Technology Hellas (CERTH)

Redistribution and use in source and binary forms, with or without modification, are permitted provided that the following conditions are met:

1. Redistributions of source code must retain the above copyright notice, this list of conditions and the following disclaimer.

2. Redistributions in binary form must reproduce the above copyright notice, this list of conditions and the following disclaimer in the documentation and/or other materials provided with the distribution.

3. Neither the name of the copyright holder nor the names of its contributors may be used to endorse or promote products derived from this software without specific prior written permission.

THIS SOFTWARE IS PROVIDED BY THE COPYRIGHT HOLDERS AND CONTRIBUTORS “AS IS” AND ANY EXPRESS OR IMPLIED WARRANTIES, INCLUDING, BUT NOT LIMITED TO, THE IMPLIED WARRANTIES OF MERCHANTABILITY AND FITNESS FOR A PARTICULAR PURPOSE ARE DISCLAIMED. IN NO EVENT SHALL THE COPYRIGHT HOLDER OR CONTRIBUTORS BE LIABLE FOR ANY DIRECT, INDIRECT, INCIDENTAL, SPECIAL, EXEMPLARY, OR CONSEQUENTIAL DAMAGES (INCLUDING, BUT NOT LIMITED TO, PROCUREMENT OF SUBSTITUTE GOODS OR SERVICES; LOSS OF USE, DATA, OR PROFITS; OR BUSINESS INTERRUPTION) HOWEVER CAUSED AND ON ANY THEORY OF LIABILITY, WHETHER IN CONTRACT, STRICT LIABILITY, OR TORT (INCLUDING NEGLIGENCE OR OTHERWISE) ARISING IN ANY WAY OUT OF THE USE OF THIS SOFTWARE, EVEN IF ADVISED OF THE POSSIBILITY OF SUCH DAMAGE.
*/
//line :1
package main

type DatasetMetadata struct {
	ObjectType	string		`json:"docType"`
	                           
	UsernameOfProvider	string		`json:"usernameOfProvider"`
	OrgOfProvider	string		`json:"orgOfProvider"`
	EmailOfProvider	string		`json:"emailOfProvider"`
	UsernameOfOwner	string		`json:"usernameOfOwner"`
	OrgOfOwner	string		`json:"orgOfOwner"`
	                        
	DatasetName	string		`json:"datasetName"`
	Identifier	string		`json:"identifier"`
	DatasetDescription	string		`json:"datasetDescription"`
	ReleaseDateTime	string		`json:"releaseDateTime"`
	ModificationDateTime	string		`json:"modificationDateTime"`
	                        
	ThemeCategory	string		`json:"themeCategory"`
	KeywordTag	string		`json:"keywordTag"`
	Language	string		`json:"language"`
	Distribution	string		`json:"distribution"`
	DataVelocity	int		`json:"dataVelocity"`
	SpatialGeographicCoverage	string		`json:"spatialGeographicCoverage"`
	TemporalCoverageStart	string		`json:"temporalCoverageStart"`
	TemporalCoverageEnd	string		`json:"temporalCoverageEnd"`
	IndustryDomain	string		`json:"industryDomain"`
	DataVolume	int		`json:"dataVolume"`
	Comments	string		`json:"comments"`
	                           
	AccessRights	[]string		`json:"accessRights"`
	ContractAgreementHash	string		`json:"contractAgreementHash"`
	ContractAgreementURL	string		`json:"contractAgreementURL"`
	TermsConditionsHash	string		`json:"termsConditionsHash"`
	TermsConditionsURL	string		`json:"termsConditionsURL"`
	CopyrightHash	string		`json:"copyrightHash"`
	CopyrightURL	string		`json:"copyrightsURL"`
	CustomAccessRights	[]string		`json:"customAccessRights"`
	                                
	Endpoint	string		`json:"endpoint"`
	Blockchain	string		`json:"blockchain"`
	Version	int		`json:"version"`
	HasAgent	string		`json:"hasAgent"`
	                      
	DataSourceMetadata	*DataSourceMetadata		`json:"dataSourceMetadata"`
}

type DataSourceMetadata struct {
	DataSourceID	string		`json:"dataSourceID"`
	DataSourceType	string		`json:"dataSourceType"`
	DataModels	*DataModelsStruct		`json:"dataModels"`
	DataProvided	*DataProvidedStruct		`json:"dataProvided"`
	Attributes	string		`json:"attributes"`
	Service	string		`json:"service"`
	ServicePath	string		`json:"servicePath"`
	Mapping	string		`json:"mapping"`
	DataportsDataModelandFormat	bool		`json:"isDataPortsCompatible"`
}

type DataModelsStruct struct {
	Type	string		`json:"type"`
	Value	string		`json:"value"`
	Metadata	string		`json:"metadata"`
}

type DataProvidedStruct struct {
	Type	string		`json:"type"`
	Value	string		`json:"value"`
	Metadata	string		`json:"metadata"`
}
