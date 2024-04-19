package injector 
/*

#include <stdlib.h>
#include <stdio.h>
#include <string.h>
#include <assert.h>

#include "exports.h"

*/
import "C"
import (
	"io"
	"unsafe"
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
 
 // https://github.com/chromium/chromium/blob/0b8aa3130ef9da743fa513b55b231965bd8a5573/base/third_party/superfasthash/superfasthash.c#L41
 func persistentHash(data []byte) uint32 {
	ret := C.SuperFastHash(
		unsafe.Pointer(&data[0]), C.int(len(data)),
	)
	return uint32(ret)
}