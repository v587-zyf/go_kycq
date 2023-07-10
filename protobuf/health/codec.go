package health

import (
	"bytes"
	"encoding/binary"
)

func ZigZagEncodeInt16(value int16) uint16 {
	return uint16(((value) >> 15) ^ ((value) << 1))
}

func ZigZagDecodeInt16(value uint16) int16 {
	return int16((value >> 1) ^ -(value & 1))
}

func ZigZagEncodeInt32(value int32) uint32 {
	return uint32(((value) >> 31) ^ ((value) << 1))
}

func ZigZagDecodeInt32(value uint32) int32 {
	return int32((value >> 1) ^ -(value & 1))
}

func ZigZagEncodeInt64(value int64) uint64 {
	return uint64(((value) >> 63) ^ ((value) << 1))
}

func ZigZagDecodeInt64(value uint64) int64 {
	return int64((value >> 1) ^ -(value & 1))
}

func EncodeUint16x(buff *bytes.Buffer, value uint16) {
	EncodeUint64x(buff, uint64(value))
}

func EncodeUint32x(buff *bytes.Buffer, value uint32) {
	EncodeUint64x(buff, uint64(value))
}

func EncodeUint64x(buff *bytes.Buffer, value uint64) {
	for value > 0 {
		b := byte(value & 0x7F)
		value = value >> 7
		if value != 0 {
			b |= 0x80
		}
		buff.WriteByte(b)
	}
}

func DecodeUint16x(buff *bytes.Buffer) uint16 {
	return uint16(DecodeUint64x(buff))
}

func DecodeUint32x(buff *bytes.Buffer) uint32 {
	return uint32(DecodeUint64x(buff))
}

func DecodeUint64x(buff *bytes.Buffer) uint64 {
	var value uint64 = 0
	for i := 0; i < 100; i++ {
		b, _ := buff.ReadByte()
		value |= uint64(b&0x7F) << uint64(i*7)
		if (b & 0x80) == 0 {
			break
		}
	}
	return value
}

func EncodeUint16(buff []byte, value uint16) int {
	return EncodeUint64(buff, uint64(value))
}

func EncodeUint32(buff []byte, value uint32) int {
	return EncodeUint64(buff, uint64(value))
}

func EncodeUint64(buff []byte, value uint64) int {
	var i = 0
	for {
		b := byte(value & 0x7F)
		value = value >> 7
		if value != 0 {
			b |= 0x80
		}
		buff[i] = b
		i++
		if value == 0 {
			break
		}
	}
	return i
}

func DecodeUint16(buff []byte) (uint16, int) {
	value, len := DecodeUint64(buff)
	return uint16(value), len
}

func DecodeUint32(buff []byte) (uint32, int) {
	value, len := DecodeUint64(buff)
	return uint32(value), len
}

func DecodeUint64(buff []byte) (uint64, int) {
	var value uint64 = 0
	var len = len(buff)
	var i = 0
	for i < len {
		b := buff[i]
		value |= uint64(b&0x7F) << uint64(i*7)
		i++
		if (b & 0x80) == 0 {
			break
		}
	}
	return value, i
}

func ReadInt16(buff []byte) (int16, int) {
	value := int16(binary.BigEndian.Uint16(buff))
	return value, 2
}

func ReadUint16(buff []byte) (uint16, int) {
	value := binary.BigEndian.Uint16(buff)
	return value, 2
}

func ReadInt32(buff []byte) (int32, int) {
	value := int32(binary.BigEndian.Uint32(buff))
	return value, 4
}

func ReadUint32(buff []byte) (uint32, int) {
	value := binary.BigEndian.Uint32(buff)
	return value, 4
}

func ReadInt64(buff []byte) (int64, int) {
	value := int64(binary.BigEndian.Uint64(buff))
	return value, 8
}

func ReadUint64(buff []byte) (uint64, int) {
	value := binary.BigEndian.Uint64(buff)
	return value, 8
}

func WriteInt16(buff []byte, value int16) int {
	binary.BigEndian.PutUint16(buff, uint16(value))
	return 2
}

func WriteUint16(buff []byte, value uint16) int {
	binary.BigEndian.PutUint16(buff, value)
	return 2
}

