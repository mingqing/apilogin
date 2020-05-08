// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: demo.proto

// 根据具体的微服务名称做更改

package test_apilogin_v1alpha1

import (
	fmt "fmt"
	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/gogo/protobuf/proto"
	types "github.com/gogo/protobuf/types"
	golang_proto "github.com/golang/protobuf/proto"
	_ "github.com/grpc-ecosystem/grpc-gateway/protoc-gen-swagger/options"
	v1 "github.com/grpc-kit/api/proto/v1"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	io "io"
	math "math"
	math_bits "math/bits"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = golang_proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion3 // please upgrade the proto package

type DemoRequest struct {
	// UUID 资源编号
	Uuid string `protobuf:"bytes,1,opt,name=uuid,proto3" json:"uuid,omitempty"`
	// Ping 资源内容
	Ping *v1.ExampleRequest `protobuf:"bytes,2,opt,name=ping,proto3" json:"ping,omitempty"`
}

func (m *DemoRequest) Reset()         { *m = DemoRequest{} }
func (m *DemoRequest) String() string { return proto.CompactTextString(m) }
func (*DemoRequest) ProtoMessage()    {}
func (*DemoRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_ca53982754088a9d, []int{0}
}
func (m *DemoRequest) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *DemoRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_DemoRequest.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *DemoRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DemoRequest.Merge(m, src)
}
func (m *DemoRequest) XXX_Size() int {
	return m.Size()
}
func (m *DemoRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_DemoRequest.DiscardUnknown(m)
}

var xxx_messageInfo_DemoRequest proto.InternalMessageInfo

func (m *DemoRequest) GetUuid() string {
	if m != nil {
		return m.Uuid
	}
	return ""
}

func (m *DemoRequest) GetPing() *v1.ExampleRequest {
	if m != nil {
		return m.Ping
	}
	return nil
}

type DemoResponse struct {
	// Pong 放回创建的资源
	Pong *DemoResponse_Pong `protobuf:"bytes,1,opt,name=pong,proto3" json:"pong,omitempty"`
	// Content 多个资源响应内容（无分页属性）
	Content []*v1.ExampleResponse `protobuf:"bytes,2,rep,name=content,proto3" json:"content,omitempty"`
	// Ping 返回更新的资源
	Ping *v1.ExampleResponse `protobuf:"bytes,3,opt,name=ping,proto3" json:"ping,omitempty"`
	// Empty 返回空的内容
	Empty *types.Empty `protobuf:"bytes,4,opt,name=empty,proto3" json:"empty,omitempty"`
}

func (m *DemoResponse) Reset()         { *m = DemoResponse{} }
func (m *DemoResponse) String() string { return proto.CompactTextString(m) }
func (*DemoResponse) ProtoMessage()    {}
func (*DemoResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_ca53982754088a9d, []int{1}
}
func (m *DemoResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *DemoResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_DemoResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *DemoResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DemoResponse.Merge(m, src)
}
func (m *DemoResponse) XXX_Size() int {
	return m.Size()
}
func (m *DemoResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_DemoResponse.DiscardUnknown(m)
}

var xxx_messageInfo_DemoResponse proto.InternalMessageInfo

func (m *DemoResponse) GetPong() *DemoResponse_Pong {
	if m != nil {
		return m.Pong
	}
	return nil
}

func (m *DemoResponse) GetContent() []*v1.ExampleResponse {
	if m != nil {
		return m.Content
	}
	return nil
}

func (m *DemoResponse) GetPing() *v1.ExampleResponse {
	if m != nil {
		return m.Ping
	}
	return nil
}

func (m *DemoResponse) GetEmpty() *types.Empty {
	if m != nil {
		return m.Empty
	}
	return nil
}

