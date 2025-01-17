package schemas

import (
	"reflect"

	. "github.com/sredevopsdev/terraform-provider-kops/pkg/schemas"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"k8s.io/apimachinery/pkg/api/resource"
	"k8s.io/kops/pkg/apis/kops"
)

var _ = Schema

func ResourceCalicoNetworkingSpec() *schema.Resource {
	res := &schema.Resource{
		Schema: map[string]*schema.Schema{
			"registry":                  OptionalString(),
			"version":                   OptionalString(),
			"allow_ip_forwarding":       OptionalBool(),
			"aws_src_dst_check":         OptionalString(),
			"bpf_enabled":               OptionalBool(),
			"bpf_external_service_mode": OptionalString(),
			"bpf_kube_proxy_iptables_cleanup_enabled": OptionalBool(),
			"bpf_log_level":                      OptionalString(),
			"chain_insert_mode":                  OptionalString(),
			"cpu_request":                        OptionalQuantity(),
			"cross_subnet":                       OptionalBool(),
			"encapsulation_mode":                 OptionalString(),
			"ip_ip_mode":                         OptionalString(),
			"ipv4_auto_detection_method":         OptionalString(),
			"ipv6_auto_detection_method":         OptionalString(),
			"iptables_backend":                   OptionalString(),
			"log_severity_screen":                OptionalString(),
			"mtu":                                OptionalInt(),
			"prometheus_metrics_enabled":         OptionalBool(),
			"prometheus_metrics_port":            OptionalInt(),
			"prometheus_go_metrics_enabled":      OptionalBool(),
			"prometheus_process_metrics_enabled": OptionalBool(),
			"typha_prometheus_metrics_enabled":   OptionalBool(),
			"typha_prometheus_metrics_port":      OptionalInt(),
			"typha_replicas":                     OptionalInt(),
			"vxlan_mode":                         OptionalString(),
			"wireguard_enabled":                  OptionalBool(),
		},
	}

	return res
}

