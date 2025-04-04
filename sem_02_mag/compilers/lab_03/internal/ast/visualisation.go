package ast

import (
	"fmt"
	"os"
	"os/exec"
)

func Show(root Node) error {
	file, err := os.Create("data/ast.dot")
	if err != nil {
		return err
	}
	defer file.Close()

	_, err = file.WriteString("digraph AST {\n")
	if err != nil {
		return err
	}

	err = root.draw(file)
	if err != nil {
		return err
	}

	_, err = file.WriteString("}\n")
	if err != nil {
		return err
	}

	cmd := exec.Command("dot", "-Tpng", "data/ast.dot", "-o", "data/ast.png")
	err = cmd.Run()
	if err != nil {
		return err
	}

	// cmd = exec.Command("open", "data/ast.png")
	// err = cmd.Run()
	// if err != nil {
	// 	return err
	// }

	return nil
}

func (n *Program) draw(file *os.File) error {
	_, err := fmt.Fprintf(file, "\"%p\" [label=\"%s\"];\n", n, "<program>")
	if err != nil {
		return err
	}

	_, err = fmt.Fprintf(file, "\"%p\" -> \"%p\";\n", n, n.Block)
	if err != nil {
		return err
	}

	return n.Block.draw(file)
}

func (n *Block) draw(file *os.File) error {
	_, err := fmt.Fprintf(file, "\"%p\" [label=\"%s\"];\n", n, "<block>")
	if err != nil {
		return err
	}

	_, err = fmt.Fprintf(file, "\"%p\" -> \"%p\";\n", n, n.OperatorList)
	if err != nil {
		return err
	}

	return n.OperatorList.draw(file)
}

func (n *OperatorList) draw(file *os.File) error {
	_, err := fmt.Fprintf(file, "\"%p\" [label=\"%s\"];\n", n, "<operator list>")
	if err != nil {
		return err
	}

	if n.Operator != nil {
		_, err = fmt.Fprintf(file, "\"%p\" -> \"%p\";\n", n, n.Operator)
		if err != nil {
			return err
		}

		err = n.Operator.draw(file)
		if err != nil {
			return err
		}
	}

	if n.OperatorList != nil {
		_, err = fmt.Fprintf(file, "\"%p\" -> \"%p\";\n", n, n.OperatorList)
		if err != nil {
			return err
		}

		err = n.OperatorList.draw(file)
		if err != nil {
			return err
		}
	}

	return nil
}

func (n *Operator) draw(file *os.File) error {
	_, err := fmt.Fprintf(file, "\"%p\" [label=\"%s\"];\n", n, "<operator>")
	if err != nil {
		return err
	}

	if n.Identifier != nil {
		_, err = fmt.Fprintf(file, "\"%p\" -> \"%p\";\n", n, n.Identifier)
		if err != nil {
			return err
		}

		err = n.Identifier.draw(file)
		if err != nil {
			return err
		}
	}

	if n.Expression != nil {
		_, err = fmt.Fprintf(file, "\"%p\" -> \"%p\";\n", n, n.Expression)
		if err != nil {
			return err
		}

		err = n.Expression.draw(file)
		if err != nil {
			return err
		}
	}

	return nil
}

func (n *Identifier) draw(file *os.File) error {
	_, err := fmt.Fprintf(file, "\"%p\" [label=\"%s (%s)\"];\n", n, "<identifier>", n.Value)
	if err != nil {
		return err
	}

	return nil
}

func (n *Expression) draw(file *os.File) error {
	_, err := fmt.Fprintf(file, "\"%p\" [label=\"%s\"];\n", n, "<expression>")
	if err != nil {
		return err
	}

	if n.Left != nil {
		_, err = fmt.Fprintf(file, "\"%p\" -> \"%p\";\n", n, n.Left)
		if err != nil {
			return err
		}

		err = n.Left.draw(file)
		if err != nil {
			return err
		}
	}

	if n.RelationOperation != nil {
		_, err = fmt.Fprintf(file, "\"%p\" -> \"%p\";\n", n, n.RelationOperation)
		if err != nil {
			return err
		}

		err = n.RelationOperation.draw(file)
		if err != nil {
			return err
		}
	}

	if n.Right != nil {
		_, err = fmt.Fprintf(file, "\"%p\" -> \"%p\";\n", n, n.Right)
		if err != nil {
			return err
		}

		err = n.Right.draw(file)
		if err != nil {
			return err
		}
	}

	return nil
}

func (n *ArithmeticalExpression) draw(file *os.File) error {
	_, err := fmt.Fprintf(file, "\"%p\" [label=\"%s\"];\n", n, "<arithmetical expression>")
	if err != nil {
		return err
	}

	if n.Term != nil {
		_, err = fmt.Fprintf(file, "\"%p\" -> \"%p\";\n", n, n.Term)
		if err != nil {
			return err
		}

		err = n.Term.draw(file)
		if err != nil {
			return err
		}
	}

	if n.ArithmeticalExpression != nil {
		_, err = fmt.Fprintf(file, "\"%p\" -> \"%p\";\n", n, n.ArithmeticalExpression)
		if err != nil {
			return err
		}

		err = n.ArithmeticalExpression.draw(file)
		if err != nil {
			return err
		}
	}

	return nil
}

