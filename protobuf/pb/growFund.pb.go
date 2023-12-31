// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: growFund.proto

package pb

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

import io "io"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// 成长基金购买
type GrowFundBuyReq struct {
}

func (m *GrowFundBuyReq) Reset()                    { *m = GrowFundBuyReq{} }
func (m *GrowFundBuyReq) String() string            { return proto.CompactTextString(m) }
func (*GrowFundBuyReq) ProtoMessage()               {}
func (*GrowFundBuyReq) Descriptor() ([]byte, []int) { return fileDescriptorGrowFund, []int{0} }

type GrowFundBuyAck struct {
	IsBuy bool `protobuf:"varint,1,opt,name=isBuy,proto3" json:"isBuy,omitempty"`
}

func (m *GrowFundBuyAck) Reset()                    { *m = GrowFundBuyAck{} }
func (m *GrowFundBuyAck) String() string            { return proto.CompactTextString(m) }
func (*GrowFundBuyAck) ProtoMessage()               {}
func (*GrowFundBuyAck) Descriptor() ([]byte, []int) { return fileDescriptorGrowFund, []int{1} }

func (m *GrowFundBuyAck) GetIsBuy() bool {
	if m != nil {
		return m.IsBuy
	}
	return false
}

// 成长基金领取
type GrowFundRewardReq struct {
	Id int32 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (m *GrowFundRewardReq) Reset()                    { *m = GrowFundRewardReq{} }
func (m *GrowFundRewardReq) String() string            { return proto.CompactTextString(m) }
func (*GrowFundRewardReq) ProtoMessage()               {}
func (*GrowFundRewardReq) Descriptor() ([]byte, []int) { return fileDescriptorGrowFund, []int{2} }

func (m *GrowFundRewardReq) GetId() int32 {
	if m != nil {
		return m.Id
	}
	return 0
}

type GrowFundRewardAck struct {
	Id    int32           `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Goods *GoodsChangeNtf `protobuf:"bytes,2,opt,name=goods" json:"goods,omitempty"`
}

func (m *GrowFundRewardAck) Reset()                    { *m = GrowFundRewardAck{} }
func (m *GrowFundRewardAck) String() string            { return proto.CompactTextString(m) }
func (*GrowFundRewardAck) ProtoMessage()               {}
func (*GrowFundRewardAck) Descriptor() ([]byte, []int) { return fileDescriptorGrowFund, []int{3} }

func (m *GrowFundRewardAck) GetId() int32 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *GrowFundRewardAck) GetGoods() *GoodsChangeNtf {
	if m != nil {
		return m.Goods
	}
	return nil
}

func init() {
	proto.RegisterType((*GrowFundBuyReq)(nil), "pb.GrowFundBuyReq")
	proto.RegisterType((*GrowFundBuyAck)(nil), "pb.GrowFundBuyAck")
	proto.RegisterType((*GrowFundRewardReq)(nil), "pb.GrowFundRewardReq")
	proto.RegisterType((*GrowFundRewardAck)(nil), "pb.GrowFundRewardAck")
}
func (m *GrowFundBuyReq) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *GrowFundBuyReq) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	return i, nil
}

func (m *GrowFundBuyAck) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *GrowFundBuyAck) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if m.IsBuy {
		dAtA[i] = 0x8
		i++
		if m.IsBuy {
			dAtA[i] = 1
		} else {
			dAtA[i] = 0
		}
		i++
	}
	return i, nil
}

func (m *GrowFundRewardReq) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *GrowFundRewardReq) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if m.Id != 0 {
		dAtA[i] = 0x8
		i++
		i = encodeVarintGrowFund(dAtA, i, uint64(m.Id))
	}
	return i, nil
}

func (m *GrowFundRewardAck) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *GrowFundRewardAck) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if m.Id != 0 {
		dAtA[i] = 0x8
		i++
		i = encodeVarintGrowFund(dAtA, i, uint64(m.Id))
	}
	if m.Goods != nil {
		dAtA[i] = 0x12
		i++
		i = encodeVarintGrowFund(dAtA, i, uint64(m.Goods.Size()))
		n1, err := m.Goods.MarshalTo(dAtA[i:])
		if err != nil {
			return 0, err
		}
		i += n1
	}
	return i, nil
}

func encodeVarintGrowFund(dAtA []byte, offset int, v uint64) int {
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return offset + 1
}
func (m *GrowFundBuyReq) Size() (n int) {
	var l int
	_ = l
	return n
}

func (m *GrowFundBuyAck) Size() (n int) {
	var l int
	_ = l
	if m.IsBuy {
		n += 2
	}
	return n
}

func (m *GrowFundRewardReq) Size() (n int) {
	var l int
	_ = l
	if m.Id != 0 {
		n += 1 + sovGrowFund(uint64(m.Id))
	}
	return n
}

func (m *GrowFundRewardAck) Size() (n int) {
	var l int
	_ = l
	if m.Id != 0 {
		n += 1 + sovGrowFund(uint64(m.Id))
	}
	if m.Goods != nil {
		l = m.Goods.Size()
		n += 1 + l + sovGrowFund(uint64(l))
	}
	return n
}

func sovGrowFund(x uint64) (n int) {
	for {
		n++
		x >>= 7
		if x == 0 {
			break
		}
	}
	return n
}
func sozGrowFund(x uint64) (n int) {
	return sovGrowFund(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *GrowFundBuyReq) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowGrowFund
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
			return fmt.Errorf("proto: GrowFundBuyReq: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: GrowFundBuyReq: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		default:
			iNdEx = preIndex
			skippy, err := skipGrowFund(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthGrowFund
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
func (m *GrowFundBuyAck) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowGrowFund
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
			return fmt.Errorf("proto: GrowFundBuyAck: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: GrowFundBuyAck: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field IsBuy", wireType)
			}
			var v int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGrowFund
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
			m.IsBuy = bool(v != 0)
		default:
			iNdEx = preIndex
			skippy, err := skipGrowFund(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthGrowFund
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
func (m *GrowFundRewardReq) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowGrowFund
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
			return fmt.Errorf("proto: GrowFundRewardReq: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: GrowFundRewardReq: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Id", wireType)
			}
			m.Id = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGrowFund
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Id |= (int32(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipGrowFund(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthGrowFund
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
func (m *GrowFundRewardAck) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowGrowFund
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
			return fmt.Errorf("proto: GrowFundRewardAck: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: GrowFundRewardAck: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Id", wireType)
			}
			m.Id = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGrowFund
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Id |= (int32(b) & 0x7F) << shift
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
					return ErrIntOverflowGrowFund
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
				return ErrInvalidLengthGrowFund
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
			skippy, err := skipGrowFund(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthGrowFund
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
func skipGrowFund(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowGrowFund
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
					return 0, ErrIntOverflowGrowFund
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
					return 0, ErrIntOverflowGrowFund
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
				return 0, ErrInvalidLengthGrowFund
			}
			return iNdEx, nil
		case 3:
			for {
				var innerWire uint64
				var start int = iNdEx
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return 0, ErrIntOverflowGrowFund
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
				next, err := skipGrowFund(dAtA[start:])
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
	ErrInvalidLengthGrowFund = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowGrowFund   = fmt.Errorf("proto: integer overflow")
)

func init() { proto.RegisterFile("growFund.proto", fileDescriptorGrowFund) }

var fileDescriptorGrowFund = []byte{
	// 188 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0x4b, 0x2f, 0xca, 0x2f,
	0x77, 0x2b, 0xcd, 0x4b, 0xd1, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x62, 0x2a, 0x48, 0x92, 0xe2,
	0x49, 0xce, 0xcf, 0xcd, 0xcd, 0xcf, 0x83, 0x88, 0x28, 0x09, 0x70, 0xf1, 0xb9, 0x43, 0xd5, 0x38,
	0x95, 0x56, 0x06, 0xa5, 0x16, 0x2a, 0xa9, 0xa1, 0x88, 0x38, 0x26, 0x67, 0x0b, 0x89, 0x70, 0xb1,
	0x66, 0x16, 0x3b, 0x95, 0x56, 0x4a, 0x30, 0x2a, 0x30, 0x6a, 0x70, 0x04, 0x41, 0x38, 0x4a, 0xca,
	0x5c, 0x82, 0x30, 0x75, 0x41, 0xa9, 0xe5, 0x89, 0x45, 0x29, 0x41, 0xa9, 0x85, 0x42, 0x7c, 0x5c,
	0x4c, 0x99, 0x29, 0x60, 0x75, 0xac, 0x41, 0x4c, 0x99, 0x29, 0x4a, 0xbe, 0xe8, 0x8a, 0x40, 0xe6,
	0xa1, 0x29, 0x12, 0xd2, 0xe0, 0x62, 0x4d, 0xcf, 0xcf, 0x4f, 0x29, 0x96, 0x60, 0x52, 0x60, 0xd4,
	0xe0, 0x36, 0x12, 0xd2, 0x2b, 0x48, 0xd2, 0x73, 0x07, 0x09, 0x38, 0x67, 0x24, 0xe6, 0xa5, 0xa7,
	0xfa, 0x95, 0xa4, 0x05, 0x41, 0x14, 0x38, 0x09, 0x9c, 0x78, 0x24, 0xc7, 0x78, 0xe1, 0x91, 0x1c,
	0xe3, 0x83, 0x47, 0x72, 0x8c, 0x33, 0x1e, 0xcb, 0x31, 0x24, 0xb1, 0x81, 0xbd, 0x61, 0x0c, 0x08,
	0x00, 0x00, 0xff, 0xff, 0x51, 0xfd, 0xb5, 0xe2, 0xea, 0x00, 0x00, 0x00,
}
