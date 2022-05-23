// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: comdex/liquidation/v1beta1/msg.proto

package types

import (
	context "context"
	fmt "fmt"
	_ "github.com/gogo/protobuf/gogoproto"
	grpc1 "github.com/gogo/protobuf/grpc"
	proto "github.com/gogo/protobuf/proto"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	io "io"
	math "math"
	math_bits "math/bits"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion3 // please upgrade the proto package

type WhitelistAppId struct {
	AppMappingId uint64 `protobuf:"varint,1,opt,name=app_mapping_id,json=appMappingId,proto3" json:"app_mapping_id,omitempty" yaml:"app_mapping_id"`
	From         string `protobuf:"bytes,2,opt,name=from,proto3" json:"from,omitempty" yaml:"from"`
}

func (m *WhitelistAppId) Reset()         { *m = WhitelistAppId{} }
func (m *WhitelistAppId) String() string { return proto.CompactTextString(m) }
func (*WhitelistAppId) ProtoMessage()    {}
func (*WhitelistAppId) Descriptor() ([]byte, []int) {
	return fileDescriptor_aee4595774e10d25, []int{0}
}
func (m *WhitelistAppId) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *WhitelistAppId) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_WhitelistAppId.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *WhitelistAppId) XXX_Merge(src proto.Message) {
	xxx_messageInfo_WhitelistAppId.Merge(m, src)
}
func (m *WhitelistAppId) XXX_Size() int {
	return m.Size()
}
func (m *WhitelistAppId) XXX_DiscardUnknown() {
	xxx_messageInfo_WhitelistAppId.DiscardUnknown(m)
}

var xxx_messageInfo_WhitelistAppId proto.InternalMessageInfo

type RemoveWhitelistAppId struct {
	AppMappingId uint64 `protobuf:"varint,1,opt,name=app_mapping_id,json=appMappingId,proto3" json:"app_mapping_id,omitempty" yaml:"app_mapping_id"`
	From         string `protobuf:"bytes,2,opt,name=from,proto3" json:"from,omitempty" yaml:"from"`
}

func (m *RemoveWhitelistAppId) Reset()         { *m = RemoveWhitelistAppId{} }
func (m *RemoveWhitelistAppId) String() string { return proto.CompactTextString(m) }
func (*RemoveWhitelistAppId) ProtoMessage()    {}
func (*RemoveWhitelistAppId) Descriptor() ([]byte, []int) {
	return fileDescriptor_aee4595774e10d25, []int{1}
}
func (m *RemoveWhitelistAppId) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *RemoveWhitelistAppId) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_RemoveWhitelistAppId.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *RemoveWhitelistAppId) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RemoveWhitelistAppId.Merge(m, src)
}
func (m *RemoveWhitelistAppId) XXX_Size() int {
	return m.Size()
}
func (m *RemoveWhitelistAppId) XXX_DiscardUnknown() {
	xxx_messageInfo_RemoveWhitelistAppId.DiscardUnknown(m)
}

var xxx_messageInfo_RemoveWhitelistAppId proto.InternalMessageInfo

type MsgWhitelistAppIdResponse struct {
}

func (m *MsgWhitelistAppIdResponse) Reset()         { *m = MsgWhitelistAppIdResponse{} }
func (m *MsgWhitelistAppIdResponse) String() string { return proto.CompactTextString(m) }
func (*MsgWhitelistAppIdResponse) ProtoMessage()    {}
func (*MsgWhitelistAppIdResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_aee4595774e10d25, []int{2}
}
func (m *MsgWhitelistAppIdResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgWhitelistAppIdResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgWhitelistAppIdResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgWhitelistAppIdResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgWhitelistAppIdResponse.Merge(m, src)
}
func (m *MsgWhitelistAppIdResponse) XXX_Size() int {
	return m.Size()
}
func (m *MsgWhitelistAppIdResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgWhitelistAppIdResponse.DiscardUnknown(m)
}

var xxx_messageInfo_MsgWhitelistAppIdResponse proto.InternalMessageInfo

type MsgRemoveWhitelistAppIdResponse struct {
}

