// Copyright 2016 <chaishushan{AT}gmail.com>. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// export cgo functions for go test

package cgo

/*
#cgo CXXFLAGS: -std=c++11
#cgo windows LDFLAGS: -Wl,--allow-multiple-definition

*/
import "C"
