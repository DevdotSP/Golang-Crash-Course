package main

import (
	"os"

	"github.com/davecgh/go-spew/spew"
)

func Dump(v any) {
	spew.Dump(v)
	os.Exit(1) // like dd()
}


func main() {

	helloWorlds := "hello"
	Dump(helloWorlds)
}