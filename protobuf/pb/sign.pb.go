// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: sign.proto

package pb

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

import io "io"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// 签到
type SignReq struct {
}

func (m *SignReq) Reset()                    { *m = SignReq{} }
func (m *SignReq) String() string            { return proto.CompactTextString(m) }
func (*SignReq) ProtoMessage()               {}
func (*SignReq) Descriptor() ([]byte, []int) { return fileDescriptorSign, []int{0} }

type SignAck struct {
	SignInfo *SignInfo       `protobuf:"bytes,1,opt,name=signInfo" json:"signInfo,omitempty"`
	Goods    *GoodsChangeNtf `protobuf:"bytes,2,opt,name=goods" json:"goods,omitempty"`
}

func (m *SignAck) Reset()                    { *m = SignAck{} }
func (m *SignAck) String() string            { return proto.CompactTextString(m) }
func (*SignAck) ProtoMessage()               {}
func (*SignAck) Descriptor() ([]byte, []int) { return fileDescriptorSign, []int{1} }

func (m *SignAck) GetSignInfo() *SignInfo {
	if m != nil {
		return m.SignInfo
	}
	return nil
}

func (m *SignAck) GetGoods() *GoodsChangeNtf {
	if m != nil {
		return m.Goods
	}
	return nil
}

// 补签
type SignRepairReq struct {
	RepairDay int32 `protobuf:"varint,1,opt,name=repairDay,proto3" json:"repairDay,omitempty"`
}

func (m *SignRepairReq) Reset()                    { *m = SignRepairReq{} }
func (m *SignRepairReq) String() string            { return proto.CompactTextString(m) }
func (*SignRepairReq) ProtoMessage()               {}
func (*SignRepairReq) Descriptor() ([]byte, []int) { return fileDescriptorSign, []int{2} }

func (m *SignRepairReq) GetRepairDay() int32 {
	if m != nil {
		return m.RepairDay
	}
	return 0
}

type SignRepairAck struct {
	SignInfo *SignInfo       `protobuf:"bytes,1,opt,name=signInfo" json:"signInfo,omitempty"`
	Goods    *GoodsChangeNtf `protobuf:"bytes,2,opt,name=goods" json:"goods,omitempty"`
}

func (m *SignRepairAck) Reset()                    { *m = SignRepairAck{} }
func (m *SignRepairAck) String() string            { return proto.CompactTextString(m) }
func (*SignRepairAck) ProtoMessage()               {}
func (*SignRepairAck) Descriptor() ([]byte, []int) { return fileDescriptorSign, []int{3} }

func (m *SignRepairAck) GetSignInfo() *SignInfo {
	if m != nil {
		return m.SignInfo
	}
	return nil
}

func (m *SignRepairAck) GetGoods() *GoodsChangeNtf {
	if m != nil {
		return m.Goods
	}
	return nil
}

// 累计奖励
type CumulativeSignReq struct {
	CumulativeDay int32 `protobuf:"varint,1,opt,name=cumulativeDay,proto3" json:"cumulativeDay,omitempty"`
}

func (m *CumulativeSignReq) Reset()                    { *m = CumulativeSignReq{} }
func (m *CumulativeSignReq) String() string            { return proto.CompactTextString(m) }
func (*CumulativeSignReq) ProtoMessage()               {}
func (*CumulativeSignReq) Descriptor() ([]byte, []int) { return fileDescriptorSign, []int{4} }

func (m *CumulativeSignReq) GetCumulativeDay() int32 {
	if m != nil {
		return m.CumulativeDay
	}
	return 0
}

type CumulativeSignAck struct {
	SignInfo *SignInfo       `protobuf:"bytes,1,opt,name=signInfo" json:"signInfo,omitempty"`
	Goods    *GoodsChangeNtf `protobuf:"bytes,2,opt,name=goods" json:"goods,omitempty"`
}

func (m *CumulativeSignAck) Reset()                    { *m = CumulativeSignAck{} }
func (m *CumulativeSignAck) String() string            { return proto.CompactTextString(m) }
func (*CumulativeSignAck) ProtoMessage()               {}
func (*CumulativeSignAck) Descriptor() ([]byte, []int) { return fileDescriptorSign, []int{5} }

