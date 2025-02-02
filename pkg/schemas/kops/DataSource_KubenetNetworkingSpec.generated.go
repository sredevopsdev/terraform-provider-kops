package schemas

import (
	. "github.com/sredevopsdev/terraform-provider-kops/pkg/schemas"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"k8s.io/kops/pkg/apis/kops"
)

var _ = Schema

func DataSourceKubenetNetworkingSpec() *schema.Resource {
	res := &schema.Resource{
		Schema: map[string]*schema.Schema{},
	}

	return res
}

func ExpandDataSourceKubenetNetworkingSpec(in map[string]interface{}) kops.KubenetNetworkingSpec {
	if in == nil {
		panic("expand KubenetNetworkingSpec failure, in is nil")
	}
	return kops.KubenetNetworkingSpec{}
}

func FlattenDataSourceKubenetNetworkingSpecInto(in kops.KubenetNetworkingSpec, out map[string]interface{}) {
}

func FlattenDataSourceKubenetNetworkingSpec(in kops.KubenetNetworkingSpec) map[string]interface{} {
	out := map[string]interface{}{}
	FlattenDataSourceKubenetNetworkingSpecInto(in, out)
	return out
}
