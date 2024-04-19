package injector 

import (
	"bytes"
	"encoding/binary"
)

const (
	SIZE_INT32 = 4
	SIZE_UINT32 = 4
	SIZE_INT64 = 8
	SIZE_UINT64 = 8
	SIZE_FLOAT = 4
	SIZE_DOUBLE = 8

	PAYLOAD_UNIT = 64
)

type Pickle struct {
	header []byte
	headerSize int 
	capacityAfterHeader int 
	writeOffset int
}

func NewPickle() *Pickle {
	p := &Pickle{}
	p.initEmpty()
	return p
}

func alignInt(i int, alignment int) int {
	return i + (alignment - (i % alignment)) % alignment
}

func (self *Pickle) initEmpty() {
	self.header = make([]byte, 0)
	self.headerSize = SIZE_UINT32
	self.capacityAfterHeader = 0
	self.writeOffset = 0
 	self.resize(PAYLOAD_UNIT)
	self.setPayloadSize(0)
}

func (self *Pickle) resize(newCapacity int) {
	newCapacity = alignInt(newCapacity, PAYLOAD_UNIT)
	self.header = append(self.header, make([]byte, newCapacity)...)
	self.capacityAfterHeader = newCapacity
}

func (self *Pickle) setPayloadSize(payloadSize int) {
	b := make([]byte, 4)
	binary.LittleEndian.PutUint32(b, uint32(payloadSize)) 
	copy(self.header[:len(b)], b)
}

func _max(a int, b int) int {
	if a > b {
		return a 
	}
	return b 
}

func (self *Pickle) WriteBytes(data []byte) bool {
	length := len(data)

	dataLength := alignInt(length, SIZE_UINT32)
	newSize := self.writeOffset + dataLength
	if (newSize > self.capacityAfterHeader) {
		self.resize(_max(self.capacityAfterHeader * 2, newSize))
	}
	startOffset := self.headerSize + self.writeOffset
	endOffset := startOffset + length
	copy(self.header[startOffset:], data)

	zerosEnd := endOffset + dataLength - length
	for i := endOffset; i <= zerosEnd; i+=1 {
		self.header[i] = 0
	}
	self.setPayloadSize(newSize)
	self.writeOffset = newSize
	return true
}

func (self *Pickle) getPayloadSize() int {
	r := binary.LittleEndian.Uint32(self.header)
	return int(r)
}

func (self *Pickle) Bytes() []byte {
	res := make([]byte, self.headerSize + self.getPayloadSize())
	copy(res, self.header[:len(res)])
	return res
}

func (self *Pickle) WriteString(value string) {
	valueBytes := []byte(value)
	self.WriteBytesString(valueBytes)
}

func (self *Pickle) WriteBytesString(valueBytes []byte) {
	self.WriteInt32(int32(len(valueBytes)), false);
	self.WriteBytes(valueBytes)
}

func (self *Pickle) WriteUInt16(number uint16, bigEndian bool) {
	data := func() []byte {
		b := make([]byte, 2)
		if bigEndian { 
			binary.BigEndian.PutUint16(b, number) 
		} else {
			binary.LittleEndian.PutUint16(b, number) 
		}
		return b 
	}()
	self.WriteBytes(data)
}

func (self *Pickle) WriteUint32(number uint32, bigEndian bool) {
	data := func() []byte {
		b := make([]byte, 4)
		if bigEndian { 
			binary.BigEndian.PutUint32(b, number) 
		} else {
			binary.LittleEndian.PutUint32(b, number) 
		}
		return b 
	}()
	self.WriteBytes(data)
}

func (self *Pickle) WriteUint64(number uint64, bigEndian bool) {
	data := func() []byte {
		b := make([]byte, 8)
		if bigEndian { 
			binary.BigEndian.PutUint64(b, number) 
		} else {
			binary.LittleEndian.PutUint64(b, number) 
		}
		return b 
	}()
	self.WriteBytes(data)
}

func (self *Pickle) WriteSigned(value interface{}, bigEndian bool) error {
	var b bytes.Buffer
	var order binary.ByteOrder
	if bigEndian {
		order = binary.BigEndian
	} else {
		order = binary.LittleEndian
	}
	err := binary.Write(&b, order, value)
	if err != nil {
		return err 
	}
	self.WriteBytes(b.Bytes())
	return nil 
}

func (self *Pickle) WriteInt32(number int32, bigEndian bool) error {
	return self.WriteSigned(&number, bigEndian)
}

func (self *Pickle) WriteInt64(number int64, bigEndian bool) error {
	return self.WriteSigned(&number, bigEndian)
}