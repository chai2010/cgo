// Copyright 2016 <chaishushan{AT}gmail.com>. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package cgo

/*
#cgo CXXFLAGS: -std=c++11

#include <stdint.h>
#include <stdlib.h>
#include <string.h>

*/
import "C"
import "unsafe"

type (
	CChar   C.char
	CInt    C.int
	CUint   C.uint
	CFloat  C.float
	CDouble C.double
	CSizeT  C.size_t

	CUint8  C.uint8_t
	CUint16 C.uint16_t
	CUint32 C.uint32_t
	CUint64 C.uint64_t

	CInt8  C.int8_t
	CInt16 C.int16_t
	CInt32 C.int32_t
	CInt64 C.int64_t

	CIntPtr  C.intptr_t
	CUIntPtr C.uintptr_t
)

var CEmptyString = CString("")

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

func CMalloc(n int) unsafe.Pointer {
	return C.malloc(C.size_t(n))
}

func CFree(p unsafe.Pointer) {
	C.free(p)
}

func CStrdup(s *CChar) *CChar {
	if s == nil {
		s = CEmptyString
	}

	d := (*CChar)(C.malloc(C.strlen((*C.char)(s)) + 1))
	if d == nil {
		return nil
	}
	C.strcpy((*C.char)(d), (*C.char)(s))
	return d
}

func CStrFree(p *CChar) {
	if p == nil || p == CEmptyString {
		return
	}
	C.free(unsafe.Pointer(p))
}