func (n *Term) draw(file *os.File) error {
	_, err := fmt.Fprintf(file, "\"%p\" [label=\"%s\"];\n", n, "<term>")
	if err != nil {
		return err
	}

	if n.Factor != nil {
		_, err = fmt.Fprintf(file, "\"%p\" -> \"%p\";\n", n, n.Factor)
		if err != nil {
			return err
		}

		err = n.Factor.draw(file)
		if err != nil {
			return err
		}
	}

	if n.Term != nil {
		_, err = fmt.Fprintf(file, "\"%p\" -> \"%p\";\n", n, n.Term)
		if err != nil {
			return err
		}

		err = n.Term.draw(file)
		if err != nil {
			return err
		}
	}

	return nil
}

func (n *Factor) draw(file *os.File) error {
	_, err := fmt.Fprintf(file, "\"%p\" [label=\"%s\"];\n", n, "<factor>")
	if err != nil {
		return err
	}

	if n.Identifier != nil {
		_, err = fmt.Fprintf(file, "\"%p\" -> \"%p\";\n", n, n.Identifier)
		if err != nil {
			return err
		}

		err = n.Identifier.draw(file)
		if err != nil {
			return err
		}
	}

	if n.Constant != nil {
		_, err = fmt.Fprintf(file, "\"%p\" -> \"%p\";\n", n, n.Constant)
		if err != nil {
			return err
		}

		err = n.Constant.draw(file)
		if err != nil {
			return err
		}
	}

	if n.ArithmeticalExpression != nil {
		_, err = fmt.Fprintf(file, "\"%p\" -> \"%p\";\n", n, n.ArithmeticalExpression)
		if err != nil {
			return err
		}

		err = n.ArithmeticalExpression.draw(file)
		if err != nil {
			return err
		}
	}

	return nil
}

func (n *Constant) draw(file *os.File) error {
	_, err := fmt.Fprintf(file, "\"%p\" [label=\"%s (%s)\"];\n", n, "<constant>", n.Value)
	if err != nil {
		return err
	}

	return nil
}

func (n *OperatorListX) draw(file *os.File) error {
	_, err := fmt.Fprintf(file, "\"%p\" [label=\"%s\"];\n", n, "<operator list>'")
	if err != nil {
		return err
	}

	if n.Operator != nil {
		_, err = fmt.Fprintf(file, "\"%p\" -> \"%p\";\n", n, n.Operator)
		if err != nil {
			return err
		}

		err = n.Operator.draw(file)
		if err != nil {
			return err
		}
	}

	if n.OperatorList != nil {
		_, err = fmt.Fprintf(file, "\"%p\" -> \"%p\";\n", n, n.OperatorList)
		if err != nil {
			return err
		}

		err = n.OperatorList.draw(file)
		if err != nil {
			return err
		}
	}

	return nil
}

func (n *RelationOperation) draw(file *os.File) error {
	_, err := fmt.Fprintf(file, "\"%p\" [label=\"%s (%s)\"];\n", n, "<relation operation>", n.Value)
	if err != nil {
		return err
	}

	return nil
}

func (n *ArithmeticalExpressionX) draw(file *os.File) error {
	_, err := fmt.Fprintf(file, "\"%p\" [label=\"%s\"];\n", n, "<arithmetical expression>'")
	if err != nil {
		return err
	}

	if n.SumOperation != nil {
		_, err = fmt.Fprintf(file, "\"%p\" -> \"%p\";\n", n, n.SumOperation)
		if err != nil {
			return err
		}

		err = n.SumOperation.draw(file)
		if err != nil {
			return err
		}
	}

	if n.Term != nil {
		_, err = fmt.Fprintf(file, "\"%p\" -> \"%p\";\n", n, n.Term)
		if err != nil {
			return err
		}

		err = n.Term.draw(file)
		if err != nil {
			return err
		}
	}

	if n.ArithmeticalExpression != nil {
		_, err = fmt.Fprintf(file, "\"%p\" -> \"%p\";\n", n, n.ArithmeticalExpression)
		if err != nil {
			return err
		}

		err = n.ArithmeticalExpression.draw(file)
		if err != nil {
			return err
		}
	}

	return nil
}

func (n *TermX) draw(file *os.File) error {
	_, err := fmt.Fprintf(file, "\"%p\" [label=\"%s\"];\n", n, "<term>'")
	if err != nil {
		return err
	}

	if n.MulOperation != nil {
		_, err = fmt.Fprintf(file, "\"%p\" -> \"%p\";\n", n, n.MulOperation)
		if err != nil {
			return err
		}

		err = n.MulOperation.draw(file)
		if err != nil {
			return err
		}
	}

	if n.Factor != nil {
		_, err = fmt.Fprintf(file, "\"%p\" -> \"%p\";\n", n, n.Factor)
		if err != nil {
			return err
		}

		err = n.Factor.draw(file)
		if err != nil {
			return err
		}
	}

	if n.Term != nil {
		_, err = fmt.Fprintf(file, "\"%p\" -> \"%p\";\n", n, n.Term)
		if err != nil {
			return err
		}

		err = n.Term.draw(file)
		if err != nil {
			return err
		}
	}

	return nil
}

func (n *SumOperation) draw(file *os.File) error {
	_, err := fmt.Fprintf(file, "\"%p\" [label=\"%s (%s)\"];\n", n, "<sum operation>", n.Value)
	if err != nil {
		return err
	}

	return nil
}

func (n *MulOperation) draw(file *os.File) error {
	_, err := fmt.Fprintf(file, "\"%p\" [label=\"%s (%s)\"];\n", n, "<mul operation>", n.Value)
	if err != nil {
		return err
	}

	return nil
}
