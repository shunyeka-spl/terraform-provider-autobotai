package pkg

type AzureIntegrations struct {
	Payload struct {
		ID             string      `json:"id,omitempty"`
		Alias          string      `json:"alias"`
		ClientId       string      `json:"clientId"`
		TenantId       string      `json:"tenantId"`
		SecretKey      string      `json:"secretKey"`
		SubscriptionId string      `json:"subscriptionId"`
		Groups         interface{} `json:"groups"`
		CspName        string      `json:"cspName"`
	} `json:"payload"`
}

type DeleteIntegrationsResponse struct {
	Message string `json:"message"`
}
type IntegrationsResponse struct {
	ID             string      `json:"id,omitempty"`
	UserId         string      `json:"userId"`
	AccountId      string      `json:"accountId"`
	CspName        string      `json:"cspName"`
	Alias          string      `json:"alias"`
	Groups         interface{} `json:"groups"`
	AccessToken    string      `json:"accessToken"`
	CreatedAt      string      `json:"createdAt"`
	UpdatedAt      string      `json:"updatedAt"`
	IndexFailures  int         `json:"indexFailures"`
	IsUnauthorized bool        `json:"isUnauthorized"`
	LastUsed       string      `json:"lastUsed"`
}
type ApiList struct {
	Message string `json:"message"`
}

type GCPIntegration struct {
	Payload struct {
		Alias                   string      `json:"alias"`
		Groups                  interface{} `json:"groups"`
		Type                    string      `json:"type"`
		ProjectId               string      `json:"project_id"`
		PrivateKeyId            string      `json:"private_key_id"`
		PrivateKey              string      `json:"private_key"`
		ClientEmail             string      `json:"client_email"`
		ClientId                string      `json:"client_id"`
		AuthUri                 string      `json:"auth_uri"`
		TokenUri                string      `json:"token_uri"`
		AuthProviderX509CertUrl string      `json:"auth_provider_x509_cert_url"`
		ClientC509CertUrl       string      `json:"client_x509_cert_url"`
		CspName                 string      `json:"cspName"`
	} `json:"payload"`
}
type AwsIntegration struct {
	ID           string      `json:"id,omitempty"`
	Alias        string      `json:"alias"`
	Groups       interface{} `json:"groups"`
	DeployedBots bool        `json:"deploy_bots"`
}
type AwsIntegrationResponse struct {
	ID              string `json:"id,omitempty"`
	Success         bool   `json:"success"`
	ExternalId      string `json:"externalId"`
	TargetPrincipal string `json:"targetPrincipal"`
}
type MSTeamsIntegration struct {
	Payload struct {
		Alias   string      `json:"alias"`
		Groups  interface{} `json:"groups"`
		CspName string      `json:"cspName"`
		Webhook string      `json:"webhook"`
	} `json:"payload"`
}
type WorkloadSecurity struct {
	Payload struct {
		Alias      string      `json:"alias"`
		Groups     interface{} `json:"groups"`
		CspName    string      `json:"cspName"`
		ApiVersion string      `json:"apiVersion"`
		SecretKey  string      `json:"secretKey"`
		Url        string      `json:"url"`
	} `json:"payload"`
}
type Conformity struct {
	Payload struct {
		Alias   string      `json:"alias"`
		Groups  interface{} `json:"groups"`
		CspName string      `json:"cspName"`
		ApiKey  string      `json:"apiKey"`
		Region  interface{} `json:"region"`
	} `json:"payload"`
}
type Git struct {
	Payload struct {
		Alias   string      `json:"alias"`
		Groups  interface{} `json:"groups"`
		CspName string      `json:"cspName"`
		Host    string      `json:"host"`
		Token   string      `json:"token"`
	} `json:"payload"`
}
