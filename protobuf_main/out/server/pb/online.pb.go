// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: online.proto

package pb

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

import io "io"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// 获取在线奖励信息
type GetOnlineAwardInfoReq struct {
}

func (m *GetOnlineAwardInfoReq) Reset()                    { *m = GetOnlineAwardInfoReq{} }
func (m *GetOnlineAwardInfoReq) String() string            { return proto.CompactTextString(m) }
func (*GetOnlineAwardInfoReq) ProtoMessage()               {}
func (*GetOnlineAwardInfoReq) Descriptor() ([]byte, []int) { return fileDescriptorOnline, []int{0} }

// 获取在线奖励信息
type GetOnlineAwardInfoAck struct {
	OnlineTime int32   `protobuf:"varint,1,opt,name=onlineTime,proto3" json:"onlineTime,omitempty"`
	GetAwardId []int32 `protobuf:"varint,2,rep,packed,name=getAwardId" json:"getAwardId,omitempty"`
}

func (m *GetOnlineAwardInfoAck) Reset()                    { *m = GetOnlineAwardInfoAck{} }
func (m *GetOnlineAwardInfoAck) String() string            { return proto.CompactTextString(m) }
func (*GetOnlineAwardInfoAck) ProtoMessage()               {}
func (*GetOnlineAwardInfoAck) Descriptor() ([]byte, []int) { return fileDescriptorOnline, []int{1} }

func (m *GetOnlineAwardInfoAck) GetOnlineTime() int32 {
	if m != nil {
		return m.OnlineTime
	}
	return 0
}

func (m *GetOnlineAwardInfoAck) GetGetAwardId() []int32 {
	if m != nil {
		return m.GetAwardId
	}
	return nil
}

// 领取在线奖励
type GetOnlineAwardReq struct {
	AwardId int32 `protobuf:"varint,1,opt,name=awardId,proto3" json:"awardId,omitempty"`
}

func (m *GetOnlineAwardReq) Reset()                    { *m = GetOnlineAwardReq{} }
func (m *GetOnlineAwardReq) String() string            { return proto.CompactTextString(m) }
func (*GetOnlineAwardReq) ProtoMessage()               {}
func (*GetOnlineAwardReq) Descriptor() ([]byte, []int) { return fileDescriptorOnline, []int{2} }

func (m *GetOnlineAwardReq) GetAwardId() int32 {
	if m != nil {
		return m.AwardId
	}
	return 0
}

// 领取在线奖励
type GetOnlineAwardAck struct {
	GetAwardId []int32 `protobuf:"varint,1,rep,packed,name=getAwardId" json:"getAwardId,omitempty"`
}

func (m *GetOnlineAwardAck) Reset()                    { *m = GetOnlineAwardAck{} }
func (m *GetOnlineAwardAck) String() string            { return proto.CompactTextString(m) }
func (*GetOnlineAwardAck) ProtoMessage()               {}
func (*GetOnlineAwardAck) Descriptor() ([]byte, []int) { return fileDescriptorOnline, []int{3} }

func (m *GetOnlineAwardAck) GetGetAwardId() []int32 {
	if m != nil {
		return m.GetAwardId
	}
	return nil
}

func init() {
	proto.RegisterType((*GetOnlineAwardInfoReq)(nil), "pb.GetOnlineAwardInfoReq")
	proto.RegisterType((*GetOnlineAwardInfoAck)(nil), "pb.GetOnlineAwardInfoAck")
	proto.RegisterType((*GetOnlineAwardReq)(nil), "pb.GetOnlineAwardReq")
	proto.RegisterType((*GetOnlineAwardAck)(nil), "pb.GetOnlineAwardAck")
}
func (m *GetOnlineAwardInfoReq) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *GetOnlineAwardInfoReq) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	return i, nil
}

func (m *GetOnlineAwardInfoAck) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *GetOnlineAwardInfoAck) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if m.OnlineTime != 0 {
		dAtA[i] = 0x8
		i++
		i = encodeVarintOnline(dAtA, i, uint64(m.OnlineTime))
	}
	if len(m.GetAwardId) > 0 {
		dAtA2 := make([]byte, len(m.GetAwardId)*10)
		var j1 int
		for _, num1 := range m.GetAwardId {
			num := uint64(num1)
			for num >= 1<<7 {
				dAtA2[j1] = uint8(uint64(num)&0x7f | 0x80)
				num >>= 7
				j1++
			}
			dAtA2[j1] = uint8(num)
			j1++
		}
		dAtA[i] = 0x12
		i++
		i = encodeVarintOnline(dAtA, i, uint64(j1))
		i += copy(dAtA[i:], dAtA2[:j1])
	}
	return i, nil
}

