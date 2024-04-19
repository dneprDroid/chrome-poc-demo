package injector 

import (
	"io"
	"crypto/sha1"
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