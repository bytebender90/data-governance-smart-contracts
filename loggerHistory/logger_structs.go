package main

type LoggerRevokedAccess struct {
	ObjectType         string   `json:"docType"`
	Dataset            string   `json:"dataset"`
	UsernameOfProvider string   `json:"usernameOfProvider"`
	DatasetProviderOrg string   `json:"datasetProviderOrg "`
	Description        string   `json:"description"`
	UsernameOfConsumer string   `json:"usernameOfConsumer"`
	Organization       string   `json:"organization"`
	Timestamp          string   `json:"timestamp"`
	RevokedPermission  string   `json:"revokedPermission"`
	CustomAccessRights []string `json:"customAccessRights"`
	Count              int      `json:"count"`
}

type LoggerRevokedAccessCounter struct {
	ObjectType             string `json:"docType"`
	KeyLoggerRevokedAccess string `json:"keyLoggerRevokedAccess"`
	Count                  int    `json:"count"`
}

type LoggerRevokedAccessOrgs struct {
	ObjectType           string   `json:"docType"`
	Dataset              string   `json:"dataset"`
	UsernameOfProvider   string   `json:"usernameOfProvider"`
	DatasetProviderOrg   string   `json:"datasetProviderOrg "`
	Description          string   `json:"description"`
	OrganizationConsumer string   `json:"organizationConsumer"`
	Timestamp            string   `json:"timestamp"`
	RevokedPermission    string   `json:"revokedPermission"`
	CustomAccessRights   []string `json:"customAccessRights"`
	Count                int      `json:"count"`
	Username             string   `json:"username"`
	Name                 string   `json:"name"`
	Surname              string   `json:"surname"`
	Role                 string   `json:"role"`
}

type LoggerRevokedAccessCounterOrg struct {
	ObjectType             string `json:"docType"`
	KeyLoggerRevokedAccess string `json:"keyLoggerRevokedAccess"`
	Count                  int    `json:"count"`
}

type LoggerRevokedAccessPublic struct {
	ObjectType         string   `json:"docType"`
	Dataset            string   `json:"dataset"`
	UsernameOfProvider string   `json:"usernameOfProvider"`
	Timestamp          string   `json:"timestamp"`
	RevokedPermission  string   `json:"revokedPermission"`
	CustomAccessRights []string `json:"customAccessRights"`
	Count              int      `json:"count"`
}

type LoggerRevokedAccessCounterPublic struct {
	ObjectType             string `json:"docType"`
	KeyLoggerRevokedAccess string `json:"keyLoggerRevokedAccess"`
	Count                  int    `json:"count"`
}

type LoggerAccess struct {
	ObjectType         string   `json:"docType"`
	Dataset            string   `json:"dataset"`
	UsernameOfProvider string   `json:"usernameOfProvider"`
	DatasetProviderOrg string   `json:"datasetProviderOrg "`
	Description        string   `json:"description"`
	UsernameOfConsumer string   `json:"usernameOfConsumer"`
	Organization       string   `json:"organization"`
	Timestamp          string   `json:"timestamp"`
	Permission         string   `json:"permission"`
	CustomAccessRights []string `json:"customAccessRights"`
	Status             string   `json:"status"`
	RequestID          string   `json:"requestID"`
	Count              int      `json:"count"`
}

type LoggerAccessCounter struct {
	ObjectType string `json:"docType"`
	Dataset    string `json:"dataset"`
	Count      int    `json:"count"`
}

type LoggerAccessOrgs struct {
	ObjectType           string   `json:"docType"`
	Dataset              string   `json:"dataset"`
	UsernameOfProvider   string   `json:"usernameOfProvider"`
	DatasetProviderOrg   string   `json:"datasetProviderOrg "`
	Description          string   `json:"description"`
	OrganizationConsumer string   `json:"organizationConsumer"`
	Timestamp            string   `json:"timestamp"`
	Permission           string   `json:"permission"`
	CustomAccessRights   []string `json:"customAccessRights"`
	Status               string   `json:"status"`
	RequestID            string   `json:"requestID"`
	Count                int      `json:"count"`
	Username             string   `json:"username"`
	Name                 string   `json:"name"`
	Surname              string   `json:"surname"`
	Role                 string   `json:"role"`
}

type LoggerAccessOrgsCounter struct {
	ObjectType string `json:"docType"`
	Dataset    string `json:"dataset"`
	Count      int    `json:"count"`
}

type LoggerAccessPublic struct {
	ObjectType         string   `json:"docType"`
	Dataset            string   `json:"dataset"`
	UsernameOfProvider string   `json:"usernameOfProvider"`
	Permission         string   `json:"permission"`
	Timestamp          string   `json:"timestamp"`
	CustomAccessRights []string `json:"customAccessRights"`
	Count              int      `json:"count"`
}

type LoggerAccessPublicCounter struct {
	ObjectType string `json:"docType"`
	Dataset    string `json:"dataset"`
	Count      int    `json:"count"`
}

type LoggerMetadata struct {
	ObjectType   string   `json:"docType"`
	Dataset      string   `json:"dataset"`
	Username     string   `json:"username"`
	Organization string   `json:"organization"`
	Changes      []string `json:"changes"`
	Comments     string   `json:"comments"`
	Timestamp    string   `json:"timestamp"`
}

type DataSourceMetadataCounter struct {
	ObjectType         string `json:"docType"`
	DataSourceMetadata string `json:"dataSourceMetadata"`
	Count              int    `json:"count"`
}

type DataSourceMetadata struct {
	ObjectType         string   `json:"docType"`
	DataSourceMetadata string   `json:"dataSourceMetadata"`
	Status             string   `json:"status"`
	Changes            []string `json:"changes"`
	Timestamp          string   `json:"timestamp"`
}

