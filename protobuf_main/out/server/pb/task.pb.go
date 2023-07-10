// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: task.proto

package pb

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

import io "io"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

type TaskDoneReq struct {
	TaskId int32 `protobuf:"varint,1,opt,name=taskId,proto3" json:"taskId,omitempty"`
}

func (m *TaskDoneReq) Reset()                    { *m = TaskDoneReq{} }
func (m *TaskDoneReq) String() string            { return proto.CompactTextString(m) }
func (*TaskDoneReq) ProtoMessage()               {}
func (*TaskDoneReq) Descriptor() ([]byte, []int) { return fileDescriptorTask, []int{0} }

func (m *TaskDoneReq) GetTaskId() int32 {
	if m != nil {
		return m.TaskId
	}
	return 0
}

type TaskDoneAck struct {
	Goods *GoodsChangeNtf `protobuf:"bytes,1,opt,name=goods" json:"goods,omitempty"`
}

func (m *TaskDoneAck) Reset()                    { *m = TaskDoneAck{} }
func (m *TaskDoneAck) String() string            { return proto.CompactTextString(m) }
func (*TaskDoneAck) ProtoMessage()               {}
func (*TaskDoneAck) Descriptor() ([]byte, []int) { return fileDescriptorTask, []int{1} }

func (m *TaskDoneAck) GetGoods() *GoodsChangeNtf {
	if m != nil {
		return m.Goods
	}
	return nil
}

type TaskNpcStateReq struct {
}

func (m *TaskNpcStateReq) Reset()                    { *m = TaskNpcStateReq{} }
func (m *TaskNpcStateReq) String() string            { return proto.CompactTextString(m) }
func (*TaskNpcStateReq) ProtoMessage()               {}
func (*TaskNpcStateReq) Descriptor() ([]byte, []int) { return fileDescriptorTask, []int{2} }

type TaskNpcStateAck struct {
	TaskId  int32 `protobuf:"varint,1,opt,name=taskId,proto3" json:"taskId,omitempty"`
	Process int32 `protobuf:"varint,2,opt,name=process,proto3" json:"process,omitempty"`
}

func (m *TaskNpcStateAck) Reset()                    { *m = TaskNpcStateAck{} }
func (m *TaskNpcStateAck) String() string            { return proto.CompactTextString(m) }
func (*TaskNpcStateAck) ProtoMessage()               {}
func (*TaskNpcStateAck) Descriptor() ([]byte, []int) { return fileDescriptorTask, []int{3} }

func (m *TaskNpcStateAck) GetTaskId() int32 {
	if m != nil {
		return m.TaskId
	}
	return 0
}

func (m *TaskNpcStateAck) GetProcess() int32 {
	if m != nil {
		return m.Process
	}
	return 0
}

func init() {
	proto.RegisterType((*TaskDoneReq)(nil), "pb.TaskDoneReq")
	proto.RegisterType((*TaskDoneAck)(nil), "pb.TaskDoneAck")
	proto.RegisterType((*TaskNpcStateReq)(nil), "pb.TaskNpcStateReq")
	proto.RegisterType((*TaskNpcStateAck)(nil), "pb.TaskNpcStateAck")
}
func (m *TaskDoneReq) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *TaskDoneReq) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if m.TaskId != 0 {
		dAtA[i] = 0x8
		i++
		i = encodeVarintTask(dAtA, i, uint64(m.TaskId))
	}
	return i, nil
}

func (m *TaskDoneAck) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *TaskDoneAck) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if m.Goods != nil {
		dAtA[i] = 0xa
		i++
		i = encodeVarintTask(dAtA, i, uint64(m.Goods.Size()))
		n1, err := m.Goods.MarshalTo(dAtA[i:])
		if err != nil {
			return 0, err
		}
		i += n1
	}
	return i, nil
}

func (m *TaskNpcStateReq) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *TaskNpcStateReq) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	return i, nil
}

func (m *TaskNpcStateAck) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *TaskNpcStateAck) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if m.TaskId != 0 {
		dAtA[i] = 0x8
		i++
		i = encodeVarintTask(dAtA, i, uint64(m.TaskId))
	}
	if m.Process != 0 {
		dAtA[i] = 0x10
		i++
		i = encodeVarintTask(dAtA, i, uint64(m.Process))
	}
	return i, nil
}

func encodeVarintTask(dAtA []byte, offset int, v uint64) int {
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return offset + 1
}
func (m *TaskDoneReq) Size() (n int) {
	var l int
	_ = l
	if m.TaskId != 0 {
		n += 1 + sovTask(uint64(m.TaskId))
	}
	return n
}

func (m *TaskDoneAck) Size() (n int) {
	var l int
	_ = l
	if m.Goods != nil {
		l = m.Goods.Size()
		n += 1 + l + sovTask(uint64(l))
	}
	return n
}

