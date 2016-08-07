// Copyright 2016 <chaishushan{AT}gmail.com>. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package cgo

import (
	"sync"
)

type GoObjectId int32

var refs struct {
	sync.Mutex
	objs map[GoObjectId]interface{}
	next GoObjectId
}

func init() {
	refs.Lock()
	defer refs.Unlock()

	refs.objs = make(map[GoObjectId]interface{})
	refs.next = 1000
}

func NewGoObjectId(obj interface{}) GoObjectId {
	refs.Lock()
	defer refs.Unlock()

	id := refs.next
	refs.next++

	refs.objs[id] = obj
	return id
}

func (id GoObjectId) IsNil() bool {
	return id == 0
}

func (id GoObjectId) Get() interface{} {
	refs.Lock()
	defer refs.Unlock()

	return refs.objs[id]
}

func (id GoObjectId) Free() interface{} {
	refs.Lock()
	defer refs.Unlock()

	obj := refs.objs[id]
	delete(refs.objs, id)
	return obj
}
