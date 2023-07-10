// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: loginserver.proto

package pbserver

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

import io "io"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

type LoginKeyVerifyReq struct {
	OpenId   string `protobuf:"bytes,1,opt,name=openId,proto3" json:"openId,omitempty"`
	UserId   uint32 `protobuf:"varint,2,opt,name=userId,proto3" json:"userId,omitempty"`
	LoginKey string `protobuf:"bytes,3,opt,name=loginKey,proto3" json:"loginKey,omitempty"`
	ClientIp string `protobuf:"bytes,4,opt,name=clientIp,proto3" json:"clientIp,omitempty"`
}

func (m *LoginKeyVerifyReq) Reset()                    { *m = LoginKeyVerifyReq{} }
func (m *LoginKeyVerifyReq) String() string            { return proto.CompactTextString(m) }
func (*LoginKeyVerifyReq) ProtoMessage()               {}
func (*LoginKeyVerifyReq) Descriptor() ([]byte, []int) { return fileDescriptorLoginserver, []int{0} }

func (m *LoginKeyVerifyReq) GetOpenId() string {
	if m != nil {
		return m.OpenId
	}
	return ""
}

func (m *LoginKeyVerifyReq) GetUserId() uint32 {
	if m != nil {
		return m.UserId
	}
	return 0
}

func (m *LoginKeyVerifyReq) GetLoginKey() string {
	if m != nil {
		return m.LoginKey
	}
	return ""
}

func (m *LoginKeyVerifyReq) GetClientIp() string {
	if m != nil {
		return m.ClientIp
	}
	return ""
}

type LoginKeyVerifyAck struct {
	Result  int32 `protobuf:"varint,1,opt,name=result,proto3" json:"result,omitempty"`
	Channel int32 `protobuf:"varint,2,opt,name=channel,proto3" json:"channel,omitempty"`
}

func (m *LoginKeyVerifyAck) Reset()                    { *m = LoginKeyVerifyAck{} }
func (m *LoginKeyVerifyAck) String() string            { return proto.CompactTextString(m) }
func (*LoginKeyVerifyAck) ProtoMessage()               {}
func (*LoginKeyVerifyAck) Descriptor() ([]byte, []int) { return fileDescriptorLoginserver, []int{1} }

func (m *LoginKeyVerifyAck) GetResult() int32 {
	if m != nil {
		return m.Result
	}
	return 0
}

func (m *LoginKeyVerifyAck) GetChannel() int32 {
	if m != nil {
		return m.Channel
	}
	return 0
}

func init() {
	proto.RegisterType((*LoginKeyVerifyReq)(nil), "pbserver.LoginKeyVerifyReq")
	proto.RegisterType((*LoginKeyVerifyAck)(nil), "pbserver.LoginKeyVerifyAck")
}
func (m *LoginKeyVerifyReq) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *LoginKeyVerifyReq) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if len(m.OpenId) > 0 {
		dAtA[i] = 0xa
		i++
		i = encodeVarintLoginserver(dAtA, i, uint64(len(m.OpenId)))
		i += copy(dAtA[i:], m.OpenId)
	}
	if m.UserId != 0 {
		dAtA[i] = 0x10
		i++
		i = encodeVarintLoginserver(dAtA, i, uint64(m.UserId))
	}
	if len(m.LoginKey) > 0 {
		dAtA[i] = 0x1a
		i++
		i = encodeVarintLoginserver(dAtA, i, uint64(len(m.LoginKey)))
		i += copy(dAtA[i:], m.LoginKey)
	}
	if len(m.ClientIp) > 0 {
		dAtA[i] = 0x22
		i++
		i = encodeVarintLoginserver(dAtA, i, uint64(len(m.ClientIp)))
		i += copy(dAtA[i:], m.ClientIp)
	}
	return i, nil
}

func (m *LoginKeyVerifyAck) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *LoginKeyVerifyAck) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if m.Result != 0 {
		dAtA[i] = 0x8
		i++
		i = encodeVarintLoginserver(dAtA, i, uint64(m.Result))
	}
	if m.Channel != 0 {
		dAtA[i] = 0x10
		i++
		i = encodeVarintLoginserver(dAtA, i, uint64(m.Channel))
	}
	return i, nil
}

func encodeVarintLoginserver(dAtA []byte, offset int, v uint64) int {
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return offset + 1
}
func (m *LoginKeyVerifyReq) Size() (n int) {
	var l int
	_ = l
	l = len(m.OpenId)
	if l > 0 {
		n += 1 + l + sovLoginserver(uint64(l))
	}
	if m.UserId != 0 {
		n += 1 + sovLoginserver(uint64(m.UserId))
	}
	l = len(m.LoginKey)
	if l > 0 {
		n += 1 + l + sovLoginserver(uint64(l))
	}
	l = len(m.ClientIp)
	if l > 0 {
		n += 1 + l + sovLoginserver(uint64(l))
	}
	return n
}

