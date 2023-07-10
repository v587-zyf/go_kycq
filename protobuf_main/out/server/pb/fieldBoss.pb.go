// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: fieldBoss.proto

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

// 打开野外boss时
type FieldBossLoadReq struct {
	Area int32 `protobuf:"varint,1,opt,name=area,proto3" json:"area,omitempty"`
}

func (m *FieldBossLoadReq) Reset()                    { *m = FieldBossLoadReq{} }
func (m *FieldBossLoadReq) String() string            { return proto.CompactTextString(m) }
func (*FieldBossLoadReq) ProtoMessage()               {}
func (*FieldBossLoadReq) Descriptor() ([]byte, []int) { return fileDescriptorFieldBoss, []int{0} }

func (m *FieldBossLoadReq) GetArea() int32 {
	if m != nil {
		return m.Area
	}
	return 0
}

type FieldBossLoadAck struct {
	FieldBoss []*FieldBossNtf `protobuf:"bytes,1,rep,name=fieldBoss" json:"fieldBoss,omitempty"`
}

func (m *FieldBossLoadAck) Reset()                    { *m = FieldBossLoadAck{} }
func (m *FieldBossLoadAck) String() string            { return proto.CompactTextString(m) }
func (*FieldBossLoadAck) ProtoMessage()               {}
func (*FieldBossLoadAck) Descriptor() ([]byte, []int) { return fileDescriptorFieldBoss, []int{1} }

func (m *FieldBossLoadAck) GetFieldBoss() []*FieldBossNtf {
	if m != nil {
		return m.FieldBoss
	}
	return nil
}

// 战斗
type EnterFieldBossFightReq struct {
	StageId int32 `protobuf:"varint,1,opt,name=stageId,proto3" json:"stageId,omitempty"`
}

func (m *EnterFieldBossFightReq) Reset()                    { *m = EnterFieldBossFightReq{} }
func (m *EnterFieldBossFightReq) String() string            { return proto.CompactTextString(m) }
func (*EnterFieldBossFightReq) ProtoMessage()               {}
func (*EnterFieldBossFightReq) Descriptor() ([]byte, []int) { return fileDescriptorFieldBoss, []int{2} }

func (m *EnterFieldBossFightReq) GetStageId() int32 {
	if m != nil {
		return m.StageId
	}
	return 0
}

type FieldBossFightResultNtf struct {
	StageId int32           `protobuf:"varint,1,opt,name=stageId,proto3" json:"stageId,omitempty"`
	Result  int32           `protobuf:"varint,2,opt,name=result,proto3" json:"result,omitempty"`
	DareNum int32           `protobuf:"varint,3,opt,name=dareNum,proto3" json:"dareNum,omitempty"`
	Winner  *BriefUserInfo  `protobuf:"bytes,4,opt,name=winner" json:"winner,omitempty"`
	Goods   *GoodsChangeNtf `protobuf:"bytes,5,opt,name=goods" json:"goods,omitempty"`
}

func (m *FieldBossFightResultNtf) Reset()                    { *m = FieldBossFightResultNtf{} }
func (m *FieldBossFightResultNtf) String() string            { return proto.CompactTextString(m) }
func (*FieldBossFightResultNtf) ProtoMessage()               {}
func (*FieldBossFightResultNtf) Descriptor() ([]byte, []int) { return fileDescriptorFieldBoss, []int{3} }

func (m *FieldBossFightResultNtf) GetStageId() int32 {
	if m != nil {
		return m.StageId
	}
	return 0
}

func (m *FieldBossFightResultNtf) GetResult() int32 {
	if m != nil {
		return m.Result
	}
	return 0
}

func (m *FieldBossFightResultNtf) GetDareNum() int32 {
	if m != nil {
		return m.DareNum
	}
	return 0
}

func (m *FieldBossFightResultNtf) GetWinner() *BriefUserInfo {
	if m != nil {
		return m.Winner
	}
	return nil
}

func (m *FieldBossFightResultNtf) GetGoods() *GoodsChangeNtf {
	if m != nil {
		return m.Goods
	}
	return nil
}

// 购买次数
type FieldBossBuyNumReq struct {
	Use bool `protobuf:"varint,1,opt,name=use,proto3" json:"use,omitempty"`
}

