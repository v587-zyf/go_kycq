// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: fashion.proto

package pb

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

import io "io"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// 时装升级（激活）
type FashionUpLevelReq struct {
	HeroIndex int32 `protobuf:"varint,1,opt,name=heroIndex,proto3" json:"heroIndex,omitempty"`
	FashionId int32 `protobuf:"varint,2,opt,name=fashionId,proto3" json:"fashionId,omitempty"`
}

func (m *FashionUpLevelReq) Reset()                    { *m = FashionUpLevelReq{} }
func (m *FashionUpLevelReq) String() string            { return proto.CompactTextString(m) }
func (*FashionUpLevelReq) ProtoMessage()               {}
func (*FashionUpLevelReq) Descriptor() ([]byte, []int) { return fileDescriptorFashion, []int{0} }

func (m *FashionUpLevelReq) GetHeroIndex() int32 {
	if m != nil {
		return m.HeroIndex
	}
	return 0
}

func (m *FashionUpLevelReq) GetFashionId() int32 {
	if m != nil {
		return m.FashionId
	}
	return 0
}

type FashionUpLevelAck struct {
	HeroIndex int32    `protobuf:"varint,1,opt,name=heroIndex,proto3" json:"heroIndex,omitempty"`
	Fashion   *Fashion `protobuf:"bytes,2,opt,name=fashion" json:"fashion,omitempty"`
}

func (m *FashionUpLevelAck) Reset()                    { *m = FashionUpLevelAck{} }
func (m *FashionUpLevelAck) String() string            { return proto.CompactTextString(m) }
func (*FashionUpLevelAck) ProtoMessage()               {}
func (*FashionUpLevelAck) Descriptor() ([]byte, []int) { return fileDescriptorFashion, []int{1} }

func (m *FashionUpLevelAck) GetHeroIndex() int32 {
	if m != nil {
		return m.HeroIndex
	}
	return 0
}

func (m *FashionUpLevelAck) GetFashion() *Fashion {
	if m != nil {
		return m.Fashion
	}
	return nil
}

// 时装穿戴
type FashionWearReq struct {
	HeroIndex int32 `protobuf:"varint,1,opt,name=heroIndex,proto3" json:"heroIndex,omitempty"`
	FashionId int32 `protobuf:"varint,2,opt,name=fashionId,proto3" json:"fashionId,omitempty"`
	IsWear    bool  `protobuf:"varint,3,opt,name=isWear,proto3" json:"isWear,omitempty"`
}

func (m *FashionWearReq) Reset()                    { *m = FashionWearReq{} }
func (m *FashionWearReq) String() string            { return proto.CompactTextString(m) }
func (*FashionWearReq) ProtoMessage()               {}
func (*FashionWearReq) Descriptor() ([]byte, []int) { return fileDescriptorFashion, []int{2} }

func (m *FashionWearReq) GetHeroIndex() int32 {
	if m != nil {
		return m.HeroIndex
	}
	return 0
}

func (m *FashionWearReq) GetFashionId() int32 {
	if m != nil {
		return m.FashionId
	}
	return 0
}

func (m *FashionWearReq) GetIsWear() bool {
	if m != nil {
		return m.IsWear
	}
	return false
}

type FashionWearAck struct {
	HeroIndex     int32 `protobuf:"varint,1,opt,name=heroIndex,proto3" json:"heroIndex,omitempty"`
	WearFashionId int32 `protobuf:"varint,2,opt,name=wearFashionId,proto3" json:"wearFashionId,omitempty"`
	IsWear        bool  `protobuf:"varint,3,opt,name=isWear,proto3" json:"isWear,omitempty"`
}

func (m *FashionWearAck) Reset()                    { *m = FashionWearAck{} }
func (m *FashionWearAck) String() string            { return proto.CompactTextString(m) }
func (*FashionWearAck) ProtoMessage()               {}
func (*FashionWearAck) Descriptor() ([]byte, []int) { return fileDescriptorFashion, []int{3} }

func (m *FashionWearAck) GetHeroIndex() int32 {
	if m != nil {
		return m.HeroIndex
	}
	return 0
}

func (m *FashionWearAck) GetWearFashionId() int32 {
	if m != nil {
		return m.WearFashionId
	}
	return 0
}

func (m *FashionWearAck) GetIsWear() bool {
	if m != nil {
		return m.IsWear
	}
	return false
}

