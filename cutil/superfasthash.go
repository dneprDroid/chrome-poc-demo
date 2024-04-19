package cutil
/*

#include <stdlib.h>
#include <stdio.h>
#include <string.h>
#include <assert.h>

#include "exports.h"

*/
import "C"
import (
	"unsafe"
)
 
 // https://github.com/chromium/chromium/blob/0b8aa3130ef9da743fa513b55b231965bd8a5573/base/third_party/superfasthash/superfasthash.c#L41
 func SuperFastHash(data []byte) uint32 {
	ret := C.SuperFastHash(
		unsafe.Pointer(&data[0]), C.int(len(data)),
	)
	return uint32(ret)
}