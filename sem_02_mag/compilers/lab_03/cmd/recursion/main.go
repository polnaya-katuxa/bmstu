package main

import (
	"fmt"

	"github.com/polnaya-katuxa/bmstu/sem_02_mag/compilers/lab_03/internal/grammar"
	"github.com/polnaya-katuxa/bmstu/sem_02_mag/compilers/lab_03/internal/recursion"
)

const (
	version = 1
)

func main() {
	g, err := grammar.NewFromFile("data/in.txt")
	if err != nil {
		panic(err)
	}

	if version == 1 {
		fmt.Println(recursion.RemoveLeftRecursionV1(g).String())
	} else {
		fmt.Println(recursion.RemoveLeftRecursionV0(g).String())
	}
}
