// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: spendRebates.proto

package pb

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

import io "io"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// 消费返利领取
type SpendRebatesRewardReq struct {
	Id int32 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (m *SpendRebatesRewardReq) Reset()         { *m = SpendRebatesRewardReq{} }
func (m *SpendRebatesRewardReq) String() string { return proto.CompactTextString(m) }
func (*SpendRebatesRewardReq) ProtoMessage()    {}
func (*SpendRebatesRewardReq) Descriptor() ([]byte, []int) {
	return fileDescriptorSpendRebates, []int{0}
}

func (m *SpendRebatesRewardReq) GetId() int32 {
	if m != nil {
		return m.Id
	}
	return 0
}

type SpendRebatesRewardAck struct {
	Id    int32           `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Goods *GoodsChangeNtf `protobuf:"bytes,2,opt,name=goods" json:"goods,omitempty"`
}

func (m *SpendRebatesRewardAck) Reset()         { *m = SpendRebatesRewardAck{} }
func (m *SpendRebatesRewardAck) String() string { return proto.CompactTextString(m) }
func (*SpendRebatesRewardAck) ProtoMessage()    {}
func (*SpendRebatesRewardAck) Descriptor() ([]byte, []int) {
	return fileDescriptorSpendRebates, []int{1}
}

func (m *SpendRebatesRewardAck) GetId() int32 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *SpendRebatesRewardAck) GetGoods() *GoodsChangeNtf {
	if m != nil {
		return m.Goods
	}
	return nil
}

// 推送消费额
type SpendRebatesNtf struct {
	CountIngot int32 `protobuf:"varint,1,opt,name=countIngot,proto3" json:"countIngot,omitempty"`
	Ingot      int32 `protobuf:"varint,2,opt,name=ingot,proto3" json:"ingot,omitempty"`
}

func (m *SpendRebatesNtf) Reset()                    { *m = SpendRebatesNtf{} }
func (m *SpendRebatesNtf) String() string            { return proto.CompactTextString(m) }
func (*SpendRebatesNtf) ProtoMessage()               {}
func (*SpendRebatesNtf) Descriptor() ([]byte, []int) { return fileDescriptorSpendRebates, []int{2} }

func (m *SpendRebatesNtf) GetCountIngot() int32 {
	if m != nil {
		return m.CountIngot
	}
	return 0
}

func (m *SpendRebatesNtf) GetIngot() int32 {
	if m != nil {
		return m.Ingot
	}
	return 0
}

func init() {
	proto.RegisterType((*SpendRebatesRewardReq)(nil), "pb.SpendRebatesRewardReq")
	proto.RegisterType((*SpendRebatesRewardAck)(nil), "pb.SpendRebatesRewardAck")
	proto.RegisterType((*SpendRebatesNtf)(nil), "pb.SpendRebatesNtf")
}
func (m *SpendRebatesRewardReq) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *SpendRebatesRewardReq) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if m.Id != 0 {
		dAtA[i] = 0x8
		i++
		i = encodeVarintSpendRebates(dAtA, i, uint64(m.Id))
	}
	return i, nil
}

func (m *SpendRebatesRewardAck) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *SpendRebatesRewardAck) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if m.Id != 0 {
		dAtA[i] = 0x8
		i++
		i = encodeVarintSpendRebates(dAtA, i, uint64(m.Id))
	}
	if m.Goods != nil {
		dAtA[i] = 0x12
		i++
		i = encodeVarintSpendRebates(dAtA, i, uint64(m.Goods.Size()))
		n1, err := m.Goods.MarshalTo(dAtA[i:])
		if err != nil {
			return 0, err
		}
		i += n1
	}
	return i, nil
}

func (m *SpendRebatesNtf) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *SpendRebatesNtf) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if m.CountIngot != 0 {
		dAtA[i] = 0x8
		i++
		i = encodeVarintSpendRebates(dAtA, i, uint64(m.CountIngot))
	}
	if m.Ingot != 0 {
		dAtA[i] = 0x10
		i++
		i = encodeVarintSpendRebates(dAtA, i, uint64(m.Ingot))
	}
	return i, nil
}

func encodeVarintSpendRebates(dAtA []byte, offset int, v uint64) int {
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return offset + 1
}
func (m *SpendRebatesRewardReq) Size() (n int) {
	var l int
	_ = l
	if m.Id != 0 {
		n += 1 + sovSpendRebates(uint64(m.Id))
	}
	return n
}

func (m *SpendRebatesRewardAck) Size() (n int) {
	var l int
	_ = l
	if m.Id != 0 {
		n += 1 + sovSpendRebates(uint64(m.Id))
	}
	if m.Goods != nil {
		l = m.Goods.Size()
		n += 1 + l + sovSpendRebates(uint64(l))
	}
	return n
}

func (m *SpendRebatesNtf) Size() (n int) {
	var l int
	_ = l
	if m.CountIngot != 0 {
		n += 1 + sovSpendRebates(uint64(m.CountIngot))
	}
	if m.Ingot != 0 {
		n += 1 + sovSpendRebates(uint64(m.Ingot))
	}
	return n
}

func sovSpendRebates(x uint64) (n int) {
	for {
		n++
		x >>= 7
		if x == 0 {
			break
		}
	}
	return n
}
func sozSpendRebates(x uint64) (n int) {
	return sovSpendRebates(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *SpendRebatesRewardReq) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowSpendRebates
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
			return fmt.Errorf("proto: SpendRebatesRewardReq: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: SpendRebatesRewardReq: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Id", wireType)
			}
			m.Id = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSpendRebates
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
			skippy, err := skipSpendRebates(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthSpendRebates
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
func (m *SpendRebatesRewardAck) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowSpendRebates
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
			return fmt.Errorf("proto: SpendRebatesRewardAck: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: SpendRebatesRewardAck: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Id", wireType)
			}
			m.Id = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSpendRebates
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
					return ErrIntOverflowSpendRebates
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
				return ErrInvalidLengthSpendRebates
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
			skippy, err := skipSpendRebates(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthSpendRebates
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
func (m *SpendRebatesNtf) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowSpendRebates
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
			return fmt.Errorf("proto: SpendRebatesNtf: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: SpendRebatesNtf: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field CountIngot", wireType)
			}
			m.CountIngot = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSpendRebates
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.CountIngot |= (int32(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Ingot", wireType)
			}
			m.Ingot = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSpendRebates
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Ingot |= (int32(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipSpendRebates(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthSpendRebates
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
func skipSpendRebates(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowSpendRebates
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
					return 0, ErrIntOverflowSpendRebates
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
					return 0, ErrIntOverflowSpendRebates
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
				return 0, ErrInvalidLengthSpendRebates
			}
			return iNdEx, nil
		case 3:
			for {
				var innerWire uint64
				var start int = iNdEx
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return 0, ErrIntOverflowSpendRebates
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
				next, err := skipSpendRebates(dAtA[start:])
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
	ErrInvalidLengthSpendRebates = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowSpendRebates   = fmt.Errorf("proto: integer overflow")
)

func init() { proto.RegisterFile("spendRebates.proto", fileDescriptorSpendRebates) }

var fileDescriptorSpendRebates = []byte{
	// 198 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x12, 0x2a, 0x2e, 0x48, 0xcd,
	0x4b, 0x09, 0x4a, 0x4d, 0x4a, 0x2c, 0x49, 0x2d, 0xd6, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x62,
	0x2a, 0x48, 0x92, 0xe2, 0x49, 0xce, 0xcf, 0xcd, 0xcd, 0xcf, 0x83, 0x88, 0x28, 0xa9, 0x73, 0x89,
	0x06, 0x23, 0xa9, 0x0b, 0x4a, 0x2d, 0x4f, 0x2c, 0x4a, 0x09, 0x4a, 0x2d, 0x14, 0xe2, 0xe3, 0x62,
	0xca, 0x4c, 0x91, 0x60, 0x54, 0x60, 0xd4, 0x60, 0x0d, 0x62, 0xca, 0x4c, 0x51, 0x0a, 0xc4, 0xa6,
	0xd0, 0x31, 0x39, 0x1b, 0x5d, 0xa1, 0x90, 0x06, 0x17, 0x6b, 0x7a, 0x7e, 0x7e, 0x4a, 0xb1, 0x04,
	0x93, 0x02, 0xa3, 0x06, 0xb7, 0x91, 0x90, 0x5e, 0x41, 0x92, 0x9e, 0x3b, 0x48, 0xc0, 0x39, 0x23,
	0x31, 0x2f, 0x3d, 0xd5, 0xaf, 0x24, 0x2d, 0x08, 0xa2, 0x40, 0xc9, 0x9d, 0x8b, 0x1f, 0xd9, 0x48,
	0xbf, 0x92, 0x34, 0x21, 0x39, 0x2e, 0xae, 0xe4, 0xfc, 0xd2, 0xbc, 0x12, 0xcf, 0xbc, 0xf4, 0xfc,
	0x12, 0xa8, 0xa1, 0x48, 0x22, 0x42, 0x22, 0x5c, 0xac, 0x99, 0x60, 0x29, 0x26, 0xb0, 0x14, 0x84,
	0xe3, 0x24, 0x70, 0xe2, 0x91, 0x1c, 0xe3, 0x85, 0x47, 0x72, 0x8c, 0x0f, 0x1e, 0xc9, 0x31, 0xce,
	0x78, 0x2c, 0xc7, 0x90, 0xc4, 0x06, 0xf6, 0x9d, 0x31, 0x20, 0x00, 0x00, 0xff, 0xff, 0xcc, 0xb3,
	0xf5, 0xde, 0x05, 0x01, 0x00, 0x00,
}
