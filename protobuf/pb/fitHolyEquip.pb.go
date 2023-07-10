// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: fitHolyEquip.proto

package pb

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

import io "io"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// 合体圣装合成
type FitHolyEquipComposeReq struct {
	EquipId int32 `protobuf:"varint,1,opt,name=equipId,proto3" json:"equipId,omitempty"`
	Pos     int32 `protobuf:"varint,2,opt,name=pos,proto3" json:"pos,omitempty"`
}

func (m *FitHolyEquipComposeReq) Reset()         { *m = FitHolyEquipComposeReq{} }
func (m *FitHolyEquipComposeReq) String() string { return proto.CompactTextString(m) }
func (*FitHolyEquipComposeReq) ProtoMessage()    {}
func (*FitHolyEquipComposeReq) Descriptor() ([]byte, []int) {
	return fileDescriptorFitHolyEquip, []int{0}
}

func (m *FitHolyEquipComposeReq) GetEquipId() int32 {
	if m != nil {
		return m.EquipId
	}
	return 0
}

func (m *FitHolyEquipComposeReq) GetPos() int32 {
	if m != nil {
		return m.Pos
	}
	return 0
}

type FitHolyEquipComposeAck struct {
	SuitType     int32             `protobuf:"varint,1,opt,name=suitType,proto3" json:"suitType,omitempty"`
	FitHolyEquip *FitHolyEquipUnit `protobuf:"bytes,2,opt,name=fitHolyEquip" json:"fitHolyEquip,omitempty"`
}

func (m *FitHolyEquipComposeAck) Reset()         { *m = FitHolyEquipComposeAck{} }
func (m *FitHolyEquipComposeAck) String() string { return proto.CompactTextString(m) }
func (*FitHolyEquipComposeAck) ProtoMessage()    {}
func (*FitHolyEquipComposeAck) Descriptor() ([]byte, []int) {
	return fileDescriptorFitHolyEquip, []int{1}
}

func (m *FitHolyEquipComposeAck) GetSuitType() int32 {
	if m != nil {
		return m.SuitType
	}
	return 0
}

func (m *FitHolyEquipComposeAck) GetFitHolyEquip() *FitHolyEquipUnit {
	if m != nil {
		return m.FitHolyEquip
	}
	return nil
}

// 合体圣装分解
type FitHolyEquipDeComposeReq struct {
	BagPos int32 `protobuf:"varint,1,opt,name=bagPos,proto3" json:"bagPos,omitempty"`
}

func (m *FitHolyEquipDeComposeReq) Reset()         { *m = FitHolyEquipDeComposeReq{} }
func (m *FitHolyEquipDeComposeReq) String() string { return proto.CompactTextString(m) }
func (*FitHolyEquipDeComposeReq) ProtoMessage()    {}
func (*FitHolyEquipDeComposeReq) Descriptor() ([]byte, []int) {
	return fileDescriptorFitHolyEquip, []int{2}
}

func (m *FitHolyEquipDeComposeReq) GetBagPos() int32 {
	if m != nil {
		return m.BagPos
	}
	return 0
}

type FitHolyEquipDeComposeAck struct {
	Goods *GoodsChangeNtf `protobuf:"bytes,1,opt,name=goods" json:"goods,omitempty"`
}

func (m *FitHolyEquipDeComposeAck) Reset()         { *m = FitHolyEquipDeComposeAck{} }
func (m *FitHolyEquipDeComposeAck) String() string { return proto.CompactTextString(m) }
func (*FitHolyEquipDeComposeAck) ProtoMessage()    {}
func (*FitHolyEquipDeComposeAck) Descriptor() ([]byte, []int) {
	return fileDescriptorFitHolyEquip, []int{3}
}

func (m *FitHolyEquipDeComposeAck) GetGoods() *GoodsChangeNtf {
	if m != nil {
		return m.Goods
	}
	return nil
}

// 合体圣装穿戴
type FitHolyEquipWearReq struct {
	BagPos   int32 `protobuf:"varint,1,opt,name=bagPos,proto3" json:"bagPos,omitempty"`
	EquipPos int32 `protobuf:"varint,2,opt,name=equipPos,proto3" json:"equipPos,omitempty"`
}

