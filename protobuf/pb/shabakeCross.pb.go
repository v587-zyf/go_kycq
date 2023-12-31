// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: shabakeCross.proto

package pb

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

import io "io"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

type ShaBaKeInfoCrossReq struct {
}

func (m *ShaBaKeInfoCrossReq) Reset()                    { *m = ShaBaKeInfoCrossReq{} }
func (m *ShaBaKeInfoCrossReq) String() string            { return proto.CompactTextString(m) }
func (*ShaBaKeInfoCrossReq) ProtoMessage()               {}
func (*ShaBaKeInfoCrossReq) Descriptor() ([]byte, []int) { return fileDescriptorShabakeCross, []int{0} }

type ShaBaKeInfoCrossAck struct {
	WinGuildUserInfo []*Info `protobuf:"bytes,1,rep,name=WinGuildUserInfo" json:"WinGuildUserInfo,omitempty"`
	FirstGuildName   string  `protobuf:"bytes,2,opt,name=firstGuildName,proto3" json:"firstGuildName,omitempty"`
	IsEnd            int32   `protobuf:"varint,3,opt,name=isEnd,proto3" json:"isEnd,omitempty"`
	ServerId         int32   `protobuf:"varint,4,opt,name=serverId,proto3" json:"serverId,omitempty"`
}

func (m *ShaBaKeInfoCrossAck) Reset()                    { *m = ShaBaKeInfoCrossAck{} }
func (m *ShaBaKeInfoCrossAck) String() string            { return proto.CompactTextString(m) }
func (*ShaBaKeInfoCrossAck) ProtoMessage()               {}
func (*ShaBaKeInfoCrossAck) Descriptor() ([]byte, []int) { return fileDescriptorShabakeCross, []int{1} }

func (m *ShaBaKeInfoCrossAck) GetWinGuildUserInfo() []*Info {
	if m != nil {
		return m.WinGuildUserInfo
	}
	return nil
}

func (m *ShaBaKeInfoCrossAck) GetFirstGuildName() string {
	if m != nil {
		return m.FirstGuildName
	}
	return ""
}

func (m *ShaBaKeInfoCrossAck) GetIsEnd() int32 {
	if m != nil {
		return m.IsEnd
	}
	return 0
}

