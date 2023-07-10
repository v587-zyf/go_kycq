// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: hellBoss.proto

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

// 加载信息
type HellBossLoadReq struct {
	Floor int32 `protobuf:"varint,1,opt,name=floor,proto3" json:"floor,omitempty"`
}

func (m *HellBossLoadReq) Reset()                    { *m = HellBossLoadReq{} }
func (m *HellBossLoadReq) String() string            { return proto.CompactTextString(m) }
func (*HellBossLoadReq) ProtoMessage()               {}
func (*HellBossLoadReq) Descriptor() ([]byte, []int) { return fileDescriptorHellBoss, []int{0} }

func (m *HellBossLoadReq) GetFloor() int32 {
	if m != nil {
		return m.Floor
	}
	return 0
}

type HellBossLoadAck struct {
	Floor int32          `protobuf:"varint,1,opt,name=floor,proto3" json:"floor,omitempty"`
	List  []*HellBossNtf `protobuf:"bytes,2,rep,name=list" json:"list,omitempty"`
}

func (m *HellBossLoadAck) Reset()                    { *m = HellBossLoadAck{} }
func (m *HellBossLoadAck) String() string            { return proto.CompactTextString(m) }
func (*HellBossLoadAck) ProtoMessage()               {}
func (*HellBossLoadAck) Descriptor() ([]byte, []int) { return fileDescriptorHellBoss, []int{1} }

func (m *HellBossLoadAck) GetFloor() int32 {
	if m != nil {
		return m.Floor
	}
	return 0
}

func (m *HellBossLoadAck) GetList() []*HellBossNtf {
	if m != nil {
		return m.List
	}
	return nil
}

// 购买次数
type HellBossBuyNumReq struct {
	Use    bool  `protobuf:"varint,1,opt,name=use,proto3" json:"use,omitempty"`
	BuyNum int32 `protobuf:"varint,2,opt,name=buyNum,proto3" json:"buyNum,omitempty"`
}

func (m *HellBossBuyNumReq) Reset()                    { *m = HellBossBuyNumReq{} }
func (m *HellBossBuyNumReq) String() string            { return proto.CompactTextString(m) }
func (*HellBossBuyNumReq) ProtoMessage()               {}
func (*HellBossBuyNumReq) Descriptor() ([]byte, []int) { return fileDescriptorHellBoss, []int{2} }

func (m *HellBossBuyNumReq) GetUse() bool {
	if m != nil {
		return m.Use
	}
	return false
}

func (m *HellBossBuyNumReq) GetBuyNum() int32 {
	if m != nil {
		return m.BuyNum
	}
	return 0
}

type HellBossBuyNumAck struct {
	BuyNum int32 `protobuf:"varint,1,opt,name=buyNum,proto3" json:"buyNum,omitempty"`
}

func (m *HellBossBuyNumAck) Reset()                    { *m = HellBossBuyNumAck{} }
func (m *HellBossBuyNumAck) String() string            { return proto.CompactTextString(m) }
func (*HellBossBuyNumAck) ProtoMessage()               {}
func (*HellBossBuyNumAck) Descriptor() ([]byte, []int) { return fileDescriptorHellBoss, []int{3} }

func (m *HellBossBuyNumAck) GetBuyNum() int32 {
	if m != nil {
		return m.BuyNum
	}
	return 0
}

// 推送挑战次数
type HellBossDareNumNtf struct {
	DareNum int32 `protobuf:"varint,1,opt,name=dareNum,proto3" json:"dareNum,omitempty"`
}

func (m *HellBossDareNumNtf) Reset()                    { *m = HellBossDareNumNtf{} }
func (m *HellBossDareNumNtf) String() string            { return proto.CompactTextString(m) }
func (*HellBossDareNumNtf) ProtoMessage()               {}
func (*HellBossDareNumNtf) Descriptor() ([]byte, []int) { return fileDescriptorHellBoss, []int{4} }

func (m *HellBossDareNumNtf) GetDareNum() int32 {
	if m != nil {
		return m.DareNum
	}
	return 0
}

// 战斗
type EnterHellBossFightReq struct {
	StageId int32 `protobuf:"varint,1,opt,name=stageId,proto3" json:"stageId,omitempty"`
}

