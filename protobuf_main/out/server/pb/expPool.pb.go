// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: expPool.proto

package pb

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

import io "io"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

type ExpPoolLoadReq struct {
}

func (m *ExpPoolLoadReq) Reset()                    { *m = ExpPoolLoadReq{} }
func (m *ExpPoolLoadReq) String() string            { return proto.CompactTextString(m) }
func (*ExpPoolLoadReq) ProtoMessage()               {}
func (*ExpPoolLoadReq) Descriptor() ([]byte, []int) { return fileDescriptorExpPool, []int{0} }

type ExpPoolLoadAck struct {
	Heros   []*HeroInfo `protobuf:"bytes,1,rep,name=heros" json:"heros,omitempty"`
	WorlLvl int32       `protobuf:"varint,2,opt,name=worlLvl,proto3" json:"worlLvl,omitempty"`
	ExpPool int32       `protobuf:"varint,3,opt,name=expPool,proto3" json:"expPool,omitempty"`
}

func (m *ExpPoolLoadAck) Reset()                    { *m = ExpPoolLoadAck{} }
func (m *ExpPoolLoadAck) String() string            { return proto.CompactTextString(m) }
func (*ExpPoolLoadAck) ProtoMessage()               {}
func (*ExpPoolLoadAck) Descriptor() ([]byte, []int) { return fileDescriptorExpPool, []int{1} }

func (m *ExpPoolLoadAck) GetHeros() []*HeroInfo {
	if m != nil {
		return m.Heros
	}
	return nil
}

func (m *ExpPoolLoadAck) GetWorlLvl() int32 {
	if m != nil {
		return m.WorlLvl
	}
	return 0
}

func (m *ExpPoolLoadAck) GetExpPool() int32 {
	if m != nil {
		return m.ExpPool
	}
	return 0
}

// 玩家升级
type ExpPoolUpGradeReq struct {
	Index int32 `protobuf:"varint,1,opt,name=index,proto3" json:"index,omitempty"`
}

func (m *ExpPoolUpGradeReq) Reset()                    { *m = ExpPoolUpGradeReq{} }
func (m *ExpPoolUpGradeReq) String() string            { return proto.CompactTextString(m) }
func (*ExpPoolUpGradeReq) ProtoMessage()               {}
func (*ExpPoolUpGradeReq) Descriptor() ([]byte, []int) { return fileDescriptorExpPool, []int{2} }

func (m *ExpPoolUpGradeReq) GetIndex() int32 {
	if m != nil {
		return m.Index
	}
	return 0
}

type ExpPoolUpGradeAck struct {
	HeroIndex int32 `protobuf:"varint,1,opt,name=heroIndex,proto3" json:"heroIndex,omitempty"`
	Lvl       int32 `protobuf:"varint,2,opt,name=Lvl,proto3" json:"Lvl,omitempty"`
}

func (m *ExpPoolUpGradeAck) Reset()                    { *m = ExpPoolUpGradeAck{} }
func (m *ExpPoolUpGradeAck) String() string            { return proto.CompactTextString(m) }
func (*ExpPoolUpGradeAck) ProtoMessage()               {}
func (*ExpPoolUpGradeAck) Descriptor() ([]byte, []int) { return fileDescriptorExpPool, []int{3} }

func (m *ExpPoolUpGradeAck) GetHeroIndex() int32 {
	if m != nil {
		return m.HeroIndex
	}
	return 0
}

func (m *ExpPoolUpGradeAck) GetLvl() int32 {
	if m != nil {
		return m.Lvl
	}
	return 0
}

func init() {
	proto.RegisterType((*ExpPoolLoadReq)(nil), "pb.ExpPoolLoadReq")
	proto.RegisterType((*ExpPoolLoadAck)(nil), "pb.ExpPoolLoadAck")
	proto.RegisterType((*ExpPoolUpGradeReq)(nil), "pb.ExpPoolUpGradeReq")
	proto.RegisterType((*ExpPoolUpGradeAck)(nil), "pb.ExpPoolUpGradeAck")
}
func (m *ExpPoolLoadReq) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *ExpPoolLoadReq) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	return i, nil
}

