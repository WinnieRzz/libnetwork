// Code generated by protoc-gen-gogo.
// source: agent.proto
// DO NOT EDIT!

/*
	Package libnetwork is a generated protocol buffer package.

	It is generated from these files:
		agent.proto

	It has these top-level messages:
		EndpointRecord
		PortConfig
*/
package libnetwork

import proto "github.com/gogo/protobuf/proto"
import fmt "fmt"
import math "math"
import _ "github.com/gogo/protobuf/gogoproto"

import strings "strings"
import github_com_gogo_protobuf_proto "github.com/gogo/protobuf/proto"
import sort "sort"
import strconv "strconv"
import reflect "reflect"

import io "io"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
const _ = proto.GoGoProtoPackageIsVersion1

type PortConfig_Protocol int32

const (
	ProtocolTCP PortConfig_Protocol = 0
	ProtocolUDP PortConfig_Protocol = 1
)

var PortConfig_Protocol_name = map[int32]string{
	0: "TCP",
	1: "UDP",
}
var PortConfig_Protocol_value = map[string]int32{
	"TCP": 0,
	"UDP": 1,
}

func (x PortConfig_Protocol) String() string {
	return proto.EnumName(PortConfig_Protocol_name, int32(x))
}
func (PortConfig_Protocol) EnumDescriptor() ([]byte, []int) { return fileDescriptorAgent, []int{1, 0} }

// EndpointRecord specifies all the endpoint specific information that
// needs to gossiped to nodes participating in the network.
type EndpointRecord struct {
	// Name of the endpoint
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	// Service name of the service to which this endpoint belongs.
	ServiceName string `protobuf:"bytes,2,opt,name=service_name,json=serviceName,proto3" json:"service_name,omitempty"`
	// Service ID of the service to which this endpoint belongs.
	ServiceID string `protobuf:"bytes,3,opt,name=service_id,json=serviceId,proto3" json:"service_id,omitempty"`
	// Virtual IP of the service to which this endpoint belongs.
	VirtualIP string `protobuf:"bytes,4,opt,name=virtual_ip,json=virtualIp,proto3" json:"virtual_ip,omitempty"`
	// IP assigned to this endpoint.
	EndpointIP string `protobuf:"bytes,5,opt,name=endpoint_ip,json=endpointIp,proto3" json:"endpoint_ip,omitempty"`
	// IngressPorts exposed by the service to which this endpoint belongs.
	IngressPorts []*PortConfig `protobuf:"bytes,6,rep,name=ingress_ports,json=ingressPorts" json:"ingress_ports,omitempty"`
	// A list of aliases which are alternate names for the service
	Aliases []string `protobuf:"bytes,7,rep,name=aliases" json:"aliases,omitempty"`
}

func (m *EndpointRecord) Reset()                    { *m = EndpointRecord{} }
func (*EndpointRecord) ProtoMessage()               {}
func (*EndpointRecord) Descriptor() ([]byte, []int) { return fileDescriptorAgent, []int{0} }

func (m *EndpointRecord) GetIngressPorts() []*PortConfig {
	if m != nil {
		return m.IngressPorts
	}
	return nil
}

// PortConfig specifies an exposed port which can be
// addressed using the given name. This can be later queried
// using a service discovery api or a DNS SRV query. The node
// port specifies a port that can be used to address this
// service external to the cluster by sending a connection
// request to this port to any node on the cluster.
type PortConfig struct {
	// Name for the port. If provided the port information can
	// be queried using the name as in a DNS SRV query.
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	// Protocol for the port which is exposed.
	Protocol PortConfig_Protocol `protobuf:"varint,2,opt,name=protocol,proto3,enum=libnetwork.PortConfig_Protocol" json:"protocol,omitempty"`
	// The port which the application is exposing and is bound to.
	TargetPort uint32 `protobuf:"varint,3,opt,name=target_port,json=targetPort,proto3" json:"target_port,omitempty"`
	// PublishedPort specifies the port on which the service is
	// exposed on all nodes on the cluster. If not specified an
	// arbitrary port in the node port range is allocated by the
	// system. If specified it should be within the node port
	// range and it should be available.
	PublishedPort uint32 `protobuf:"varint,4,opt,name=published_port,json=publishedPort,proto3" json:"published_port,omitempty"`
}

