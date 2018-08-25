// Copyright 2018 ConsenSys AG
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package app

import "github.com/gopherjs/gopherjs/js"

// JSArray is a wrapper around a JS Array with a convenient UpdateEntry method
type JSArray struct {
	*js.Object
	headers    *js.Object `js:"headers"`
	rows       *js.Object `js:"rows"`
	indices    *js.Object `js:"indices"`
	nbElements int        `js:"nbElements"`
}

func NewJSArray(headers ...interface{}) *JSArray {
	toReturn := &JSArray{Object: newJSObject()}
	toReturn.headers = js.Global.Get("Array").New(headers...)
	toReturn.rows = js.Global.Get("Object").New(js.Global.Get("Array").New())
	toReturn.indices = newJSObject()
	toReturn.nbElements = 0
	return toReturn
}

// UpdateEntry update or insert row with key key.
// if key doesn't exist, add the new row on top of the array (shift the rest)
func (array *JSArray) UpdateEntry(key string, cells ...JSCell) {
	// convert to interface{}
	var jsCells []interface{}
	for _, c := range cells {
		jsCells = append(jsCells, c.toJSObject())
	}

	// check if record already exists
	if id := array.indices.Get(key); id != nil && id.String() != "undefined" {
		// we need to update the row
		idx := array.nbElements - 1 - id.Int()
		array.rows.SetIndex(idx, js.Global.Get("Array").New(jsCells...))
	} else {
		// we need to insert a row
		// meaning we need to shift every entry a slot after
		for i := array.nbElements; i > 0; i-- {
			array.rows.SetIndex(i, array.rows.Index(i-1))
		}
		array.rows.SetIndex(0, js.Global.Get("Array").New(jsCells...))
		array.indices.Set(key, array.nbElements)
		array.nbElements = array.nbElements + 1
	}
}

type JSCell struct {
	Type  interface{}
	Value interface{}
}

func (jsCell *JSCell) toJSObject() *js.Object {
	toReturn := newJSObject()
	toReturn.Set("type", jsCell.Type)
	toReturn.Set("value", jsCell.Value)
	return toReturn
}
