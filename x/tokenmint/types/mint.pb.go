// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: comdex/tokenmint/v1beta1/mint.proto

package types

import (
	fmt "fmt"
	github_com_cosmos_cosmos_sdk_types "github.com/cosmos/cosmos-sdk/types"
	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/gogo/protobuf/proto"
	github_com_gogo_protobuf_types "github.com/gogo/protobuf/types"
	_ "github.com/golang/protobuf/ptypes/timestamp"
	io "io"
	math "math"
	math_bits "math/bits"
	time "time"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf
var _ = time.Kitchen

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion3 // please upgrade the proto package

//app_vault_type_id will be the key for  the KVStore for this value.
type TokenMint struct {
	AppId        uint64          `protobuf:"varint,1,opt,name=app_id,json=appId,proto3" json:"app_id,omitempty" yaml:"app_id"`
	MintedTokens []*MintedTokens `protobuf:"bytes,2,rep,name=minted_tokens,json=mintedTokens,proto3" json:"minted_tokens,omitempty" yaml:"minted_tokens"`
}

func (m *TokenMint) Reset()         { *m = TokenMint{} }
func (m *TokenMint) String() string { return proto.CompactTextString(m) }
func (*TokenMint) ProtoMessage()    {}
func (*TokenMint) Descriptor() ([]byte, []int) {
	return fileDescriptor_0f3526562ed06c1c, []int{0}
}
func (m *TokenMint) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *TokenMint) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_TokenMint.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *TokenMint) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TokenMint.Merge(m, src)
}
func (m *TokenMint) XXX_Size() int {
	return m.Size()
}
func (m *TokenMint) XXX_DiscardUnknown() {
	xxx_messageInfo_TokenMint.DiscardUnknown(m)
}

var xxx_messageInfo_TokenMint proto.InternalMessageInfo

type MintedTokens struct {
	AssetId       uint64                                 `protobuf:"varint,1,opt,name=asset_id,json=assetId,proto3" json:"asset_id,omitempty" yaml:"asset_id"`
	GenesisSupply github_com_cosmos_cosmos_sdk_types.Int `protobuf:"bytes,2,opt,name=genesis_supply,json=genesisSupply,proto3,customtype=github.com/cosmos/cosmos-sdk/types.Int" json:"genesis_supply" yaml:"genesis_supply"`
	CreatedAt     time.Time                              `protobuf:"bytes,3,opt,name=created_at,json=createdAt,proto3,stdtime" json:"created_at" yaml:"created_at"`
	CurrentSupply github_com_cosmos_cosmos_sdk_types.Int `protobuf:"bytes,4,opt,name=current_supply,json=currentSupply,proto3,customtype=github.com/cosmos/cosmos-sdk/types.Int" json:"current_supply" yaml:"current_supply"`
}

func (m *MintedTokens) Reset()         { *m = MintedTokens{} }
func (m *MintedTokens) String() string { return proto.CompactTextString(m) }
func (*MintedTokens) ProtoMessage()    {}
func (*MintedTokens) Descriptor() ([]byte, []int) {
	return fileDescriptor_0f3526562ed06c1c, []int{1}
}
func (m *MintedTokens) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MintedTokens) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MintedTokens.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MintedTokens) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MintedTokens.Merge(m, src)
}
func (m *MintedTokens) XXX_Size() int {
	return m.Size()
}
func (m *MintedTokens) XXX_DiscardUnknown() {
	xxx_messageInfo_MintedTokens.DiscardUnknown(m)
}

var xxx_messageInfo_MintedTokens proto.InternalMessageInfo

func init() {
	proto.RegisterType((*TokenMint)(nil), "comdex.tokenmint.v1beta1.TokenMint")
	proto.RegisterType((*MintedTokens)(nil), "comdex.tokenmint.v1beta1.MintedTokens")
}

func init() {
	proto.RegisterFile("comdex/tokenmint/v1beta1/mint.proto", fileDescriptor_0f3526562ed06c1c)
}