func (m *MsgRemoveWhitelistAppIdResponse) Reset()         { *m = MsgRemoveWhitelistAppIdResponse{} }
func (m *MsgRemoveWhitelistAppIdResponse) String() string { return proto.CompactTextString(m) }
func (*MsgRemoveWhitelistAppIdResponse) ProtoMessage()    {}
func (*MsgRemoveWhitelistAppIdResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_aee4595774e10d25, []int{3}
}
func (m *MsgRemoveWhitelistAppIdResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgRemoveWhitelistAppIdResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgRemoveWhitelistAppIdResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgRemoveWhitelistAppIdResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgRemoveWhitelistAppIdResponse.Merge(m, src)
}
func (m *MsgRemoveWhitelistAppIdResponse) XXX_Size() int {
	return m.Size()
}
func (m *MsgRemoveWhitelistAppIdResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgRemoveWhitelistAppIdResponse.DiscardUnknown(m)
}

var xxx_messageInfo_MsgRemoveWhitelistAppIdResponse proto.InternalMessageInfo

func init() {
	proto.RegisterType((*WhitelistAppId)(nil), "comdex.liquidation.v1beta1.WhitelistAppId")
	proto.RegisterType((*RemoveWhitelistAppId)(nil), "comdex.liquidation.v1beta1.RemoveWhitelistAppId")
	proto.RegisterType((*MsgWhitelistAppIdResponse)(nil), "comdex.liquidation.v1beta1.MsgWhitelistAppIdResponse")
	proto.RegisterType((*MsgRemoveWhitelistAppIdResponse)(nil), "comdex.liquidation.v1beta1.MsgRemoveWhitelistAppIdResponse")
}

func init() {
	proto.RegisterFile("comdex/liquidation/v1beta1/msg.proto", fileDescriptor_aee4595774e10d25)
}

var fileDescriptor_aee4595774e10d25 = []byte{
	// 346 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x52, 0x49, 0xce, 0xcf, 0x4d,
	0x49, 0xad, 0xd0, 0xcf, 0xc9, 0x2c, 0x2c, 0xcd, 0x4c, 0x49, 0x2c, 0xc9, 0xcc, 0xcf, 0xd3, 0x2f,
	0x33, 0x4c, 0x4a, 0x2d, 0x49, 0x34, 0xd4, 0xcf, 0x2d, 0x4e, 0xd7, 0x2b, 0x28, 0xca, 0x2f, 0xc9,
	0x17, 0x92, 0x82, 0xa8, 0xd2, 0x43, 0x52, 0xa5, 0x07, 0x55, 0x25, 0x25, 0x92, 0x9e, 0x9f, 0x9e,
	0x0f, 0x56, 0xa6, 0x0f, 0x62, 0x41, 0x74, 0x28, 0x95, 0x71, 0xf1, 0x85, 0x67, 0x64, 0x96, 0xa4,
	0xe6, 0x64, 0x16, 0x97, 0x38, 0x16, 0x14, 0x78, 0xa6, 0x08, 0xd9, 0x73, 0xf1, 0x25, 0x16, 0x14,
	0xc4, 0xe7, 0x26, 0x16, 0x14, 0x64, 0xe6, 0xa5, 0xc7, 0x67, 0xa6, 0x48, 0x30, 0x2a, 0x30, 0x6a,
	0xb0, 0x38, 0x49, 0x7e, 0xba, 0x27, 0x2f, 0x5a, 0x99, 0x98, 0x9b, 0x63, 0xa5, 0x84, 0x2a, 0xaf,
	0x14, 0xc4, 0x93, 0x58, 0x50, 0xe0, 0x0b, 0xe1, 0x7b, 0xa6, 0x08, 0x29, 0x73, 0xb1, 0xa4, 0x15,
	0xe5, 0xe7, 0x4a, 0x30, 0x29, 0x30, 0x6a, 0x70, 0x3a, 0xf1, 0x7f, 0xba, 0x27, 0xcf, 0x0d, 0xd1,
	0x06, 0x12, 0x55, 0x0a, 0x02, 0x4b, 0x2a, 0xd5, 0x70, 0x89, 0x04, 0xa5, 0xe6, 0xe6, 0x97, 0xa5,
	0x0e, 0x88, 0xed, 0xd2, 0x5c, 0x92, 0xbe, 0xc5, 0xe9, 0xa8, 0x56, 0x07, 0xa5, 0x16, 0x17, 0xe4,
	0xe7, 0x15, 0xa7, 0x2a, 0x29, 0x72, 0xc9, 0xfb, 0x16, 0xa7, 0x63, 0x73, 0x1d, 0x4c, 0x89, 0xd1,
	0x3f, 0x46, 0x2e, 0x66, 0xdf, 0xe2, 0x74, 0xa1, 0x42, 0x2e, 0x1e, 0x64, 0x15, 0x42, 0x5a, 0x7a,
	0xb8, 0x23, 0x40, 0x0f, 0xd5, 0x2c, 0x29, 0x53, 0x7c, 0x6a, 0x71, 0xba, 0x4e, 0xa8, 0x99, 0x91,
	0x4b, 0x08, 0xd3, 0x6d, 0x42, 0x06, 0xf8, 0x4c, 0xc3, 0xe6, 0x17, 0x29, 0x6b, 0x02, 0xf6, 0xe3,
	0x0b, 0x00, 0xa7, 0xf0, 0x13, 0x0f, 0xe5, 0x18, 0x56, 0x3c, 0x92, 0x63, 0x38, 0xf1, 0x48, 0x8e,
	0xf1, 0xc2, 0x23, 0x39, 0xc6, 0x07, 0x8f, 0xe4, 0x18, 0x27, 0x3c, 0x96, 0x63, 0xb8, 0xf0, 0x58,
	0x8e, 0xe1, 0xc6, 0x63, 0x39, 0x86, 0x28, 0xd3, 0xf4, 0xcc, 0x92, 0x8c, 0xd2, 0x24, 0x90, 0x25,
	0xfa, 0x10, 0x8b, 0x74, 0xf3, 0xd3, 0xd2, 0x32, 0x93, 0x33, 0x13, 0x73, 0xa0, 0x7c, 0x7d, 0xd4,
	0xd4, 0x5c, 0x52, 0x59, 0x90, 0x5a, 0x9c, 0xc4, 0x06, 0x4e, 0x96, 0xc6, 0x80, 0x00, 0x00, 0x00,
	0xff, 0xff, 0x4a, 0x6a, 0xab, 0x6e, 0xf0, 0x02, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// MsgClient is the client API for Msg service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type MsgClient interface {
	WhitelistApp(ctx context.Context, in *WhitelistAppId, opts ...grpc.CallOption) (*MsgWhitelistAppIdResponse, error)
	RemoveWhitelistApp(ctx context.Context, in *RemoveWhitelistAppId, opts ...grpc.CallOption) (*MsgRemoveWhitelistAppIdResponse, error)
}

