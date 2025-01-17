package schemas

import (
	. "github.com/sredevopsdev/terraform-provider-kops/pkg/schemas"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"k8s.io/kops/pkg/apis/kops"
)

var _ = Schema

func DataSourceClassicNetworkingSpec() *schema.Resource {
	res := &schema.Resource{
		Schema: map[string]*schema.Schema{},
	}

	return res
}

func ExpandDataSourceClassicNetworkingSpec(in map[string]interface{}) kops.ClassicNetworkingSpec {
	if in == nil {
		panic("expand ClassicNetworkingSpec failure, in is nil")
	}
	return kops.ClassicNetworkingSpec{}
}

func FlattenDataSourceClassicNetworkingSpecInto(in kops.ClassicNetworkingSpec, out map[string]interface{}) {
}

func FlattenDataSourceClassicNetworkingSpec(in kops.ClassicNetworkingSpec) map[string]interface{} {
	out := map[string]interface{}{}
	FlattenDataSourceClassicNetworkingSpecInto(in, out)
	return out
}