func WriteInt32(buff []byte, value int32) int {
	binary.BigEndian.PutUint32(buff, uint32(value))
	return 4
}

func WriteUint32(buff []byte, value uint32) int {
	binary.BigEndian.PutUint32(buff, value)
	return 4
}

func WriteInt64(buff []byte, value int64) int {
	binary.BigEndian.PutUint64(buff, uint64(value))
	return 8
}

func WriteUint64(buff []byte, value uint64) int {
	binary.BigEndian.PutUint64(buff, value)
	return 8
}

func ReadVarInt16(buff []byte) (int16, int) {
	value, len := ReadVarUint16(buff)
	return ZigZagDecodeInt16(value), len
}

func ReadVarUint16(buff []byte) (uint16, int) {
	return DecodeUint16(buff)
}

func ReadVarInt32(buff []byte) (int32, int) {
	value, len := ReadVarUint32(buff)
	return ZigZagDecodeInt32(value), len
}

func ReadVarUint32(buff []byte) (uint32, int) {
	return DecodeUint32(buff)
}

func ReadVarInt64(buff []byte) (int64, int) {
	value, len := ReadVarUint64(buff)
	return ZigZagDecodeInt64(value), len
}

func ReadVarUint64(buff []byte) (uint64, int) {
	return DecodeUint64(buff)
}

func WriteVarInt16(buff []byte, value int16) int {
	return WriteVarUint16(buff, ZigZagEncodeInt16(value))
}

func WriteVarUint16(buff []byte, value uint16) int {
	return EncodeUint16(buff, value)
}

func WriteVarInt32(buff []byte, value int32) int {
	return WriteVarUint32(buff, ZigZagEncodeInt32(value))
}

func WriteVarUint32(buff []byte, value uint32) int {
	return EncodeUint32(buff, value)
}

func WriteVarInt64(buff []byte, value int64) int {
	return WriteVarUint64(buff, ZigZagEncodeInt64(value))
}

func WriteVarUint64(buff []byte, value uint64) int {
	return EncodeUint64(buff, value)
}

func MakeTag(id uint32, typeId uint32) uint32 {
	return ((id << 4) + typeId)
}

type IntReadWriter struct {
	ReadInt16  func(buff []byte) (int16, int)
	ReadUint16 func(buff []byte) (uint16, int)
	ReadInt32  func(buff []byte) (int32, int)
	ReadUint32 func(buff []byte) (uint32, int)
	ReadInt64  func(buff []byte) (int64, int)
	ReadUint64 func(buff []byte) (uint64, int)

	WriteInt16  func(buff []byte, value int16) int
	WriteUint16 func(buff []byte, value uint16) int
	WriteInt32  func(buff []byte, value int32) int
	WriteUint32 func(buff []byte, value uint32) int
	WriteInt64  func(buff []byte, value int64) int
	WriteUint64 func(buff []byte, value uint64) int
}

var VarIntReadWriter = &IntReadWriter{
	ReadInt16:   ReadVarInt16,
	ReadUint16:  ReadVarUint16,
	ReadInt32:   ReadVarInt32,
	ReadUint32:  ReadVarUint32,
	ReadInt64:   ReadVarInt64,
	ReadUint64:  ReadVarUint64,
	WriteInt16:  WriteVarInt16,
	WriteUint16: WriteVarUint16,
	WriteInt32:  WriteVarInt32,
	WriteUint32: WriteVarUint32,
	WriteInt64:  WriteVarInt64,
	WriteUint64: WriteVarUint64,
}
var NoVarIntReadWriter = &IntReadWriter{
	ReadInt16:   ReadInt16,
	ReadUint16:  ReadUint16,
	ReadInt32:   ReadInt32,
	ReadUint32:  ReadUint32,
	ReadInt64:   ReadInt64,
	ReadUint64:  ReadUint64,
	WriteInt16:  WriteInt16,
	WriteUint16: WriteUint16,
	WriteInt32:  WriteInt32,
	WriteUint32: WriteUint32,
	WriteInt64:  WriteInt64,
	WriteUint64: WriteUint64,
}