type msgClient struct {
	cc grpc1.ClientConn
}

func NewMsgClient(cc grpc1.ClientConn) MsgClient {
	return &msgClient{cc}
}

func (c *msgClient) WhitelistApp(ctx context.Context, in *WhitelistAppId, opts ...grpc.CallOption) (*MsgWhitelistAppIdResponse, error) {
	out := new(MsgWhitelistAppIdResponse)
	err := c.cc.Invoke(ctx, "/comdex.liquidation.v1beta1.Msg/WhitelistApp", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *msgClient) RemoveWhitelistApp(ctx context.Context, in *RemoveWhitelistAppId, opts ...grpc.CallOption) (*MsgRemoveWhitelistAppIdResponse, error) {
	out := new(MsgRemoveWhitelistAppIdResponse)
	err := c.cc.Invoke(ctx, "/comdex.liquidation.v1beta1.Msg/RemoveWhitelistApp", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// MsgServer is the server API for Msg service.
type MsgServer interface {
	WhitelistApp(context.Context, *WhitelistAppId) (*MsgWhitelistAppIdResponse, error)
	RemoveWhitelistApp(context.Context, *RemoveWhitelistAppId) (*MsgRemoveWhitelistAppIdResponse, error)
}

// UnimplementedMsgServer can be embedded to have forward compatible implementations.
type UnimplementedMsgServer struct {
}

func (*UnimplementedMsgServer) WhitelistApp(ctx context.Context, req *WhitelistAppId) (*MsgWhitelistAppIdResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method WhitelistApp not implemented")
}
func (*UnimplementedMsgServer) RemoveWhitelistApp(ctx context.Context, req *RemoveWhitelistAppId) (*MsgRemoveWhitelistAppIdResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RemoveWhitelistApp not implemented")
}

func RegisterMsgServer(s grpc1.Server, srv MsgServer) {
	s.RegisterService(&_Msg_serviceDesc, srv)
}

func _Msg_WhitelistApp_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(WhitelistAppId)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServer).WhitelistApp(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/comdex.liquidation.v1beta1.Msg/WhitelistApp",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServer).WhitelistApp(ctx, req.(*WhitelistAppId))
	}
	return interceptor(ctx, in, info, handler)
}

func _Msg_RemoveWhitelistApp_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RemoveWhitelistAppId)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServer).RemoveWhitelistApp(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/comdex.liquidation.v1beta1.Msg/RemoveWhitelistApp",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServer).RemoveWhitelistApp(ctx, req.(*RemoveWhitelistAppId))
	}
	return interceptor(ctx, in, info, handler)
}

