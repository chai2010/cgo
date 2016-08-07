// Copyright 2016 <chaishushan{AT}gmail.com>. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package cgo

//#include <stdio.h>
import "C"

type CFILE C.FILE

func CPuts(s *CChar) int {
	return int(C.puts((*C.char)(s)))
}

// TODO: support <stdio.h>
