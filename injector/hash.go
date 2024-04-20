package injector 

import (
	"io"
	"crypto/sha1"
	"crypto/sha256"
	"hash/crc32"
	"encoding/binary"

	"chrome-poc/cutil"
)

func entryHash(cacheKey string) uint64 {
	h := sha1.New()
	io.WriteString(h, cacheKey)
	hashBytes := h.Sum(nil)
	eHash := binary.LittleEndian.Uint64(hashBytes)
	return eHash
}
 
func persistentHash(data []byte) uint32 {
	return cutil.SuperFastHash(data)
}

func crc32hash(data []byte) uint32 {
	return crc32.ChecksumIEEE(data)
}

func metadataHash(key string) []byte {
	h := sha256.New()
	h.Write([]byte(key))
	hash := h.Sum(nil)
	return hash 
}