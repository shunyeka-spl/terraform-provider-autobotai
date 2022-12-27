package integration

import (
	"context"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/http/httptest"

	"os"
	"regexp"

	"autobotai_integration/pkg"
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
	client, server := createAutobotMock()
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

func createAutobotMock() (*pkg.Client, *httptest.Server) {

	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		var postAutobotAzureIntegration = regexp.MustCompile(`^/integrations$`)
		var getAutobotAzureIntegration = regexp.MustCompile(`^/integrations/id/(.*)$`)
		var deleteAutobotAzureIntegration = regexp.MustCompile(`^/integrations/(.*)$`)
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
		switch {

		case postAutobotAzureIntegration.MatchString(r.URL.Path) && r.Method == "POST":

			w.Write([]byte(azureIntegrationResponse))

		case getAutobotAzureIntegration.MatchString(r.URL.Path) && r.Method == "GET":

			w.Write([]byte(azureIntegrationResponse))
		case deleteAutobotAzureIntegration.MatchString(r.URL.Path) && r.Method == "DELETE":

			w.Write([]byte(`{"deleted":true}`))

		}

	}))
	client := pkg.Client{Apikey: "TEST-APIKEY", Url: server.URL, HttpClient: server.Client()}
	return &client, server
}
