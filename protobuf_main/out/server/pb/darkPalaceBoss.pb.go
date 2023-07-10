// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: darkPalaceBoss.proto

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

// 加载boss信息
type DarkPalaceLoadReq struct {
	Floor int32 `protobuf:"varint,1,opt,name=floor,proto3" json:"floor,omitempty"`
}

func (m *DarkPalaceLoadReq) Reset()                    { *m = DarkPalaceLoadReq{} }
func (m *DarkPalaceLoadReq) String() string            { return proto.CompactTextString(m) }
func (*DarkPalaceLoadReq) ProtoMessage()               {}
func (*DarkPalaceLoadReq) Descriptor() ([]byte, []int) { return fileDescriptorDarkPalaceBoss, []int{0} }

func (m *DarkPalaceLoadReq) GetFloor() int32 {
	if m != nil {
		return m.Floor
	}
	return 0
}

type DarkPalaceLoadAck struct {
	DarkPalaceBoss []*DarkPalaceBossNtf `protobuf:"bytes,1,rep,name=darkPalaceBoss" json:"darkPalaceBoss,omitempty"`
}

func (m *DarkPalaceLoadAck) Reset()                    { *m = DarkPalaceLoadAck{} }
func (m *DarkPalaceLoadAck) String() string            { return proto.CompactTextString(m) }
func (*DarkPalaceLoadAck) ProtoMessage()               {}
func (*DarkPalaceLoadAck) Descriptor() ([]byte, []int) { return fileDescriptorDarkPalaceBoss, []int{1} }

func (m *DarkPalaceLoadAck) GetDarkPalaceBoss() []*DarkPalaceBossNtf {
	if m != nil {
		return m.DarkPalaceBoss
	}
	return nil
}

// 进入战斗
type EnterDarkPalaceFightReq struct {
	StageId int32 `protobuf:"varint,1,opt,name=stageId,proto3" json:"stageId,omitempty"`
}

func (m *EnterDarkPalaceFightReq) Reset()         { *m = EnterDarkPalaceFightReq{} }
func (m *EnterDarkPalaceFightReq) String() string { return proto.CompactTextString(m) }
func (*EnterDarkPalaceFightReq) ProtoMessage()    {}
func (*EnterDarkPalaceFightReq) Descriptor() ([]byte, []int) {
	return fileDescriptorDarkPalaceBoss, []int{2}
}

func (m *EnterDarkPalaceFightReq) GetStageId() int32 {
	if m != nil {
		return m.StageId
	}
	return 0
}

type DarkPalaceFightResultNtf struct {
	StageId int32           `protobuf:"varint,1,opt,name=stageId,proto3" json:"stageId,omitempty"`
	Result  int32           `protobuf:"varint,2,opt,name=result,proto3" json:"result,omitempty"`
	DareNum int32           `protobuf:"varint,3,opt,name=dareNum,proto3" json:"dareNum,omitempty"`
	Goods   *GoodsChangeNtf `protobuf:"bytes,4,opt,name=goods" json:"goods,omitempty"`
	Winner  *BriefUserInfo  `protobuf:"bytes,5,opt,name=winner" json:"winner,omitempty"`
}

func (m *DarkPalaceFightResultNtf) Reset()         { *m = DarkPalaceFightResultNtf{} }
func (m *DarkPalaceFightResultNtf) String() string { return proto.CompactTextString(m) }
func (*DarkPalaceFightResultNtf) ProtoMessage()    {}
func (*DarkPalaceFightResultNtf) Descriptor() ([]byte, []int) {
	return fileDescriptorDarkPalaceBoss, []int{3}
}

func (m *DarkPalaceFightResultNtf) GetStageId() int32 {
	if m != nil {
		return m.StageId
	}
	return 0
}

func (m *DarkPalaceFightResultNtf) GetResult() int32 {
	if m != nil {
		return m.Result
	}
	return 0
}

