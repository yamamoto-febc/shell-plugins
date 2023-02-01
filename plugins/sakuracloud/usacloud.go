package sakuracloud

import (
	"github.com/1Password/shell-plugins/sdk"
	"github.com/1Password/shell-plugins/sdk/needsauth"
	"github.com/1Password/shell-plugins/sdk/schema"
	"github.com/1Password/shell-plugins/sdk/schema/credname"
)

func SakuraCloudCLI() schema.Executable {
	return schema.Executable{
		Name:    "SAKURA Cloud CLI",
		Runs:    []string{"usacloud"},
		DocsURL: sdk.URL("https://docs.usacloud.jp/usacloud"),
		NeedsAuth: needsauth.IfAll(
			needsauth.NotForHelpOrVersion(),
			needsauth.NotWithoutArgs(),
		),
		Uses: []schema.CredentialUsage{
			{
				Name: credname.APIKey,
			},
		},
	}
}