func (m *CumulativeSignAck) GetSignInfo() *SignInfo {
	if m != nil {
		return m.SignInfo
	}
	return nil
}

func (m *CumulativeSignAck) GetGoods() *GoodsChangeNtf {
	if m != nil {
		return m.Goods
	}
	return nil
}

// 签到周期第一天推送
type SignResetNtf struct {
	SignInfo *SignInfo `protobuf:"bytes,1,opt,name=signInfo" json:"signInfo,omitempty"`
}

func (m *SignResetNtf) Reset()                    { *m = SignResetNtf{} }
func (m *SignResetNtf) String() string            { return proto.CompactTextString(m) }
func (*SignResetNtf) ProtoMessage()               {}
func (*SignResetNtf) Descriptor() ([]byte, []int) { return fileDescriptorSign, []int{6} }

func (m *SignResetNtf) GetSignInfo() *SignInfo {
	if m != nil {
		return m.SignInfo
	}
	return nil
}

func init() {
	proto.RegisterType((*SignReq)(nil), "pb.SignReq")
	proto.RegisterType((*SignAck)(nil), "pb.SignAck")
	proto.RegisterType((*SignRepairReq)(nil), "pb.SignRepairReq")
	proto.RegisterType((*SignRepairAck)(nil), "pb.SignRepairAck")
	proto.RegisterType((*CumulativeSignReq)(nil), "pb.CumulativeSignReq")
	proto.RegisterType((*CumulativeSignAck)(nil), "pb.CumulativeSignAck")
	proto.RegisterType((*SignResetNtf)(nil), "pb.SignResetNtf")
}
func (m *SignReq) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *SignReq) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	return i, nil
}

func (m *SignAck) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *SignAck) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if m.SignInfo != nil {
		dAtA[i] = 0xa
		i++
		i = encodeVarintSign(dAtA, i, uint64(m.SignInfo.Size()))
		n1, err := m.SignInfo.MarshalTo(dAtA[i:])
		if err != nil {
			return 0, err
		}
		i += n1
	}
	if m.Goods != nil {
		dAtA[i] = 0x12
		i++
		i = encodeVarintSign(dAtA, i, uint64(m.Goods.Size()))
		n2, err := m.Goods.MarshalTo(dAtA[i:])
		if err != nil {
			return 0, err
		}
		i += n2
	}
	return i, nil
}

func (m *SignRepairReq) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *SignRepairReq) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if m.RepairDay != 0 {
		dAtA[i] = 0x8
		i++
		i = encodeVarintSign(dAtA, i, uint64(m.RepairDay))
	}
	return i, nil
}

func (m *SignRepairAck) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *SignRepairAck) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if m.SignInfo != nil {
		dAtA[i] = 0xa
		i++
		i = encodeVarintSign(dAtA, i, uint64(m.SignInfo.Size()))
		n3, err := m.SignInfo.MarshalTo(dAtA[i:])
		if err != nil {
			return 0, err
		}
		i += n3
	}
	if m.Goods != nil {
		dAtA[i] = 0x12
		i++
		i = encodeVarintSign(dAtA, i, uint64(m.Goods.Size()))
		n4, err := m.Goods.MarshalTo(dAtA[i:])
		if err != nil {
			return 0, err
		}
		i += n4
	}
	return i, nil
}

func (m *CumulativeSignReq) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *CumulativeSignReq) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if m.CumulativeDay != 0 {
		dAtA[i] = 0x8
		i++
		i = encodeVarintSign(dAtA, i, uint64(m.CumulativeDay))
	}
	return i, nil
}

func (m *CumulativeSignAck) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *CumulativeSignAck) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if m.SignInfo != nil {
		dAtA[i] = 0xa
		i++
		i = encodeVarintSign(dAtA, i, uint64(m.SignInfo.Size()))
		n5, err := m.SignInfo.MarshalTo(dAtA[i:])
		if err != nil {
			return 0, err
		}
		i += n5
	}
	if m.Goods != nil {
		dAtA[i] = 0x12
		i++
		i = encodeVarintSign(dAtA, i, uint64(m.Goods.Size()))
		n6, err := m.Goods.MarshalTo(dAtA[i:])
		if err != nil {
			return 0, err
		}
		i += n6
	}
	return i, nil
}

