// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: lib/types/header.proto

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

type HeaderPb struct {
	ParentHash []byte `protobuf:"bytes,1,opt,name=parentHash,proto3" json:"parenthash" xml:"parenthash"`
	Version    uint64 `protobuf:"varint,2,opt,name=version,proto3" json:"version" xml:"version"`
	Timestamp  uint64 `protobuf:"varint,3,opt,name=timestamp,proto3" json:"timestamp" xml:"timestamp"`
	Extra      []byte `protobuf:"bytes,4,opt,name=extra,proto3" json:"extra" xml:"extra"`
}

func (m *HeaderPb) Reset()         { *m = HeaderPb{} }
func (m *HeaderPb) String() string { return proto.CompactTextString(m) }
func (*HeaderPb) ProtoMessage()    {}
func (*HeaderPb) Descriptor() ([]byte, []int) {
	return fileDescriptor_87736b9d033c718d, []int{0}
}
func (m *HeaderPb) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *HeaderPb) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_HeaderPb.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *HeaderPb) XXX_Merge(src proto.Message) {
	xxx_messageInfo_HeaderPb.Merge(m, src)
}
func (m *HeaderPb) XXX_Size() int {
	return m.ProtoSize()
}
func (m *HeaderPb) XXX_DiscardUnknown() {
	xxx_messageInfo_HeaderPb.DiscardUnknown(m)
}

var xxx_messageInfo_HeaderPb proto.InternalMessageInfo

func init() {
	proto.RegisterType((*HeaderPb)(nil), "types.HeaderPb")
}

func init() { proto.RegisterFile("lib/types/header.proto", fileDescriptor_87736b9d033c718d) }

var fileDescriptor_87736b9d033c718d = []byte{
	// 308 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x4c, 0x91, 0xbd, 0x4e, 0xf3, 0x30,
	0x14, 0x86, 0xe3, 0xef, 0x6b, 0xf9, 0x31, 0x20, 0x50, 0x06, 0x14, 0x21, 0x64, 0x57, 0x66, 0xa9,
	0x84, 0xa8, 0x07, 0x24, 0x06, 0x06, 0x86, 0x20, 0x41, 0x17, 0x24, 0x94, 0x91, 0xcd, 0x01, 0xd3,
	0x44, 0x6a, 0xea, 0xc8, 0x76, 0xa1, 0xdc, 0x05, 0x97, 0xc0, 0xe5, 0x74, 0xcc, 0xc8, 0x64, 0xa9,
	0xcd, 0x96, 0x31, 0x03, 0x33, 0xd2, 0x49, 0xdb, 0xb0, 0xf9, 0x7d, 0x5e, 0x3f, 0xd2, 0xd1, 0x39,
	0xf8, 0x78, 0x9c, 0xc6, 0xdc, 0x7e, 0xe4, 0xd2, 0xf0, 0x44, 0x8a, 0x17, 0xa9, 0x07, 0xb9, 0x56,
	0x56, 0xf9, 0x5d, 0x60, 0x27, 0x67, 0x5a, 0xe6, 0xca, 0x70, 0x60, 0xf1, 0xf4, 0x95, 0x8f, 0xd4,
	0x48, 0x41, 0x80, 0x57, 0xf3, 0x97, 0xfd, 0x20, 0xbc, 0x33, 0x04, 0xf9, 0x31, 0xf6, 0x43, 0x8c,
	0x73, 0xa1, 0xe5, 0xc4, 0x0e, 0x85, 0x49, 0x02, 0xd4, 0x43, 0xfd, 0xfd, 0x90, 0x55, 0x8e, 0xae,
	0x68, 0x22, 0x4c, 0x52, 0x3b, 0x7a, 0x34, 0xcb, 0xc6, 0xd7, 0xac, 0x45, 0x2c, 0xfa, 0x63, 0xf9,
	0x57, 0x78, 0xfb, 0x4d, 0x6a, 0x93, 0xaa, 0x49, 0xf0, 0xaf, 0x87, 0xfa, 0x9d, 0xf0, 0xb4, 0x72,
	0x74, 0x8d, 0x6a, 0x47, 0x0f, 0xc0, 0x5e, 0x65, 0x16, 0xad, 0x1b, 0xff, 0x06, 0xef, 0xda, 0x34,
	0x93, 0xc6, 0x8a, 0x2c, 0x0f, 0xfe, 0x83, 0xd9, 0xab, 0x1c, 0x6d, 0x61, 0xed, 0xe8, 0x21, 0xb8,
	0x1b, 0xc2, 0xa2, 0xb6, 0xf5, 0x07, 0xb8, 0x2b, 0x67, 0x56, 0x8b, 0xa0, 0x03, 0x63, 0x07, 0x95,
	0xa3, 0x0d, 0xa8, 0x1d, 0xdd, 0x03, 0x0f, 0x12, 0x8b, 0x1a, 0x1a, 0x3e, 0xcc, 0x17, 0xc4, 0x2b,
	0x16, 0xc4, 0x9b, 0x2f, 0x09, 0x2a, 0x96, 0x04, 0x7d, 0x96, 0xc4, 0xfb, 0x2a, 0x09, 0x2a, 0x4a,
	0xe2, 0x7d, 0x97, 0xc4, 0x7b, 0x3a, 0x1f, 0xa5, 0x36, 0x99, 0xc6, 0x83, 0x67, 0x95, 0xf1, 0x48,
	0x19, 0x69, 0xad, 0xb8, 0x1b, 0xab, 0x77, 0x7e, 0x2b, 0xb4, 0x4e, 0xa5, 0xbe, 0xb8, 0x57, 0x7c,
	0x73, 0x80, 0x78, 0x0b, 0xd6, 0x79, 0xf9, 0x1b, 0x00, 0x00, 0xff, 0xff, 0x4b, 0xc4, 0xbc, 0x20,
	0x94, 0x01, 0x00, 0x00,
}