func (m *FitHolyEquipWearReq) Reset()                    { *m = FitHolyEquipWearReq{} }
func (m *FitHolyEquipWearReq) String() string            { return proto.CompactTextString(m) }
func (*FitHolyEquipWearReq) ProtoMessage()               {}
func (*FitHolyEquipWearReq) Descriptor() ([]byte, []int) { return fileDescriptorFitHolyEquip, []int{4} }

func (m *FitHolyEquipWearReq) GetBagPos() int32 {
	if m != nil {
		return m.BagPos
	}
	return 0
}

func (m *FitHolyEquipWearReq) GetEquipPos() int32 {
	if m != nil {
		return m.EquipPos
	}
	return 0
}

type FitHolyEquipWearAck struct {
	SuitType     int32             `protobuf:"varint,1,opt,name=suitType,proto3" json:"suitType,omitempty"`
	FitHolyEquip *FitHolyEquipUnit `protobuf:"bytes,2,opt,name=fitHolyEquip" json:"fitHolyEquip,omitempty"`
}

func (m *FitHolyEquipWearAck) Reset()                    { *m = FitHolyEquipWearAck{} }
func (m *FitHolyEquipWearAck) String() string            { return proto.CompactTextString(m) }
func (*FitHolyEquipWearAck) ProtoMessage()               {}
func (*FitHolyEquipWearAck) Descriptor() ([]byte, []int) { return fileDescriptorFitHolyEquip, []int{5} }

func (m *FitHolyEquipWearAck) GetSuitType() int32 {
	if m != nil {
		return m.SuitType
	}
	return 0
}

func (m *FitHolyEquipWearAck) GetFitHolyEquip() *FitHolyEquipUnit {
	if m != nil {
		return m.FitHolyEquip
	}
	return nil
}

// 合体圣装卸下
type FitHolyEquipRemoveReq struct {
	Pos      int32 `protobuf:"varint,1,opt,name=pos,proto3" json:"pos,omitempty"`
	SuitType int32 `protobuf:"varint,2,opt,name=suitType,proto3" json:"suitType,omitempty"`
}

func (m *FitHolyEquipRemoveReq) Reset()         { *m = FitHolyEquipRemoveReq{} }
func (m *FitHolyEquipRemoveReq) String() string { return proto.CompactTextString(m) }
func (*FitHolyEquipRemoveReq) ProtoMessage()    {}
func (*FitHolyEquipRemoveReq) Descriptor() ([]byte, []int) {
	return fileDescriptorFitHolyEquip, []int{6}
}

func (m *FitHolyEquipRemoveReq) GetPos() int32 {
	if m != nil {
		return m.Pos
	}
	return 0
}

func (m *FitHolyEquipRemoveReq) GetSuitType() int32 {
	if m != nil {
		return m.SuitType
	}
	return 0
}

type FitHolyEquipRemoveAck struct {
	SuitType     int32             `protobuf:"varint,1,opt,name=suitType,proto3" json:"suitType,omitempty"`
	FitHolyEquip *FitHolyEquipUnit `protobuf:"bytes,2,opt,name=fitHolyEquip" json:"fitHolyEquip,omitempty"`
}

func (m *FitHolyEquipRemoveAck) Reset()         { *m = FitHolyEquipRemoveAck{} }
func (m *FitHolyEquipRemoveAck) String() string { return proto.CompactTextString(m) }
func (*FitHolyEquipRemoveAck) ProtoMessage()    {}
func (*FitHolyEquipRemoveAck) Descriptor() ([]byte, []int) {
	return fileDescriptorFitHolyEquip, []int{7}
}

func (m *FitHolyEquipRemoveAck) GetSuitType() int32 {
	if m != nil {
		return m.SuitType
	}
	return 0
}

func (m *FitHolyEquipRemoveAck) GetFitHolyEquip() *FitHolyEquipUnit {
	if m != nil {
		return m.FitHolyEquip
	}
	return nil
}

// 合体圣装套装技能更换
type FitHolyEquipSuitSkillChangeReq struct {
	SuitId int32 `protobuf:"varint,1,opt,name=suitId,proto3" json:"suitId,omitempty"`
}

