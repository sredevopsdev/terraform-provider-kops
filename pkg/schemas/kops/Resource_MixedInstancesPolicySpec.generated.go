package schemas

import (
	"reflect"

	. "github.com/sredevopsdev/terraform-provider-kops/pkg/schemas"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"k8s.io/kops/pkg/apis/kops"
)

var _ = Schema

func ResourceMixedInstancesPolicySpec() *schema.Resource {
	res := &schema.Resource{
		Schema: map[string]*schema.Schema{
			"instances":                     OptionalList(String()),
			"instance_requirements":         OptionalStruct(ResourceInstanceRequirementsSpec()),
			"on_demand_allocation_strategy": OptionalString(),
			"on_demand_base":                Nullable(OptionalInt()),
			"on_demand_above_base":          Nullable(OptionalInt()),
			"spot_allocation_strategy":      OptionalString(),
			"spot_instance_pools":           OptionalInt(),
		},
	}

	return res
}

func ExpandResourceMixedInstancesPolicySpec(in map[string]interface{}) kops.MixedInstancesPolicySpec {
	if in == nil {
		panic("expand MixedInstancesPolicySpec failure, in is nil")
	}
	return kops.MixedInstancesPolicySpec{
		Instances: func(in interface{}) []string {
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
		}(in["instances"]),
		InstanceRequirements: func(in interface{}) *kops.InstanceRequirementsSpec {
			return func(in interface{}) *kops.InstanceRequirementsSpec {
				if in == nil {
					return nil
				}
				if _, ok := in.([]interface{}); ok && len(in.([]interface{})) == 0 {
					return nil
				}
				return func(in kops.InstanceRequirementsSpec) *kops.InstanceRequirementsSpec {
					return &in
				}(func(in interface{}) kops.InstanceRequirementsSpec {
					if in, ok := in.([]interface{}); ok && len(in) == 1 && in[0] != nil {
						return ExpandResourceInstanceRequirementsSpec(in[0].(map[string]interface{}))
					}
					return kops.InstanceRequirementsSpec{}
				}(in))
			}(in)
		}(in["instance_requirements"]),
		OnDemandAllocationStrategy: func(in interface{}) *string {
			if in == nil {
				return nil
			}
			if reflect.DeepEqual(in, reflect.Zero(reflect.TypeOf(in)).Interface()) {
				return nil
			}
			return func(in interface{}) *string {
				if in == nil {
					return nil
				}
				if _, ok := in.([]interface{}); ok && len(in.([]interface{})) == 0 {
					return nil
				}
				return func(in string) *string {
					return &in
				}(string(ExpandString(in)))
			}(in)
		}(in["on_demand_allocation_strategy"]),
		OnDemandBase: func(in interface{}) *int64 {
			if in == nil {
				return nil
			}
			if in, ok := in.([]interface{}); ok && len(in) == 1 {
				return func(in interface{}) *int64 {
					return func(in interface{}) *int64 {
						if in == nil {
							return nil
						}
						if _, ok := in.([]interface{}); ok && len(in.([]interface{})) == 0 {
							return nil
						}
						return func(in int64) *int64 {
							return &in
						}(int64(ExpandInt(in)))
					}(in)
				}(in[0].(map[string]interface{})["value"])
			}
			return nil
		}(in["on_demand_base"]),
		OnDemandAboveBase: func(in interface{}) *int64 {
			if in == nil {
				return nil
			}
			if in, ok := in.([]interface{}); ok && len(in) == 1 {
				return func(in interface{}) *int64 {
					return func(in interface{}) *int64 {
						if in == nil {
							return nil
						}
						if _, ok := in.([]interface{}); ok && len(in.([]interface{})) == 0 {
							return nil
						}
						return func(in int64) *int64 {
							return &in
						}(int64(ExpandInt(in)))
					}(in)
				}(in[0].(map[string]interface{})["value"])
			}
			return nil
		}(in["on_demand_above_base"]),
		SpotAllocationStrategy: func(in interface{}) *string {
			if in == nil {
				return nil
			}
			if reflect.DeepEqual(in, reflect.Zero(reflect.TypeOf(in)).Interface()) {
				return nil
			}
			return func(in interface{}) *string {
				if in == nil {
					return nil
				}
				if _, ok := in.([]interface{}); ok && len(in.([]interface{})) == 0 {
					return nil
				}
				return func(in string) *string {
					return &in
				}(string(ExpandString(in)))
			}(in)
		}(in["spot_allocation_strategy"]),
		SpotInstancePools: func(in interface{}) *int64 {
			if in == nil {
				return nil
			}
			if reflect.DeepEqual(in, reflect.Zero(reflect.TypeOf(in)).Interface()) {
				return nil
			}
			return func(in interface{}) *int64 {
				if in == nil {
					return nil
				}
				if _, ok := in.([]interface{}); ok && len(in.([]interface{})) == 0 {
					return nil
				}
				return func(in int64) *int64 {
					return &in
				}(int64(ExpandInt(in)))
			}(in)
		}(in["spot_instance_pools"]),
	}
}