func (m *PortConfig) Reset()                    { *m = PortConfig{} }
func (*PortConfig) ProtoMessage()               {}
func (*PortConfig) Descriptor() ([]byte, []int) { return fileDescriptorAgent, []int{1} }

func init() {
	proto.RegisterType((*EndpointRecord)(nil), "libnetwork.EndpointRecord")
	proto.RegisterType((*PortConfig)(nil), "libnetwork.PortConfig")
	proto.RegisterEnum("libnetwork.PortConfig_Protocol", PortConfig_Protocol_name, PortConfig_Protocol_value)
}
func (this *EndpointRecord) GoString() string {
	if this == nil {
		return "nil"
	}
	s := make([]string, 0, 11)
	s = append(s, "&libnetwork.EndpointRecord{")
	s = append(s, "Name: "+fmt.Sprintf("%#v", this.Name)+",\n")
	s = append(s, "ServiceName: "+fmt.Sprintf("%#v", this.ServiceName)+",\n")
	s = append(s, "ServiceID: "+fmt.Sprintf("%#v", this.ServiceID)+",\n")
	s = append(s, "VirtualIP: "+fmt.Sprintf("%#v", this.VirtualIP)+",\n")
	s = append(s, "EndpointIP: "+fmt.Sprintf("%#v", this.EndpointIP)+",\n")
	if this.IngressPorts != nil {
		s = append(s, "IngressPorts: "+fmt.Sprintf("%#v", this.IngressPorts)+",\n")
	}
	s = append(s, "Aliases: "+fmt.Sprintf("%#v", this.Aliases)+",\n")
	s = append(s, "}")
	return strings.Join(s, "")
}
func (this *PortConfig) GoString() string {
	if this == nil {
		return "nil"
	}
	s := make([]string, 0, 8)
	s = append(s, "&libnetwork.PortConfig{")
	s = append(s, "Name: "+fmt.Sprintf("%#v", this.Name)+",\n")
	s = append(s, "Protocol: "+fmt.Sprintf("%#v", this.Protocol)+",\n")
	s = append(s, "TargetPort: "+fmt.Sprintf("%#v", this.TargetPort)+",\n")
	s = append(s, "PublishedPort: "+fmt.Sprintf("%#v", this.PublishedPort)+",\n")
	s = append(s, "}")
	return strings.Join(s, "")
}
func valueToGoStringAgent(v interface{}, typ string) string {
	rv := reflect.ValueOf(v)
	if rv.IsNil() {
		return "nil"
	}
	pv := reflect.Indirect(rv).Interface()
	return fmt.Sprintf("func(v %v) *%v { return &v } ( %#v )", typ, typ, pv)
}
func extensionToGoStringAgent(e map[int32]github_com_gogo_protobuf_proto.Extension) string {
	if e == nil {
		return "nil"
	}
	s := "map[int32]proto.Extension{"
	keys := make([]int, 0, len(e))
	for k := range e {
		keys = append(keys, int(k))
	}
	sort.Ints(keys)
	ss := []string{}
	for _, k := range keys {
		ss = append(ss, strconv.Itoa(k)+": "+e[int32(k)].GoString())
	}
	s += strings.Join(ss, ",") + "}"
	return s
}
func (m *EndpointRecord) Marshal() (data []byte, err error) {
	size := m.Size()
	data = make([]byte, size)
	n, err := m.MarshalTo(data)
	if err != nil {
		return nil, err
	}
	return data[:n], nil
}