func (m *ExpPoolLoadAck) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *ExpPoolLoadAck) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if len(m.Heros) > 0 {
		for _, msg := range m.Heros {
			dAtA[i] = 0xa
			i++
			i = encodeVarintExpPool(dAtA, i, uint64(msg.Size()))
			n, err := msg.MarshalTo(dAtA[i:])
			if err != nil {
				return 0, err
			}
			i += n
		}
	}
	if m.WorlLvl != 0 {
		dAtA[i] = 0x10
		i++
		i = encodeVarintExpPool(dAtA, i, uint64(m.WorlLvl))
	}
	if m.ExpPool != 0 {
		dAtA[i] = 0x18
		i++
		i = encodeVarintExpPool(dAtA, i, uint64(m.ExpPool))
	}
	return i, nil
}

func (m *ExpPoolUpGradeReq) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *ExpPoolUpGradeReq) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if m.Index != 0 {
		dAtA[i] = 0x8
		i++
		i = encodeVarintExpPool(dAtA, i, uint64(m.Index))
	}
	return i, nil
}

func (m *ExpPoolUpGradeAck) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *ExpPoolUpGradeAck) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if m.HeroIndex != 0 {
		dAtA[i] = 0x8
		i++
		i = encodeVarintExpPool(dAtA, i, uint64(m.HeroIndex))
	}
	if m.Lvl != 0 {
		dAtA[i] = 0x10
		i++
		i = encodeVarintExpPool(dAtA, i, uint64(m.Lvl))
	}
	return i, nil
}

func encodeVarintExpPool(dAtA []byte, offset int, v uint64) int {
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return offset + 1
}
func (m *ExpPoolLoadReq) Size() (n int) {
	var l int
	_ = l
	return n
}

func (m *ExpPoolLoadAck) Size() (n int) {
	var l int
	_ = l
	if len(m.Heros) > 0 {
		for _, e := range m.Heros {
			l = e.Size()
			n += 1 + l + sovExpPool(uint64(l))
		}
	}
	if m.WorlLvl != 0 {
		n += 1 + sovExpPool(uint64(m.WorlLvl))
	}
	if m.ExpPool != 0 {
		n += 1 + sovExpPool(uint64(m.ExpPool))
	}
	return n
}

func (m *ExpPoolUpGradeReq) Size() (n int) {
	var l int
	_ = l
	if m.Index != 0 {
		n += 1 + sovExpPool(uint64(m.Index))
	}
	return n
}

func (m *ExpPoolUpGradeAck) Size() (n int) {
	var l int
	_ = l
	if m.HeroIndex != 0 {
		n += 1 + sovExpPool(uint64(m.HeroIndex))
	}
	if m.Lvl != 0 {
		n += 1 + sovExpPool(uint64(m.Lvl))
	}
	return n
}

