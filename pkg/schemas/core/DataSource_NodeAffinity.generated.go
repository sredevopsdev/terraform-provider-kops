package schemas

import (
	. "github.com/sredevopsdev/terraform-provider-kops/pkg/schemas"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	core "k8s.io/api/core/v1"
)

var _ = Schema

func DataSourceNodeAffinity() *schema.Resource {
	res := &schema.Resource{
		Schema: map[string]*schema.Schema{
			"required_during_scheduling_ignored_during_execution":  ComputedStruct(DataSourceNodeSelector()),
			"preferred_during_scheduling_ignored_during_execution": ComputedList(DataSourcePreferredSchedulingTerm()),
		},
	}

	return res
}

func ExpandDataSourceNodeAffinity(in map[string]interface{}) core.NodeAffinity {
	if in == nil {
		panic("expand NodeAffinity failure, in is nil")
	}
	return core.NodeAffinity{
		RequiredDuringSchedulingIgnoredDuringExecution: func(in interface{}) *core.NodeSelector {
			return func(in interface{}) *core.NodeSelector {
				if in == nil {
					return nil
				}
				if _, ok := in.([]interface{}); ok && len(in.([]interface{})) == 0 {
					return nil
				}
				return func(in core.NodeSelector) *core.NodeSelector {
					return &in
				}(func(in interface{}) core.NodeSelector {
					if in, ok := in.([]interface{}); ok && len(in) == 1 && in[0] != nil {
						return ExpandDataSourceNodeSelector(in[0].(map[string]interface{}))
					}
					return core.NodeSelector{}
				}(in))
			}(in)
		}(in["required_during_scheduling_ignored_during_execution"]),
		PreferredDuringSchedulingIgnoredDuringExecution: func(in interface{}) []core.PreferredSchedulingTerm {
			return func(in interface{}) []core.PreferredSchedulingTerm {
				if in == nil {
					return nil
				}
				var out []core.PreferredSchedulingTerm
				for _, in := range in.([]interface{}) {
					out = append(out, func(in interface{}) core.PreferredSchedulingTerm {
						if in == nil {
							return core.PreferredSchedulingTerm{}
						}
						return ExpandDataSourcePreferredSchedulingTerm(in.(map[string]interface{}))
					}(in))
				}
				return out
			}(in)
		}(in["preferred_during_scheduling_ignored_during_execution"]),
	}
}

func FlattenDataSourceNodeAffinityInto(in core.NodeAffinity, out map[string]interface{}) {
	out["required_during_scheduling_ignored_during_execution"] = func(in *core.NodeSelector) interface{} {
		return func(in *core.NodeSelector) interface{} {
			if in == nil {
				return nil
			}
			return func(in core.NodeSelector) interface{} {
				return func(in core.NodeSelector) []interface{} {
					return []interface{}{FlattenDataSourceNodeSelector(in)}
				}(in)
			}(*in)
		}(in)
	}(in.RequiredDuringSchedulingIgnoredDuringExecution)
	out["preferred_during_scheduling_ignored_during_execution"] = func(in []core.PreferredSchedulingTerm) interface{} {
		return func(in []core.PreferredSchedulingTerm) []interface{} {
			var out []interface{}
			for _, in := range in {
				out = append(out, func(in core.PreferredSchedulingTerm) interface{} {
					return FlattenDataSourcePreferredSchedulingTerm(in)
				}(in))
			}
			return out
		}(in)
	}(in.PreferredDuringSchedulingIgnoredDuringExecution)
}

func FlattenDataSourceNodeAffinity(in core.NodeAffinity) map[string]interface{} {
	out := map[string]interface{}{}
	FlattenDataSourceNodeAffinityInto(in, out)
	return out
}
