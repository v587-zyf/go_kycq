// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: sevenInvestment.proto

package pb

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

import io "io"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

type SevenInvestmentLoadReq struct {
}

func (m *SevenInvestmentLoadReq) Reset()         { *m = SevenInvestmentLoadReq{} }
func (m *SevenInvestmentLoadReq) String() string { return proto.CompactTextString(m) }
func (*SevenInvestmentLoadReq) ProtoMessage()    {}
func (*SevenInvestmentLoadReq) Descriptor() ([]byte, []int) {
	return fileDescriptorSevenInvestment, []int{0}
}

type SevenInvestmentLoadAck struct {
	HaveGetIds  []int32 `protobuf:"varint,1,rep,packed,name=haveGetIds" json:"haveGetIds,omitempty"`
	ActivateDay int32   `protobuf:"varint,2,opt,name=activateDay,proto3" json:"activateDay,omitempty"`
}

func (m *SevenInvestmentLoadAck) Reset()         { *m = SevenInvestmentLoadAck{} }
func (m *SevenInvestmentLoadAck) String() string { return proto.CompactTextString(m) }
func (*SevenInvestmentLoadAck) ProtoMessage()    {}
func (*SevenInvestmentLoadAck) Descriptor() ([]byte, []int) {
	return fileDescriptorSevenInvestment, []int{1}
}

func (m *SevenInvestmentLoadAck) GetHaveGetIds() []int32 {
	if m != nil {
		return m.HaveGetIds
	}
	return nil
}

func (m *SevenInvestmentLoadAck) GetActivateDay() int32 {
	if m != nil {
		return m.ActivateDay
	}
	return 0
}

type GetSevenInvestmentAwardReq struct {
	Id int32 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (m *GetSevenInvestmentAwardReq) Reset()         { *m = GetSevenInvestmentAwardReq{} }
func (m *GetSevenInvestmentAwardReq) String() string { return proto.CompactTextString(m) }
func (*GetSevenInvestmentAwardReq) ProtoMessage()    {}
func (*GetSevenInvestmentAwardReq) Descriptor() ([]byte, []int) {
	return fileDescriptorSevenInvestment, []int{2}
}

func (m *GetSevenInvestmentAwardReq) GetId() int32 {
	if m != nil {
		return m.Id
	}
	return 0
}

type GetSevenInvestmentAwardAck struct {
	HaveGetIds []int32 `protobuf:"varint,1,rep,packed,name=haveGetIds" json:"haveGetIds,omitempty"`
}

func (m *GetSevenInvestmentAwardAck) Reset()         { *m = GetSevenInvestmentAwardAck{} }
func (m *GetSevenInvestmentAwardAck) String() string { return proto.CompactTextString(m) }
func (*GetSevenInvestmentAwardAck) ProtoMessage()    {}
func (*GetSevenInvestmentAwardAck) Descriptor() ([]byte, []int) {
	return fileDescriptorSevenInvestment, []int{3}
}

func (m *GetSevenInvestmentAwardAck) GetHaveGetIds() []int32 {
	if m != nil {
		return m.HaveGetIds
	}
	return nil
}

func init() {
	proto.RegisterType((*SevenInvestmentLoadReq)(nil), "pb.SevenInvestmentLoadReq")
	proto.RegisterType((*SevenInvestmentLoadAck)(nil), "pb.SevenInvestmentLoadAck")
	proto.RegisterType((*GetSevenInvestmentAwardReq)(nil), "pb.GetSevenInvestmentAwardReq")
	proto.RegisterType((*GetSevenInvestmentAwardAck)(nil), "pb.GetSevenInvestmentAwardAck")
}
func (m *SevenInvestmentLoadReq) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *SevenInvestmentLoadReq) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	return i, nil
}

func (m *SevenInvestmentLoadAck) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *SevenInvestmentLoadAck) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if len(m.HaveGetIds) > 0 {
		dAtA2 := make([]byte, len(m.HaveGetIds)*10)
		var j1 int
		for _, num1 := range m.HaveGetIds {
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
		i = encodeVarintSevenInvestment(dAtA, i, uint64(j1))
		i += copy(dAtA[i:], dAtA2[:j1])
	}
	if m.ActivateDay != 0 {
		dAtA[i] = 0x10
		i++
		i = encodeVarintSevenInvestment(dAtA, i, uint64(m.ActivateDay))
	}
	return i, nil
}

