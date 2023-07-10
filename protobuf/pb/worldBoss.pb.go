// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: worldBoss.proto

package pb

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

import io "io"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// 战斗
type EnterWorldBossFightReq struct {
	StageId int32 `protobuf:"varint,1,opt,name=stageId,proto3" json:"stageId,omitempty"`
}

func (m *EnterWorldBossFightReq) Reset()                    { *m = EnterWorldBossFightReq{} }
func (m *EnterWorldBossFightReq) String() string            { return proto.CompactTextString(m) }
func (*EnterWorldBossFightReq) ProtoMessage()               {}
func (*EnterWorldBossFightReq) Descriptor() ([]byte, []int) { return fileDescriptorWorldBoss, []int{0} }

func (m *EnterWorldBossFightReq) GetStageId() int32 {
	if m != nil {
		return m.StageId
	}
	return 0
}

type WorldBossFightResultNtf struct {
	StageId int32           `protobuf:"varint,1,opt,name=stageId,proto3" json:"stageId,omitempty"`
	Rank    int32           `protobuf:"varint,2,opt,name=rank,proto3" json:"rank,omitempty"`
	Goods   *GoodsChangeNtf `protobuf:"bytes,3,opt,name=goods" json:"goods,omitempty"`
}

func (m *WorldBossFightResultNtf) Reset()                    { *m = WorldBossFightResultNtf{} }
func (m *WorldBossFightResultNtf) String() string            { return proto.CompactTextString(m) }
func (*WorldBossFightResultNtf) ProtoMessage()               {}
func (*WorldBossFightResultNtf) Descriptor() ([]byte, []int) { return fileDescriptorWorldBoss, []int{1} }

func (m *WorldBossFightResultNtf) GetStageId() int32 {
	if m != nil {
		return m.StageId
	}
	return 0
}

func (m *WorldBossFightResultNtf) GetRank() int32 {
	if m != nil {
		return m.Rank
	}
	return 0
}

func (m *WorldBossFightResultNtf) GetGoods() *GoodsChangeNtf {
	if m != nil {
		return m.Goods
	}
	return nil
}

func init() {
	proto.RegisterType((*EnterWorldBossFightReq)(nil), "pb.EnterWorldBossFightReq")
	proto.RegisterType((*WorldBossFightResultNtf)(nil), "pb.WorldBossFightResultNtf")
}
func (m *EnterWorldBossFightReq) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *EnterWorldBossFightReq) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if m.StageId != 0 {
		dAtA[i] = 0x8
		i++
		i = encodeVarintWorldBoss(dAtA, i, uint64(m.StageId))
	}
	return i, nil
}

func (m *WorldBossFightResultNtf) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *WorldBossFightResultNtf) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if m.StageId != 0 {
		dAtA[i] = 0x8
		i++
		i = encodeVarintWorldBoss(dAtA, i, uint64(m.StageId))
	}
	if m.Rank != 0 {
		dAtA[i] = 0x10
		i++
		i = encodeVarintWorldBoss(dAtA, i, uint64(m.Rank))
	}
	if m.Goods != nil {
		dAtA[i] = 0x1a
		i++
		i = encodeVarintWorldBoss(dAtA, i, uint64(m.Goods.Size()))
		n1, err := m.Goods.MarshalTo(dAtA[i:])
		if err != nil {
			return 0, err
		}
		i += n1
	}
	return i, nil
}

func encodeVarintWorldBoss(dAtA []byte, offset int, v uint64) int {
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return offset + 1
}
func (m *EnterWorldBossFightReq) Size() (n int) {
	var l int
	_ = l
	if m.StageId != 0 {
		n += 1 + sovWorldBoss(uint64(m.StageId))
	}
	return n
}

func (m *WorldBossFightResultNtf) Size() (n int) {
	var l int
	_ = l
	if m.StageId != 0 {
		n += 1 + sovWorldBoss(uint64(m.StageId))
	}
	if m.Rank != 0 {
		n += 1 + sovWorldBoss(uint64(m.Rank))
	}
	if m.Goods != nil {
		l = m.Goods.Size()
		n += 1 + l + sovWorldBoss(uint64(l))
	}
	return n
}

