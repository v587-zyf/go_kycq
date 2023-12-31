// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: shabake.proto

package pbserver

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

import io "io"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

type GsToCcsBackGuildInfoNtf struct {
	FirstGuildInfo string `protobuf:"bytes,1,opt,name=firstGuildInfo,proto3" json:"firstGuildInfo,omitempty"`
	BenFuShaBaKe   int32  `protobuf:"varint,2,opt,name=benFuShaBaKe,proto3" json:"benFuShaBaKe,omitempty"`
	CrossFsId      int32  `protobuf:"varint,4,opt,name=crossFsId,proto3" json:"crossFsId,omitempty"`
}

func (m *GsToCcsBackGuildInfoNtf) Reset()                    { *m = GsToCcsBackGuildInfoNtf{} }
func (m *GsToCcsBackGuildInfoNtf) String() string            { return proto.CompactTextString(m) }
func (*GsToCcsBackGuildInfoNtf) ProtoMessage()               {}
func (*GsToCcsBackGuildInfoNtf) Descriptor() ([]byte, []int) { return fileDescriptorShabake, []int{0} }

func (m *GsToCcsBackGuildInfoNtf) GetFirstGuildInfo() string {
	if m != nil {
		return m.FirstGuildInfo
	}
	return ""
}

func (m *GsToCcsBackGuildInfoNtf) GetBenFuShaBaKe() int32 {
	if m != nil {
		return m.BenFuShaBaKe
	}
	return 0
}

func (m *GsToCcsBackGuildInfoNtf) GetCrossFsId() int32 {
	if m != nil {
		return m.CrossFsId
	}
	return 0
}

type CcsToGsBroadShaBakeFirstGuildInfo struct {
	FirstGuildInfo string `protobuf:"bytes,1,opt,name=firstGuildInfo,proto3" json:"firstGuildInfo,omitempty"`
	BenFuShaBaKe   int32  `protobuf:"varint,2,opt,name=benFuShaBaKe,proto3" json:"benFuShaBaKe,omitempty"`
}

func (m *CcsToGsBroadShaBakeFirstGuildInfo) Reset()         { *m = CcsToGsBroadShaBakeFirstGuildInfo{} }
func (m *CcsToGsBroadShaBakeFirstGuildInfo) String() string { return proto.CompactTextString(m) }
func (*CcsToGsBroadShaBakeFirstGuildInfo) ProtoMessage()    {}
func (*CcsToGsBroadShaBakeFirstGuildInfo) Descriptor() ([]byte, []int) {
	return fileDescriptorShabake, []int{1}
}

func (m *CcsToGsBroadShaBakeFirstGuildInfo) GetFirstGuildInfo() string {
	if m != nil {
		return m.FirstGuildInfo
	}
	return ""
}

func (m *CcsToGsBroadShaBakeFirstGuildInfo) GetBenFuShaBaKe() int32 {
	if m != nil {
		return m.BenFuShaBaKe
	}
	return 0
}

type Info struct {
	NickName string   `protobuf:"bytes,1,opt,name=nickName,proto3" json:"nickName,omitempty"`
	Sex      int32    `protobuf:"varint,2,opt,name=sex,proto3" json:"sex,omitempty"`
	Job      int32    `protobuf:"varint,3,opt,name=job,proto3" json:"job,omitempty"`
	Position int32    `protobuf:"varint,4,opt,name=position,proto3" json:"position,omitempty"`
	Display  *Display `protobuf:"bytes,5,opt,name=display" json:"display,omitempty"`
}

func (m *Info) Reset()                    { *m = Info{} }
func (m *Info) String() string            { return proto.CompactTextString(m) }
func (*Info) ProtoMessage()               {}
func (*Info) Descriptor() ([]byte, []int) { return fileDescriptorShabake, []int{2} }

func (m *Info) GetNickName() string {
	if m != nil {
		return m.NickName
	}
	return ""
}

func (m *Info) GetSex() int32 {
	if m != nil {
		return m.Sex
	}
	return 0
}

func (m *Info) GetJob() int32 {
	if m != nil {
		return m.Job
	}
	return 0
}

func (m *Info) GetPosition() int32 {
	if m != nil {
		return m.Position
	}
	return 0
}

