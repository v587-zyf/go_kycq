// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: magicTower.proto

package pb

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

import io "io"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// 九重魔塔结算
type MagicTowerEndNtf struct {
	Rank int32 `protobuf:"varint,1,opt,name=rank,proto3" json:"rank,omitempty"`
}

func (m *MagicTowerEndNtf) Reset()                    { *m = MagicTowerEndNtf{} }
func (m *MagicTowerEndNtf) String() string            { return proto.CompactTextString(m) }
func (*MagicTowerEndNtf) ProtoMessage()               {}
func (*MagicTowerEndNtf) Descriptor() ([]byte, []int) { return fileDescriptorMagicTower, []int{0} }

func (m *MagicTowerEndNtf) GetRank() int32 {
	if m != nil {
		return m.Rank
	}
	return 0
}

// 获取玩家当前分数
type MagicTowerGetUserInfoReq struct {
}

func (m *MagicTowerGetUserInfoReq) Reset()         { *m = MagicTowerGetUserInfoReq{} }
func (m *MagicTowerGetUserInfoReq) String() string { return proto.CompactTextString(m) }
func (*MagicTowerGetUserInfoReq) ProtoMessage()    {}
func (*MagicTowerGetUserInfoReq) Descriptor() ([]byte, []int) {
	return fileDescriptorMagicTower, []int{1}
}

// 获取玩家当前分数
type MagicTowerGetUserInfoAck struct {
	Score      int32 `protobuf:"varint,1,opt,name=score,proto3" json:"score,omitempty"`
	IsGetAward int32 `protobuf:"varint,2,opt,name=isGetAward,proto3" json:"isGetAward,omitempty"`
}

func (m *MagicTowerGetUserInfoAck) Reset()         { *m = MagicTowerGetUserInfoAck{} }
func (m *MagicTowerGetUserInfoAck) String() string { return proto.CompactTextString(m) }
func (*MagicTowerGetUserInfoAck) ProtoMessage()    {}
func (*MagicTowerGetUserInfoAck) Descriptor() ([]byte, []int) {
	return fileDescriptorMagicTower, []int{2}
}

func (m *MagicTowerGetUserInfoAck) GetScore() int32 {
	if m != nil {
		return m.Score
	}
	return 0
}

func (m *MagicTowerGetUserInfoAck) GetIsGetAward() int32 {
	if m != nil {
		return m.IsGetAward
	}
	return 0
}

// 获取玩家当前分数
type MagicTowerlayerAwardReq struct {
}

func (m *MagicTowerlayerAwardReq) Reset()         { *m = MagicTowerlayerAwardReq{} }
func (m *MagicTowerlayerAwardReq) String() string { return proto.CompactTextString(m) }
func (*MagicTowerlayerAwardReq) ProtoMessage()    {}
func (*MagicTowerlayerAwardReq) Descriptor() ([]byte, []int) {
	return fileDescriptorMagicTower, []int{3}
}

// 获取玩家当前分数
type MagicTowerlayerAwardAck struct {
	Goods *GoodsChangeNtf `protobuf:"bytes,1,opt,name=goods" json:"goods,omitempty"`
}

func (m *MagicTowerlayerAwardAck) Reset()         { *m = MagicTowerlayerAwardAck{} }
func (m *MagicTowerlayerAwardAck) String() string { return proto.CompactTextString(m) }
func (*MagicTowerlayerAwardAck) ProtoMessage()    {}
func (*MagicTowerlayerAwardAck) Descriptor() ([]byte, []int) {
	return fileDescriptorMagicTower, []int{4}
}

func (m *MagicTowerlayerAwardAck) GetGoods() *GoodsChangeNtf {
	if m != nil {
		return m.Goods
	}
	return nil
}