func (m *HeaderPb) Marshal() (dAtA []byte, err error) {
	size := m.ProtoSize()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *HeaderPb) MarshalTo(dAtA []byte) (int, error) {
	size := m.ProtoSize()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *HeaderPb) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Extra) > 0 {
		i -= len(m.Extra)
		copy(dAtA[i:], m.Extra)
		i = encodeVarintHeader(dAtA, i, uint64(len(m.Extra)))
		i--
		dAtA[i] = 0x22
	}
	if m.Timestamp != 0 {
		i = encodeVarintHeader(dAtA, i, uint64(m.Timestamp))
		i--
		dAtA[i] = 0x18
	}
	if m.Version != 0 {
		i = encodeVarintHeader(dAtA, i, uint64(m.Version))
		i--
		dAtA[i] = 0x10
	}
	if len(m.ParentHash) > 0 {
		i -= len(m.ParentHash)
		copy(dAtA[i:], m.ParentHash)
		i = encodeVarintHeader(dAtA, i, uint64(len(m.ParentHash)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintHeader(dAtA []byte, offset int, v uint64) int {
	offset -= sovHeader(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *HeaderPb) ProtoSize() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.ParentHash)
	if l > 0 {
		n += 1 + l + sovHeader(uint64(l))
	}
	if m.Version != 0 {
		n += 1 + sovHeader(uint64(m.Version))
	}
	if m.Timestamp != 0 {
		n += 1 + sovHeader(uint64(m.Timestamp))
	}
	l = len(m.Extra)
	if l > 0 {
		n += 1 + l + sovHeader(uint64(l))
	}
	return n
}

func sovHeader(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozHeader(x uint64) (n int) {
	return sovHeader(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *HeaderPb) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowHeader
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
			return fmt.Errorf("proto: HeaderPb: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: HeaderPb: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ParentHash", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowHeader
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				byteLen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if byteLen < 0 {
				return ErrInvalidLengthHeader
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthHeader
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ParentHash = append(m.ParentHash[:0], dAtA[iNdEx:postIndex]...)
			if m.ParentHash == nil {
				m.ParentHash = []byte{}
			}
			iNdEx = postIndex
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Version", wireType)
			}
			m.Version = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowHeader
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Version |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Timestamp", wireType)
			}
			m.Timestamp = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowHeader
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Timestamp |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Extra", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowHeader
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				byteLen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if byteLen < 0 {
				return ErrInvalidLengthHeader
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthHeader
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Extra = append(m.Extra[:0], dAtA[iNdEx:postIndex]...)
			if m.Extra == nil {
				m.Extra = []byte{}
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipHeader(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthHeader
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
func skipHeader(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowHeader
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
					return 0, ErrIntOverflowHeader
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
					return 0, ErrIntOverflowHeader
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
				return 0, ErrInvalidLengthHeader
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupHeader
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthHeader
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthHeader        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowHeader          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupHeader = fmt.Errorf("proto: unexpected end of group")
)