func ExpandResourceCalicoNetworkingSpec(in map[string]interface{}) kops.CalicoNetworkingSpec {
	if in == nil {
		panic("expand CalicoNetworkingSpec failure, in is nil")
	}
	return kops.CalicoNetworkingSpec{
		Registry: func(in interface{}) string {
			return string(ExpandString(in))
		}(in["registry"]),
		Version: func(in interface{}) string {
			return string(ExpandString(in))
		}(in["version"]),
		AllowIPForwarding: func(in interface{}) bool {
			return bool(ExpandBool(in))
		}(in["allow_ip_forwarding"]),
		AWSSrcDstCheck: func(in interface{}) string {
			return string(ExpandString(in))
		}(in["aws_src_dst_check"]),
		BPFEnabled: func(in interface{}) bool {
			return bool(ExpandBool(in))
		}(in["bpf_enabled"]),
		BPFExternalServiceMode: func(in interface{}) string {
			return string(ExpandString(in))
		}(in["bpf_external_service_mode"]),
		BPFKubeProxyIptablesCleanupEnabled: func(in interface{}) bool {
			return bool(ExpandBool(in))
		}(in["bpf_kube_proxy_iptables_cleanup_enabled"]),
		BPFLogLevel: func(in interface{}) string {
			return string(ExpandString(in))
		}(in["bpf_log_level"]),
		ChainInsertMode: func(in interface{}) string {
			return string(ExpandString(in))
		}(in["chain_insert_mode"]),
		CPURequest: func(in interface{}) *resource.Quantity {
			if in == nil {
				return nil
			}
			if reflect.DeepEqual(in, reflect.Zero(reflect.TypeOf(in)).Interface()) {
				return nil
			}
			return func(in interface{}) *resource.Quantity {
				if in == nil {
					return nil
				}
				if _, ok := in.([]interface{}); ok && len(in.([]interface{})) == 0 {
					return nil
				}
				return func(in resource.Quantity) *resource.Quantity {
					return &in
				}(ExpandQuantity(in))
			}(in)
		}(in["cpu_request"]),
		CrossSubnet: func(in interface{}) *bool {
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
		}(in["cross_subnet"]),
		EncapsulationMode: func(in interface{}) string {
			return string(ExpandString(in))
		}(in["encapsulation_mode"]),
		IPIPMode: func(in interface{}) string {
			return string(ExpandString(in))
		}(in["ip_ip_mode"]),
		IPv4AutoDetectionMethod: func(in interface{}) string {
			return string(ExpandString(in))
		}(in["ipv4_auto_detection_method"]),
		IPv6AutoDetectionMethod: func(in interface{}) string {
			return string(ExpandString(in))
		}(in["ipv6_auto_detection_method"]),
		IptablesBackend: func(in interface{}) string {
			return string(ExpandString(in))
		}(in["iptables_backend"]),
		LogSeverityScreen: func(in interface{}) string {
			return string(ExpandString(in))
		}(in["log_severity_screen"]),
		MTU: func(in interface{}) *int32 {
			if in == nil {
				return nil
			}
			if reflect.DeepEqual(in, reflect.Zero(reflect.TypeOf(in)).Interface()) {
				return nil
			}
			return func(in interface{}) *int32 {
				if in == nil {
					return nil
				}
				if _, ok := in.([]interface{}); ok && len(in.([]interface{})) == 0 {
					return nil
				}
				return func(in int32) *int32 {
					return &in
				}(int32(ExpandInt(in)))
			}(in)
		}(in["mtu"]),
		PrometheusMetricsEnabled: func(in interface{}) bool {
			return bool(ExpandBool(in))
		}(in["prometheus_metrics_enabled"]),
		PrometheusMetricsPort: func(in interface{}) int32 {
			return int32(ExpandInt(in))
		}(in["prometheus_metrics_port"]),
		PrometheusGoMetricsEnabled: func(in interface{}) bool {
			return bool(ExpandBool(in))
		}(in["prometheus_go_metrics_enabled"]),
		PrometheusProcessMetricsEnabled: func(in interface{}) bool {
			return bool(ExpandBool(in))
		}(in["prometheus_process_metrics_enabled"]),
		TyphaPrometheusMetricsEnabled: func(in interface{}) bool {
			return bool(ExpandBool(in))
		}(in["typha_prometheus_metrics_enabled"]),
		TyphaPrometheusMetricsPort: func(in interface{}) int32 {
			return int32(ExpandInt(in))
		}(in["typha_prometheus_metrics_port"]),
		TyphaReplicas: func(in interface{}) int32 {
			return int32(ExpandInt(in))
		}(in["typha_replicas"]),
		VXLANMode: func(in interface{}) string {
			return string(ExpandString(in))
		}(in["vxlan_mode"]),
		WireguardEnabled: func(in interface{}) bool {
			return bool(ExpandBool(in))
		}(in["wireguard_enabled"]),
	}
}

