package main

type DatasetLimits struct {
	ObjectType       string `json:"docType"`
	Dataset_Name     string `json:"dataset_name"`
	Dataset_Provider string `json:"dataset_provider"`
	DateLimit        string `json:"date"`
	Fee              int    `json:"fee"`
}

type RequestAccess struct {
	ObjectType       string `json:"docType"`
	RequestID        string `json:"requestid"`
	Dataset_Name     string `json:"dataset_name"`
	Dataset_Provider string `json:"dataset_provider"`
	DatasetProviderOrg       string `json:"datasetProviderOrg "`
	Description         string `json:"description"`
	Permission       string `json:"permission"`
	CustomAccessRights []string `json:"customAccessRights"`
	Name             string `json:"name"`
	Surname          string `json:"surname"`
	Organization     string `json:"organization"`
	Username         string `json:"username"`
	Email            string `json:"email"`
	Role             string `json:"role"`
	DateTime         string `json:"dateTime"`
}

type RequestAccessByOrg struct {
	ObjectType       string `json:"docType"`
	RequestID        string `json:"requestid"`
	Dataset_Name     string `json:"dataset_name"`
	Dataset_Provider string `json:"dataset_provider"`
	DatasetProviderOrg       string `json:"datasetProviderOrg "`
	Description         string `json:"description"`
	Permission       string `json:"permission"`
	CustomAccessRights []string `json:"customAccessRights"`
	Organization     string `json:"organization"`
	DateTime         string `json:"dateTime"`
	Username         string `json:"username"`
	Name             string `json:"name"`
	Surname          string `json:"surname"`
	Role             string `json:"role"`
}

type RequestRevokeAccessByOrg struct {
	ObjectType       string `json:"docType"`
	RequestID        string `json:"requestid"`
	Dataset_Name     string `json:"dataset_name"`
	Dataset_Provider string `json:"dataset_provider"`
	Permission       string `json:"permission"`
	CustomAccessRights []string `json:"customAccessRights"`
	Organization     string `json:"organization"`
	DateTime         string `json:"dateTime"`
	Username         string `json:"username"`
	Name             string `json:"name"`
	Surname          string `json:"surname"`
	Role             string `json:"role"`
}

type RequestRevokeAccessByUser struct {
	ObjectType       string `json:"docType"`
	RequestID        string `json:"requestid"`
	Dataset_Name     string `json:"dataset_name"`
	Dataset_Provider string `json:"dataset_provider"`
	Permission       string `json:"permission"`
	CustomAccessRights []string `json:"customAccessRights"`
	Name             string `json:"name"`
	Surname          string `json:"surname"`
	Organization     string `json:"organization"`
	Username         string `json:"username"`
	Email            string `json:"email"`
	Role             string `json:"role"`
	DateTime         string `json:"dateTime"`
}

type AuthorizedOrgs struct {
	ObjectType       string   `json:"docType"`
	Dataset_Name     string   `json:"dataset_name"`
	Dataset_Provider string   `json:"dataset_provider"`
	DatasetProviderOrg       string `json:"datasetProviderOrg "`
	Description         string `json:"description"`
	Permission       []string `json:"permission"`
	CustomAccessRights []string `json:"customAccessRights"`
	Organization     string   `json:"organization"`
	Username         string   `json:"username"`
	Name             string   `json:"name"`
	Surname          string   `json:"surname"`
	Role             string   `json:"role"`
}

type AuthorizedUsers struct {
	ObjectType       string                      `json:"docType"`
	Dataset_Name     string                      `json:"dataset_name"`
	Dataset_Provider string                      `json:"dataset_provider"`
	DatasetProviderOrg       string `json:"datasetProviderOrg "`
	Description         string `json:"description"`
	Users            *CredentialsAuthorizedUsers `json:"users"`
}

type CredentialsAuthorizedUsers struct {
	Permission   []string `json:"permission"`
	CustomAccessRights []string `json:"customAccessRights"`
	Name         string   `json:"name"`
	Surname      string   `json:"surname"`
	Organization string   `json:"organization"`
	Username     string   `json:"username"`
	Email        string   `json:"email"`
	Role         string   `json:"role"`
}

type PublicDataset struct {
	ObjectType         string   `json:"docType"`
	Dataset_Name       string   `json:"dataset_name"`
	Dataset_Provider   string   `json:"dataset_provider"`
	Permission         []string `json:"permission"`
	CustomAccessRights []string `json:"customAccessRights"`
}