type DemoResponse_Pong struct {
	// UUID 资源编号
	Uuid string `protobuf:"bytes,1,opt,name=uuid,proto3" json:"uuid,omitempty"`
	// Pong 单个资源响应内容
	Pong *v1.ExampleResponse `protobuf:"bytes,2,opt,name=pong,proto3" json:"pong,omitempty"`
}

func (m *DemoResponse_Pong) Reset()         { *m = DemoResponse_Pong{} }
func (m *DemoResponse_Pong) String() string { return proto.CompactTextString(m) }
func (*DemoResponse_Pong) ProtoMessage()    {}
func (*DemoResponse_Pong) Descriptor() ([]byte, []int) {
	return fileDescriptor_ca53982754088a9d, []int{1, 0}
}
func (m *DemoResponse_Pong) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *DemoResponse_Pong) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_DemoResponse_Pong.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *DemoResponse_Pong) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DemoResponse_Pong.Merge(m, src)
}
func (m *DemoResponse_Pong) XXX_Size() int {
	return m.Size()
}
func (m *DemoResponse_Pong) XXX_DiscardUnknown() {
	xxx_messageInfo_DemoResponse_Pong.DiscardUnknown(m)
}

var xxx_messageInfo_DemoResponse_Pong proto.InternalMessageInfo

func (m *DemoResponse_Pong) GetUuid() string {
	if m != nil {
		return m.Uuid
	}
	return ""
}

func (m *DemoResponse_Pong) GetPong() *v1.ExampleResponse {
	if m != nil {
		return m.Pong
	}
	return nil
}

func init() {
	proto.RegisterType((*DemoRequest)(nil), "test.apilogin.v1alpha1.DemoRequest")
	golang_proto.RegisterType((*DemoRequest)(nil), "test.apilogin.v1alpha1.DemoRequest")
	proto.RegisterType((*DemoResponse)(nil), "test.apilogin.v1alpha1.DemoResponse")
	golang_proto.RegisterType((*DemoResponse)(nil), "test.apilogin.v1alpha1.DemoResponse")
	proto.RegisterType((*DemoResponse_Pong)(nil), "test.apilogin.v1alpha1.DemoResponse.Pong")
	golang_proto.RegisterType((*DemoResponse_Pong)(nil), "test.apilogin.v1alpha1.DemoResponse.Pong")
}

func init() { proto.RegisterFile("demo.proto", fileDescriptor_ca53982754088a9d) }
func init() { golang_proto.RegisterFile("demo.proto", fileDescriptor_ca53982754088a9d) }

