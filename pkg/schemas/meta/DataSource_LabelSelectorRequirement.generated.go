package schemas

import (
	. "github.com/sredevopsdev/terraform-provider-kops/pkg/schemas"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	meta "k8s.io/apimachinery/pkg/apis/meta/v1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

var _ = Schema

func DataSourceLabelSelectorRequirement() *schema.Resource {
	res := &schema.Resource{
		Schema: map[string]*schema.Schema{
			"key":      ComputedString(),
			"operator": ComputedString(),
			"values":   ComputedList(String()),
		},
	}

	return res
}

func ExpandDataSourceLabelSelectorRequirement(in map[string]interface{}) meta.LabelSelectorRequirement {
	if in == nil {
		panic("expand LabelSelectorRequirement failure, in is nil")
	}
	return meta.LabelSelectorRequirement{
		Key: func(in interface{}) string {
			return string(ExpandString(in))
		}(in["key"]),
		Operator: func(in interface{}) meta.LabelSelectorOperator {
			return v1.LabelSelectorOperator(ExpandString(in))
		}(in["operator"]),
		Values: func(in interface{}) []string {
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
		}(in["values"]),
	}
}

func FlattenDataSourceLabelSelectorRequirementInto(in meta.LabelSelectorRequirement, out map[string]interface{}) {
	out["key"] = func(in string) interface{} {
		return FlattenString(string(in))
	}(in.Key)
	out["operator"] = func(in meta.LabelSelectorOperator) interface{} {
		return FlattenString(string(in))
	}(in.Operator)
	out["values"] = func(in []string) interface{} {
		return func(in []string) []interface{} {
			var out []interface{}
			for _, in := range in {
				out = append(out, FlattenString(string(in)))
			}
			return out
		}(in)
	}(in.Values)
}

func FlattenDataSourceLabelSelectorRequirement(in meta.LabelSelectorRequirement) map[string]interface{} {
	out := map[string]interface{}{}
	FlattenDataSourceLabelSelectorRequirementInto(in, out)
	return out
}
