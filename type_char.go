// Copyright 2016 <chaishushan{AT}gmail.com>. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package cgo

//#include <stdlib.h>
//#include <string.h>
import "C"

import (
	"fmt"
	"unsafe"
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

func NewCharFormat(format string, args ...interface{}) *CChar {
	s := fmt.Sprintf(format, args...)
	return CString(s)
}

func (s *CChar) Slice(n int) []byte {
	return ((*[1 << 31]byte)(unsafe.Pointer(s)))[0:n:n]
}
func (p *CChar) Free() {
	C.free(unsafe.Pointer(p))
}

func (s *CChar) IsEmpty() bool {
	return s == nil || *s == 0
}

func (s *CChar) Len() int {
	return int(C.strlen((*C.char)(s)))
}

func (s *CChar) Dup() *CChar {
	d := (*CChar)(C.malloc(C.strlen((*C.char)(s)) + 1))
	if d == nil {
		return nil
	}
	C.strcpy((*C.char)(d), (*C.char)(s))
	return d
}

// C string to Go string
func (s *CChar) GoString() string {
	return C.GoString((*C.char)(s))
}

// C data with explicit length to Go string
func (s *CChar) GoStringN(n int) string {
	return C.GoStringN((*C.char)(s), C.int(n))
}
