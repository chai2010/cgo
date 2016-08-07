// Copyright 2016 <chaishushan{AT}gmail.com>. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package cgo

//#include <string.h>
import "C"

import (
	"unsafe"
)

func Memcpy(dst, src unsafe.Pointer, num int) unsafe.Pointer {
	return C.memcpy(dst, src, C.size_t(num))
}

func Memset(p unsafe.Pointer, value, num int) unsafe.Pointer {
	return C.memset(p, C.int(value), C.size_t(num))
}

func Memmove(dst, src unsafe.Pointer, num int) unsafe.Pointer {
	return C.memmove(dst, src, C.size_t(num))
}

func Memchr(p unsafe.Pointer, value, num int) unsafe.Pointer {
	return C.memchr(p, C.int(value), C.size_t(num))
}

func Strlen(s *CChar) int {
	return int(C.strlen((*C.char)(s)))
}

func Strcpy(dst, src *CChar) *CChar {
	return (*CChar)(C.strcpy((*C.char)(dst), (*C.char)(src)))
}

func Strncpy(dst, src *CChar, num int) *CChar {
	return (*CChar)(C.strncpy((*C.char)(dst), (*C.char)(src), C.size_t(num)))
}

func Strcat(dst, src *CChar) *CChar {
	return (*CChar)(C.strcat((*C.char)(dst), (*C.char)(src)))
}

func Strncat(dst, src *CChar, num int) *CChar {
	return (*CChar)(C.strncat((*C.char)(dst), (*C.char)(src), C.size_t(num)))
}

func Strchr(p *CChar, ch int) *CChar {
	return (*CChar)(C.strchr((*C.char)(p), C.int(ch)))
}

func Strstr(s1, s2 *CChar) *CChar {
	return (*CChar)(C.strstr((*C.char)(s1), (*C.char)(s2)))
}

func Strrchr(s1 *CChar, ch int) *CChar {
	return (*CChar)(C.strrchr((*C.char)(s1), (C.int)(ch)))
}

func Strcspn(s1, s2 *CChar) int {
	return int(C.strcspn((*C.char)(s1), (*C.char)(s2)))
}

func Strspn(s1, s2 *CChar) int {
	return (int)(C.strspn((*C.char)(s1), (*C.char)(s2)))
}

func Strpbrk(s1, s2 *CChar) *CChar {
	return (*CChar)(C.strpbrk((*C.char)(s1), (*C.char)(s2)))
}

func Strtok(s, delimiters *CChar) *CChar {
	return (*CChar)(C.strtok((*C.char)(s), (*C.char)(delimiters)))
}

func Strerr(errnum int) *CChar {
	return (*CChar)(C.strerror(C.int(errnum)))
}
