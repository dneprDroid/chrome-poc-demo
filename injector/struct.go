package injector 

import (
	"bytes"
	"encoding/binary"
)

type StructDump struct {
	data []byte
	cursor int
}

func (self *StructDump) Bytes() []byte {
	return self.data
}

func NewStructDump(data []byte) *StructDump {
	return &StructDump{data: data}
}

func (self *StructDump) WriteBytes(data []byte) {
	self.cursor += len(data)
	self.data = append(self.data, data...)
}

func (self *StructDump) WriteUInt16(number uint16, bigEndian bool) {
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

func (self *StructDump) WriteUint32(number uint32, bigEndian bool) {
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

func (self *StructDump) WriteUint64(number uint64, bigEndian bool) {
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

func (self *StructDump) WriteSigned(value interface{}, bigEndian bool) error {
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

func (self *StructDump) WriteInt32(number int32, bigEndian bool) error {
	return self.WriteSigned(&number, bigEndian)
}

func (self *StructDump) WriteInt64(number int64, bigEndian bool) error {
	return self.WriteSigned(&number, bigEndian)
}