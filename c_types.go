// Copyright 2016 <chaishushan{AT}gmail.com>. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package cgo

//#include <stdbool.h>
//#include <stdint.h>
import "C"
import "unsafe"

type (
	CBool   C.bool
	CChar   C.char
	CInt    C.int
	CUint   C.uint
	CFloat  C.float
	CDouble C.double
	CSizeT  C.size_t

	CInt8   C.int8_t
	CInt16  C.int16_t
	CInt32  C.int32_t
	CInt64  C.int64_t
	CIntPtr C.intptr_t

	CUint8   C.uint8_t
	CUint16  C.uint16_t
	CUint32  C.uint32_t
	CUint64  C.uint64_t
	CUIntPtr C.uintptr_t

	UnsafePointer uintptr
)

func (s *CBool) GoSliceNoCopy(n int) []CBool {
	return ((*[1 << 31]CBool)(unsafe.Pointer(s)))[0:n:n]
}
func (s *CChar) GoSliceNoCopy(n int) []byte {
	return ((*[1 << 31]byte)(unsafe.Pointer(s)))[0:n:n]
}

func (s *CInt) GoSliceNoCopy(n int) []CInt {
	return ((*[1 << 29]CInt)(unsafe.Pointer(s)))[0:n:n]
}
func (s *CUint) GoSliceNoCopy(n int) []CUint {
	return ((*[1 << 29]CUint)(unsafe.Pointer(s)))[0:n:n]
}

func (s *CFloat) GoSliceNoCopy(n int) []float32 {
	return ((*[1 << 29]float32)(unsafe.Pointer(s)))[0:n:n]
}
func (s *CDouble) GoSliceNoCopy(n int) []float64 {
	return ((*[1 << 28]float64)(unsafe.Pointer(s)))[0:n:n]
}

func (s *CSizeT) GoSliceNoCopy(n int) []CSizeT {
	return ((*[1 << 28]CSizeT)(unsafe.Pointer(s)))[0:n:n]
}

func (s *CInt8) GoSliceNoCopy(n int) []int8 {
	return ((*[1 << 31]int8)(unsafe.Pointer(s)))[0:n:n]
}
func (s *CUint8) GoSliceNoCopy(n int) []uint8 {
	return ((*[1 << 31]uint8)(unsafe.Pointer(s)))[0:n:n]
}

func (s *CInt16) GoSliceNoCopy(n int) []int16 {
	return ((*[1 << 30]int16)(unsafe.Pointer(s)))[0:n:n]
}
func (s *CUint16) GoSliceNoCopy(n int) []uint16 {
	return ((*[1 << 30]uint16)(unsafe.Pointer(s)))[0:n:n]
}

func (s *CInt32) GoSliceNoCopy(n int) []int32 {
	return ((*[1 << 29]int32)(unsafe.Pointer(s)))[0:n:n]
}
func (s *CUint32) GoSliceNoCopy(n int) []uint32 {
	return ((*[1 << 29]uint32)(unsafe.Pointer(s)))[0:n:n]
}

func (s *CInt64) GoSliceNoCopy(n int) []int64 {
	return ((*[1 << 28]int64)(unsafe.Pointer(s)))[0:n:n]
}
func (s *CUint64) GoSliceNoCopy(n int) []uint64 {
	return ((*[1 << 28]uint64)(unsafe.Pointer(s)))[0:n:n]
}

func (s UnsafePointer) GoBytesNoCopy(n int) []byte {
	return ((*[1 << 31]byte)(unsafe.Pointer(s)))[0:n:n]
}