func (m *GetOnlineAwardReq) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *GetOnlineAwardReq) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if m.AwardId != 0 {
		dAtA[i] = 0x8
		i++
		i = encodeVarintOnline(dAtA, i, uint64(m.AwardId))
	}
	return i, nil
}

func (m *GetOnlineAwardAck) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *GetOnlineAwardAck) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if len(m.GetAwardId) > 0 {
		dAtA4 := make([]byte, len(m.GetAwardId)*10)
		var j3 int
		for _, num1 := range m.GetAwardId {
			num := uint64(num1)
			for num >= 1<<7 {
				dAtA4[j3] = uint8(uint64(num)&0x7f | 0x80)
				num >>= 7
				j3++
			}
			dAtA4[j3] = uint8(num)
			j3++
		}
		dAtA[i] = 0xa
		i++
		i = encodeVarintOnline(dAtA, i, uint64(j3))
		i += copy(dAtA[i:], dAtA4[:j3])
	}
	return i, nil
}

func encodeVarintOnline(dAtA []byte, offset int, v uint64) int {
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return offset + 1
}
func (m *GetOnlineAwardInfoReq) Size() (n int) {
	var l int
	_ = l
	return n
}

func (m *GetOnlineAwardInfoAck) Size() (n int) {
	var l int
	_ = l
	if m.OnlineTime != 0 {
		n += 1 + sovOnline(uint64(m.OnlineTime))
	}
	if len(m.GetAwardId) > 0 {
		l = 0
		for _, e := range m.GetAwardId {
			l += sovOnline(uint64(e))
		}
		n += 1 + sovOnline(uint64(l)) + l
	}
	return n
}

func (m *GetOnlineAwardReq) Size() (n int) {
	var l int
	_ = l
	if m.AwardId != 0 {
		n += 1 + sovOnline(uint64(m.AwardId))
	}
	return n
}

func (m *GetOnlineAwardAck) Size() (n int) {
	var l int
	_ = l
	if len(m.GetAwardId) > 0 {
		l = 0
		for _, e := range m.GetAwardId {
			l += sovOnline(uint64(e))
		}
		n += 1 + sovOnline(uint64(l)) + l
	}
	return n
}

