// Copyright 2016 <chaishushan{AT}gmail.com>. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// https://github.com/chai2010/cgo

#pragma once

#ifndef CHAI2010_CGO_C_STRING_H_
#define CHAI2010_CGO_C_STRING_H_

#include <stdarg.h>
#include <stddef.h>
#include <stdint.h>
#include <stdlib.h>

#if defined(__cplusplus)
extern "C" {
#endif

// ----------------------------------------------------------------------------
// types
// ----------------------------------------------------------------------------

typedef struct chai2010_cgo_String_T     chai2010_cgo_String_T;
typedef struct chai2010_cgo_StringList_T chai2010_cgo_StringList_T;

// ----------------------------------------------------------------------------
// string
// ----------------------------------------------------------------------------

chai2010_cgo_String_T* chai2010_cgo_String_New();
chai2010_cgo_String_T* chai2010_cgo_String_NewFromStr(const char* str);
chai2010_cgo_String_T* chai2010_cgo_String_NewFromStrN(const char* str, int n);
void chai2010_cgo_String_Delete(chai2010_cgo_String_T* p);

int  chai2010_cgo_String_Size(chai2010_cgo_String_T* p);
void chai2010_cgo_String_Resize(chai2010_cgo_String_T* p, int newSize);
void chai2010_cgo_String_Clear(chai2010_cgo_String_T* p);

const char* chai2010_cgo_String_CStr(chai2010_cgo_String_T* p);
const char* chai2010_cgo_String_Data(chai2010_cgo_String_T* p);

void chai2010_cgo_String_AssignStr(chai2010_cgo_String_T* p, const char* str);
void chai2010_cgo_String_AssignStrN(chai2010_cgo_String_T* p, const char* str, int size);

// ----------------------------------------------------------------------------
// string list
// ----------------------------------------------------------------------------

chai2010_cgo_StringList_T* chai2010_cgo_StringList_New();
void chai2010_cgo_StringList_Delete(chai2010_cgo_StringList_T* p);

int  chai2010_cgo_StringList_Size(chai2010_cgo_StringList_T* p);
void chai2010_cgo_StringList_Resize(chai2010_cgo_StringList_T* p, int newSize);
void chai2010_cgo_StringList_Clear(chai2010_cgo_StringList_T* p);

int  chai2010_cgo_StringList_StrSize(chai2010_cgo_StringList_T* p, int i);
void chai2010_cgo_StringList_StrResize(chai2010_cgo_StringList_T* p, int i, int newSize);
void chai2010_cgo_StringList_StrClear(chai2010_cgo_StringList_T* p, int i);

const char* chai2010_cgo_StringList_StrData(chai2010_cgo_StringList_T* p, int i);
const char* chai2010_cgo_StringList_StrCStr(chai2010_cgo_StringList_T* p, int i);

void chai2010_cgo_StringList_AssignStr(chai2010_cgo_StringList_T* p, int i, const char* str);
void chai2010_cgo_StringList_AssignStrN(chai2010_cgo_StringList_T* p, int i, const char* str, int size);

void chai2010_cgo_StringList_AppendStr(chai2010_cgo_StringList_T* p, const char* str);
void chai2010_cgo_StringList_AppendStrN(chai2010_cgo_StringList_T* p, const char* str, int size);

// ----------------------------------------------------------------------------
// END
// ----------------------------------------------------------------------------

#if defined(__cplusplus)
}
#endif
#endif // CHAI2010_CGO_C_STRING_H_
