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

type LoggerRevokedAccess struct {
	ObjectType	string		`json:"docType"`
	Dataset	string		`json:"dataset"`
	UsernameOfProvider	string		`json:"usernameOfProvider"`
	DatasetProviderOrg	string		`json:"datasetProviderOrg "`
	Description	string		`json:"description"`
	UsernameOfConsumer	string		`json:"usernameOfConsumer"`
	Organization	string		`json:"organization"`
	Timestamp	string		`json:"timestamp"`
	RevokedPermission	string		`json:"revokedPermission"`
	CustomAccessRights	[]string		`json:"customAccessRights"`
	Count	int		`json:"count"`
}

type LoggerRevokedAccessCounter struct {
	ObjectType	string		`json:"docType"`
	KeyLoggerRevokedAccess	string		`json:"keyLoggerRevokedAccess"`
	Count	int		`json:"count"`
}

type LoggerRevokedAccessOrgs struct {
	ObjectType	string		`json:"docType"`
	Dataset	string		`json:"dataset"`
	UsernameOfProvider	string		`json:"usernameOfProvider"`
	DatasetProviderOrg	string		`json:"datasetProviderOrg "`
	Description	string		`json:"description"`
	OrganizationConsumer	string		`json:"organizationConsumer"`
	Timestamp	string		`json:"timestamp"`
	RevokedPermission	string		`json:"revokedPermission"`
	CustomAccessRights	[]string		`json:"customAccessRights"`
	Count	int		`json:"count"`
	Username	string		`json:"username"`
	Name	string		`json:"name"`
	Surname	string		`json:"surname"`
	Role	string		`json:"role"`
}

type LoggerRevokedAccessCounterOrg struct {
	ObjectType	string		`json:"docType"`
	KeyLoggerRevokedAccess	string		`json:"keyLoggerRevokedAccess"`
	Count	int		`json:"count"`
}

type LoggerRevokedAccessPublic struct {
	ObjectType	string		`json:"docType"`
	Dataset	string		`json:"dataset"`
	UsernameOfProvider	string		`json:"usernameOfProvider"`
	Timestamp	string		`json:"timestamp"`
	RevokedPermission	string		`json:"revokedPermission"`
	CustomAccessRights	[]string		`json:"customAccessRights"`
	Count	int		`json:"count"`
}

type LoggerRevokedAccessCounterPublic struct {
	ObjectType	string		`json:"docType"`
	KeyLoggerRevokedAccess	string		`json:"keyLoggerRevokedAccess"`
	Count	int		`json:"count"`
}

type LoggerAccess struct {
	ObjectType	string		`json:"docType"`
	Dataset	string		`json:"dataset"`
	UsernameOfProvider	string		`json:"usernameOfProvider"`
	DatasetProviderOrg	string		`json:"datasetProviderOrg "`
	Description	string		`json:"description"`
	UsernameOfConsumer	string		`json:"usernameOfConsumer"`
	Organization	string		`json:"organization"`
	Timestamp	string		`json:"timestamp"`
	Permission	string		`json:"permission"`
	CustomAccessRights	[]string		`json:"customAccessRights"`
	Status	string		`json:"status"`
	RequestID	string		`json:"requestID"`
	Count	int		`json:"count"`
}

type LoggerAccessCounter struct {
	ObjectType	string		`json:"docType"`
	Dataset	string		`json:"dataset"`
	Count	int		`json:"count"`
}

type LoggerAccessOrgs struct {
	ObjectType	string		`json:"docType"`
	Dataset	string		`json:"dataset"`
	UsernameOfProvider	string		`json:"usernameOfProvider"`
	DatasetProviderOrg	string		`json:"datasetProviderOrg "`
	Description	string		`json:"description"`
	OrganizationConsumer	string		`json:"organizationConsumer"`
	Timestamp	string		`json:"timestamp"`
	Permission	string		`json:"permission"`
	CustomAccessRights	[]string		`json:"customAccessRights"`
	Status	string		`json:"status"`
	RequestID	string		`json:"requestID"`
	Count	int		`json:"count"`
	Username	string		`json:"username"`
	Name	string		`json:"name"`
	Surname	string		`json:"surname"`
	Role	string		`json:"role"`
}

type LoggerAccessOrgsCounter struct {
	ObjectType	string		`json:"docType"`
	Dataset	string		`json:"dataset"`
	Count	int		`json:"count"`
}

type LoggerAccessPublic struct {
	ObjectType	string		`json:"docType"`
	Dataset	string		`json:"dataset"`
	UsernameOfProvider	string		`json:"usernameOfProvider"`
	Permission	string		`json:"permission"`
	Timestamp	string		`json:"timestamp"`
	CustomAccessRights	[]string		`json:"customAccessRights"`
	Count	int		`json:"count"`
}

type LoggerAccessPublicCounter struct {
	ObjectType	string		`json:"docType"`
	Dataset	string		`json:"dataset"`
	Count	int		`json:"count"`
}

type LoggerMetadata struct {
	ObjectType	string		`json:"docType"`
	Dataset	string		`json:"dataset"`
	Username	string		`json:"username"`
	Organization	string		`json:"organization"`
	Changes	[]string		`json:"changes"`
	Comments	string		`json:"comments"`
	Timestamp	string		`json:"timestamp"`
}

type DataSourceMetadataCounter struct {
	ObjectType	string		`json:"docType"`
	DataSourceMetadata	string		`json:"dataSourceMetadata"`
	Count	int		`json:"count"`
}

type DataSourceMetadata struct {
	ObjectType	string		`json:"docType"`
	DataSourceMetadata	string		`json:"dataSourceMetadata"`
	Status	string		`json:"status"`
	Changes	[]string		`json:"changes"`
	Timestamp	string		`json:"timestamp"`
}