func FlattenResourceMixedInstancesPolicySpecInto(in kops.MixedInstancesPolicySpec, out map[string]interface{}) {
	out["instances"] = func(in []string) interface{} {
		return func(in []string) []interface{} {
			var out []interface{}
			for _, in := range in {
				out = append(out, FlattenString(string(in)))
			}
			return out
		}(in)
	}(in.Instances)
	out["instance_requirements"] = func(in *kops.InstanceRequirementsSpec) interface{} {
		return func(in *kops.InstanceRequirementsSpec) interface{} {
			if in == nil {
				return nil
			}
			return func(in kops.InstanceRequirementsSpec) interface{} {
				return func(in kops.InstanceRequirementsSpec) []interface{} {
					return []interface{}{FlattenResourceInstanceRequirementsSpec(in)}
				}(in)
			}(*in)
		}(in)
	}(in.InstanceRequirements)
	out["on_demand_allocation_strategy"] = func(in *string) interface{} {
		return func(in *string) interface{} {
			if in == nil {
				return nil
			}
			return func(in string) interface{} {
				return FlattenString(string(in))
			}(*in)
		}(in)
	}(in.OnDemandAllocationStrategy)
	out["on_demand_base"] = func(in *int64) interface{} {
		if in == nil {
			return nil
		}
		return []interface{}{map[string]interface{}{"value": func(in *int64) interface{} {
			if in == nil {
				return nil
			}
			return func(in int64) interface{} {
				return FlattenInt(int(in))
			}(*in)
		}(in)}}
	}(in.OnDemandBase)
	out["on_demand_above_base"] = func(in *int64) interface{} {
		if in == nil {
			return nil
		}
		return []interface{}{map[string]interface{}{"value": func(in *int64) interface{} {
			if in == nil {
				return nil
			}
			return func(in int64) interface{} {
				return FlattenInt(int(in))
			}(*in)
		}(in)}}
	}(in.OnDemandAboveBase)
	out["spot_allocation_strategy"] = func(in *string) interface{} {
		return func(in *string) interface{} {
			if in == nil {
				return nil
			}
			return func(in string) interface{} {
				return FlattenString(string(in))
			}(*in)
		}(in)
	}(in.SpotAllocationStrategy)
	out["spot_instance_pools"] = func(in *int64) interface{} {
		return func(in *int64) interface{} {
			if in == nil {
				return nil
			}
			return func(in int64) interface{} {
				return FlattenInt(int(in))
			}(*in)
		}(in)
	}(in.SpotInstancePools)
}

func FlattenResourceMixedInstancesPolicySpec(in kops.MixedInstancesPolicySpec) map[string]interface{} {
	out := map[string]interface{}{}
	FlattenResourceMixedInstancesPolicySpecInto(in, out)
	return out
}