func (m *ShaBaKeInfoCrossAck) GetServerId() int32 {
	if m != nil {
		return m.ServerId
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
func (*Info) Descriptor() ([]byte, []int) { return fileDescriptorShabakeCross, []int{2} }

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

// 战斗
type EnterCrossShaBaKeFightReq struct {
}

func (m *EnterCrossShaBaKeFightReq) Reset()         { *m = EnterCrossShaBaKeFightReq{} }
func (m *EnterCrossShaBaKeFightReq) String() string { return proto.CompactTextString(m) }
func (*EnterCrossShaBaKeFightReq) ProtoMessage()    {}
func (*EnterCrossShaBaKeFightReq) Descriptor() ([]byte, []int) {
	return fileDescriptorShabakeCross, []int{3}
}

type EnterCrossShaBaKeFightAck struct {
	State bool `protobuf:"varint,1,opt,name=state,proto3" json:"state,omitempty"`
}

func (m *EnterCrossShaBaKeFightAck) Reset()         { *m = EnterCrossShaBaKeFightAck{} }
func (m *EnterCrossShaBaKeFightAck) String() string { return proto.CompactTextString(m) }
func (*EnterCrossShaBaKeFightAck) ProtoMessage()    {}
func (*EnterCrossShaBaKeFightAck) Descriptor() ([]byte, []int) {
	return fileDescriptorShabakeCross, []int{4}
}

func (m *EnterCrossShaBaKeFightAck) GetState() bool {
	if m != nil {
		return m.State
	}
	return false
}

type CrossShaBaKeFightEndNtf struct {
	ServerRank []*ShabakeRankScore `protobuf:"bytes,1,rep,name=serverRank" json:"serverRank,omitempty"`
}

func (m *CrossShaBaKeFightEndNtf) Reset()         { *m = CrossShaBaKeFightEndNtf{} }
func (m *CrossShaBaKeFightEndNtf) String() string { return proto.CompactTextString(m) }
func (*CrossShaBaKeFightEndNtf) ProtoMessage()    {}
func (*CrossShaBaKeFightEndNtf) Descriptor() ([]byte, []int) {
	return fileDescriptorShabakeCross, []int{5}
}

func (m *CrossShaBaKeFightEndNtf) GetServerRank() []*ShabakeRankScore {
	if m != nil {
		return m.ServerRank
	}
	return nil
}

type CrossShabakeOpenNtf struct {
	IsOpen bool `protobuf:"varint,1,opt,name=isOpen,proto3" json:"isOpen,omitempty"`
}

func (m *CrossShabakeOpenNtf) Reset()                    { *m = CrossShabakeOpenNtf{} }
func (m *CrossShabakeOpenNtf) String() string            { return proto.CompactTextString(m) }
func (*CrossShabakeOpenNtf) ProtoMessage()               {}
func (*CrossShabakeOpenNtf) Descriptor() ([]byte, []int) { return fileDescriptorShabakeCross, []int{6} }

func (m *CrossShabakeOpenNtf) GetIsOpen() bool {
	if m != nil {
		return m.IsOpen
	}
	return false
}

type ShabakeRankScore struct {
	Id    int32 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Score int32 `protobuf:"varint,2,opt,name=score,proto3" json:"score,omitempty"`
}

func (m *ShabakeRankScore) Reset()                    { *m = ShabakeRankScore{} }
func (m *ShabakeRankScore) String() string            { return proto.CompactTextString(m) }
func (*ShabakeRankScore) ProtoMessage()               {}
func (*ShabakeRankScore) Descriptor() ([]byte, []int) { return fileDescriptorShabakeCross, []int{7} }

func (m *ShabakeRankScore) GetId() int32 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *ShabakeRankScore) GetScore() int32 {
	if m != nil {
		return m.Score
	}
	return 0
}

func init() {
	proto.RegisterType((*ShaBaKeInfoCrossReq)(nil), "pb.ShaBaKeInfoCrossReq")
	proto.RegisterType((*ShaBaKeInfoCrossAck)(nil), "pb.ShaBaKeInfoCrossAck")
	proto.RegisterType((*Info)(nil), "pb.Info")
	proto.RegisterType((*EnterCrossShaBaKeFightReq)(nil), "pb.EnterCrossShaBaKeFightReq")
	proto.RegisterType((*EnterCrossShaBaKeFightAck)(nil), "pb.EnterCrossShaBaKeFightAck")
	proto.RegisterType((*CrossShaBaKeFightEndNtf)(nil), "pb.CrossShaBaKeFightEndNtf")
	proto.RegisterType((*CrossShabakeOpenNtf)(nil), "pb.CrossShabakeOpenNtf")
	proto.RegisterType((*ShabakeRankScore)(nil), "pb.ShabakeRankScore")
}
func (m *ShaBaKeInfoCrossReq) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *ShaBaKeInfoCrossReq) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	return i, nil
}

