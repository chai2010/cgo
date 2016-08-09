// Copyright 2016 <chaishushan{AT}gmail.com>. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

#include "./c_string.h"

#include <string>
#include <vector>

// ----------------------------------------------------------------------------
// types
// ----------------------------------------------------------------------------

struct chai2010_cgo_String_T {
	std::string v;
};

struct chai2010_cgo_StringList_T {
	std::vector<std::string> v;
};

// ----------------------------------------------------------------------------
// string
// ----------------------------------------------------------------------------

chai2010_cgo_String_T* chai2010_cgo_String_New() {
	return new chai2010_cgo_String_T();
}
chai2010_cgo_String_T* chai2010_cgo_String_NewFromStr(const char* str) {
	auto p = new chai2010_cgo_String_T();
	p->v.assign(str);
	return p;
}
chai2010_cgo_String_T* chai2010_cgo_String_NewFromStrN(const char* str, int n) {
	auto p = new chai2010_cgo_String_T();
	p->v.assign(str, size_t(n));
	return p;
}
void chai2010_cgo_String_Delete(chai2010_cgo_String_T* p) {
	delete p;
}

int chai2010_cgo_String_Size(chai2010_cgo_String_T* p) {
	return p->v.size();
}
void chai2010_cgo_String_Resize(chai2010_cgo_String_T* p, int newSize) {
	p->v.resize(newSize);
}
void chai2010_cgo_String_Clear(chai2010_cgo_String_T* p) {
	p->v.clear();
}

const char* chai2010_cgo_String_CStr(chai2010_cgo_String_T* p) {
	return p->v.data();
}
const char* chai2010_cgo_String_Data(chai2010_cgo_String_T* p) {
	return p->v.c_str();
}

void chai2010_cgo_String_AssignStr(chai2010_cgo_String_T* p, const char* str) {
	p->v.assign(str);
}
void chai2010_cgo_String_AssignStrN(chai2010_cgo_String_T* p, const char* str, int size) {
	p->v.assign(str, size_t(size));
}

// ----------------------------------------------------------------------------
// string list
// ----------------------------------------------------------------------------

chai2010_cgo_StringList_T* chai2010_cgo_StringList_New() {
	return new chai2010_cgo_StringList_T();
}
void chai2010_cgo_StringList_Delete(chai2010_cgo_StringList_T* p) {
	delete p;
}

int chai2010_cgo_StringList_Size(chai2010_cgo_StringList_T* p) {
	return p->v.size();
}
void chai2010_cgo_StringList_Resize(chai2010_cgo_StringList_T* p, int newSize) {
	p->v.resize(newSize);
}
void chai2010_cgo_StringList_Clear(chai2010_cgo_StringList_T* p) {
	p->v.clear();
}

int chai2010_cgo_StringList_StrSize(chai2010_cgo_StringList_T* p, int i) {
	return p->v[i].size();
}
void chai2010_cgo_StringList_StrResize(chai2010_cgo_StringList_T* p, int i, int newSize) {
	p->v[i].resize(newSize);
}
void chai2010_cgo_StringList_StrClear(chai2010_cgo_StringList_T* p, int i) {
	p->v[i].clear();
}

const char* chai2010_cgo_StringList_StrData(chai2010_cgo_StringList_T* p, int i) {
	return p->v[i].data();
}
const char* chai2010_cgo_StringList_StrCStr(chai2010_cgo_StringList_T* p, int i) {
	return p->v[i].c_str();
}

void chai2010_cgo_StringList_AssignStr(chai2010_cgo_StringList_T* p, int i, const char* str) {
	p->v[i].assign(str);
}
void chai2010_cgo_StringList_AssignStrN(chai2010_cgo_StringList_T* p, int i, const char* str, int size) {
	p->v[i].assign(str, size_t(size));
}

void chai2010_cgo_StringList_AppendStr(chai2010_cgo_StringList_T* p, const char* str) {
	p->v.push_back(std::string(str));
}
void chai2010_cgo_StringList_AppendStrN(chai2010_cgo_StringList_T* p, const char* str, int size) {
	p->v.push_back(std::string(str, size_t(size)));
}
