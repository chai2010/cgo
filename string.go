// Copyright 2016 <chaishushan{AT}gmail.com>. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package cgo

//#include "./c_string.h"
import "C"

import (
	"fmt"
	"reflect"
	"unsafe"
)

type String C.chai2010_cgo_String_T

func NewString(s string) *String {
	p := (*String)(C.chai2010_cgo_String_New())
	p.AssignByGoString(s)
	return p
}
func NewStringFormat(format string, args ...interface{}) *String {
	return NewString(fmt.Sprintf(format, args...))
}
func NewStringFromStr(s *Char) *String {
	return (*String)(C.chai2010_cgo_String_NewFromStr((*C.char)(s)))
}
func NewStringFromStrN(s *Char, n int) *String {
	return (*String)(C.chai2010_cgo_String_NewFromStrN((*C.char)(s), C.int(n)))
}
func (s *String) Delete() {
	C.chai2010_cgo_String_Delete((*C.chai2010_cgo_String_T)(s))
}

func (s *String) Size() int {
	return int(C.chai2010_cgo_String_Size((*C.chai2010_cgo_String_T)(s)))
}
func (s *String) Resize(newSize int) {
	C.chai2010_cgo_String_Resize((*C.chai2010_cgo_String_T)(s), C.int(newSize))
}
func (s *String) Clear() {
	C.chai2010_cgo_String_Clear((*C.chai2010_cgo_String_T)(s))
}

func (s *String) Data() *Char {
	return (*Char)(C.chai2010_cgo_String_Data((*C.chai2010_cgo_String_T)(s)))
}
func (s *String) CStr() *Char {
	return (*Char)(C.chai2010_cgo_String_CStr((*C.chai2010_cgo_String_T)(s)))
}

func (s *String) AssignStr(str *Char) {
	C.chai2010_cgo_String_AssignStr((*C.chai2010_cgo_String_T)(s), (*C.char)(str))
}
func (s *String) AssignStrN(str *Char, n int) {
	C.chai2010_cgo_String_AssignStrN((*C.chai2010_cgo_String_T)(s), (*C.char)(str), C.int(n))
}

func (s *String) AssignByGoString(str string) {
	C.chai2010_cgo_String_AssignStrN((*C.chai2010_cgo_String_T)(s),
		(*C.char)(unsafe.Pointer((*reflect.StringHeader)(unsafe.Pointer(&str)).Data)),
		C.int(len(str)),
	)
}
func (s *String) AssignByGoBytes(str []byte) {
	C.chai2010_cgo_String_AssignStrN((*C.chai2010_cgo_String_T)(s),
		(*C.char)(unsafe.Pointer((*reflect.SliceHeader)(unsafe.Pointer(&str)).Data)),
		C.int(len(str)),
	)
}

func (s *String) GoString() string {
	return C.GoString((*C.char)(s.Data()))
}

func (s *String) GoSlice() []byte {
	if p, n := s.Data(), s.Size(); n > 0 {
		return (*[1 << 31]byte)(unsafe.Pointer(p))[0:n:n]
	}
	return nil
}