func (m *LoginKeyVerifyAck) Size() (n int) {
	var l int
	_ = l
	if m.Result != 0 {
		n += 1 + sovLoginserver(uint64(m.Result))
	}
	if m.Channel != 0 {
		n += 1 + sovLoginserver(uint64(m.Channel))
	}
	return n
}

func sovLoginserver(x uint64) (n int) {
	for {
		n++
		x >>= 7
		if x == 0 {
			break
		}
	}
	return n
}
func sozLoginserver(x uint64) (n int) {
	return sovLoginserver(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *LoginKeyVerifyReq) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowLoginserver
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
			return fmt.Errorf("proto: LoginKeyVerifyReq: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: LoginKeyVerifyReq: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field OpenId", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowLoginserver
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
				return ErrInvalidLengthLoginserver
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.OpenId = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field UserId", wireType)
			}
			m.UserId = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowLoginserver
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.UserId |= (uint32(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field LoginKey", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowLoginserver
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
				return ErrInvalidLengthLoginserver
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.LoginKey = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ClientIp", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowLoginserver
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
				return ErrInvalidLengthLoginserver
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ClientIp = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipLoginserver(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthLoginserver
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
func (m *LoginKeyVerifyAck) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowLoginserver
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
			return fmt.Errorf("proto: LoginKeyVerifyAck: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: LoginKeyVerifyAck: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Result", wireType)
			}
			m.Result = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowLoginserver
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
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Channel", wireType)
			}
			m.Channel = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowLoginserver
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Channel |= (int32(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipLoginserver(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthLoginserver
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
func skipLoginserver(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowLoginserver
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
					return 0, ErrIntOverflowLoginserver
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
					return 0, ErrIntOverflowLoginserver
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
				return 0, ErrInvalidLengthLoginserver
			}
			return iNdEx, nil
		case 3:
			for {
				var innerWire uint64
				var start int = iNdEx
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return 0, ErrIntOverflowLoginserver
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
				next, err := skipLoginserver(dAtA[start:])
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
	ErrInvalidLengthLoginserver = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowLoginserver   = fmt.Errorf("proto: integer overflow")
)

func init() { proto.RegisterFile("loginserver.proto", fileDescriptorLoginserver) }

var fileDescriptorLoginserver = []byte{
	// 194 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x12, 0xcc, 0xc9, 0x4f, 0xcf,
	0xcc, 0x2b, 0x4e, 0x2d, 0x2a, 0x4b, 0x2d, 0xd2, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0xe2, 0x28,
	0x48, 0x82, 0xf0, 0x95, 0xaa, 0xb9, 0x04, 0x7d, 0x40, 0xd2, 0xde, 0xa9, 0x95, 0x61, 0xa9, 0x45,
	0x99, 0x69, 0x95, 0x41, 0xa9, 0x85, 0x42, 0x62, 0x5c, 0x6c, 0xf9, 0x05, 0xa9, 0x79, 0x9e, 0x29,
	0x12, 0x8c, 0x0a, 0x8c, 0x1a, 0x9c, 0x41, 0x50, 0x1e, 0x48, 0xbc, 0xb4, 0x38, 0xb5, 0xc8, 0x33,
	0x45, 0x82, 0x49, 0x81, 0x51, 0x83, 0x37, 0x08, 0xca, 0x13, 0x92, 0xe2, 0xe2, 0xc8, 0x81, 0x1a,
	0x22, 0xc1, 0x0c, 0xd6, 0x01, 0xe7, 0x83, 0xe4, 0x92, 0x73, 0x32, 0x53, 0xf3, 0x4a, 0x3c, 0x0b,
	0x24, 0x58, 0x20, 0x72, 0x30, 0xbe, 0x92, 0x2b, 0xba, 0xe5, 0x8e, 0xc9, 0xd9, 0x20, 0x4b, 0x8a,
	0x52, 0x8b, 0x4b, 0x73, 0x4a, 0xc0, 0x96, 0xb3, 0x06, 0x41, 0x79, 0x42, 0x12, 0x5c, 0xec, 0xc9,
	0x19, 0x89, 0x79, 0x79, 0xa9, 0x39, 0x60, 0xdb, 0x59, 0x83, 0x60, 0x5c, 0x27, 0x81, 0x13, 0x8f,
	0xe4, 0x18, 0x2f, 0x3c, 0x92, 0x63, 0x7c, 0xf0, 0x48, 0x8e, 0x71, 0xc6, 0x63, 0x39, 0x86, 0x24,
	0x36, 0xb0, 0x37, 0x8d, 0x01, 0x01, 0x00, 0x00, 0xff, 0xff, 0xe2, 0xef, 0xa3, 0xe7, 0xfb, 0x00,
	0x00, 0x00,
}