func (m *FitHolyEquipSuitSkillChangeReq) Reset()         { *m = FitHolyEquipSuitSkillChangeReq{} }
func (m *FitHolyEquipSuitSkillChangeReq) String() string { return proto.CompactTextString(m) }
func (*FitHolyEquipSuitSkillChangeReq) ProtoMessage()    {}
func (*FitHolyEquipSuitSkillChangeReq) Descriptor() ([]byte, []int) {
	return fileDescriptorFitHolyEquip, []int{8}
}

func (m *FitHolyEquipSuitSkillChangeReq) GetSuitId() int32 {
	if m != nil {
		return m.SuitId
	}
	return 0
}

type FitHolyEquipSuitSkillChangeAck struct {
	SuitId int32 `protobuf:"varint,1,opt,name=suitId,proto3" json:"suitId,omitempty"`
}

func (m *FitHolyEquipSuitSkillChangeAck) Reset()         { *m = FitHolyEquipSuitSkillChangeAck{} }
func (m *FitHolyEquipSuitSkillChangeAck) String() string { return proto.CompactTextString(m) }
func (*FitHolyEquipSuitSkillChangeAck) ProtoMessage()    {}
func (*FitHolyEquipSuitSkillChangeAck) Descriptor() ([]byte, []int) {
	return fileDescriptorFitHolyEquip, []int{9}
}

func (m *FitHolyEquipSuitSkillChangeAck) GetSuitId() int32 {
	if m != nil {
		return m.SuitId
	}
	return 0
}

func init() {
	proto.RegisterType((*FitHolyEquipComposeReq)(nil), "pb.FitHolyEquipComposeReq")
	proto.RegisterType((*FitHolyEquipComposeAck)(nil), "pb.FitHolyEquipComposeAck")
	proto.RegisterType((*FitHolyEquipDeComposeReq)(nil), "pb.FitHolyEquipDeComposeReq")
	proto.RegisterType((*FitHolyEquipDeComposeAck)(nil), "pb.FitHolyEquipDeComposeAck")
	proto.RegisterType((*FitHolyEquipWearReq)(nil), "pb.FitHolyEquipWearReq")
	proto.RegisterType((*FitHolyEquipWearAck)(nil), "pb.FitHolyEquipWearAck")
	proto.RegisterType((*FitHolyEquipRemoveReq)(nil), "pb.FitHolyEquipRemoveReq")
	proto.RegisterType((*FitHolyEquipRemoveAck)(nil), "pb.FitHolyEquipRemoveAck")
	proto.RegisterType((*FitHolyEquipSuitSkillChangeReq)(nil), "pb.FitHolyEquipSuitSkillChangeReq")
	proto.RegisterType((*FitHolyEquipSuitSkillChangeAck)(nil), "pb.FitHolyEquipSuitSkillChangeAck")
}
func (m *FitHolyEquipComposeReq) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *FitHolyEquipComposeReq) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if m.EquipId != 0 {
		dAtA[i] = 0x8
		i++
		i = encodeVarintFitHolyEquip(dAtA, i, uint64(m.EquipId))
	}
	if m.Pos != 0 {
		dAtA[i] = 0x10
		i++
		i = encodeVarintFitHolyEquip(dAtA, i, uint64(m.Pos))
	}
	return i, nil
}

func (m *FitHolyEquipComposeAck) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *FitHolyEquipComposeAck) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if m.SuitType != 0 {
		dAtA[i] = 0x8
		i++
		i = encodeVarintFitHolyEquip(dAtA, i, uint64(m.SuitType))
	}
	if m.FitHolyEquip != nil {
		dAtA[i] = 0x12
		i++
		i = encodeVarintFitHolyEquip(dAtA, i, uint64(m.FitHolyEquip.Size()))
		n1, err := m.FitHolyEquip.MarshalTo(dAtA[i:])
		if err != nil {
			return 0, err
		}
		i += n1
	}
	return i, nil
}

func (m *FitHolyEquipDeComposeReq) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *FitHolyEquipDeComposeReq) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if m.BagPos != 0 {
		dAtA[i] = 0x8
		i++
		i = encodeVarintFitHolyEquip(dAtA, i, uint64(m.BagPos))
	}
	return i, nil
}

