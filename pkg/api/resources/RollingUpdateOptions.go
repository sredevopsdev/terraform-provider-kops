package resources

import "github.com/sredevopsdev/terraform-provider-kops/pkg/api/utils"

type RollingUpdateOptions struct {
	// Skip allows skipping cluster rolling update
	Skip bool
	utils.RollingUpdateOptions
}