var fileDescriptor_0f3526562ed06c1c = []byte{
	// 456 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xa4, 0x92, 0x3f, 0x8f, 0xd3, 0x30,
	0x18, 0xc6, 0xe3, 0xfb, 0x07, 0xf5, 0x5d, 0x41, 0x84, 0x43, 0x8a, 0x2a, 0x11, 0x57, 0x46, 0x3a,
	0x75, 0x39, 0x5b, 0xbd, 0xdb, 0xd8, 0x2e, 0x0b, 0xea, 0x70, 0x4b, 0x38, 0x24, 0xc4, 0x52, 0xb9,
	0x89, 0x1b, 0xa2, 0x6b, 0x62, 0xab, 0x76, 0x11, 0xfd, 0x16, 0xf7, 0x25, 0x90, 0xf8, 0x28, 0x1d,
	0x6f, 0x44, 0x0c, 0x01, 0xd2, 0x6f, 0x90, 0x89, 0x11, 0xd9, 0x4e, 0xb9, 0x76, 0x60, 0x62, 0x8a,
	0xdf, 0x37, 0x3f, 0x3f, 0xef, 0xf3, 0xe4, 0x0d, 0x7c, 0x95, 0x88, 0x22, 0xe5, 0x9f, 0xa9, 0x16,
	0xb7, 0xbc, 0x2c, 0xf2, 0x52, 0xd3, 0x4f, 0xc3, 0x09, 0xd7, 0x6c, 0x48, 0x4d, 0x41, 0xe4, 0x5c,
	0x68, 0xe1, 0x07, 0x0e, 0x22, 0x7f, 0x21, 0xd2, 0x42, 0xbd, 0xd3, 0x4c, 0x64, 0xc2, 0x42, 0xd4,
	0x9c, 0x1c, 0xdf, 0x43, 0x99, 0x10, 0xd9, 0x8c, 0x53, 0x5b, 0x4d, 0x16, 0x53, 0xaa, 0xf3, 0x82,
	0x2b, 0xcd, 0x0a, 0xe9, 0x00, 0xfc, 0x05, 0xc0, 0xce, 0x8d, 0x11, 0xbb, 0xce, 0x4b, 0xed, 0x0f,
	0xe1, 0x11, 0x93, 0x72, 0x9c, 0xa7, 0x01, 0xe8, 0x83, 0xc1, 0x41, 0xd4, 0xab, 0x2b, 0x74, 0x78,
	0x25, 0xe5, 0x28, 0x6d, 0x2a, 0xd4, 0x5d, 0xb2, 0x62, 0xf6, 0x1a, 0x3b, 0x00, 0xc7, 0x87, 0xcc,
	0xf4, 0x7d, 0x0e, 0xbb, 0xc6, 0x07, 0x4f, 0xc7, 0xd6, 0x93, 0x0a, 0xf6, 0xfa, 0xfb, 0x83, 0xe3,
	0x8b, 0x33, 0xf2, 0x2f, 0xa7, 0xe4, 0xda, 0xe2, 0x76, 0xa8, 0x8a, 0x82, 0xa6, 0x42, 0xa7, 0x4e,
	0x78, 0x47, 0x06, 0xc7, 0x27, 0xc5, 0x16, 0x87, 0x7f, 0xef, 0xc1, 0x93, 0xed, 0x8b, 0x3e, 0x81,
	0x8f, 0x99, 0x52, 0x5c, 0x3f, 0x98, 0x7d, 0xde, 0x54, 0xe8, 0x69, 0xeb, 0xb1, 0x7d, 0x83, 0xe3,
	0x47, 0xf6, 0x38, 0x4a, 0xfd, 0x12, 0x3e, 0xc9, 0x78, 0xc9, 0x55, 0xae, 0xc6, 0x6a, 0x21, 0xe5,
	0x6c, 0x19, 0xec, 0xf5, 0xc1, 0xa0, 0x13, 0xbd, 0x59, 0x55, 0xc8, 0xfb, 0x5e, 0xa1, 0xb3, 0x2c,
	0xd7, 0x1f, 0x17, 0x13, 0x63, 0x9b, 0x26, 0x42, 0x15, 0x42, 0xb5, 0x8f, 0x73, 0x95, 0xde, 0x52,
	0xbd, 0x94, 0x5c, 0x91, 0x51, 0xa9, 0x9b, 0x0a, 0xbd, 0x70, 0x33, 0x76, 0xd5, 0x70, 0xdc, 0x6d,
	0x1b, 0x6f, 0x6d, 0xed, 0xbf, 0x87, 0x30, 0x99, 0x73, 0x66, 0x12, 0x31, 0x1d, 0xec, 0xf7, 0xc1,
	0xe0, 0xf8, 0xa2, 0x47, 0xdc, 0x3a, 0xc8, 0x66, 0x1d, 0xe4, 0x66, 0xb3, 0x8e, 0xe8, 0xa5, 0xf1,
	0xd1, 0x54, 0xe8, 0x99, 0x53, 0x7f, 0xb8, 0x8b, 0xef, 0x7e, 0x20, 0x10, 0x77, 0xda, 0xc6, 0x95,
	0x36, 0x49, 0x92, 0xc5, 0x7c, 0xce, 0x4b, 0xbd, 0x49, 0x72, 0xf0, 0x7f, 0x49, 0x76, 0xd5, 0x70,
	0xdc, 0x6d, 0x1b, 0x2e, 0x49, 0xf4, 0x6e, 0xf5, 0x2b, 0xf4, 0xbe, 0xd6, 0xa1, 0xb7, 0xaa, 0x43,
	0x70, 0x5f, 0x87, 0xe0, 0x67, 0x1d, 0x82, 0xbb, 0x75, 0xe8, 0xdd, 0xaf, 0x43, 0xef, 0xdb, 0x3a,
	0xf4, 0x3e, 0x5c, 0xee, 0x4c, 0x34, 0x6b, 0x3f, 0x17, 0xd3, 0x69, 0x9e, 0xe4, 0x6c, 0xd6, 0xd6,
	0x74, 0xfb, 0xbf, 0xb6, 0x16, 0x26, 0x47, 0xf6, 0x23, 0x5c, 0xfe, 0x09, 0x00, 0x00, 0xff, 0xff,
	0xf5, 0xe3, 0x28, 0x4c, 0xf8, 0x02, 0x00, 0x00,
}

