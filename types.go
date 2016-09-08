// Copyright 2016 <chaishushan{AT}gmail.com>. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package cgo

//#include <stdint.h>
//#include <stdlib.h>
import "C"
import "unsafe"

type (
	Bool   C.char
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

	VoidPointer uintptr
)

// -----------------------------------------------------------------------------

func bool2Byte(b bool) byte {
	if b {
		return 1
	}
	return 0
}

func NewBool(firstValue bool, moreValues ...bool) *Bool {
	n := len(moreValues) + 1
	p := NewBoolN(n)
	s := p.Slice(n)

	s[0] = bool2Byte(firstValue)
	for i, v := range moreValues {
		s[i+1] = bool2Byte(v)
	}
	return p
}

func NewBoolN(n int) *Bool {
	p := C.calloc(C.size_t(n), C.size_t(unsafe.Sizeof(Char(0))))
	return (*Bool)(p)
}

func (s *Bool) First() bool {
	return *s != 0
}

func (s *Bool) Slice(n int) []byte {
	return ((*[1 << 31]byte)(unsafe.Pointer(s)))[0:n:n]
}

func (s *Bool) Free() {
	C.free(unsafe.Pointer(s))
}

// -----------------------------------------------------------------------------

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

func (s *Char) First() byte {
	return byte(*s)
}

func (s *Char) Slice(n int) []byte {
	if n == 0 {
		n = s.Strlen()
	}
	return ((*[1 << 31]byte)(unsafe.Pointer(s)))[0:n:n]
}

func (s *Char) Free() {
	C.free(unsafe.Pointer(s))
}

// -----------------------------------------------------------------------------

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

func (p *Int) First() int {
	return int(*p)
}

func (p *Int) Slice(n int) []Int {
	return ((*[1 << 29]Int)(unsafe.Pointer(p)))[0:n:n]
}

func (p *Int) Free() {
	C.free(unsafe.Pointer(p))
}

// -----------------------------------------------------------------------------

func NewUInt(firstValue uint, moreValues ...uint) *UInt {
	n := len(moreValues) + 1
	p := NewUIntN(n)
	s := p.Slice(n)

	s[0] = UInt(firstValue)
	for i, v := range moreValues {
		s[i+1] = UInt(v)
	}
	return p
}

func NewUIntN(n int) *UInt {
	p := C.calloc(C.size_t(n), C.size_t(unsafe.Sizeof(UInt(0))))
	return (*UInt)(p)
}

func (p *UInt) First() uint {
	return uint(*p)
}

func (p *UInt) Slice(n int) []UInt {
	return ((*[1 << 29]UInt)(unsafe.Pointer(p)))[0:n:n]
}

func (p *UInt) Free() {
	C.free(unsafe.Pointer(p))
}

// -----------------------------------------------------------------------------

func NewFloat(firstValue float32, moreValues ...float32) *Float {
	n := len(moreValues) + 1
	p := NewFloatN(n)
	s := p.Slice(n)

	s[0] = firstValue
	for i, v := range moreValues {
		s[i+1] = v
	}
	return p
}

func NewFloatN(n int) *Float {
	p := C.calloc(C.size_t(n), C.size_t(unsafe.Sizeof(Float(0))))
	return (*Float)(p)
}

func (p *Float) First() float32 {
	return float32(*p)
}

func (p *Float) Slice(n int) []float32 {
	return ((*[1 << 29]float32)(unsafe.Pointer(p)))[0:n:n]
}

func (p *Float) Free() {
	C.free(unsafe.Pointer(p))
}

// -----------------------------------------------------------------------------

func NewDouble(firstValue float64, moreValues ...float64) *Double {
	n := len(moreValues) + 1
	p := NewDoubleN(n)
	s := p.Slice(n)

	s[0] = firstValue
	for i, v := range moreValues {
		s[i+1] = v
	}
	return p
}

func NewDoubleN(n int) *Double {
	p := C.calloc(C.size_t(n), C.size_t(unsafe.Sizeof(Double(0))))
	return (*Double)(p)
}

func (p *Double) First() float64 {
	return float64(*p)
}

func (p *Double) Slice(n int) []float64 {
	return ((*[1 << 28]float64)(unsafe.Pointer(p)))[0:n:n]
}

func (p *Double) Free() {
	C.free(unsafe.Pointer(p))
}

// -----------------------------------------------------------------------------

func NewSizeT(firstValue int, moreValues ...int) *SizeT {
	n := len(moreValues) + 1
	p := NewSizeTN(n)
	s := p.Slice(n)

	s[0] = SizeT(firstValue)
	for i, v := range moreValues {
		s[i+1] = SizeT(v)
	}
	return p
}