func (m *ShaBaKeInfoCrossAck) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *ShaBaKeInfoCrossAck) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if len(m.WinGuildUserInfo) > 0 {
		for _, msg := range m.WinGuildUserInfo {
			dAtA[i] = 0xa
			i++
			i = encodeVarintShabakeCross(dAtA, i, uint64(msg.Size()))
			n, err := msg.MarshalTo(dAtA[i:])
			if err != nil {
				return 0, err
			}
			i += n
		}
	}
	if len(m.FirstGuildName) > 0 {
		dAtA[i] = 0x12
		i++
		i = encodeVarintShabakeCross(dAtA, i, uint64(len(m.FirstGuildName)))
		i += copy(dAtA[i:], m.FirstGuildName)
	}
	if m.IsEnd != 0 {
		dAtA[i] = 0x18
		i++
		i = encodeVarintShabakeCross(dAtA, i, uint64(m.IsEnd))
	}
	if m.ServerId != 0 {
		dAtA[i] = 0x20
		i++
		i = encodeVarintShabakeCross(dAtA, i, uint64(m.ServerId))
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
		i = encodeVarintShabakeCross(dAtA, i, uint64(len(m.NickName)))
		i += copy(dAtA[i:], m.NickName)
	}
	if m.Sex != 0 {
		dAtA[i] = 0x10
		i++
		i = encodeVarintShabakeCross(dAtA, i, uint64(m.Sex))
	}
	if m.Job != 0 {
		dAtA[i] = 0x18
		i++
		i = encodeVarintShabakeCross(dAtA, i, uint64(m.Job))
	}
	if m.Position != 0 {
		dAtA[i] = 0x20
		i++
		i = encodeVarintShabakeCross(dAtA, i, uint64(m.Position))
	}
	if m.Display != nil {
		dAtA[i] = 0x2a
		i++
		i = encodeVarintShabakeCross(dAtA, i, uint64(m.Display.Size()))
		n1, err := m.Display.MarshalTo(dAtA[i:])
		if err != nil {
			return 0, err
		}
		i += n1
	}
	return i, nil
}

func (m *EnterCrossShaBaKeFightReq) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *EnterCrossShaBaKeFightReq) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	return i, nil
}

func (m *EnterCrossShaBaKeFightAck) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *EnterCrossShaBaKeFightAck) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if m.State {
		dAtA[i] = 0x8
		i++
		if m.State {
			dAtA[i] = 1
		} else {
			dAtA[i] = 0
		}
		i++
	}
	return i, nil
}

func (m *CrossShaBaKeFightEndNtf) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *CrossShaBaKeFightEndNtf) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if len(m.ServerRank) > 0 {
		for _, msg := range m.ServerRank {
			dAtA[i] = 0xa
			i++
			i = encodeVarintShabakeCross(dAtA, i, uint64(msg.Size()))
			n, err := msg.MarshalTo(dAtA[i:])
			if err != nil {
				return 0, err
			}
			i += n
		}
	}
	return i, nil
}

func (m *CrossShabakeOpenNtf) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *CrossShabakeOpenNtf) MarshalTo(dAtA []byte) (int, error) {
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

func (m *ShabakeRankScore) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *ShabakeRankScore) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if m.Id != 0 {
		dAtA[i] = 0x8
		i++
		i = encodeVarintShabakeCross(dAtA, i, uint64(m.Id))
	}
	if m.Score != 0 {
		dAtA[i] = 0x10
		i++
		i = encodeVarintShabakeCross(dAtA, i, uint64(m.Score))
	}
	return i, nil
}

func encodeVarintShabakeCross(dAtA []byte, offset int, v uint64) int {
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return offset + 1
}
func (m *ShaBaKeInfoCrossReq) Size() (n int) {
	var l int
	_ = l
	return n
}

func (m *ShaBaKeInfoCrossAck) Size() (n int) {
	var l int
	_ = l
	if len(m.WinGuildUserInfo) > 0 {
		for _, e := range m.WinGuildUserInfo {
			l = e.Size()
			n += 1 + l + sovShabakeCross(uint64(l))
		}
	}
	l = len(m.FirstGuildName)
	if l > 0 {
		n += 1 + l + sovShabakeCross(uint64(l))
	}
	if m.IsEnd != 0 {
		n += 1 + sovShabakeCross(uint64(m.IsEnd))
	}
	if m.ServerId != 0 {
		n += 1 + sovShabakeCross(uint64(m.ServerId))
	}
	return n
}

