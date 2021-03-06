// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: login.proto

// 根据具体的微服务名称做更改

package test_apilogin_v1alpha1

import (
	fmt "fmt"
	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/gogo/protobuf/proto"
	_ "github.com/gogo/protobuf/types"
	golang_proto "github.com/golang/protobuf/proto"
	_ "github.com/grpc-ecosystem/grpc-gateway/protoc-gen-swagger/options"
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

type BaseAuthRequest struct {
}

func (m *BaseAuthRequest) Reset()         { *m = BaseAuthRequest{} }
func (m *BaseAuthRequest) String() string { return proto.CompactTextString(m) }
func (*BaseAuthRequest) ProtoMessage()    {}
func (*BaseAuthRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_67c21677aa7f4e4f, []int{0}
}
func (m *BaseAuthRequest) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *BaseAuthRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_BaseAuthRequest.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *BaseAuthRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_BaseAuthRequest.Merge(m, src)
}
func (m *BaseAuthRequest) XXX_Size() int {
	return m.Size()
}
func (m *BaseAuthRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_BaseAuthRequest.DiscardUnknown(m)
}

var xxx_messageInfo_BaseAuthRequest proto.InternalMessageInfo

type BaseAuthResponse struct {
	LoginTime string `protobuf:"bytes,1,opt,name=login_time,json=loginTime,proto3" json:"login_time,omitempty"`
}

func (m *BaseAuthResponse) Reset()         { *m = BaseAuthResponse{} }
func (m *BaseAuthResponse) String() string { return proto.CompactTextString(m) }
func (*BaseAuthResponse) ProtoMessage()    {}
func (*BaseAuthResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_67c21677aa7f4e4f, []int{1}
}
func (m *BaseAuthResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *BaseAuthResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_BaseAuthResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *BaseAuthResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_BaseAuthResponse.Merge(m, src)
}
func (m *BaseAuthResponse) XXX_Size() int {
	return m.Size()
}
func (m *BaseAuthResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_BaseAuthResponse.DiscardUnknown(m)
}

var xxx_messageInfo_BaseAuthResponse proto.InternalMessageInfo

func (m *BaseAuthResponse) GetLoginTime() string {
	if m != nil {
		return m.LoginTime
	}
	return ""
}

type JWTAuthRequest struct {
}

func (m *JWTAuthRequest) Reset()         { *m = JWTAuthRequest{} }
func (m *JWTAuthRequest) String() string { return proto.CompactTextString(m) }
func (*JWTAuthRequest) ProtoMessage()    {}
func (*JWTAuthRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_67c21677aa7f4e4f, []int{2}
}
func (m *JWTAuthRequest) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *JWTAuthRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_JWTAuthRequest.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *JWTAuthRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_JWTAuthRequest.Merge(m, src)
}
func (m *JWTAuthRequest) XXX_Size() int {
	return m.Size()
}
func (m *JWTAuthRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_JWTAuthRequest.DiscardUnknown(m)
}

var xxx_messageInfo_JWTAuthRequest proto.InternalMessageInfo

type JWTAuthResponse struct {
	LoginTime string `protobuf:"bytes,1,opt,name=login_time,json=loginTime,proto3" json:"login_time,omitempty"`
}

func (m *JWTAuthResponse) Reset()         { *m = JWTAuthResponse{} }
func (m *JWTAuthResponse) String() string { return proto.CompactTextString(m) }
func (*JWTAuthResponse) ProtoMessage()    {}
func (*JWTAuthResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_67c21677aa7f4e4f, []int{3}
}
func (m *JWTAuthResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *JWTAuthResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_JWTAuthResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *JWTAuthResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_JWTAuthResponse.Merge(m, src)
}
func (m *JWTAuthResponse) XXX_Size() int {
	return m.Size()
}
func (m *JWTAuthResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_JWTAuthResponse.DiscardUnknown(m)
}

var xxx_messageInfo_JWTAuthResponse proto.InternalMessageInfo

func (m *JWTAuthResponse) GetLoginTime() string {
	if m != nil {
		return m.LoginTime
	}
	return ""
}

func init() {
	proto.RegisterType((*BaseAuthRequest)(nil), "test.apilogin.v1alpha1.BaseAuthRequest")
	golang_proto.RegisterType((*BaseAuthRequest)(nil), "test.apilogin.v1alpha1.BaseAuthRequest")
	proto.RegisterType((*BaseAuthResponse)(nil), "test.apilogin.v1alpha1.BaseAuthResponse")
	golang_proto.RegisterType((*BaseAuthResponse)(nil), "test.apilogin.v1alpha1.BaseAuthResponse")
	proto.RegisterType((*JWTAuthRequest)(nil), "test.apilogin.v1alpha1.JWTAuthRequest")
	golang_proto.RegisterType((*JWTAuthRequest)(nil), "test.apilogin.v1alpha1.JWTAuthRequest")
	proto.RegisterType((*JWTAuthResponse)(nil), "test.apilogin.v1alpha1.JWTAuthResponse")
	golang_proto.RegisterType((*JWTAuthResponse)(nil), "test.apilogin.v1alpha1.JWTAuthResponse")
}

