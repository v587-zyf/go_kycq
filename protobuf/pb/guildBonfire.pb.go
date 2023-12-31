// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: guildBonfire.proto

package pb

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

import binary "encoding/binary"

import io "io"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

type GuildBonfireLoadReq struct {
}

func (m *GuildBonfireLoadReq) Reset()                    { *m = GuildBonfireLoadReq{} }
func (m *GuildBonfireLoadReq) String() string            { return proto.CompactTextString(m) }
func (*GuildBonfireLoadReq) ProtoMessage()               {}
func (*GuildBonfireLoadReq) Descriptor() ([]byte, []int) { return fileDescriptorGuildBonfire, []int{0} }

type GuildBonfireLoadAck struct {
	ExpAddPercent float32       `protobuf:"fixed32,1,opt,name=expAddPercent,proto3" json:"expAddPercent,omitempty"`
	PeopleList    []*WoodPeople `protobuf:"bytes,2,rep,name=peopleList" json:"peopleList,omitempty"`
}

func (m *GuildBonfireLoadAck) Reset()                    { *m = GuildBonfireLoadAck{} }
func (m *GuildBonfireLoadAck) String() string            { return proto.CompactTextString(m) }
func (*GuildBonfireLoadAck) ProtoMessage()               {}
func (*GuildBonfireLoadAck) Descriptor() ([]byte, []int) { return fileDescriptorGuildBonfire, []int{1} }

func (m *GuildBonfireLoadAck) GetExpAddPercent() float32 {
	if m != nil {
		return m.ExpAddPercent
	}
	return 0
}

func (m *GuildBonfireLoadAck) GetPeopleList() []*WoodPeople {
	if m != nil {
		return m.PeopleList
	}
	return nil
}

type GuildBonfireAddExpReq struct {
	ConsumptionType int32 `protobuf:"varint,1,opt,name=consumptionType,proto3" json:"consumptionType,omitempty"`
}

func (m *GuildBonfireAddExpReq) Reset()         { *m = GuildBonfireAddExpReq{} }
func (m *GuildBonfireAddExpReq) String() string { return proto.CompactTextString(m) }
func (*GuildBonfireAddExpReq) ProtoMessage()    {}
func (*GuildBonfireAddExpReq) Descriptor() ([]byte, []int) {
	return fileDescriptorGuildBonfire, []int{2}
}

func (m *GuildBonfireAddExpReq) GetConsumptionType() int32 {
	if m != nil {
		return m.ConsumptionType
	}
	return 0
}

type GuildBonfireAddExpAck struct {
	ExpAddPercent float32       `protobuf:"fixed32,1,opt,name=expAddPercent,proto3" json:"expAddPercent,omitempty"`
	PeopleList    []*WoodPeople `protobuf:"bytes,2,rep,name=peopleList" json:"peopleList,omitempty"`
}

func (m *GuildBonfireAddExpAck) Reset()         { *m = GuildBonfireAddExpAck{} }
func (m *GuildBonfireAddExpAck) String() string { return proto.CompactTextString(m) }
func (*GuildBonfireAddExpAck) ProtoMessage()    {}
func (*GuildBonfireAddExpAck) Descriptor() ([]byte, []int) {
	return fileDescriptorGuildBonfire, []int{3}
}

func (m *GuildBonfireAddExpAck) GetExpAddPercent() float32 {
	if m != nil {
		return m.ExpAddPercent
	}
	return 0
}

func (m *GuildBonfireAddExpAck) GetPeopleList() []*WoodPeople {
	if m != nil {
		return m.PeopleList
	}
	return nil
}

type WoodPeople struct {
	NickName string `protobuf:"bytes,1,opt,name=nickName,proto3" json:"nickName,omitempty"`
	Avatar   string `protobuf:"bytes,2,opt,name=avatar,proto3" json:"avatar,omitempty"`
	Times    int32  `protobuf:"varint,3,opt,name=times,proto3" json:"times,omitempty"`
	Types    int32  `protobuf:"varint,4,opt,name=types,proto3" json:"types,omitempty"`
}

func (m *WoodPeople) Reset()                    { *m = WoodPeople{} }
func (m *WoodPeople) String() string            { return proto.CompactTextString(m) }
func (*WoodPeople) ProtoMessage()               {}
func (*WoodPeople) Descriptor() ([]byte, []int) { return fileDescriptorGuildBonfire, []int{4} }

func (m *WoodPeople) GetNickName() string {
	if m != nil {
		return m.NickName
	}
	return ""
}

func (m *WoodPeople) GetAvatar() string {
	if m != nil {
		return m.Avatar
	}
	return ""
}

