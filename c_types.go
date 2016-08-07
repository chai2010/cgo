// Copyright 2016 <chaishushan{AT}gmail.com>. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package cgo

//#include <stdbool.h>
//#include <stdint.h>
import "C"

type (
	CBool   C.bool
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
