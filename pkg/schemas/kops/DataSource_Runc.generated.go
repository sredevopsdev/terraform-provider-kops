package schemas

import (
	"reflect"

	. "github.com/sredevopsdev/terraform-provider-kops/pkg/schemas"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"k8s.io/kops/pkg/apis/kops"
)

var _ = Schema

func DataSourceRunc() *schema.Resource {
	res := &schema.Resource{
		Schema: map[string]*schema.Schema{
			"version":  ComputedString(),
			"packages": ComputedStruct(DataSourcePackagesConfig()),
		},
	}

	return res
}

func ExpandDataSourceRunc(in map[string]interface{}) kops.Runc {
	if in == nil {
		panic("expand Runc failure, in is nil")
	}
	return kops.Runc{
		Version: func(in interface{}) *string {
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
		}(in["version"]),
		Packages: func(in interface{}) *kops.PackagesConfig {
			return func(in interface{}) *kops.PackagesConfig {
				if in == nil {
					return nil
				}
				if _, ok := in.([]interface{}); ok && len(in.([]interface{})) == 0 {
					return nil
				}
				return func(in kops.PackagesConfig) *kops.PackagesConfig {
					return &in
				}(func(in interface{}) kops.PackagesConfig {
					if in, ok := in.([]interface{}); ok && len(in) == 1 && in[0] != nil {
						return ExpandDataSourcePackagesConfig(in[0].(map[string]interface{}))
					}
					return kops.PackagesConfig{}
				}(in))
			}(in)
		}(in["packages"]),
	}
}

func FlattenDataSourceRuncInto(in kops.Runc, out map[string]interface{}) {
	out["version"] = func(in *string) interface{} {
		return func(in *string) interface{} {
			if in == nil {
				return nil
			}
			return func(in string) interface{} {
				return FlattenString(string(in))
			}(*in)
		}(in)
	}(in.Version)
	out["packages"] = func(in *kops.PackagesConfig) interface{} {
		return func(in *kops.PackagesConfig) interface{} {
			if in == nil {
				return nil
			}
			return func(in kops.PackagesConfig) interface{} {
				return func(in kops.PackagesConfig) []interface{} {
					return []interface{}{FlattenDataSourcePackagesConfig(in)}
				}(in)
			}(*in)
		}(in)
	}(in.Packages)
}

func FlattenDataSourceRunc(in kops.Runc) map[string]interface{} {
	out := map[string]interface{}{}
	FlattenDataSourceRuncInto(in, out)
	return out
}