func (m *SignResetNtf) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *SignResetNtf) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if m.SignInfo != nil {
		dAtA[i] = 0xa
		i++
		i = encodeVarintSign(dAtA, i, uint64(m.SignInfo.Size()))
		n7, err := m.SignInfo.MarshalTo(dAtA[i:])
		if err != nil {
			return 0, err
		}
		i += n7
	}
	return i, nil
}

func encodeVarintSign(dAtA []byte, offset int, v uint64) int {
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return offset + 1
}
func (m *SignReq) Size() (n int) {
	var l int
	_ = l
	return n
}

func (m *SignAck) Size() (n int) {
	var l int
	_ = l
	if m.SignInfo != nil {
		l = m.SignInfo.Size()
		n += 1 + l + sovSign(uint64(l))
	}
	if m.Goods != nil {
		l = m.Goods.Size()
		n += 1 + l + sovSign(uint64(l))
	}
	return n
}

func (m *SignRepairReq) Size() (n int) {
	var l int
	_ = l
	if m.RepairDay != 0 {
		n += 1 + sovSign(uint64(m.RepairDay))
	}
	return n
}

func (m *SignRepairAck) Size() (n int) {
	var l int
	_ = l
	if m.SignInfo != nil {
		l = m.SignInfo.Size()
		n += 1 + l + sovSign(uint64(l))
	}
	if m.Goods != nil {
		l = m.Goods.Size()
		n += 1 + l + sovSign(uint64(l))
	}
	return n
}

func (m *CumulativeSignReq) Size() (n int) {
	var l int
	_ = l
	if m.CumulativeDay != 0 {
		n += 1 + sovSign(uint64(m.CumulativeDay))
	}
	return n
}

func (m *CumulativeSignAck) Size() (n int) {
	var l int
	_ = l
	if m.SignInfo != nil {
		l = m.SignInfo.Size()
		n += 1 + l + sovSign(uint64(l))
	}
	if m.Goods != nil {
		l = m.Goods.Size()
		n += 1 + l + sovSign(uint64(l))
	}
	return n
}

func (m *SignResetNtf) Size() (n int) {
	var l int
	_ = l
	if m.SignInfo != nil {
		l = m.SignInfo.Size()
		n += 1 + l + sovSign(uint64(l))
	}
	return n
}