func (m *DarkPalaceFightResultNtf) GetDareNum() int32 {
	if m != nil {
		return m.DareNum
	}
	return 0
}

func (m *DarkPalaceFightResultNtf) GetGoods() *GoodsChangeNtf {
	if m != nil {
		return m.Goods
	}
	return nil
}

func (m *DarkPalaceFightResultNtf) GetWinner() *BriefUserInfo {
	if m != nil {
		return m.Winner
	}
	return nil
}

// 购买次数
type DarkPalaceBuyNumReq struct {
	Use bool `protobuf:"varint,1,opt,name=use,proto3" json:"use,omitempty"`
}

func (m *DarkPalaceBuyNumReq) Reset()         { *m = DarkPalaceBuyNumReq{} }
func (m *DarkPalaceBuyNumReq) String() string { return proto.CompactTextString(m) }
func (*DarkPalaceBuyNumReq) ProtoMessage()    {}
func (*DarkPalaceBuyNumReq) Descriptor() ([]byte, []int) {
	return fileDescriptorDarkPalaceBoss, []int{4}
}

func (m *DarkPalaceBuyNumReq) GetUse() bool {
	if m != nil {
		return m.Use
	}
	return false
}

type DarkPalaceBuyNumAck struct {
	BuyNum int32 `protobuf:"varint,1,opt,name=buyNum,proto3" json:"buyNum,omitempty"`
}

func (m *DarkPalaceBuyNumAck) Reset()         { *m = DarkPalaceBuyNumAck{} }
func (m *DarkPalaceBuyNumAck) String() string { return proto.CompactTextString(m) }
func (*DarkPalaceBuyNumAck) ProtoMessage()    {}
func (*DarkPalaceBuyNumAck) Descriptor() ([]byte, []int) {
	return fileDescriptorDarkPalaceBoss, []int{5}
}

func (m *DarkPalaceBuyNumAck) GetBuyNum() int32 {
	if m != nil {
		return m.BuyNum
	}
	return 0
}

// 推送boss状态
type DarkPalaceBossNtf struct {
	StageId    int32   `protobuf:"varint,1,opt,name=stageId,proto3" json:"stageId,omitempty"`
	Blood      float32 `protobuf:"fixed32,2,opt,name=blood,proto3" json:"blood,omitempty"`
	ReliveTime int64   `protobuf:"varint,3,opt,name=reliveTime,proto3" json:"reliveTime,omitempty"`
}

func (m *DarkPalaceBossNtf) Reset()                    { *m = DarkPalaceBossNtf{} }
func (m *DarkPalaceBossNtf) String() string            { return proto.CompactTextString(m) }
func (*DarkPalaceBossNtf) ProtoMessage()               {}
func (*DarkPalaceBossNtf) Descriptor() ([]byte, []int) { return fileDescriptorDarkPalaceBoss, []int{6} }

func (m *DarkPalaceBossNtf) GetStageId() int32 {
	if m != nil {
		return m.StageId
	}
	return 0
}

func (m *DarkPalaceBossNtf) GetBlood() float32 {
	if m != nil {
		return m.Blood
	}
	return 0
}

func (m *DarkPalaceBossNtf) GetReliveTime() int64 {
	if m != nil {
		return m.ReliveTime
	}
	return 0
}

// 协助
type EnterDarkPalaceHelpFightReq struct {
}

func (m *EnterDarkPalaceHelpFightReq) Reset()         { *m = EnterDarkPalaceHelpFightReq{} }
func (m *EnterDarkPalaceHelpFightReq) String() string { return proto.CompactTextString(m) }
func (*EnterDarkPalaceHelpFightReq) ProtoMessage()    {}
func (*EnterDarkPalaceHelpFightReq) Descriptor() ([]byte, []int) {
	return fileDescriptorDarkPalaceBoss, []int{7}
}

type DarkPalaceHelpFightResultNtf struct {
}

