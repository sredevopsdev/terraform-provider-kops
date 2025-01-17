package schemas

import (
	. "github.com/sredevopsdev/terraform-provider-kops/pkg/schemas"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"k8s.io/kops/pkg/apis/kops"
)

var _ = Schema

func DataSourceAuthorizationSpec() *schema.Resource {
	res := &schema.Resource{
		Schema: map[string]*schema.Schema{
			"always_allow": ComputedStruct(DataSourceAlwaysAllowAuthorizationSpec()),
			"rbac":         ComputedStruct(DataSourceRBACAuthorizationSpec()),
		},
	}

	return res
}

func ExpandDataSourceAuthorizationSpec(in map[string]interface{}) kops.AuthorizationSpec {
	if in == nil {
		panic("expand AuthorizationSpec failure, in is nil")
	}
	return kops.AuthorizationSpec{
		AlwaysAllow: func(in interface{}) *kops.AlwaysAllowAuthorizationSpec {
			return func(in interface{}) *kops.AlwaysAllowAuthorizationSpec {
				if in == nil {
					return nil
				}
				if _, ok := in.([]interface{}); ok && len(in.([]interface{})) == 0 {
					return nil
				}
				return func(in kops.AlwaysAllowAuthorizationSpec) *kops.AlwaysAllowAuthorizationSpec {
					return &in
				}(func(in interface{}) kops.AlwaysAllowAuthorizationSpec {
					if in, ok := in.([]interface{}); ok && len(in) == 1 && in[0] != nil {
						return ExpandDataSourceAlwaysAllowAuthorizationSpec(in[0].(map[string]interface{}))
					}
					return kops.AlwaysAllowAuthorizationSpec{}
				}(in))
			}(in)
		}(in["always_allow"]),
		RBAC: func(in interface{}) *kops.RBACAuthorizationSpec {
			return func(in interface{}) *kops.RBACAuthorizationSpec {
				if in == nil {
					return nil
				}
				if _, ok := in.([]interface{}); ok && len(in.([]interface{})) == 0 {
					return nil
				}
				return func(in kops.RBACAuthorizationSpec) *kops.RBACAuthorizationSpec {
					return &in
				}(func(in interface{}) kops.RBACAuthorizationSpec {
					if in, ok := in.([]interface{}); ok && len(in) == 1 && in[0] != nil {
						return ExpandDataSourceRBACAuthorizationSpec(in[0].(map[string]interface{}))
					}
					return kops.RBACAuthorizationSpec{}
				}(in))
			}(in)
		}(in["rbac"]),
	}
}

func FlattenDataSourceAuthorizationSpecInto(in kops.AuthorizationSpec, out map[string]interface{}) {
	out["always_allow"] = func(in *kops.AlwaysAllowAuthorizationSpec) interface{} {
		return func(in *kops.AlwaysAllowAuthorizationSpec) interface{} {
			if in == nil {
				return nil
			}
			return func(in kops.AlwaysAllowAuthorizationSpec) interface{} {
				return func(in kops.AlwaysAllowAuthorizationSpec) []interface{} {
					return []interface{}{FlattenDataSourceAlwaysAllowAuthorizationSpec(in)}
				}(in)
			}(*in)
		}(in)
	}(in.AlwaysAllow)
	out["rbac"] = func(in *kops.RBACAuthorizationSpec) interface{} {
		return func(in *kops.RBACAuthorizationSpec) interface{} {
			if in == nil {
				return nil
			}
			return func(in kops.RBACAuthorizationSpec) interface{} {
				return func(in kops.RBACAuthorizationSpec) []interface{} {
					return []interface{}{FlattenDataSourceRBACAuthorizationSpec(in)}
				}(in)
			}(*in)
		}(in)
	}(in.RBAC)
}

func FlattenDataSourceAuthorizationSpec(in kops.AuthorizationSpec) map[string]interface{} {
	out := map[string]interface{}{}
	FlattenDataSourceAuthorizationSpecInto(in, out)
	return out
}
