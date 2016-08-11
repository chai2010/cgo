// Copyright 2016 <chaishushan{AT}gmail.com>. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package cgo

/*
#include <stdio.h>
#include <stdlib.h>
#include <string.h>
#include <errno.h>

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
	"errors"
	"io"
	"unsafe"
)

type File C.FILE

var (
	_ io.Seeker = (*File)(nil)
	_ io.Reader = (*File)(nil)
	_ io.Writer = (*File)(nil)
	_ io.Closer = (*File)(nil)

	_ io.ByteReader = (*File)(nil)
	_ io.ByteWriter = (*File)(nil)
)

const (
	EOF = C.EOF

	SEEK_CUR = C.SEEK_CUR
	SEEK_SET = C.SEEK_SET
	SEEK_END = C.SEEK_END
)

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

func (f *File) Close() error {
	n := int(C.fclose((*C.FILE)(f)))
	if n != 0 {
		return io.EOF
	}
	return nil
}

func (f *File) Flush() {
	C.fflush((*C.FILE)(f))
}

func (f *File) Read(buf []byte) (int, error) {
	n := int(C.fread(unsafe.Pointer(&buf[0]), C.size_t(1), C.size_t(len(buf)), (*C.FILE)(f)))
	if n == len(buf) {
		return n, nil
	}
	if f.Eof() != 0 {
		return 0, io.EOF
	}
	err := errors.New(C.GoString(C.strerror(C.int(f.Error()))))
	return 0, err
}

func (f *File) Write(buf []byte) (int, error) {
	n := int(C.fwrite(unsafe.Pointer(&buf[0]), C.size_t(1), C.size_t(len(buf)), (*C.FILE)(f)))
	if n == len(buf) {
		return n, nil
	}
	if f.Eof() != 0 {
		return 0, io.EOF
	}
	err := errors.New(C.GoString(C.strerror(C.int(f.Error()))))
	return 0, err
}

func (f *File) ReadByte() (c byte, err error) {
	if ch := f.Getc(); ch >= 0 {
		return byte(ch), nil
	}
	if f.Eof() != 0 {
		return 0, io.EOF
	}
	err = errors.New(C.GoString(C.strerror(C.int(f.Error()))))
	return 0, err
}
func (f *File) WriteByte(c byte) error {
	if n := f.Putc(int(c)); n != int(C.EOF) {
		return nil
	}
	if f.Eof() != 0 {
		return io.EOF
	}
	err := errors.New(C.GoString(C.strerror(C.int(f.Error()))))
	return err
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

func (f *File) Tell() int64 {
	return int64(C.ftell((*C.FILE)(f)))
}
func (f *File) Seek(off int64, origin int) (int64, error) {
	n := int(C.fseek((*C.FILE)(f), C.long(off), C.int(origin)))
	if n == 0 {
		return f.Tell(), nil
	}
	err := errors.New(C.GoString(C.strerror(C.int(f.Error()))))
	return 0, err
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
