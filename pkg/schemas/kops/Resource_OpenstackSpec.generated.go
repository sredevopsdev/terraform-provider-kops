package schemas

import (
	"reflect"

	. "github.com/sredevopsdev/terraform-provider-kops/pkg/schemas"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"k8s.io/kops/pkg/apis/kops"
)

var _ = Schema

func ResourceOpenstackSpec() *schema.Resource {
	res := &schema.Resource{
		Schema: map[string]*schema.Schema{
			"loadbalancer":         OptionalStruct(ResourceOpenstackLoadbalancerConfig()),
			"monitor":              OptionalStruct(ResourceOpenstackMonitor()),
			"router":               OptionalStruct(ResourceOpenstackRouter()),
			"block_storage":        OptionalStruct(ResourceOpenstackBlockStorageConfig()),
			"insecure_skip_verify": OptionalBool(),
			"network":              OptionalStruct(ResourceOpenstackNetwork()),
			"metadata":             OptionalStruct(ResourceOpenstackMetadata()),
		},
	}

	return res
}

func ExpandResourceOpenstackSpec(in map[string]interface{}) kops.OpenstackSpec {
	if in == nil {
		panic("expand OpenstackSpec failure, in is nil")
	}
	return kops.OpenstackSpec{
		Loadbalancer: func(in interface{}) *kops.OpenstackLoadbalancerConfig {
			return func(in interface{}) *kops.OpenstackLoadbalancerConfig {
				if in == nil {
					return nil
				}
				if _, ok := in.([]interface{}); ok && len(in.([]interface{})) == 0 {
					return nil
				}
				return func(in kops.OpenstackLoadbalancerConfig) *kops.OpenstackLoadbalancerConfig {
					return &in
				}(func(in interface{}) kops.OpenstackLoadbalancerConfig {
					if in, ok := in.([]interface{}); ok && len(in) == 1 && in[0] != nil {
						return ExpandResourceOpenstackLoadbalancerConfig(in[0].(map[string]interface{}))
					}
					return kops.OpenstackLoadbalancerConfig{}
				}(in))
			}(in)
		}(in["loadbalancer"]),
		Monitor: func(in interface{}) *kops.OpenstackMonitor {
			return func(in interface{}) *kops.OpenstackMonitor {
				if in == nil {
					return nil
				}
				if _, ok := in.([]interface{}); ok && len(in.([]interface{})) == 0 {
					return nil
				}
				return func(in kops.OpenstackMonitor) *kops.OpenstackMonitor {
					return &in
				}(func(in interface{}) kops.OpenstackMonitor {
					if in, ok := in.([]interface{}); ok && len(in) == 1 && in[0] != nil {
						return ExpandResourceOpenstackMonitor(in[0].(map[string]interface{}))
					}
					return kops.OpenstackMonitor{}
				}(in))
			}(in)
		}(in["monitor"]),
		Router: func(in interface{}) *kops.OpenstackRouter {
			return func(in interface{}) *kops.OpenstackRouter {
				if in == nil {
					return nil
				}
				if _, ok := in.([]interface{}); ok && len(in.([]interface{})) == 0 {
					return nil
				}
				return func(in kops.OpenstackRouter) *kops.OpenstackRouter {
					return &in
				}(func(in interface{}) kops.OpenstackRouter {
					if in, ok := in.([]interface{}); ok && len(in) == 1 && in[0] != nil {
						return ExpandResourceOpenstackRouter(in[0].(map[string]interface{}))
					}
					return kops.OpenstackRouter{}
				}(in))
			}(in)
		}(in["router"]),
		BlockStorage: func(in interface{}) *kops.OpenstackBlockStorageConfig {
			return func(in interface{}) *kops.OpenstackBlockStorageConfig {
				if in == nil {
					return nil
				}
				if _, ok := in.([]interface{}); ok && len(in.([]interface{})) == 0 {
					return nil
				}
				return func(in kops.OpenstackBlockStorageConfig) *kops.OpenstackBlockStorageConfig {
					return &in
				}(func(in interface{}) kops.OpenstackBlockStorageConfig {
					if in, ok := in.([]interface{}); ok && len(in) == 1 && in[0] != nil {
						return ExpandResourceOpenstackBlockStorageConfig(in[0].(map[string]interface{}))
					}
					return kops.OpenstackBlockStorageConfig{}
				}(in))
			}(in)
		}(in["block_storage"]),
		InsecureSkipVerify: func(in interface{}) *bool {
			if in == nil {
				return nil
			}
			if reflect.DeepEqual(in, reflect.Zero(reflect.TypeOf(in)).Interface()) {
				return nil
			}
			return func(in interface{}) *bool {
				if in == nil {
					return nil
				}
				if _, ok := in.([]interface{}); ok && len(in.([]interface{})) == 0 {
					return nil
				}
				return func(in bool) *bool {
					return &in
				}(bool(ExpandBool(in)))
			}(in)
		}(in["insecure_skip_verify"]),
		Network: func(in interface{}) *kops.OpenstackNetwork {
			return func(in interface{}) *kops.OpenstackNetwork {
				if in == nil {
					return nil
				}
				if _, ok := in.([]interface{}); ok && len(in.([]interface{})) == 0 {
					return nil
				}
				return func(in kops.OpenstackNetwork) *kops.OpenstackNetwork {
					return &in
				}(func(in interface{}) kops.OpenstackNetwork {
					if in, ok := in.([]interface{}); ok && len(in) == 1 && in[0] != nil {
						return ExpandResourceOpenstackNetwork(in[0].(map[string]interface{}))
					}
					return kops.OpenstackNetwork{}
				}(in))
			}(in)
		}(in["network"]),
		Metadata: func(in interface{}) *kops.OpenstackMetadata {
			return func(in interface{}) *kops.OpenstackMetadata {
				if in == nil {
					return nil
				}
				if _, ok := in.([]interface{}); ok && len(in.([]interface{})) == 0 {
					return nil
				}
				return func(in kops.OpenstackMetadata) *kops.OpenstackMetadata {
					return &in
				}(func(in interface{}) kops.OpenstackMetadata {
					if in, ok := in.([]interface{}); ok && len(in) == 1 && in[0] != nil {
						return ExpandResourceOpenstackMetadata(in[0].(map[string]interface{}))
					}
					return kops.OpenstackMetadata{}
				}(in))
			}(in)
		}(in["metadata"]),
	}
}

