package injector 

import (
	"bytes"
	"encoding/binary"
)

type StructDump struct {
	data []byte
	cursor int
}

func NewStructDump(data []byte) *StructDump {
	return &StructDump{data: data}
}

func (self *StructDump) Bytes() []byte {
	return self.data
}

func (self *StructDump) WriteBytes(data []byte) {
	self.cursor += len(data)
	self.data = append(self.data, data...)
}

func (self *StructDump) WriteUInt16(number uint16, bigEndian bool) {
	data := make([]byte, 2)
	byteOrder(bigEndian).PutUint16(data, number) 

	self.WriteBytes(data)
}

func (self *StructDump) WriteUint32(number uint32, bigEndian bool) {
	data := make([]byte, 4)
	byteOrder(bigEndian).PutUint32(data, number) 

	self.WriteBytes(data)
}

func (self *StructDump) WriteUint64(number uint64, bigEndian bool) {
	data := make([]byte, 8)
	byteOrder(bigEndian).PutUint64(data, number) 

	self.WriteBytes(data)
}

func (self *StructDump) WriteSigned(value interface{}, bigEndian bool) error {
	var b bytes.Buffer
	err := binary.Write(&b, byteOrder(bigEndian), value)
	if err != nil {
		return err 
	}
	self.WriteBytes(b.Bytes())
	return nil 
}

func (self *StructDump) WriteInt32(number int32, bigEndian bool) error {
	return self.WriteSigned(&number, bigEndian)
}

func (self *StructDump) WriteInt64(number int64, bigEndian bool) error {
	return self.WriteSigned(&number, bigEndian)
}