func (m *EnterHellBossFightReq) Reset()                    { *m = EnterHellBossFightReq{} }
func (m *EnterHellBossFightReq) String() string            { return proto.CompactTextString(m) }
func (*EnterHellBossFightReq) ProtoMessage()               {}
func (*EnterHellBossFightReq) Descriptor() ([]byte, []int) { return fileDescriptorHellBoss, []int{5} }

func (m *EnterHellBossFightReq) GetStageId() int32 {
	if m != nil {
		return m.StageId
	}
	return 0
}

type HellBossFightResultNtf struct {
	StageId  int32           `protobuf:"varint,1,opt,name=stageId,proto3" json:"stageId,omitempty"`
	Result   int32           `protobuf:"varint,2,opt,name=result,proto3" json:"result,omitempty"`
	DareNum  int32           `protobuf:"varint,3,opt,name=dareNum,proto3" json:"dareNum,omitempty"`
	Goods    *GoodsChangeNtf `protobuf:"bytes,4,opt,name=goods" json:"goods,omitempty"`
	Winner   *BriefUserInfo  `protobuf:"bytes,5,opt,name=winner" json:"winner,omitempty"`
	IsHelper bool            `protobuf:"varint,6,opt,name=isHelper,proto3" json:"isHelper,omitempty"`
	HelpNum  int32           `protobuf:"varint,7,opt,name=helpNum,proto3" json:"helpNum,omitempty"`
}

func (m *HellBossFightResultNtf) Reset()                    { *m = HellBossFightResultNtf{} }
func (m *HellBossFightResultNtf) String() string            { return proto.CompactTextString(m) }
func (*HellBossFightResultNtf) ProtoMessage()               {}
func (*HellBossFightResultNtf) Descriptor() ([]byte, []int) { return fileDescriptorHellBoss, []int{6} }

func (m *HellBossFightResultNtf) GetStageId() int32 {
	if m != nil {
		return m.StageId
	}
	return 0
}

func (m *HellBossFightResultNtf) GetResult() int32 {
	if m != nil {
		return m.Result
	}
	return 0
}

func (m *HellBossFightResultNtf) GetDareNum() int32 {
	if m != nil {
		return m.DareNum
	}
	return 0
}

func (m *HellBossFightResultNtf) GetGoods() *GoodsChangeNtf {
	if m != nil {
		return m.Goods
	}
	return nil
}

func (m *HellBossFightResultNtf) GetWinner() *BriefUserInfo {
	if m != nil {
		return m.Winner
	}
	return nil
}

func (m *HellBossFightResultNtf) GetIsHelper() bool {
	if m != nil {
		return m.IsHelper
	}
	return false
}

func (m *HellBossFightResultNtf) GetHelpNum() int32 {
	if m != nil {
		return m.HelpNum
	}
	return 0
}

type HellBossNtf struct {
	StageId    int32   `protobuf:"varint,1,opt,name=stageId,proto3" json:"stageId,omitempty"`
	Blood      float32 `protobuf:"fixed32,2,opt,name=blood,proto3" json:"blood,omitempty"`
	ReliveTime int64   `protobuf:"varint,3,opt,name=reliveTime,proto3" json:"reliveTime,omitempty"`
}

func (m *HellBossNtf) Reset()                    { *m = HellBossNtf{} }
func (m *HellBossNtf) String() string            { return proto.CompactTextString(m) }
func (*HellBossNtf) ProtoMessage()               {}
func (*HellBossNtf) Descriptor() ([]byte, []int) { return fileDescriptorHellBoss, []int{7} }

func (m *HellBossNtf) GetStageId() int32 {
	if m != nil {
		return m.StageId
	}
	return 0
}

func (m *HellBossNtf) GetBlood() float32 {
	if m != nil {
		return m.Blood
	}
	return 0
}

func (m *HellBossNtf) GetReliveTime() int64 {
	if m != nil {
		return m.ReliveTime
	}
	return 0
}

