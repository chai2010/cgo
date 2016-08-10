// Copyright 2016 <chaishushan{AT}gmail.com>. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// https://github.com/chai2010/cgo

#pragma once

#ifndef CHAI2010_CGO_C_ERROR_H_
#define CHAI2010_CGO_C_ERROR_H_

#include <stdarg.h>
#include <stddef.h>
#include <stdint.h>
#include <stdlib.h>

#if defined(__cplusplus)
extern "C" {
#endif

typedef struct chai2010_cgo_Error_T chai2010_cgo_Error_T;

chai2010_cgo_Error_T* chai2010_cgo_Error_New(int errCode, const char* errText);
void chai2010_cgo_Error_Delete(chai2010_cgo_Error_T* p);

void chai2010_cgo_Error_SetError(chai2010_cgo_Error_T* p, int errCode, const char* errText);

int chai2010_cgo_Error_GetCode(chai2010_cgo_Error_T* p);
const char* chai2010_cgo_Error_GetText(chai2010_cgo_Error_T* p);

#if defined(__cplusplus)
}
#endif
#endif // CHAI2010_CGO_C_ERROR_H_
