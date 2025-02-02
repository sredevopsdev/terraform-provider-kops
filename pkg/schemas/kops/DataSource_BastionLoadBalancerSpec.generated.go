package schemas

import (
	. "github.com/sredevopsdev/terraform-provider-kops/pkg/schemas"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"k8s.io/kops/pkg/apis/kops"
)

var _ = Schema

func DataSourceBastionLoadBalancerSpec() *schema.Resource {
	res := &schema.Resource{
		Schema: map[string]*schema.Schema{
			"additional_security_groups": ComputedList(String()),
			"type":                       ComputedString(),
		},
	}

	return res
}

func ExpandDataSourceBastionLoadBalancerSpec(in map[string]interface{}) kops.BastionLoadBalancerSpec {
	if in == nil {
		panic("expand BastionLoadBalancerSpec failure, in is nil")
	}
	return kops.BastionLoadBalancerSpec{
		AdditionalSecurityGroups: func(in interface{}) []string {
			return func(in interface{}) []string {
				if in == nil {
					return nil
				}
				var out []string
				for _, in := range in.([]interface{}) {
					out = append(out, string(ExpandString(in)))
				}
				return out
			}(in)
		}(in["additional_security_groups"]),
		Type: func(in interface{}) kops.LoadBalancerType {
			return kops.LoadBalancerType(ExpandString(in))
		}(in["type"]),
	}
}

func FlattenDataSourceBastionLoadBalancerSpecInto(in kops.BastionLoadBalancerSpec, out map[string]interface{}) {
	out["additional_security_groups"] = func(in []string) interface{} {
		return func(in []string) []interface{} {
			var out []interface{}
			for _, in := range in {
				out = append(out, FlattenString(string(in)))
			}
			return out
		}(in)
	}(in.AdditionalSecurityGroups)
	out["type"] = func(in kops.LoadBalancerType) interface{} {
		return FlattenString(string(in))
	}(in.Type)
}

func FlattenDataSourceBastionLoadBalancerSpec(in kops.BastionLoadBalancerSpec) map[string]interface{} {
	out := map[string]interface{}{}
	FlattenDataSourceBastionLoadBalancerSpecInto(in, out)
	return out
}
