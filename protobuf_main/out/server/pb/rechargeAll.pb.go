// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: rechargeAll.proto

package pb

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

import io "io"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// 累计充值礼包领取
type RechargeAllGetReq struct {
	Id int32 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (m *RechargeAllGetReq) Reset()                    { *m = RechargeAllGetReq{} }
func (m *RechargeAllGetReq) String() string            { return proto.CompactTextString(m) }
func (*RechargeAllGetReq) ProtoMessage()               {}
func (*RechargeAllGetReq) Descriptor() ([]byte, []int) { return fileDescriptorRechargeAll, []int{0} }

func (m *RechargeAllGetReq) GetId() int32 {
	if m != nil {
		return m.Id
	}
	return 0
}

type RechargeAllGetAck struct {
	RechargeGetGetIds []int32 `protobuf:"varint,1,rep,packed,name=rechargeGetGetIds" json:"rechargeGetGetIds,omitempty"`
	RechargeAll       int32   `protobuf:"varint,2,opt,name=rechargeAll,proto3" json:"rechargeAll,omitempty"`
}

func (m *RechargeAllGetAck) Reset()                    { *m = RechargeAllGetAck{} }
func (m *RechargeAllGetAck) String() string            { return proto.CompactTextString(m) }
func (*RechargeAllGetAck) ProtoMessage()               {}
func (*RechargeAllGetAck) Descriptor() ([]byte, []int) { return fileDescriptorRechargeAll, []int{1} }

func (m *RechargeAllGetAck) GetRechargeGetGetIds() []int32 {
	if m != nil {
		return m.RechargeGetGetIds
	}
	return nil
}

func (m *RechargeAllGetAck) GetRechargeAll() int32 {
	if m != nil {
		return m.RechargeAll
	}
	return 0
}

func init() {
	proto.RegisterType((*RechargeAllGetReq)(nil), "pb.RechargeAllGetReq")
	proto.RegisterType((*RechargeAllGetAck)(nil), "pb.RechargeAllGetAck")
}
func (m *RechargeAllGetReq) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *RechargeAllGetReq) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if m.Id != 0 {
		dAtA[i] = 0x8
		i++
		i = encodeVarintRechargeAll(dAtA, i, uint64(m.Id))
	}
	return i, nil
}

func (m *RechargeAllGetAck) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *RechargeAllGetAck) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if len(m.RechargeGetGetIds) > 0 {
		dAtA2 := make([]byte, len(m.RechargeGetGetIds)*10)
		var j1 int
		for _, num1 := range m.RechargeGetGetIds {
			num := uint64(num1)
			for num >= 1<<7 {
				dAtA2[j1] = uint8(uint64(num)&0x7f | 0x80)
				num >>= 7
				j1++
			}
			dAtA2[j1] = uint8(num)
			j1++
		}
		dAtA[i] = 0xa
		i++
		i = encodeVarintRechargeAll(dAtA, i, uint64(j1))
		i += copy(dAtA[i:], dAtA2[:j1])
	}
	if m.RechargeAll != 0 {
		dAtA[i] = 0x10
		i++
		i = encodeVarintRechargeAll(dAtA, i, uint64(m.RechargeAll))
	}
	return i, nil
}

func encodeVarintRechargeAll(dAtA []byte, offset int, v uint64) int {
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return offset + 1
}
func (m *RechargeAllGetReq) Size() (n int) {
	var l int
	_ = l
	if m.Id != 0 {
		n += 1 + sovRechargeAll(uint64(m.Id))
	}
	return n
}

func (m *RechargeAllGetAck) Size() (n int) {
	var l int
	_ = l
	if len(m.RechargeGetGetIds) > 0 {
		l = 0
		for _, e := range m.RechargeGetGetIds {
			l += sovRechargeAll(uint64(e))
		}
		n += 1 + sovRechargeAll(uint64(l)) + l
	}
	if m.RechargeAll != 0 {
		n += 1 + sovRechargeAll(uint64(m.RechargeAll))
	}
	return n
}

