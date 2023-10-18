// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: lavanet/lava/subscription/cu_tracker.proto

package types

import (
	fmt "fmt"
	proto "github.com/cosmos/gogoproto/proto"
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

type TrackedCu struct {
	CuWithQos    uint64 `protobuf:"varint,1,opt,name=cu_with_qos,json=cuWithQos,proto3" json:"cu_with_qos,omitempty"`
	CuWithoutQos uint64 `protobuf:"varint,2,opt,name=cu_without_qos,json=cuWithoutQos,proto3" json:"cu_without_qos,omitempty"`
}

func (m *TrackedCu) Reset()         { *m = TrackedCu{} }
func (m *TrackedCu) String() string { return proto.CompactTextString(m) }
func (*TrackedCu) ProtoMessage()    {}
func (*TrackedCu) Descriptor() ([]byte, []int) {
	return fileDescriptor_5974e118ddf7c543, []int{0}
}
func (m *TrackedCu) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *TrackedCu) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_TrackedCu.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *TrackedCu) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TrackedCu.Merge(m, src)
}
func (m *TrackedCu) XXX_Size() int {
	return m.Size()
}
func (m *TrackedCu) XXX_DiscardUnknown() {
	xxx_messageInfo_TrackedCu.DiscardUnknown(m)
}

var xxx_messageInfo_TrackedCu proto.InternalMessageInfo

func (m *TrackedCu) GetCuWithQos() uint64 {
	if m != nil {
		return m.CuWithQos
	}
	return 0
}

func (m *TrackedCu) GetCuWithoutQos() uint64 {
	if m != nil {
		return m.CuWithoutQos
	}
	return 0
}

func init() {
	proto.RegisterType((*TrackedCu)(nil), "lavanet.lava.subscription.TrackedCu")
}

func init() {
	proto.RegisterFile("lavanet/lava/subscription/cu_tracker.proto", fileDescriptor_5974e118ddf7c543)
}

var fileDescriptor_5974e118ddf7c543 = []byte{
	// 197 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xd2, 0xca, 0x49, 0x2c, 0x4b,
	0xcc, 0x4b, 0x2d, 0xd1, 0x07, 0xd1, 0xfa, 0xc5, 0xa5, 0x49, 0xc5, 0xc9, 0x45, 0x99, 0x05, 0x25,
	0x99, 0xf9, 0x79, 0xfa, 0xc9, 0xa5, 0xf1, 0x25, 0x45, 0x89, 0xc9, 0xd9, 0xa9, 0x45, 0x7a, 0x05,
	0x45, 0xf9, 0x25, 0xf9, 0x42, 0x92, 0x50, 0xb5, 0x7a, 0x20, 0x5a, 0x0f, 0x59, 0xad, 0x52, 0x20,
	0x17, 0x67, 0x08, 0x58, 0x6d, 0x8a, 0x73, 0xa9, 0x90, 0x1c, 0x17, 0x77, 0x72, 0x69, 0x7c, 0x79,
	0x66, 0x49, 0x46, 0x7c, 0x61, 0x7e, 0xb1, 0x04, 0xa3, 0x02, 0xa3, 0x06, 0x4b, 0x10, 0x67, 0x72,
	0x69, 0x78, 0x66, 0x49, 0x46, 0x60, 0x7e, 0xb1, 0x90, 0x0a, 0x17, 0x1f, 0x54, 0x3e, 0xbf, 0xb4,
	0x04, 0xac, 0x84, 0x09, 0xac, 0x84, 0x07, 0xa2, 0x24, 0xbf, 0xb4, 0x24, 0x30, 0xbf, 0xd8, 0xc9,
	0xed, 0xc4, 0x23, 0x39, 0xc6, 0x0b, 0x8f, 0xe4, 0x18, 0x1f, 0x3c, 0x92, 0x63, 0x9c, 0xf0, 0x58,
	0x8e, 0xe1, 0xc2, 0x63, 0x39, 0x86, 0x1b, 0x8f, 0xe5, 0x18, 0xa2, 0x74, 0xd2, 0x33, 0x4b, 0x32,
	0x4a, 0x93, 0xf4, 0x92, 0xf3, 0x73, 0xf5, 0x51, 0x9c, 0x5f, 0x81, 0xea, 0x81, 0x92, 0xca, 0x82,
	0xd4, 0xe2, 0x24, 0x36, 0xb0, 0xe3, 0x8d, 0x01, 0x01, 0x00, 0x00, 0xff, 0xff, 0xb3, 0x44, 0xa1,
	0x3a, 0xea, 0x00, 0x00, 0x00,
}

func (m *TrackedCu) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *TrackedCu) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *TrackedCu) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.CuWithoutQos != 0 {
		i = encodeVarintCuTracker(dAtA, i, uint64(m.CuWithoutQos))
		i--
		dAtA[i] = 0x10
	}
	if m.CuWithQos != 0 {
		i = encodeVarintCuTracker(dAtA, i, uint64(m.CuWithQos))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func encodeVarintCuTracker(dAtA []byte, offset int, v uint64) int {
	offset -= sovCuTracker(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *TrackedCu) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.CuWithQos != 0 {
		n += 1 + sovCuTracker(uint64(m.CuWithQos))
	}
	if m.CuWithoutQos != 0 {
		n += 1 + sovCuTracker(uint64(m.CuWithoutQos))
	}
	return n
}

func sovCuTracker(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozCuTracker(x uint64) (n int) {
	return sovCuTracker(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *TrackedCu) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowCuTracker
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
			return fmt.Errorf("proto: TrackedCu: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: TrackedCu: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field CuWithQos", wireType)
			}
			m.CuWithQos = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCuTracker
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.CuWithQos |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field CuWithoutQos", wireType)
			}
			m.CuWithoutQos = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCuTracker
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.CuWithoutQos |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipCuTracker(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthCuTracker
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
func skipCuTracker(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowCuTracker
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
					return 0, ErrIntOverflowCuTracker
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
					return 0, ErrIntOverflowCuTracker
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
				return 0, ErrInvalidLengthCuTracker
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupCuTracker
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthCuTracker
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthCuTracker        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowCuTracker          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupCuTracker = fmt.Errorf("proto: unexpected end of group")
)
