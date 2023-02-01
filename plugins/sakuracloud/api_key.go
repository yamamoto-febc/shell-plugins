package sakuracloud

import (
	"github.com/1Password/shell-plugins/sdk"
	"github.com/1Password/shell-plugins/sdk/importer"
	"github.com/1Password/shell-plugins/sdk/provision"
	"github.com/1Password/shell-plugins/sdk/schema"
	"github.com/1Password/shell-plugins/sdk/schema/credname"
	"github.com/1Password/shell-plugins/sdk/schema/fieldname"
)

func APIKey() schema.CredentialType {
	return schema.CredentialType{
		Name:          credname.APIKey,
		DocsURL:       sdk.URL("https://cloud.sakura.ad.jp"),
		ManagementURL: sdk.URL("https://secure.sakura.ad.jp/cloud/#!/apikey/top"),
		Fields: []schema.CredentialField{
			{
				Name:                fieldname.Token,
				MarkdownDescription: "Token used to authenticate to SAKURA Cloud.",
				Secret:              true,
				Composition: &schema.ValueComposition{
					Length: 36,
					Charset: schema.Charset{
						Lowercase: true,
						Digits:    true,
					},
				},
			},
			{
				Name:                fieldname.Secret,
				MarkdownDescription: "Secret used to authenticate to SAKURA Cloud.",
				Secret:              true,
				Composition: &schema.ValueComposition{
					Length: 65,
					Charset: schema.Charset{
						Lowercase: false,
						Digits:    true,
					},
				},
			},
		},
		DefaultProvisioner: provision.EnvVars(map[string]sdk.FieldName{
			"SAKURACLOUD_ACCESS_TOKEN":        fieldname.Token,
			"SAKURACLOUD_ACCESS_TOKEN_SECRET": fieldname.Secret,
		}),
		Importer: importer.TryAll(
			importer.TryAllEnvVars(fieldname.Token, "SAKURACLOUD_ACCESS_TOKEN"),
			importer.TryAllEnvVars(fieldname.Secret, "SAKURACLOUD_ACCESS_TOKEN_SECRET"),
		),
	}
}
