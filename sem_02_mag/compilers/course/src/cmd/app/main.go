package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/polnaya-katuxa/bmstu/sem_02_mag/compilers/course/src/internal/compiler"
)

// TODO: cmp int & float
// TODO: null checks

var (
	src = flag.String("i", "examples/wiki.lua", "source .lua file")
)

func main() {
	flag.Parse()

	srcFilename := "examples/wiki.lua"
	if src != nil {
		srcFilename = *src
	}
	err := compiler.Compile(srcFilename, "examples/out")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
