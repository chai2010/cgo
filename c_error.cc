// Copyright 2016 <chaishushan{AT}gmail.com>. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

#include "c_error.h"

#include <string>

struct chai2010_cgo_Error_T {
	int         errCode;
	std::string errText;
};

chai2010_cgo_Error_T* chai2010_cgo_Error_New(int errCode, const char* errText) {
	auto p = new chai2010_cgo_Error_T();
	p->errCode = errCode;
	p->errText = errText? errText: "";
	return p;
}
void chai2010_cgo_Error_Delete(chai2010_cgo_Error_T* p) {
	delete p;
}

void chai2010_cgo_Error_SetError(chai2010_cgo_Error_T* p, int errCode, const char* errText) {
	p->errCode = errCode;
	p->errText = errText? errText: "";
}

int chai2010_cgo_Error_GetCode(chai2010_cgo_Error_T* p) {
	return p->errCode;
}
const char* chai2010_cgo_Error_GetText(chai2010_cgo_Error_T* p) {
	return p->errText.c_str();
}
