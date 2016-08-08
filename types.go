// Copyright 2016 <chaishushan{AT}gmail.com>. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package cgo

//#include <stdint.h>
//#include <stdlib.h>
import "C"
import "unsafe"

type (
	Char   C.char
	Int    C.int
	UInt   C.uint
	Float  C.float
	Double C.double
	SizeT  C.size_t

	Int8   C.int8_t
	Int16  C.int16_t
	Int32  C.int32_t
	Int64  C.int64_t
	IntPtr C.intptr_t

	UInt8   C.uint8_t
	UInt16  C.uint16_t
	UInt32  C.uint32_t
	UInt64  C.uint64_t
	UIntPtr C.uintptr_t

	UnsafePointer uintptr
)

func NewChar(firstValue int, moreValues ...int) *Char {
	n := len(moreValues) + 1
	p := NewCharN(n)
	s := p.Slice(n)

	s[0] = byte(firstValue)
	for i, v := range moreValues {
		s[i+1] = byte(v)
	}
	return p
}
func NewCharN(n int) *Char {
	p := C.calloc(C.size_t(n), C.size_t(unsafe.Sizeof(Char(0))))
	return (*Char)(p)
}
func (s *Char) Slice(n int) []byte {
	return ((*[1 << 31]byte)(unsafe.Pointer(s)))[0:n:n]
}
func (s *Char) Free() {
	C.free(unsafe.Pointer(s))
}

func NewInt(firstValue int, moreValues ...int) *Int {
	n := len(moreValues) + 1
	p := NewIntN(n)
	s := p.Slice(n)

	s[0] = Int(firstValue)
	for i, v := range moreValues {
		s[i+1] = Int(v)
	}
	return p
}
func NewIntN(n int) *Int {
	p := C.calloc(C.size_t(n), C.size_t(unsafe.Sizeof(Int(0))))
	return (*Int)(p)
}
func (s *Int) Slice(n int) []Int {
	return ((*[1 << 29]Int)(unsafe.Pointer(s)))[0:n:n]
}
func (p *Int) Free() {
	C.free(unsafe.Pointer(p))
}

func (s *UInt) Slice(n int) []UInt {
	return ((*[1 << 29]UInt)(unsafe.Pointer(s)))[0:n:n]
}

func (s *Float) Slice(n int) []float32 {
	return ((*[1 << 29]float32)(unsafe.Pointer(s)))[0:n:n]
}
func (s *Double) Slice(n int) []float64 {
	return ((*[1 << 28]float64)(unsafe.Pointer(s)))[0:n:n]
}

func (s *SizeT) Slice(n int) []SizeT {
	return ((*[1 << 28]SizeT)(unsafe.Pointer(s)))[0:n:n]
}

func (s *Int8) Slice(n int) []int8 {
	return ((*[1 << 31]int8)(unsafe.Pointer(s)))[0:n:n]
}
func (s *UInt8) Slice(n int) []uint8 {
	return ((*[1 << 31]uint8)(unsafe.Pointer(s)))[0:n:n]
}

func (s *Int16) Slice(n int) []int16 {
	return ((*[1 << 30]int16)(unsafe.Pointer(s)))[0:n:n]
}
func (s *UInt16) Slice(n int) []uint16 {
	return ((*[1 << 30]uint16)(unsafe.Pointer(s)))[0:n:n]
}

func (s *Int32) Slice(n int) []int32 {
	return ((*[1 << 29]int32)(unsafe.Pointer(s)))[0:n:n]
}
func (s *UInt32) Slice(n int) []uint32 {
	return ((*[1 << 29]uint32)(unsafe.Pointer(s)))[0:n:n]
}

func (s *Int64) Slice(n int) []int64 {
	return ((*[1 << 28]int64)(unsafe.Pointer(s)))[0:n:n]
}
func (s *UInt64) Slice(n int) []uint64 {
	return ((*[1 << 28]uint64)(unsafe.Pointer(s)))[0:n:n]
}
