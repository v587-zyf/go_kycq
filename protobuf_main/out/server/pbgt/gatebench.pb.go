// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: gatebench.proto

package pbgt

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

import io "io"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

type ClientMsgType int32

const (
	ClientMsgType__       ClientMsgType = 0
	ClientMsgType_ChatReq ClientMsgType = 11
	ClientMsgType_ChatAck ClientMsgType = 12
	ClientMsgType_ChatNtf ClientMsgType = 13
	ClientMsgType_MoveRpt ClientMsgType = 556
	ClientMsgType_MoveNtf ClientMsgType = 557
)

var ClientMsgType_name = map[int32]string{
	0:   "_",
	11:  "ChatReq",
	12:  "ChatAck",
	13:  "ChatNtf",
	556: "MoveRpt",
	557: "MoveNtf",
}
var ClientMsgType_value = map[string]int32{
	"_":       0,
	"ChatReq": 11,
	"ChatAck": 12,
	"ChatNtf": 13,
	"MoveRpt": 556,
	"MoveNtf": 557,
}

func (x ClientMsgType) String() string {
	return proto.EnumName(ClientMsgType_name, int32(x))
}
func (ClientMsgType) EnumDescriptor() ([]byte, []int) { return fileDescriptorGatebench, []int{0} }

type ClientChatReq struct {
	Timestamp int64  `protobuf:"varint,1,opt,name=timestamp,proto3" json:"timestamp,omitempty"`
	SenderId  int32  `protobuf:"varint,2,opt,name=senderId,proto3" json:"senderId,omitempty"`
	Msg       string `protobuf:"bytes,3,opt,name=msg,proto3" json:"msg,omitempty"`
	Broadcast bool   `protobuf:"varint,4,opt,name=broadcast,proto3" json:"broadcast,omitempty"`
}

func (m *ClientChatReq) Reset()                    { *m = ClientChatReq{} }
func (m *ClientChatReq) String() string            { return proto.CompactTextString(m) }
func (*ClientChatReq) ProtoMessage()               {}
func (*ClientChatReq) Descriptor() ([]byte, []int) { return fileDescriptorGatebench, []int{0} }

func (m *ClientChatReq) GetTimestamp() int64 {
	if m != nil {
		return m.Timestamp
	}
	return 0
}

func (m *ClientChatReq) GetSenderId() int32 {
	if m != nil {
		return m.SenderId
	}
	return 0
}

func (m *ClientChatReq) GetMsg() string {
	if m != nil {
		return m.Msg
	}
	return ""
}

func (m *ClientChatReq) GetBroadcast() bool {
	if m != nil {
		return m.Broadcast
	}
	return false
}

type ClientChatAck struct {
	SendTimestamp int64 `protobuf:"varint,1,opt,name=sendTimestamp,proto3" json:"sendTimestamp,omitempty"`
	SenderId      int32 `protobuf:"varint,2,opt,name=senderId,proto3" json:"senderId,omitempty"`
}

func (m *ClientChatAck) Reset()                    { *m = ClientChatAck{} }
func (m *ClientChatAck) String() string            { return proto.CompactTextString(m) }
func (*ClientChatAck) ProtoMessage()               {}
func (*ClientChatAck) Descriptor() ([]byte, []int) { return fileDescriptorGatebench, []int{1} }

func (m *ClientChatAck) GetSendTimestamp() int64 {
	if m != nil {
		return m.SendTimestamp
	}
	return 0
}

func (m *ClientChatAck) GetSenderId() int32 {
	if m != nil {
		return m.SenderId
	}
	return 0
}

type ClientChatNtf struct {
	SendTimestamp int64  `protobuf:"varint,1,opt,name=sendTimestamp,proto3" json:"sendTimestamp,omitempty"`
	SenderId      int32  `protobuf:"varint,2,opt,name=senderId,proto3" json:"senderId,omitempty"`
	Msg           string `protobuf:"bytes,3,opt,name=msg,proto3" json:"msg,omitempty"`
}

func (m *ClientChatNtf) Reset()                    { *m = ClientChatNtf{} }
func (m *ClientChatNtf) String() string            { return proto.CompactTextString(m) }
func (*ClientChatNtf) ProtoMessage()               {}
func (*ClientChatNtf) Descriptor() ([]byte, []int) { return fileDescriptorGatebench, []int{2} }