func sovWorldBoss(x uint64) (n int) {
	for {
		n++
		x >>= 7
		if x == 0 {
			break
		}
	}
	return n
}
func sozWorldBoss(x uint64) (n int) {
	return sovWorldBoss(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *EnterWorldBossFightReq) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowWorldBoss
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
			return fmt.Errorf("proto: EnterWorldBossFightReq: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: EnterWorldBossFightReq: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field StageId", wireType)
			}
			m.StageId = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowWorldBoss
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.StageId |= (int32(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipWorldBoss(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthWorldBoss
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
func (m *WorldBossFightResultNtf) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowWorldBoss
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
			return fmt.Errorf("proto: WorldBossFightResultNtf: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: WorldBossFightResultNtf: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field StageId", wireType)
			}
			m.StageId = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowWorldBoss
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.StageId |= (int32(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Rank", wireType)
			}
			m.Rank = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowWorldBoss
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
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Goods", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowWorldBoss
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
				return ErrInvalidLengthWorldBoss
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
			skippy, err := skipWorldBoss(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthWorldBoss
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
func skipWorldBoss(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowWorldBoss
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
					return 0, ErrIntOverflowWorldBoss
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
					return 0, ErrIntOverflowWorldBoss
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
				return 0, ErrInvalidLengthWorldBoss
			}
			return iNdEx, nil
		case 3:
			for {
				var innerWire uint64
				var start int = iNdEx
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return 0, ErrIntOverflowWorldBoss
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
				next, err := skipWorldBoss(dAtA[start:])
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
	ErrInvalidLengthWorldBoss = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowWorldBoss   = fmt.Errorf("proto: integer overflow")
)

func init() { proto.RegisterFile("worldBoss.proto", fileDescriptorWorldBoss) }

var fileDescriptorWorldBoss = []byte{
	// 184 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0x2f, 0xcf, 0x2f, 0xca,
	0x49, 0x71, 0xca, 0x2f, 0x2e, 0xd6, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x62, 0x2a, 0x48, 0x92,
	0xe2, 0x49, 0xce, 0xcf, 0xcd, 0xcd, 0xcf, 0x83, 0x88, 0x28, 0x19, 0x71, 0x89, 0xb9, 0xe6, 0x95,
	0xa4, 0x16, 0x85, 0xc3, 0x54, 0xba, 0x65, 0xa6, 0x67, 0x94, 0x04, 0xa5, 0x16, 0x0a, 0x49, 0x70,
	0xb1, 0x17, 0x97, 0x24, 0xa6, 0xa7, 0x7a, 0xa6, 0x48, 0x30, 0x2a, 0x30, 0x6a, 0xb0, 0x06, 0xc1,
	0xb8, 0x4a, 0x85, 0x5c, 0xe2, 0xe8, 0xca, 0x8b, 0x4b, 0x73, 0x4a, 0xfc, 0x4a, 0xd2, 0x70, 0x6b,
	0x12, 0x12, 0xe2, 0x62, 0x29, 0x4a, 0xcc, 0xcb, 0x96, 0x60, 0x02, 0x0b, 0x83, 0xd9, 0x42, 0x1a,
	0x5c, 0xac, 0xe9, 0xf9, 0xf9, 0x29, 0xc5, 0x12, 0xcc, 0x0a, 0x8c, 0x1a, 0xdc, 0x46, 0x42, 0x7a,
	0x05, 0x49, 0x7a, 0xee, 0x20, 0x01, 0xe7, 0x8c, 0xc4, 0xbc, 0xf4, 0x54, 0xbf, 0x92, 0xb4, 0x20,
	0x88, 0x02, 0x27, 0x81, 0x13, 0x8f, 0xe4, 0x18, 0x2f, 0x3c, 0x92, 0x63, 0x7c, 0xf0, 0x48, 0x8e,
	0x71, 0xc6, 0x63, 0x39, 0x86, 0x24, 0x36, 0xb0, 0xfb, 0x8d, 0x01, 0x01, 0x00, 0x00, 0xff, 0xff,
	0x60, 0xd4, 0xa7, 0xcc, 0xe4, 0x00, 0x00, 0x00,
}