func (m *WoodPeople) GetTimes() int32 {
	if m != nil {
		return m.Times
	}
	return 0
}

func (m *WoodPeople) GetTypes() int32 {
	if m != nil {
		return m.Types
	}
	return 0
}

// 进入战斗
type EnterGuildBonfireFightReq struct {
}

func (m *EnterGuildBonfireFightReq) Reset()         { *m = EnterGuildBonfireFightReq{} }
func (m *EnterGuildBonfireFightReq) String() string { return proto.CompactTextString(m) }
func (*EnterGuildBonfireFightReq) ProtoMessage()    {}
func (*EnterGuildBonfireFightReq) Descriptor() ([]byte, []int) {
	return fileDescriptorGuildBonfire, []int{5}
}

type GuildBonfireFightNtf struct {
	Result int32 `protobuf:"varint,1,opt,name=result,proto3" json:"result,omitempty"`
}

func (m *GuildBonfireFightNtf) Reset()                    { *m = GuildBonfireFightNtf{} }
func (m *GuildBonfireFightNtf) String() string            { return proto.CompactTextString(m) }
func (*GuildBonfireFightNtf) ProtoMessage()               {}
func (*GuildBonfireFightNtf) Descriptor() ([]byte, []int) { return fileDescriptorGuildBonfire, []int{6} }

func (m *GuildBonfireFightNtf) GetResult() int32 {
	if m != nil {
		return m.Result
	}
	return 0
}

type GuildBonfireOpenStateNtf struct {
	IsOpen bool `protobuf:"varint,1,opt,name=isOpen,proto3" json:"isOpen,omitempty"`
}

func (m *GuildBonfireOpenStateNtf) Reset()         { *m = GuildBonfireOpenStateNtf{} }
func (m *GuildBonfireOpenStateNtf) String() string { return proto.CompactTextString(m) }
func (*GuildBonfireOpenStateNtf) ProtoMessage()    {}
func (*GuildBonfireOpenStateNtf) Descriptor() ([]byte, []int) {
	return fileDescriptorGuildBonfire, []int{7}
}

func (m *GuildBonfireOpenStateNtf) GetIsOpen() bool {
	if m != nil {
		return m.IsOpen
	}
	return false
}

func init() {
	proto.RegisterType((*GuildBonfireLoadReq)(nil), "pb.GuildBonfireLoadReq")
	proto.RegisterType((*GuildBonfireLoadAck)(nil), "pb.GuildBonfireLoadAck")
	proto.RegisterType((*GuildBonfireAddExpReq)(nil), "pb.GuildBonfireAddExpReq")
	proto.RegisterType((*GuildBonfireAddExpAck)(nil), "pb.GuildBonfireAddExpAck")
	proto.RegisterType((*WoodPeople)(nil), "pb.WoodPeople")
	proto.RegisterType((*EnterGuildBonfireFightReq)(nil), "pb.EnterGuildBonfireFightReq")
	proto.RegisterType((*GuildBonfireFightNtf)(nil), "pb.GuildBonfireFightNtf")
	proto.RegisterType((*GuildBonfireOpenStateNtf)(nil), "pb.GuildBonfireOpenStateNtf")
}
func (m *GuildBonfireLoadReq) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *GuildBonfireLoadReq) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	return i, nil
}

func (m *GuildBonfireLoadAck) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *GuildBonfireLoadAck) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if m.ExpAddPercent != 0 {
		dAtA[i] = 0xd
		i++
		binary.LittleEndian.PutUint32(dAtA[i:], uint32(math.Float32bits(float32(m.ExpAddPercent))))
		i += 4
	}
	if len(m.PeopleList) > 0 {
		for _, msg := range m.PeopleList {
			dAtA[i] = 0x12
			i++
			i = encodeVarintGuildBonfire(dAtA, i, uint64(msg.Size()))
			n, err := msg.MarshalTo(dAtA[i:])
			if err != nil {
				return 0, err
			}
			i += n
		}
	}
	return i, nil
}

func (m *GuildBonfireAddExpReq) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *GuildBonfireAddExpReq) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if m.ConsumptionType != 0 {
		dAtA[i] = 0x8
		i++
		i = encodeVarintGuildBonfire(dAtA, i, uint64(m.ConsumptionType))
	}
	return i, nil
}