func (m *DarkPalaceHelpFightResultNtf) Reset()         { *m = DarkPalaceHelpFightResultNtf{} }
func (m *DarkPalaceHelpFightResultNtf) String() string { return proto.CompactTextString(m) }
func (*DarkPalaceHelpFightResultNtf) ProtoMessage()    {}
func (*DarkPalaceHelpFightResultNtf) Descriptor() ([]byte, []int) {
	return fileDescriptorDarkPalaceBoss, []int{8}
}

func init() {
	proto.RegisterType((*DarkPalaceLoadReq)(nil), "pb.DarkPalaceLoadReq")
	proto.RegisterType((*DarkPalaceLoadAck)(nil), "pb.DarkPalaceLoadAck")
	proto.RegisterType((*EnterDarkPalaceFightReq)(nil), "pb.EnterDarkPalaceFightReq")
	proto.RegisterType((*DarkPalaceFightResultNtf)(nil), "pb.DarkPalaceFightResultNtf")
	proto.RegisterType((*DarkPalaceBuyNumReq)(nil), "pb.DarkPalaceBuyNumReq")
	proto.RegisterType((*DarkPalaceBuyNumAck)(nil), "pb.DarkPalaceBuyNumAck")
	proto.RegisterType((*DarkPalaceBossNtf)(nil), "pb.DarkPalaceBossNtf")
	proto.RegisterType((*EnterDarkPalaceHelpFightReq)(nil), "pb.EnterDarkPalaceHelpFightReq")
	proto.RegisterType((*DarkPalaceHelpFightResultNtf)(nil), "pb.DarkPalaceHelpFightResultNtf")
}
func (m *DarkPalaceLoadReq) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *DarkPalaceLoadReq) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if m.Floor != 0 {
		dAtA[i] = 0x8
		i++
		i = encodeVarintDarkPalaceBoss(dAtA, i, uint64(m.Floor))
	}
	return i, nil
}

func (m *DarkPalaceLoadAck) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *DarkPalaceLoadAck) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if len(m.DarkPalaceBoss) > 0 {
		for _, msg := range m.DarkPalaceBoss {
			dAtA[i] = 0xa
			i++
			i = encodeVarintDarkPalaceBoss(dAtA, i, uint64(msg.Size()))
			n, err := msg.MarshalTo(dAtA[i:])
			if err != nil {
				return 0, err
			}
			i += n
		}
	}
	return i, nil
}

func (m *EnterDarkPalaceFightReq) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *EnterDarkPalaceFightReq) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if m.StageId != 0 {
		dAtA[i] = 0x8
		i++
		i = encodeVarintDarkPalaceBoss(dAtA, i, uint64(m.StageId))
	}
	return i, nil
}

func (m *DarkPalaceFightResultNtf) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *DarkPalaceFightResultNtf) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if m.StageId != 0 {
		dAtA[i] = 0x8
		i++
		i = encodeVarintDarkPalaceBoss(dAtA, i, uint64(m.StageId))
	}
	if m.Result != 0 {
		dAtA[i] = 0x10
		i++
		i = encodeVarintDarkPalaceBoss(dAtA, i, uint64(m.Result))
	}
	if m.DareNum != 0 {
		dAtA[i] = 0x18
		i++
		i = encodeVarintDarkPalaceBoss(dAtA, i, uint64(m.DareNum))
	}
	if m.Goods != nil {
		dAtA[i] = 0x22
		i++
		i = encodeVarintDarkPalaceBoss(dAtA, i, uint64(m.Goods.Size()))
		n1, err := m.Goods.MarshalTo(dAtA[i:])
		if err != nil {
			return 0, err
		}
		i += n1
	}
	if m.Winner != nil {
		dAtA[i] = 0x2a
		i++
		i = encodeVarintDarkPalaceBoss(dAtA, i, uint64(m.Winner.Size()))
		n2, err := m.Winner.MarshalTo(dAtA[i:])
		if err != nil {
			return 0, err
		}
		i += n2
	}
	return i, nil
}

