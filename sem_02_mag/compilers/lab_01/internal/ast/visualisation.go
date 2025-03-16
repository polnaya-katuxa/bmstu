package ast

import (
	"fmt"
	"os"
	"os/exec"
)

func printAST(node Node, prefix string, isLast bool) {
	branch := "├── "
	if isLast {
		branch = "└── "
	}

	fmt.Print(prefix + branch)
	fmt.Printf("%s\n", node)

	newPrefix := prefix
	if isLast {
		newPrefix += "    "
	} else {
		newPrefix += "│   "
	}

	switch n := node.(type) {
	case *Concatenation:
		printAST(n.Right, newPrefix, false)
		printAST(n.Left, newPrefix, true)
	case *Alternation:
		printAST(n.Right, newPrefix, false)
		printAST(n.Left, newPrefix, true)
	case *KleeneStar:
		printAST(n.Child, newPrefix, true)
	case *Symbol:
	}
}

func (a *AST) Show() error {
	file, err := os.Create("temp/ast.dot")
	if err != nil {
		return err
	}
	defer file.Close()

	_, err = file.WriteString("digraph AST {\n")
	if err != nil {
		return err
	}

	err = a.writeNode(file, a.root)
	if err != nil {
		return err
	}

	_, err = file.WriteString("}\n")
	if err != nil {
		return err
	}

	cmd := exec.Command("dot", "-Tpng", "temp/ast.dot", "-o", "temp/ast.png")
	err = cmd.Run()
	if err != nil {
		return err
	}

	cmd = exec.Command("open", "temp/ast.png")
	err = cmd.Run()
	if err != nil {
		return err
	}

	return nil
}

func (a *AST) writeNode(file *os.File, node Node) error {
	_, err := fmt.Fprintf(file, "\"%p\" [label=\"%s\"];\n", node, node.String())
	if err != nil {
		return err
	}

	switch n := node.(type) {
	case *Concatenation:
		_, err = fmt.Fprintf(file, "\"%p\" -> \"%p\";\n", n, n.Left)
		if err != nil {
			return err
		}
		_, err = fmt.Fprintf(file, "\"%p\" -> \"%p\";\n", n, n.Right)
		if err != nil {
			return err
		}
		err = a.writeNode(file, n.Left)
		if err != nil {
			return err
		}
		err = a.writeNode(file, n.Right)
		if err != nil {
			return err
		}
	case *Alternation:
		_, err = fmt.Fprintf(file, "\"%p\" -> \"%p\";\n", n, n.Left)
		if err != nil {
			return err
		}
		_, err = fmt.Fprintf(file, "\"%p\" -> \"%p\";\n", n, n.Right)
		if err != nil {
			return err
		}
		err = a.writeNode(file, n.Left)
		if err != nil {
			return err
		}
		err = a.writeNode(file, n.Right)
		if err != nil {
			return err
		}
	case *KleeneStar:
		_, err = fmt.Fprintf(file, "\"%p\" -> \"%p\";\n", n, n.Child)
		if err != nil {
			return err
		}
		err = a.writeNode(file, n.Child)
		if err != nil {
			return err
		}
	case *Symbol:
	}
	return nil
}
