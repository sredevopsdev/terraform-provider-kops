package schemas

import (
	. "github.com/sredevopsdev/terraform-provider-kops/pkg/schemas"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"k8s.io/kops/pkg/apis/kops"
)

var _ = Schema

func DataSourceUserData() *schema.Resource {
	res := &schema.Resource{
		Schema: map[string]*schema.Schema{
			"name":    ComputedString(),
			"type":    ComputedString(),
			"content": ComputedString(),
		},
	}

	return res
}

func ExpandDataSourceUserData(in map[string]interface{}) kops.UserData {
	if in == nil {
		panic("expand UserData failure, in is nil")
	}
	return kops.UserData{
		Name: func(in interface{}) string {
			return string(ExpandString(in))
		}(in["name"]),
		Type: func(in interface{}) string {
			return string(ExpandString(in))
		}(in["type"]),
		Content: func(in interface{}) string {
			return string(ExpandString(in))
		}(in["content"]),
	}
}

func FlattenDataSourceUserDataInto(in kops.UserData, out map[string]interface{}) {
	out["name"] = func(in string) interface{} {
		return FlattenString(string(in))
	}(in.Name)
	out["type"] = func(in string) interface{} {
		return FlattenString(string(in))
	}(in.Type)
	out["content"] = func(in string) interface{} {
		return FlattenString(string(in))
	}(in.Content)
}

func FlattenDataSourceUserData(in kops.UserData) map[string]interface{} {
	out := map[string]interface{}{}
	FlattenDataSourceUserDataInto(in, out)
	return out
}