func init() { proto.RegisterFile("login.proto", fileDescriptor_67c21677aa7f4e4f) }
func init() { golang_proto.RegisterFile("login.proto", fileDescriptor_67c21677aa7f4e4f) }

var fileDescriptor_67c21677aa7f4e4f = []byte{
	// 295 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x90, 0xb1, 0x4e, 0xf3, 0x30,
	0x14, 0x85, 0xe3, 0xe5, 0x97, 0xea, 0x5f, 0xa2, 0xa5, 0x03, 0x42, 0x45, 0x58, 0xa8, 0x13, 0x03,
	0x69, 0xa8, 0x98, 0x19, 0xe8, 0xc8, 0x58, 0x2a, 0x31, 0x22, 0x37, 0xba, 0xb8, 0x96, 0x12, 0x5f,
	0x13, 0xdf, 0x50, 0xe5, 0x2d, 0x78, 0x8c, 0x3e, 0x06, 0x63, 0xc6, 0x8e, 0x8c, 0x10, 0xbf, 0x08,
	0x92, 0x13, 0x95, 0x22, 0x18, 0xd8, 0xee, 0x39, 0xbe, 0xe7, 0xf3, 0xd1, 0xe5, 0xff, 0x33, 0x54,
	0xda, 0x4c, 0x6c, 0x81, 0x84, 0xc3, 0x23, 0x02, 0x47, 0x13, 0x69, 0x75, 0x6b, 0x3e, 0x4f, 0x65,
	0x66, 0x57, 0x72, 0x3a, 0x8a, 0x95, 0xa6, 0x55, 0xb9, 0x9c, 0xa4, 0x98, 0x27, 0x0a, 0x15, 0x26,
	0x61, 0x7d, 0x59, 0x3e, 0x06, 0x15, 0x44, 0x98, 0x5a, 0xcc, 0xe8, 0xfa, 0xdb, 0x3a, 0xaa, 0x0c,
	0xa4, 0xd5, 0xee, 0xe7, 0x98, 0x48, 0xab, 0x13, 0x69, 0x0c, 0x92, 0x24, 0x8d, 0xc6, 0x75, 0xf1,
	0xbb, 0xfd, 0x78, 0x61, 0xd3, 0x18, 0x52, 0x74, 0x95, 0x23, 0xe8, 0xa4, 0x92, 0x04, 0x6b, 0x59,
	0xb5, 0x25, 0xd2, 0x58, 0x81, 0x89, 0xdd, 0x5a, 0x2a, 0x05, 0x45, 0x82, 0x36, 0x80, 0x7e, 0x81,
	0x9e, 0x74, 0x5f, 0xee, 0x9a, 0x43, 0x6e, 0xa9, 0x6a, 0x1f, 0xc7, 0x87, 0xbc, 0x3f, 0x93, 0x0e,
	0x6e, 0x4a, 0x5a, 0xcd, 0xe1, 0xa9, 0x04, 0x47, 0xe3, 0x29, 0x1f, 0x7c, 0x59, 0xce, 0xa2, 0x71,
	0x30, 0x3c, 0xe5, 0x3c, 0x1c, 0xe6, 0x81, 0x74, 0x0e, 0xc7, 0xec, 0x8c, 0x9d, 0xf7, 0xe6, 0xbd,
	0xe0, 0x2c, 0x74, 0x0e, 0xe3, 0x01, 0x3f, 0xb8, 0xbd, 0x5f, 0xec, 0x43, 0x2e, 0x79, 0x7f, 0xe7,
	0xfc, 0x89, 0x31, 0xbb, 0xa8, 0x3f, 0x04, 0xdb, 0x34, 0x82, 0xd5, 0x8d, 0x60, 0xdb, 0x46, 0xb0,
	0xf7, 0x46, 0xb0, 0x17, 0x2f, 0xa2, 0x8d, 0x17, 0xec, 0xd5, 0x0b, 0x56, 0x7b, 0x11, 0x6d, 0xbd,
	0x88, 0xde, 0xbc, 0x88, 0x96, 0xff, 0x42, 0xfd, 0xab, 0xcf, 0x00, 0x00, 0x00, 0xff, 0xff, 0xa2,
	0x20, 0x56, 0x14, 0xc5, 0x01, 0x00, 0x00,
}