var _Msg_serviceDesc = grpc.ServiceDesc{
	ServiceName: "comdex.liquidation.v1beta1.Msg",
	HandlerType: (*MsgServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "WhitelistApp",
			Handler:    _Msg_WhitelistApp_Handler,
		},
		{
			MethodName: "RemoveWhitelistApp",
			Handler:    _Msg_RemoveWhitelistApp_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "comdex/liquidation/v1beta1/msg.proto",
}

func (m *WhitelistAppId) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *WhitelistAppId) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *WhitelistAppId) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.From) > 0 {
		i -= len(m.From)
		copy(dAtA[i:], m.From)
		i = encodeVarintMsg(dAtA, i, uint64(len(m.From)))
		i--
		dAtA[i] = 0x12
	}
	if m.AppMappingId != 0 {
		i = encodeVarintMsg(dAtA, i, uint64(m.AppMappingId))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func (m *RemoveWhitelistAppId) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *RemoveWhitelistAppId) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *RemoveWhitelistAppId) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.From) > 0 {
		i -= len(m.From)
		copy(dAtA[i:], m.From)
		i = encodeVarintMsg(dAtA, i, uint64(len(m.From)))
		i--
		dAtA[i] = 0x12
	}
	if m.AppMappingId != 0 {
		i = encodeVarintMsg(dAtA, i, uint64(m.AppMappingId))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func (m *MsgWhitelistAppIdResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgWhitelistAppIdResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgWhitelistAppIdResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	return len(dAtA) - i, nil
}

func (m *MsgRemoveWhitelistAppIdResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgRemoveWhitelistAppIdResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgRemoveWhitelistAppIdResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	return len(dAtA) - i, nil
}

func encodeVarintMsg(dAtA []byte, offset int, v uint64) int {
	offset -= sovMsg(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *WhitelistAppId) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.AppMappingId != 0 {
		n += 1 + sovMsg(uint64(m.AppMappingId))
	}
	l = len(m.From)
	if l > 0 {
		n += 1 + l + sovMsg(uint64(l))
	}
	return n
}

func (m *RemoveWhitelistAppId) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.AppMappingId != 0 {
		n += 1 + sovMsg(uint64(m.AppMappingId))
	}
	l = len(m.From)
	if l > 0 {
		n += 1 + l + sovMsg(uint64(l))
	}
	return n
}

func (m *MsgWhitelistAppIdResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	return n
}

func (m *MsgRemoveWhitelistAppIdResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	return n
}

func sovMsg(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozMsg(x uint64) (n int) {
	return sovMsg(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *WhitelistAppId) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowMsg
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
			return fmt.Errorf("proto: WhitelistAppId: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: WhitelistAppId: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field AppMappingId", wireType)
			}
			m.AppMappingId = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMsg
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.AppMappingId |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field From", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMsg
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
				return ErrInvalidLengthMsg
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthMsg
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.From = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipMsg(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthMsg
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
func (m *RemoveWhitelistAppId) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowMsg
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
			return fmt.Errorf("proto: RemoveWhitelistAppId: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: RemoveWhitelistAppId: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field AppMappingId", wireType)
			}
			m.AppMappingId = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMsg
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.AppMappingId |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field From", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMsg
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
				return ErrInvalidLengthMsg
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthMsg
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.From = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipMsg(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthMsg
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
func (m *MsgWhitelistAppIdResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowMsg
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
			return fmt.Errorf("proto: MsgWhitelistAppIdResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgWhitelistAppIdResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		default:
			iNdEx = preIndex
			skippy, err := skipMsg(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthMsg
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
func (m *MsgRemoveWhitelistAppIdResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowMsg
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
			return fmt.Errorf("proto: MsgRemoveWhitelistAppIdResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgRemoveWhitelistAppIdResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		default:
			iNdEx = preIndex
			skippy, err := skipMsg(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthMsg
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
func skipMsg(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowMsg
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
					return 0, ErrIntOverflowMsg
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
					return 0, ErrIntOverflowMsg
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
				return 0, ErrInvalidLengthMsg
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupMsg
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthMsg
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthMsg        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowMsg          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupMsg = fmt.Errorf("proto: unexpected end of group")
)