func (m *GuildBonfireAddExpAck) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *GuildBonfireAddExpAck) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if m.ExpAddPercent != 0 {
		dAtA[i] = 0xd
		i++
		binary.LittleEndian.PutUint32(dAtA[i:], uint32(math.Float32bits(float32(m.ExpAddPercent))))
		i += 4
	}
	if len(m.PeopleList) > 0 {
		for _, msg := range m.PeopleList {
			dAtA[i] = 0x12
			i++
			i = encodeVarintGuildBonfire(dAtA, i, uint64(msg.Size()))
			n, err := msg.MarshalTo(dAtA[i:])
			if err != nil {
				return 0, err
			}
			i += n
		}
	}
	return i, nil
}

func (m *WoodPeople) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *WoodPeople) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if len(m.NickName) > 0 {
		dAtA[i] = 0xa
		i++
		i = encodeVarintGuildBonfire(dAtA, i, uint64(len(m.NickName)))
		i += copy(dAtA[i:], m.NickName)
	}
	if len(m.Avatar) > 0 {
		dAtA[i] = 0x12
		i++
		i = encodeVarintGuildBonfire(dAtA, i, uint64(len(m.Avatar)))
		i += copy(dAtA[i:], m.Avatar)
	}
	if m.Times != 0 {
		dAtA[i] = 0x18
		i++
		i = encodeVarintGuildBonfire(dAtA, i, uint64(m.Times))
	}
	if m.Types != 0 {
		dAtA[i] = 0x20
		i++
		i = encodeVarintGuildBonfire(dAtA, i, uint64(m.Types))
	}
	return i, nil
}

func (m *EnterGuildBonfireFightReq) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *EnterGuildBonfireFightReq) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	return i, nil
}

func (m *GuildBonfireFightNtf) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *GuildBonfireFightNtf) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if m.Result != 0 {
		dAtA[i] = 0x8
		i++
		i = encodeVarintGuildBonfire(dAtA, i, uint64(m.Result))
	}
	return i, nil
}

func (m *GuildBonfireOpenStateNtf) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *GuildBonfireOpenStateNtf) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if m.IsOpen {
		dAtA[i] = 0x8
		i++
		if m.IsOpen {
			dAtA[i] = 1
		} else {
			dAtA[i] = 0
		}
		i++
	}
	return i, nil
}

func encodeVarintGuildBonfire(dAtA []byte, offset int, v uint64) int {
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return offset + 1
}
func (m *GuildBonfireLoadReq) Size() (n int) {
	var l int
	_ = l
	return n
}

func (m *GuildBonfireLoadAck) Size() (n int) {
	var l int
	_ = l
	if m.ExpAddPercent != 0 {
		n += 5
	}
	if len(m.PeopleList) > 0 {
		for _, e := range m.PeopleList {
			l = e.Size()
			n += 1 + l + sovGuildBonfire(uint64(l))
		}
	}
	return n
}

func (m *GuildBonfireAddExpReq) Size() (n int) {
	var l int
	_ = l
	if m.ConsumptionType != 0 {
		n += 1 + sovGuildBonfire(uint64(m.ConsumptionType))
	}
	return n
}

func (m *GuildBonfireAddExpAck) Size() (n int) {
	var l int
	_ = l
	if m.ExpAddPercent != 0 {
		n += 5
	}
	if len(m.PeopleList) > 0 {
		for _, e := range m.PeopleList {
			l = e.Size()
			n += 1 + l + sovGuildBonfire(uint64(l))
		}
	}
	return n
}

func (m *WoodPeople) Size() (n int) {
	var l int
	_ = l
	l = len(m.NickName)
	if l > 0 {
		n += 1 + l + sovGuildBonfire(uint64(l))
	}
	l = len(m.Avatar)
	if l > 0 {
		n += 1 + l + sovGuildBonfire(uint64(l))
	}
	if m.Times != 0 {
		n += 1 + sovGuildBonfire(uint64(m.Times))
	}
	if m.Types != 0 {
		n += 1 + sovGuildBonfire(uint64(m.Types))
	}
	return n
}

func (m *EnterGuildBonfireFightReq) Size() (n int) {
	var l int
	_ = l
	return n
}

func (m *GuildBonfireFightNtf) Size() (n int) {
	var l int
	_ = l
	if m.Result != 0 {
		n += 1 + sovGuildBonfire(uint64(m.Result))
	}
	return n
}

func (m *GuildBonfireOpenStateNtf) Size() (n int) {
	var l int
	_ = l
	if m.IsOpen {
		n += 2
	}
	return n
}