func sovRechargeAll(x uint64) (n int) {
	for {
		n++
		x >>= 7
		if x == 0 {
			break
		}
	}
	return n
}
func sozRechargeAll(x uint64) (n int) {
	return sovRechargeAll(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *RechargeAllGetReq) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowRechargeAll
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
			return fmt.Errorf("proto: RechargeAllGetReq: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: RechargeAllGetReq: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Id", wireType)
			}
			m.Id = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRechargeAll
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
		default:
			iNdEx = preIndex
			skippy, err := skipRechargeAll(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthRechargeAll
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
func (m *RechargeAllGetAck) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowRechargeAll
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
			return fmt.Errorf("proto: RechargeAllGetAck: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: RechargeAllGetAck: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType == 0 {
				var v int32
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return ErrIntOverflowRechargeAll
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
				m.RechargeGetGetIds = append(m.RechargeGetGetIds, v)
			} else if wireType == 2 {
				var packedLen int
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return ErrIntOverflowRechargeAll
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
					return ErrInvalidLengthRechargeAll
				}
				postIndex := iNdEx + packedLen
				if postIndex > l {
					return io.ErrUnexpectedEOF
				}
				for iNdEx < postIndex {
					var v int32
					for shift := uint(0); ; shift += 7 {
						if shift >= 64 {
							return ErrIntOverflowRechargeAll
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
					m.RechargeGetGetIds = append(m.RechargeGetGetIds, v)
				}
			} else {
				return fmt.Errorf("proto: wrong wireType = %d for field RechargeGetGetIds", wireType)
			}
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field RechargeAll", wireType)
			}
			m.RechargeAll = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRechargeAll
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.RechargeAll |= (int32(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipRechargeAll(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthRechargeAll
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
func skipRechargeAll(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowRechargeAll
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
					return 0, ErrIntOverflowRechargeAll
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
					return 0, ErrIntOverflowRechargeAll
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
				return 0, ErrInvalidLengthRechargeAll
			}
			return iNdEx, nil
		case 3:
			for {
				var innerWire uint64
				var start int = iNdEx
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return 0, ErrIntOverflowRechargeAll
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
				next, err := skipRechargeAll(dAtA[start:])
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
	ErrInvalidLengthRechargeAll = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowRechargeAll   = fmt.Errorf("proto: integer overflow")
)

func init() { proto.RegisterFile("rechargeAll.proto", fileDescriptorRechargeAll) }

var fileDescriptorRechargeAll = []byte{
	// 142 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x12, 0x2c, 0x4a, 0x4d, 0xce,
	0x48, 0x2c, 0x4a, 0x4f, 0x75, 0xcc, 0xc9, 0xd1, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x62, 0x2a,
	0x48, 0x52, 0x52, 0xe6, 0x12, 0x0c, 0x42, 0x48, 0xb8, 0xa7, 0x96, 0x04, 0xa5, 0x16, 0x0a, 0xf1,
	0x71, 0x31, 0x65, 0xa6, 0x48, 0x30, 0x2a, 0x30, 0x6a, 0xb0, 0x06, 0x31, 0x65, 0xa6, 0x28, 0x25,
	0xa3, 0x2b, 0x72, 0x4c, 0xce, 0x16, 0xd2, 0x41, 0x18, 0xe9, 0x9e, 0x5a, 0xe2, 0x9e, 0x5a, 0xe2,
	0x99, 0x52, 0x2c, 0xc1, 0xa8, 0xc0, 0xac, 0xc1, 0x1a, 0x84, 0x29, 0x21, 0xa4, 0xc0, 0xc5, 0x8d,
	0xe4, 0x00, 0x09, 0x26, 0xb0, 0xd9, 0xc8, 0x42, 0x4e, 0x02, 0x27, 0x1e, 0xc9, 0x31, 0x5e, 0x78,
	0x24, 0xc7, 0xf8, 0xe0, 0x91, 0x1c, 0xe3, 0x8c, 0xc7, 0x72, 0x0c, 0x49, 0x6c, 0x60, 0x67, 0x1a,
	0x03, 0x02, 0x00, 0x00, 0xff, 0xff, 0xa0, 0x83, 0x91, 0x8b, 0xbb, 0x00, 0x00, 0x00,
}
