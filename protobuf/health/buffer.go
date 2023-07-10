package health

type Buffer struct {
	buf      []byte
	pos      int
	isVarInt bool
	intRW    *IntReadWriter
	Err      error
}

const (
	VAR_INT_MAGIC_ENABLED  = 0xaa
	VAR_INT_MAGIC_DISABLED = 0x99
)

const (
	TYPE_ID_VARINT           = 0
	TYPE_ID_1_BYTE           = 1
	TYPE_ID_2_BYTE           = 2
	TYPE_ID_4_BYTE           = 3
	TYPE_ID_8_BYTE           = 4
	TYPE_ID_LENGTH_DELIMITED = 5
)

func NewBuffer(buf []byte, isVarInt bool) *Buffer {
	var intRW = NoVarIntReadWriter
	if isVarInt {
		intRW = VarIntReadWriter
	}
	return &Buffer{
		buf:      buf,
		pos:      0,
		Err:      nil,
		isVarInt: isVarInt,
		intRW:    intRW,
	}
}

func (this *Buffer) Pos() int {
	return this.pos
}

func (this *Buffer) SkipField(typ uint32) {
	switch typ {
	case TYPE_ID_VARINT:
		_, n := VarIntReadWriter.ReadUint64(this.buf[this.pos:])
		this.pos += n
	case TYPE_ID_1_BYTE:
		this.pos += 1
	case TYPE_ID_2_BYTE:
		this.pos += 2
	case TYPE_ID_4_BYTE:
		this.pos += 4
	case TYPE_ID_8_BYTE:
		this.pos += 8
	case TYPE_ID_LENGTH_DELIMITED:
		len, n := NoVarIntReadWriter.ReadUint32(this.buf[this.pos:])
		this.pos += int(len) + n
	default:
	}
}

// func (this *Buffer) WriteBegin() {
// 	if this.isVarInt {
// 		this.WriteUint8(VAR_INT_MAGIC_ENABLED)
// 	} else {
// 		this.WriteUint8(VAR_INT_MAGIC_DISABLED)
// 	}
// 	this.pos += 4 // 预留长度字段
// }

// func (this *Buffer) WriteEnd() {
// 	NoVarIntReadWriter.WriteUint32(this.buf[1:], uint32(this.pos))
// 	this.pos += 4
// }

func (this *Buffer) ReadTag() (uint32, uint32) {
	tag, len := VarIntReadWriter.ReadUint32(this.buf[this.pos:])
	this.pos += len
	return tag >> 4, tag & 0xF
}

func (this *Buffer) WriteTag(seq uint32, typeId uint32) *Buffer {
	if this.isVarInt && typeId >= 2 && typeId <= 4 {
		typeId = TYPE_ID_VARINT
	}
	len := WriteVarUint32(this.buf[this.pos:], MakeTag(seq, typeId))
	this.pos += len
	return this
}

func (this *Buffer) ReadLength() uint32 {
	value, n := NoVarIntReadWriter.ReadUint32(this.buf[this.pos:])
	this.pos += n
	return value
}

func (this *Buffer) WriteLength(pos int, n int) {
	NoVarIntReadWriter.WriteUint32(this.buf[pos:], uint32(n))
}

func (this *Buffer) ReserveLength() int {
	oldPos := this.pos
	this.pos += 4
	return oldPos
}

func (this *Buffer) ReadInt8() int8 {
	value := int8(this.buf[this.pos])
	this.pos += 1
	return value
}

func (this *Buffer) WriteInt8(value int8) {
	this.buf[this.pos] = byte(value)
	this.pos += 1
}

func (this *Buffer) ReadInt16() int16 {
	value, len := this.intRW.ReadInt16(this.buf[this.pos:])
	this.pos += len
	return value
}

func (this *Buffer) WriteInt16(value int16) {
	len := this.intRW.WriteInt16(this.buf[this.pos:], value)
	this.pos += len
}

func (this *Buffer) ReadInt32() int32 {
	value, len := this.intRW.ReadInt32(this.buf[this.pos:])
	this.pos += len
	return value
}

func (this *Buffer) WriteInt32(value int32) {
	len := this.intRW.WriteInt32(this.buf[this.pos:], value)
	this.pos += len
}

func (this *Buffer) ReadInt64() int64 {
	value, len := this.intRW.ReadInt64(this.buf[this.pos:])
	this.pos += len
	return value
}

func (this *Buffer) WriteInt64(value int64) {
	len := this.intRW.WriteInt64(this.buf[this.pos:], value)
	this.pos += len
}

func (this *Buffer) ReadUint8() uint8 {
	value := uint8(this.buf[this.pos])
	this.pos += 1
	return value
}

func (this *Buffer) WriteUint8(value uint8) {
	this.buf[this.pos] = byte(value)
	this.pos += 1
}

func (this *Buffer) ReadUint16() uint16 {
	value, len := this.intRW.ReadUint16(this.buf[this.pos:])
	this.pos += len
	return value
}

func (this *Buffer) WriteUint16(value uint16) {
	len := this.intRW.WriteUint16(this.buf[this.pos:], value)
	this.pos += len
}

func (this *Buffer) ReadUint32() uint32 {
	value, len := this.intRW.ReadUint32(this.buf[this.pos:])
	this.pos += len
	return value
}

func (this *Buffer) WriteUint32(value uint32) {
	len := this.intRW.WriteUint32(this.buf[this.pos:], value)
	this.pos += len
}

func (this *Buffer) ReadUint64() uint64 {
	value, len := this.intRW.ReadUint64(this.buf[this.pos:])
	this.pos += len
	return value
}

func (this *Buffer) WriteUint64(value uint64) {
	len := this.intRW.WriteUint64(this.buf[this.pos:], value)
	this.pos += len
}

func (this *Buffer) ReadString() string {
	strLen, n := NoVarIntReadWriter.ReadUint32(this.buf[this.pos:])
	this.pos += n
	if strLen == 0 {
		return ""
	}
	var rb = make([]byte, strLen)
	copy(rb, this.buf[this.pos:this.pos+int(strLen)])
	this.pos += int(strLen)
	return string(rb)
}

func (this *Buffer) WriteString(str string) {
	strLen := len(str)
	n := NoVarIntReadWriter.WriteUint32(this.buf[this.pos:], uint32(strLen))
	this.pos += n
	copy(this.buf[this.pos:], []byte(str))
	this.pos += strLen
}

//func (this *Buffer) WriteBlock(msg Message) {
//	oldPos := this.pos
//	this.pos += 4
//	n, err := msg.MarshalTo(this.buf[this.pos:], this.isVarInt)
//	if err != nil {
//		this.Err = err
//		return
//	}
//	this.pos += n
//	this.WriteLength(oldPos, n)
//}
//
//func (this *Buffer) ReadBlock(msg Message) {
//	msgLen := int(this.ReadLength())
//	err := msg.Unmarshal(this.buf[this.pos:this.pos+msgLen], this.isVarInt)
//	if err != nil {
//		this.Err = err
//	} else {
//		this.pos += msgLen
//	}
//}
