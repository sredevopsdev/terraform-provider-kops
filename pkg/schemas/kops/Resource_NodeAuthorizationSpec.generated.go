package schemas

import (
	. "github.com/sredevopsdev/terraform-provider-kops/pkg/schemas"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"k8s.io/kops/pkg/apis/kops"
)

var _ = Schema

func ResourceNodeAuthorizationSpec() *schema.Resource {
	res := &schema.Resource{
		Schema: map[string]*schema.Schema{
			"node_authorizer": OptionalStruct(ResourceNodeAuthorizerSpec()),
		},
	}

	return res
}

func ExpandResourceNodeAuthorizationSpec(in map[string]interface{}) kops.NodeAuthorizationSpec {
	if in == nil {
		panic("expand NodeAuthorizationSpec failure, in is nil")
	}
	return kops.NodeAuthorizationSpec{
		NodeAuthorizer: func(in interface{}) *kops.NodeAuthorizerSpec {
			return func(in interface{}) *kops.NodeAuthorizerSpec {
				if in == nil {
					return nil
				}
				if _, ok := in.([]interface{}); ok && len(in.([]interface{})) == 0 {
					return nil
				}
				return func(in kops.NodeAuthorizerSpec) *kops.NodeAuthorizerSpec {
					return &in
				}(func(in interface{}) kops.NodeAuthorizerSpec {
					if in, ok := in.([]interface{}); ok && len(in) == 1 && in[0] != nil {
						return ExpandResourceNodeAuthorizerSpec(in[0].(map[string]interface{}))
					}
					return kops.NodeAuthorizerSpec{}
				}(in))
			}(in)
		}(in["node_authorizer"]),
	}
}

func FlattenResourceNodeAuthorizationSpecInto(in kops.NodeAuthorizationSpec, out map[string]interface{}) {
	out["node_authorizer"] = func(in *kops.NodeAuthorizerSpec) interface{} {
		return func(in *kops.NodeAuthorizerSpec) interface{} {
			if in == nil {
				return nil
			}
			return func(in kops.NodeAuthorizerSpec) interface{} {
				return func(in kops.NodeAuthorizerSpec) []interface{} {
					return []interface{}{FlattenResourceNodeAuthorizerSpec(in)}
				}(in)
			}(*in)
		}(in)
	}(in.NodeAuthorizer)
}

func FlattenResourceNodeAuthorizationSpec(in kops.NodeAuthorizationSpec) map[string]interface{} {
	out := map[string]interface{}{}
	FlattenResourceNodeAuthorizationSpecInto(in, out)
	return out
}