func NewSizeTN(n int) *SizeT {
	p := C.calloc(C.size_t(n), C.size_t(unsafe.Sizeof(SizeT(0))))
	return (*SizeT)(p)
}

func (p *SizeT) First() int {
	return int(*p)
}

func (p *SizeT) Slice(n int) []SizeT {
	return ((*[1 << 28]SizeT)(unsafe.Pointer(p)))[0:n:n]
}

func (p *SizeT) Free() {
	C.free(unsafe.Pointer(p))
}

// -----------------------------------------------------------------------------

func NewInt8(firstValue int8, moreValues ...int8) *Int8 {
	n := len(moreValues) + 1
	p := NewInt8N(n)
	s := p.Slice(n)

	s[0] = firstValue
	for i, v := range moreValues {
		s[i+1] = v
	}
	return p
}

func NewInt8N(n int) *Int8 {
	p := C.calloc(C.size_t(n), C.size_t(unsafe.Sizeof(Int8(0))))
	return (*Int8)(p)
}

func (p *Int8) First() int8 {
	return int8(*p)
}

func (p *Int8) Slice(n int) []int8 {
	return ((*[1 << 31]int8)(unsafe.Pointer(p)))[0:n:n]
}

func (p *Int8) Free() {
	C.free(unsafe.Pointer(p))
}

// -----------------------------------------------------------------------------

func NewUInt8(firstValue uint8, moreValues ...uint8) *UInt8 {
	n := len(moreValues) + 1
	p := NewUInt8N(n)
	s := p.Slice(n)

	s[0] = firstValue
	for i, v := range moreValues {
		s[i+1] = v
	}
	return p
}

func NewUInt8N(n int) *UInt8 {
	p := C.calloc(C.size_t(n), C.size_t(unsafe.Sizeof(UInt8(0))))
	return (*UInt8)(p)
}

func (p *UInt8) First() uint8 {
	return uint8(*p)
}

func (p *UInt8) Slice(n int) []uint8 {
	return ((*[1 << 31]uint8)(unsafe.Pointer(p)))[0:n:n]
}

func (p *UInt8) Free() {
	C.free(unsafe.Pointer(p))
}

// -----------------------------------------------------------------------------

func NewInt16(firstValue int16, moreValues ...int16) *Int16 {
	n := len(moreValues) + 1
	p := NewInt16N(n)
	s := p.Slice(n)

	s[0] = firstValue
	for i, v := range moreValues {
		s[i+1] = v
	}
	return p
}

func NewInt16N(n int) *Int16 {
	p := C.calloc(C.size_t(n), C.size_t(unsafe.Sizeof(Int16(0))))
	return (*Int16)(p)
}

func (p *Int16) First() int16 {
	return int16(*p)
}

func (p *Int16) Slice(n int) []int16 {
	return ((*[1 << 30]int16)(unsafe.Pointer(p)))[0:n:n]
}

func (p *Int16) Free() {
	C.free(unsafe.Pointer(p))
}

// -----------------------------------------------------------------------------

func NewUInt16(firstValue uint16, moreValues ...uint16) *UInt16 {
	n := len(moreValues) + 1
	p := NewUInt16N(n)
	s := p.Slice(n)

	s[0] = firstValue
	for i, v := range moreValues {
		s[i+1] = v
	}
	return p
}

func NewUInt16N(n int) *UInt16 {
	p := C.calloc(C.size_t(n), C.size_t(unsafe.Sizeof(UInt16(0))))
	return (*UInt16)(p)
}

func (p *UInt16) First() uint16 {
	return uint16(*p)
}

func (p *UInt16) Slice(n int) []uint16 {
	return ((*[1 << 30]uint16)(unsafe.Pointer(p)))[0:n:n]
}

func (p *UInt16) Free() {
	C.free(unsafe.Pointer(p))
}

// -----------------------------------------------------------------------------

func NewInt32(firstValue int32, moreValues ...int32) *Int32 {
	n := len(moreValues) + 1
	p := NewInt32N(n)
	s := p.Slice(n)

	s[0] = firstValue
	for i, v := range moreValues {
		s[i+1] = v
	}
	return p
}

func NewInt32N(n int) *Int32 {
	p := C.calloc(C.size_t(n), C.size_t(unsafe.Sizeof(Int32(0))))
	return (*Int32)(p)
}

func (p *Int32) First() int32 {
	return int32(*p)
}

func (p *Int32) Slice(n int) []int32 {
	return ((*[1 << 29]int32)(unsafe.Pointer(p)))[0:n:n]
}

