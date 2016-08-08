// Copyright 2016 <chaishushan{AT}gmail.com>. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package cgo

/*
#include <errno.h>
#include <string.h>

static int cgo_get_errno() {
	return errno;
}
static void cgo_set_errno(int value) {
	errno = value;
}
*/
import "C"

func Errno() int {
	return int(C.cgo_get_errno())
}

func ErrnoSet(v int) {
	C.cgo_set_errno(C.int(v))
}

func ErrnoDesc(errnum int) *Char {
	return (*Char)(C.strerror(C.int(errnum)))
}