func (m *GetSevenInvestmentAwardReq) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *GetSevenInvestmentAwardReq) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if m.Id != 0 {
		dAtA[i] = 0x8
		i++
		i = encodeVarintSevenInvestment(dAtA, i, uint64(m.Id))
	}
	return i, nil
}

func (m *GetSevenInvestmentAwardAck) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *GetSevenInvestmentAwardAck) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if len(m.HaveGetIds) > 0 {
		dAtA4 := make([]byte, len(m.HaveGetIds)*10)
		var j3 int
		for _, num1 := range m.HaveGetIds {
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
		i = encodeVarintSevenInvestment(dAtA, i, uint64(j3))
		i += copy(dAtA[i:], dAtA4[:j3])
	}
	return i, nil
}

func encodeVarintSevenInvestment(dAtA []byte, offset int, v uint64) int {
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return offset + 1
}
func (m *SevenInvestmentLoadReq) Size() (n int) {
	var l int
	_ = l
	return n
}

func (m *SevenInvestmentLoadAck) Size() (n int) {
	var l int
	_ = l
	if len(m.HaveGetIds) > 0 {
		l = 0
		for _, e := range m.HaveGetIds {
			l += sovSevenInvestment(uint64(e))
		}
		n += 1 + sovSevenInvestment(uint64(l)) + l
	}
	if m.ActivateDay != 0 {
		n += 1 + sovSevenInvestment(uint64(m.ActivateDay))
	}
	return n
}

func (m *GetSevenInvestmentAwardReq) Size() (n int) {
	var l int
	_ = l
	if m.Id != 0 {
		n += 1 + sovSevenInvestment(uint64(m.Id))
	}
	return n
}

func (m *GetSevenInvestmentAwardAck) Size() (n int) {
	var l int
	_ = l
	if len(m.HaveGetIds) > 0 {
		l = 0
		for _, e := range m.HaveGetIds {
			l += sovSevenInvestment(uint64(e))
		}
		n += 1 + sovSevenInvestment(uint64(l)) + l
	}
	return n
}

