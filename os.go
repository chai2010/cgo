// Copyright 2016 <chaishushan{AT}gmail.com>. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package cgo

/*
#include <stdio.h>
*/
import "C"

func Remove(filename *CChar) int {
	return int(C.remove((*C.char)(filename)))
}

func Rename(oldname, newname *CChar) int {
	return int(C.rename((*C.char)(oldname), (*C.char)(newname)))
}

func Tmpfile() *CFile {
	return (*CFile)(C.tmpfile())
}
