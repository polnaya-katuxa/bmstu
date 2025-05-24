package compiler

import (
	_ "embed"
	"encoding/json"
	"fmt"
	"os"
	"os/exec"

	"github.com/antlr4-go/antlr/v4"
	"github.com/polnaya-katuxa/bmstu/sem_02_mag/compilers/course/src/internal/ast"
	"github.com/polnaya-katuxa/bmstu/sem_02_mag/compilers/course/src/internal/parser"
	"github.com/polnaya-katuxa/bmstu/sem_02_mag/compilers/course/src/internal/visitor"
)

func Compile(in, out string) error {
	inputStream, err := antlr.NewFileStream(in)
	if err != nil {
		return fmt.Errorf("cannot create input stream: %w", err)
	}

	lexer := parser.NewLuaLexer(inputStream)
	stream := antlr.NewCommonTokenStream(lexer, 0)
	p := parser.NewLuaParser(stream)
	lexerErrors := &customErrorListener{}
	p.AddErrorListener(lexerErrors)
	tree := p.Chunk()
	if lexerErrors.IsError {
		return fmt.Errorf("there are syntax errors")
	}

	t := ast.BuildAst(tree.(antlr.ParseTree))
	treeJson, err := json.MarshalIndent(t, "", "  ")
	if err != nil {
		return fmt.Errorf("there are errors during json marshal")
	}

	err = os.WriteFile("examples/tree.json", treeJson, 0666)
	if err != nil {
		return fmt.Errorf("write tree to file: %w", err)
	}

	v := visitor.New()
	v.Visit(tree)
	if len(v.ErrorList) > 0 {
		for _, v := range v.ErrorList {
			fmt.Println(v)
		}

		return fmt.Errorf("there are errors during compilation")
	}

	ir := v.IR.String()

	llFile := out + ".ll"

	err = os.WriteFile(llFile, []byte(ir), 0666)
	if err != nil {
		return fmt.Errorf("write ir to file: %w", err)
	}

	cmd := exec.Command("clang", llFile, "-o", out, "-lm")
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	err = cmd.Run()
	if err != nil {
		return fmt.Errorf("compile llvm to program: %w", err)
	}

	return nil
}