func init() {
	proto.RegisterType((*FashionUpLevelReq)(nil), "pb.FashionUpLevelReq")
	proto.RegisterType((*FashionUpLevelAck)(nil), "pb.FashionUpLevelAck")
	proto.RegisterType((*FashionWearReq)(nil), "pb.FashionWearReq")
	proto.RegisterType((*FashionWearAck)(nil), "pb.FashionWearAck")
}
func (m *FashionUpLevelReq) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *FashionUpLevelReq) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if m.HeroIndex != 0 {
		dAtA[i] = 0x8
		i++
		i = encodeVarintFashion(dAtA, i, uint64(m.HeroIndex))
	}
	if m.FashionId != 0 {
		dAtA[i] = 0x10
		i++
		i = encodeVarintFashion(dAtA, i, uint64(m.FashionId))
	}
	return i, nil
}

func (m *FashionUpLevelAck) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *FashionUpLevelAck) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if m.HeroIndex != 0 {
		dAtA[i] = 0x8
		i++
		i = encodeVarintFashion(dAtA, i, uint64(m.HeroIndex))
	}
	if m.Fashion != nil {
		dAtA[i] = 0x12
		i++
		i = encodeVarintFashion(dAtA, i, uint64(m.Fashion.Size()))
		n1, err := m.Fashion.MarshalTo(dAtA[i:])
		if err != nil {
			return 0, err
		}
		i += n1
	}
	return i, nil
}

func (m *FashionWearReq) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *FashionWearReq) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if m.HeroIndex != 0 {
		dAtA[i] = 0x8
		i++
		i = encodeVarintFashion(dAtA, i, uint64(m.HeroIndex))
	}
	if m.FashionId != 0 {
		dAtA[i] = 0x10
		i++
		i = encodeVarintFashion(dAtA, i, uint64(m.FashionId))
	}
	if m.IsWear {
		dAtA[i] = 0x18
		i++
		if m.IsWear {
			dAtA[i] = 1
		} else {
			dAtA[i] = 0
		}
		i++
	}
	return i, nil
}

func (m *FashionWearAck) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *FashionWearAck) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if m.HeroIndex != 0 {
		dAtA[i] = 0x8
		i++
		i = encodeVarintFashion(dAtA, i, uint64(m.HeroIndex))
	}
	if m.WearFashionId != 0 {
		dAtA[i] = 0x10
		i++
		i = encodeVarintFashion(dAtA, i, uint64(m.WearFashionId))
	}
	if m.IsWear {
		dAtA[i] = 0x18
		i++
		if m.IsWear {
			dAtA[i] = 1
		} else {
			dAtA[i] = 0
		}
		i++
	}
	return i, nil
}

func encodeVarintFashion(dAtA []byte, offset int, v uint64) int {
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return offset + 1
}
func (m *FashionUpLevelReq) Size() (n int) {
	var l int
	_ = l
	if m.HeroIndex != 0 {
		n += 1 + sovFashion(uint64(m.HeroIndex))
	}
	if m.FashionId != 0 {
		n += 1 + sovFashion(uint64(m.FashionId))
	}
	return n
}

func (m *FashionUpLevelAck) Size() (n int) {
	var l int
	_ = l
	if m.HeroIndex != 0 {
		n += 1 + sovFashion(uint64(m.HeroIndex))
	}
	if m.Fashion != nil {
		l = m.Fashion.Size()
		n += 1 + l + sovFashion(uint64(l))
	}
	return n
}

func (m *FashionWearReq) Size() (n int) {
	var l int
	_ = l
	if m.HeroIndex != 0 {
		n += 1 + sovFashion(uint64(m.HeroIndex))
	}
	if m.FashionId != 0 {
		n += 1 + sovFashion(uint64(m.FashionId))
	}
	if m.IsWear {
		n += 2
	}
	return n
}

func (m *FashionWearAck) Size() (n int) {
	var l int
	_ = l
	if m.HeroIndex != 0 {
		n += 1 + sovFashion(uint64(m.HeroIndex))
	}
	if m.WearFashionId != 0 {
		n += 1 + sovFashion(uint64(m.WearFashionId))
	}
	if m.IsWear {
		n += 2
	}
	return n
}

