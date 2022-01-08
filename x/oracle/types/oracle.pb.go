// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: comdex/oracle/v1beta1/oracle.proto

package types

import (
	fmt "fmt"
	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/gogo/protobuf/proto"
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

type Market struct {
	Symbol   string `protobuf:"bytes,1,opt,name=symbol,proto3" json:"symbol,omitempty" yaml:"symbol"`
	ScriptID uint64 `protobuf:"varint,2,opt,name=script_id,json=scriptId,proto3" json:"script_id,omitempty" yaml:"script_id"`
}

func (m *Market) Reset()         { *m = Market{} }
func (m *Market) String() string { return proto.CompactTextString(m) }
func (*Market) ProtoMessage()    {}
func (*Market) Descriptor() ([]byte, []int) {
	return fileDescriptor_52ea76e22e2125a9, []int{0}
}
func (m *Market) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Market) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Market.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Market) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Market.Merge(m, src)
}
func (m *Market) XXX_Size() int {
	return m.Size()
}
func (m *Market) XXX_DiscardUnknown() {
	xxx_messageInfo_Market.DiscardUnknown(m)
}

var xxx_messageInfo_Market proto.InternalMessageInfo

type FetchPriceCallData struct {
	Symbols    []string `protobuf:"bytes,1,rep,name=symbols,proto3" json:"symbols,omitempty" yaml:"symbols"`
	Multiplier uint64   `protobuf:"varint,2,opt,name=multiplier,proto3" json:"multiplier,omitempty" yaml:"multiplier"`
}

func (m *FetchPriceCallData) Reset()         { *m = FetchPriceCallData{} }
func (m *FetchPriceCallData) String() string { return proto.CompactTextString(m) }
func (*FetchPriceCallData) ProtoMessage()    {}
func (*FetchPriceCallData) Descriptor() ([]byte, []int) {
	return fileDescriptor_52ea76e22e2125a9, []int{1}
}
func (m *FetchPriceCallData) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *FetchPriceCallData) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_FetchPriceCallData.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *FetchPriceCallData) XXX_Merge(src proto.Message) {
	xxx_messageInfo_FetchPriceCallData.Merge(m, src)
}
func (m *FetchPriceCallData) XXX_Size() int {
	return m.Size()
}
func (m *FetchPriceCallData) XXX_DiscardUnknown() {
	xxx_messageInfo_FetchPriceCallData.DiscardUnknown(m)
}

var xxx_messageInfo_FetchPriceCallData proto.InternalMessageInfo

type FetchPriceResult struct {
	Rates []uint64 `protobuf:"varint,1,rep,packed,name=rates,proto3" json:"rates,omitempty" yaml:"rates"`
}

func (m *FetchPriceResult) Reset()         { *m = FetchPriceResult{} }
func (m *FetchPriceResult) String() string { return proto.CompactTextString(m) }
func (*FetchPriceResult) ProtoMessage()    {}
func (*FetchPriceResult) Descriptor() ([]byte, []int) {
	return fileDescriptor_52ea76e22e2125a9, []int{2}
}
func (m *FetchPriceResult) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *FetchPriceResult) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_FetchPriceResult.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *FetchPriceResult) XXX_Merge(src proto.Message) {
	xxx_messageInfo_FetchPriceResult.Merge(m, src)
}
func (m *FetchPriceResult) XXX_Size() int {
	return m.Size()
}
func (m *FetchPriceResult) XXX_DiscardUnknown() {
	xxx_messageInfo_FetchPriceResult.DiscardUnknown(m)
}

var xxx_messageInfo_FetchPriceResult proto.InternalMessageInfo

func init() {
	proto.RegisterType((*Market)(nil), "comdex.oracle.v1beta1.Market")
	proto.RegisterType((*FetchPriceCallData)(nil), "comdex.oracle.v1beta1.FetchPriceCallData")
	proto.RegisterType((*FetchPriceResult)(nil), "comdex.oracle.v1beta1.FetchPriceResult")
}

func init() {
	proto.RegisterFile("comdex/oracle/v1beta1/oracle.proto", fileDescriptor_52ea76e22e2125a9)
}

