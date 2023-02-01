package sakuracloud

import (
	"github.com/1Password/shell-plugins/sdk"
	"github.com/1Password/shell-plugins/sdk/schema"
)

func New() schema.Plugin {
	return schema.Plugin{
		Name: "sakuracloud",
		Platform: schema.PlatformInfo{
			Name:     "SAKURA Cloud",
			Homepage: sdk.URL("https://cloud.sakura.ad.jp"),
		},
		Credentials: []schema.CredentialType{
			APIKey(),
		},
		Executables: []schema.Executable{
			SakuraCloudCLI(),
		},
	}
}
