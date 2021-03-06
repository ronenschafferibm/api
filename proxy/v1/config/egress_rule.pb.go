// Code generated by protoc-gen-go.
// source: proxy/v1/config/egress_rule.proto
// DO NOT EDIT!

package istio_proxy_v1_config

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// Egress rules describe the properties of a service outside Istio. When transparent proxying
// is used, egress rules signify a white listed set of domains that microserves in the mesh
// are allowed to access. A subset of routing rules and all destination policies can be applied
// on the service targeted by an egress rule. The destination of an egress rule is allowed to
// contain wildcards (e.g., *.foo.com). Currently, only HTTP-based services can be expressed
// through the egress rule. If TLS origination from the sidecar is desired, the protocol
// associated with the service port must be marked as HTTPS, and the service is expected to
// be accessed over HTTP (e.g., http://gmail.com:443). The sidecar will automatically upgrade
// the connection to TLS when initiating a connection with the external service.
//
// For example, the following egress rule describes the set of services hosted under the *.foo.com domain
//
//     kind: EgressRule
//     metadata:
//       name: foo-egress-rule
//     spec:
//       destination:
//         service: *.foo.com
//       ports:
//         - port: 80
//           protocol: http
//         - port: 443
//           protocol: https
//
type EgressRule struct {
	// REQUIRED: Hostname or a wildcard domain name associated with the external service.
	// ONLY the "service" field of destination will be taken into consideration. Name,
	// namespace, domain and labels are ignored. Routing rules and destination policies that
	// refer to these external services must have identical specification for the destination
	// as the corresponding egress rule. Wildcard domain specifications must conform to format
	// allowed by Envoy's Virtual Host specification, such as “*.foo.com” or “*-bar.foo.com”.
	// The character '*' in a domain specification indicates a non-empty string. Hence, a wildcard
	// domain of form “*-bar.foo.com” will match “baz-bar.foo.com” but not “-bar.foo.com”.
	Destination *IstioService `protobuf:"bytes,1,opt,name=destination" json:"destination,omitempty"`
	// REQUIRED: list of ports on which the external service is available.
	Ports []*Port `protobuf:"bytes,2,rep,name=ports" json:"ports,omitempty"`
	// Forward all the external traffic through a dedicated egress proxy. It is used in some scenarios
	// where there is a requirement that all the external traffic goes through special dedicated nodes/pods.
	// These dedicated egress nodes could then be more closely monitored for security vulnerabilities.
	//
	// The default is false, i.e. the sidecar forwards external traffic directly to the external service.
	UseEgressProxy bool `protobuf:"varint,3,opt,name=use_egress_proxy,json=useEgressProxy" json:"use_egress_proxy,omitempty"`
}

func (m *EgressRule) Reset()                    { *m = EgressRule{} }
func (m *EgressRule) String() string            { return proto.CompactTextString(m) }
func (*EgressRule) ProtoMessage()               {}
func (*EgressRule) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{0} }

func (m *EgressRule) GetDestination() *IstioService {
	if m != nil {
		return m.Destination
	}
	return nil
}

func (m *EgressRule) GetPorts() []*Port {
	if m != nil {
		return m.Ports
	}
	return nil
}

func (m *EgressRule) GetUseEgressProxy() bool {
	if m != nil {
		return m.UseEgressProxy
	}
	return false
}

func init() {
	proto.RegisterType((*EgressRule)(nil), "istio.proxy.v1.config.EgressRule")
}

func init() { proto.RegisterFile("proxy/v1/config/egress_rule.proto", fileDescriptor1) }

var fileDescriptor1 = []byte{
	// 199 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x52, 0x2c, 0x28, 0xca, 0xaf,
	0xa8, 0xd4, 0x2f, 0x33, 0xd4, 0x4f, 0xce, 0xcf, 0x4b, 0xcb, 0x4c, 0xd7, 0x4f, 0x4d, 0x2f, 0x4a,
	0x2d, 0x2e, 0x8e, 0x2f, 0x2a, 0xcd, 0x49, 0xd5, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x12, 0xcd,
	0x2c, 0x2e, 0xc9, 0xcc, 0xd7, 0x03, 0x2b, 0xd4, 0x2b, 0x33, 0xd4, 0x83, 0x28, 0x94, 0x52, 0x40,
	0xd7, 0x59, 0x94, 0x5f, 0x5a, 0x92, 0x8a, 0xa4, 0x51, 0x69, 0x03, 0x23, 0x17, 0x97, 0x2b, 0xd8,
	0xb8, 0xa0, 0xd2, 0x9c, 0x54, 0x21, 0x57, 0x2e, 0xee, 0x94, 0xd4, 0xe2, 0x92, 0xcc, 0xbc, 0xc4,
	0x92, 0xcc, 0xfc, 0x3c, 0x09, 0x46, 0x05, 0x46, 0x0d, 0x6e, 0x23, 0x65, 0x3d, 0xac, 0xa6, 0xeb,
	0x79, 0x82, 0x44, 0x83, 0x53, 0x8b, 0xca, 0x32, 0x93, 0x53, 0x83, 0x90, 0xf5, 0x09, 0x19, 0x72,
	0xb1, 0x16, 0xe4, 0x17, 0x95, 0x14, 0x4b, 0x30, 0x29, 0x30, 0x6b, 0x70, 0x1b, 0x49, 0xe3, 0x30,
	0x20, 0x20, 0xbf, 0xa8, 0x24, 0x08, 0xa2, 0x52, 0x48, 0x83, 0x4b, 0xa0, 0xb4, 0x38, 0x35, 0x1e,
	0xea, 0x35, 0xb0, 0x4a, 0x09, 0x66, 0x05, 0x46, 0x0d, 0x8e, 0x20, 0xbe, 0xd2, 0xe2, 0x54, 0x88,
	0x13, 0x03, 0x40, 0xa2, 0x49, 0x6c, 0x60, 0x97, 0x1b, 0x03, 0x02, 0x00, 0x00, 0xff, 0xff, 0xd0,
	0xb2, 0xc0, 0x9c, 0x17, 0x01, 0x00, 0x00,
}
