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

// C data with explicit length to Go []byte
func (s UnsafePointer) GoBytes(n int) []byte {
	return C.GoBytes(unsafe.Pointer(s), C.int(n))
}

func (s UnsafePointer) Slice(n int) []byte {
	return ((*[1 << 31]byte)(unsafe.Pointer(s)))[0:n:n]
}

func (s UnsafePointer) Int8Slice(n int) []int8 {
	return ((*[1 << 31]int8)(unsafe.Pointer(s)))[0:n:n]
}
func (s UnsafePointer) Int16Slice(n int) []int16 {
	return ((*[1 << 30]int16)(unsafe.Pointer(s)))[0:n:n]
}
func (s UnsafePointer) Int32Slice(n int) []int32 {
	return ((*[1 << 29]int32)(unsafe.Pointer(s)))[0:n:n]
}
func (s UnsafePointer) Int64Slice(n int) []int64 {
	return ((*[1 << 28]int64)(unsafe.Pointer(s)))[0:n:n]
}

func (s UnsafePointer) UInt8Slice(n int) []uint8 {
	return ((*[1 << 31]uint8)(unsafe.Pointer(s)))[0:n:n]
}
func (s UnsafePointer) UInt16Slice(n int) []uint16 {
	return ((*[1 << 30]uint16)(unsafe.Pointer(s)))[0:n:n]
}
func (s UnsafePointer) UInt32Slice(n int) []uint32 {
	return ((*[1 << 29]uint32)(unsafe.Pointer(s)))[0:n:n]
}
func (s UnsafePointer) UInt64Slice(n int) []uint64 {
	return ((*[1 << 28]uint64)(unsafe.Pointer(s)))[0:n:n]
}

func (s UnsafePointer) Float32Slice(n int) []float32 {
	return ((*[1 << 29]float32)(unsafe.Pointer(s)))[0:n:n]
}
func (s UnsafePointer) Float64Slice(n int) []float64 {
	return ((*[1 << 28]float64)(unsafe.Pointer(s)))[0:n:n]
}

func (p UnsafePointer) Memcpy(src UnsafePointer, num int) {
	C.memcpy(unsafe.Pointer(p), unsafe.Pointer(src), C.size_t(num))
}

func (p UnsafePointer) Memset(value, num int) {
	C.memset(unsafe.Pointer(p), C.int(value), C.size_t(num))
}

func (p UnsafePointer) Memmove(src UnsafePointer, num int) {
	C.memmove(unsafe.Pointer(p), unsafe.Pointer(src), C.size_t(num))
}

func (p UnsafePointer) Memchr(value, num int) {
	C.memchr(unsafe.Pointer(p), C.int(value), C.size_t(num))
}
