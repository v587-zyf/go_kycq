// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: vip.proto

package pb

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

import io "io"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// vip礼包领取
type VipGiftGetReq struct {
	Lv int32 `protobuf:"varint,1,opt,name=lv,proto3" json:"lv,omitempty"`
}

func (m *VipGiftGetReq) Reset()                    { *m = VipGiftGetReq{} }
func (m *VipGiftGetReq) String() string            { return proto.CompactTextString(m) }
func (*VipGiftGetReq) ProtoMessage()               {}
func (*VipGiftGetReq) Descriptor() ([]byte, []int) { return fileDescriptorVip, []int{0} }

func (m *VipGiftGetReq) GetLv() int32 {
	if m != nil {
		return m.Lv
	}
	return 0
}

type VipGiftGetAck struct {
	Lv int32 `protobuf:"varint,1,opt,name=lv,proto3" json:"lv,omitempty"`
}

func (m *VipGiftGetAck) Reset()                    { *m = VipGiftGetAck{} }
func (m *VipGiftGetAck) String() string            { return proto.CompactTextString(m) }
func (*VipGiftGetAck) ProtoMessage()               {}
func (*VipGiftGetAck) Descriptor() ([]byte, []int) { return fileDescriptorVip, []int{1} }

func (m *VipGiftGetAck) GetLv() int32 {
	if m != nil {
		return m.Lv
	}
	return 0
}

func init() {
	proto.RegisterType((*VipGiftGetReq)(nil), "pb.VipGiftGetReq")
	proto.RegisterType((*VipGiftGetAck)(nil), "pb.VipGiftGetAck")
}
func (m *VipGiftGetReq) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *VipGiftGetReq) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if m.Lv != 0 {
		dAtA[i] = 0x8
		i++
		i = encodeVarintVip(dAtA, i, uint64(m.Lv))
	}
	return i, nil
}

func (m *VipGiftGetAck) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *VipGiftGetAck) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if m.Lv != 0 {
		dAtA[i] = 0x8
		i++
		i = encodeVarintVip(dAtA, i, uint64(m.Lv))
	}
	return i, nil
}

func encodeVarintVip(dAtA []byte, offset int, v uint64) int {
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return offset + 1
}
func (m *VipGiftGetReq) Size() (n int) {
	var l int
	_ = l
	if m.Lv != 0 {
		n += 1 + sovVip(uint64(m.Lv))
	}
	return n
}

func (m *VipGiftGetAck) Size() (n int) {
	var l int
	_ = l
	if m.Lv != 0 {
		n += 1 + sovVip(uint64(m.Lv))
	}
	return n
}

func sovVip(x uint64) (n int) {
	for {
		n++
		x >>= 7
		if x == 0 {
			break
		}
	}
	return n
}
func sozVip(x uint64) (n int) {
	return sovVip(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *VipGiftGetReq) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowVip
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: VipGiftGetReq: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: VipGiftGetReq: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Lv", wireType)
			}
			m.Lv = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowVip
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Lv |= (int32(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipVip(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthVip
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
func (m *VipGiftGetAck) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowVip
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: VipGiftGetAck: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: VipGiftGetAck: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Lv", wireType)
			}
			m.Lv = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowVip
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Lv |= (int32(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipVip(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthVip
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
func skipVip(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowVip
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
					return 0, ErrIntOverflowVip
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				iNdEx++
				if dAtA[iNdEx-1] < 0x80 {
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
					return 0, ErrIntOverflowVip
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
			iNdEx += length
			if length < 0 {
				return 0, ErrInvalidLengthVip
			}
			return iNdEx, nil
		case 3:
			for {
				var innerWire uint64
				var start int = iNdEx
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return 0, ErrIntOverflowVip
					}
					if iNdEx >= l {
						return 0, io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
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
				next, err := skipVip(dAtA[start:])
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
	ErrInvalidLengthVip = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowVip   = fmt.Errorf("proto: integer overflow")
)

func init() { proto.RegisterFile("vip.proto", fileDescriptorVip) }

var fileDescriptorVip = []byte{
	// 103 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0x2c, 0xcb, 0x2c, 0xd0,
	0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x62, 0x2a, 0x48, 0x52, 0x92, 0xe7, 0xe2, 0x0d, 0xcb, 0x2c,
	0x70, 0xcf, 0x4c, 0x2b, 0x71, 0x4f, 0x2d, 0x09, 0x4a, 0x2d, 0x14, 0xe2, 0xe3, 0x62, 0xca, 0x29,
	0x93, 0x60, 0x54, 0x60, 0xd4, 0x60, 0x0d, 0x62, 0xca, 0x29, 0x43, 0x55, 0xe0, 0x98, 0x9c, 0x8d,
	0xae, 0xc0, 0x49, 0xe0, 0xc4, 0x23, 0x39, 0xc6, 0x0b, 0x8f, 0xe4, 0x18, 0x1f, 0x3c, 0x92, 0x63,
	0x9c, 0xf1, 0x58, 0x8e, 0x21, 0x89, 0x0d, 0x6c, 0xbc, 0x31, 0x20, 0x00, 0x00, 0xff, 0xff, 0xad,
	0x3d, 0xf6, 0x50, 0x6b, 0x00, 0x00, 0x00,
}
