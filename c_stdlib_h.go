// Copyright 2016 <chaishushan{AT}gmail.com>. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package cgo

//#include <stdlib.h>
import "C"

import (
	"unsafe"
)

func Malloc(n int) UnsafePointer {
	return UnsafePointer(C.malloc(C.size_t(n)))
}

func Calloc(num, size int) UnsafePointer {
	return UnsafePointer(C.calloc(C.size_t(num), C.size_t(size)))
}

func Realloc(p UnsafePointer, newSize int) UnsafePointer {
	return UnsafePointer(C.realloc(unsafe.Pointer(p), C.size_t(newSize)))
}

func (p UnsafePointer) Free() {
	C.free(unsafe.Pointer(p))
}