func FlattenResourceCalicoNetworkingSpecInto(in kops.CalicoNetworkingSpec, out map[string]interface{}) {
	out["registry"] = func(in string) interface{} {
		return FlattenString(string(in))
	}(in.Registry)
	out["version"] = func(in string) interface{} {
		return FlattenString(string(in))
	}(in.Version)
	out["allow_ip_forwarding"] = func(in bool) interface{} {
		return FlattenBool(bool(in))
	}(in.AllowIPForwarding)
	out["aws_src_dst_check"] = func(in string) interface{} {
		return FlattenString(string(in))
	}(in.AWSSrcDstCheck)
	out["bpf_enabled"] = func(in bool) interface{} {
		return FlattenBool(bool(in))
	}(in.BPFEnabled)
	out["bpf_external_service_mode"] = func(in string) interface{} {
		return FlattenString(string(in))
	}(in.BPFExternalServiceMode)
	out["bpf_kube_proxy_iptables_cleanup_enabled"] = func(in bool) interface{} {
		return FlattenBool(bool(in))
	}(in.BPFKubeProxyIptablesCleanupEnabled)
	out["bpf_log_level"] = func(in string) interface{} {
		return FlattenString(string(in))
	}(in.BPFLogLevel)
	out["chain_insert_mode"] = func(in string) interface{} {
		return FlattenString(string(in))
	}(in.ChainInsertMode)
	out["cpu_request"] = func(in *resource.Quantity) interface{} {
		return func(in *resource.Quantity) interface{} {
			if in == nil {
				return nil
			}
			return func(in resource.Quantity) interface{} {
				return FlattenQuantity(in)
			}(*in)
		}(in)
	}(in.CPURequest)
	out["cross_subnet"] = func(in *bool) interface{} {
		return func(in *bool) interface{} {
			if in == nil {
				return nil
			}
			return func(in bool) interface{} {
				return FlattenBool(bool(in))
			}(*in)
		}(in)
	}(in.CrossSubnet)
	out["encapsulation_mode"] = func(in string) interface{} {
		return FlattenString(string(in))
	}(in.EncapsulationMode)
	out["ip_ip_mode"] = func(in string) interface{} {
		return FlattenString(string(in))
	}(in.IPIPMode)
	out["ipv4_auto_detection_method"] = func(in string) interface{} {
		return FlattenString(string(in))
	}(in.IPv4AutoDetectionMethod)
	out["ipv6_auto_detection_method"] = func(in string) interface{} {
		return FlattenString(string(in))
	}(in.IPv6AutoDetectionMethod)
	out["iptables_backend"] = func(in string) interface{} {
		return FlattenString(string(in))
	}(in.IptablesBackend)
	out["log_severity_screen"] = func(in string) interface{} {
		return FlattenString(string(in))
	}(in.LogSeverityScreen)
	out["mtu"] = func(in *int32) interface{} {
		return func(in *int32) interface{} {
			if in == nil {
				return nil
			}
			return func(in int32) interface{} {
				return FlattenInt(int(in))
			}(*in)
		}(in)
	}(in.MTU)
	out["prometheus_metrics_enabled"] = func(in bool) interface{} {
		return FlattenBool(bool(in))
	}(in.PrometheusMetricsEnabled)
	out["prometheus_metrics_port"] = func(in int32) interface{} {
		return FlattenInt(int(in))
	}(in.PrometheusMetricsPort)
	out["prometheus_go_metrics_enabled"] = func(in bool) interface{} {
		return FlattenBool(bool(in))
	}(in.PrometheusGoMetricsEnabled)
	out["prometheus_process_metrics_enabled"] = func(in bool) interface{} {
		return FlattenBool(bool(in))
	}(in.PrometheusProcessMetricsEnabled)
	out["typha_prometheus_metrics_enabled"] = func(in bool) interface{} {
		return FlattenBool(bool(in))
	}(in.TyphaPrometheusMetricsEnabled)
	out["typha_prometheus_metrics_port"] = func(in int32) interface{} {
		return FlattenInt(int(in))
	}(in.TyphaPrometheusMetricsPort)
	out["typha_replicas"] = func(in int32) interface{} {
		return FlattenInt(int(in))
	}(in.TyphaReplicas)
	out["vxlan_mode"] = func(in string) interface{} {
		return FlattenString(string(in))
	}(in.VXLANMode)
	out["wireguard_enabled"] = func(in bool) interface{} {
		return FlattenBool(bool(in))
	}(in.WireguardEnabled)
}

func FlattenResourceCalicoNetworkingSpec(in kops.CalicoNetworkingSpec) map[string]interface{} {
	out := map[string]interface{}{}
	FlattenResourceCalicoNetworkingSpecInto(in, out)
	return out
}
