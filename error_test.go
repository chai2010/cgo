// Copyright 2016 <chaishushan{AT}gmail.com>. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package cgo

import (
	"testing"
)

func TestError(t *testing.T) {
	p := NewErrorFrom(9527, "hi")
	defer p.Delete()

	code := p.GetCode()
	text := p.GetText().GoString()

	if code != 9527 {
		t.Fatalf("expect = %v, got = %v", 9527, code)
	}
	if text != "hi" {
		t.Fatalf("expect = %v, got = %v", "hi", text)
	}
}

func TestError_empty(t *testing.T) {
	p := NewErrorFrom(0, "")
	defer p.Delete()

	code := p.GetCode()
	text := p.GetText().GoString()

	if code != 0 {
		t.Fatalf("expect = %v, got = %v", 0, code)
	}
	if text != "" {
		t.Fatalf("expect = <empty>, got = %v", text)
	}
}

func TestError_init(t *testing.T) {
	p := NewErrorFrom(0, "")
	defer p.Delete()

	p.SetErrorFrom(9527, "hi")

	code := p.GetCode()
	text := p.GetText().GoString()

	if code != 9527 {
		t.Fatalf("expect = %v, got = %v", 9527, code)
	}
	if text != "hi" {
		t.Fatalf("expect = %v, got = %v", "hi", text)
	}
}