func (m *ClientChatNtf) GetSendTimestamp() int64 {
	if m != nil {
		return m.SendTimestamp
	}
	return 0
}

func (m *ClientChatNtf) GetSenderId() int32 {
	if m != nil {
		return m.SenderId
	}
	return 0
}

func (m *ClientChatNtf) GetMsg() string {
	if m != nil {
		return m.Msg
	}
	return ""
}

type ClientMoveRpt struct {
	Timestamp int64 `protobuf:"varint,1,opt,name=timestamp,proto3" json:"timestamp,omitempty"`
	SenderId  int32 `protobuf:"varint,2,opt,name=senderId,proto3" json:"senderId,omitempty"`
	X         int32 `protobuf:"varint,3,opt,name=x,proto3" json:"x,omitempty"`
	Y         int32 `protobuf:"varint,4,opt,name=y,proto3" json:"y,omitempty"`
}

func (m *ClientMoveRpt) Reset()                    { *m = ClientMoveRpt{} }
func (m *ClientMoveRpt) String() string            { return proto.CompactTextString(m) }
func (*ClientMoveRpt) ProtoMessage()               {}
func (*ClientMoveRpt) Descriptor() ([]byte, []int) { return fileDescriptorGatebench, []int{3} }

func (m *ClientMoveRpt) GetTimestamp() int64 {
	if m != nil {
		return m.Timestamp
	}
	return 0
}

func (m *ClientMoveRpt) GetSenderId() int32 {
	if m != nil {
		return m.SenderId
	}
	return 0
}

func (m *ClientMoveRpt) GetX() int32 {
	if m != nil {
		return m.X
	}
	return 0
}

func (m *ClientMoveRpt) GetY() int32 {
	if m != nil {
		return m.Y
	}
	return 0
}

type ClientMoveNtf struct {
	SendTimestamp int64 `protobuf:"varint,1,opt,name=sendTimestamp,proto3" json:"sendTimestamp,omitempty"`
	SenderId      int32 `protobuf:"varint,2,opt,name=senderId,proto3" json:"senderId,omitempty"`
	X             int32 `protobuf:"varint,3,opt,name=x,proto3" json:"x,omitempty"`
	Y             int32 `protobuf:"varint,4,opt,name=y,proto3" json:"y,omitempty"`
}

func (m *ClientMoveNtf) Reset()                    { *m = ClientMoveNtf{} }
func (m *ClientMoveNtf) String() string            { return proto.CompactTextString(m) }
func (*ClientMoveNtf) ProtoMessage()               {}
func (*ClientMoveNtf) Descriptor() ([]byte, []int) { return fileDescriptorGatebench, []int{4} }

func (m *ClientMoveNtf) GetSendTimestamp() int64 {
	if m != nil {
		return m.SendTimestamp
	}
	return 0
}

func (m *ClientMoveNtf) GetSenderId() int32 {
	if m != nil {
		return m.SenderId
	}
	return 0
}

func (m *ClientMoveNtf) GetX() int32 {
	if m != nil {
		return m.X
	}
	return 0
}

func (m *ClientMoveNtf) GetY() int32 {
	if m != nil {
		return m.Y
	}
	return 0
}

func init() {
	proto.RegisterType((*ClientChatReq)(nil), "pbgt.ClientChatReq")
	proto.RegisterType((*ClientChatAck)(nil), "pbgt.ClientChatAck")
	proto.RegisterType((*ClientChatNtf)(nil), "pbgt.ClientChatNtf")
	proto.RegisterType((*ClientMoveRpt)(nil), "pbgt.ClientMoveRpt")
	proto.RegisterType((*ClientMoveNtf)(nil), "pbgt.ClientMoveNtf")
	proto.RegisterEnum("pbgt.ClientMsgType", ClientMsgType_name, ClientMsgType_value)
}
func (m *ClientChatReq) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *ClientChatReq) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if m.Timestamp != 0 {
		dAtA[i] = 0x8
		i++
		i = encodeVarintGatebench(dAtA, i, uint64(m.Timestamp))
	}
	if m.SenderId != 0 {
		dAtA[i] = 0x10
		i++
		i = encodeVarintGatebench(dAtA, i, uint64(m.SenderId))
	}
	if len(m.Msg) > 0 {
		dAtA[i] = 0x1a
		i++
		i = encodeVarintGatebench(dAtA, i, uint64(len(m.Msg)))
		i += copy(dAtA[i:], m.Msg)
	}
	if m.Broadcast {
		dAtA[i] = 0x20
		i++
		if m.Broadcast {
			dAtA[i] = 1
		} else {
			dAtA[i] = 0
		}
		i++
	}
	return i, nil
}

