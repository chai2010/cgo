// Copyright 2016 <chaishushan{AT}gmail.com>. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package cgo

import (
	"testing"
)

func TestString(t *testing.T) {
	s := NewString("")
	defer s.Delete()

	tAssertStrEQ(t, s, "")

	s.AssignByGoString("abc")
	tAssertStrEQ(t, s, "abc")

	s.AssignByGoString("")
	tAssertStrEQ(t, s, "")

	s.AssignByGoBytes([]byte("123"))
	tAssertStrEQ(t, s, "123")

	s.AssignByGoBytes(nil)
	tAssertStrEQ(t, s, "")
}

func tAssertStrEQ(t *testing.T, s1 *String, s2 string) {
	if s1.Size() != len(s2) {
		t.Fatalf("%q: bad size", s2)
	}

	if s1.Data().Strlen() != len(s2) {
		t.Fatalf("%q: bad Strlen", s2)
	}
	if s1.CStr().Strlen() != len(s2) {
		t.Fatalf("%q: bad Strlen", s2)
	}

	if string(s1.GoSlice()) != s2 {
		t.Fatalf("%q: bad GoSlice", s2)
	}
	if s1.GoString() != s2 {
		t.Fatalf("%q: bad GoString", s2)
	}
}
