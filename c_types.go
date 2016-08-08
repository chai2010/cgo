// Copyright 2016 <chaishushan{AT}gmail.com>. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package cgo

//#include <stdint.h>
//#include <stdlib.h>
import "C"
import "unsafe"

type (
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

func NewChar(firstValue int, moreValues ...int) *CChar {
	n := len(moreValues) + 1
	p := NewCharN(n)
	s := p.Slice(n)

	s[0] = byte(firstValue)
	for i, v := range moreValues {
		s[i+1] = byte(v)
	}
	return p
}
func NewCharN(n int) *CChar {
	p := C.calloc(C.size_t(n), C.size_t(unsafe.Sizeof(CChar(0))))
	return (*CChar)(p)
}
func (s *CChar) Slice(n int) []byte {
	return ((*[1 << 31]byte)(unsafe.Pointer(s)))[0:n:n]
}
func (p *CChar) Free() {
	C.free(unsafe.Pointer(p))
}

func NewInt(firstValue int, moreValues ...int) *CInt {
	n := len(moreValues) + 1
	p := NewIntN(n)
	s := p.Slice(n)

	s[0] = CInt(firstValue)
	for i, v := range moreValues {
		s[i+1] = CInt(v)
	}
	return p
}
func NewIntN(n int) *CInt {
	p := C.calloc(C.size_t(n), C.size_t(unsafe.Sizeof(CInt(0))))
	return (*CInt)(p)
}
func (s *CInt) Slice(n int) []CInt {
	return ((*[1 << 29]CInt)(unsafe.Pointer(s)))[0:n:n]
}
func (p *CInt) Free() {
	C.free(unsafe.Pointer(p))
}

func (s *CUint) Slice(n int) []CUint {
	return ((*[1 << 29]CUint)(unsafe.Pointer(s)))[0:n:n]
}

func (s *CFloat) Slice(n int) []float32 {
	return ((*[1 << 29]float32)(unsafe.Pointer(s)))[0:n:n]
}
func (s *CDouble) Slice(n int) []float64 {
	return ((*[1 << 28]float64)(unsafe.Pointer(s)))[0:n:n]
}

func (s *CSizeT) Slice(n int) []CSizeT {
	return ((*[1 << 28]CSizeT)(unsafe.Pointer(s)))[0:n:n]
}

func (s *CInt8) Slice(n int) []int8 {
	return ((*[1 << 31]int8)(unsafe.Pointer(s)))[0:n:n]
}
func (s *CUint8) Slice(n int) []uint8 {
	return ((*[1 << 31]uint8)(unsafe.Pointer(s)))[0:n:n]
}

func (s *CInt16) Slice(n int) []int16 {
	return ((*[1 << 30]int16)(unsafe.Pointer(s)))[0:n:n]
}
func (s *CUint16) Slice(n int) []uint16 {
	return ((*[1 << 30]uint16)(unsafe.Pointer(s)))[0:n:n]
}

func (s *CInt32) Slice(n int) []int32 {
	return ((*[1 << 29]int32)(unsafe.Pointer(s)))[0:n:n]
}
func (s *CUint32) Slice(n int) []uint32 {
	return ((*[1 << 29]uint32)(unsafe.Pointer(s)))[0:n:n]
}

func (s *CInt64) Slice(n int) []int64 {
	return ((*[1 << 28]int64)(unsafe.Pointer(s)))[0:n:n]
}
func (s *CUint64) Slice(n int) []uint64 {
	return ((*[1 << 28]uint64)(unsafe.Pointer(s)))[0:n:n]
}

func (s UnsafePointer) Slice(n int) []byte {
	return ((*[1 << 31]byte)(unsafe.Pointer(s)))[0:n:n]
}
