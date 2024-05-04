package main

import (
	"io/ioutil"
)

func main() {
	// This should throw linter errors
	ioutil.TempDir("", "")
}