func (m *FieldBossBuyNumReq) Reset()                    { *m = FieldBossBuyNumReq{} }
func (m *FieldBossBuyNumReq) String() string            { return proto.CompactTextString(m) }
func (*FieldBossBuyNumReq) ProtoMessage()               {}
func (*FieldBossBuyNumReq) Descriptor() ([]byte, []int) { return fileDescriptorFieldBoss, []int{4} }

func (m *FieldBossBuyNumReq) GetUse() bool {
	if m != nil {
		return m.Use
	}
	return false
}

type FieldBossBuyNumAck struct {
	BuyNum int32 `protobuf:"varint,1,opt,name=buyNum,proto3" json:"buyNum,omitempty"`
}

func (m *FieldBossBuyNumAck) Reset()                    { *m = FieldBossBuyNumAck{} }
func (m *FieldBossBuyNumAck) String() string            { return proto.CompactTextString(m) }
func (*FieldBossBuyNumAck) ProtoMessage()               {}
func (*FieldBossBuyNumAck) Descriptor() ([]byte, []int) { return fileDescriptorFieldBoss, []int{5} }

func (m *FieldBossBuyNumAck) GetBuyNum() int32 {
	if m != nil {
		return m.BuyNum
	}
	return 0
}

// 推送野外首领信息
type FieldBossNtf struct {
	StageId    int32   `protobuf:"varint,1,opt,name=stageId,proto3" json:"stageId,omitempty"`
	Blood      float32 `protobuf:"fixed32,2,opt,name=blood,proto3" json:"blood,omitempty"`
	ReliveTime int64   `protobuf:"varint,3,opt,name=reliveTime,proto3" json:"reliveTime,omitempty"`
	Area       int32   `protobuf:"varint,4,opt,name=area,proto3" json:"area,omitempty"`
}

func (m *FieldBossNtf) Reset()                    { *m = FieldBossNtf{} }
func (m *FieldBossNtf) String() string            { return proto.CompactTextString(m) }
func (*FieldBossNtf) ProtoMessage()               {}
func (*FieldBossNtf) Descriptor() ([]byte, []int) { return fileDescriptorFieldBoss, []int{6} }

func (m *FieldBossNtf) GetStageId() int32 {
	if m != nil {
		return m.StageId
	}
	return 0
}

func (m *FieldBossNtf) GetBlood() float32 {
	if m != nil {
		return m.Blood
	}
	return 0
}

func (m *FieldBossNtf) GetReliveTime() int64 {
	if m != nil {
		return m.ReliveTime
	}
	return 0
}

func (m *FieldBossNtf) GetArea() int32 {
	if m != nil {
		return m.Area
	}
	return 0
}

func init() {
	proto.RegisterType((*FieldBossLoadReq)(nil), "pb.FieldBossLoadReq")
	proto.RegisterType((*FieldBossLoadAck)(nil), "pb.FieldBossLoadAck")
	proto.RegisterType((*EnterFieldBossFightReq)(nil), "pb.EnterFieldBossFightReq")
	proto.RegisterType((*FieldBossFightResultNtf)(nil), "pb.FieldBossFightResultNtf")
	proto.RegisterType((*FieldBossBuyNumReq)(nil), "pb.FieldBossBuyNumReq")
	proto.RegisterType((*FieldBossBuyNumAck)(nil), "pb.FieldBossBuyNumAck")
	proto.RegisterType((*FieldBossNtf)(nil), "pb.FieldBossNtf")
}
func (m *FieldBossLoadReq) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *FieldBossLoadReq) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if m.Area != 0 {
		dAtA[i] = 0x8
		i++
		i = encodeVarintFieldBoss(dAtA, i, uint64(m.Area))
	}
	return i, nil
}

func (m *FieldBossLoadAck) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *FieldBossLoadAck) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if len(m.FieldBoss) > 0 {
		for _, msg := range m.FieldBoss {
			dAtA[i] = 0xa
			i++
			i = encodeVarintFieldBoss(dAtA, i, uint64(msg.Size()))
			n, err := msg.MarshalTo(dAtA[i:])
			if err != nil {
				return 0, err
			}
			i += n
		}
	}
	return i, nil
}

func (m *EnterFieldBossFightReq) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *EnterFieldBossFightReq) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if m.StageId != 0 {
		dAtA[i] = 0x8
		i++
		i = encodeVarintFieldBoss(dAtA, i, uint64(m.StageId))
	}
	return i, nil
}

