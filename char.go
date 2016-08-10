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

// Go string to C string
// The C string is allocated in the C heap using malloc.
// It is the caller's responsibility to arrange for it to be
// freed, such as by calling C.free (be sure to include stdlib.h
// if C.free is needed).
func NewCharString(s string) *Char {
	return (*Char)(C.CString(s))
}

func NewCharFormat(format string, args ...interface{}) *Char {
	s := fmt.Sprintf(format, args...)
	return (*Char)(C.CString(s))
}

func (s *Char) IsEmpty() bool {
	return s == nil || *s == 0
}

// C string to Go string
func (s *Char) GoString() string {
	return C.GoString((*C.char)(s))
}

// C data with explicit length to Go string
func (s *Char) GoStringN(n int) string {
	return C.GoStringN((*C.char)(s), C.int(n))
}

func (s *Char) GoBytes(n int) []byte {
	return C.GoBytes(unsafe.Pointer(s), C.int(n))
}

func (s *Char) Strlen() int {
	return int(C.strlen((*C.char)(s)))
}

func (s *Char) Strdup() *Char {
	d := (*Char)(C.malloc(C.strlen((*C.char)(s)) + 1))
	if d == nil {
		return nil
	}
	C.strcpy((*C.char)(d), (*C.char)(s))
	return d
}

func (dst *Char) Strcpy(src *Char) *Char {
	return (*Char)(C.strcpy((*C.char)(dst), (*C.char)(src)))
}

func (dst *Char) Strncpy(src *Char, num int) *Char {
	return (*Char)(C.strncpy((*C.char)(dst), (*C.char)(src), C.size_t(num)))
}

func (dst *Char) Strcat(src *Char) *Char {
	return (*Char)(C.strcat((*C.char)(dst), (*C.char)(src)))
}

func (dst *Char) Strncat(src *Char, num int) *Char {
	return (*Char)(C.strncat((*C.char)(dst), (*C.char)(src), C.size_t(num)))
}

func (s *Char) Strchr(ch int) *Char {
	return (*Char)(C.strchr((*C.char)(s), C.int(ch)))
}

func (s *Char) Strrchr(ch int) *Char {
	return (*Char)(C.strrchr((*C.char)(s), (C.int)(ch)))
}

func (s *Char) Strstr(s2 *Char) *Char {
	return (*Char)(C.strstr((*C.char)(s), (*C.char)(s2)))
}

func (s *Char) Strcspn(s2 *Char) int {
	return int(C.strcspn((*C.char)(s), (*C.char)(s2)))
}

func (s *Char) Strspn(s2 *Char) int {
	return (int)(C.strspn((*C.char)(s), (*C.char)(s2)))
}

func (s *Char) Strpbrk(s2 *Char) *Char {
	return (*Char)(C.strpbrk((*C.char)(s), (*C.char)(s2)))
}

func (s *Char) Strtok(delimiters *Char) *Char {
	return (*Char)(C.strtok((*C.char)(s), (*C.char)(delimiters)))
}
