package schemas

import (
	. "github.com/sredevopsdev/terraform-provider-kops/pkg/schemas"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"k8s.io/kops/pkg/apis/kops"
)

var _ = Schema

func ResourceGCENetworkingSpec() *schema.Resource {
	res := &schema.Resource{
		Schema: map[string]*schema.Schema{},
	}

	return res
}

func ExpandResourceGCENetworkingSpec(in map[string]interface{}) kops.GCENetworkingSpec {
	if in == nil {
		panic("expand GCENetworkingSpec failure, in is nil")
	}
	return kops.GCENetworkingSpec{}
}

func FlattenResourceGCENetworkingSpecInto(in kops.GCENetworkingSpec, out map[string]interface{}) {
}

func FlattenResourceGCENetworkingSpec(in kops.GCENetworkingSpec) map[string]interface{} {
	out := map[string]interface{}{}
	FlattenResourceGCENetworkingSpecInto(in, out)
	return out
}