func (m *FitHolyEquipDeComposeAck) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *FitHolyEquipDeComposeAck) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if m.Goods != nil {
		dAtA[i] = 0xa
		i++
		i = encodeVarintFitHolyEquip(dAtA, i, uint64(m.Goods.Size()))
		n2, err := m.Goods.MarshalTo(dAtA[i:])
		if err != nil {
			return 0, err
		}
		i += n2
	}
	return i, nil
}

func (m *FitHolyEquipWearReq) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *FitHolyEquipWearReq) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if m.BagPos != 0 {
		dAtA[i] = 0x8
		i++
		i = encodeVarintFitHolyEquip(dAtA, i, uint64(m.BagPos))
	}
	if m.EquipPos != 0 {
		dAtA[i] = 0x10
		i++
		i = encodeVarintFitHolyEquip(dAtA, i, uint64(m.EquipPos))
	}
	return i, nil
}

func (m *FitHolyEquipWearAck) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *FitHolyEquipWearAck) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if m.SuitType != 0 {
		dAtA[i] = 0x8
		i++
		i = encodeVarintFitHolyEquip(dAtA, i, uint64(m.SuitType))
	}
	if m.FitHolyEquip != nil {
		dAtA[i] = 0x12
		i++
		i = encodeVarintFitHolyEquip(dAtA, i, uint64(m.FitHolyEquip.Size()))
		n3, err := m.FitHolyEquip.MarshalTo(dAtA[i:])
		if err != nil {
			return 0, err
		}
		i += n3
	}
	return i, nil
}

func (m *FitHolyEquipRemoveReq) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *FitHolyEquipRemoveReq) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if m.Pos != 0 {
		dAtA[i] = 0x8
		i++
		i = encodeVarintFitHolyEquip(dAtA, i, uint64(m.Pos))
	}
	if m.SuitType != 0 {
		dAtA[i] = 0x10
		i++
		i = encodeVarintFitHolyEquip(dAtA, i, uint64(m.SuitType))
	}
	return i, nil
}

func (m *FitHolyEquipRemoveAck) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *FitHolyEquipRemoveAck) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if m.SuitType != 0 {
		dAtA[i] = 0x8
		i++
		i = encodeVarintFitHolyEquip(dAtA, i, uint64(m.SuitType))
	}
	if m.FitHolyEquip != nil {
		dAtA[i] = 0x12
		i++
		i = encodeVarintFitHolyEquip(dAtA, i, uint64(m.FitHolyEquip.Size()))
		n4, err := m.FitHolyEquip.MarshalTo(dAtA[i:])
		if err != nil {
			return 0, err
		}
		i += n4
	}
	return i, nil
}

func (m *FitHolyEquipSuitSkillChangeReq) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *FitHolyEquipSuitSkillChangeReq) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if m.SuitId != 0 {
		dAtA[i] = 0x8
		i++
		i = encodeVarintFitHolyEquip(dAtA, i, uint64(m.SuitId))
	}
	return i, nil
}

func (m *FitHolyEquipSuitSkillChangeAck) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *FitHolyEquipSuitSkillChangeAck) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if m.SuitId != 0 {
		dAtA[i] = 0x8
		i++
		i = encodeVarintFitHolyEquip(dAtA, i, uint64(m.SuitId))
	}
	return i, nil
}

func encodeVarintFitHolyEquip(dAtA []byte, offset int, v uint64) int {
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return offset + 1
}
func (m *FitHolyEquipComposeReq) Size() (n int) {
	var l int
	_ = l
	if m.EquipId != 0 {
		n += 1 + sovFitHolyEquip(uint64(m.EquipId))
	}
	if m.Pos != 0 {
		n += 1 + sovFitHolyEquip(uint64(m.Pos))
	}
	return n
}

func (m *FitHolyEquipComposeAck) Size() (n int) {
	var l int
	_ = l
	if m.SuitType != 0 {
		n += 1 + sovFitHolyEquip(uint64(m.SuitType))
	}
	if m.FitHolyEquip != nil {
		l = m.FitHolyEquip.Size()
		n += 1 + l + sovFitHolyEquip(uint64(l))
	}
	return n
}

func (m *FitHolyEquipDeComposeReq) Size() (n int) {
	var l int
	_ = l
	if m.BagPos != 0 {
		n += 1 + sovFitHolyEquip(uint64(m.BagPos))
	}
	return n
}

