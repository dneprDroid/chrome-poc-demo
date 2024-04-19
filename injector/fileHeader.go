
package injector
 
const (
	kFileHeaderInitialMagicNumber = uint64(0xfcfb6d1ba7725c30)
	kFileHeaderEntryVersionOnDisk = uint32(5)
)

type FileHeader struct {
	initial_magic_number uint64
	version uint32
	key_length uint32
	key_hash uint32
}
 
func (self *FileHeader) toBytes() []byte {
	 pickle := NewPickle()
	 bigEndian := false
	 pickle.WriteUint64(self.initial_magic_number, bigEndian)
	 pickle.WriteUint32(self.version, bigEndian)
	 pickle.WriteUint32(self.key_length, bigEndian)
	 pickle.WriteUint32(self.key_hash, bigEndian)
	 return pickle.Bytes()
 }
 
func fileHeader(key string) []byte {
	header := &FileHeader{}
	header.initial_magic_number = kFileHeaderInitialMagicNumber
	header.version = kFileHeaderEntryVersionOnDisk
 
	keyData := []byte(key)
	header.key_length = uint32(len(keyData))
	header.key_hash = persistentHash(keyData)
	return header.toBytes()
 }