func (m *Info) GetDisplay() *Display {
	if m != nil {
		return m.Display
	}
	return nil
}

type Display struct {
	ClothItemId     int32 `protobuf:"varint,1,opt,name=clothItemId,proto3" json:"clothItemId,omitempty"`
	ClothType       int32 `protobuf:"varint,2,opt,name=clothType,proto3" json:"clothType,omitempty"`
	WeaponItemId    int32 `protobuf:"varint,3,opt,name=weaponItemId,proto3" json:"weaponItemId,omitempty"`
	WeaponType      int32 `protobuf:"varint,4,opt,name=weaponType,proto3" json:"weaponType,omitempty"`
	WingId          int32 `protobuf:"varint,5,opt,name=wingId,proto3" json:"wingId,omitempty"`
	MagicCircleLvId int32 `protobuf:"varint,6,opt,name=magicCircleLvId,proto3" json:"magicCircleLvId,omitempty"`
}

func (m *Display) Reset()                    { *m = Display{} }
func (m *Display) String() string            { return proto.CompactTextString(m) }
func (*Display) ProtoMessage()               {}
func (*Display) Descriptor() ([]byte, []int) { return fileDescriptorShabake, []int{3} }

func (m *Display) GetClothItemId() int32 {
	if m != nil {
		return m.ClothItemId
	}
	return 0
}

func (m *Display) GetClothType() int32 {
	if m != nil {
		return m.ClothType
	}
	return 0
}

func (m *Display) GetWeaponItemId() int32 {
	if m != nil {
		return m.WeaponItemId
	}
	return 0
}

func (m *Display) GetWeaponType() int32 {
	if m != nil {
		return m.WeaponType
	}
	return 0
}

func (m *Display) GetWingId() int32 {
	if m != nil {
		return m.WingId
	}
	return 0
}

func (m *Display) GetMagicCircleLvId() int32 {
	if m != nil {
		return m.MagicCircleLvId
	}
	return 0
}

func init() {
	proto.RegisterType((*GsToCcsBackGuildInfoNtf)(nil), "pbserver.GsToCcsBackGuildInfoNtf")
	proto.RegisterType((*CcsToGsBroadShaBakeFirstGuildInfo)(nil), "pbserver.CcsToGsBroadShaBakeFirstGuildInfo")
	proto.RegisterType((*Info)(nil), "pbserver.Info")
	proto.RegisterType((*Display)(nil), "pbserver.Display")
}
func (m *GsToCcsBackGuildInfoNtf) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *GsToCcsBackGuildInfoNtf) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if len(m.FirstGuildInfo) > 0 {
		dAtA[i] = 0xa
		i++
		i = encodeVarintShabake(dAtA, i, uint64(len(m.FirstGuildInfo)))
		i += copy(dAtA[i:], m.FirstGuildInfo)
	}
	if m.BenFuShaBaKe != 0 {
		dAtA[i] = 0x10
		i++
		i = encodeVarintShabake(dAtA, i, uint64(m.BenFuShaBaKe))
	}
	if m.CrossFsId != 0 {
		dAtA[i] = 0x20
		i++
		i = encodeVarintShabake(dAtA, i, uint64(m.CrossFsId))
	}
	return i, nil
}

func (m *CcsToGsBroadShaBakeFirstGuildInfo) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *CcsToGsBroadShaBakeFirstGuildInfo) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if len(m.FirstGuildInfo) > 0 {
		dAtA[i] = 0xa
		i++
		i = encodeVarintShabake(dAtA, i, uint64(len(m.FirstGuildInfo)))
		i += copy(dAtA[i:], m.FirstGuildInfo)
	}
	if m.BenFuShaBaKe != 0 {
		dAtA[i] = 0x10
		i++
		i = encodeVarintShabake(dAtA, i, uint64(m.BenFuShaBaKe))
	}
	return i, nil
}

