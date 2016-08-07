// Copyright 2016 <chaishushan{AT}gmail.com>. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package cgo

//#include <stdlib.h>
import "C"

import (
	"unsafe"
)

func CMalloc(n int) unsafe.Pointer {
	return C.malloc(C.size_t(n))
}

func CCalloc(num, size int) unsafe.Pointer {
	return C.calloc(C.size_t(num), C.size_t(size))
}

func CRealloc(p unsafe.Pointer, newSize int) unsafe.Pointer {
	return C.realloc(p, C.size_t(newSize))
}

func CFree(p unsafe.Pointer) {
	C.free(p)
}
