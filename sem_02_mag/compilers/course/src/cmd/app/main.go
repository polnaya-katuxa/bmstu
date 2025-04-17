package main

import (
	"fmt"
	"os"

	"github.com/polnaya-katuxa/bmstu/sem_02_mag/compilers/course/src/internal/compiler"
)

// TODO: cmp int & float
// TODO: null checks

func main() {
	err := compiler.Compile("examples/wiki.lua", "examples/out")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
