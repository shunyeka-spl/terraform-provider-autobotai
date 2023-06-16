package autobotai

import (
	"context"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"net/http/httptest"

	"os"
	"regexp"

	"autobotAI/pkg"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

var testAccAutobotAIProviders map[string]*schema.Provider
var testAccAutobotAIProvider *schema.Provider

var testServer *httptest.Server

func init() {
	testAccAutobotAIProvider = Provider()
	testAccAutobotAIProvider.ConfigureContextFunc = testProviderConfigure
	testAccAutobotAIProviders = map[string]*schema.Provider{
		"autobotai": testAccAutobotAIProvider,
	}
}

func TestProvider(t *testing.T) {
	if err := Provider().InternalValidate(); err != nil {
		t.Fatalf("err: %s", err)
	}
}
func testAccAutobotAIPreCheck(t *testing.T) {
	os.Setenv("ApiKey", "test-apikey")
	os.Setenv("url", "test-url")
}

func testProviderConfigure(ctx context.Context, d *schema.ResourceData) (interface{}, diag.Diagnostics) {
	var diags diag.Diagnostics
	client, server := createAutobotAIMock()
	testServer = server
	return client, diags

}

func readRequestBody(r *http.Request, payload interface{}) (interface{}, error) {

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(body, &payload)
	if err != nil {
		return nil, err
	}

	return payload, nil
}

func createAutobotAIMock() (*pkg.Client, *httptest.Server) {

	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		//Azure Integration
		var postAutobotAIAzureIntegration = regexp.MustCompile(`^/integrations$`)
		var getAutobotAIAzureIntegration = regexp.MustCompile(`^/integrations/id/(.*)$`)
		var deleteAutobotAIAzureIntegration = regexp.MustCompile(`^/integrations/(.*)$`)

		//Listener
		var postAutobotAIListener = regexp.MustCompile(`^/listeners$`)
		var getAutobotAIListener = regexp.MustCompile(`^/listeners/(.*)$`)
		var putAutobotAIListener = regexp.MustCompile(`^/listeners/(.*)$`)
		var deleteAUtobotAIListener = regexp.MustCompile(`^/listeners/(.*)$`)

		//Fetcher
		var postAutobotAIFetcher = regexp.MustCompile(`^/fetchers$`)
		var getAutobotAIFetcher = regexp.MustCompile(`^/fetchers/(.*)$`)
		var putAutobotAIFetcher = regexp.MustCompile(`^/fetchers/(.*)$`)
		var deleteAutobotAIFetcher = regexp.MustCompile(`^/fetchers/(.*)$`)

		//Automation
		var postAutobotAIAutomation = regexp.MustCompile(`^/automations$`)
		var getAutobotAIAutomation = regexp.MustCompile(`^/automations/(.*)$`)
		var putAutobotAIAutomation = regexp.MustCompile(`^/automations/(.*)$`)
		var deleteAutobotAIAutomation = regexp.MustCompile(`^/automations/(.*)$`)

		//Evaluator
		var postAutobotAIEvaluator = regexp.MustCompile(`^/evaluators$`)
		var getAutobotAIEvaluator = regexp.MustCompile(`^/evaluators/(.*)$`)
		var putAutobotAIEvaluator = regexp.MustCompile(`^/evaluators/(.*)$`)
		var deleteAutobotAIEvaluator = regexp.MustCompile(`^/evaluators/(.*)$`)

		//Bot
		var postAutobotAIBot = regexp.MustCompile(`^/bots$`)
		var getAutobotAIBot = regexp.MustCompile(`^/bots/(.*)$`)
		var putAutobotAIBot = regexp.MustCompile(`^/bots/(.*)$`)
		var deleteAutobotAIBot = regexp.MustCompile(`^/bots/(.*)$`)

		//Inventory Schedule
		var postAutobotAIInventory = regexp.MustCompile(`^/inventory_schedules/$`)
		var deleteAutobotAIInventory = regexp.MustCompile(`^/inventory_schedule/(.*)$`)
		var inventoryResponse = `
		{
			"_id": "6412ec551a2c65c862f70d67",
			"integration_id": "qwedwqdwq_sdnbnghncsdcdsdsc",
			"cron_expression": "0 0 * * *",
			"run_at": "2023-03-17T00:00:00",
			"integration_type": "azure",
			"root_user_id": "amit@shunyeka.com",
			"user_id": "amit@shunyeka.com",
			"created_at": "2023-03-16T10:15:47.419720",
			"updated_at": "2023-03-16T10:15:49.260255"
		}
		`

		var azureIntegrationResponse = `{
			"userId": "user_id",
			"accountId": "account_id",
			"cspName": "azure",
			"alias": "MyProject",
			"groups": [
				"test"
			],
			"accessToken": "",
			"createdAt": "2022-12-23T07:47:36.439839",
			"updatedAt": "2022-12-23T07:47:36.439839",
			"indexFailures": 0,
			"isUnauthorized": false,
			"lastUsed": null
		}`

		var fetcherResponse = `
		{
			"_id": "63ef28572f0ff602dbdcb842",
			"resource_type": "fetcher",
			"name": "testing-fetcher-gcp",
			"arn": "arn:fetcher:e16bef29e64157d797ca9f5d345b0c95:testing-fetcher-gcp",
			"root_user_id": "amit@shunyeka.com",
			"user_id": "amit@shunyeka.com",
			"clients": [
				"gcp_audit_policy"
			],
			"code": "",
			"created_at": "2023-02-17T07:10:15.491000",
			"is_global": false,
			"updated_at": "2023-02-17T07:10:15.491000",
			"integration_type": "gcp",
			"data_schema": null,
			"data_keys": [
				{
					"name": "service.gcp_audit_policy",
					"type": "str"
				},
				{
					"name": "audit_log_configs.gcp_audit_policy",
					"type": "str"
				},
				{
					"name": "akas.gcp_audit_policy",
					"type": "str"
				},
				{
					"name": "location.gcp_audit_policy",
					"type": "str"
				},
				{
					"name": "project.gcp_audit_policy",
					"type": "str"
				}
			],
			"type": "no_code"
		}
		`
		var listenerResponse = `{
			"_id": "63b81e9a0e85d9b882b92903",
			"resource_type": "listener",
			"name": "testing",
			"arn": "arn:listener:e16bef29e64157d797ca9f5d345b0c95:testing",
			"root_user_id": "abc@shunyeka.com",
			"user_id": "abc@shunyeka.com",
			"description": "testing",
			"type": null,
			"created_at": "2023-01-06T13:14:02.141627",
			"webhook_source_integration_type": "conformity",
			"secret": "47ab6d5fc62748499348a9e1bee47a58",
			"updated_at": "2023-01-06T13:14:02.141627",
			"sample_input_json": null,
			"data_schema": null,
			"data_keys": null
		  }
		`
		var automationResponse = `{
			"_id": "63b81f820e85d9b882b9292b",
			"resource_type": "automation",
			"name": "testing",
			"arn": "arn",
			"root_user_id": "abc@shunyeka.com",
			"user_id": "abc@shunyeka.com",
			"title": "testing",
			"code" : "# This automation fetch all resources.\n\n# 'clients' is a dict that contains all the clients mentioned in automation setup\n# 'params' is a dict where key is mentioned in automation setup and value is provided in bot setup\n# 'resources' is a list of resources fetched by fetcher and filtered via evaluator\n# 'test' is a boolean that could be used for conditional testing\n\ndef execute(clients, params, resources=[], test=False):\n    # To get the provided clients, proceed as per Example :\n    resource_client = clients['ResourceManagementClient']\n    # Retrieve the list of resource groups\n    resource_group_response = resource_client.resource_groups.list()\n    # <> Your Code goes here.\n\n    # resources=[]\n    # for resource_group in list(resource_group_response):\n    # # Retrieve the list of resources for each resource group\n    #     resource_list = resource_client.resources.list_by_resource_group(resource_group.name)\n    #     for resource in list(resource_list):\n    #         if resource.type=='Microsoft.Network/networkSecurityGroups':\n    #             resources.append({\n    #                 'id': resource.id,\n    #                 'name':resource.name,\n    #                 'type':resource.type,\n    #                 'region':resource.location,\n    #                 'resource_group_name':resource_group.name\n    #             })\n    # return resources\n",
			"integration_type": "azure",
			"clients": [
			  "ComputeManagementClient"
			],
			"webhook_url": null,
			"type": "communication",
			"required_payload": null,
			"approval_required": true
		  }
		`
		var evaluatorResponse = `{
			"_id": "63bafb3f47a9176af4a97a93",
			"resource_type": "evaluator",
			"name": "testing-bot",
			"arn": "arn:evaluator:e16bef29e64157d797ca9f5d345b0c95:eval-testing-bot-pulkit-faA434",
			"root_user_id": "amit@shunyeka.com",
			"user_id": "amit@shunyeka.com",
			"integration_type": null,
			"eval_details": {
			  "code": "null",
			  "qb_rules": {
				"id": "g-0.04594995406042268",
				"rules": [
					
					{
						"id": "r-0.03355383276017343",
						"field": "type",
						"operator": "is_not_empty",
						"valueSource": "value"
					},
					{
						"id": "r-0.9009325597055955",
						"field": "name",
						"operator": "is_not_empty",
						"valueSource": "value"
					}
				],
				"combinator": "and",
				"not": false
			  },
			  "preference": "qb_rules"
			},
			"created_at": "2023-01-08T17:19:59.770000",
			"updated_at": "2023-01-08T17:20:11.138935"
		  }
			`
		var botResponse = `{
			"_id": "63b6bc1be2fb758eb099aaad",
			"resource_type": "bot",
			"name": "testing",
			"arn": "arn:bot:e16bef29e64157d797ca9f5d345b0c95:testing-bots-e86BcA",
			"root_user_id": "amit@shunyeka.com",
			"user_id": "amit@shunyeka.com",
			"cron_expression": null,
			"topic": "testing",
			"subject": null,
			"description": null,
			"links": null,
			"category": "Reliability",
			"importance": "Medium",
			"examples": null,
			"permissions": null,
			"fetcher_id": "63454da5a6ac05002369749412345",
			"listener_id": null,
			"evaluator_id": "63b6bc1b0e5f0e6ae94dc5e1",
			"actions": [
				{
					"bot_id": null,
					"automation_id": "63171d962268c207516223463456",
					"params": [
						{
							"description": "Name of the bucket",
							"name": "bucket_name",
							"type": "string",
							"values": "autobot-bucket-automation-test"
						}
					],
					"approval_details": {
						"required": false
					},
					"batch_size": null,
					"integration_id": null,
					"integration": null,
					"type": "mutation"
				}
			],
			"integration_id": "e9cb562a-1f12-42c1-8e68-586e140aa6c12345",
			"integration_type": "azure",
			"run_at": null,
			"status": "disabled",
			"fleet_in_sync": true,
			"fleet_id": null
		}`
		switch {

		case postAutobotAIInventory.MatchString(r.URL.Path) && r.Method == "POST":
			w.Write([]byte(inventoryResponse))
		case deleteAutobotAIInventory.MatchString(r.URL.Path) && r.Method == "DELETE":
			w.Write([]byte(`true`))

		//Bot
		case postAutobotAIBot.MatchString(r.URL.Path) && r.Method == "POST":
			w.Write([]byte(botResponse))
		case getAutobotAIBot.MatchString(r.URL.Path) && r.Method == "GET":
			w.Write([]byte(botResponse))
		case putAutobotAIBot.MatchString(r.URL.Path) && r.Method == "PUT":
			w.Write([]byte(botResponse))
		case deleteAutobotAIBot.MatchString(r.URL.Path) && r.Method == "DELETE":
			w.Write([]byte(`true`))

		//Evaluator
		case postAutobotAIEvaluator.MatchString(r.URL.Path) && r.Method == "POST":
			log.Println("[DEBUG] Inside the Post and the response is ", evaluatorResponse)
			w.Write([]byte(evaluatorResponse))
		case getAutobotAIEvaluator.MatchString(r.URL.Path) && r.Method == "GET":
			w.Write([]byte(evaluatorResponse))
		case putAutobotAIEvaluator.MatchString(r.URL.Path) && r.Method == "PUT":
			w.Write([]byte(evaluatorResponse))
		case deleteAutobotAIEvaluator.MatchString(r.URL.Path) && r.Method == "DELETE":
			w.Write([]byte(`true`))

		//Automation
		case postAutobotAIAutomation.MatchString(r.URL.Path) && r.Method == "POST":
			w.Write([]byte(automationResponse))
		case getAutobotAIAutomation.MatchString(r.URL.Path) && r.Method == "GET":
			w.Write([]byte(automationResponse))
		case putAutobotAIAutomation.MatchString(r.URL.Path) && r.Method == "PUT":
			w.Write([]byte(automationResponse))
		case deleteAutobotAIAutomation.MatchString(r.URL.Path) && r.Method == "DELETE":
			w.Write([]byte(`true`))

		//Azure
		case postAutobotAIAzureIntegration.MatchString(r.URL.Path) && r.Method == "POST":
			w.Write([]byte(azureIntegrationResponse))

		case getAutobotAIAzureIntegration.MatchString(r.URL.Path) && r.Method == "GET":
			w.Write([]byte(azureIntegrationResponse))

		case deleteAutobotAIAzureIntegration.MatchString(r.URL.Path) && r.Method == "DELETE":
			w.Write([]byte(`{"deleted":true}`))

		//Listener
		case postAutobotAIListener.MatchString(r.URL.Path) && r.Method == "POST":
			w.Write([]byte(listenerResponse))
		case getAutobotAIListener.MatchString(r.URL.Path) && r.Method == "GET":
			w.Write([]byte(listenerResponse))
		case putAutobotAIListener.MatchString(r.URL.Path) && r.Method == "PUT":
			w.Write([]byte(listenerResponse))
		case deleteAUtobotAIListener.MatchString(r.URL.Path) && r.Method == "DELETE":
			w.Write([]byte(`true`))

		//Fetcher
		case postAutobotAIFetcher.MatchString(r.URL.Path) && r.Method == "POST":
			w.Write([]byte(fetcherResponse))
		case getAutobotAIFetcher.MatchString(r.URL.Path) && r.Method == "GET":
			w.Write([]byte(fetcherResponse))
		case putAutobotAIFetcher.MatchString(r.URL.Path) && r.Method == "PUT":
			w.Write([]byte(fetcherResponse))
		case deleteAutobotAIFetcher.MatchString(r.URL.Path) && r.Method == "DELETE":
			w.Write([]byte(`true`))

		}

	}))
	client := pkg.Client{Apikey: "TEST-APIKEY", Url: server.URL, HttpClient: server.Client()}
	return &client, server
}