func sovGuildBonfire(x uint64) (n int) {
	for {
		n++
		x >>= 7
		if x == 0 {
			break
		}
	}
	return n
}
func sozGuildBonfire(x uint64) (n int) {
	return sovGuildBonfire(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *GuildBonfireLoadReq) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowGuildBonfire
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
			return fmt.Errorf("proto: GuildBonfireLoadReq: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: GuildBonfireLoadReq: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		default:
			iNdEx = preIndex
			skippy, err := skipGuildBonfire(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthGuildBonfire
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
func (m *GuildBonfireLoadAck) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowGuildBonfire
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
			return fmt.Errorf("proto: GuildBonfireLoadAck: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: GuildBonfireLoadAck: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 5 {
				return fmt.Errorf("proto: wrong wireType = %d for field ExpAddPercent", wireType)
			}
			var v uint32
			if (iNdEx + 4) > l {
				return io.ErrUnexpectedEOF
			}
			v = uint32(binary.LittleEndian.Uint32(dAtA[iNdEx:]))
			iNdEx += 4
			m.ExpAddPercent = float32(math.Float32frombits(v))
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field PeopleList", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGuildBonfire
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
				return ErrInvalidLengthGuildBonfire
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.PeopleList = append(m.PeopleList, &WoodPeople{})
			if err := m.PeopleList[len(m.PeopleList)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipGuildBonfire(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthGuildBonfire
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
func (m *GuildBonfireAddExpReq) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowGuildBonfire
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
			return fmt.Errorf("proto: GuildBonfireAddExpReq: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: GuildBonfireAddExpReq: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field ConsumptionType", wireType)
			}
			m.ConsumptionType = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGuildBonfire
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.ConsumptionType |= (int32(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipGuildBonfire(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthGuildBonfire
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
func (m *GuildBonfireAddExpAck) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowGuildBonfire
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
			return fmt.Errorf("proto: GuildBonfireAddExpAck: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: GuildBonfireAddExpAck: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 5 {
				return fmt.Errorf("proto: wrong wireType = %d for field ExpAddPercent", wireType)
			}
			var v uint32
			if (iNdEx + 4) > l {
				return io.ErrUnexpectedEOF
			}
			v = uint32(binary.LittleEndian.Uint32(dAtA[iNdEx:]))
			iNdEx += 4
			m.ExpAddPercent = float32(math.Float32frombits(v))
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field PeopleList", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGuildBonfire
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
				return ErrInvalidLengthGuildBonfire
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.PeopleList = append(m.PeopleList, &WoodPeople{})
			if err := m.PeopleList[len(m.PeopleList)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipGuildBonfire(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthGuildBonfire
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
func (m *WoodPeople) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowGuildBonfire
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
			return fmt.Errorf("proto: WoodPeople: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: WoodPeople: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field NickName", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGuildBonfire
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
				return ErrInvalidLengthGuildBonfire
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.NickName = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Avatar", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGuildBonfire
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
				return ErrInvalidLengthGuildBonfire
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Avatar = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Times", wireType)
			}
			m.Times = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGuildBonfire
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Times |= (int32(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 4:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Types", wireType)
			}
			m.Types = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGuildBonfire
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Types |= (int32(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipGuildBonfire(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthGuildBonfire
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
func (m *EnterGuildBonfireFightReq) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowGuildBonfire
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
			return fmt.Errorf("proto: EnterGuildBonfireFightReq: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: EnterGuildBonfireFightReq: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		default:
			iNdEx = preIndex
			skippy, err := skipGuildBonfire(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthGuildBonfire
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
func (m *GuildBonfireFightNtf) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowGuildBonfire
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
			return fmt.Errorf("proto: GuildBonfireFightNtf: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: GuildBonfireFightNtf: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Result", wireType)
			}
			m.Result = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGuildBonfire
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Result |= (int32(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipGuildBonfire(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthGuildBonfire
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
func (m *GuildBonfireOpenStateNtf) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowGuildBonfire
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
			return fmt.Errorf("proto: GuildBonfireOpenStateNtf: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: GuildBonfireOpenStateNtf: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field IsOpen", wireType)
			}
			var v int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGuildBonfire
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
			m.IsOpen = bool(v != 0)
		default:
			iNdEx = preIndex
			skippy, err := skipGuildBonfire(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthGuildBonfire
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
func skipGuildBonfire(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowGuildBonfire
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
					return 0, ErrIntOverflowGuildBonfire
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
					return 0, ErrIntOverflowGuildBonfire
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
				return 0, ErrInvalidLengthGuildBonfire
			}
			return iNdEx, nil
		case 3:
			for {
				var innerWire uint64
				var start int = iNdEx
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return 0, ErrIntOverflowGuildBonfire
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
				next, err := skipGuildBonfire(dAtA[start:])
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
	ErrInvalidLengthGuildBonfire = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowGuildBonfire   = fmt.Errorf("proto: integer overflow")
)

func init() { proto.RegisterFile("guildBonfire.proto", fileDescriptorGuildBonfire) }

var fileDescriptorGuildBonfire = []byte{
	// 324 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xb4, 0x92, 0xcf, 0x4e, 0xf2, 0x40,
	0x14, 0xc5, 0xbf, 0x96, 0x0f, 0x82, 0xd7, 0xf8, 0x27, 0x23, 0x98, 0xaa, 0x49, 0x43, 0x1a, 0x17,
	0xac, 0xba, 0xc0, 0x27, 0x28, 0x09, 0xba, 0x21, 0x48, 0x46, 0x13, 0xd7, 0xa5, 0x73, 0xc1, 0x09,
	0xed, 0xcc, 0xd0, 0x0e, 0x06, 0xde, 0xc4, 0x47, 0x72, 0xe9, 0x23, 0x18, 0x7c, 0x11, 0xd3, 0x61,
	0xa2, 0x55, 0xdc, 0xba, 0x3c, 0xbf, 0x73, 0xef, 0x9c, 0x7b, 0x67, 0x06, 0xc8, 0x6c, 0xc9, 0x53,
	0xd6, 0x97, 0x62, 0xca, 0x73, 0x0c, 0x55, 0x2e, 0xb5, 0x24, 0xae, 0x9a, 0x04, 0x6d, 0x38, 0xb9,
	0xa9, 0x38, 0x43, 0x19, 0x33, 0x8a, 0x8b, 0x60, 0xbe, 0x8b, 0xa3, 0x64, 0x4e, 0x2e, 0xe1, 0x00,
	0x57, 0x2a, 0x62, 0x6c, 0x8c, 0x79, 0x82, 0x42, 0x7b, 0x4e, 0xc7, 0xe9, 0xba, 0xf4, 0x3b, 0x24,
	0x21, 0x80, 0x42, 0xa9, 0x52, 0x1c, 0xf2, 0x42, 0x7b, 0x6e, 0xa7, 0xd6, 0xdd, 0xef, 0x1d, 0x86,
	0x6a, 0x12, 0x3e, 0x48, 0xc9, 0xc6, 0xc6, 0xa1, 0x95, 0x8a, 0x20, 0x82, 0x76, 0x35, 0x2c, 0x62,
	0x6c, 0xb0, 0x52, 0x14, 0x17, 0xa4, 0x0b, 0x47, 0x89, 0x14, 0xc5, 0x32, 0x53, 0x9a, 0x4b, 0x71,
	0xbf, 0x56, 0x68, 0x02, 0xeb, 0xf4, 0x27, 0x0e, 0xb2, 0xdf, 0x8e, 0xf8, 0xbb, 0x89, 0x53, 0x80,
	0x2f, 0x87, 0x9c, 0x43, 0x53, 0xf0, 0x64, 0x3e, 0x8a, 0xb3, 0xed, 0x7c, 0x7b, 0xf4, 0x53, 0x93,
	0x53, 0x68, 0xc4, 0x4f, 0xb1, 0x8e, 0x73, 0xcf, 0x35, 0x8e, 0x55, 0xa4, 0x05, 0x75, 0xcd, 0x33,
	0x2c, 0xbc, 0x9a, 0x59, 0x68, 0x2b, 0x0c, 0x5d, 0x2b, 0x2c, 0xbc, 0xff, 0x96, 0x96, 0x22, 0xb8,
	0x80, 0xb3, 0x81, 0xd0, 0x98, 0x57, 0x37, 0xbc, 0xe6, 0xb3, 0x47, 0x5d, 0xbe, 0x54, 0x08, 0xad,
	0x1d, 0x3e, 0xd2, 0xd3, 0x32, 0x38, 0xc7, 0x62, 0x99, 0x6a, 0x7b, 0x65, 0x56, 0x05, 0x3d, 0xf0,
	0xaa, 0xf5, 0xb7, 0x0a, 0xc5, 0x9d, 0x8e, 0x35, 0xda, 0x1e, 0x5e, 0x94, 0xc4, 0xf4, 0x34, 0xa9,
	0x55, 0xfd, 0xe3, 0x97, 0x8d, 0xef, 0xbc, 0x6e, 0x7c, 0xe7, 0x6d, 0xe3, 0x3b, 0xcf, 0xef, 0xfe,
	0xbf, 0x49, 0xc3, 0xfc, 0xa0, 0xab, 0x8f, 0x00, 0x00, 0x00, 0xff, 0xff, 0x51, 0x84, 0xc4, 0xbd,
	0x57, 0x02, 0x00, 0x00,
}