func (m *TokenMint) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *TokenMint) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *TokenMint) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.MintedTokens) > 0 {
		for iNdEx := len(m.MintedTokens) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.MintedTokens[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintMint(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x12
		}
	}
	if m.AppId != 0 {
		i = encodeVarintMint(dAtA, i, uint64(m.AppId))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func (m *MintedTokens) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MintedTokens) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MintedTokens) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	{
		size := m.CurrentSupply.Size()
		i -= size
		if _, err := m.CurrentSupply.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintMint(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x22
	n1, err1 := github_com_gogo_protobuf_types.StdTimeMarshalTo(m.CreatedAt, dAtA[i-github_com_gogo_protobuf_types.SizeOfStdTime(m.CreatedAt):])
	if err1 != nil {
		return 0, err1
	}
	i -= n1
	i = encodeVarintMint(dAtA, i, uint64(n1))
	i--
	dAtA[i] = 0x1a
	{
		size := m.GenesisSupply.Size()
		i -= size
		if _, err := m.GenesisSupply.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintMint(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x12
	if m.AssetId != 0 {
		i = encodeVarintMint(dAtA, i, uint64(m.AssetId))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func encodeVarintMint(dAtA []byte, offset int, v uint64) int {
	offset -= sovMint(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *TokenMint) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.AppId != 0 {
		n += 1 + sovMint(uint64(m.AppId))
	}
	if len(m.MintedTokens) > 0 {
		for _, e := range m.MintedTokens {
			l = e.Size()
			n += 1 + l + sovMint(uint64(l))
		}
	}
	return n
}

func (m *MintedTokens) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.AssetId != 0 {
		n += 1 + sovMint(uint64(m.AssetId))
	}
	l = m.GenesisSupply.Size()
	n += 1 + l + sovMint(uint64(l))
	l = github_com_gogo_protobuf_types.SizeOfStdTime(m.CreatedAt)
	n += 1 + l + sovMint(uint64(l))
	l = m.CurrentSupply.Size()
	n += 1 + l + sovMint(uint64(l))
	return n
}

func sovMint(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozMint(x uint64) (n int) {
	return sovMint(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *TokenMint) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowMint
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
			return fmt.Errorf("proto: TokenMint: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: TokenMint: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field AppId", wireType)
			}
			m.AppId = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMint
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.AppId |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field MintedTokens", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMint
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
				return ErrInvalidLengthMint
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthMint
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.MintedTokens = append(m.MintedTokens, &MintedTokens{})
			if err := m.MintedTokens[len(m.MintedTokens)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipMint(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthMint
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
func (m *MintedTokens) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowMint
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
			return fmt.Errorf("proto: MintedTokens: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MintedTokens: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field AssetId", wireType)
			}
			m.AssetId = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMint
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.AssetId |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field GenesisSupply", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMint
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
				return ErrInvalidLengthMint
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthMint
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.GenesisSupply.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field CreatedAt", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMint
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
				return ErrInvalidLengthMint
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthMint
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := github_com_gogo_protobuf_types.StdTimeUnmarshal(&m.CreatedAt, dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field CurrentSupply", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMint
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
				return ErrInvalidLengthMint
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthMint
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.CurrentSupply.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipMint(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthMint
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
func skipMint(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowMint
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
					return 0, ErrIntOverflowMint
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
					return 0, ErrIntOverflowMint
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
				return 0, ErrInvalidLengthMint
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupMint
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthMint
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthMint        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowMint          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupMint = fmt.Errorf("proto: unexpected end of group")
)