func sovSevenInvestment(x uint64) (n int) {
	for {
		n++
		x >>= 7
		if x == 0 {
			break
		}
	}
	return n
}
func sozSevenInvestment(x uint64) (n int) {
	return sovSevenInvestment(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *SevenInvestmentLoadReq) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowSevenInvestment
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
			return fmt.Errorf("proto: SevenInvestmentLoadReq: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: SevenInvestmentLoadReq: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		default:
			iNdEx = preIndex
			skippy, err := skipSevenInvestment(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthSevenInvestment
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
func (m *SevenInvestmentLoadAck) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowSevenInvestment
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
			return fmt.Errorf("proto: SevenInvestmentLoadAck: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: SevenInvestmentLoadAck: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType == 0 {
				var v int32
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return ErrIntOverflowSevenInvestment
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
				m.HaveGetIds = append(m.HaveGetIds, v)
			} else if wireType == 2 {
				var packedLen int
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return ErrIntOverflowSevenInvestment
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
					return ErrInvalidLengthSevenInvestment
				}
				postIndex := iNdEx + packedLen
				if postIndex > l {
					return io.ErrUnexpectedEOF
				}
				for iNdEx < postIndex {
					var v int32
					for shift := uint(0); ; shift += 7 {
						if shift >= 64 {
							return ErrIntOverflowSevenInvestment
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
					m.HaveGetIds = append(m.HaveGetIds, v)
				}
			} else {
				return fmt.Errorf("proto: wrong wireType = %d for field HaveGetIds", wireType)
			}
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field ActivateDay", wireType)
			}
			m.ActivateDay = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSevenInvestment
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.ActivateDay |= (int32(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipSevenInvestment(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthSevenInvestment
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
func (m *GetSevenInvestmentAwardReq) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowSevenInvestment
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
			return fmt.Errorf("proto: GetSevenInvestmentAwardReq: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: GetSevenInvestmentAwardReq: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Id", wireType)
			}
			m.Id = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSevenInvestment
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
			skippy, err := skipSevenInvestment(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthSevenInvestment
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
func (m *GetSevenInvestmentAwardAck) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowSevenInvestment
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
			return fmt.Errorf("proto: GetSevenInvestmentAwardAck: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: GetSevenInvestmentAwardAck: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType == 0 {
				var v int32
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return ErrIntOverflowSevenInvestment
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
				m.HaveGetIds = append(m.HaveGetIds, v)
			} else if wireType == 2 {
				var packedLen int
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return ErrIntOverflowSevenInvestment
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
					return ErrInvalidLengthSevenInvestment
				}
				postIndex := iNdEx + packedLen
				if postIndex > l {
					return io.ErrUnexpectedEOF
				}
				for iNdEx < postIndex {
					var v int32
					for shift := uint(0); ; shift += 7 {
						if shift >= 64 {
							return ErrIntOverflowSevenInvestment
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
					m.HaveGetIds = append(m.HaveGetIds, v)
				}
			} else {
				return fmt.Errorf("proto: wrong wireType = %d for field HaveGetIds", wireType)
			}
		default:
			iNdEx = preIndex
			skippy, err := skipSevenInvestment(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthSevenInvestment
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
func skipSevenInvestment(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowSevenInvestment
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
					return 0, ErrIntOverflowSevenInvestment
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
					return 0, ErrIntOverflowSevenInvestment
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
				return 0, ErrInvalidLengthSevenInvestment
			}
			return iNdEx, nil
		case 3:
			for {
				var innerWire uint64
				var start int = iNdEx
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return 0, ErrIntOverflowSevenInvestment
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
				next, err := skipSevenInvestment(dAtA[start:])
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
	ErrInvalidLengthSevenInvestment = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowSevenInvestment   = fmt.Errorf("proto: integer overflow")
)

func init() { proto.RegisterFile("sevenInvestment.proto", fileDescriptorSevenInvestment) }

var fileDescriptorSevenInvestment = []byte{
	// 179 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x12, 0x2d, 0x4e, 0x2d, 0x4b,
	0xcd, 0xf3, 0xcc, 0x2b, 0x4b, 0x2d, 0x2e, 0xc9, 0x4d, 0xcd, 0x2b, 0xd1, 0x2b, 0x28, 0xca, 0x2f,
	0xc9, 0x17, 0x62, 0x2a, 0x48, 0x52, 0x92, 0xe0, 0x12, 0x0b, 0x46, 0x95, 0xf4, 0xc9, 0x4f, 0x4c,
	0x09, 0x4a, 0x2d, 0x54, 0x8a, 0xc2, 0x2a, 0xe3, 0x98, 0x9c, 0x2d, 0x24, 0xc7, 0xc5, 0x95, 0x91,
	0x58, 0x96, 0xea, 0x9e, 0x5a, 0xe2, 0x99, 0x52, 0x2c, 0xc1, 0xa8, 0xc0, 0xac, 0xc1, 0x1a, 0x84,
	0x24, 0x22, 0xa4, 0xc0, 0xc5, 0x9d, 0x98, 0x5c, 0x92, 0x59, 0x96, 0x58, 0x92, 0xea, 0x92, 0x58,
	0x29, 0xc1, 0xa4, 0xc0, 0xa8, 0xc1, 0x1a, 0x84, 0x2c, 0xa4, 0xa4, 0xc3, 0x25, 0xe5, 0x9e, 0x5a,
	0x82, 0x66, 0xbc, 0x63, 0x79, 0x62, 0x11, 0xc8, 0x66, 0x21, 0x3e, 0x2e, 0xa6, 0xcc, 0x14, 0x09,
	0x46, 0xb0, 0x36, 0xa6, 0xcc, 0x14, 0x25, 0x1b, 0x9c, 0xaa, 0x89, 0x70, 0x8d, 0x93, 0xc0, 0x89,
	0x47, 0x72, 0x8c, 0x17, 0x1e, 0xc9, 0x31, 0x3e, 0x78, 0x24, 0xc7, 0x38, 0xe3, 0xb1, 0x1c, 0x43,
	0x12, 0x1b, 0xd8, 0xfb, 0xc6, 0x80, 0x00, 0x00, 0x00, 0xff, 0xff, 0x3e, 0x05, 0x8e, 0x2b, 0x17,
	0x01, 0x00, 0x00,
}
