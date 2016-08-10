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

import (
	"unsafe"
)

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

func (f *File) Flush() {
	C.fflush((*C.FILE)(f))
}

func (f *File) Read(buf []byte) int {
	return int(C.fread(unsafe.Pointer(&buf[0]), C.size_t(1), C.size_t(len(buf)), (*C.FILE)(f)))
}

func (f *File) Write(buf []byte) int {
	return int(C.fwrite(unsafe.Pointer(&buf[0]), C.size_t(1), C.size_t(len(buf)), (*C.FILE)(f)))
}

func (f *File) Puts(s *Char) {
	C.fputs((*C.char)(s), (*C.FILE)(f))
}

func (f *File) Getc() int {
	return int(C.fgetc((*C.FILE)(f)))
}

func (f *File) Gets(buf []byte) *Char {
	return (*Char)(C.fgets((*C.char)(unsafe.Pointer(&buf[0])), C.int(len(buf)), (*C.FILE)(f)))
}

func (f *File) Putc(ch int) int {
	return int(C.fputc(C.int(ch), (*C.FILE)(f)))
}

func (f *File) Ungetc(ch int) int {
	return int(C.ungetc(C.int(ch), (*C.FILE)(f)))
}

func (f *File) Tell() int {
	return int(C.ftell((*C.FILE)(f)))
}
func (f *File) Seek(off, origin int) int {
	return int(C.fseek((*C.FILE)(f), C.long(off), C.int(origin)))
}
func (f *File) Rewind() {
	C.rewind((*C.FILE)(f))
}

func (f *File) Eof() int {
	return int(C.feof((*C.FILE)(f)))
}
func (f *File) Error() int {
	return int(C.ferror((*C.FILE)(f)))
}
func (f *File) ClearErr() {
	C.clearerr((*C.FILE)(f))
}
