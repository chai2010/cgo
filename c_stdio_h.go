// Copyright 2016 <chaishushan{AT}gmail.com>. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package cgo

//#include <stdio.h>
import "C"

type CFile C.FILE

func Remove(filename *CChar) int {
	return int(C.remove((*C.char)(filename)))
}

func Rename(oldname, newname *CChar) int {
	return int(C.rename((*C.char)(oldname), (*C.char)(newname)))
}

func Tmpfile() *CFile {
	return (*CFile)(C.tmpfile())
}

func Fopen(filename, mode *CChar) *CFile {
	return (*CFile)(C.fopen((*C.char)(filename), (*C.char)(mode)))
}

func Fclose(p *CFile) int {
	return int(C.fclose((*C.FILE)(p)))
}

func Fflush(p *CFile) int {
	return int(C.fflush((*C.FILE)(p)))
}

func CPuts(s *CChar) int {
	return int(C.puts((*C.char)(s)))
}

// TODO: support <stdio.h>

// TODO: rename package cgo to crt