func (m *TaskNpcStateReq) Size() (n int) {
	var l int
	_ = l
	return n
}

func (m *TaskNpcStateAck) Size() (n int) {
	var l int
	_ = l
	if m.TaskId != 0 {
		n += 1 + sovTask(uint64(m.TaskId))
	}
	if m.Process != 0 {
		n += 1 + sovTask(uint64(m.Process))
	}
	return n
}

func sovTask(x uint64) (n int) {
	for {
		n++
		x >>= 7
		if x == 0 {
			break
		}
	}
	return n
}
func sozTask(x uint64) (n int) {
	return sovTask(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *TaskDoneReq) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTask
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
			return fmt.Errorf("proto: TaskDoneReq: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: TaskDoneReq: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field TaskId", wireType)
			}
			m.TaskId = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTask
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.TaskId |= (int32(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipTask(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthTask
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
func (m *TaskDoneAck) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTask
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
			return fmt.Errorf("proto: TaskDoneAck: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: TaskDoneAck: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Goods", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTask
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
				return ErrInvalidLengthTask
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
			skippy, err := skipTask(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthTask
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
func (m *TaskNpcStateReq) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTask
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
			return fmt.Errorf("proto: TaskNpcStateReq: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: TaskNpcStateReq: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		default:
			iNdEx = preIndex
			skippy, err := skipTask(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthTask
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
func (m *TaskNpcStateAck) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTask
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
			return fmt.Errorf("proto: TaskNpcStateAck: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: TaskNpcStateAck: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field TaskId", wireType)
			}
			m.TaskId = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTask
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.TaskId |= (int32(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Process", wireType)
			}
			m.Process = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTask
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Process |= (int32(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipTask(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthTask
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
func skipTask(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowTask
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
					return 0, ErrIntOverflowTask
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
					return 0, ErrIntOverflowTask
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
				return 0, ErrInvalidLengthTask
			}
			return iNdEx, nil
		case 3:
			for {
				var innerWire uint64
				var start int = iNdEx
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return 0, ErrIntOverflowTask
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
				next, err := skipTask(dAtA[start:])
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
	ErrInvalidLengthTask = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowTask   = fmt.Errorf("proto: integer overflow")
)

func init() { proto.RegisterFile("task.proto", fileDescriptorTask) }

var fileDescriptorTask = []byte{
	// 186 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0x2a, 0x49, 0x2c, 0xce,
	0xd6, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x62, 0x2a, 0x48, 0x92, 0xe2, 0x49, 0xce, 0xcf, 0xcd,
	0xcd, 0xcf, 0x83, 0x88, 0x28, 0xa9, 0x72, 0x71, 0x87, 0x24, 0x16, 0x67, 0xbb, 0xe4, 0xe7, 0xa5,
	0x06, 0xa5, 0x16, 0x0a, 0x89, 0x71, 0xb1, 0x81, 0x94, 0x7b, 0xa6, 0x48, 0x30, 0x2a, 0x30, 0x6a,
	0xb0, 0x06, 0x41, 0x79, 0x4a, 0xe6, 0x08, 0x65, 0x8e, 0xc9, 0xd9, 0x42, 0x1a, 0x5c, 0xac, 0xe9,
	0xf9, 0xf9, 0x29, 0xc5, 0x60, 0x55, 0xdc, 0x46, 0x42, 0x7a, 0x05, 0x49, 0x7a, 0xee, 0x20, 0x01,
	0xe7, 0x8c, 0xc4, 0xbc, 0xf4, 0x54, 0xbf, 0x92, 0xb4, 0x20, 0x88, 0x02, 0x25, 0x41, 0x2e, 0x7e,
	0x90, 0x46, 0xbf, 0x82, 0xe4, 0xe0, 0x92, 0xc4, 0x12, 0x90, 0x1d, 0x4a, 0xce, 0xa8, 0x42, 0x20,
	0xf3, 0x70, 0x58, 0x2b, 0x24, 0xc1, 0xc5, 0x5e, 0x50, 0x94, 0x9f, 0x9c, 0x5a, 0x5c, 0x2c, 0xc1,
	0x04, 0x96, 0x80, 0x71, 0x9d, 0x04, 0x4e, 0x3c, 0x92, 0x63, 0xbc, 0xf0, 0x48, 0x8e, 0xf1, 0xc1,
	0x23, 0x39, 0xc6, 0x19, 0x8f, 0xe5, 0x18, 0x92, 0xd8, 0xc0, 0x1e, 0x32, 0x06, 0x04, 0x00, 0x00,
	0xff, 0xff, 0xb2, 0xf7, 0xe5, 0x70, 0xf0, 0x00, 0x00, 0x00,
}