var fileDescriptor_52ea76e22e2125a9 = []byte{
	// 348 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x4c, 0x91, 0xc1, 0x4a, 0xfb, 0x40,
	0x10, 0xc6, 0x93, 0xff, 0xbf, 0xd6, 0x76, 0x51, 0x69, 0x17, 0x0b, 0xc5, 0xc3, 0x26, 0xec, 0x41,
	0x2a, 0x68, 0x63, 0x11, 0x2f, 0x05, 0x2f, 0xb1, 0x08, 0x3d, 0x08, 0x12, 0x6f, 0x5e, 0x64, 0x93,
	0x6e, 0xdb, 0xc5, 0x0d, 0x1b, 0x36, 0x5b, 0x31, 0x6f, 0xe1, 0x63, 0xf8, 0x28, 0x3d, 0xf6, 0xe8,
	0x29, 0x68, 0xfa, 0x06, 0x79, 0x02, 0x31, 0xbb, 0xb5, 0xbd, 0xcd, 0x7c, 0xf3, 0x9b, 0x99, 0x0f,
	0x3e, 0x80, 0x23, 0x11, 0x4f, 0xe8, 0x9b, 0x27, 0x24, 0x89, 0x38, 0xf5, 0x5e, 0x07, 0x21, 0x55,
	0x64, 0x60, 0xda, 0x7e, 0x22, 0x85, 0x12, 0xb0, 0xa3, 0x99, 0xbe, 0x11, 0x0d, 0x73, 0x72, 0x3c,
	0x13, 0x33, 0x51, 0x11, 0xde, 0x6f, 0xa5, 0x61, 0x2c, 0x41, 0xfd, 0x9e, 0xc8, 0x17, 0xaa, 0xe0,
	0x19, 0xa8, 0xa7, 0x59, 0x1c, 0x0a, 0xde, 0xb5, 0x5d, 0xbb, 0xd7, 0xf4, 0xdb, 0x65, 0xee, 0x1c,
	0x66, 0x24, 0xe6, 0x43, 0xac, 0x75, 0x1c, 0x18, 0x00, 0xde, 0x80, 0x66, 0x1a, 0x49, 0x96, 0xa8,
	0x67, 0x36, 0xe9, 0xfe, 0x73, 0xed, 0x5e, 0xcd, 0x77, 0x8b, 0xdc, 0x69, 0x3c, 0x56, 0xe2, 0x78,
	0x54, 0xe6, 0x4e, 0xcb, 0x6c, 0x6e, 0x30, 0x1c, 0x34, 0x74, 0x3d, 0x9e, 0xe0, 0x0c, 0xc0, 0x3b,
	0xaa, 0xa2, 0xf9, 0x83, 0x64, 0x11, 0xbd, 0x25, 0x9c, 0x8f, 0x88, 0x22, 0xf0, 0x1c, 0xec, 0xeb,
	0xf3, 0x69, 0xd7, 0x76, 0xff, 0xf7, 0x9a, 0x3e, 0x2c, 0x73, 0xe7, 0x68, 0xd7, 0x40, 0x8a, 0x83,
	0x0d, 0x02, 0xaf, 0x01, 0x88, 0x17, 0x5c, 0xb1, 0x84, 0x33, 0x2a, 0x8d, 0x87, 0x4e, 0x99, 0x3b,
	0x6d, 0xbd, 0xb0, 0x9d, 0xe1, 0x60, 0x07, 0xc4, 0x43, 0xd0, 0xda, 0xbe, 0x0e, 0x68, 0xba, 0xe0,
	0x0a, 0x9e, 0x82, 0x3d, 0x49, 0x14, 0xd5, 0x6f, 0x6b, 0x7e, 0xab, 0xcc, 0x9d, 0x03, 0x7d, 0xa5,
	0x92, 0x71, 0xa0, 0xc7, 0x7e, 0xb0, 0xfc, 0x46, 0xd6, 0x47, 0x81, 0xac, 0x65, 0x81, 0xec, 0x55,
	0x81, 0xec, 0xaf, 0x02, 0xd9, 0xef, 0x6b, 0x64, 0xad, 0xd6, 0xc8, 0xfa, 0x5c, 0x23, 0xeb, 0xe9,
	0x72, 0xc6, 0xd4, 0x7c, 0x11, 0xf6, 0x23, 0x11, 0x7b, 0x3a, 0x84, 0x0b, 0x31, 0x9d, 0xb2, 0x88,
	0x11, 0x6e, 0x7a, 0xef, 0x2f, 0x3a, 0x95, 0x25, 0x34, 0x0d, 0xeb, 0x55, 0x0a, 0x57, 0x3f, 0x01,
	0x00, 0x00, 0xff, 0xff, 0xc4, 0xe1, 0x8d, 0xa3, 0xd8, 0x01, 0x00, 0x00,
}

func (m *Market) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Market) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Market) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.ScriptID != 0 {
		i = encodeVarintOracle(dAtA, i, uint64(m.ScriptID))
		i--
		dAtA[i] = 0x10
	}
	if len(m.Symbol) > 0 {
		i -= len(m.Symbol)
		copy(dAtA[i:], m.Symbol)
		i = encodeVarintOracle(dAtA, i, uint64(len(m.Symbol)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *FetchPriceCallData) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *FetchPriceCallData) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *FetchPriceCallData) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.Multiplier != 0 {
		i = encodeVarintOracle(dAtA, i, uint64(m.Multiplier))
		i--
		dAtA[i] = 0x10
	}
	if len(m.Symbols) > 0 {
		for iNdEx := len(m.Symbols) - 1; iNdEx >= 0; iNdEx-- {
			i -= len(m.Symbols[iNdEx])
			copy(dAtA[i:], m.Symbols[iNdEx])
			i = encodeVarintOracle(dAtA, i, uint64(len(m.Symbols[iNdEx])))
			i--
			dAtA[i] = 0xa
		}
	}
	return len(dAtA) - i, nil
}