func (m *ClientChatAck) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *ClientChatAck) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if m.SendTimestamp != 0 {
		dAtA[i] = 0x8
		i++
		i = encodeVarintGatebench(dAtA, i, uint64(m.SendTimestamp))
	}
	if m.SenderId != 0 {
		dAtA[i] = 0x10
		i++
		i = encodeVarintGatebench(dAtA, i, uint64(m.SenderId))
	}
	return i, nil
}

func (m *ClientChatNtf) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *ClientChatNtf) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if m.SendTimestamp != 0 {
		dAtA[i] = 0x8
		i++
		i = encodeVarintGatebench(dAtA, i, uint64(m.SendTimestamp))
	}
	if m.SenderId != 0 {
		dAtA[i] = 0x10
		i++
		i = encodeVarintGatebench(dAtA, i, uint64(m.SenderId))
	}
	if len(m.Msg) > 0 {
		dAtA[i] = 0x1a
		i++
		i = encodeVarintGatebench(dAtA, i, uint64(len(m.Msg)))
		i += copy(dAtA[i:], m.Msg)
	}
	return i, nil
}

func (m *ClientMoveRpt) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *ClientMoveRpt) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if m.Timestamp != 0 {
		dAtA[i] = 0x8
		i++
		i = encodeVarintGatebench(dAtA, i, uint64(m.Timestamp))
	}
	if m.SenderId != 0 {
		dAtA[i] = 0x10
		i++
		i = encodeVarintGatebench(dAtA, i, uint64(m.SenderId))
	}
	if m.X != 0 {
		dAtA[i] = 0x18
		i++
		i = encodeVarintGatebench(dAtA, i, uint64(m.X))
	}
	if m.Y != 0 {
		dAtA[i] = 0x20
		i++
		i = encodeVarintGatebench(dAtA, i, uint64(m.Y))
	}
	return i, nil
}

func (m *ClientMoveNtf) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *ClientMoveNtf) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if m.SendTimestamp != 0 {
		dAtA[i] = 0x8
		i++
		i = encodeVarintGatebench(dAtA, i, uint64(m.SendTimestamp))
	}
	if m.SenderId != 0 {
		dAtA[i] = 0x10
		i++
		i = encodeVarintGatebench(dAtA, i, uint64(m.SenderId))
	}
	if m.X != 0 {
		dAtA[i] = 0x18
		i++
		i = encodeVarintGatebench(dAtA, i, uint64(m.X))
	}
	if m.Y != 0 {
		dAtA[i] = 0x20
		i++
		i = encodeVarintGatebench(dAtA, i, uint64(m.Y))
	}
	return i, nil
}

func encodeVarintGatebench(dAtA []byte, offset int, v uint64) int {
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return offset + 1
}
func (m *ClientChatReq) Size() (n int) {
	var l int
	_ = l
	if m.Timestamp != 0 {
		n += 1 + sovGatebench(uint64(m.Timestamp))
	}
	if m.SenderId != 0 {
		n += 1 + sovGatebench(uint64(m.SenderId))
	}
	l = len(m.Msg)
	if l > 0 {
		n += 1 + l + sovGatebench(uint64(l))
	}
	if m.Broadcast {
		n += 2
	}
	return n
}

func (m *ClientChatAck) Size() (n int) {
	var l int
	_ = l
	if m.SendTimestamp != 0 {
		n += 1 + sovGatebench(uint64(m.SendTimestamp))
	}
	if m.SenderId != 0 {
		n += 1 + sovGatebench(uint64(m.SenderId))
	}
	return n
}

func (m *ClientChatNtf) Size() (n int) {
	var l int
	_ = l
	if m.SendTimestamp != 0 {
		n += 1 + sovGatebench(uint64(m.SendTimestamp))
	}
	if m.SenderId != 0 {
		n += 1 + sovGatebench(uint64(m.SenderId))
	}
	l = len(m.Msg)
	if l > 0 {
		n += 1 + l + sovGatebench(uint64(l))
	}
	return n
}

