package main

import (
	"fmt"

	"github.com/polnaya-katuxa/bmstu/sem_02_mag/compilers/lab_02/internal/grammar"
	"github.com/polnaya-katuxa/bmstu/sem_02_mag/compilers/lab_02/internal/greibach"
)

func main() {
	g, err := grammar.NewFromFile("data/in.txt")
	if err != nil {
		panic(err)
	}

	fmt.Println(greibach.Greibach(g).String())
}
