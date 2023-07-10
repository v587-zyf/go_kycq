// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: firstRecharge.proto

package pb

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

import io "io"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// 领取礼包
type FirstRechargeRewardReq struct {
	Day int32 `protobuf:"varint,1,opt,name=day,proto3" json:"day,omitempty"`
}

func (m *FirstRechargeRewardReq) Reset()         { *m = FirstRechargeRewardReq{} }
func (m *FirstRechargeRewardReq) String() string { return proto.CompactTextString(m) }
func (*FirstRechargeRewardReq) ProtoMessage()    {}
func (*FirstRechargeRewardReq) Descriptor() ([]byte, []int) {
	return fileDescriptorFirstRecharge, []int{0}
}

func (m *FirstRechargeRewardReq) GetDay() int32 {
	if m != nil {
		return m.Day
	}
	return 0
}

type FirstRechargeRewardAck struct {
	Day   int32           `protobuf:"varint,1,opt,name=day,proto3" json:"day,omitempty"`
	Goods *GoodsChangeNtf `protobuf:"bytes,2,opt,name=goods" json:"goods,omitempty"`
}

func (m *FirstRechargeRewardAck) Reset()         { *m = FirstRechargeRewardAck{} }
func (m *FirstRechargeRewardAck) String() string { return proto.CompactTextString(m) }
func (*FirstRechargeRewardAck) ProtoMessage()    {}
func (*FirstRechargeRewardAck) Descriptor() ([]byte, []int) {
	return fileDescriptorFirstRecharge, []int{1}
}

func (m *FirstRechargeRewardAck) GetDay() int32 {
	if m != nil {
		return m.Day
	}
	return 0
}

func (m *FirstRechargeRewardAck) GetGoods() *GoodsChangeNtf {
	if m != nil {
		return m.Goods
	}
	return nil
}

// 推送已首充
type FirstRechargeNtf struct {
	IsRecharge bool  `protobuf:"varint,1,opt,name=isRecharge,proto3" json:"isRecharge,omitempty"`
	OpenDay    int64 `protobuf:"varint,2,opt,name=openDay,proto3" json:"openDay,omitempty"`
}

func (m *FirstRechargeNtf) Reset()                    { *m = FirstRechargeNtf{} }
func (m *FirstRechargeNtf) String() string            { return proto.CompactTextString(m) }
func (*FirstRechargeNtf) ProtoMessage()               {}
func (*FirstRechargeNtf) Descriptor() ([]byte, []int) { return fileDescriptorFirstRecharge, []int{2} }

func (m *FirstRechargeNtf) GetIsRecharge() bool {
	if m != nil {
		return m.IsRecharge
	}
	return false
}

func (m *FirstRechargeNtf) GetOpenDay() int64 {
	if m != nil {
		return m.OpenDay
	}
	return 0
}

func init() {
	proto.RegisterType((*FirstRechargeRewardReq)(nil), "pb.FirstRechargeRewardReq")
	proto.RegisterType((*FirstRechargeRewardAck)(nil), "pb.FirstRechargeRewardAck")
	proto.RegisterType((*FirstRechargeNtf)(nil), "pb.FirstRechargeNtf")
}
func (m *FirstRechargeRewardReq) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *FirstRechargeRewardReq) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if m.Day != 0 {
		dAtA[i] = 0x8
		i++
		i = encodeVarintFirstRecharge(dAtA, i, uint64(m.Day))
	}
	return i, nil
}

func (m *FirstRechargeRewardAck) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *FirstRechargeRewardAck) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if m.Day != 0 {
		dAtA[i] = 0x8
		i++
		i = encodeVarintFirstRecharge(dAtA, i, uint64(m.Day))
	}
	if m.Goods != nil {
		dAtA[i] = 0x12
		i++
		i = encodeVarintFirstRecharge(dAtA, i, uint64(m.Goods.Size()))
		n1, err := m.Goods.MarshalTo(dAtA[i:])
		if err != nil {
			return 0, err
		}
		i += n1
	}
	return i, nil
}

func (m *FirstRechargeNtf) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *FirstRechargeNtf) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if m.IsRecharge {
		dAtA[i] = 0x8
		i++
		if m.IsRecharge {
			dAtA[i] = 1
		} else {
			dAtA[i] = 0
		}
		i++
	}
	if m.OpenDay != 0 {
		dAtA[i] = 0x10
		i++
		i = encodeVarintFirstRecharge(dAtA, i, uint64(m.OpenDay))
	}
	return i, nil
}

func encodeVarintFirstRecharge(dAtA []byte, offset int, v uint64) int {
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return offset + 1
}
func (m *FirstRechargeRewardReq) Size() (n int) {
	var l int
	_ = l
	if m.Day != 0 {
		n += 1 + sovFirstRecharge(uint64(m.Day))
	}
	return n
}

func (m *FirstRechargeRewardAck) Size() (n int) {
	var l int
	_ = l
	if m.Day != 0 {
		n += 1 + sovFirstRecharge(uint64(m.Day))
	}
	if m.Goods != nil {
		l = m.Goods.Size()
		n += 1 + l + sovFirstRecharge(uint64(l))
	}
	return n
}

