package main

import (
	"fmt"
	"os"

	"github.com/polnaya-katuxa/bmstu/sem_02_mag/compilers/lab_03/internal/ast"
	"github.com/polnaya-katuxa/bmstu/sem_02_mag/compilers/lab_03/internal/token"
)

func main() {
	tokens, err := token.ReadTokens("data/program.txt")
	if err != nil {
		fmt.Println("read tokens: ", err)
		os.Exit(1)
	}

	tree, err := ast.New(tokens)
	if err != nil {
		fmt.Println("make ast: ", err)
		os.Exit(1)
	}

	fmt.Println(ast.DepthPass(tree))

	err = ast.Show(tree)
	if err != nil {
		fmt.Println("show ast: ", err)
		os.Exit(1)
	}

	ast.Postfix = true
	err = ast.Show(tree)
	if err != nil {
		fmt.Println("show ast postfix: ", err)
		os.Exit(1)
	}
}
