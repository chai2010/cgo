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

type CFile C.FILE

var (
	Stdin  = (*CFile)(C.cgo_get_stdin())
	Stdout = (*CFile)(C.cgo_get_stdout())
	Stderr = (*CFile)(C.cgo_get_stderr())
)

func Fopen(filename, mode *CChar) *CFile {
	return (*CFile)(C.fopen((*C.char)(filename), (*C.char)(mode)))
}

func (f *CFile) Close() int {
	return int(C.fclose((*C.FILE)(f)))
}

func (f *CFile) Puts(s *CChar) {
	C.fputs((*C.char)(s), (*C.FILE)(f))
}

func (f *CFile) Flush() {
	C.fflush((*C.FILE)(f))
}

func Remove(filename *CChar) int {
	return int(C.remove((*C.char)(filename)))
}

func Rename(oldname, newname *CChar) int {
	return int(C.rename((*C.char)(oldname), (*C.char)(newname)))
}

func Tmpfile() *CFile {
	return (*CFile)(C.tmpfile())
}

func Fclose(p *CFile) int {
	return int(C.fclose((*C.FILE)(p)))
}

func Fflush(p *CFile) int {
	return int(C.fflush((*C.FILE)(p)))
}

func Puts(s *CChar) int {
	return int(C.puts((*C.char)(s)))
}