func sovFashion(x uint64) (n int) {
	for {
		n++
		x >>= 7
		if x == 0 {
			break
		}
	}
	return n
}
func sozFashion(x uint64) (n int) {
	return sovFashion(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *FashionUpLevelReq) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowFashion
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
			return fmt.Errorf("proto: FashionUpLevelReq: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: FashionUpLevelReq: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field HeroIndex", wireType)
			}
			m.HeroIndex = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowFashion
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
				return fmt.Errorf("proto: wrong wireType = %d for field FashionId", wireType)
			}
			m.FashionId = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowFashion
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.FashionId |= (int32(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipFashion(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthFashion
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
func (m *FashionUpLevelAck) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowFashion
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
			return fmt.Errorf("proto: FashionUpLevelAck: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: FashionUpLevelAck: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field HeroIndex", wireType)
			}
			m.HeroIndex = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowFashion
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
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Fashion", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowFashion
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
				return ErrInvalidLengthFashion
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Fashion == nil {
				m.Fashion = &Fashion{}
			}
			if err := m.Fashion.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipFashion(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthFashion
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
func (m *FashionWearReq) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowFashion
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
			return fmt.Errorf("proto: FashionWearReq: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: FashionWearReq: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field HeroIndex", wireType)
			}
			m.HeroIndex = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowFashion
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
				return fmt.Errorf("proto: wrong wireType = %d for field FashionId", wireType)
			}
			m.FashionId = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowFashion
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.FashionId |= (int32(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field IsWear", wireType)
			}
			var v int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowFashion
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
			m.IsWear = bool(v != 0)
		default:
			iNdEx = preIndex
			skippy, err := skipFashion(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthFashion
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
func (m *FashionWearAck) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowFashion
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
			return fmt.Errorf("proto: FashionWearAck: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: FashionWearAck: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field HeroIndex", wireType)
			}
			m.HeroIndex = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowFashion
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
				return fmt.Errorf("proto: wrong wireType = %d for field WearFashionId", wireType)
			}
			m.WearFashionId = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowFashion
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.WearFashionId |= (int32(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field IsWear", wireType)
			}
			var v int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowFashion
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
			m.IsWear = bool(v != 0)
		default:
			iNdEx = preIndex
			skippy, err := skipFashion(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthFashion
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
func skipFashion(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowFashion
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
					return 0, ErrIntOverflowFashion
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
					return 0, ErrIntOverflowFashion
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
				return 0, ErrInvalidLengthFashion
			}
			return iNdEx, nil
		case 3:
			for {
				var innerWire uint64
				var start int = iNdEx
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return 0, ErrIntOverflowFashion
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
				next, err := skipFashion(dAtA[start:])
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
	ErrInvalidLengthFashion = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowFashion   = fmt.Errorf("proto: integer overflow")
)

func init() { proto.RegisterFile("fashion.proto", fileDescriptorFashion) }

var fileDescriptorFashion = []byte{
	// 205 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0x4d, 0x4b, 0x2c, 0xce,
	0xc8, 0xcc, 0xcf, 0xd3, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x62, 0x2a, 0x48, 0x92, 0xe2, 0x49,
	0xce, 0xcf, 0xcd, 0x85, 0x89, 0x28, 0xf9, 0x73, 0x09, 0xba, 0x41, 0x94, 0x84, 0x16, 0xf8, 0xa4,
	0x96, 0xa5, 0xe6, 0x04, 0xa5, 0x16, 0x0a, 0xc9, 0x70, 0x71, 0x66, 0xa4, 0x16, 0xe5, 0x7b, 0xe6,
	0xa5, 0xa4, 0x56, 0x48, 0x30, 0x2a, 0x30, 0x6a, 0xb0, 0x06, 0x21, 0x04, 0x40, 0xb2, 0x50, 0x53,
	0x3d, 0x53, 0x24, 0x98, 0x20, 0xb2, 0x70, 0x01, 0xa5, 0x08, 0x74, 0x03, 0x1d, 0x93, 0xb3, 0x09,
	0x18, 0xa8, 0xca, 0xc5, 0x0e, 0xd5, 0x0f, 0x36, 0x8e, 0xdb, 0x88, 0x5b, 0xaf, 0x20, 0x49, 0x0f,
	0x6a, 0x4a, 0x10, 0x4c, 0x4e, 0x29, 0x85, 0x8b, 0x0f, 0x2a, 0x16, 0x9e, 0x9a, 0x58, 0x44, 0xa1,
	0x3b, 0x85, 0xc4, 0xb8, 0xd8, 0x32, 0x8b, 0x41, 0x06, 0x49, 0x30, 0x2b, 0x30, 0x6a, 0x70, 0x04,
	0x41, 0x79, 0x4a, 0x39, 0x28, 0xb6, 0x10, 0x76, 0xbc, 0x0a, 0x17, 0x6f, 0x79, 0x6a, 0x62, 0x91,
	0x1b, 0x9a, 0x4d, 0xa8, 0x82, 0xb8, 0x6c, 0x73, 0x12, 0x38, 0xf1, 0x48, 0x8e, 0xf1, 0xc2, 0x23,
	0x39, 0xc6, 0x07, 0x8f, 0xe4, 0x18, 0x67, 0x3c, 0x96, 0x63, 0x48, 0x62, 0x03, 0xc7, 0x8b, 0x31,
	0x20, 0x00, 0x00, 0xff, 0xff, 0xa5, 0x8e, 0xe9, 0x7e, 0xba, 0x01, 0x00, 0x00,
}
