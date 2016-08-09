// Copyright 2016 <chaishushan{AT}gmail.com>. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package cgo

//#include "./c_string.h"
import "C"

import (
	"reflect"
	"unsafe"
)

type StringList C.chai2010_cgo_StringList_T

func NewStringList(args ...string) *StringList {
	p := (*StringList)(C.chai2010_cgo_StringList_New())
	p.Resize(len(args))
	for i, s := range args {
		p.AssignByGoString(i, s)
	}
	return p
}
func (p *StringList) Delete() {
	C.chai2010_cgo_StringList_Delete((*C.chai2010_cgo_StringList_T)(p))
}

func (p *StringList) Size() int {
	return int(C.chai2010_cgo_StringList_Size((*C.chai2010_cgo_StringList_T)(p)))
}
func (p *StringList) Resize(newSize int) {
	C.chai2010_cgo_StringList_Resize((*C.chai2010_cgo_StringList_T)(p), C.int(newSize))
}
func (p *StringList) Clear() {
	C.chai2010_cgo_StringList_Clear((*C.chai2010_cgo_StringList_T)(p))
}

func (p *StringList) StrSize(i int) int {
	return int(C.chai2010_cgo_StringList_StrSize((*C.chai2010_cgo_StringList_T)(p), C.int(i)))
}
func (p *StringList) StrResize(i, newSize int) {
	C.chai2010_cgo_StringList_StrResize((*C.chai2010_cgo_StringList_T)(p), C.int(i), C.int(newSize))
}
func (p *StringList) StrClear(i int) {
	C.chai2010_cgo_StringList_StrClear((*C.chai2010_cgo_StringList_T)(p), C.int(i))
}

func (p *StringList) StrData(i int) *Char {
	return (*Char)(C.chai2010_cgo_StringList_StrData((*C.chai2010_cgo_StringList_T)(p), C.int(i)))
}
func (p *StringList) StrCStr(i int) *Char {
	return (*Char)(C.chai2010_cgo_StringList_StrCStr((*C.chai2010_cgo_StringList_T)(p), C.int(i)))
}

func (p *StringList) StrGoString(i int) string {
	return C.GoString((*C.char)(p.StrData(i)))
}
func (p *StringList) StrGoData(i int) []byte {
	if p, n := p.StrData(i), p.StrSize(i); n > 0 {
		return (*[1 << 31]byte)(unsafe.Pointer(p))[0:n:n]
	}
	return nil
}

func (p *StringList) AssignStr(i int, s *Char) {
	C.chai2010_cgo_StringList_AssignStr((*C.chai2010_cgo_StringList_T)(p), C.int(i), (*C.char)(s))
}
func (p *StringList) AssignStrN(i int, s *Char, n int) {
	C.chai2010_cgo_StringList_AssignStrN((*C.chai2010_cgo_StringList_T)(p), C.int(i), (*C.char)(s), C.int(n))
}

func (p *StringList) AssignByGoString(i int, s string) {
	C.chai2010_cgo_StringList_AssignStrN((*C.chai2010_cgo_StringList_T)(p), C.int(i),
		(*C.char)(unsafe.Pointer((*reflect.StringHeader)(unsafe.Pointer(&s)).Data)),
		C.int(len(s)),
	)
}
func (p *StringList) AssignByGoSlice(i int, s []byte) {
	C.chai2010_cgo_StringList_AssignStrN((*C.chai2010_cgo_StringList_T)(p), C.int(i),
		(*C.char)(unsafe.Pointer((*reflect.SliceHeader)(unsafe.Pointer(&s)).Data)),
		C.int(len(s)),
	)
}

func (p *StringList) AppendStr(str *Char) {
	C.chai2010_cgo_StringList_AppendStr((*C.chai2010_cgo_StringList_T)(p), (*C.char)(str))
}
func (p *StringList) AppendStrN(str *Char, n int) {
	C.chai2010_cgo_StringList_AppendStrN((*C.chai2010_cgo_StringList_T)(p), (*C.char)(str), C.int(n))
}

func (p *StringList) AppendByGoString(s string) {
	C.chai2010_cgo_StringList_AppendStrN((*C.chai2010_cgo_StringList_T)(p),
		(*C.char)(unsafe.Pointer((*reflect.StringHeader)(unsafe.Pointer(&s)).Data)),
		C.int(len(s)),
	)
}
func (p *StringList) AppendByGoSlice(s []byte) {
	C.chai2010_cgo_StringList_AppendStrN((*C.chai2010_cgo_StringList_T)(p),
		(*C.char)(unsafe.Pointer((*reflect.SliceHeader)(unsafe.Pointer(&s)).Data)),
		C.int(len(s)),
	)
}