func FlattenResourceOpenstackSpecInto(in kops.OpenstackSpec, out map[string]interface{}) {
	out["loadbalancer"] = func(in *kops.OpenstackLoadbalancerConfig) interface{} {
		return func(in *kops.OpenstackLoadbalancerConfig) interface{} {
			if in == nil {
				return nil
			}
			return func(in kops.OpenstackLoadbalancerConfig) interface{} {
				return func(in kops.OpenstackLoadbalancerConfig) []interface{} {
					return []interface{}{FlattenResourceOpenstackLoadbalancerConfig(in)}
				}(in)
			}(*in)
		}(in)
	}(in.Loadbalancer)
	out["monitor"] = func(in *kops.OpenstackMonitor) interface{} {
		return func(in *kops.OpenstackMonitor) interface{} {
			if in == nil {
				return nil
			}
			return func(in kops.OpenstackMonitor) interface{} {
				return func(in kops.OpenstackMonitor) []interface{} {
					return []interface{}{FlattenResourceOpenstackMonitor(in)}
				}(in)
			}(*in)
		}(in)
	}(in.Monitor)
	out["router"] = func(in *kops.OpenstackRouter) interface{} {
		return func(in *kops.OpenstackRouter) interface{} {
			if in == nil {
				return nil
			}
			return func(in kops.OpenstackRouter) interface{} {
				return func(in kops.OpenstackRouter) []interface{} {
					return []interface{}{FlattenResourceOpenstackRouter(in)}
				}(in)
			}(*in)
		}(in)
	}(in.Router)
	out["block_storage"] = func(in *kops.OpenstackBlockStorageConfig) interface{} {
		return func(in *kops.OpenstackBlockStorageConfig) interface{} {
			if in == nil {
				return nil
			}
			return func(in kops.OpenstackBlockStorageConfig) interface{} {
				return func(in kops.OpenstackBlockStorageConfig) []interface{} {
					return []interface{}{FlattenResourceOpenstackBlockStorageConfig(in)}
				}(in)
			}(*in)
		}(in)
	}(in.BlockStorage)
	out["insecure_skip_verify"] = func(in *bool) interface{} {
		return func(in *bool) interface{} {
			if in == nil {
				return nil
			}
			return func(in bool) interface{} {
				return FlattenBool(bool(in))
			}(*in)
		}(in)
	}(in.InsecureSkipVerify)
	out["network"] = func(in *kops.OpenstackNetwork) interface{} {
		return func(in *kops.OpenstackNetwork) interface{} {
			if in == nil {
				return nil
			}
			return func(in kops.OpenstackNetwork) interface{} {
				return func(in kops.OpenstackNetwork) []interface{} {
					return []interface{}{FlattenResourceOpenstackNetwork(in)}
				}(in)
			}(*in)
		}(in)
	}(in.Network)
	out["metadata"] = func(in *kops.OpenstackMetadata) interface{} {
		return func(in *kops.OpenstackMetadata) interface{} {
			if in == nil {
				return nil
			}
			return func(in kops.OpenstackMetadata) interface{} {
				return func(in kops.OpenstackMetadata) []interface{} {
					return []interface{}{FlattenResourceOpenstackMetadata(in)}
				}(in)
			}(*in)
		}(in)
	}(in.Metadata)
}

func FlattenResourceOpenstackSpec(in kops.OpenstackSpec) map[string]interface{} {
	out := map[string]interface{}{}
	FlattenResourceOpenstackSpecInto(in, out)
	return out
}