func (m *FirstRechargeNtf) Size() (n int) {
	var l int
	_ = l
	if m.IsRecharge {
		n += 2
	}
	if m.OpenDay != 0 {
		n += 1 + sovFirstRecharge(uint64(m.OpenDay))
	}
	return n
}

func sovFirstRecharge(x uint64) (n int) {
	for {
		n++
		x >>= 7
		if x == 0 {
			break
		}
	}
	return n
}
func sozFirstRecharge(x uint64) (n int) {
	return sovFirstRecharge(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *FirstRechargeRewardReq) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowFirstRecharge
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
			return fmt.Errorf("proto: FirstRechargeRewardReq: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: FirstRechargeRewardReq: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Day", wireType)
			}
			m.Day = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowFirstRecharge
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Day |= (int32(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipFirstRecharge(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthFirstRecharge
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
func (m *FirstRechargeRewardAck) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowFirstRecharge
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
			return fmt.Errorf("proto: FirstRechargeRewardAck: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: FirstRechargeRewardAck: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Day", wireType)
			}
			m.Day = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowFirstRecharge
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Day |= (int32(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Goods", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowFirstRecharge
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthFirstRecharge
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Goods == nil {
				m.Goods = &GoodsChangeNtf{}
			}
			if err := m.Goods.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipFirstRecharge(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthFirstRecharge
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
func (m *FirstRechargeNtf) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowFirstRecharge
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
			return fmt.Errorf("proto: FirstRechargeNtf: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: FirstRechargeNtf: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field IsRecharge", wireType)
			}
			var v int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowFirstRecharge
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				v |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			m.IsRecharge = bool(v != 0)
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field OpenDay", wireType)
			}
			m.OpenDay = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowFirstRecharge
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.OpenDay |= (int64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipFirstRecharge(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthFirstRecharge
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
func skipFirstRecharge(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowFirstRecharge
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
					return 0, ErrIntOverflowFirstRecharge
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
					return 0, ErrIntOverflowFirstRecharge
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
				return 0, ErrInvalidLengthFirstRecharge
			}
			return iNdEx, nil
		case 3:
			for {
				var innerWire uint64
				var start int = iNdEx
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return 0, ErrIntOverflowFirstRecharge
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
				next, err := skipFirstRecharge(dAtA[start:])
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
	ErrInvalidLengthFirstRecharge = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowFirstRecharge   = fmt.Errorf("proto: integer overflow")
)

func init() { proto.RegisterFile("firstRecharge.proto", fileDescriptorFirstRecharge) }

var fileDescriptorFirstRecharge = []byte{
	// 200 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x12, 0x4e, 0xcb, 0x2c, 0x2a,
	0x2e, 0x09, 0x4a, 0x4d, 0xce, 0x48, 0x2c, 0x4a, 0x4f, 0xd5, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17,
	0x62, 0x2a, 0x48, 0x92, 0xe2, 0x49, 0xce, 0xcf, 0xcd, 0xcd, 0xcf, 0x83, 0x88, 0x28, 0x69, 0x71,
	0x89, 0xb9, 0x21, 0x2b, 0x0c, 0x4a, 0x2d, 0x4f, 0x2c, 0x4a, 0x09, 0x4a, 0x2d, 0x14, 0x12, 0xe0,
	0x62, 0x4e, 0x49, 0xac, 0x94, 0x60, 0x54, 0x60, 0xd4, 0x60, 0x0d, 0x02, 0x31, 0x95, 0x42, 0xb0,
	0xaa, 0x75, 0x4c, 0xce, 0xc6, 0x54, 0x2b, 0xa4, 0xc1, 0xc5, 0x9a, 0x9e, 0x9f, 0x9f, 0x52, 0x2c,
	0xc1, 0xa4, 0xc0, 0xa8, 0xc1, 0x6d, 0x24, 0xa4, 0x57, 0x90, 0xa4, 0xe7, 0x0e, 0x12, 0x70, 0xce,
	0x48, 0xcc, 0x4b, 0x4f, 0xf5, 0x2b, 0x49, 0x0b, 0x82, 0x28, 0x50, 0xf2, 0xe1, 0x12, 0x40, 0x31,
	0xd5, 0xaf, 0x24, 0x4d, 0x48, 0x8e, 0x8b, 0x2b, 0xb3, 0x18, 0x26, 0x00, 0x36, 0x96, 0x23, 0x08,
	0x49, 0x44, 0x48, 0x82, 0x8b, 0x3d, 0xbf, 0x20, 0x35, 0xcf, 0x25, 0xb1, 0x12, 0x6c, 0x3e, 0x73,
	0x10, 0x8c, 0xeb, 0x24, 0x70, 0xe2, 0x91, 0x1c, 0xe3, 0x85, 0x47, 0x72, 0x8c, 0x0f, 0x1e, 0xc9,
	0x31, 0xce, 0x78, 0x2c, 0xc7, 0x90, 0xc4, 0x06, 0xf6, 0xa8, 0x31, 0x20, 0x00, 0x00, 0xff, 0xff,
	0xf0, 0x0c, 0x83, 0x37, 0x11, 0x01, 0x00, 0x00,
}