var fileDescriptor_ca53982754088a9d = []byte{
	// 601 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x91, 0xcf, 0x4f, 0xd4, 0x4e,
	0x18, 0xc6, 0xb7, 0xb0, 0xdf, 0xaf, 0x71, 0xd6, 0x53, 0x0f, 0x64, 0x83, 0x49, 0xad, 0x13, 0x35,
	0xab, 0xa1, 0xb3, 0x6e, 0x05, 0x64, 0x9b, 0x90, 0x88, 0x91, 0x3b, 0xa9, 0xd1, 0xfb, 0x6c, 0x77,
	0x76, 0x68, 0xd8, 0xce, 0x8c, 0x74, 0x16, 0x44, 0xc2, 0x85, 0x80, 0x37, 0xc5, 0xf0, 0xc3, 0x03,
	0x57, 0x63, 0xc2, 0x49, 0x34, 0x1c, 0x08, 0x26, 0x44, 0x8f, 0x5c, 0x4c, 0x38, 0x7a, 0xd4, 0x2d,
	0x5d, 0xfc, 0x33, 0x4c, 0xa7, 0x45, 0x51, 0x37, 0xfe, 0xb8, 0xcd, 0x9b, 0x79, 0x3f, 0xcf, 0xfb,
	0xbc, 0xef, 0x03, 0x40, 0x9d, 0x04, 0x1c, 0x89, 0x69, 0x2e, 0xb9, 0xde, 0x27, 0x49, 0x28, 0x11,
	0x16, 0x7e, 0x93, 0x53, 0x9f, 0xa1, 0x99, 0x0a, 0x6e, 0x8a, 0x49, 0x5c, 0xe9, 0xb7, 0xa8, 0x2f,
	0x27, 0x5b, 0x35, 0xe4, 0xf1, 0xa0, 0x4c, 0x39, 0xe5, 0x65, 0xd5, 0x5e, 0x6b, 0x35, 0x54, 0xa5,
	0x0a, 0xf5, 0x4a, 0x65, 0xfa, 0x47, 0x7f, 0x68, 0xe7, 0xb4, 0x49, 0xb0, 0xf0, 0xc3, 0x5f, 0x9f,
	0x65, 0x2c, 0xfc, 0x32, 0x66, 0x8c, 0x4b, 0x2c, 0x7d, 0xce, 0xc2, 0x0c, 0xbf, 0x7b, 0x1a, 0x9f,
	0x16, 0x9e, 0x45, 0x3c, 0x1e, 0xce, 0x85, 0x92, 0x64, 0x25, 0xc5, 0x92, 0xcc, 0xe2, 0xb9, 0xd4,
	0x84, 0x67, 0x51, 0xc2, 0xac, 0x70, 0x16, 0x53, 0x4a, 0xa6, 0xcb, 0x5c, 0x28, 0xa1, 0x2e, 0xa2,
	0xe7, 0xb3, 0x91, 0xdf, 0x9c, 0x93, 0x40, 0xc8, 0xb9, 0xec, 0x13, 0xfd, 0x3c, 0x71, 0xca, 0x97,
	0xca, 0x5c, 0xba, 0xda, 0x4c, 0xa5, 0x4c, 0x1e, 0xe2, 0x40, 0x34, 0x49, 0xda, 0x0f, 0x77, 0x34,
	0x50, 0xb8, 0x43, 0x02, 0xee, 0x92, 0x07, 0x2d, 0x12, 0x4a, 0x5d, 0x07, 0xf9, 0x56, 0xcb, 0xaf,
	0x17, 0x35, 0x53, 0x2b, 0x9d, 0x75, 0xd5, 0x5b, 0xaf, 0x82, 0xbc, 0xf0, 0x19, 0x2d, 0xf6, 0x98,
	0x5a, 0xa9, 0x60, 0x5f, 0x46, 0x89, 0x2e, 0x9a, 0xf2, 0xd5, 0x79, 0x53, 0x19, 0x34, 0x53, 0x41,
	0xe3, 0xa9, 0x6e, 0x26, 0xe4, 0x2a, 0xc4, 0xb9, 0xb7, 0x3a, 0xe6, 0x82, 0x6b, 0x76, 0x29, 0x19,
	0xd1, 0x79, 0xbe, 0xd7, 0xd9, 0xd9, 0xf8, 0xb2, 0xbd, 0xdb, 0x79, 0xb3, 0x18, 0xaf, 0x6c, 0x1f,
	0xed, 0xbf, 0x3b, 0x5e, 0xdf, 0x3a, 0xde, 0x5c, 0xea, 0x3c, 0x7d, 0x19, 0xaf, 0xbc, 0x88, 0x57,
	0x16, 0x3b, 0x1b, 0xaf, 0xed, 0x4b, 0x3a, 0x9c, 0x37, 0x61, 0xc2, 0x42, 0xc7, 0x9c, 0x37, 0x21,
	0xc3, 0x01, 0x81, 0x8e, 0x09, 0x4f, 0x76, 0x81, 0xe6, 0x82, 0xb9, 0x00, 0x3f, 0xf4, 0x82, 0x73,
	0xa9, 0xeb, 0x50, 0x70, 0x16, 0x12, 0x7d, 0x14, 0xe4, 0x05, 0x67, 0x54, 0xd9, 0x2e, 0xd8, 0x57,
	0x51, 0xf7, 0xf4, 0xd1, 0x69, 0x06, 0x4d, 0x70, 0x46, 0x5d, 0x85, 0xe9, 0xb7, 0xc0, 0x19, 0x8f,
	0x33, 0x49, 0x98, 0x2c, 0xf6, 0x98, 0xbd, 0xa5, 0x82, 0x7d, 0xe5, 0x4f, 0x4b, 0xa6, 0x1a, 0xee,
	0x09, 0xa6, 0x3b, 0xd9, 0x8d, 0x7a, 0x95, 0x81, 0xbf, 0xc5, 0x15, 0xa3, 0x0f, 0x80, 0xff, 0x54,
	0x84, 0xc5, 0xbc, 0x82, 0xfb, 0x50, 0x1a, 0x30, 0x3a, 0x09, 0x18, 0x8d, 0x27, 0xbf, 0x6e, 0xda,
	0xd4, 0x7f, 0x1f, 0xe4, 0x13, 0xe7, 0x5d, 0x93, 0x72, 0xb2, 0x33, 0xf4, 0xfc, 0xa3, 0x0b, 0xce,
	0xa8, 0xf3, 0x68, 0x75, 0x6c, 0x16, 0x5c, 0xb4, 0x2f, 0x7c, 0x8f, 0x2a, 0x5e, 0x7b, 0x12, 0xbf,
	0x5d, 0x3f, 0xde, 0x5c, 0x8a, 0x97, 0x77, 0x8f, 0xf6, 0xd7, 0xe2, 0xc7, 0xcb, 0xf1, 0xab, 0x3d,
	0xdb, 0xd5, 0x27, 0xe6, 0x61, 0x32, 0x0d, 0x3a, 0xb0, 0x5a, 0x6d, 0x10, 0xdc, 0xa8, 0x0d, 0x59,
	0x35, 0x52, 0x1f, 0xb6, 0x06, 0xeb, 0xb8, 0x61, 0x55, 0xed, 0x9b, 0xd8, 0x1a, 0xae, 0x62, 0x1b,
	0xd7, 0x46, 0xae, 0x7b, 0x83, 0x23, 0x43, 0x70, 0xc0, 0x84, 0xc9, 0x80, 0xdf, 0xe4, 0x79, 0x7b,
	0xe0, 0xe0, 0xb3, 0xa1, 0x6d, 0xb5, 0x0d, 0xed, 0xa0, 0x6d, 0x68, 0x87, 0x6d, 0x43, 0xfb, 0xd4,
	0x36, 0xb4, 0x67, 0x91, 0x91, 0xdb, 0x8a, 0x0c, 0xed, 0x7d, 0x64, 0x68, 0x07, 0x91, 0x91, 0x3b,
	0x8c, 0x8c, 0xdc, 0xc7, 0xc8, 0xc8, 0xd5, 0xfe, 0x57, 0x9b, 0xdc, 0xf8, 0x1a, 0x00, 0x00, 0xff,
	0xff, 0x39, 0xd7, 0x72, 0x29, 0xf0, 0x03, 0x00, 0x00,
}