func (m *DarkPalaceBuyNumReq) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *DarkPalaceBuyNumReq) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if m.Use {
		dAtA[i] = 0x8
		i++
		if m.Use {
			dAtA[i] = 1
		} else {
			dAtA[i] = 0
		}
		i++
	}
	return i, nil
}

func (m *DarkPalaceBuyNumAck) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *DarkPalaceBuyNumAck) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if m.BuyNum != 0 {
		dAtA[i] = 0x8
		i++
		i = encodeVarintDarkPalaceBoss(dAtA, i, uint64(m.BuyNum))
	}
	return i, nil
}

func (m *DarkPalaceBossNtf) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *DarkPalaceBossNtf) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if m.StageId != 0 {
		dAtA[i] = 0x8
		i++
		i = encodeVarintDarkPalaceBoss(dAtA, i, uint64(m.StageId))
	}
	if m.Blood != 0 {
		dAtA[i] = 0x15
		i++
		binary.LittleEndian.PutUint32(dAtA[i:], uint32(math.Float32bits(float32(m.Blood))))
		i += 4
	}
	if m.ReliveTime != 0 {
		dAtA[i] = 0x18
		i++
		i = encodeVarintDarkPalaceBoss(dAtA, i, uint64(m.ReliveTime))
	}
	return i, nil
}

func (m *EnterDarkPalaceHelpFightReq) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *EnterDarkPalaceHelpFightReq) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	return i, nil
}

func (m *DarkPalaceHelpFightResultNtf) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *DarkPalaceHelpFightResultNtf) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	return i, nil
}

func encodeVarintDarkPalaceBoss(dAtA []byte, offset int, v uint64) int {
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return offset + 1
}
func (m *DarkPalaceLoadReq) Size() (n int) {
	var l int
	_ = l
	if m.Floor != 0 {
		n += 1 + sovDarkPalaceBoss(uint64(m.Floor))
	}
	return n
}

func (m *DarkPalaceLoadAck) Size() (n int) {
	var l int
	_ = l
	if len(m.DarkPalaceBoss) > 0 {
		for _, e := range m.DarkPalaceBoss {
			l = e.Size()
			n += 1 + l + sovDarkPalaceBoss(uint64(l))
		}
	}
	return n
}

func (m *EnterDarkPalaceFightReq) Size() (n int) {
	var l int
	_ = l
	if m.StageId != 0 {
		n += 1 + sovDarkPalaceBoss(uint64(m.StageId))
	}
	return n
}

func (m *DarkPalaceFightResultNtf) Size() (n int) {
	var l int
	_ = l
	if m.StageId != 0 {
		n += 1 + sovDarkPalaceBoss(uint64(m.StageId))
	}
	if m.Result != 0 {
		n += 1 + sovDarkPalaceBoss(uint64(m.Result))
	}
	if m.DareNum != 0 {
		n += 1 + sovDarkPalaceBoss(uint64(m.DareNum))
	}
	if m.Goods != nil {
		l = m.Goods.Size()
		n += 1 + l + sovDarkPalaceBoss(uint64(l))
	}
	if m.Winner != nil {
		l = m.Winner.Size()
		n += 1 + l + sovDarkPalaceBoss(uint64(l))
	}
	return n
}

func (m *DarkPalaceBuyNumReq) Size() (n int) {
	var l int
	_ = l
	if m.Use {
		n += 2
	}
	return n
}

func (m *DarkPalaceBuyNumAck) Size() (n int) {
	var l int
	_ = l
	if m.BuyNum != 0 {
		n += 1 + sovDarkPalaceBoss(uint64(m.BuyNum))
	}
	return n
}