func init() {
	proto.RegisterType((*HellBossLoadReq)(nil), "pb.HellBossLoadReq")
	proto.RegisterType((*HellBossLoadAck)(nil), "pb.HellBossLoadAck")
	proto.RegisterType((*HellBossBuyNumReq)(nil), "pb.HellBossBuyNumReq")
	proto.RegisterType((*HellBossBuyNumAck)(nil), "pb.HellBossBuyNumAck")
	proto.RegisterType((*HellBossDareNumNtf)(nil), "pb.HellBossDareNumNtf")
	proto.RegisterType((*EnterHellBossFightReq)(nil), "pb.EnterHellBossFightReq")
	proto.RegisterType((*HellBossFightResultNtf)(nil), "pb.HellBossFightResultNtf")
	proto.RegisterType((*HellBossNtf)(nil), "pb.HellBossNtf")
}
func (m *HellBossLoadReq) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *HellBossLoadReq) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if m.Floor != 0 {
		dAtA[i] = 0x8
		i++
		i = encodeVarintHellBoss(dAtA, i, uint64(m.Floor))
	}
	return i, nil
}

func (m *HellBossLoadAck) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *HellBossLoadAck) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if m.Floor != 0 {
		dAtA[i] = 0x8
		i++
		i = encodeVarintHellBoss(dAtA, i, uint64(m.Floor))
	}
	if len(m.List) > 0 {
		for _, msg := range m.List {
			dAtA[i] = 0x12
			i++
			i = encodeVarintHellBoss(dAtA, i, uint64(msg.Size()))
			n, err := msg.MarshalTo(dAtA[i:])
			if err != nil {
				return 0, err
			}
			i += n
		}
	}
	return i, nil
}

func (m *HellBossBuyNumReq) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *HellBossBuyNumReq) MarshalTo(dAtA []byte) (int, error) {
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
	if m.BuyNum != 0 {
		dAtA[i] = 0x10
		i++
		i = encodeVarintHellBoss(dAtA, i, uint64(m.BuyNum))
	}
	return i, nil
}

func (m *HellBossBuyNumAck) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *HellBossBuyNumAck) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if m.BuyNum != 0 {
		dAtA[i] = 0x8
		i++
		i = encodeVarintHellBoss(dAtA, i, uint64(m.BuyNum))
	}
	return i, nil
}

func (m *HellBossDareNumNtf) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *HellBossDareNumNtf) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if m.DareNum != 0 {
		dAtA[i] = 0x8
		i++
		i = encodeVarintHellBoss(dAtA, i, uint64(m.DareNum))
	}
	return i, nil
}

func (m *EnterHellBossFightReq) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *EnterHellBossFightReq) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if m.StageId != 0 {
		dAtA[i] = 0x8
		i++
		i = encodeVarintHellBoss(dAtA, i, uint64(m.StageId))
	}
	return i, nil
}

func (m *HellBossFightResultNtf) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *HellBossFightResultNtf) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if m.StageId != 0 {
		dAtA[i] = 0x8
		i++
		i = encodeVarintHellBoss(dAtA, i, uint64(m.StageId))
	}
	if m.Result != 0 {
		dAtA[i] = 0x10
		i++
		i = encodeVarintHellBoss(dAtA, i, uint64(m.Result))
	}
	if m.DareNum != 0 {
		dAtA[i] = 0x18
		i++
		i = encodeVarintHellBoss(dAtA, i, uint64(m.DareNum))
	}
	if m.Goods != nil {
		dAtA[i] = 0x22
		i++
		i = encodeVarintHellBoss(dAtA, i, uint64(m.Goods.Size()))
		n1, err := m.Goods.MarshalTo(dAtA[i:])
		if err != nil {
			return 0, err
		}
		i += n1
	}
	if m.Winner != nil {
		dAtA[i] = 0x2a
		i++
		i = encodeVarintHellBoss(dAtA, i, uint64(m.Winner.Size()))
		n2, err := m.Winner.MarshalTo(dAtA[i:])
		if err != nil {
			return 0, err
		}
		i += n2
	}
	if m.IsHelper {
		dAtA[i] = 0x30
		i++
		if m.IsHelper {
			dAtA[i] = 1
		} else {
			dAtA[i] = 0
		}
		i++
	}
	if m.HelpNum != 0 {
		dAtA[i] = 0x38
		i++
		i = encodeVarintHellBoss(dAtA, i, uint64(m.HelpNum))
	}
	return i, nil
}

func (m *HellBossNtf) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *HellBossNtf) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if m.StageId != 0 {
		dAtA[i] = 0x8
		i++
		i = encodeVarintHellBoss(dAtA, i, uint64(m.StageId))
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
		i = encodeVarintHellBoss(dAtA, i, uint64(m.ReliveTime))
	}
	return i, nil
}