func (m *EndpointRecord) MarshalTo(data []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if len(m.Name) > 0 {
		data[i] = 0xa
		i++
		i = encodeVarintAgent(data, i, uint64(len(m.Name)))
		i += copy(data[i:], m.Name)
	}
	if len(m.ServiceName) > 0 {
		data[i] = 0x12
		i++
		i = encodeVarintAgent(data, i, uint64(len(m.ServiceName)))
		i += copy(data[i:], m.ServiceName)
	}
	if len(m.ServiceID) > 0 {
		data[i] = 0x1a
		i++
		i = encodeVarintAgent(data, i, uint64(len(m.ServiceID)))
		i += copy(data[i:], m.ServiceID)
	}
	if len(m.VirtualIP) > 0 {
		data[i] = 0x22
		i++
		i = encodeVarintAgent(data, i, uint64(len(m.VirtualIP)))
		i += copy(data[i:], m.VirtualIP)
	}
	if len(m.EndpointIP) > 0 {
		data[i] = 0x2a
		i++
		i = encodeVarintAgent(data, i, uint64(len(m.EndpointIP)))
		i += copy(data[i:], m.EndpointIP)
	}
	if len(m.IngressPorts) > 0 {
		for _, msg := range m.IngressPorts {
			data[i] = 0x32
			i++
			i = encodeVarintAgent(data, i, uint64(msg.Size()))
			n, err := msg.MarshalTo(data[i:])
			if err != nil {
				return 0, err
			}
			i += n
		}
	}
	if len(m.Aliases) > 0 {
		for _, s := range m.Aliases {
			data[i] = 0x3a
			i++
			l = len(s)
			for l >= 1<<7 {
				data[i] = uint8(uint64(l)&0x7f | 0x80)
				l >>= 7
				i++
			}
			data[i] = uint8(l)
			i++
			i += copy(data[i:], s)
		}
	}
	return i, nil
}

func (m *PortConfig) Marshal() (data []byte, err error) {
	size := m.Size()
	data = make([]byte, size)
	n, err := m.MarshalTo(data)
	if err != nil {
		return nil, err
	}
	return data[:n], nil
}

func (m *PortConfig) MarshalTo(data []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if len(m.Name) > 0 {
		data[i] = 0xa
		i++
		i = encodeVarintAgent(data, i, uint64(len(m.Name)))
		i += copy(data[i:], m.Name)
	}
	if m.Protocol != 0 {
		data[i] = 0x10
		i++
		i = encodeVarintAgent(data, i, uint64(m.Protocol))
	}
	if m.TargetPort != 0 {
		data[i] = 0x18
		i++
		i = encodeVarintAgent(data, i, uint64(m.TargetPort))
	}
	if m.PublishedPort != 0 {
		data[i] = 0x20
		i++
		i = encodeVarintAgent(data, i, uint64(m.PublishedPort))
	}
	return i, nil
}

