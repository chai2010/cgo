// Copyright 2016 <chaishushan{AT}gmail.com>. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package cgo

//#include <stdlib.h>
//#include <string.h>
import "C"

import (
	"fmt"
)

func Sprintf(format string, args ...interface{}) *CChar {
	s := fmt.Sprintf(format, args...)
	return CString(s)
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

func (s *CChar) GoString() string {
	return C.GoString((*C.char)(s))
}
