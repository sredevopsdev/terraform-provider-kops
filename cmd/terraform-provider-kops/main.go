package main

import (
	"github.com/sredevopsdev/terraform-provider-kops/pkg/provider"
	"github.com/hashicorp/terraform-plugin-sdk/v2/plugin"
)

func main() {
	plugin.Serve(&plugin.ServeOpts{ProviderFunc: provider.NewProvider})
}