func (m *FitHolyEquipDeComposeAck) Size() (n int) {
	var l int
	_ = l
	if m.Goods != nil {
		l = m.Goods.Size()
		n += 1 + l + sovFitHolyEquip(uint64(l))
	}
	return n
}

func (m *FitHolyEquipWearReq) Size() (n int) {
	var l int
	_ = l
	if m.BagPos != 0 {
		n += 1 + sovFitHolyEquip(uint64(m.BagPos))
	}
	if m.EquipPos != 0 {
		n += 1 + sovFitHolyEquip(uint64(m.EquipPos))
	}
	return n
}

func (m *FitHolyEquipWearAck) Size() (n int) {
	var l int
	_ = l
	if m.SuitType != 0 {
		n += 1 + sovFitHolyEquip(uint64(m.SuitType))
	}
	if m.FitHolyEquip != nil {
		l = m.FitHolyEquip.Size()
		n += 1 + l + sovFitHolyEquip(uint64(l))
	}
	return n
}

func (m *FitHolyEquipRemoveReq) Size() (n int) {
	var l int
	_ = l
	if m.Pos != 0 {
		n += 1 + sovFitHolyEquip(uint64(m.Pos))
	}
	if m.SuitType != 0 {
		n += 1 + sovFitHolyEquip(uint64(m.SuitType))
	}
	return n
}

func (m *FitHolyEquipRemoveAck) Size() (n int) {
	var l int
	_ = l
	if m.SuitType != 0 {
		n += 1 + sovFitHolyEquip(uint64(m.SuitType))
	}
	if m.FitHolyEquip != nil {
		l = m.FitHolyEquip.Size()
		n += 1 + l + sovFitHolyEquip(uint64(l))
	}
	return n
}

func (m *FitHolyEquipSuitSkillChangeReq) Size() (n int) {
	var l int
	_ = l
	if m.SuitId != 0 {
		n += 1 + sovFitHolyEquip(uint64(m.SuitId))
	}
	return n
}

func (m *FitHolyEquipSuitSkillChangeAck) Size() (n int) {
	var l int
	_ = l
	if m.SuitId != 0 {
		n += 1 + sovFitHolyEquip(uint64(m.SuitId))
	}
	return n
}