func (m *ClientMoveRpt) Size() (n int) {
	var l int
	_ = l
	if m.Timestamp != 0 {
		n += 1 + sovGatebench(uint64(m.Timestamp))
	}
	if m.SenderId != 0 {
		n += 1 + sovGatebench(uint64(m.SenderId))
	}
	if m.X != 0 {
		n += 1 + sovGatebench(uint64(m.X))
	}
	if m.Y != 0 {
		n += 1 + sovGatebench(uint64(m.Y))
	}
	return n
}

func (m *ClientMoveNtf) Size() (n int) {
	var l int
	_ = l
	if m.SendTimestamp != 0 {
		n += 1 + sovGatebench(uint64(m.SendTimestamp))
	}
	if m.SenderId != 0 {
		n += 1 + sovGatebench(uint64(m.SenderId))
	}
	if m.X != 0 {
		n += 1 + sovGatebench(uint64(m.X))
	}
	if m.Y != 0 {
		n += 1 + sovGatebench(uint64(m.Y))
	}
	return n
}

func sovGatebench(x uint64) (n int) {
	for {
		n++
		x >>= 7
		if x == 0 {
			break
		}
	}
	return n
}
func sozGatebench(x uint64) (n int) {
	return sovGatebench(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *ClientChatReq) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowGatebench
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
			return fmt.Errorf("proto: ClientChatReq: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: ClientChatReq: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Timestamp", wireType)
			}
			m.Timestamp = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGatebench
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Timestamp |= (int64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field SenderId", wireType)
			}
			m.SenderId = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGatebench
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.SenderId |= (int32(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Msg", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGatebench
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
				return ErrInvalidLengthGatebench
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Msg = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 4:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Broadcast", wireType)
			}
			var v int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGatebench
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
			m.Broadcast = bool(v != 0)
		default:
			iNdEx = preIndex
			skippy, err := skipGatebench(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthGatebench
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
func (m *ClientChatAck) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowGatebench
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
			return fmt.Errorf("proto: ClientChatAck: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: ClientChatAck: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field SendTimestamp", wireType)
			}
			m.SendTimestamp = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGatebench
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.SendTimestamp |= (int64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field SenderId", wireType)
			}
			m.SenderId = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGatebench
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.SenderId |= (int32(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipGatebench(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthGatebench
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
func (m *ClientChatNtf) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowGatebench
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
			return fmt.Errorf("proto: ClientChatNtf: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: ClientChatNtf: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field SendTimestamp", wireType)
			}
			m.SendTimestamp = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGatebench
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.SendTimestamp |= (int64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field SenderId", wireType)
			}
			m.SenderId = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGatebench
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.SenderId |= (int32(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Msg", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGatebench
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
				return ErrInvalidLengthGatebench
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Msg = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipGatebench(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthGatebench
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
func (m *ClientMoveRpt) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowGatebench
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
			return fmt.Errorf("proto: ClientMoveRpt: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: ClientMoveRpt: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Timestamp", wireType)
			}
			m.Timestamp = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGatebench
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Timestamp |= (int64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field SenderId", wireType)
			}
			m.SenderId = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGatebench
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.SenderId |= (int32(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field X", wireType)
			}
			m.X = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGatebench
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.X |= (int32(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 4:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Y", wireType)
			}
			m.Y = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGatebench
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Y |= (int32(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipGatebench(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthGatebench
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
func (m *ClientMoveNtf) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowGatebench
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
			return fmt.Errorf("proto: ClientMoveNtf: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: ClientMoveNtf: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field SendTimestamp", wireType)
			}
			m.SendTimestamp = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGatebench
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.SendTimestamp |= (int64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field SenderId", wireType)
			}
			m.SenderId = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGatebench
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.SenderId |= (int32(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field X", wireType)
			}
			m.X = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGatebench
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.X |= (int32(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 4:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Y", wireType)
			}
			m.Y = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGatebench
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Y |= (int32(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipGatebench(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthGatebench
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
func skipGatebench(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowGatebench
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
					return 0, ErrIntOverflowGatebench
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
					return 0, ErrIntOverflowGatebench
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
				return 0, ErrInvalidLengthGatebench
			}
			return iNdEx, nil
		case 3:
			for {
				var innerWire uint64
				var start int = iNdEx
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return 0, ErrIntOverflowGatebench
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
				next, err := skipGatebench(dAtA[start:])
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
	ErrInvalidLengthGatebench = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowGatebench   = fmt.Errorf("proto: integer overflow")
)

func init() { proto.RegisterFile("gatebench.proto", fileDescriptorGatebench) }

var fileDescriptorGatebench = []byte{
	// 297 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0x4f, 0x4f, 0x2c, 0x49,
	0x4d, 0x4a, 0xcd, 0x4b, 0xce, 0xd0, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x62, 0x29, 0x48, 0x4a,
	0x2f, 0x51, 0xaa, 0xe4, 0xe2, 0x75, 0xce, 0xc9, 0x4c, 0xcd, 0x2b, 0x71, 0xce, 0x48, 0x2c, 0x09,
	0x4a, 0x2d, 0x14, 0x92, 0xe1, 0xe2, 0x2c, 0xc9, 0xcc, 0x4d, 0x2d, 0x2e, 0x49, 0xcc, 0x2d, 0x90,
	0x60, 0x54, 0x60, 0xd4, 0x60, 0x0e, 0x42, 0x08, 0x08, 0x49, 0x71, 0x71, 0x14, 0xa7, 0xe6, 0xa5,
	0xa4, 0x16, 0x79, 0xa6, 0x48, 0x30, 0x29, 0x30, 0x6a, 0xb0, 0x06, 0xc1, 0xf9, 0x42, 0x02, 0x5c,
	0xcc, 0xb9, 0xc5, 0xe9, 0x12, 0xcc, 0x0a, 0x8c, 0x1a, 0x9c, 0x41, 0x20, 0x26, 0xc8, 0xac, 0xa4,
	0xa2, 0xfc, 0xc4, 0x94, 0xe4, 0xc4, 0xe2, 0x12, 0x09, 0x16, 0x05, 0x46, 0x0d, 0x8e, 0x20, 0x84,
	0x80, 0x52, 0x20, 0xb2, 0xd5, 0x8e, 0xc9, 0xd9, 0x42, 0x2a, 0x5c, 0xbc, 0x20, 0xc3, 0x42, 0xd0,
	0xac, 0x47, 0x15, 0xc4, 0xe7, 0x04, 0xa5, 0x64, 0x64, 0x23, 0xfd, 0x4a, 0xd2, 0x28, 0x37, 0x12,
	0xd3, 0x57, 0x4a, 0xa9, 0x30, 0x4b, 0x7c, 0xf3, 0xcb, 0x52, 0x83, 0x0a, 0x4a, 0x28, 0x08, 0x32,
	0x1e, 0x2e, 0xc6, 0x0a, 0xb0, 0xd1, 0xac, 0x41, 0x8c, 0x15, 0x20, 0x5e, 0x25, 0x38, 0x98, 0x58,
	0x83, 0x18, 0x2b, 0x95, 0x72, 0x91, 0xad, 0xa1, 0x8e, 0x5f, 0xf0, 0x58, 0xa7, 0x15, 0x09, 0xb7,
	0xae, 0x38, 0x3d, 0xa4, 0xb2, 0x20, 0x55, 0x88, 0x95, 0x8b, 0x31, 0x5e, 0x80, 0x41, 0x88, 0x9b,
	0x8b, 0x1d, 0x9a, 0x34, 0x04, 0xb8, 0x61, 0x1c, 0xc7, 0xe4, 0x6c, 0x01, 0x1e, 0x18, 0xc7, 0xaf,
	0x24, 0x4d, 0x80, 0x57, 0x88, 0x87, 0x8b, 0x1d, 0x1a, 0x1c, 0x02, 0x6b, 0x58, 0x60, 0x3c, 0x90,
	0xd4, 0x5a, 0x16, 0x27, 0x81, 0x13, 0x8f, 0xe4, 0x18, 0x2f, 0x3c, 0x92, 0x63, 0x7c, 0xf0, 0x48,
	0x8e, 0x71, 0xc6, 0x63, 0x39, 0x86, 0x24, 0x36, 0x70, 0x12, 0x34, 0x06, 0x04, 0x00, 0x00, 0xff,
	0xff, 0x4e, 0x1e, 0xba, 0x57, 0x95, 0x02, 0x00, 0x00,
}