func encodeFixed64Agent(data []byte, offset int, v uint64) int {
	data[offset] = uint8(v)
	data[offset+1] = uint8(v >> 8)
	data[offset+2] = uint8(v >> 16)
	data[offset+3] = uint8(v >> 24)
	data[offset+4] = uint8(v >> 32)
	data[offset+5] = uint8(v >> 40)
	data[offset+6] = uint8(v >> 48)
	data[offset+7] = uint8(v >> 56)
	return offset + 8
}
func encodeFixed32Agent(data []byte, offset int, v uint32) int {
	data[offset] = uint8(v)
	data[offset+1] = uint8(v >> 8)
	data[offset+2] = uint8(v >> 16)
	data[offset+3] = uint8(v >> 24)
	return offset + 4
}
func encodeVarintAgent(data []byte, offset int, v uint64) int {
	for v >= 1<<7 {
		data[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	data[offset] = uint8(v)
	return offset + 1
}
func (m *EndpointRecord) Size() (n int) {
	var l int
	_ = l
	l = len(m.Name)
	if l > 0 {
		n += 1 + l + sovAgent(uint64(l))
	}
	l = len(m.ServiceName)
	if l > 0 {
		n += 1 + l + sovAgent(uint64(l))
	}
	l = len(m.ServiceID)
	if l > 0 {
		n += 1 + l + sovAgent(uint64(l))
	}
	l = len(m.VirtualIP)
	if l > 0 {
		n += 1 + l + sovAgent(uint64(l))
	}
	l = len(m.EndpointIP)
	if l > 0 {
		n += 1 + l + sovAgent(uint64(l))
	}
	if len(m.IngressPorts) > 0 {
		for _, e := range m.IngressPorts {
			l = e.Size()
			n += 1 + l + sovAgent(uint64(l))
		}
	}
	if len(m.Aliases) > 0 {
		for _, s := range m.Aliases {
			l = len(s)
			n += 1 + l + sovAgent(uint64(l))
		}
	}
	return n
}

func (m *PortConfig) Size() (n int) {
	var l int
	_ = l
	l = len(m.Name)
	if l > 0 {
		n += 1 + l + sovAgent(uint64(l))
	}
	if m.Protocol != 0 {
		n += 1 + sovAgent(uint64(m.Protocol))
	}
	if m.TargetPort != 0 {
		n += 1 + sovAgent(uint64(m.TargetPort))
	}
	if m.PublishedPort != 0 {
		n += 1 + sovAgent(uint64(m.PublishedPort))
	}
	return n
}

func sovAgent(x uint64) (n int) {
	for {
		n++
		x >>= 7
		if x == 0 {
			break
		}
	}
	return n
}
func sozAgent(x uint64) (n int) {
	return sovAgent(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (this *EndpointRecord) String() string {
	if this == nil {
		return "nil"
	}
	s := strings.Join([]string{`&EndpointRecord{`,
		`Name:` + fmt.Sprintf("%v", this.Name) + `,`,
		`ServiceName:` + fmt.Sprintf("%v", this.ServiceName) + `,`,
		`ServiceID:` + fmt.Sprintf("%v", this.ServiceID) + `,`,
		`VirtualIP:` + fmt.Sprintf("%v", this.VirtualIP) + `,`,
		`EndpointIP:` + fmt.Sprintf("%v", this.EndpointIP) + `,`,
		`IngressPorts:` + strings.Replace(fmt.Sprintf("%v", this.IngressPorts), "PortConfig", "PortConfig", 1) + `,`,
		`Aliases:` + fmt.Sprintf("%v", this.Aliases) + `,`,
		`}`,
	}, "")
	return s
}
func (this *PortConfig) String() string {
	if this == nil {
		return "nil"
	}
	s := strings.Join([]string{`&PortConfig{`,
		`Name:` + fmt.Sprintf("%v", this.Name) + `,`,
		`Protocol:` + fmt.Sprintf("%v", this.Protocol) + `,`,
		`TargetPort:` + fmt.Sprintf("%v", this.TargetPort) + `,`,
		`PublishedPort:` + fmt.Sprintf("%v", this.PublishedPort) + `,`,
		`}`,
	}, "")
	return s
}
func valueToStringAgent(v interface{}) string {
	rv := reflect.ValueOf(v)
	if rv.IsNil() {
		return "nil"
	}
	pv := reflect.Indirect(rv).Interface()
	return fmt.Sprintf("*%v", pv)
}
func (m *EndpointRecord) Unmarshal(data []byte) error {
	l := len(data)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowAgent
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := data[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: EndpointRecord: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: EndpointRecord: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Name", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowAgent
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := data[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthAgent
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Name = string(data[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ServiceName", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowAgent
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := data[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthAgent
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ServiceName = string(data[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ServiceID", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowAgent
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := data[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthAgent
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ServiceID = string(data[iNdEx:postIndex])
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field VirtualIP", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowAgent
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := data[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthAgent
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.VirtualIP = string(data[iNdEx:postIndex])
			iNdEx = postIndex
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field EndpointIP", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowAgent
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := data[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthAgent
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.EndpointIP = string(data[iNdEx:postIndex])
			iNdEx = postIndex
		case 6:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field IngressPorts", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowAgent
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := data[iNdEx]
				iNdEx++
				msglen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthAgent
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.IngressPorts = append(m.IngressPorts, &PortConfig{})
			if err := m.IngressPorts[len(m.IngressPorts)-1].Unmarshal(data[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 7:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Aliases", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowAgent
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := data[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthAgent
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Aliases = append(m.Aliases, string(data[iNdEx:postIndex]))
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipAgent(data[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthAgent
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *PortConfig) Unmarshal(data []byte) error {
	l := len(data)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowAgent
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := data[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: PortConfig: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: PortConfig: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Name", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowAgent
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := data[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthAgent
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Name = string(data[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Protocol", wireType)
			}
			m.Protocol = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowAgent
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := data[iNdEx]
				iNdEx++
				m.Protocol |= (PortConfig_Protocol(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field TargetPort", wireType)
			}
			m.TargetPort = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowAgent
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := data[iNdEx]
				iNdEx++
				m.TargetPort |= (uint32(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 4:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field PublishedPort", wireType)
			}
			m.PublishedPort = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowAgent
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := data[iNdEx]
				iNdEx++
				m.PublishedPort |= (uint32(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipAgent(data[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthAgent
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func skipAgent(data []byte) (n int, err error) {
	l := len(data)
	iNdEx := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowAgent
			}
			if iNdEx >= l {
				return 0, io.ErrUnexpectedEOF
			}
			b := data[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		wireType := int(wire & 0x7)
		switch wireType {
		case 0:
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowAgent
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				iNdEx++
				if data[iNdEx-1] < 0x80 {
					break
				}
			}
			return iNdEx, nil
		case 1:
			iNdEx += 8
			return iNdEx, nil
		case 2:
			var length int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowAgent
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				b := data[iNdEx]
				iNdEx++
				length |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			iNdEx += length
			if length < 0 {
				return 0, ErrInvalidLengthAgent
			}
			return iNdEx, nil
		case 3:
			for {
				var innerWire uint64
				var start int = iNdEx
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return 0, ErrIntOverflowAgent
					}
					if iNdEx >= l {
						return 0, io.ErrUnexpectedEOF
					}
					b := data[iNdEx]
					iNdEx++
					innerWire |= (uint64(b) & 0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				innerWireType := int(innerWire & 0x7)
				if innerWireType == 4 {
					break
				}
				next, err := skipAgent(data[start:])
				if err != nil {
					return 0, err
				}
				iNdEx = start + next
			}
			return iNdEx, nil
		case 4:
			return iNdEx, nil
		case 5:
			iNdEx += 4
			return iNdEx, nil
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
	}
	panic("unreachable")
}

var (
	ErrInvalidLengthAgent = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowAgent   = fmt.Errorf("proto: integer overflow")
)

var fileDescriptorAgent = []byte{
	// 397 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x6c, 0x90, 0xbf, 0xae, 0xd3, 0x30,
	0x14, 0xc6, 0x9b, 0xdb, 0x70, 0x6f, 0x73, 0xd2, 0x84, 0xca, 0x42, 0x28, 0xea, 0x90, 0x96, 0x4a,
	0x48, 0x1d, 0x50, 0x2a, 0x95, 0xb1, 0x5b, 0x5b, 0x86, 0x2c, 0x28, 0x32, 0x7f, 0xd6, 0x2a, 0x6d,
	0x4c, 0xb0, 0x08, 0x71, 0x64, 0xbb, 0x65, 0x65, 0x44, 0xbc, 0x03, 0x13, 0x23, 0x2f, 0xc2, 0xc8,
	0xc8, 0x84, 0x68, 0x57, 0x16, 0x1e, 0x01, 0xdb, 0x49, 0x5a, 0x21, 0x75, 0xb0, 0xe4, 0xfc, 0xce,
	0xef, 0x4b, 0x4e, 0x3e, 0x70, 0xd3, 0x9c, 0x94, 0x32, 0xaa, 0x38, 0x93, 0x0c, 0x41, 0x41, 0xb7,
	0x25, 0x91, 0x1f, 0x18, 0x7f, 0x37, 0x7c, 0x90, 0xb3, 0x9c, 0x19, 0x3c, 0xd3, 0xb7, 0xda, 0x98,
	0x7c, 0xbb, 0x01, 0xff, 0x59, 0x99, 0x55, 0x8c, 0x96, 0x12, 0x93, 0x1d, 0xe3, 0x19, 0x42, 0x60,
	0x97, 0xe9, 0x7b, 0x12, 0x58, 0x63, 0x6b, 0xea, 0x60, 0x73, 0x47, 0x8f, 0xa0, 0x2f, 0x08, 0x3f,
	0xd0, 0x1d, 0xd9, 0x98, 0xd9, 0x8d, 0x99, 0xb9, 0x0d, 0x7b, 0xae, 0x95, 0x27, 0x00, 0xad, 0x42,
	0xb3, 0xa0, 0xab, 0x85, 0xa5, 0x77, 0xfa, 0x35, 0x72, 0x5e, 0xd4, 0x34, 0x5e, 0x63, 0xa7, 0x11,
	0xe2, 0x4c, 0xdb, 0x07, 0xca, 0xe5, 0x3e, 0x2d, 0x36, 0xb4, 0x0a, 0xec, 0x8b, 0xfd, 0xba, 0xa6,
	0x71, 0x82, 0x9d, 0x46, 0x88, 0x2b, 0x34, 0x03, 0x97, 0x34, 0x4b, 0x6a, 0xfd, 0x9e, 0xd1, 0x7d,
	0xa5, 0x43, 0xbb, 0xbb, 0xf2, 0xa1, 0x55, 0x54, 0x60, 0x01, 0x1e, 0x2d, 0x73, 0x4e, 0x84, 0xd8,
	0x54, 0x8c, 0x4b, 0x11, 0xdc, 0x8e, 0xbb, 0x53, 0x77, 0xfe, 0x30, 0xba, 0x14, 0x12, 0x25, 0x6a,
	0xb0, 0x62, 0xe5, 0x1b, 0x9a, 0xe3, 0x7e, 0x23, 0x6b, 0x24, 0x50, 0x00, 0x77, 0x69, 0x41, 0x53,
	0x41, 0x44, 0x70, 0xa7, 0x62, 0x0e, 0x6e, 0x1f, 0x27, 0x7f, 0x2c, 0x80, 0x4b, 0xec, 0x6a, 0x53,
	0x0b, 0xe8, 0x99, 0x66, 0x77, 0xac, 0x30, 0x2d, 0xf9, 0xf3, 0xd1, 0xf5, 0x8f, 0x46, 0x49, 0xa3,
	0xe1, 0x73, 0x00, 0x8d, 0xc0, 0x95, 0x29, 0xcf, 0x89, 0x34, 0x5b, 0x9b, 0x12, 0x3d, 0x0c, 0x35,
	0xd2, 0x49, 0xf4, 0x18, 0xfc, 0x6a, 0xbf, 0x2d, 0xa8, 0x78, 0x4b, 0xb2, 0xda, 0xb1, 0x8d, 0xe3,
	0x9d, 0xa9, 0xd6, 0x26, 0x6b, 0xe8, 0xb5, 0x6f, 0x57, 0x7f, 0xd3, 0x7d, 0xb9, 0x4a, 0x06, 0x9d,
	0xe1, 0xfd, 0xcf, 0x5f, 0xc6, 0x6e, 0x8b, 0x15, 0xd2, 0x93, 0x57, 0xeb, 0x64, 0x60, 0xfd, 0x3f,
	0x51, 0x68, 0x68, 0x7f, 0xfa, 0x1a, 0x76, 0x96, 0xc1, 0xcf, 0x63, 0xd8, 0xf9, 0x7b, 0x0c, 0xad,
	0x8f, 0xa7, 0xd0, 0xfa, 0xae, 0xce, 0x0f, 0x75, 0x7e, 0xab, 0xb3, 0xbd, 0x35, 0x1b, 0x3f, 0xfd,
	0x17, 0x00, 0x00, 0xff, 0xff, 0xc5, 0x58, 0xc7, 0xbd, 0x6d, 0x02, 0x00, 0x00,
}