func sovFitHolyEquip(x uint64) (n int) {
	for {
		n++
		x >>= 7
		if x == 0 {
			break
		}
	}
	return n
}
func sozFitHolyEquip(x uint64) (n int) {
	return sovFitHolyEquip(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *FitHolyEquipComposeReq) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowFitHolyEquip
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
			return fmt.Errorf("proto: FitHolyEquipComposeReq: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: FitHolyEquipComposeReq: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field EquipId", wireType)
			}
			m.EquipId = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowFitHolyEquip
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.EquipId |= (int32(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Pos", wireType)
			}
			m.Pos = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowFitHolyEquip
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Pos |= (int32(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipFitHolyEquip(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthFitHolyEquip
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
func (m *FitHolyEquipComposeAck) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowFitHolyEquip
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
			return fmt.Errorf("proto: FitHolyEquipComposeAck: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: FitHolyEquipComposeAck: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field SuitType", wireType)
			}
			m.SuitType = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowFitHolyEquip
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.SuitType |= (int32(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field FitHolyEquip", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowFitHolyEquip
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
				return ErrInvalidLengthFitHolyEquip
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.FitHolyEquip == nil {
				m.FitHolyEquip = &FitHolyEquipUnit{}
			}
			if err := m.FitHolyEquip.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipFitHolyEquip(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthFitHolyEquip
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
func (m *FitHolyEquipDeComposeReq) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowFitHolyEquip
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
			return fmt.Errorf("proto: FitHolyEquipDeComposeReq: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: FitHolyEquipDeComposeReq: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field BagPos", wireType)
			}
			m.BagPos = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowFitHolyEquip
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.BagPos |= (int32(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipFitHolyEquip(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthFitHolyEquip
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
func (m *FitHolyEquipDeComposeAck) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowFitHolyEquip
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
			return fmt.Errorf("proto: FitHolyEquipDeComposeAck: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: FitHolyEquipDeComposeAck: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Goods", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowFitHolyEquip
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
				return ErrInvalidLengthFitHolyEquip
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
			skippy, err := skipFitHolyEquip(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthFitHolyEquip
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
func (m *FitHolyEquipWearReq) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowFitHolyEquip
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
			return fmt.Errorf("proto: FitHolyEquipWearReq: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: FitHolyEquipWearReq: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field BagPos", wireType)
			}
			m.BagPos = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowFitHolyEquip
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.BagPos |= (int32(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field EquipPos", wireType)
			}
			m.EquipPos = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowFitHolyEquip
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.EquipPos |= (int32(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipFitHolyEquip(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthFitHolyEquip
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
func (m *FitHolyEquipWearAck) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowFitHolyEquip
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
			return fmt.Errorf("proto: FitHolyEquipWearAck: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: FitHolyEquipWearAck: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field SuitType", wireType)
			}
			m.SuitType = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowFitHolyEquip
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.SuitType |= (int32(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field FitHolyEquip", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowFitHolyEquip
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
				return ErrInvalidLengthFitHolyEquip
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.FitHolyEquip == nil {
				m.FitHolyEquip = &FitHolyEquipUnit{}
			}
			if err := m.FitHolyEquip.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipFitHolyEquip(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthFitHolyEquip
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
func (m *FitHolyEquipRemoveReq) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowFitHolyEquip
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
			return fmt.Errorf("proto: FitHolyEquipRemoveReq: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: FitHolyEquipRemoveReq: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Pos", wireType)
			}
			m.Pos = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowFitHolyEquip
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Pos |= (int32(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field SuitType", wireType)
			}
			m.SuitType = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowFitHolyEquip
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.SuitType |= (int32(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipFitHolyEquip(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthFitHolyEquip
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
func (m *FitHolyEquipRemoveAck) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowFitHolyEquip
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
			return fmt.Errorf("proto: FitHolyEquipRemoveAck: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: FitHolyEquipRemoveAck: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field SuitType", wireType)
			}
			m.SuitType = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowFitHolyEquip
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.SuitType |= (int32(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field FitHolyEquip", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowFitHolyEquip
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
				return ErrInvalidLengthFitHolyEquip
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.FitHolyEquip == nil {
				m.FitHolyEquip = &FitHolyEquipUnit{}
			}
			if err := m.FitHolyEquip.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipFitHolyEquip(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthFitHolyEquip
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
func (m *FitHolyEquipSuitSkillChangeReq) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowFitHolyEquip
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
			return fmt.Errorf("proto: FitHolyEquipSuitSkillChangeReq: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: FitHolyEquipSuitSkillChangeReq: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field SuitId", wireType)
			}
			m.SuitId = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowFitHolyEquip
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.SuitId |= (int32(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipFitHolyEquip(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthFitHolyEquip
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
func (m *FitHolyEquipSuitSkillChangeAck) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowFitHolyEquip
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
			return fmt.Errorf("proto: FitHolyEquipSuitSkillChangeAck: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: FitHolyEquipSuitSkillChangeAck: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field SuitId", wireType)
			}
			m.SuitId = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowFitHolyEquip
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.SuitId |= (int32(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipFitHolyEquip(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthFitHolyEquip
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
func skipFitHolyEquip(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowFitHolyEquip
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
					return 0, ErrIntOverflowFitHolyEquip
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
					return 0, ErrIntOverflowFitHolyEquip
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
				return 0, ErrInvalidLengthFitHolyEquip
			}
			return iNdEx, nil
		case 3:
			for {
				var innerWire uint64
				var start int = iNdEx
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return 0, ErrIntOverflowFitHolyEquip
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
				next, err := skipFitHolyEquip(dAtA[start:])
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
	ErrInvalidLengthFitHolyEquip = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowFitHolyEquip   = fmt.Errorf("proto: integer overflow")
)

func init() { proto.RegisterFile("fitHolyEquip.proto", fileDescriptorFitHolyEquip) }

var fileDescriptorFitHolyEquip = []byte{
	// 328 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xb4, 0x53, 0xc1, 0x4e, 0xc2, 0x40,
	0x14, 0x74, 0x31, 0xa0, 0x79, 0x70, 0x20, 0xab, 0x92, 0x86, 0x43, 0x63, 0xf6, 0xc4, 0x89, 0x03,
	0x5e, 0xb8, 0x2a, 0xa0, 0x72, 0x31, 0xa4, 0x68, 0x3c, 0x53, 0x58, 0x70, 0x03, 0xed, 0x5b, 0xe8,
	0xd6, 0x84, 0x3f, 0xf1, 0x93, 0x3c, 0xfa, 0x09, 0xa6, 0xfe, 0x88, 0xd9, 0x76, 0x6d, 0xb6, 0x11,
	0xf4, 0xc4, 0xad, 0xf3, 0x32, 0x33, 0x6f, 0x3a, 0xaf, 0x05, 0x3a, 0x17, 0xea, 0x1e, 0x57, 0xdb,
	0xc1, 0x3a, 0x16, 0xb2, 0x2d, 0x37, 0xa8, 0x90, 0x96, 0xa4, 0xdf, 0xac, 0x4d, 0x31, 0x08, 0x30,
	0xcc, 0x26, 0xac, 0x0f, 0x8d, 0x5b, 0x8b, 0xd7, 0xc3, 0x40, 0x62, 0xc4, 0x3d, 0xbe, 0xa6, 0x0e,
	0x9c, 0x70, 0x3d, 0x1a, 0xce, 0x1c, 0x72, 0x49, 0x5a, 0x65, 0xef, 0x07, 0xd2, 0x3a, 0x1c, 0x4b,
	0x8c, 0x9c, 0x52, 0x3a, 0xd5, 0x8f, 0x2c, 0xdc, 0xe9, 0x72, 0x3d, 0x5d, 0xd2, 0x26, 0x9c, 0x46,
	0xb1, 0x50, 0x8f, 0x5b, 0xc9, 0x8d, 0x4d, 0x8e, 0x69, 0x17, 0x6a, 0x76, 0xc6, 0xd4, 0xb0, 0xda,
	0x39, 0x6f, 0x4b, 0xbf, 0x6d, 0xbb, 0x3d, 0x85, 0x42, 0x79, 0x05, 0x26, 0xeb, 0x80, 0x63, 0x33,
	0xfa, 0xdc, 0xca, 0xdd, 0x80, 0x8a, 0x3f, 0x59, 0x8c, 0x30, 0x32, 0xfb, 0x0c, 0x62, 0xfd, 0x3d,
	0x1a, 0x9d, 0xb2, 0x05, 0xe5, 0x05, 0xe2, 0x2c, 0x93, 0x54, 0x3b, 0x54, 0x47, 0xb8, 0xd3, 0x83,
	0xde, 0xcb, 0x24, 0x5c, 0xf0, 0x07, 0x35, 0xf7, 0x32, 0x02, 0x1b, 0xc2, 0x99, 0xed, 0xf2, 0xcc,
	0x27, 0x9b, 0x3f, 0x96, 0xea, 0xd7, 0x4f, 0x5b, 0x1b, 0xe5, 0x7d, 0xe5, 0x98, 0x2d, 0x7f, 0x5b,
	0x1d, 0xae, 0xb1, 0x01, 0x5c, 0xd8, 0x0c, 0x8f, 0x07, 0xf8, 0x9a, 0xd6, 0x65, 0x8e, 0x49, 0xf2,
	0x63, 0x16, 0x02, 0x94, 0x8a, 0x01, 0x58, 0xb0, 0xcb, 0xe6, 0x70, 0xa9, 0xbb, 0xe0, 0xda, 0x8c,
	0x71, 0x2c, 0xd4, 0x78, 0x29, 0x56, 0xab, 0xec, 0x2c, 0xa6, 0x78, 0xbd, 0x27, 0xff, 0x48, 0x0d,
	0xfa, 0x47, 0xa9, 0x13, 0xef, 0x51, 0xde, 0xd4, 0xdf, 0x13, 0x97, 0x7c, 0x24, 0x2e, 0xf9, 0x4c,
	0x5c, 0xf2, 0xf6, 0xe5, 0x1e, 0xf9, 0x95, 0xf4, 0x57, 0xb9, 0xfa, 0x0e, 0x00, 0x00, 0xff, 0xff,
	0x2c, 0x44, 0x53, 0x6c, 0x52, 0x03, 0x00, 0x00,
}