func sovOnline(x uint64) (n int) {
	for {
		n++
		x >>= 7
		if x == 0 {
			break
		}
	}
	return n
}
func sozOnline(x uint64) (n int) {
	return sovOnline(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *GetOnlineAwardInfoReq) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowOnline
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
			return fmt.Errorf("proto: GetOnlineAwardInfoReq: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: GetOnlineAwardInfoReq: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		default:
			iNdEx = preIndex
			skippy, err := skipOnline(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthOnline
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
func (m *GetOnlineAwardInfoAck) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowOnline
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
			return fmt.Errorf("proto: GetOnlineAwardInfoAck: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: GetOnlineAwardInfoAck: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field OnlineTime", wireType)
			}
			m.OnlineTime = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowOnline
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.OnlineTime |= (int32(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType == 0 {
				var v int32
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return ErrIntOverflowOnline
					}
					if iNdEx >= l {
						return io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					v |= (int32(b) & 0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				m.GetAwardId = append(m.GetAwardId, v)
			} else if wireType == 2 {
				var packedLen int
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return ErrIntOverflowOnline
					}
					if iNdEx >= l {
						return io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					packedLen |= (int(b) & 0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				if packedLen < 0 {
					return ErrInvalidLengthOnline
				}
				postIndex := iNdEx + packedLen
				if postIndex > l {
					return io.ErrUnexpectedEOF
				}
				for iNdEx < postIndex {
					var v int32
					for shift := uint(0); ; shift += 7 {
						if shift >= 64 {
							return ErrIntOverflowOnline
						}
						if iNdEx >= l {
							return io.ErrUnexpectedEOF
						}
						b := dAtA[iNdEx]
						iNdEx++
						v |= (int32(b) & 0x7F) << shift
						if b < 0x80 {
							break
						}
					}
					m.GetAwardId = append(m.GetAwardId, v)
				}
			} else {
				return fmt.Errorf("proto: wrong wireType = %d for field GetAwardId", wireType)
			}
		default:
			iNdEx = preIndex
			skippy, err := skipOnline(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthOnline
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
func (m *GetOnlineAwardReq) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowOnline
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
			return fmt.Errorf("proto: GetOnlineAwardReq: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: GetOnlineAwardReq: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field AwardId", wireType)
			}
			m.AwardId = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowOnline
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.AwardId |= (int32(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipOnline(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthOnline
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
func (m *GetOnlineAwardAck) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowOnline
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
			return fmt.Errorf("proto: GetOnlineAwardAck: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: GetOnlineAwardAck: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType == 0 {
				var v int32
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return ErrIntOverflowOnline
					}
					if iNdEx >= l {
						return io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					v |= (int32(b) & 0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				m.GetAwardId = append(m.GetAwardId, v)
			} else if wireType == 2 {
				var packedLen int
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return ErrIntOverflowOnline
					}
					if iNdEx >= l {
						return io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					packedLen |= (int(b) & 0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				if packedLen < 0 {
					return ErrInvalidLengthOnline
				}
				postIndex := iNdEx + packedLen
				if postIndex > l {
					return io.ErrUnexpectedEOF
				}
				for iNdEx < postIndex {
					var v int32
					for shift := uint(0); ; shift += 7 {
						if shift >= 64 {
							return ErrIntOverflowOnline
						}
						if iNdEx >= l {
							return io.ErrUnexpectedEOF
						}
						b := dAtA[iNdEx]
						iNdEx++
						v |= (int32(b) & 0x7F) << shift
						if b < 0x80 {
							break
						}
					}
					m.GetAwardId = append(m.GetAwardId, v)
				}
			} else {
				return fmt.Errorf("proto: wrong wireType = %d for field GetAwardId", wireType)
			}
		default:
			iNdEx = preIndex
			skippy, err := skipOnline(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthOnline
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
func skipOnline(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowOnline
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
					return 0, ErrIntOverflowOnline
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
					return 0, ErrIntOverflowOnline
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
				return 0, ErrInvalidLengthOnline
			}
			return iNdEx, nil
		case 3:
			for {
				var innerWire uint64
				var start int = iNdEx
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return 0, ErrIntOverflowOnline
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
				next, err := skipOnline(dAtA[start:])
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
	ErrInvalidLengthOnline = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowOnline   = fmt.Errorf("proto: integer overflow")
)

func init() { proto.RegisterFile("online.proto", fileDescriptorOnline) }

var fileDescriptorOnline = []byte{
	// 162 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0xc9, 0xcf, 0xcb, 0xc9,
	0xcc, 0x4b, 0xd5, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x62, 0x2a, 0x48, 0x52, 0x12, 0xe7, 0x12,
	0x75, 0x4f, 0x2d, 0xf1, 0x07, 0x0b, 0x3b, 0x96, 0x27, 0x16, 0xa5, 0x78, 0xe6, 0xa5, 0xe5, 0x07,
	0xa5, 0x16, 0x2a, 0x85, 0x63, 0x93, 0x70, 0x4c, 0xce, 0x16, 0x92, 0xe3, 0xe2, 0x82, 0x98, 0x12,
	0x92, 0x99, 0x9b, 0x2a, 0xc1, 0xa8, 0xc0, 0xa8, 0xc1, 0x1a, 0x84, 0x24, 0x02, 0x92, 0x4f, 0x4f,
	0x2d, 0x81, 0x68, 0x49, 0x91, 0x60, 0x52, 0x60, 0x06, 0xc9, 0x23, 0x44, 0x94, 0x74, 0xb9, 0x04,
	0x51, 0x0d, 0x0e, 0x4a, 0x2d, 0x14, 0x92, 0xe0, 0x62, 0x4f, 0x84, 0xea, 0x80, 0x98, 0x08, 0xe3,
	0x2a, 0x19, 0xa3, 0x2b, 0x87, 0xba, 0x01, 0xc9, 0x0e, 0x46, 0x74, 0x3b, 0x9c, 0x04, 0x4e, 0x3c,
	0x92, 0x63, 0xbc, 0xf0, 0x48, 0x8e, 0xf1, 0xc1, 0x23, 0x39, 0xc6, 0x19, 0x8f, 0xe5, 0x18, 0x92,
	0xd8, 0xc0, 0x5e, 0x36, 0x06, 0x04, 0x00, 0x00, 0xff, 0xff, 0xcb, 0xfa, 0xbc, 0xc2, 0x02, 0x01,
	0x00, 0x00,
}
