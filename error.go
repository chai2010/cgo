// Copyright 2016 <chaishushan{AT}gmail.com>. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package cgo

//#include "./c_error.h"
import "C"

type Error C.chai2010_cgo_Error_T

func NewError(code int, desc *Char) *Error {
	return (*Error)(C.chai2010_cgo_Error_New(C.int(code), (*C.char)(desc)))
}
func NewErrorFrom(code int, desc string) *Error {
	cs := NewCharString(desc)
	defer cs.Free()

	return (*Error)(C.chai2010_cgo_Error_New(C.int(code), (*C.char)(cs)))
}
func (p *Error) Delete() {
	C.chai2010_cgo_Error_Delete((*C.chai2010_cgo_Error_T)(p))
}

func (p *Error) SetError(code int, desc *Char) {
	C.chai2010_cgo_Error_SetError((*C.chai2010_cgo_Error_T)(p), C.int(code), (*C.char)(desc))
}
func (p *Error) SetErrorFrom(code int, desc string) {
	cs := NewCharString(desc)
	defer cs.Free()

	C.chai2010_cgo_Error_SetError((*C.chai2010_cgo_Error_T)(p), C.int(code), (*C.char)(cs))
}

func (p *Error) GetCode() int {
	return int(C.chai2010_cgo_Error_GetCode((*C.chai2010_cgo_Error_T)(p)))
}

func (p *Error) GetText() *Char {
	return (*Char)(C.chai2010_cgo_Error_GetText((*C.chai2010_cgo_Error_T)(p)))
}