func encodeVarintHellBoss(dAtA []byte, offset int, v uint64) int {
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return offset + 1
}
func (m *HellBossLoadReq) Size() (n int) {
	var l int
	_ = l
	if m.Floor != 0 {
		n += 1 + sovHellBoss(uint64(m.Floor))
	}
	return n
}

func (m *HellBossLoadAck) Size() (n int) {
	var l int
	_ = l
	if m.Floor != 0 {
		n += 1 + sovHellBoss(uint64(m.Floor))
	}
	if len(m.List) > 0 {
		for _, e := range m.List {
			l = e.Size()
			n += 1 + l + sovHellBoss(uint64(l))
		}
	}
	return n
}

func (m *HellBossBuyNumReq) Size() (n int) {
	var l int
	_ = l
	if m.Use {
		n += 2
	}
	if m.BuyNum != 0 {
		n += 1 + sovHellBoss(uint64(m.BuyNum))
	}
	return n
}

func (m *HellBossBuyNumAck) Size() (n int) {
	var l int
	_ = l
	if m.BuyNum != 0 {
		n += 1 + sovHellBoss(uint64(m.BuyNum))
	}
	return n
}

func (m *HellBossDareNumNtf) Size() (n int) {
	var l int
	_ = l
	if m.DareNum != 0 {
		n += 1 + sovHellBoss(uint64(m.DareNum))
	}
	return n
}

func (m *EnterHellBossFightReq) Size() (n int) {
	var l int
	_ = l
	if m.StageId != 0 {
		n += 1 + sovHellBoss(uint64(m.StageId))
	}
	return n
}

func (m *HellBossFightResultNtf) Size() (n int) {
	var l int
	_ = l
	if m.StageId != 0 {
		n += 1 + sovHellBoss(uint64(m.StageId))
	}
	if m.Result != 0 {
		n += 1 + sovHellBoss(uint64(m.Result))
	}
	if m.DareNum != 0 {
		n += 1 + sovHellBoss(uint64(m.DareNum))
	}
	if m.Goods != nil {
		l = m.Goods.Size()
		n += 1 + l + sovHellBoss(uint64(l))
	}
	if m.Winner != nil {
		l = m.Winner.Size()
		n += 1 + l + sovHellBoss(uint64(l))
	}
	if m.IsHelper {
		n += 2
	}
	if m.HelpNum != 0 {
		n += 1 + sovHellBoss(uint64(m.HelpNum))
	}
	return n
}

func (m *HellBossNtf) Size() (n int) {
	var l int
	_ = l
	if m.StageId != 0 {
		n += 1 + sovHellBoss(uint64(m.StageId))
	}
	if m.Blood != 0 {
		n += 5
	}
	if m.ReliveTime != 0 {
		n += 1 + sovHellBoss(uint64(m.ReliveTime))
	}
	return n
}

