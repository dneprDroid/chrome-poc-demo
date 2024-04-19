package injector 

import (
	"io"
	"crypto/sha1"
	"encoding/binary"
)

func entryHash(cacheKey string) uint64 {
	h := sha1.New()
	io.WriteString(h, cacheKey)
	hashBytes := h.Sum(nil)
	eHash := binary.LittleEndian.Uint64(hashBytes)
	return eHash
}