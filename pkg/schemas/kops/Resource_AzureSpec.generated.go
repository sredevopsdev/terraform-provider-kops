package schemas

import (
	. "github.com/sredevopsdev/terraform-provider-kops/pkg/schemas"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"k8s.io/kops/pkg/apis/kops"
)

var _ = Schema

func ResourceAzureSpec() *schema.Resource {
	res := &schema.Resource{
		Schema: map[string]*schema.Schema{
			"subscription_id":     OptionalString(),
			"tenant_id":           OptionalString(),
			"resource_group_name": OptionalString(),
			"route_table_name":    OptionalString(),
			"admin_user":          OptionalString(),
		},
	}

	return res
}

func ExpandResourceAzureSpec(in map[string]interface{}) kops.AzureSpec {
	if in == nil {
		panic("expand AzureSpec failure, in is nil")
	}
	return kops.AzureSpec{
		SubscriptionID: func(in interface{}) string {
			return string(ExpandString(in))
		}(in["subscription_id"]),
		TenantID: func(in interface{}) string {
			return string(ExpandString(in))
		}(in["tenant_id"]),
		ResourceGroupName: func(in interface{}) string {
			return string(ExpandString(in))
		}(in["resource_group_name"]),
		RouteTableName: func(in interface{}) string {
			return string(ExpandString(in))
		}(in["route_table_name"]),
		AdminUser: func(in interface{}) string {
			return string(ExpandString(in))
		}(in["admin_user"]),
	}
}

func FlattenResourceAzureSpecInto(in kops.AzureSpec, out map[string]interface{}) {
	out["subscription_id"] = func(in string) interface{} {
		return FlattenString(string(in))
	}(in.SubscriptionID)
	out["tenant_id"] = func(in string) interface{} {
		return FlattenString(string(in))
	}(in.TenantID)
	out["resource_group_name"] = func(in string) interface{} {
		return FlattenString(string(in))
	}(in.ResourceGroupName)
	out["route_table_name"] = func(in string) interface{} {
		return FlattenString(string(in))
	}(in.RouteTableName)
	out["admin_user"] = func(in string) interface{} {
		return FlattenString(string(in))
	}(in.AdminUser)
}

func FlattenResourceAzureSpec(in kops.AzureSpec) map[string]interface{} {
	out := map[string]interface{}{}
	FlattenResourceAzureSpecInto(in, out)
	return out
}