func (m *Info) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Info) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if len(m.NickName) > 0 {
		dAtA[i] = 0xa
		i++
		i = encodeVarintShabake(dAtA, i, uint64(len(m.NickName)))
		i += copy(dAtA[i:], m.NickName)
	}
	if m.Sex != 0 {
		dAtA[i] = 0x10
		i++
		i = encodeVarintShabake(dAtA, i, uint64(m.Sex))
	}
	if m.Job != 0 {
		dAtA[i] = 0x18
		i++
		i = encodeVarintShabake(dAtA, i, uint64(m.Job))
	}
	if m.Position != 0 {
		dAtA[i] = 0x20
		i++
		i = encodeVarintShabake(dAtA, i, uint64(m.Position))
	}
	if m.Display != nil {
		dAtA[i] = 0x2a
		i++
		i = encodeVarintShabake(dAtA, i, uint64(m.Display.Size()))
		n1, err := m.Display.MarshalTo(dAtA[i:])
		if err != nil {
			return 0, err
		}
		i += n1
	}
	return i, nil
}

func (m *Display) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Display) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if m.ClothItemId != 0 {
		dAtA[i] = 0x8
		i++
		i = encodeVarintShabake(dAtA, i, uint64(m.ClothItemId))
	}
	if m.ClothType != 0 {
		dAtA[i] = 0x10
		i++
		i = encodeVarintShabake(dAtA, i, uint64(m.ClothType))
	}
	if m.WeaponItemId != 0 {
		dAtA[i] = 0x18
		i++
		i = encodeVarintShabake(dAtA, i, uint64(m.WeaponItemId))
	}
	if m.WeaponType != 0 {
		dAtA[i] = 0x20
		i++
		i = encodeVarintShabake(dAtA, i, uint64(m.WeaponType))
	}
	if m.WingId != 0 {
		dAtA[i] = 0x28
		i++
		i = encodeVarintShabake(dAtA, i, uint64(m.WingId))
	}
	if m.MagicCircleLvId != 0 {
		dAtA[i] = 0x30
		i++
		i = encodeVarintShabake(dAtA, i, uint64(m.MagicCircleLvId))
	}
	return i, nil
}

func encodeVarintShabake(dAtA []byte, offset int, v uint64) int {
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return offset + 1
}
func (m *GsToCcsBackGuildInfoNtf) Size() (n int) {
	var l int
	_ = l
	l = len(m.FirstGuildInfo)
	if l > 0 {
		n += 1 + l + sovShabake(uint64(l))
	}
	if m.BenFuShaBaKe != 0 {
		n += 1 + sovShabake(uint64(m.BenFuShaBaKe))
	}
	if m.CrossFsId != 0 {
		n += 1 + sovShabake(uint64(m.CrossFsId))
	}
	return n
}

func (m *CcsToGsBroadShaBakeFirstGuildInfo) Size() (n int) {
	var l int
	_ = l
	l = len(m.FirstGuildInfo)
	if l > 0 {
		n += 1 + l + sovShabake(uint64(l))
	}
	if m.BenFuShaBaKe != 0 {
		n += 1 + sovShabake(uint64(m.BenFuShaBaKe))
	}
	return n
}

func (m *Info) Size() (n int) {
	var l int
	_ = l
	l = len(m.NickName)
	if l > 0 {
		n += 1 + l + sovShabake(uint64(l))
	}
	if m.Sex != 0 {
		n += 1 + sovShabake(uint64(m.Sex))
	}
	if m.Job != 0 {
		n += 1 + sovShabake(uint64(m.Job))
	}
	if m.Position != 0 {
		n += 1 + sovShabake(uint64(m.Position))
	}
	if m.Display != nil {
		l = m.Display.Size()
		n += 1 + l + sovShabake(uint64(l))
	}
	return n
}

func (m *Display) Size() (n int) {
	var l int
	_ = l
	if m.ClothItemId != 0 {
		n += 1 + sovShabake(uint64(m.ClothItemId))
	}
	if m.ClothType != 0 {
		n += 1 + sovShabake(uint64(m.ClothType))
	}
	if m.WeaponItemId != 0 {
		n += 1 + sovShabake(uint64(m.WeaponItemId))
	}
	if m.WeaponType != 0 {
		n += 1 + sovShabake(uint64(m.WeaponType))
	}
	if m.WingId != 0 {
		n += 1 + sovShabake(uint64(m.WingId))
	}
	if m.MagicCircleLvId != 0 {
		n += 1 + sovShabake(uint64(m.MagicCircleLvId))
	}
	return n
}

