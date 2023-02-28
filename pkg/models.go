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
type DeleteResponse struct {
	Deleted bool `json:"deleted"`
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
	ExternalId     string      `json:"externalId,omitempty"`
	StackId        string      `json:"stackId,omitempty"`
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
type AwsSesConfigure struct {
	DependsOn string `json:"depends_on"`
	Payload   struct {
		Alias         string `json:"alias"`
		Region        string `json:"region"`
		IntegrationID string `json:"integration_id"`
		CspName       string `json:"cspName"`
	} `json:"payload"`
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

type Fetcher struct {
	Clients          interface{} `json:"clients"`
	Code             string      `json:"code"`
	Integration_Type string      `json:"integration_type"`
	Name             string      `json:"name"`
	Type             string      `json:"type"`
	DataKeys         interface{} `json:"data_keys"`
}

type Listener struct {
	Description                  string `json:"description"`
	Name                         string `json:"name"`
	WebhookSourceIntegrationType string `json:"webhook_source_integration_type"`
}
type Automation struct {
	ApprovalRequired bool        `json:"approval_required"`
	Clients          interface{} `json:"clients"`
	Code             string      `json:"code"`
	Integration_Type string      `json:"integration_type"`
	Name             string      `json:"name"`
	Title            string      `json:"title"`
	Type             string      `json:"type"`
}

type AutomationResponse struct {
	ID               string      `json:"_id,omitempty"`
	ResourceType     string      `json:"resource_type"`
	Name             string      `json:"name"`
	Arn              string      `json:"arn"`
	RootUserId       string      `json:"root_user_id"`
	UserId           string      `json:"user_id"`
	Title            string      `json:"title"`
	Code             string      `json:"code"`
	Type             string      `json:"type"`
	Integration_Type string      `json:"integration_type"`
	Clients          interface{} `json:"clients"`
	WebhookUrl       string      `json:"webhook_url"`
	Requiredpayload  interface{} `json:"required_payload"`
	ApprovalRequired bool        `json:"approval_required"`
}
type FetcherResponse struct {
	ID               string      `json:"_id,omitempty"`
	Arn              string      `json:"arn"`
	Name             string      `json:"name"`
	ResourceType     string      `json:"resource_type"`
	RootUserId       string      `json:"root_user_id"`
	CreatedAt        string      `json:"created_at"`
	UpdatedAt        string      `json:"updated_at"`
	UserId           string      `json:"user_id"`
	DataKeys         interface{} `json:"data_keys"`
	DataSchema       string      `json:"data_schema"`
	Clients          interface{} `json:"clients"`
	Code             string      `json:"code"`
	Integration_Type string      `json:"integration_type"`
	IsGlobal         bool        `json:"is_global"`
	Type             string      `json:"type"`
}

type ListenerResponse struct {
	ID                           string `json:"_id,omitempty"`
	Arn                          string `json:"arn"`
	Name                         string `json:"name"`
	ResourceType                 string `json:"resource_type"`
	RootUserId                   string `json:"root_user_id"`
	CreatedAt                    string `json:"createdAt"`
	UpdatedAt                    string `json:"updatedAt"`
	UserId                       string `json:"user_id"`
	DataKeys                     string `json:"data_keys"`
	DataSchema                   string `json:"data_schema"`
	Description                  string `json:"description"`
	WebhookSourceIntegrationType string `json:"webhook_source_integration_type"`
	SampleInputJson              string `json:"sample_input_json"`
	Secret                       string `json:"secret"`
	Type                         string `json:"type"`
}

type QbRules struct {
	ID         string  `json:"id"`
	Rules      []Rules `json:"rules"`
	Combinator string  `json:"combinator"`
	Not        bool    `json:"not"`
}
type Rules struct {
	ID          string  `json:"id"`
	Field       string  `json:"field,omitempty"`
	Operator    string  `json:"operator,omitempty"`
	ValueSource string  `json:"valueSource,omitempty"`
	Value       string  `json:"value,omitempty"`
	Rules       []Rules `json:"rules,omitempty"`
	Combinator  string  `json:"combinator,omitempty"`
	Not         bool    `json:"not,omitempty"`
}
type Evaluator struct {
	Name        string `json:"name"`
	EvaluatorId string `json:"evaluator_id"`
	EvalDetails struct {
		Code       interface{} `json:"code"`
		QbRules    QbRules     `json:"qb_rules"`
		Preference string      `json:"preference"`
	} `json:"eval_details"`
}

type EvaluatorResponse struct {
	ID              string      `json:"_id"`
	ResourceType    string      `json:"resource_type"`
	Name            string      `json:"name"`
	Arn             string      `json:"arn"`
	RootUserID      string      `json:"root_user_id"`
	UserID          string      `json:"user_id"`
	IntegrationType interface{} `json:"integration_type"`
	EvalDetails     struct {
		Code    string `json:"code"`
		QbRules struct {
			ID         string  `json:"id"`
			Rules      []Rules `json:"rules"`
			Combinator string  `json:"combinator"`
			Not        bool    `json:"not"`
		} `json:"qb_rules"`
		Preference string `json:"preference"`
	} `json:"eval_details"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
}

type Params struct {
	Description string      `json:"description"`
	Name        string      `json:"name"`
	Type        string      `json:"type"`
	Values      interface{} `json:"values"`
}
type Action struct {
	Type            string   `json:"type"`
	AutomationID    string   `json:"automation_id"`
	Params          []Params `json:"params"`
	ApprovalDetails struct {
		Required bool `json:"required"`
	} `json:"approval_details"`
	ActionID string `json:"actionId"`
}
type Bot struct {
	Name            string   `json:"name"`
	Topic           string   `json:"topic"`
	Category        string   `json:"category"`
	Importance      string   `json:"importance"`
	FetcherID       string   `json:"fetcher_id"`
	IntegrationID   string   `json:"integration_id"`
	IntegrationType string   `json:"integration_type"`
	Actions         []Action `json:"actions"`
	EvaluatorID     string   `json:"evaluator_id"`
}

type BotResponse struct {
	ID              string      `json:"_id,omitempty"`
	ResourceType    string      `json:"resource_type"`
	Name            string      `json:"name"`
	Arn             string      `json:"arn"`
	RootUserID      string      `json:"root_user_id"`
	UserID          string      `json:"user_id"`
	CronExpression  interface{} `json:"cron_expression"`
	Topic           string      `json:"topic"`
	Subject         interface{} `json:"subject"`
	Description     interface{} `json:"description"`
	Links           interface{} `json:"links"`
	Category        string      `json:"category"`
	Importance      string      `json:"importance"`
	Examples        interface{} `json:"examples"`
	Permissions     interface{} `json:"permissions"`
	FetcherID       string      `json:"fetcher_id"`
	ListenerID      interface{} `json:"listener_id"`
	EvaluatorID     string      `json:"evaluator_id"`
	Action          []Actions   `json:"actions"`
	IntegrationID   string      `json:"integration_id"`
	IntegrationType string      `json:"integration_type"`
	RunAt           interface{} `json:"run_at"`
	Status          string      `json:"status"`
	FleetInSync     bool        `json:"fleet_in_sync"`
	FleetID         interface{} `json:"fleet_id"`
}

type Actions struct {
	BotID           interface{} `json:"bot_id"`
	AutomationID    string      `json:"automation_id"`
	Params          []Params    `json:"params"`
	ApprovalDetails struct {
		Required bool `json:"required"`
	} `json:"approval_details"`
	BatchSize     interface{} `json:"batch_size"`
	IntegrationID interface{} `json:"integration_id"`
	Integration   interface{} `json:"integration"`
	Type          string      `json:"type"`
	ActionID      string      `json:"actionId"`
}

type GoogleChatIntegration struct {
	Payload struct {
		Alias   string      `json:"alias"`
		Webhook string      `json:"webhook"`
		Groups  interface{} `json:"groups"`
		CspName string      `json:"cspName"`
	} `json:"payload"`
}