func init() {
	proto.RegisterType((*MagicTowerEndNtf)(nil), "pb.MagicTowerEndNtf")
	proto.RegisterType((*MagicTowerGetUserInfoReq)(nil), "pb.MagicTowerGetUserInfoReq")
	proto.RegisterType((*MagicTowerGetUserInfoAck)(nil), "pb.MagicTowerGetUserInfoAck")
	proto.RegisterType((*MagicTowerlayerAwardReq)(nil), "pb.MagicTowerlayerAwardReq")
	proto.RegisterType((*MagicTowerlayerAwardAck)(nil), "pb.MagicTowerlayerAwardAck")
}
func (m *MagicTowerEndNtf) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MagicTowerEndNtf) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if m.Rank != 0 {
		dAtA[i] = 0x8
		i++
		i = encodeVarintMagicTower(dAtA, i, uint64(m.Rank))
	}
	return i, nil
}

func (m *MagicTowerGetUserInfoReq) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MagicTowerGetUserInfoReq) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	return i, nil
}

func (m *MagicTowerGetUserInfoAck) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MagicTowerGetUserInfoAck) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if m.Score != 0 {
		dAtA[i] = 0x8
		i++
		i = encodeVarintMagicTower(dAtA, i, uint64(m.Score))
	}
	if m.IsGetAward != 0 {
		dAtA[i] = 0x10
		i++
		i = encodeVarintMagicTower(dAtA, i, uint64(m.IsGetAward))
	}
	return i, nil
}

func (m *MagicTowerlayerAwardReq) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MagicTowerlayerAwardReq) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	return i, nil
}

func (m *MagicTowerlayerAwardAck) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MagicTowerlayerAwardAck) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if m.Goods != nil {
		dAtA[i] = 0xa
		i++
		i = encodeVarintMagicTower(dAtA, i, uint64(m.Goods.Size()))
		n1, err := m.Goods.MarshalTo(dAtA[i:])
		if err != nil {
			return 0, err
		}
		i += n1
	}
	return i, nil
}

func encodeVarintMagicTower(dAtA []byte, offset int, v uint64) int {
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return offset + 1
}
func (m *MagicTowerEndNtf) Size() (n int) {
	var l int
	_ = l
	if m.Rank != 0 {
		n += 1 + sovMagicTower(uint64(m.Rank))
	}
	return n
}

func (m *MagicTowerGetUserInfoReq) Size() (n int) {
	var l int
	_ = l
	return n
}

func (m *MagicTowerGetUserInfoAck) Size() (n int) {
	var l int
	_ = l
	if m.Score != 0 {
		n += 1 + sovMagicTower(uint64(m.Score))
	}
	if m.IsGetAward != 0 {
		n += 1 + sovMagicTower(uint64(m.IsGetAward))
	}
	return n
}

func (m *MagicTowerlayerAwardReq) Size() (n int) {
	var l int
	_ = l
	return n
}

func (m *MagicTowerlayerAwardAck) Size() (n int) {
	var l int
	_ = l
	if m.Goods != nil {
		l = m.Goods.Size()
		n += 1 + l + sovMagicTower(uint64(l))
	}
	return n
}

