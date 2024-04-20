package injector

const (
	FileEOF_FLAG_HAS_CRC_32 = (uint32(1) << 0)
	FileEOF_FLAG_HAS_KEY_SHA256 = (uint32(1) << 1)
 
	kFileFinalMagicNumber = uint64(0xf4fa6f45970d41d8)
 )

type FileEOF struct {
	final_magic_number uint64
	flags uint32
	data_crc32 uint32
	stream_size uint32
}

func (self *FileEOF) toBytes() []byte {
	sd := NewStructDump([]byte{})
	bigEndian := false
	sd.WriteUint64(self.final_magic_number, bigEndian)
	sd.WriteUint32(self.flags, bigEndian)
	sd.WriteUint32(self.data_crc32, bigEndian)
	sd.WriteUint32(self.stream_size, bigEndian)
	sd.WriteUint32(0, bigEndian)
	return sd.Bytes()
}

func fileEofData(streamData []byte, streamIndex int) []byte {
	hasCrc32 := true 
	entryStatDataSize := len(streamData)
	dataCrc32 := crc32hash(streamData)

	eofRecord := &FileEOF{}
	eofRecord.stream_size = uint32(entryStatDataSize)
	eofRecord.final_magic_number = kFileFinalMagicNumber
	eofRecord.flags = 0

	if (hasCrc32) {
		eofRecord.flags |= FileEOF_FLAG_HAS_CRC_32
	}
	if (streamIndex == 0) {
		eofRecord.flags |= FileEOF_FLAG_HAS_KEY_SHA256
	}
	eofRecord.data_crc32 = dataCrc32;
	return eofRecord.toBytes()
}