func (this *DemoRequest) Compare(that interface{}) int {
	if that == nil {
		if this == nil {
			return 0
		}
		return 1
	}

	that1, ok := that.(*DemoRequest)
	if !ok {
		that2, ok := that.(DemoRequest)
		if ok {
			that1 = &that2
		} else {
			return 1
		}
	}
	if that1 == nil {
		if this == nil {
			return 0
		}
		return 1
	} else if this == nil {
		return -1
	}
	if this.Uuid != that1.Uuid {
		if this.Uuid < that1.Uuid {
			return -1
		}
		return 1
	}
	if c := this.Ping.Compare(that1.Ping); c != 0 {
		return c
	}
	return 0
}
func (this *DemoResponse) Compare(that interface{}) int {
	if that == nil {
		if this == nil {
			return 0
		}
		return 1
	}

	that1, ok := that.(*DemoResponse)
	if !ok {
		that2, ok := that.(DemoResponse)
		if ok {
			that1 = &that2
		} else {
			return 1
		}
	}
	if that1 == nil {
		if this == nil {
			return 0
		}
		return 1
	} else if this == nil {
		return -1
	}
	if c := this.Pong.Compare(that1.Pong); c != 0 {
		return c
	}
	if len(this.Content) != len(that1.Content) {
		if len(this.Content) < len(that1.Content) {
			return -1
		}
		return 1
	}
	for i := range this.Content {
		if c := this.Content[i].Compare(that1.Content[i]); c != 0 {
			return c
		}
	}
	if c := this.Ping.Compare(that1.Ping); c != 0 {
		return c
	}
	if c := this.Empty.Compare(that1.Empty); c != 0 {
		return c
	}
	return 0
}
func (this *DemoResponse_Pong) Compare(that interface{}) int {
	if that == nil {
		if this == nil {
			return 0
		}
		return 1
	}

	that1, ok := that.(*DemoResponse_Pong)
	if !ok {
		that2, ok := that.(DemoResponse_Pong)
		if ok {
			that1 = &that2
		} else {
			return 1
		}
	}
	if that1 == nil {
		if this == nil {
			return 0
		}
		return 1
	} else if this == nil {
		return -1
	}
	if this.Uuid != that1.Uuid {
		if this.Uuid < that1.Uuid {
			return -1
		}
		return 1
	}
	if c := this.Pong.Compare(that1.Pong); c != 0 {
		return c
	}
	return 0
}
func (this *DemoRequest) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*DemoRequest)
	if !ok {
		that2, ok := that.(DemoRequest)
		if ok {
			that1 = &that2
		} else {
			return false
		}
	}
	if that1 == nil {
		return this == nil
	} else if this == nil {
		return false
	}
	if this.Uuid != that1.Uuid {
		return false
	}
	if !this.Ping.Equal(that1.Ping) {
		return false
	}
	return true
}
func (this *DemoResponse) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*DemoResponse)
	if !ok {
		that2, ok := that.(DemoResponse)
		if ok {
			that1 = &that2
		} else {
			return false
		}
	}
	if that1 == nil {
		return this == nil
	} else if this == nil {
		return false
	}
	if !this.Pong.Equal(that1.Pong) {
		return false
	}
	if len(this.Content) != len(that1.Content) {
		return false
	}
	for i := range this.Content {
		if !this.Content[i].Equal(that1.Content[i]) {
			return false
		}
	}
	if !this.Ping.Equal(that1.Ping) {
		return false
	}
	if !this.Empty.Equal(that1.Empty) {
		return false
	}
	return true
}
func (this *DemoResponse_Pong) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*DemoResponse_Pong)
	if !ok {
		that2, ok := that.(DemoResponse_Pong)
		if ok {
			that1 = &that2
		} else {
			return false
		}
	}
	if that1 == nil {
		return this == nil
	} else if this == nil {
		return false
	}
	if this.Uuid != that1.Uuid {
		return false
	}
	if !this.Pong.Equal(that1.Pong) {
		return false
	}
	return true
}
func (m *DemoRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *DemoRequest) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *DemoRequest) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.Ping != nil {
		{
			size, err := m.Ping.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintDemo(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x12
	}
	if len(m.Uuid) > 0 {
		i -= len(m.Uuid)
		copy(dAtA[i:], m.Uuid)
		i = encodeVarintDemo(dAtA, i, uint64(len(m.Uuid)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *DemoResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *DemoResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *DemoResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.Empty != nil {
		{
			size, err := m.Empty.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintDemo(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x22
	}
	if m.Ping != nil {
		{
			size, err := m.Ping.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintDemo(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x1a
	}
	if len(m.Content) > 0 {
		for iNdEx := len(m.Content) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.Content[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintDemo(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x12
		}
	}
	if m.Pong != nil {
		{
			size, err := m.Pong.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintDemo(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *DemoResponse_Pong) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *DemoResponse_Pong) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *DemoResponse_Pong) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.Pong != nil {
		{
			size, err := m.Pong.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintDemo(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x12
	}
	if len(m.Uuid) > 0 {
		i -= len(m.Uuid)
		copy(dAtA[i:], m.Uuid)
		i = encodeVarintDemo(dAtA, i, uint64(len(m.Uuid)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintDemo(dAtA []byte, offset int, v uint64) int {
	offset -= sovDemo(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *DemoRequest) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Uuid)
	if l > 0 {
		n += 1 + l + sovDemo(uint64(l))
	}
	if m.Ping != nil {
		l = m.Ping.Size()
		n += 1 + l + sovDemo(uint64(l))
	}
	return n
}

func (m *DemoResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Pong != nil {
		l = m.Pong.Size()
		n += 1 + l + sovDemo(uint64(l))
	}
	if len(m.Content) > 0 {
		for _, e := range m.Content {
			l = e.Size()
			n += 1 + l + sovDemo(uint64(l))
		}
	}
	if m.Ping != nil {
		l = m.Ping.Size()
		n += 1 + l + sovDemo(uint64(l))
	}
	if m.Empty != nil {
		l = m.Empty.Size()
		n += 1 + l + sovDemo(uint64(l))
	}
	return n
}

func (m *DemoResponse_Pong) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Uuid)
	if l > 0 {
		n += 1 + l + sovDemo(uint64(l))
	}
	if m.Pong != nil {
		l = m.Pong.Size()
		n += 1 + l + sovDemo(uint64(l))
	}
	return n
}

func sovDemo(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozDemo(x uint64) (n int) {
	return sovDemo(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *DemoRequest) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowDemo
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: DemoRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: DemoRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Uuid", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowDemo
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthDemo
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthDemo
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Uuid = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Ping", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowDemo
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthDemo
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthDemo
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Ping == nil {
				m.Ping = &v1.ExampleRequest{}
			}
			if err := m.Ping.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipDemo(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthDemo
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthDemo
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
func (m *DemoResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowDemo
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: DemoResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: DemoResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Pong", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowDemo
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthDemo
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthDemo
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Pong == nil {
				m.Pong = &DemoResponse_Pong{}
			}
			if err := m.Pong.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Content", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowDemo
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthDemo
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthDemo
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Content = append(m.Content, &v1.ExampleResponse{})
			if err := m.Content[len(m.Content)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Ping", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowDemo
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthDemo
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthDemo
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Ping == nil {
				m.Ping = &v1.ExampleResponse{}
			}
			if err := m.Ping.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Empty", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowDemo
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthDemo
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthDemo
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Empty == nil {
				m.Empty = &types.Empty{}
			}
			if err := m.Empty.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipDemo(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthDemo
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthDemo
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
func (m *DemoResponse_Pong) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowDemo
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: Pong: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Pong: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Uuid", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowDemo
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthDemo
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthDemo
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Uuid = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Pong", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowDemo
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthDemo
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthDemo
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Pong == nil {
				m.Pong = &v1.ExampleResponse{}
			}
			if err := m.Pong.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipDemo(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthDemo
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthDemo
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
func skipDemo(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowDemo
			}
			if iNdEx >= l {
				return 0, io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
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
					return 0, ErrIntOverflowDemo
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				iNdEx++
				if dAtA[iNdEx-1] < 0x80 {
					break
				}
			}
		case 1:
			iNdEx += 8
		case 2:
			var length int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowDemo
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				length |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if length < 0 {
				return 0, ErrInvalidLengthDemo
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupDemo
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthDemo
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthDemo        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowDemo          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupDemo = fmt.Errorf("proto: unexpected end of group")
)