func sovSign(x uint64) (n int) {
	for {
		n++
		x >>= 7
		if x == 0 {
			break
		}
	}
	return n
}
func sozSign(x uint64) (n int) {
	return sovSign(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *SignReq) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowSign
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
			return fmt.Errorf("proto: SignReq: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: SignReq: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		default:
			iNdEx = preIndex
			skippy, err := skipSign(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthSign
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
func (m *SignAck) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowSign
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
			return fmt.Errorf("proto: SignAck: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: SignAck: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field SignInfo", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSign
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
				return ErrInvalidLengthSign
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.SignInfo == nil {
				m.SignInfo = &SignInfo{}
			}
			if err := m.SignInfo.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Goods", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSign
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
				return ErrInvalidLengthSign
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
			skippy, err := skipSign(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthSign
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
func (m *SignRepairReq) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowSign
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
			return fmt.Errorf("proto: SignRepairReq: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: SignRepairReq: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field RepairDay", wireType)
			}
			m.RepairDay = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSign
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.RepairDay |= (int32(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipSign(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthSign
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
func (m *SignRepairAck) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowSign
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
			return fmt.Errorf("proto: SignRepairAck: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: SignRepairAck: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field SignInfo", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSign
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
				return ErrInvalidLengthSign
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.SignInfo == nil {
				m.SignInfo = &SignInfo{}
			}
			if err := m.SignInfo.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Goods", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSign
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
				return ErrInvalidLengthSign
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
			skippy, err := skipSign(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthSign
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
func (m *CumulativeSignReq) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowSign
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
			return fmt.Errorf("proto: CumulativeSignReq: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: CumulativeSignReq: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field CumulativeDay", wireType)
			}
			m.CumulativeDay = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSign
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.CumulativeDay |= (int32(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipSign(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthSign
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
func (m *CumulativeSignAck) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowSign
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
			return fmt.Errorf("proto: CumulativeSignAck: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: CumulativeSignAck: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field SignInfo", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSign
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
				return ErrInvalidLengthSign
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.SignInfo == nil {
				m.SignInfo = &SignInfo{}
			}
			if err := m.SignInfo.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Goods", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSign
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
				return ErrInvalidLengthSign
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
			skippy, err := skipSign(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthSign
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
func (m *SignResetNtf) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowSign
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
			return fmt.Errorf("proto: SignResetNtf: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: SignResetNtf: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field SignInfo", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSign
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
				return ErrInvalidLengthSign
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.SignInfo == nil {
				m.SignInfo = &SignInfo{}
			}
			if err := m.SignInfo.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipSign(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthSign
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
func skipSign(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowSign
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
					return 0, ErrIntOverflowSign
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
					return 0, ErrIntOverflowSign
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
				return 0, ErrInvalidLengthSign
			}
			return iNdEx, nil
		case 3:
			for {
				var innerWire uint64
				var start int = iNdEx
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return 0, ErrIntOverflowSign
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
				next, err := skipSign(dAtA[start:])
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
	ErrInvalidLengthSign = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowSign   = fmt.Errorf("proto: integer overflow")
)

func init() { proto.RegisterFile("sign.proto", fileDescriptorSign) }

var fileDescriptorSign = []byte{
	// 233 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0x2a, 0xce, 0x4c, 0xcf,
	0xd3, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x62, 0x2a, 0x48, 0x92, 0xe2, 0x49, 0xce, 0xcf, 0xcd,
	0xcd, 0x87, 0x8a, 0x28, 0x71, 0x72, 0xb1, 0x07, 0x67, 0xa6, 0xe7, 0x05, 0xa5, 0x16, 0x2a, 0xc5,
	0x42, 0x98, 0x8e, 0xc9, 0xd9, 0x42, 0x1a, 0x5c, 0x1c, 0x20, 0x5d, 0x9e, 0x79, 0x69, 0xf9, 0x12,
	0x8c, 0x0a, 0x8c, 0x1a, 0xdc, 0x46, 0x3c, 0x7a, 0x05, 0x49, 0x7a, 0xc1, 0x50, 0xb1, 0x20, 0xb8,
	0xac, 0x90, 0x06, 0x17, 0x6b, 0x7a, 0x7e, 0x7e, 0x4a, 0xb1, 0x04, 0x13, 0x58, 0x99, 0x10, 0x48,
	0x99, 0x3b, 0x48, 0xc0, 0x39, 0x23, 0x31, 0x2f, 0x3d, 0xd5, 0xaf, 0x24, 0x2d, 0x08, 0xa2, 0x40,
	0x49, 0x97, 0x8b, 0x17, 0x62, 0x53, 0x41, 0x62, 0x66, 0x51, 0x50, 0x6a, 0xa1, 0x90, 0x0c, 0x17,
	0x67, 0x11, 0x98, 0xe3, 0x92, 0x58, 0x09, 0xb6, 0x85, 0x35, 0x08, 0x21, 0xa0, 0x94, 0x8c, 0xac,
	0x9c, 0x56, 0x6e, 0xb2, 0xe4, 0x12, 0x74, 0x2e, 0xcd, 0x2d, 0xcd, 0x49, 0x2c, 0xc9, 0x2c, 0x4b,
	0x85, 0x86, 0x83, 0x90, 0x0a, 0x17, 0x6f, 0x32, 0x5c, 0x10, 0xe1, 0x36, 0x54, 0x41, 0xa5, 0x74,
	0x74, 0xad, 0xb4, 0x72, 0xa3, 0x05, 0x17, 0x0f, 0xc4, 0x65, 0xc5, 0xa9, 0x25, 0x7e, 0x25, 0x69,
	0xc4, 0xdb, 0xe1, 0x24, 0x70, 0xe2, 0x91, 0x1c, 0xe3, 0x85, 0x47, 0x72, 0x8c, 0x0f, 0x1e, 0xc9,
	0x31, 0xce, 0x78, 0x2c, 0xc7, 0x90, 0xc4, 0x06, 0x8e, 0x74, 0x63, 0x40, 0x00, 0x00, 0x00, 0xff,
	0xff, 0xf2, 0xa5, 0x2c, 0x79, 0x14, 0x02, 0x00, 0x00,
}