func (m *FieldBossFightResultNtf) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *FieldBossFightResultNtf) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if m.StageId != 0 {
		dAtA[i] = 0x8
		i++
		i = encodeVarintFieldBoss(dAtA, i, uint64(m.StageId))
	}
	if m.Result != 0 {
		dAtA[i] = 0x10
		i++
		i = encodeVarintFieldBoss(dAtA, i, uint64(m.Result))
	}
	if m.DareNum != 0 {
		dAtA[i] = 0x18
		i++
		i = encodeVarintFieldBoss(dAtA, i, uint64(m.DareNum))
	}
	if m.Winner != nil {
		dAtA[i] = 0x22
		i++
		i = encodeVarintFieldBoss(dAtA, i, uint64(m.Winner.Size()))
		n1, err := m.Winner.MarshalTo(dAtA[i:])
		if err != nil {
			return 0, err
		}
		i += n1
	}
	if m.Goods != nil {
		dAtA[i] = 0x2a
		i++
		i = encodeVarintFieldBoss(dAtA, i, uint64(m.Goods.Size()))
		n2, err := m.Goods.MarshalTo(dAtA[i:])
		if err != nil {
			return 0, err
		}
		i += n2
	}
	return i, nil
}

func (m *FieldBossBuyNumReq) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *FieldBossBuyNumReq) MarshalTo(dAtA []byte) (int, error) {
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

func (m *FieldBossBuyNumAck) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *FieldBossBuyNumAck) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if m.BuyNum != 0 {
		dAtA[i] = 0x8
		i++
		i = encodeVarintFieldBoss(dAtA, i, uint64(m.BuyNum))
	}
	return i, nil
}

func (m *FieldBossNtf) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *FieldBossNtf) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if m.StageId != 0 {
		dAtA[i] = 0x8
		i++
		i = encodeVarintFieldBoss(dAtA, i, uint64(m.StageId))
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
		i = encodeVarintFieldBoss(dAtA, i, uint64(m.ReliveTime))
	}
	if m.Area != 0 {
		dAtA[i] = 0x20
		i++
		i = encodeVarintFieldBoss(dAtA, i, uint64(m.Area))
	}
	return i, nil
}

func encodeVarintFieldBoss(dAtA []byte, offset int, v uint64) int {
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return offset + 1
}
func (m *FieldBossLoadReq) Size() (n int) {
	var l int
	_ = l
	if m.Area != 0 {
		n += 1 + sovFieldBoss(uint64(m.Area))
	}
	return n
}

func (m *FieldBossLoadAck) Size() (n int) {
	var l int
	_ = l
	if len(m.FieldBoss) > 0 {
		for _, e := range m.FieldBoss {
			l = e.Size()
			n += 1 + l + sovFieldBoss(uint64(l))
		}
	}
	return n
}

func (m *EnterFieldBossFightReq) Size() (n int) {
	var l int
	_ = l
	if m.StageId != 0 {
		n += 1 + sovFieldBoss(uint64(m.StageId))
	}
	return n
}

func (m *FieldBossFightResultNtf) Size() (n int) {
	var l int
	_ = l
	if m.StageId != 0 {
		n += 1 + sovFieldBoss(uint64(m.StageId))
	}
	if m.Result != 0 {
		n += 1 + sovFieldBoss(uint64(m.Result))
	}
	if m.DareNum != 0 {
		n += 1 + sovFieldBoss(uint64(m.DareNum))
	}
	if m.Winner != nil {
		l = m.Winner.Size()
		n += 1 + l + sovFieldBoss(uint64(l))
	}
	if m.Goods != nil {
		l = m.Goods.Size()
		n += 1 + l + sovFieldBoss(uint64(l))
	}
	return n
}

func (m *FieldBossBuyNumReq) Size() (n int) {
	var l int
	_ = l
	if m.Use {
		n += 2
	}
	return n
}

func (m *FieldBossBuyNumAck) Size() (n int) {
	var l int
	_ = l
	if m.BuyNum != 0 {
		n += 1 + sovFieldBoss(uint64(m.BuyNum))
	}
	return n
}

func (m *FieldBossNtf) Size() (n int) {
	var l int
	_ = l
	if m.StageId != 0 {
		n += 1 + sovFieldBoss(uint64(m.StageId))
	}
	if m.Blood != 0 {
		n += 5
	}
	if m.ReliveTime != 0 {
		n += 1 + sovFieldBoss(uint64(m.ReliveTime))
	}
	if m.Area != 0 {
		n += 1 + sovFieldBoss(uint64(m.Area))
	}
	return n
}

