package schemas

import (
	. "github.com/sredevopsdev/terraform-provider-kops/pkg/schemas"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"k8s.io/kops/pkg/apis/kops"
)

var _ = Schema

func ResourceAuthorizationSpec() *schema.Resource {
	res := &schema.Resource{
		Schema: map[string]*schema.Schema{
			"always_allow": OptionalStruct(ResourceAlwaysAllowAuthorizationSpec()),
			"rbac":         OptionalStruct(ResourceRBACAuthorizationSpec()),
		},
	}

	return res
}

func ExpandResourceAuthorizationSpec(in map[string]interface{}) kops.AuthorizationSpec {
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
						return ExpandResourceAlwaysAllowAuthorizationSpec(in[0].(map[string]interface{}))
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
						return ExpandResourceRBACAuthorizationSpec(in[0].(map[string]interface{}))
					}
					return kops.RBACAuthorizationSpec{}
				}(in))
			}(in)
		}(in["rbac"]),
	}
}

func FlattenResourceAuthorizationSpecInto(in kops.AuthorizationSpec, out map[string]interface{}) {
	out["always_allow"] = func(in *kops.AlwaysAllowAuthorizationSpec) interface{} {
		return func(in *kops.AlwaysAllowAuthorizationSpec) interface{} {
			if in == nil {
				return nil
			}
			return func(in kops.AlwaysAllowAuthorizationSpec) interface{} {
				return func(in kops.AlwaysAllowAuthorizationSpec) []interface{} {
					return []interface{}{FlattenResourceAlwaysAllowAuthorizationSpec(in)}
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
					return []interface{}{FlattenResourceRBACAuthorizationSpec(in)}
				}(in)
			}(*in)
		}(in)
	}(in.RBAC)
}

func FlattenResourceAuthorizationSpec(in kops.AuthorizationSpec) map[string]interface{} {
	out := map[string]interface{}{}
	FlattenResourceAuthorizationSpecInto(in, out)
	return out
}
