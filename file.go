// Copyright 2016 <chaishushan{AT}gmail.com>. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package cgo

/*
#include <stdio.h>

// on mingw, stderr and stdout are defined as &_iob[FILENO]
// on netbsd, they are defined as &__sF[FILENO]
// and cgo doesn't recognize them, so write a function to get them,
// instead of depending on internals of libc implementation.

static FILE *cgo_get_stdin(void)  { return stdin;  }
static FILE *cgo_get_stdout(void) { return stdout; }
static FILE *cgo_get_stderr(void) { return stderr; }

*/
import "C"

type File C.FILE

var (
	Stdin  = (*File)(C.cgo_get_stdin())
	Stdout = (*File)(C.cgo_get_stdout())
	Stderr = (*File)(C.cgo_get_stderr())
)

func OpenFile(filename, mode *Char) *File {
	return (*File)(C.fopen((*C.char)(filename), (*C.char)(mode)))
}

func TempFile() *File {
	return (*File)(C.tmpfile())
}

func (f *File) Close() int {
	return int(C.fclose((*C.FILE)(f)))
}

func (f *File) Puts(s *Char) {
	C.fputs((*C.char)(s), (*C.FILE)(f))
}

func (f *File) Flush() {
	C.fflush((*C.FILE)(f))
}
