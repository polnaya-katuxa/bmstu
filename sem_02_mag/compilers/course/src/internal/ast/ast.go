package ast

import (
	"fmt"

	"github.com/antlr4-go/antlr/v4"
)

func BuildAst(tree antlr.ParseTree) *Node {
	res := &Node{
		Name:     fmt.Sprintf("%T", tree),
		Value:    tree.GetText(),
		Children: []*Node{},
	}

	for _, v := range tree.GetChildren() {
		res.Children = append(res.Children, BuildAst(v.(antlr.ParseTree)))
	}

	return res
}