func sovFieldBoss(x uint64) (n int) {
	for {
		n++
		x >>= 7
		if x == 0 {
			break
		}
	}
	return n
}
func sozFieldBoss(x uint64) (n int) {
	return sovFieldBoss(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *FieldBossLoadReq) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowFieldBoss
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
			return fmt.Errorf("proto: FieldBossLoadReq: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: FieldBossLoadReq: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Area", wireType)
			}
			m.Area = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowFieldBoss
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Area |= (int32(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipFieldBoss(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthFieldBoss
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
func (m *FieldBossLoadAck) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowFieldBoss
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
			return fmt.Errorf("proto: FieldBossLoadAck: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: FieldBossLoadAck: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field FieldBoss", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowFieldBoss
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
				return ErrInvalidLengthFieldBoss
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.FieldBoss = append(m.FieldBoss, &FieldBossNtf{})
			if err := m.FieldBoss[len(m.FieldBoss)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipFieldBoss(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthFieldBoss
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
func (m *EnterFieldBossFightReq) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowFieldBoss
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
			return fmt.Errorf("proto: EnterFieldBossFightReq: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: EnterFieldBossFightReq: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field StageId", wireType)
			}
			m.StageId = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowFieldBoss
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
			skippy, err := skipFieldBoss(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthFieldBoss
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
func (m *FieldBossFightResultNtf) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowFieldBoss
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
			return fmt.Errorf("proto: FieldBossFightResultNtf: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: FieldBossFightResultNtf: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field StageId", wireType)
			}
			m.StageId = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowFieldBoss
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
					return ErrIntOverflowFieldBoss
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
					return ErrIntOverflowFieldBoss
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
				return fmt.Errorf("proto: wrong wireType = %d for field Winner", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowFieldBoss
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
				return ErrInvalidLengthFieldBoss
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
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Goods", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowFieldBoss
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
				return ErrInvalidLengthFieldBoss
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
			skippy, err := skipFieldBoss(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthFieldBoss
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
func (m *FieldBossBuyNumReq) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowFieldBoss
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
			return fmt.Errorf("proto: FieldBossBuyNumReq: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: FieldBossBuyNumReq: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Use", wireType)
			}
			var v int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowFieldBoss
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
			skippy, err := skipFieldBoss(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthFieldBoss
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
func (m *FieldBossBuyNumAck) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowFieldBoss
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
			return fmt.Errorf("proto: FieldBossBuyNumAck: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: FieldBossBuyNumAck: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field BuyNum", wireType)
			}
			m.BuyNum = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowFieldBoss
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
			skippy, err := skipFieldBoss(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthFieldBoss
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
func (m *FieldBossNtf) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowFieldBoss
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
			return fmt.Errorf("proto: FieldBossNtf: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: FieldBossNtf: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field StageId", wireType)
			}
			m.StageId = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowFieldBoss
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
					return ErrIntOverflowFieldBoss
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
		case 4:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Area", wireType)
			}
			m.Area = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowFieldBoss
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Area |= (int32(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipFieldBoss(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthFieldBoss
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
func skipFieldBoss(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowFieldBoss
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
					return 0, ErrIntOverflowFieldBoss
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
					return 0, ErrIntOverflowFieldBoss
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
				return 0, ErrInvalidLengthFieldBoss
			}
			return iNdEx, nil
		case 3:
			for {
				var innerWire uint64
				var start int = iNdEx
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return 0, ErrIntOverflowFieldBoss
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
				next, err := skipFieldBoss(dAtA[start:])
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
	ErrInvalidLengthFieldBoss = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowFieldBoss   = fmt.Errorf("proto: integer overflow")
)

func init() { proto.RegisterFile("fieldBoss.proto", fileDescriptorFieldBoss) }

var fileDescriptorFieldBoss = []byte{
	// 363 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x92, 0x51, 0x4e, 0xab, 0x40,
	0x14, 0x86, 0xef, 0x94, 0xc2, 0xbd, 0xf7, 0xb4, 0x89, 0x38, 0x31, 0x95, 0xf8, 0x40, 0x1a, 0x1e,
	0x1a, 0x4c, 0x0c, 0x0f, 0x75, 0x05, 0x62, 0xac, 0x69, 0x62, 0x78, 0x20, 0xba, 0x00, 0x28, 0x07,
	0x4a, 0x04, 0xa6, 0x0e, 0xa0, 0x71, 0x27, 0xae, 0xc5, 0x15, 0xf8, 0xe8, 0x12, 0x4c, 0xdd, 0x88,
	0x99, 0x81, 0xd2, 0xa6, 0x46, 0xdf, 0xe6, 0x3f, 0xf3, 0x9d, 0x9c, 0xf9, 0xff, 0x33, 0x70, 0x10,
	0xa7, 0x98, 0x45, 0x2e, 0x2b, 0x4b, 0x67, 0xc5, 0x59, 0xc5, 0x68, 0x6f, 0x15, 0x9e, 0x0c, 0x17,
	0x2c, 0xcf, 0x59, 0xd1, 0x54, 0xac, 0x09, 0xe8, 0xb3, 0x0d, 0x74, 0xc3, 0x82, 0xc8, 0xc7, 0x07,
	0x4a, 0xa1, 0x1f, 0x70, 0x0c, 0x0c, 0x32, 0x26, 0xb6, 0xea, 0xcb, 0xb3, 0xe5, 0xee, 0x71, 0x17,
	0x8b, 0x7b, 0xea, 0xc0, 0xff, 0x6e, 0x80, 0x41, 0xc6, 0x8a, 0x3d, 0x98, 0xea, 0xce, 0x2a, 0x74,
	0x3a, 0xd0, 0xab, 0x62, 0x7f, 0x8b, 0x58, 0x53, 0x18, 0x5d, 0x15, 0x15, 0xf2, 0xee, 0x7e, 0x96,
	0x26, 0xcb, 0x4a, 0x4c, 0x34, 0xe0, 0x6f, 0x59, 0x05, 0x09, 0xce, 0xa3, 0x76, 0xe8, 0x46, 0x5a,
	0xaf, 0x04, 0x8e, 0xf7, 0xf9, 0xb2, 0xce, 0x2a, 0xaf, 0x8a, 0x7f, 0xee, 0xa2, 0x23, 0xd0, 0xb8,
	0xc4, 0x8c, 0x9e, 0xbc, 0x68, 0x95, 0xe8, 0x88, 0x02, 0x8e, 0x5e, 0x9d, 0x1b, 0x4a, 0xd3, 0xd1,
	0x4a, 0x7a, 0x0a, 0xda, 0x53, 0x5a, 0x14, 0xc8, 0x8d, 0xfe, 0x98, 0xd8, 0x83, 0xe9, 0xa1, 0x30,
	0xe2, 0xf2, 0x14, 0xe3, 0xbb, 0x12, 0xf9, 0xbc, 0x88, 0x99, 0xdf, 0x02, 0xd4, 0x06, 0x35, 0x61,
	0x2c, 0x2a, 0x0d, 0x55, 0x92, 0x54, 0x90, 0xd7, 0xa2, 0x70, 0xb9, 0x0c, 0x8a, 0x04, 0x85, 0xe9,
	0x06, 0xb0, 0x26, 0x40, 0xbb, 0xb7, 0xbb, 0xf5, 0xb3, 0x57, 0xe7, 0xc2, 0xac, 0x0e, 0x4a, 0x5d,
	0xa2, 0x7c, 0xf2, 0x3f, 0x5f, 0x1c, 0xad, 0xb3, 0x6f, 0x9c, 0x88, 0x77, 0x04, 0x5a, 0x28, 0x45,
	0xeb, 0xae, 0x55, 0x16, 0x87, 0xe1, 0x6e, 0xc2, 0xbf, 0xc4, 0x70, 0x04, 0x6a, 0x98, 0x31, 0x16,
	0xc9, 0x14, 0x7a, 0x7e, 0x23, 0xa8, 0x09, 0xc0, 0x31, 0x4b, 0x1f, 0xf1, 0x36, 0xcd, 0x51, 0xe6,
	0xa0, 0xf8, 0x3b, 0x95, 0x6e, 0xfd, 0xfd, 0xed, 0xfa, 0x5d, 0xfd, 0x6d, 0x6d, 0x92, 0xf7, 0xb5,
	0x49, 0x3e, 0xd6, 0x26, 0x79, 0xf9, 0x34, 0xff, 0x84, 0x9a, 0xfc, 0x3f, 0xe7, 0x5f, 0x01, 0x00,
	0x00, 0xff, 0xff, 0x91, 0xc6, 0x04, 0x81, 0x64, 0x02, 0x00, 0x00,
}