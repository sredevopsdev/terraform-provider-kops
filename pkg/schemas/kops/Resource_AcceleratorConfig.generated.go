package schemas

import (
	. "github.com/sredevopsdev/terraform-provider-kops/pkg/schemas"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"k8s.io/kops/pkg/apis/kops"
)

var _ = Schema

func ResourceAcceleratorConfig() *schema.Resource {
	res := &schema.Resource{
		Schema: map[string]*schema.Schema{
			"accelerator_count": OptionalInt(),
			"accelerator_type":  OptionalString(),
		},
	}

	return res
}

func ExpandResourceAcceleratorConfig(in map[string]interface{}) kops.AcceleratorConfig {
	if in == nil {
		panic("expand AcceleratorConfig failure, in is nil")
	}
	return kops.AcceleratorConfig{
		AcceleratorCount: func(in interface{}) int64 {
			return int64(ExpandInt(in))
		}(in["accelerator_count"]),
		AcceleratorType: func(in interface{}) string {
			return string(ExpandString(in))
		}(in["accelerator_type"]),
	}
}

func FlattenResourceAcceleratorConfigInto(in kops.AcceleratorConfig, out map[string]interface{}) {
	out["accelerator_count"] = func(in int64) interface{} {
		return FlattenInt(int(in))
	}(in.AcceleratorCount)
	out["accelerator_type"] = func(in string) interface{} {
		return FlattenString(string(in))
	}(in.AcceleratorType)
}

func FlattenResourceAcceleratorConfig(in kops.AcceleratorConfig) map[string]interface{} {
	out := map[string]interface{}{}
	FlattenResourceAcceleratorConfigInto(in, out)
	return out
}
