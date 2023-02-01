package sakuracloud

import (
	"testing"

	"github.com/1Password/shell-plugins/sdk"
	"github.com/1Password/shell-plugins/sdk/plugintest"
	"github.com/1Password/shell-plugins/sdk/schema/fieldname"
)

func TestPersonalAccessTokenImporter(t *testing.T) {
	plugintest.TestImporter(t, APIKey().Importer, map[string]plugintest.ImportCase{
		"environment": {
			Environment: map[string]string{
				"SAKURACLOUD_ACCESS_TOKEN":        "xxxxxxxx-xxxx-xxxx-xxxx-xxxxxxxxxxxx",
				"SAKURACLOUD_ACCESS_TOKEN_SECRET": "example-sEcRET",
			},
			ExpectedCandidates: []sdk.ImportCandidate{
				{
					Fields: map[sdk.FieldName]string{
						fieldname.Token: "xxxxxxxx-xxxx-xxxx-xxxx-xxxxxxxxxxxx",
					},
				},
				{
					Fields: map[sdk.FieldName]string{
						fieldname.Secret: "example-sEcRET",
					},
				},
			},
		},
	})
}

func TestPersonalAccessTokenProvisioner(t *testing.T) {
	plugintest.TestProvisioner(t, APIKey().DefaultProvisioner, map[string]plugintest.ProvisionCase{
		"default": {
			ItemFields: map[sdk.FieldName]string{
				fieldname.Token:  "xxxxxxxx-xxxx-xxxx-xxxx-xxxxxxxxxxxx",
				fieldname.Secret: "example-sEcRET",
			},
			ExpectedOutput: sdk.ProvisionOutput{
				Environment: map[string]string{
					"SAKURACLOUD_ACCESS_TOKEN":        "xxxxxxxx-xxxx-xxxx-xxxx-xxxxxxxxxxxx",
					"SAKURACLOUD_ACCESS_TOKEN_SECRET": "example-sEcRET",
				},
			},
		},
	})
}