func sovExpPool(x uint64) (n int) {
	for {
		n++
		x >>= 7
		if x == 0 {
			break
		}
	}
	return n
}
func sozExpPool(x uint64) (n int) {
	return sovExpPool(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *ExpPoolLoadReq) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowExpPool
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
			return fmt.Errorf("proto: ExpPoolLoadReq: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: ExpPoolLoadReq: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		default:
			iNdEx = preIndex
			skippy, err := skipExpPool(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthExpPool
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
func (m *ExpPoolLoadAck) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowExpPool
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
			return fmt.Errorf("proto: ExpPoolLoadAck: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: ExpPoolLoadAck: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Heros", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowExpPool
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
				return ErrInvalidLengthExpPool
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Heros = append(m.Heros, &HeroInfo{})
			if err := m.Heros[len(m.Heros)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field WorlLvl", wireType)
			}
			m.WorlLvl = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowExpPool
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.WorlLvl |= (int32(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field ExpPool", wireType)
			}
			m.ExpPool = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowExpPool
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.ExpPool |= (int32(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipExpPool(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthExpPool
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
func (m *ExpPoolUpGradeReq) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowExpPool
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
			return fmt.Errorf("proto: ExpPoolUpGradeReq: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: ExpPoolUpGradeReq: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Index", wireType)
			}
			m.Index = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowExpPool
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Index |= (int32(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipExpPool(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthExpPool
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
func (m *ExpPoolUpGradeAck) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowExpPool
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
			return fmt.Errorf("proto: ExpPoolUpGradeAck: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: ExpPoolUpGradeAck: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field HeroIndex", wireType)
			}
			m.HeroIndex = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowExpPool
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.HeroIndex |= (int32(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Lvl", wireType)
			}
			m.Lvl = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowExpPool
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Lvl |= (int32(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipExpPool(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthExpPool
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
func skipExpPool(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowExpPool
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
					return 0, ErrIntOverflowExpPool
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
					return 0, ErrIntOverflowExpPool
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
				return 0, ErrInvalidLengthExpPool
			}
			return iNdEx, nil
		case 3:
			for {
				var innerWire uint64
				var start int = iNdEx
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return 0, ErrIntOverflowExpPool
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
				next, err := skipExpPool(dAtA[start:])
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
	ErrInvalidLengthExpPool = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowExpPool   = fmt.Errorf("proto: integer overflow")
)

func init() { proto.RegisterFile("expPool.proto", fileDescriptorExpPool) }

var fileDescriptorExpPool = []byte{
	// 215 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0x4d, 0xad, 0x28, 0x08,
	0xc8, 0xcf, 0xcf, 0xd1, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x62, 0x2a, 0x48, 0x92, 0xe2, 0x49,
	0xce, 0xcf, 0xcd, 0xcd, 0xcf, 0x83, 0x88, 0x28, 0x09, 0x70, 0xf1, 0xb9, 0x42, 0x94, 0xf8, 0xe4,
	0x27, 0xa6, 0x04, 0xa5, 0x16, 0x2a, 0x65, 0xa0, 0x88, 0x38, 0x26, 0x67, 0x0b, 0x29, 0x71, 0xb1,
	0x66, 0xa4, 0x16, 0xe5, 0x17, 0x4b, 0x30, 0x2a, 0x30, 0x6b, 0x70, 0x1b, 0xf1, 0xe8, 0x15, 0x24,
	0xe9, 0x79, 0xa4, 0x16, 0xe5, 0x7b, 0xe6, 0xa5, 0xe5, 0x07, 0x41, 0xa4, 0x84, 0x24, 0xb8, 0xd8,
	0xcb, 0xf3, 0x8b, 0x72, 0x7c, 0xca, 0x72, 0x24, 0x98, 0x14, 0x18, 0x35, 0x58, 0x83, 0x60, 0x5c,
	0x90, 0x0c, 0xd4, 0x11, 0x12, 0xcc, 0x10, 0x19, 0x28, 0x57, 0x49, 0x93, 0x4b, 0x10, 0x6a, 0x53,
	0x68, 0x81, 0x7b, 0x51, 0x62, 0x4a, 0x6a, 0x50, 0x6a, 0xa1, 0x90, 0x08, 0x17, 0x6b, 0x66, 0x5e,
	0x4a, 0x6a, 0x85, 0x04, 0x23, 0x58, 0x31, 0x84, 0xa3, 0xe4, 0x8c, 0xae, 0x14, 0xe4, 0x2e, 0x19,
	0x2e, 0xce, 0x0c, 0xb0, 0x33, 0x10, 0xca, 0x11, 0x02, 0x42, 0x02, 0x5c, 0xcc, 0x08, 0xd7, 0x80,
	0x98, 0x4e, 0x02, 0x27, 0x1e, 0xc9, 0x31, 0x5e, 0x78, 0x24, 0xc7, 0xf8, 0xe0, 0x91, 0x1c, 0xe3,
	0x8c, 0xc7, 0x72, 0x0c, 0x49, 0x6c, 0xe0, 0x40, 0x30, 0x06, 0x04, 0x00, 0x00, 0xff, 0xff, 0x50,
	0xaa, 0xb0, 0x9c, 0x27, 0x01, 0x00, 0x00,
}