func (m *Info) Size() (n int) {
	var l int
	_ = l
	l = len(m.NickName)
	if l > 0 {
		n += 1 + l + sovShabakeCross(uint64(l))
	}
	if m.Sex != 0 {
		n += 1 + sovShabakeCross(uint64(m.Sex))
	}
	if m.Job != 0 {
		n += 1 + sovShabakeCross(uint64(m.Job))
	}
	if m.Position != 0 {
		n += 1 + sovShabakeCross(uint64(m.Position))
	}
	if m.Display != nil {
		l = m.Display.Size()
		n += 1 + l + sovShabakeCross(uint64(l))
	}
	return n
}

func (m *EnterCrossShaBaKeFightReq) Size() (n int) {
	var l int
	_ = l
	return n
}

func (m *EnterCrossShaBaKeFightAck) Size() (n int) {
	var l int
	_ = l
	if m.State {
		n += 2
	}
	return n
}

func (m *CrossShaBaKeFightEndNtf) Size() (n int) {
	var l int
	_ = l
	if len(m.ServerRank) > 0 {
		for _, e := range m.ServerRank {
			l = e.Size()
			n += 1 + l + sovShabakeCross(uint64(l))
		}
	}
	return n
}

func (m *CrossShabakeOpenNtf) Size() (n int) {
	var l int
	_ = l
	if m.IsOpen {
		n += 2
	}
	return n
}

func (m *ShabakeRankScore) Size() (n int) {
	var l int
	_ = l
	if m.Id != 0 {
		n += 1 + sovShabakeCross(uint64(m.Id))
	}
	if m.Score != 0 {
		n += 1 + sovShabakeCross(uint64(m.Score))
	}
	return n
}

