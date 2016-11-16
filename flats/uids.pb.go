// Code generated by protoc-gen-gogo.
// source: uids.proto
// DO NOT EDIT!

/*
Package flats is a generated protocol buffer package.

It is generated from these files:
	uids.proto

It has these top-level messages:
	UidList
*/
package flats

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

import io "io"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type UidList struct {
	Uid []uint64 `protobuf:"varint,1,rep,name=uid" json:"uid,omitempty"`
}

func (m *UidList) Reset()                    { *m = UidList{} }
func (m *UidList) String() string            { return proto.CompactTextString(m) }
func (*UidList) ProtoMessage()               {}
func (*UidList) Descriptor() ([]byte, []int) { return fileDescriptorUids, []int{0} }

func init() {
	proto.RegisterType((*UidList)(nil), "flats.UidList")
}
func (m *UidList) Marshal() (data []byte, err error) {
	size := m.Size()
	data = make([]byte, size)
	n, err := m.MarshalTo(data)
	if err != nil {
		return nil, err
	}
	return data[:n], nil
}

func (m *UidList) MarshalTo(data []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if len(m.Uid) > 0 {
		for _, num := range m.Uid {
			data[i] = 0x8
			i++
			i = encodeVarintUids(data, i, uint64(num))
		}
	}
	return i, nil
}

func encodeFixed64Uids(data []byte, offset int, v uint64) int {
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
func encodeFixed32Uids(data []byte, offset int, v uint32) int {
	data[offset] = uint8(v)
	data[offset+1] = uint8(v >> 8)
	data[offset+2] = uint8(v >> 16)
	data[offset+3] = uint8(v >> 24)
	return offset + 4
}
func encodeVarintUids(data []byte, offset int, v uint64) int {
	for v >= 1<<7 {
		data[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	data[offset] = uint8(v)
	return offset + 1
}
func (m *UidList) Size() (n int) {
	var l int
	_ = l
	if len(m.Uid) > 0 {
		for _, e := range m.Uid {
			n += 1 + sovUids(uint64(e))
		}
	}
	return n
}

func sovUids(x uint64) (n int) {
	for {
		n++
		x >>= 7
		if x == 0 {
			break
		}
	}
	return n
}
func sozUids(x uint64) (n int) {
	return sovUids(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *UidList) Unmarshal(data []byte) error {
	l := len(data)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowUids
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
			return fmt.Errorf("proto: UidList: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: UidList: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Uid", wireType)
			}
			var v uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowUids
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := data[iNdEx]
				iNdEx++
				v |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			m.Uid = append(m.Uid, v)
		default:
			iNdEx = preIndex
			skippy, err := skipUids(data[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthUids
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
func skipUids(data []byte) (n int, err error) {
	l := len(data)
	iNdEx := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowUids
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
					return 0, ErrIntOverflowUids
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
					return 0, ErrIntOverflowUids
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
				return 0, ErrInvalidLengthUids
			}
			return iNdEx, nil
		case 3:
			for {
				var innerWire uint64
				var start int = iNdEx
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return 0, ErrIntOverflowUids
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
				next, err := skipUids(data[start:])
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
	ErrInvalidLengthUids = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowUids   = fmt.Errorf("proto: integer overflow")
)

func init() { proto.RegisterFile("uids.proto", fileDescriptorUids) }

var fileDescriptorUids = []byte{
	// 94 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0xe2, 0xe2, 0x2a, 0xcd, 0x4c, 0x29,
	0xd6, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x62, 0x4d, 0xcb, 0x49, 0x2c, 0x29, 0x56, 0x92, 0xe6,
	0x62, 0x0f, 0xcd, 0x4c, 0xf1, 0xc9, 0x2c, 0x2e, 0x11, 0x12, 0xe0, 0x62, 0x2e, 0xcd, 0x4c, 0x91,
	0x60, 0x54, 0x60, 0xd6, 0x60, 0x09, 0x02, 0x31, 0x9d, 0x04, 0x4e, 0x3c, 0x92, 0x63, 0xbc, 0xf0,
	0x48, 0x8e, 0xf1, 0xc1, 0x23, 0x39, 0xc6, 0x19, 0x8f, 0xe5, 0x18, 0x92, 0xd8, 0xc0, 0x9a, 0x8d,
	0x01, 0x01, 0x00, 0x00, 0xff, 0xff, 0x0b, 0xb8, 0xe7, 0x71, 0x4a, 0x00, 0x00, 0x00,
}