func (this *BaseAuthRequest) Compare(that interface{}) int {
	if that == nil {
		if this == nil {
			return 0
		}
		return 1
	}

	that1, ok := that.(*BaseAuthRequest)
	if !ok {
		that2, ok := that.(BaseAuthRequest)
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
	return 0
}
func (this *BaseAuthResponse) Compare(that interface{}) int {
	if that == nil {
		if this == nil {
			return 0
		}
		return 1
	}

	that1, ok := that.(*BaseAuthResponse)
	if !ok {
		that2, ok := that.(BaseAuthResponse)
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
	if this.LoginTime != that1.LoginTime {
		if this.LoginTime < that1.LoginTime {
			return -1
		}
		return 1
	}
	return 0
}
func (this *JWTAuthRequest) Compare(that interface{}) int {
	if that == nil {
		if this == nil {
			return 0
		}
		return 1
	}

	that1, ok := that.(*JWTAuthRequest)
	if !ok {
		that2, ok := that.(JWTAuthRequest)
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
	return 0
}
func (this *JWTAuthResponse) Compare(that interface{}) int {
	if that == nil {
		if this == nil {
			return 0
		}
		return 1
	}

	that1, ok := that.(*JWTAuthResponse)
	if !ok {
		that2, ok := that.(JWTAuthResponse)
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
	if this.LoginTime != that1.LoginTime {
		if this.LoginTime < that1.LoginTime {
			return -1
		}
		return 1
	}
	return 0
}
func (this *BaseAuthRequest) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*BaseAuthRequest)
	if !ok {
		that2, ok := that.(BaseAuthRequest)
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
	return true
}
func (this *BaseAuthResponse) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*BaseAuthResponse)
	if !ok {
		that2, ok := that.(BaseAuthResponse)
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
	if this.LoginTime != that1.LoginTime {
		return false
	}
	return true
}
func (this *JWTAuthRequest) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*JWTAuthRequest)
	if !ok {
		that2, ok := that.(JWTAuthRequest)
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
	return true
}
func (this *JWTAuthResponse) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*JWTAuthResponse)
	if !ok {
		that2, ok := that.(JWTAuthResponse)
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
	if this.LoginTime != that1.LoginTime {
		return false
	}
	return true
}
func (m *BaseAuthRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *BaseAuthRequest) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *BaseAuthRequest) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	return len(dAtA) - i, nil
}

func (m *BaseAuthResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *BaseAuthResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *BaseAuthResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.LoginTime) > 0 {
		i -= len(m.LoginTime)
		copy(dAtA[i:], m.LoginTime)
		i = encodeVarintLogin(dAtA, i, uint64(len(m.LoginTime)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *JWTAuthRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *JWTAuthRequest) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *JWTAuthRequest) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	return len(dAtA) - i, nil
}

func (m *JWTAuthResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *JWTAuthResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *JWTAuthResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.LoginTime) > 0 {
		i -= len(m.LoginTime)
		copy(dAtA[i:], m.LoginTime)
		i = encodeVarintLogin(dAtA, i, uint64(len(m.LoginTime)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintLogin(dAtA []byte, offset int, v uint64) int {
	offset -= sovLogin(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *BaseAuthRequest) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	return n
}

func (m *BaseAuthResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.LoginTime)
	if l > 0 {
		n += 1 + l + sovLogin(uint64(l))
	}
	return n
}

func (m *JWTAuthRequest) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	return n
}

func (m *JWTAuthResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.LoginTime)
	if l > 0 {
		n += 1 + l + sovLogin(uint64(l))
	}
	return n
}

func sovLogin(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozLogin(x uint64) (n int) {
	return sovLogin(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *BaseAuthRequest) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowLogin
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
			return fmt.Errorf("proto: BaseAuthRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: BaseAuthRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		default:
			iNdEx = preIndex
			skippy, err := skipLogin(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthLogin
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthLogin
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
func (m *BaseAuthResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowLogin
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
			return fmt.Errorf("proto: BaseAuthResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: BaseAuthResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field LoginTime", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowLogin
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
				return ErrInvalidLengthLogin
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthLogin
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.LoginTime = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipLogin(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthLogin
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthLogin
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
func (m *JWTAuthRequest) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowLogin
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
			return fmt.Errorf("proto: JWTAuthRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: JWTAuthRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		default:
			iNdEx = preIndex
			skippy, err := skipLogin(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthLogin
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthLogin
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
func (m *JWTAuthResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowLogin
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
			return fmt.Errorf("proto: JWTAuthResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: JWTAuthResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field LoginTime", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowLogin
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
				return ErrInvalidLengthLogin
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthLogin
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.LoginTime = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipLogin(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthLogin
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthLogin
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
func skipLogin(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowLogin
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
					return 0, ErrIntOverflowLogin
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
					return 0, ErrIntOverflowLogin
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
				return 0, ErrInvalidLengthLogin
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupLogin
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthLogin
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthLogin        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowLogin          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupLogin = fmt.Errorf("proto: unexpected end of group")
)