func (m *FetchPriceResult) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *FetchPriceResult) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *FetchPriceResult) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Rates) > 0 {
		dAtA2 := make([]byte, len(m.Rates)*10)
		var j1 int
		for _, num := range m.Rates {
			for num >= 1<<7 {
				dAtA2[j1] = uint8(uint64(num)&0x7f | 0x80)
				num >>= 7
				j1++
			}
			dAtA2[j1] = uint8(num)
			j1++
		}
		i -= j1
		copy(dAtA[i:], dAtA2[:j1])
		i = encodeVarintOracle(dAtA, i, uint64(j1))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintOracle(dAtA []byte, offset int, v uint64) int {
	offset -= sovOracle(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *Market) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Symbol)
	if l > 0 {
		n += 1 + l + sovOracle(uint64(l))
	}
	if m.ScriptID != 0 {
		n += 1 + sovOracle(uint64(m.ScriptID))
	}
	return n
}

func (m *FetchPriceCallData) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if len(m.Symbols) > 0 {
		for _, s := range m.Symbols {
			l = len(s)
			n += 1 + l + sovOracle(uint64(l))
		}
	}
	if m.Multiplier != 0 {
		n += 1 + sovOracle(uint64(m.Multiplier))
	}
	return n
}

func (m *FetchPriceResult) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if len(m.Rates) > 0 {
		l = 0
		for _, e := range m.Rates {
			l += sovOracle(uint64(e))
		}
		n += 1 + sovOracle(uint64(l)) + l
	}
	return n
}

func sovOracle(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozOracle(x uint64) (n int) {
	return sovOracle(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *Market) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowOracle
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
			return fmt.Errorf("proto: Market: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Market: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Symbol", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowOracle
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
				return ErrInvalidLengthOracle
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthOracle
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Symbol = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field ScriptID", wireType)
			}
			m.ScriptID = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowOracle
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.ScriptID |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipOracle(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthOracle
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
func (m *FetchPriceCallData) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowOracle
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
			return fmt.Errorf("proto: FetchPriceCallData: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: FetchPriceCallData: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Symbols", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowOracle
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
				return ErrInvalidLengthOracle
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthOracle
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Symbols = append(m.Symbols, string(dAtA[iNdEx:postIndex]))
			iNdEx = postIndex
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Multiplier", wireType)
			}
			m.Multiplier = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowOracle
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Multiplier |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipOracle(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthOracle
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
func (m *FetchPriceResult) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowOracle
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
			return fmt.Errorf("proto: FetchPriceResult: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: FetchPriceResult: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType == 0 {
				var v uint64
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return ErrIntOverflowOracle
					}
					if iNdEx >= l {
						return io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					v |= uint64(b&0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				m.Rates = append(m.Rates, v)
			} else if wireType == 2 {
				var packedLen int
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return ErrIntOverflowOracle
					}
					if iNdEx >= l {
						return io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					packedLen |= int(b&0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				if packedLen < 0 {
					return ErrInvalidLengthOracle
				}
				postIndex := iNdEx + packedLen
				if postIndex < 0 {
					return ErrInvalidLengthOracle
				}
				if postIndex > l {
					return io.ErrUnexpectedEOF
				}
				var elementCount int
				var count int
				for _, integer := range dAtA[iNdEx:postIndex] {
					if integer < 128 {
						count++
					}
				}
				elementCount = count
				if elementCount != 0 && len(m.Rates) == 0 {
					m.Rates = make([]uint64, 0, elementCount)
				}
				for iNdEx < postIndex {
					var v uint64
					for shift := uint(0); ; shift += 7 {
						if shift >= 64 {
							return ErrIntOverflowOracle
						}
						if iNdEx >= l {
							return io.ErrUnexpectedEOF
						}
						b := dAtA[iNdEx]
						iNdEx++
						v |= uint64(b&0x7F) << shift
						if b < 0x80 {
							break
						}
					}
					m.Rates = append(m.Rates, v)
				}
			} else {
				return fmt.Errorf("proto: wrong wireType = %d for field Rates", wireType)
			}
		default:
			iNdEx = preIndex
			skippy, err := skipOracle(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthOracle
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
func skipOracle(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowOracle
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
					return 0, ErrIntOverflowOracle
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
					return 0, ErrIntOverflowOracle
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
				return 0, ErrInvalidLengthOracle
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupOracle
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthOracle
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthOracle        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowOracle          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupOracle = fmt.Errorf("proto: unexpected end of group")
)
