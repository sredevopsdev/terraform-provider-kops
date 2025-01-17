package schemas

import (
	. "github.com/sredevopsdev/terraform-provider-kops/pkg/schemas"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"k8s.io/kops/pkg/apis/kops"
)

var _ = Schema

func DataSourceVolumeMountSpec() *schema.Resource {
	res := &schema.Resource{
		Schema: map[string]*schema.Schema{
			"device":         ComputedString(),
			"filesystem":     ComputedString(),
			"format_options": ComputedList(String()),
			"mount_options":  ComputedList(String()),
			"path":           ComputedString(),
		},
	}

	return res
}

func ExpandDataSourceVolumeMountSpec(in map[string]interface{}) kops.VolumeMountSpec {
	if in == nil {
		panic("expand VolumeMountSpec failure, in is nil")
	}
	return kops.VolumeMountSpec{
		Device: func(in interface{}) string {
			return string(ExpandString(in))
		}(in["device"]),
		Filesystem: func(in interface{}) string {
			return string(ExpandString(in))
		}(in["filesystem"]),
		FormatOptions: func(in interface{}) []string {
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
		}(in["format_options"]),
		MountOptions: func(in interface{}) []string {
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
		}(in["mount_options"]),
		Path: func(in interface{}) string {
			return string(ExpandString(in))
		}(in["path"]),
	}
}

func FlattenDataSourceVolumeMountSpecInto(in kops.VolumeMountSpec, out map[string]interface{}) {
	out["device"] = func(in string) interface{} {
		return FlattenString(string(in))
	}(in.Device)
	out["filesystem"] = func(in string) interface{} {
		return FlattenString(string(in))
	}(in.Filesystem)
	out["format_options"] = func(in []string) interface{} {
		return func(in []string) []interface{} {
			var out []interface{}
			for _, in := range in {
				out = append(out, FlattenString(string(in)))
			}
			return out
		}(in)
	}(in.FormatOptions)
	out["mount_options"] = func(in []string) interface{} {
		return func(in []string) []interface{} {
			var out []interface{}
			for _, in := range in {
				out = append(out, FlattenString(string(in)))
			}
			return out
		}(in)
	}(in.MountOptions)
	out["path"] = func(in string) interface{} {
		return FlattenString(string(in))
	}(in.Path)
}

func FlattenDataSourceVolumeMountSpec(in kops.VolumeMountSpec) map[string]interface{} {
	out := map[string]interface{}{}
	FlattenDataSourceVolumeMountSpecInto(in, out)
	return out
}
