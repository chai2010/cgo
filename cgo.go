// Copyright 2016 <chaishushan{AT}gmail.com>. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package cgo

/*
#cgo CXXFLAGS: -std=c++11

#include <stdbool.h>
#include <stdint.h>
#include <stdlib.h>
#include <string.h>
#include <stdio.h>

*/
import "C"

import (
	"unsafe"
)

// Go string to C string
// The C string is allocated in the C heap using malloc.
// It is the caller's responsibility to arrange for it to be
// freed, such as by calling C.free (be sure to include stdlib.h
// if C.free is needed).
func CString(s string) *CChar {
	return (*CChar)(C.CString(s))
}

// Go []byte slice to C array
// The C array is allocated in the C heap using malloc.
// It is the caller's responsibility to arrange for it to be
// freed, such as by calling C.free (be sure to include stdlib.h
// if C.free is needed).
func CBytes(s []byte) unsafe.Pointer {
	return C.CBytes(s)
}

// C string to Go string
func GoString(s *CChar) string {
	return C.GoString((*C.char)(s))
}

// C data with explicit length to Go string
func GoStringN(s *CChar, n int) string {
	return C.GoStringN((*C.char)(s), C.int(n))
}

// C data with explicit length to Go []byte
func GoBytes(s unsafe.Pointer, n int) []byte {
	return C.GoBytes(s, C.int(n))
}

func GoBytesNoCopy(s unsafe.Pointer, n int) []byte {
	return ((*[1 << 30]byte)(unsafe.Pointer(s)))[0:n:n]
}
