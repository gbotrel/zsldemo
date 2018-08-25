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

import (
	"crypto/rand"
	"encoding/hex"

	"github.com/gopherjs/gopherjs/js"
)

func copyToClipboard(txt string) bool {
	// get application div
	divApplication := js.Global.Get("document").Call("querySelector", ".application")

	// create element
	textArea := js.Global.Get("document").Call("createElement", "textarea")

	// append element
	divApplication.Call("appendChild", textArea)

	textArea.Set("value", txt)
	textArea.Call("select")

	toReturn := js.Global.Get("document").Call("execCommand", "copy").Bool()
	divApplication.Call("removeChild", textArea)
	return toReturn
}

func truncateString(input string, num int) string {
	toReturn := input
	if len(input) > num {
		toReturn = input[0:num] + ".."
	}
	return toReturn
}

func newJSObject() *js.Object {
	return js.Global.Get("Object").New()
}

func hexEncode0x(input []byte) string {
	return "0x" + hex.EncodeToString(input)
}

func hexDecode0x(input string) []byte {
	bValue, err := hex.DecodeString(input[2:])
	if err != nil {
		panic(err)
	}
	return bValue
}

func randomBytes(length uint) []byte {
	// importing "crypto/rand" adds 0.3M to JS output.
	toReturn := make([]byte, length)
	if _, err := rand.Read(toReturn); err != nil {
		panic(err)
	}
	return toReturn
}

func parseTreePath(treePath []string) [][]byte {
	toReturn := make([][]byte, len(treePath))
	for i, v := range treePath {
		toReturn[i] = hexDecode0x(v)
	}
	return toReturn
}
