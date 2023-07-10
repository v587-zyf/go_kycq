// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: panacea.proto

package pb

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

import io "io"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

type PanaceaUseReq struct {
	Id int32 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (m *PanaceaUseReq) Reset()                    { *m = PanaceaUseReq{} }
func (m *PanaceaUseReq) String() string            { return proto.CompactTextString(m) }
func (*PanaceaUseReq) ProtoMessage()               {}
func (*PanaceaUseReq) Descriptor() ([]byte, []int) { return fileDescriptorPanacea, []int{0} }

func (m *PanaceaUseReq) GetId() int32 {
	if m != nil {
		return m.Id
	}
	return 0
}

type PanaceaUseAck struct {
	Id      int32        `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Panacea *PanaceaInfo `protobuf:"bytes,2,opt,name=panacea" json:"panacea,omitempty"`
}

func (m *PanaceaUseAck) Reset()                    { *m = PanaceaUseAck{} }
func (m *PanaceaUseAck) String() string            { return proto.CompactTextString(m) }
func (*PanaceaUseAck) ProtoMessage()               {}
func (*PanaceaUseAck) Descriptor() ([]byte, []int) { return fileDescriptorPanacea, []int{1} }

func (m *PanaceaUseAck) GetId() int32 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *PanaceaUseAck) GetPanacea() *PanaceaInfo {
	if m != nil {
		return m.Panacea
	}
	return nil
}

func init() {
	proto.RegisterType((*PanaceaUseReq)(nil), "pb.PanaceaUseReq")
	proto.RegisterType((*PanaceaUseAck)(nil), "pb.PanaceaUseAck")
}
func (m *PanaceaUseReq) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *PanaceaUseReq) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if m.Id != 0 {
		dAtA[i] = 0x8
		i++
		i = encodeVarintPanacea(dAtA, i, uint64(m.Id))
	}
	return i, nil
}

func (m *PanaceaUseAck) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *PanaceaUseAck) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if m.Id != 0 {
		dAtA[i] = 0x8
		i++
		i = encodeVarintPanacea(dAtA, i, uint64(m.Id))
	}
	if m.Panacea != nil {
		dAtA[i] = 0x12
		i++
		i = encodeVarintPanacea(dAtA, i, uint64(m.Panacea.Size()))
		n1, err := m.Panacea.MarshalTo(dAtA[i:])
		if err != nil {
			return 0, err
		}
		i += n1
	}
	return i, nil
}

func encodeVarintPanacea(dAtA []byte, offset int, v uint64) int {
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return offset + 1
}
func (m *PanaceaUseReq) Size() (n int) {
	var l int
	_ = l
	if m.Id != 0 {
		n += 1 + sovPanacea(uint64(m.Id))
	}
	return n
}

func (m *PanaceaUseAck) Size() (n int) {
	var l int
	_ = l
	if m.Id != 0 {
		n += 1 + sovPanacea(uint64(m.Id))
	}
	if m.Panacea != nil {
		l = m.Panacea.Size()
		n += 1 + l + sovPanacea(uint64(l))
	}
	return n
}

func sovPanacea(x uint64) (n int) {
	for {
		n++
		x >>= 7
		if x == 0 {
			break
		}
	}
	return n
}
func sozPanacea(x uint64) (n int) {
	return sovPanacea(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *PanaceaUseReq) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowPanacea
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
			return fmt.Errorf("proto: PanaceaUseReq: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: PanaceaUseReq: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Id", wireType)
			}
			m.Id = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPanacea
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
			skippy, err := skipPanacea(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthPanacea
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
func (m *PanaceaUseAck) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowPanacea
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
			return fmt.Errorf("proto: PanaceaUseAck: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: PanaceaUseAck: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Id", wireType)
			}
			m.Id = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPanacea
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
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Panacea", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPanacea
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
				return ErrInvalidLengthPanacea
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Panacea == nil {
				m.Panacea = &PanaceaInfo{}
			}
			if err := m.Panacea.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipPanacea(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthPanacea
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
func skipPanacea(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowPanacea
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
					return 0, ErrIntOverflowPanacea
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
					return 0, ErrIntOverflowPanacea
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
				return 0, ErrInvalidLengthPanacea
			}
			return iNdEx, nil
		case 3:
			for {
				var innerWire uint64
				var start int = iNdEx
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return 0, ErrIntOverflowPanacea
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
				next, err := skipPanacea(dAtA[start:])
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
	ErrInvalidLengthPanacea = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowPanacea   = fmt.Errorf("proto: integer overflow")
)

func init() { proto.RegisterFile("panacea.proto", fileDescriptorPanacea) }

var fileDescriptorPanacea = []byte{
	// 142 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0x2d, 0x48, 0xcc, 0x4b,
	0x4c, 0x4e, 0x4d, 0xd4, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x62, 0x2a, 0x48, 0x92, 0xe2, 0x49,
	0xce, 0xcf, 0xcd, 0xcd, 0xcf, 0x83, 0x88, 0x28, 0xc9, 0x73, 0xf1, 0x06, 0x40, 0x94, 0x84, 0x16,
	0xa7, 0x06, 0xa5, 0x16, 0x0a, 0xf1, 0x71, 0x31, 0x65, 0xa6, 0x48, 0x30, 0x2a, 0x30, 0x6a, 0xb0,
	0x06, 0x31, 0x65, 0xa6, 0x28, 0x79, 0x21, 0x2b, 0x70, 0x4c, 0xce, 0x46, 0x57, 0x20, 0xa4, 0xc9,
	0xc5, 0x0e, 0xb5, 0x44, 0x82, 0x49, 0x81, 0x51, 0x83, 0xdb, 0x88, 0x5f, 0xaf, 0x20, 0x49, 0x0f,
	0xaa, 0xc7, 0x33, 0x2f, 0x2d, 0x3f, 0x08, 0x26, 0xef, 0x24, 0x70, 0xe2, 0x91, 0x1c, 0xe3, 0x85,
	0x47, 0x72, 0x8c, 0x0f, 0x1e, 0xc9, 0x31, 0xce, 0x78, 0x2c, 0xc7, 0x90, 0xc4, 0x06, 0x76, 0x85,
	0x31, 0x20, 0x00, 0x00, 0xff, 0xff, 0x7b, 0xa1, 0xfa, 0xf0, 0xa8, 0x00, 0x00, 0x00,
}