func sovMagicTower(x uint64) (n int) {
	for {
		n++
		x >>= 7
		if x == 0 {
			break
		}
	}
	return n
}
func sozMagicTower(x uint64) (n int) {
	return sovMagicTower(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *MagicTowerEndNtf) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowMagicTower
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
			return fmt.Errorf("proto: MagicTowerEndNtf: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MagicTowerEndNtf: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Rank", wireType)
			}
			m.Rank = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMagicTower
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Rank |= (int32(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipMagicTower(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthMagicTower
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
func (m *MagicTowerGetUserInfoReq) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowMagicTower
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
			return fmt.Errorf("proto: MagicTowerGetUserInfoReq: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MagicTowerGetUserInfoReq: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		default:
			iNdEx = preIndex
			skippy, err := skipMagicTower(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthMagicTower
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
func (m *MagicTowerGetUserInfoAck) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowMagicTower
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
			return fmt.Errorf("proto: MagicTowerGetUserInfoAck: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MagicTowerGetUserInfoAck: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Score", wireType)
			}
			m.Score = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMagicTower
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Score |= (int32(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field IsGetAward", wireType)
			}
			m.IsGetAward = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMagicTower
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.IsGetAward |= (int32(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipMagicTower(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthMagicTower
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
func (m *MagicTowerlayerAwardReq) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowMagicTower
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
			return fmt.Errorf("proto: MagicTowerlayerAwardReq: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MagicTowerlayerAwardReq: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		default:
			iNdEx = preIndex
			skippy, err := skipMagicTower(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthMagicTower
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
func (m *MagicTowerlayerAwardAck) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowMagicTower
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
			return fmt.Errorf("proto: MagicTowerlayerAwardAck: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MagicTowerlayerAwardAck: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Goods", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMagicTower
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
				return ErrInvalidLengthMagicTower
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
			skippy, err := skipMagicTower(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthMagicTower
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
func skipMagicTower(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowMagicTower
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
					return 0, ErrIntOverflowMagicTower
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
					return 0, ErrIntOverflowMagicTower
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
				return 0, ErrInvalidLengthMagicTower
			}
			return iNdEx, nil
		case 3:
			for {
				var innerWire uint64
				var start int = iNdEx
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return 0, ErrIntOverflowMagicTower
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
				next, err := skipMagicTower(dAtA[start:])
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
	ErrInvalidLengthMagicTower = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowMagicTower   = fmt.Errorf("proto: integer overflow")
)

func init() { proto.RegisterFile("magicTower.proto", fileDescriptorMagicTower) }

var fileDescriptorMagicTower = []byte{
	// 222 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x12, 0xc8, 0x4d, 0x4c, 0xcf,
	0x4c, 0x0e, 0xc9, 0x2f, 0x4f, 0x2d, 0xd2, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x62, 0x2a, 0x48,
	0x92, 0xe2, 0x49, 0xce, 0xcf, 0xcd, 0xcd, 0xcf, 0x83, 0x88, 0x28, 0xa9, 0x71, 0x09, 0xf8, 0xc2,
	0x55, 0xb9, 0xe6, 0xa5, 0xf8, 0x95, 0xa4, 0x09, 0x09, 0x71, 0xb1, 0x14, 0x25, 0xe6, 0x65, 0x4b,
	0x30, 0x2a, 0x30, 0x6a, 0xb0, 0x06, 0x81, 0xd9, 0x4a, 0x52, 0x5c, 0x12, 0x08, 0x75, 0xee, 0xa9,
	0x25, 0xa1, 0xc5, 0xa9, 0x45, 0x9e, 0x79, 0x69, 0xf9, 0x41, 0xa9, 0x85, 0x4a, 0x01, 0x38, 0xe4,
	0x1c, 0x93, 0xb3, 0x85, 0x44, 0xb8, 0x58, 0x8b, 0x93, 0xf3, 0x8b, 0x52, 0xa1, 0x86, 0x41, 0x38,
	0x42, 0x72, 0x5c, 0x5c, 0x99, 0xc5, 0xee, 0xa9, 0x25, 0x8e, 0xe5, 0x89, 0x45, 0x29, 0x12, 0x4c,
	0x60, 0x29, 0x24, 0x11, 0x25, 0x49, 0x2e, 0x71, 0x84, 0x89, 0x39, 0x89, 0x95, 0xa9, 0x45, 0x60,
	0x71, 0x90, 0x65, 0xce, 0xd8, 0xa5, 0x40, 0x76, 0x69, 0x70, 0xb1, 0xa6, 0xe7, 0xe7, 0xa7, 0x14,
	0x83, 0xed, 0xe2, 0x36, 0x12, 0xd2, 0x2b, 0x48, 0xd2, 0x73, 0x07, 0x09, 0x38, 0x67, 0x24, 0xe6,
	0xa5, 0xa7, 0xfa, 0x95, 0xa4, 0x05, 0x41, 0x14, 0x38, 0x09, 0x9c, 0x78, 0x24, 0xc7, 0x78, 0xe1,
	0x91, 0x1c, 0xe3, 0x83, 0x47, 0x72, 0x8c, 0x33, 0x1e, 0xcb, 0x31, 0x24, 0xb1, 0x81, 0x83, 0xc3,
	0x18, 0x10, 0x00, 0x00, 0xff, 0xff, 0xa6, 0x18, 0xd4, 0x03, 0x34, 0x01, 0x00, 0x00,
}
