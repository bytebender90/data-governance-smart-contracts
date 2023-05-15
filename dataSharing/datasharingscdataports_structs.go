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

type DatasetLimits struct {
	ObjectType	string		`json:"docType"`
	Dataset_Name	string		`json:"dataset_name"`
	Dataset_Provider	string		`json:"dataset_provider"`
	DateLimit	string		`json:"date"`
	Fee	int		`json:"fee"`
}

type RequestAccess struct {
	ObjectType	string		`json:"docType"`
	RequestID	string		`json:"requestid"`
	Dataset_Name	string		`json:"dataset_name"`
	Dataset_Provider	string		`json:"dataset_provider"`
	DatasetProviderOrg	string		`json:"datasetProviderOrg "`
	Description	string		`json:"description"`
	Permission	string		`json:"permission"`
	CustomAccessRights	[]string		`json:"customAccessRights"`
	Name	string		`json:"name"`
	Surname	string		`json:"surname"`
	Organization	string		`json:"organization"`
	Username	string		`json:"username"`
	Email	string		`json:"email"`
	Role	string		`json:"role"`
	DateTime	string		`json:"dateTime"`
}

type RequestAccessByOrg struct {
	ObjectType	string		`json:"docType"`
	RequestID	string		`json:"requestid"`
	Dataset_Name	string		`json:"dataset_name"`
	Dataset_Provider	string		`json:"dataset_provider"`
	DatasetProviderOrg	string		`json:"datasetProviderOrg "`
	Description	string		`json:"description"`
	Permission	string		`json:"permission"`
	CustomAccessRights	[]string		`json:"customAccessRights"`
	Organization	string		`json:"organization"`
	DateTime	string		`json:"dateTime"`
	Username	string		`json:"username"`
	Name	string		`json:"name"`
	Surname	string		`json:"surname"`
	Role	string		`json:"role"`
}

type RequestRevokeAccessByOrg struct {
	ObjectType	string		`json:"docType"`
	RequestID	string		`json:"requestid"`
	Dataset_Name	string		`json:"dataset_name"`
	Dataset_Provider	string		`json:"dataset_provider"`
	Permission	string		`json:"permission"`
	CustomAccessRights	[]string		`json:"customAccessRights"`
	Organization	string		`json:"organization"`
	DateTime	string		`json:"dateTime"`
	Username	string		`json:"username"`
	Name	string		`json:"name"`
	Surname	string		`json:"surname"`
	Role	string		`json:"role"`
}

type RequestRevokeAccessByUser struct {
	ObjectType	string		`json:"docType"`
	RequestID	string		`json:"requestid"`
	Dataset_Name	string		`json:"dataset_name"`
	Dataset_Provider	string		`json:"dataset_provider"`
	Permission	string		`json:"permission"`
	CustomAccessRights	[]string		`json:"customAccessRights"`
	Name	string		`json:"name"`
	Surname	string		`json:"surname"`
	Organization	string		`json:"organization"`
	Username	string		`json:"username"`
	Email	string		`json:"email"`
	Role	string		`json:"role"`
	DateTime	string		`json:"dateTime"`
}

type AuthorizedOrgs struct {
	ObjectType	string		`json:"docType"`
	Dataset_Name	string		`json:"dataset_name"`
	Dataset_Provider	string		`json:"dataset_provider"`
	DatasetProviderOrg	string		`json:"datasetProviderOrg "`
	Description	string		`json:"description"`
	Permission	[]string		`json:"permission"`
	CustomAccessRights	[]string		`json:"customAccessRights"`
	Organization	string		`json:"organization"`
	Username	string		`json:"username"`
	Name	string		`json:"name"`
	Surname	string		`json:"surname"`
	Role	string		`json:"role"`
}

type AuthorizedUsers struct {
	ObjectType	string		`json:"docType"`
	Dataset_Name	string		`json:"dataset_name"`
	Dataset_Provider	string		`json:"dataset_provider"`
	DatasetProviderOrg	string		`json:"datasetProviderOrg "`
	Description	string		`json:"description"`
	Users	*CredentialsAuthorizedUsers		`json:"users"`
}

type CredentialsAuthorizedUsers struct {
	Permission	[]string		`json:"permission"`
	CustomAccessRights	[]string		`json:"customAccessRights"`
	Name	string		`json:"name"`
	Surname	string		`json:"surname"`
	Organization	string		`json:"organization"`
	Username	string		`json:"username"`
	Email	string		`json:"email"`
	Role	string		`json:"role"`
}

type PublicDataset struct {
	ObjectType	string		`json:"docType"`
	Dataset_Name	string		`json:"dataset_name"`
	Dataset_Provider	string		`json:"dataset_provider"`
	Permission	[]string		`json:"permission"`
	CustomAccessRights	[]string		`json:"customAccessRights"`
}