func sovShabake(x uint64) (n int) {
	for {
		n++
		x >>= 7
		if x == 0 {
			break
		}
	}
	return n
}
func sozShabake(x uint64) (n int) {
	return sovShabake(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *GsToCcsBackGuildInfoNtf) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowShabake
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
			return fmt.Errorf("proto: GsToCcsBackGuildInfoNtf: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: GsToCcsBackGuildInfoNtf: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field FirstGuildInfo", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowShabake
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthShabake
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.FirstGuildInfo = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field BenFuShaBaKe", wireType)
			}
			m.BenFuShaBaKe = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowShabake
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.BenFuShaBaKe |= (int32(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 4:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field CrossFsId", wireType)
			}
			m.CrossFsId = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowShabake
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.CrossFsId |= (int32(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipShabake(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthShabake
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
func (m *CcsToGsBroadShaBakeFirstGuildInfo) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowShabake
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
			return fmt.Errorf("proto: CcsToGsBroadShaBakeFirstGuildInfo: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: CcsToGsBroadShaBakeFirstGuildInfo: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field FirstGuildInfo", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowShabake
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthShabake
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.FirstGuildInfo = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field BenFuShaBaKe", wireType)
			}
			m.BenFuShaBaKe = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowShabake
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.BenFuShaBaKe |= (int32(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipShabake(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthShabake
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
func (m *Info) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowShabake
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
			return fmt.Errorf("proto: Info: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Info: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field NickName", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowShabake
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthShabake
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.NickName = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Sex", wireType)
			}
			m.Sex = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowShabake
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Sex |= (int32(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Job", wireType)
			}
			m.Job = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowShabake
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Job |= (int32(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 4:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Position", wireType)
			}
			m.Position = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowShabake
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Position |= (int32(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Display", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowShabake
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
				return ErrInvalidLengthShabake
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Display == nil {
				m.Display = &Display{}
			}
			if err := m.Display.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipShabake(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthShabake
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
func (m *Display) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowShabake
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
			return fmt.Errorf("proto: Display: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Display: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field ClothItemId", wireType)
			}
			m.ClothItemId = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowShabake
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.ClothItemId |= (int32(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field ClothType", wireType)
			}
			m.ClothType = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowShabake
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.ClothType |= (int32(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field WeaponItemId", wireType)
			}
			m.WeaponItemId = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowShabake
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.WeaponItemId |= (int32(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 4:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field WeaponType", wireType)
			}
			m.WeaponType = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowShabake
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.WeaponType |= (int32(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 5:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field WingId", wireType)
			}
			m.WingId = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowShabake
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.WingId |= (int32(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 6:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field MagicCircleLvId", wireType)
			}
			m.MagicCircleLvId = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowShabake
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.MagicCircleLvId |= (int32(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipShabake(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthShabake
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
func skipShabake(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowShabake
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
					return 0, ErrIntOverflowShabake
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
					return 0, ErrIntOverflowShabake
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
				return 0, ErrInvalidLengthShabake
			}
			return iNdEx, nil
		case 3:
			for {
				var innerWire uint64
				var start int = iNdEx
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return 0, ErrIntOverflowShabake
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
				next, err := skipShabake(dAtA[start:])
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
	ErrInvalidLengthShabake = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowShabake   = fmt.Errorf("proto: integer overflow")
)

func init() { proto.RegisterFile("shabake.proto", fileDescriptorShabake) }

var fileDescriptorShabake = []byte{
	// 371 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xac, 0x92, 0x4b, 0x4e, 0xeb, 0x30,
	0x14, 0x86, 0xaf, 0x6f, 0x9b, 0x3e, 0x4e, 0xef, 0xa3, 0x78, 0x00, 0x11, 0x42, 0x51, 0xc8, 0x00,
	0x45, 0x42, 0xea, 0x00, 0x76, 0x90, 0xa2, 0x56, 0x11, 0xa8, 0x83, 0xd0, 0x0d, 0x38, 0x8e, 0xdb,
	0xba, 0x49, 0xe3, 0x28, 0x4e, 0x5b, 0x3a, 0x66, 0x01, 0x4c, 0x59, 0x12, 0x33, 0x58, 0x02, 0x2a,
	0x1b, 0x41, 0x71, 0xd2, 0xe7, 0x98, 0x99, 0xff, 0xcf, 0xff, 0x39, 0xfe, 0xcf, 0x91, 0xe1, 0xaf,
	0x9c, 0x10, 0x9f, 0x84, 0xac, 0x93, 0xa4, 0x22, 0x13, 0xb8, 0x91, 0xf8, 0x92, 0xa5, 0x0b, 0x96,
	0x5a, 0xcf, 0x08, 0xce, 0xfa, 0x72, 0x28, 0xba, 0x54, 0x3a, 0x84, 0x86, 0xfd, 0x39, 0x8f, 0x02,
	0x37, 0x1e, 0x89, 0x41, 0x36, 0xc2, 0x57, 0xf0, 0x6f, 0xc4, 0x53, 0x99, 0x6d, 0xa1, 0x8e, 0x4c,
	0x64, 0x37, 0xbd, 0x23, 0x8a, 0x2d, 0xf8, 0xe3, 0xb3, 0xb8, 0x37, 0x7f, 0x9c, 0x10, 0x87, 0xdc,
	0x33, 0xfd, 0xb7, 0x89, 0x6c, 0xcd, 0x3b, 0x60, 0xf8, 0x02, 0x9a, 0x34, 0x15, 0x52, 0xf6, 0xa4,
	0x1b, 0xe8, 0x55, 0x65, 0xd8, 0x01, 0x4b, 0xc0, 0x65, 0x97, 0xca, 0xa1, 0xe8, 0x4b, 0x27, 0x15,
	0x24, 0x50, 0x45, 0x21, 0xeb, 0x1d, 0x3e, 0xf3, 0x83, 0x71, 0xac, 0x17, 0x04, 0x55, 0x65, 0x3e,
	0x87, 0x46, 0xcc, 0x69, 0x38, 0x20, 0x33, 0x56, 0xb6, 0xdb, 0x6a, 0xdc, 0x86, 0x8a, 0x64, 0x4f,
	0x65, 0x7d, 0x7e, 0xcc, 0xc9, 0x54, 0xf8, 0x7a, 0xa5, 0x20, 0x53, 0xe1, 0xe7, 0xf5, 0x89, 0x90,
	0x3c, 0xe3, 0x22, 0x2e, 0xc7, 0xda, 0x6a, 0x7c, 0x0d, 0xf5, 0x80, 0xcb, 0x24, 0x22, 0x2b, 0x5d,
	0x33, 0x91, 0xdd, 0xba, 0x39, 0xe9, 0x6c, 0xf6, 0xde, 0xb9, 0x2b, 0x2e, 0xbc, 0x8d, 0xc3, 0x7a,
	0x47, 0x50, 0x2f, 0x21, 0x36, 0xa1, 0x45, 0x23, 0x91, 0x4d, 0xdc, 0x8c, 0xcd, 0xdc, 0x40, 0xe5,
	0xd2, 0xbc, 0x7d, 0xa4, 0xd6, 0x99, 0xcb, 0xe1, 0x2a, 0xd9, 0x0c, 0xb8, 0x03, 0xf9, 0x06, 0x96,
	0x8c, 0x24, 0x22, 0x2e, 0x1b, 0x14, 0x79, 0x0f, 0x18, 0x36, 0x00, 0x0a, 0xad, 0x5a, 0x14, 0xd1,
	0xf7, 0x08, 0x3e, 0x85, 0xda, 0x92, 0xc7, 0x63, 0x37, 0x50, 0xd9, 0x35, 0xaf, 0x54, 0xd8, 0x86,
	0xff, 0x33, 0x32, 0xe6, 0xb4, 0xcb, 0x53, 0x1a, 0xb1, 0x87, 0x85, 0x1b, 0xe8, 0x35, 0x65, 0x38,
	0xc6, 0x4e, 0xfb, 0x6d, 0x6d, 0xa0, 0x8f, 0xb5, 0x81, 0x3e, 0xd7, 0x06, 0x7a, 0xfd, 0x32, 0x7e,
	0xf9, 0x35, 0xf5, 0xfb, 0x6e, 0xbf, 0x03, 0x00, 0x00, 0xff, 0xff, 0xd9, 0x99, 0x1e, 0xef, 0x8e,
	0x02, 0x00, 0x00,
}
