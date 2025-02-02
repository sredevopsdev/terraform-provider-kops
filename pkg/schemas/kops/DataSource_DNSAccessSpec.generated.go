package schemas

import (
	. "github.com/sredevopsdev/terraform-provider-kops/pkg/schemas"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"k8s.io/kops/pkg/apis/kops"
)

var _ = Schema

func DataSourceDNSAccessSpec() *schema.Resource {
	res := &schema.Resource{
		Schema: map[string]*schema.Schema{},
	}

	return res
}

func ExpandDataSourceDNSAccessSpec(in map[string]interface{}) kops.DNSAccessSpec {
	if in == nil {
		panic("expand DNSAccessSpec failure, in is nil")
	}
	return kops.DNSAccessSpec{}
}

func FlattenDataSourceDNSAccessSpecInto(in kops.DNSAccessSpec, out map[string]interface{}) {
}

func FlattenDataSourceDNSAccessSpec(in kops.DNSAccessSpec) map[string]interface{} {
	out := map[string]interface{}{}
	FlattenDataSourceDNSAccessSpecInto(in, out)
	return out
}
