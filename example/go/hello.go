/*
Copyright 2014 Google Inc.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

     http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package main

import (
	"log"
	"os"
	"path/filepath"
)

const FILE = "test.txt"

func main() {
	dir, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}
	f, err := os.OpenFile(filepath.Join(dir, FILE), os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatal(err)
	}
	var array []string
	array = append(array, "Item1", "Item2")
	for _, item := range array {
		if _, err := f.WriteString(item + "\n"); err != nil {
			log.Fatal(err)
		}
	}
	err = f.Close()
	if err != nil {
		log.Fatal(err)
	}
}
