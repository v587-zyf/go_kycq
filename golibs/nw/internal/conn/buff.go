package conn

import (
	"bytes"
)

type DataBuff struct {
	buff     *bytes.Buffer
	buffSize int
	enabled  bool
}

const MinMergedWriteBuffSize = 100 * 1024

func NewDataBuff(buffSize int, enabled bool) *DataBuff {
	if !enabled {
		return &DataBuff{enabled: enabled}
	}
	return &DataBuff{
		buff:     bytes.NewBuffer(make([]byte, buffSize+2*1024)), // buff cap is a little bigger than buffSize
		buffSize: buffSize,
	}
}

// GetData return merged buff from channel c
func (this *DataBuff) GetData(data []byte, c <-chan []byte) ([]byte, int) {
	if !this.enabled || len(c) == 0 {
		return data, 1
	}

	buff, buffSize, count := this.buff, this.buffSize, 0
	buff.Reset()
	for {
		count++
		buff.Write(data)
		if len(c) == 0 || buff.Len() >= buffSize {
			break
		}
		data = <-c
	}
	return buff.Bytes(), count
}
