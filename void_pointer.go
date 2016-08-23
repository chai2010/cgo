// Copyright 2016 <chaishushan{AT}gmail.com>. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package cgo

//#include <string.h>
//#include <stdlib.h>
import "C"

import (
	"unsafe"
)

// Go []byte slice to C array
// The C array is allocated in the C heap using malloc.
// It is the caller's responsibility to arrange for it to be
// freed, such as by calling C.free (be sure to include stdlib.h
// if C.free is needed).
func NewBytes(s []byte) VoidPointer {
	p := Malloc(len(s))
	copy(p.ByteSlice(len(s)), s)
	return p
}

func Malloc(n int) VoidPointer {
	return VoidPointer(C.malloc(C.size_t(n)))
}

func Calloc(num, size int) VoidPointer {
	return VoidPointer(C.calloc(C.size_t(num), C.size_t(size)))
}

func (p *VoidPointer) Realloc(newSize int) VoidPointer {
	*p = VoidPointer(C.realloc(unsafe.Pointer(*p), C.size_t(newSize)))
	return *p
}

func (p VoidPointer) Free() {
	C.free(unsafe.Pointer(p))
}

// C string to Go string
func (p VoidPointer) GoString() string {
	return (*Char)(unsafe.Pointer(p)).GoString()
}

// C data with explicit length to Go string
func (p VoidPointer) GoStringN(n int) string {
	return (*Char)(unsafe.Pointer(p)).GoStringN(n)
}

// C data with explicit length to Go []byte
func (p VoidPointer) GoBytes(n int) []byte {
	return C.GoBytes(unsafe.Pointer(p), C.int(n))
}

func (p VoidPointer) ByteSlice(n int) []byte {
	return ((*[1 << 31]byte)(unsafe.Pointer(p)))[0:n:n]
}

func (p VoidPointer) Int8Slice(n int) []int8 {
	return ((*[1 << 31]int8)(unsafe.Pointer(p)))[0:n:n]
}
func (p VoidPointer) Int16Slice(n int) []int16 {
	return ((*[1 << 30]int16)(unsafe.Pointer(p)))[0:n:n]
}
func (p VoidPointer) Int32Slice(n int) []int32 {
	return ((*[1 << 29]int32)(unsafe.Pointer(p)))[0:n:n]
}
func (p VoidPointer) Int64Slice(n int) []int64 {
	return ((*[1 << 28]int64)(unsafe.Pointer(p)))[0:n:n]
}

func (p VoidPointer) UInt8Slice(n int) []uint8 {
	return ((*[1 << 31]uint8)(unsafe.Pointer(p)))[0:n:n]
}
func (p VoidPointer) UInt16Slice(n int) []uint16 {
	return ((*[1 << 30]uint16)(unsafe.Pointer(p)))[0:n:n]
}
func (p VoidPointer) UInt32Slice(n int) []uint32 {
	return ((*[1 << 29]uint32)(unsafe.Pointer(p)))[0:n:n]
}
func (p VoidPointer) UInt64Slice(n int) []uint64 {
	return ((*[1 << 28]uint64)(unsafe.Pointer(p)))[0:n:n]
}

func (p VoidPointer) Float32Slice(n int) []float32 {
	return ((*[1 << 29]float32)(unsafe.Pointer(p)))[0:n:n]
}
func (p VoidPointer) Float64Slice(n int) []float64 {
	return ((*[1 << 28]float64)(unsafe.Pointer(p)))[0:n:n]
}

func (p VoidPointer) Memcpy(src VoidPointer, num int) VoidPointer {
	return VoidPointer(C.memcpy(unsafe.Pointer(p), unsafe.Pointer(src), C.size_t(num)))
}

func (p VoidPointer) Memset(value, num int) VoidPointer {
	return VoidPointer(C.memset(unsafe.Pointer(p), C.int(value), C.size_t(num)))
}

func (p VoidPointer) Memmove(src VoidPointer, num int) VoidPointer {
	return VoidPointer(C.memmove(unsafe.Pointer(p), unsafe.Pointer(src), C.size_t(num)))
}

func (p VoidPointer) Memchr(value, num int) VoidPointer {
	return VoidPointer(C.memchr(unsafe.Pointer(p), C.int(value), C.size_t(num)))
}
