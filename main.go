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

package main

import (
	"log"
	"net/http"

	"github.com/gobuffalo/packr"
)

//go:generate packr build
//go:generate echo "built fat http server binary with assets"
func main() {
	box := packr.NewBox("./frontend") // no need to pack the whole folder here, a dist subfolder would be better

	log.Println("serving zsldemo static website...")
	http.Handle("/", http.FileServer(box))
	log.Fatal(http.ListenAndServe(":8001", nil))
}