func (m *DarkPalaceBossNtf) Size() (n int) {
	var l int
	_ = l
	if m.StageId != 0 {
		n += 1 + sovDarkPalaceBoss(uint64(m.StageId))
	}
	if m.Blood != 0 {
		n += 5
	}
	if m.ReliveTime != 0 {
		n += 1 + sovDarkPalaceBoss(uint64(m.ReliveTime))
	}
	return n
}

func (m *EnterDarkPalaceHelpFightReq) Size() (n int) {
	var l int
	_ = l
	return n
}

func (m *DarkPalaceHelpFightResultNtf) Size() (n int) {
	var l int
	_ = l
	return n
}

func sovDarkPalaceBoss(x uint64) (n int) {
	for {
		n++
		x >>= 7
		if x == 0 {
			break
		}
	}
	return n
}
func sozDarkPalaceBoss(x uint64) (n int) {
	return sovDarkPalaceBoss(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *DarkPalaceLoadReq) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowDarkPalaceBoss
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
			return fmt.Errorf("proto: DarkPalaceLoadReq: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: DarkPalaceLoadReq: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Floor", wireType)
			}
			m.Floor = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowDarkPalaceBoss
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Floor |= (int32(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipDarkPalaceBoss(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthDarkPalaceBoss
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
func (m *DarkPalaceLoadAck) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowDarkPalaceBoss
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
			return fmt.Errorf("proto: DarkPalaceLoadAck: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: DarkPalaceLoadAck: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field DarkPalaceBoss", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowDarkPalaceBoss
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
				return ErrInvalidLengthDarkPalaceBoss
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.DarkPalaceBoss = append(m.DarkPalaceBoss, &DarkPalaceBossNtf{})
			if err := m.DarkPalaceBoss[len(m.DarkPalaceBoss)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipDarkPalaceBoss(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthDarkPalaceBoss
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
func (m *EnterDarkPalaceFightReq) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowDarkPalaceBoss
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
			return fmt.Errorf("proto: EnterDarkPalaceFightReq: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: EnterDarkPalaceFightReq: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field StageId", wireType)
			}
			m.StageId = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowDarkPalaceBoss
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
			skippy, err := skipDarkPalaceBoss(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthDarkPalaceBoss
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
func (m *DarkPalaceFightResultNtf) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowDarkPalaceBoss
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
			return fmt.Errorf("proto: DarkPalaceFightResultNtf: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: DarkPalaceFightResultNtf: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field StageId", wireType)
			}
			m.StageId = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowDarkPalaceBoss
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
				return fmt.Errorf("proto: wrong wireType = %d for field Result", wireType)
			}
			m.Result = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowDarkPalaceBoss
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
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field DareNum", wireType)
			}
			m.DareNum = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowDarkPalaceBoss
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.DareNum |= (int32(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Goods", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowDarkPalaceBoss
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
				return ErrInvalidLengthDarkPalaceBoss
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
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Winner", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowDarkPalaceBoss
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
				return ErrInvalidLengthDarkPalaceBoss
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Winner == nil {
				m.Winner = &BriefUserInfo{}
			}
			if err := m.Winner.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipDarkPalaceBoss(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthDarkPalaceBoss
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
func (m *DarkPalaceBuyNumReq) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowDarkPalaceBoss
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
			return fmt.Errorf("proto: DarkPalaceBuyNumReq: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: DarkPalaceBuyNumReq: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Use", wireType)
			}
			var v int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowDarkPalaceBoss
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
			m.Use = bool(v != 0)
		default:
			iNdEx = preIndex
			skippy, err := skipDarkPalaceBoss(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthDarkPalaceBoss
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
func (m *DarkPalaceBuyNumAck) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowDarkPalaceBoss
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
			return fmt.Errorf("proto: DarkPalaceBuyNumAck: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: DarkPalaceBuyNumAck: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field BuyNum", wireType)
			}
			m.BuyNum = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowDarkPalaceBoss
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.BuyNum |= (int32(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipDarkPalaceBoss(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthDarkPalaceBoss
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
func (m *DarkPalaceBossNtf) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowDarkPalaceBoss
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
			return fmt.Errorf("proto: DarkPalaceBossNtf: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: DarkPalaceBossNtf: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field StageId", wireType)
			}
			m.StageId = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowDarkPalaceBoss
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
			if wireType != 5 {
				return fmt.Errorf("proto: wrong wireType = %d for field Blood", wireType)
			}
			var v uint32
			if (iNdEx + 4) > l {
				return io.ErrUnexpectedEOF
			}
			v = uint32(binary.LittleEndian.Uint32(dAtA[iNdEx:]))
			iNdEx += 4
			m.Blood = float32(math.Float32frombits(v))
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field ReliveTime", wireType)
			}
			m.ReliveTime = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowDarkPalaceBoss
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.ReliveTime |= (int64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipDarkPalaceBoss(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthDarkPalaceBoss
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
func (m *EnterDarkPalaceHelpFightReq) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowDarkPalaceBoss
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
			return fmt.Errorf("proto: EnterDarkPalaceHelpFightReq: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: EnterDarkPalaceHelpFightReq: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		default:
			iNdEx = preIndex
			skippy, err := skipDarkPalaceBoss(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthDarkPalaceBoss
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
func (m *DarkPalaceHelpFightResultNtf) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowDarkPalaceBoss
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
			return fmt.Errorf("proto: DarkPalaceHelpFightResultNtf: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: DarkPalaceHelpFightResultNtf: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		default:
			iNdEx = preIndex
			skippy, err := skipDarkPalaceBoss(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthDarkPalaceBoss
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
func skipDarkPalaceBoss(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowDarkPalaceBoss
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
					return 0, ErrIntOverflowDarkPalaceBoss
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
					return 0, ErrIntOverflowDarkPalaceBoss
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
				return 0, ErrInvalidLengthDarkPalaceBoss
			}
			return iNdEx, nil
		case 3:
			for {
				var innerWire uint64
				var start int = iNdEx
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return 0, ErrIntOverflowDarkPalaceBoss
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
				next, err := skipDarkPalaceBoss(dAtA[start:])
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
	ErrInvalidLengthDarkPalaceBoss = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowDarkPalaceBoss   = fmt.Errorf("proto: integer overflow")
)

func init() { proto.RegisterFile("darkPalaceBoss.proto", fileDescriptorDarkPalaceBoss) }

var fileDescriptorDarkPalaceBoss = []byte{
	// 387 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x92, 0x4b, 0xae, 0xd3, 0x30,
	0x14, 0x86, 0x71, 0x43, 0x02, 0x3a, 0x45, 0xa8, 0x35, 0xa5, 0x58, 0x3c, 0xa2, 0xca, 0x13, 0xd2,
	0x01, 0x1d, 0xb4, 0x63, 0x06, 0x94, 0xf2, 0xa8, 0x84, 0x22, 0x64, 0xc1, 0x02, 0xf2, 0x70, 0xd2,
	0xa8, 0x49, 0x1c, 0x9c, 0x04, 0xc4, 0x4e, 0xd8, 0x0b, 0x1b, 0x60, 0xc8, 0x12, 0x50, 0xd9, 0xc8,
	0x95, 0x1d, 0xf7, 0xde, 0x3e, 0xee, 0xbd, 0xb3, 0xfc, 0xe7, 0x7c, 0x47, 0x39, 0xe7, 0xff, 0x0d,
	0xa3, 0x38, 0x90, 0xdb, 0xcf, 0x41, 0x1e, 0x44, 0x7c, 0x29, 0xea, 0x7a, 0x56, 0x49, 0xd1, 0x08,
	0xdc, 0xab, 0xc2, 0xa7, 0x0f, 0x22, 0x51, 0x14, 0xa2, 0xec, 0x2a, 0x74, 0x0a, 0xc3, 0xd5, 0x25,
	0xf9, 0x49, 0x04, 0x31, 0xe3, 0xdf, 0xf0, 0x08, 0xec, 0x24, 0x17, 0x42, 0x12, 0x34, 0x41, 0x9e,
	0xcd, 0x3a, 0x41, 0xd9, 0x29, 0xfa, 0x26, 0xda, 0xe2, 0xd7, 0xf0, 0xf0, 0xf8, 0x4f, 0x04, 0x4d,
	0x2c, 0xaf, 0x3f, 0x7f, 0x3c, 0xab, 0xc2, 0xd9, 0xea, 0xa8, 0xe3, 0x37, 0x09, 0x3b, 0x81, 0xe9,
	0x02, 0x9e, 0xbc, 0x2b, 0x1b, 0x2e, 0xaf, 0xc8, 0xf7, 0x59, 0xba, 0x69, 0xd4, 0x12, 0x04, 0xee,
	0xd5, 0x4d, 0x90, 0xf2, 0x75, 0x6c, 0xd6, 0xd8, 0x4b, 0xfa, 0x1b, 0x01, 0x39, 0x1b, 0xa8, 0xdb,
	0xbc, 0xf1, 0x9b, 0xe4, 0xe6, 0x31, 0x3c, 0x06, 0x47, 0x6a, 0x8c, 0xf4, 0x74, 0xc3, 0x28, 0x35,
	0x11, 0x07, 0x92, 0xfb, 0x6d, 0x41, 0xac, 0x6e, 0xc2, 0x48, 0xec, 0x81, 0x9d, 0x0a, 0x11, 0xd7,
	0xe4, 0xee, 0x04, 0x79, 0xfd, 0x39, 0x56, 0x37, 0x7d, 0x50, 0x85, 0xb7, 0x9b, 0xa0, 0x4c, 0xb9,
	0x3a, 0xa8, 0x03, 0xf0, 0x14, 0x9c, 0x1f, 0x59, 0x59, 0x72, 0x49, 0x6c, 0x8d, 0x0e, 0x15, 0xba,
	0x94, 0x19, 0x4f, 0xbe, 0xd6, 0x5c, 0xae, 0xcb, 0x44, 0x30, 0x03, 0xd0, 0x97, 0xf0, 0xe8, 0xc0,
	0x97, 0xf6, 0xa7, 0xdf, 0x16, 0xea, 0xdc, 0x01, 0x58, 0x6d, 0xcd, 0xf5, 0xce, 0xf7, 0x99, 0xfa,
	0xa4, 0xaf, 0xce, 0x41, 0xe5, 0xf8, 0x18, 0x9c, 0x50, 0x0b, 0x73, 0x9f, 0x51, 0x34, 0x3a, 0x8c,
	0xc7, 0xf8, 0x7d, 0x8b, 0x1b, 0x23, 0xb0, 0xc3, 0x5c, 0x88, 0x58, 0x9b, 0xd1, 0x63, 0x9d, 0xc0,
	0x2e, 0x80, 0xe4, 0x79, 0xf6, 0x9d, 0x7f, 0xc9, 0x0a, 0xae, 0xed, 0xb0, 0xd8, 0x41, 0x85, 0xbe,
	0x80, 0x67, 0x27, 0x79, 0x7d, 0xe4, 0x79, 0xb5, 0xcf, 0x8c, 0xba, 0xf0, 0xfc, 0xda, 0x8e, 0x09,
	0x67, 0x39, 0xf8, 0xb3, 0x73, 0xd1, 0xdf, 0x9d, 0x8b, 0xfe, 0xed, 0x5c, 0xf4, 0xeb, 0xbf, 0x7b,
	0x27, 0x74, 0xf4, 0x33, 0x5c, 0x5c, 0x04, 0x00, 0x00, 0xff, 0xff, 0xcd, 0x39, 0x30, 0xa7, 0xb0,
	0x02, 0x00, 0x00,
}
