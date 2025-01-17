package schemas

import (
	"reflect"

	. "github.com/sredevopsdev/terraform-provider-kops/pkg/schemas"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"k8s.io/kops/pkg/apis/kops"
)

var _ = Schema

func DataSourceHookSpec() *schema.Resource {
	res := &schema.Resource{
		Schema: map[string]*schema.Schema{
			"name":             ComputedString(),
			"enabled":          ComputedBool(),
			"roles":            ComputedList(String()),
			"requires":         ComputedList(String()),
			"before":           ComputedList(String()),
			"exec_container":   ComputedStruct(DataSourceExecContainerAction()),
			"manifest":         ComputedString(),
			"use_raw_manifest": ComputedBool(),
		},
	}

	return res
}

func ExpandDataSourceHookSpec(in map[string]interface{}) kops.HookSpec {
	if in == nil {
		panic("expand HookSpec failure, in is nil")
	}
	return kops.HookSpec{
		Name: func(in interface{}) string {
			return string(ExpandString(in))
		}(in["name"]),
		Enabled: func(in interface{}) *bool {
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
		}(in["enabled"]),
		Roles: func(in interface{}) []kops.InstanceGroupRole {
			return func(in interface{}) []kops.InstanceGroupRole {
				if in == nil {
					return nil
				}
				var out []kops.InstanceGroupRole
				for _, in := range in.([]interface{}) {
					out = append(out, kops.InstanceGroupRole(ExpandString(in)))
				}
				return out
			}(in)
		}(in["roles"]),
		Requires: func(in interface{}) []string {
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
		}(in["requires"]),
		Before: func(in interface{}) []string {
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
		}(in["before"]),
		ExecContainer: func(in interface{}) *kops.ExecContainerAction {
			return func(in interface{}) *kops.ExecContainerAction {
				if in == nil {
					return nil
				}
				if _, ok := in.([]interface{}); ok && len(in.([]interface{})) == 0 {
					return nil
				}
				return func(in kops.ExecContainerAction) *kops.ExecContainerAction {
					return &in
				}(func(in interface{}) kops.ExecContainerAction {
					if in, ok := in.([]interface{}); ok && len(in) == 1 && in[0] != nil {
						return ExpandDataSourceExecContainerAction(in[0].(map[string]interface{}))
					}
					return kops.ExecContainerAction{}
				}(in))
			}(in)
		}(in["exec_container"]),
		Manifest: func(in interface{}) string {
			return string(ExpandString(in))
		}(in["manifest"]),
		UseRawManifest: func(in interface{}) bool {
			return bool(ExpandBool(in))
		}(in["use_raw_manifest"]),
	}
}

func FlattenDataSourceHookSpecInto(in kops.HookSpec, out map[string]interface{}) {
	out["name"] = func(in string) interface{} {
		return FlattenString(string(in))
	}(in.Name)
	out["enabled"] = func(in *bool) interface{} {
		return func(in *bool) interface{} {
			if in == nil {
				return nil
			}
			return func(in bool) interface{} {
				return FlattenBool(bool(in))
			}(*in)
		}(in)
	}(in.Enabled)
	out["roles"] = func(in []kops.InstanceGroupRole) interface{} {
		return func(in []kops.InstanceGroupRole) []interface{} {
			var out []interface{}
			for _, in := range in {
				out = append(out, FlattenString(string(in)))
			}
			return out
		}(in)
	}(in.Roles)
	out["requires"] = func(in []string) interface{} {
		return func(in []string) []interface{} {
			var out []interface{}
			for _, in := range in {
				out = append(out, FlattenString(string(in)))
			}
			return out
		}(in)
	}(in.Requires)
	out["before"] = func(in []string) interface{} {
		return func(in []string) []interface{} {
			var out []interface{}
			for _, in := range in {
				out = append(out, FlattenString(string(in)))
			}
			return out
		}(in)
	}(in.Before)
	out["exec_container"] = func(in *kops.ExecContainerAction) interface{} {
		return func(in *kops.ExecContainerAction) interface{} {
			if in == nil {
				return nil
			}
			return func(in kops.ExecContainerAction) interface{} {
				return func(in kops.ExecContainerAction) []interface{} {
					return []interface{}{FlattenDataSourceExecContainerAction(in)}
				}(in)
			}(*in)
		}(in)
	}(in.ExecContainer)
	out["manifest"] = func(in string) interface{} {
		return FlattenString(string(in))
	}(in.Manifest)
	out["use_raw_manifest"] = func(in bool) interface{} {
		return FlattenBool(bool(in))
	}(in.UseRawManifest)
}

func FlattenDataSourceHookSpec(in kops.HookSpec) map[string]interface{} {
	out := map[string]interface{}{}
	FlattenDataSourceHookSpecInto(in, out)
	return out
}
