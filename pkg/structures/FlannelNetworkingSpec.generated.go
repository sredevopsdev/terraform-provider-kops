package structures

import (
	"k8s.io/kops/pkg/apis/kops"
)

func ExpandFlannelNetworkingSpec(in map[string]interface{}) kops.FlannelNetworkingSpec {
	if in == nil {
		panic("expand FlannelNetworkingSpec failure, in is nil")
	}
	return kops.FlannelNetworkingSpec{
		Backend: func(in interface{}) string {
			value := string(ExpandString(in))
			return value
		}(in["backend"]),
		DisableTxChecksumOffloading: func(in interface{}) bool {
			value := bool(ExpandBool(in))
			return value
		}(in["disable_tx_checksum_offloading"]),
		IptablesResyncSeconds: func(in interface{}) *int32 {
			value := func(in interface{}) *int32 {
				if in == nil {
					return nil
				}
				if slice, ok := in.([]interface{}); ok && len(slice) == 0 {
					return nil
				}
				tmp := func(in int32) *int32 {
					// if reflect.DeepEqual(in, reflect.Zero(reflect.TypeOf(in)).Interface()) {
					// 	return nil
					// }
					return &in
				}(int32(ExpandInt(in)))
				return tmp
			}(in)
			return value
		}(in["iptables_resync_seconds"]),
	}
}

func FlattenFlannelNetworkingSpec(in kops.FlannelNetworkingSpec) map[string]interface{} {
	return map[string]interface{}{
		"backend": func(in string) interface{} {
			value := FlattenString(string(in))
			return value
		}(in.Backend),
		"disable_tx_checksum_offloading": func(in bool) interface{} {
			value := FlattenBool(bool(in))
			return value
		}(in.DisableTxChecksumOffloading),
		"iptables_resync_seconds": func(in *int32) interface{} {
			value := func(in *int32) interface{} {
				if in == nil {
					return nil
				}
				return func(in int32) interface{} {
					return FlattenInt(int(in))
				}(*in)
			}(in)
			return value
		}(in.IptablesResyncSeconds),
	}
}