func (p *Int32) Free() {
	C.free(unsafe.Pointer(p))
}

// -----------------------------------------------------------------------------

func NewUInt32(firstValue uint32, moreValues ...uint32) *UInt32 {
	n := len(moreValues) + 1
	p := NewUInt32N(n)
	s := p.Slice(n)

	s[0] = firstValue
	for i, v := range moreValues {
		s[i+1] = v
	}
	return p
}

func NewUInt32N(n int) *UInt32 {
	p := C.calloc(C.size_t(n), C.size_t(unsafe.Sizeof(UInt32(0))))
	return (*UInt32)(p)
}

func (p *UInt32) First() uint32 {
	return uint32(*p)
}

func (p *UInt32) Slice(n int) []uint32 {
	return ((*[1 << 29]uint32)(unsafe.Pointer(p)))[0:n:n]
}

func (p *UInt32) Free() {
	C.free(unsafe.Pointer(p))
}

// -----------------------------------------------------------------------------

func NewInt64(firstValue int64, moreValues ...int64) *Int64 {
	n := len(moreValues) + 1
	p := NewInt64N(n)
	s := p.Slice(n)

	s[0] = firstValue
	for i, v := range moreValues {
		s[i+1] = v
	}
	return p
}

func NewInt64N(n int) *Int64 {
	p := C.calloc(C.size_t(n), C.size_t(unsafe.Sizeof(Int64(0))))
	return (*Int64)(p)
}

func (p *Int64) First() int64 {
	return int64(*p)
}

func (p *Int64) Slice(n int) []int64 {
	return ((*[1 << 28]int64)(unsafe.Pointer(p)))[0:n:n]
}

func (p *Int64) Free() {
	C.free(unsafe.Pointer(p))
}

// -----------------------------------------------------------------------------

func NewUInt64(firstValue uint64, moreValues ...uint64) *UInt64 {
	n := len(moreValues) + 1
	p := NewUInt64N(n)
	s := p.Slice(n)

	s[0] = firstValue
	for i, v := range moreValues {
		s[i+1] = v
	}
	return p
}

func NewUInt64N(n int) *UInt64 {
	p := C.calloc(C.size_t(n), C.size_t(unsafe.Sizeof(UInt64(0))))
	return (*UInt64)(p)
}

func (p *UInt64) First() uint64 {
	return uint64(*p)
}

func (p *UInt64) Slice(n int) []uint64 {
	return ((*[1 << 28]uint64)(unsafe.Pointer(p)))[0:n:n]
}

func (p *UInt64) Free() {
	C.free(unsafe.Pointer(p))
}

// -----------------------------------------------------------------------------

func NewIntPtr(firstValue uintptr, moreValues ...uintptr) *IntPtr {
	n := len(moreValues) + 1
	p := NewIntPtrN(n)
	s := p.Slice(n)

	s[0] = IntPtr(firstValue)
	for i, v := range moreValues {
		s[i+1] = IntPtr(v)
	}
	return p
}

func NewIntPtrN(n int) *IntPtr {
	p := C.calloc(C.size_t(n), C.size_t(unsafe.Sizeof(IntPtr(0))))
	return (*IntPtr)(p)
}

func (p *IntPtr) First() uintptr {
	return uintptr(*p)
}

func (p *IntPtr) Slice(n int) []IntPtr {
	return ((*[1 << 28]IntPtr)(unsafe.Pointer(p)))[0:n:n]
}

func (p *IntPtr) Free() {
	C.free(unsafe.Pointer(p))
}

// -----------------------------------------------------------------------------

func NewUIntPtr(firstValue uintptr, moreValues ...uintptr) *UIntPtr {
	n := len(moreValues) + 1
	p := NewUIntPtrN(n)
	s := p.Slice(n)

	s[0] = UIntPtr(firstValue)
	for i, v := range moreValues {
		s[i+1] = UIntPtr(v)
	}
	return p
}

func NewUIntPtrN(n int) *UIntPtr {
	p := C.calloc(C.size_t(n), C.size_t(unsafe.Sizeof(UIntPtr(0))))
	return (*UIntPtr)(p)
}

func (p *UIntPtr) First() uintptr {
	return uintptr(*p)
}

func (p *UIntPtr) Slice(n int) []UIntPtr {
	return ((*[1 << 28]UIntPtr)(unsafe.Pointer(p)))[0:n:n]
}

func (p *UIntPtr) Free() {
	C.free(unsafe.Pointer(p))
}

// -----------------------------------------------------------------------------