func sovShabakeCross(x uint64) (n int) {
	for {
		n++
		x >>= 7
		if x == 0 {
			break
		}
	}
	return n
}
func sozShabakeCross(x uint64) (n int) {
	return sovShabakeCross(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *ShaBaKeInfoCrossReq) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowShabakeCross
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
			return fmt.Errorf("proto: ShaBaKeInfoCrossReq: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: ShaBaKeInfoCrossReq: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		default:
			iNdEx = preIndex
			skippy, err := skipShabakeCross(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthShabakeCross
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
func (m *ShaBaKeInfoCrossAck) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowShabakeCross
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
			return fmt.Errorf("proto: ShaBaKeInfoCrossAck: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: ShaBaKeInfoCrossAck: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field WinGuildUserInfo", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowShabakeCross
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
				return ErrInvalidLengthShabakeCross
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.WinGuildUserInfo = append(m.WinGuildUserInfo, &Info{})
			if err := m.WinGuildUserInfo[len(m.WinGuildUserInfo)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field FirstGuildName", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowShabakeCross
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
				return ErrInvalidLengthShabakeCross
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.FirstGuildName = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field IsEnd", wireType)
			}
			m.IsEnd = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowShabakeCross
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.IsEnd |= (int32(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 4:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field ServerId", wireType)
			}
			m.ServerId = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowShabakeCross
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.ServerId |= (int32(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipShabakeCross(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthShabakeCross
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
				return ErrIntOverflowShabakeCross
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
					return ErrIntOverflowShabakeCross
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
				return ErrInvalidLengthShabakeCross
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
					return ErrIntOverflowShabakeCross
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
					return ErrIntOverflowShabakeCross
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
					return ErrIntOverflowShabakeCross
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
					return ErrIntOverflowShabakeCross
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
				return ErrInvalidLengthShabakeCross
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
			skippy, err := skipShabakeCross(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthShabakeCross
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
func (m *EnterCrossShaBaKeFightReq) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowShabakeCross
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
			return fmt.Errorf("proto: EnterCrossShaBaKeFightReq: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: EnterCrossShaBaKeFightReq: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		default:
			iNdEx = preIndex
			skippy, err := skipShabakeCross(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthShabakeCross
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
func (m *EnterCrossShaBaKeFightAck) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowShabakeCross
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
			return fmt.Errorf("proto: EnterCrossShaBaKeFightAck: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: EnterCrossShaBaKeFightAck: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field State", wireType)
			}
			var v int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowShabakeCross
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
			m.State = bool(v != 0)
		default:
			iNdEx = preIndex
			skippy, err := skipShabakeCross(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthShabakeCross
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
func (m *CrossShaBaKeFightEndNtf) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowShabakeCross
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
			return fmt.Errorf("proto: CrossShaBaKeFightEndNtf: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: CrossShaBaKeFightEndNtf: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ServerRank", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowShabakeCross
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
				return ErrInvalidLengthShabakeCross
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ServerRank = append(m.ServerRank, &ShabakeRankScore{})
			if err := m.ServerRank[len(m.ServerRank)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipShabakeCross(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthShabakeCross
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
func (m *CrossShabakeOpenNtf) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowShabakeCross
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
			return fmt.Errorf("proto: CrossShabakeOpenNtf: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: CrossShabakeOpenNtf: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field IsOpen", wireType)
			}
			var v int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowShabakeCross
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
			skippy, err := skipShabakeCross(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthShabakeCross
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
func (m *ShabakeRankScore) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowShabakeCross
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
			return fmt.Errorf("proto: ShabakeRankScore: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: ShabakeRankScore: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Id", wireType)
			}
			m.Id = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowShabakeCross
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
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Score", wireType)
			}
			m.Score = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowShabakeCross
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
		default:
			iNdEx = preIndex
			skippy, err := skipShabakeCross(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthShabakeCross
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
func skipShabakeCross(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowShabakeCross
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
					return 0, ErrIntOverflowShabakeCross
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
					return 0, ErrIntOverflowShabakeCross
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
				return 0, ErrInvalidLengthShabakeCross
			}
			return iNdEx, nil
		case 3:
			for {
				var innerWire uint64
				var start int = iNdEx
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return 0, ErrIntOverflowShabakeCross
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
				next, err := skipShabakeCross(dAtA[start:])
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
	ErrInvalidLengthShabakeCross = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowShabakeCross   = fmt.Errorf("proto: integer overflow")
)

func init() { proto.RegisterFile("shabakeCross.proto", fileDescriptorShabakeCross) }

var fileDescriptorShabakeCross = []byte{
	// 402 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x74, 0x52, 0x41, 0x8e, 0xd3, 0x30,
	0x14, 0xc5, 0xe9, 0x64, 0x08, 0xbf, 0x68, 0x14, 0x79, 0x0a, 0x84, 0x22, 0x45, 0x51, 0x24, 0x50,
	0x36, 0x54, 0x62, 0x98, 0x05, 0x5b, 0x0a, 0x05, 0x55, 0x48, 0xad, 0xe4, 0x0a, 0xb1, 0x76, 0x12,
	0x97, 0x9a, 0xb4, 0x76, 0xb0, 0x0d, 0x82, 0x23, 0x70, 0x03, 0x2e, 0xc0, 0x5d, 0x58, 0x72, 0x04,
	0x54, 0x2e, 0x82, 0x6c, 0x27, 0x15, 0x6a, 0xc5, 0xce, 0xef, 0xfd, 0x97, 0xf7, 0x7e, 0x9e, 0x0d,
	0x58, 0x6f, 0x68, 0x49, 0x1b, 0xf6, 0x42, 0x49, 0xad, 0x27, 0xad, 0x92, 0x46, 0xe2, 0xa0, 0x2d,
	0xc7, 0xb7, 0x2b, 0xb9, 0xdb, 0x49, 0xe1, 0x99, 0xfc, 0x0e, 0x5c, 0xae, 0x36, 0x74, 0x4a, 0xdf,
	0xb0, 0xb9, 0x58, 0x4b, 0xa7, 0x25, 0xec, 0x63, 0xfe, 0x03, 0x9d, 0xf2, 0xcf, 0xab, 0x06, 0x5f,
	0x43, 0xfc, 0x8e, 0x8b, 0xd7, 0x9f, 0xf8, 0xb6, 0x7e, 0xab, 0x99, 0xb2, 0xb3, 0x04, 0x65, 0x83,
	0x62, 0x78, 0x15, 0x4d, 0xda, 0x72, 0x62, 0x31, 0x39, 0x51, 0xe0, 0x47, 0x70, 0xb1, 0xe6, 0x4a,
	0x1b, 0xc7, 0x2e, 0xe8, 0x8e, 0x25, 0x41, 0x86, 0x8a, 0x5b, 0xe4, 0x88, 0xc5, 0x23, 0x08, 0xb9,
	0x9e, 0x89, 0x3a, 0x19, 0x64, 0xa8, 0x08, 0x89, 0x07, 0x78, 0x0c, 0x91, 0x66, 0xea, 0x33, 0x53,
	0xf3, 0x3a, 0x39, 0x73, 0x83, 0x03, 0xce, 0xbf, 0x21, 0x38, 0x73, 0x11, 0x63, 0x88, 0x04, 0xaf,
	0x1a, 0x67, 0x8e, 0x9c, 0xf9, 0x01, 0xe3, 0x18, 0x06, 0x9a, 0x7d, 0x71, 0x99, 0x21, 0xb1, 0x47,
	0xcb, 0x7c, 0x90, 0x65, 0x17, 0x63, 0x8f, 0xf6, 0xfb, 0x56, 0x6a, 0x6e, 0xb8, 0x14, 0x7d, 0x48,
	0x8f, 0xf1, 0x43, 0xb8, 0x59, 0x73, 0xdd, 0x6e, 0xe9, 0xd7, 0x24, 0xcc, 0x50, 0x31, 0xbc, 0x1a,
	0xda, 0x7f, 0x7d, 0xe9, 0x29, 0xd2, 0xcf, 0xf2, 0x07, 0x70, 0x7f, 0x26, 0x0c, 0x53, 0xae, 0xac,
	0xae, 0xbc, 0x57, 0xfc, 0xfd, 0xc6, 0xd8, 0x42, 0x9f, 0xfc, 0x6f, 0x68, 0x5b, 0x1d, 0x41, 0xa8,
	0x0d, 0x35, 0x7e, 0xf3, 0x88, 0x78, 0x90, 0x2f, 0xe1, 0xde, 0x89, 0x7a, 0x26, 0xea, 0x85, 0x59,
	0xe3, 0x6b, 0x00, 0x5f, 0x01, 0xa1, 0xa2, 0xe9, 0x2e, 0x60, 0x64, 0x97, 0x5a, 0xf9, 0x3b, 0xb7,
	0xf4, 0xaa, 0x92, 0x8a, 0x91, 0x7f, 0x74, 0xf9, 0x63, 0xb8, 0xec, 0x0d, 0xad, 0x68, 0xd9, 0x32,
	0x61, 0xcd, 0xee, 0xc2, 0x39, 0xd7, 0x16, 0x74, 0xf1, 0x1d, 0xca, 0x9f, 0x41, 0x7c, 0x6c, 0x87,
	0x2f, 0x20, 0xe0, 0xb5, 0xd3, 0x85, 0x24, 0xe0, 0xb5, 0xdb, 0xdc, 0x0e, 0xba, 0x72, 0x3d, 0x98,
	0xc6, 0x3f, 0xf7, 0x29, 0xfa, 0xb5, 0x4f, 0xd1, 0xef, 0x7d, 0x8a, 0xbe, 0xff, 0x49, 0x6f, 0x94,
	0xe7, 0xee, 0xb5, 0x3d, 0xfd, 0x1b, 0x00, 0x00, 0xff, 0xff, 0xb8, 0xbc, 0xad, 0xf5, 0x95, 0x02,
	0x00, 0x00,
}