func sovHellBoss(x uint64) (n int) {
	for {
		n++
		x >>= 7
		if x == 0 {
			break
		}
	}
	return n
}
func sozHellBoss(x uint64) (n int) {
	return sovHellBoss(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *HellBossLoadReq) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowHellBoss
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
			return fmt.Errorf("proto: HellBossLoadReq: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: HellBossLoadReq: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Floor", wireType)
			}
			m.Floor = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowHellBoss
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
			skippy, err := skipHellBoss(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthHellBoss
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
func (m *HellBossLoadAck) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowHellBoss
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
			return fmt.Errorf("proto: HellBossLoadAck: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: HellBossLoadAck: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Floor", wireType)
			}
			m.Floor = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowHellBoss
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
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field List", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowHellBoss
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
				return ErrInvalidLengthHellBoss
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.List = append(m.List, &HellBossNtf{})
			if err := m.List[len(m.List)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipHellBoss(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthHellBoss
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
func (m *HellBossBuyNumReq) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowHellBoss
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
			return fmt.Errorf("proto: HellBossBuyNumReq: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: HellBossBuyNumReq: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Use", wireType)
			}
			var v int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowHellBoss
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
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field BuyNum", wireType)
			}
			m.BuyNum = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowHellBoss
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
			skippy, err := skipHellBoss(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthHellBoss
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
func (m *HellBossBuyNumAck) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowHellBoss
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
			return fmt.Errorf("proto: HellBossBuyNumAck: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: HellBossBuyNumAck: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field BuyNum", wireType)
			}
			m.BuyNum = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowHellBoss
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
			skippy, err := skipHellBoss(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthHellBoss
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
func (m *HellBossDareNumNtf) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowHellBoss
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
			return fmt.Errorf("proto: HellBossDareNumNtf: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: HellBossDareNumNtf: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field DareNum", wireType)
			}
			m.DareNum = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowHellBoss
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
		default:
			iNdEx = preIndex
			skippy, err := skipHellBoss(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthHellBoss
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
func (m *EnterHellBossFightReq) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowHellBoss
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
			return fmt.Errorf("proto: EnterHellBossFightReq: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: EnterHellBossFightReq: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field StageId", wireType)
			}
			m.StageId = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowHellBoss
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
			skippy, err := skipHellBoss(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthHellBoss
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
func (m *HellBossFightResultNtf) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowHellBoss
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
			return fmt.Errorf("proto: HellBossFightResultNtf: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: HellBossFightResultNtf: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field StageId", wireType)
			}
			m.StageId = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowHellBoss
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
					return ErrIntOverflowHellBoss
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
					return ErrIntOverflowHellBoss
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
					return ErrIntOverflowHellBoss
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
				return ErrInvalidLengthHellBoss
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
					return ErrIntOverflowHellBoss
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
				return ErrInvalidLengthHellBoss
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
		case 6:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field IsHelper", wireType)
			}
			var v int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowHellBoss
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
			m.IsHelper = bool(v != 0)
		case 7:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field HelpNum", wireType)
			}
			m.HelpNum = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowHellBoss
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.HelpNum |= (int32(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipHellBoss(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthHellBoss
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
func (m *HellBossNtf) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowHellBoss
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
			return fmt.Errorf("proto: HellBossNtf: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: HellBossNtf: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field StageId", wireType)
			}
			m.StageId = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowHellBoss
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
					return ErrIntOverflowHellBoss
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
			skippy, err := skipHellBoss(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthHellBoss
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
func skipHellBoss(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowHellBoss
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
					return 0, ErrIntOverflowHellBoss
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
					return 0, ErrIntOverflowHellBoss
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
				return 0, ErrInvalidLengthHellBoss
			}
			return iNdEx, nil
		case 3:
			for {
				var innerWire uint64
				var start int = iNdEx
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return 0, ErrIntOverflowHellBoss
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
				next, err := skipHellBoss(dAtA[start:])
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
	ErrInvalidLengthHellBoss = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowHellBoss   = fmt.Errorf("proto: integer overflow")
)

func init() { proto.RegisterFile("hellBoss.proto", fileDescriptorHellBoss) }

var fileDescriptorHellBoss = []byte{
	// 407 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x92, 0x4d, 0x6e, 0xd4, 0x40,
	0x10, 0x85, 0xe9, 0x71, 0xec, 0x89, 0x6a, 0x10, 0x49, 0x5a, 0x21, 0x6a, 0x65, 0x61, 0x8d, 0xcc,
	0x02, 0x23, 0x24, 0x4b, 0x84, 0x35, 0x0b, 0xcc, 0x5f, 0x22, 0x45, 0x5e, 0x58, 0xb0, 0x64, 0x61,
	0xc7, 0xe5, 0x1f, 0xd1, 0x76, 0x9b, 0xee, 0x36, 0x88, 0x9b, 0x70, 0x24, 0x96, 0x1c, 0x01, 0x0d,
	0x07, 0xe0, 0x0a, 0xa8, 0xdb, 0x36, 0xcc, 0x44, 0xa3, 0xec, 0xfc, 0xaa, 0xbe, 0xe7, 0x7a, 0x55,
	0x36, 0x3c, 0xa8, 0x91, 0xf3, 0x58, 0x28, 0x15, 0xf5, 0x52, 0x68, 0x41, 0x17, 0x7d, 0x7e, 0x7e,
	0xff, 0x46, 0xb4, 0xad, 0xe8, 0xc6, 0x4a, 0xf0, 0x18, 0x8e, 0x2e, 0x27, 0xe6, 0x5a, 0x64, 0x45,
	0x8a, 0x9f, 0xe9, 0x29, 0xb8, 0x25, 0x17, 0x42, 0x32, 0xb2, 0x26, 0xa1, 0x9b, 0x8e, 0x22, 0xb8,
	0xde, 0x05, 0x5f, 0xde, 0x7c, 0xda, 0x0f, 0xd2, 0x47, 0x70, 0xc0, 0x1b, 0xa5, 0xd9, 0x62, 0xed,
	0x84, 0xab, 0x8b, 0xa3, 0xa8, 0xcf, 0xa3, 0xd9, 0x98, 0xe8, 0x32, 0xb5, 0xcd, 0xe0, 0x05, 0x9c,
	0xcc, 0xc5, 0x78, 0xf8, 0x96, 0x0c, 0xad, 0x19, 0x7c, 0x0c, 0xce, 0xa0, 0xd0, 0xbe, 0xed, 0x30,
	0x35, 0x8f, 0xf4, 0x0c, 0xbc, 0xdc, 0xb6, 0xd9, 0xc2, 0x8e, 0x98, 0x54, 0xf0, 0xf4, 0xb6, 0xdd,
	0xc4, 0xf9, 0x0f, 0x93, 0x1d, 0x38, 0x02, 0x3a, 0xc3, 0xaf, 0x33, 0x89, 0xc9, 0xd0, 0x26, 0xba,
	0xa4, 0x0c, 0x96, 0xc5, 0xa8, 0x26, 0x7c, 0x96, 0xc1, 0x33, 0x78, 0xf8, 0xa6, 0xd3, 0x28, 0x67,
	0xd3, 0xdb, 0xa6, 0xaa, 0xb5, 0xc9, 0xc7, 0x60, 0xa9, 0x74, 0x56, 0xe1, 0x55, 0x31, 0x5b, 0x26,
	0x19, 0xfc, 0x21, 0x70, 0x76, 0x0b, 0x57, 0x03, 0xd7, 0xd3, 0x9c, 0xfd, 0x26, 0x93, 0x57, 0x5a,
	0x6c, 0x5e, 0x6e, 0x54, 0xdb, 0xc9, 0x9c, 0x9d, 0x64, 0x34, 0x04, 0xb7, 0x12, 0xa2, 0x50, 0xec,
	0x60, 0x4d, 0xc2, 0xd5, 0x05, 0x35, 0xb7, 0x7d, 0x67, 0x0a, 0xaf, 0xea, 0xac, 0xab, 0xd0, 0x9c,
	0x77, 0x04, 0xe8, 0x13, 0xf0, 0xbe, 0x36, 0x5d, 0x87, 0x92, 0xb9, 0x16, 0x3d, 0x31, 0x68, 0x2c,
	0x1b, 0x2c, 0x3f, 0x28, 0x94, 0x57, 0x5d, 0x29, 0xd2, 0x09, 0xa0, 0xe7, 0x70, 0xd8, 0xa8, 0x4b,
	0xe4, 0x3d, 0x4a, 0xe6, 0xd9, 0xd3, 0xff, 0xd3, 0x26, 0x4a, 0x8d, 0xbc, 0x37, 0x51, 0x96, 0x63,
	0x94, 0x49, 0x06, 0x1f, 0x61, 0xb5, 0xf5, 0x55, 0xef, 0xd8, 0xf2, 0x14, 0xdc, 0x9c, 0x0b, 0x51,
	0xd8, 0x25, 0x17, 0xe9, 0x28, 0xa8, 0x0f, 0x20, 0x91, 0x37, 0x5f, 0xf0, 0x7d, 0xd3, 0xa2, 0x5d,
	0xd3, 0x49, 0xb7, 0x2a, 0xf1, 0xf1, 0x8f, 0x8d, 0x4f, 0x7e, 0x6e, 0x7c, 0xf2, 0x6b, 0xe3, 0x93,
	0xef, 0xbf, 0xfd, 0x7b, 0xb9, 0x67, 0xff, 0xd7, 0xe7, 0x7f, 0x03, 0x00, 0x00, 0xff, 0xff, 0xd5,
	0x4a, 0xb9, 0xb0, 0xd3, 0x02, 0x00, 0x00,
}
