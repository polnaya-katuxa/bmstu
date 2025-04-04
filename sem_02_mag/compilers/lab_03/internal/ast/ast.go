package ast

import (
	"fmt"
	"os"

	"github.com/polnaya-katuxa/bmstu/sem_02_mag/compilers/lab_03/internal/token"
)

type Node interface {
	draw(file *os.File) error
}

func New(tokens []token.Token) (Node, error) {
	p := newParser(tokens)
	program, err := p.parseProgram()
	if err != nil {
		return nil, fmt.Errorf("parse program: %w", err)
	}

	return program, nil
}

type Program struct {
	Block *Block
}

type Block struct {
	OperatorList *OperatorList
}

type OperatorList struct {
	Operator     *Operator
	OperatorList *OperatorListX
}

type Operator struct {
	Identifier *Identifier
	Expression *Expression
}

type Expression struct {
	Left              *ArithmeticalExpression
	RelationOperation *RelationOperation
	Right             *ArithmeticalExpression
}

type ArithmeticalExpression struct {
	Term                   *Term
	ArithmeticalExpression *ArithmeticalExpressionX
}

type Term struct {
	Factor *Factor
	Term   *TermX
}

type Factor struct {
	Identifier             *Identifier
	Constant               *Constant
	ArithmeticalExpression *ArithmeticalExpression
}

type RelationOperation struct {
	Value string
}

type SumOperation struct {
	Value string
}

type MulOperation struct {
	Value string
}

type OperatorListX struct {
	Operator     *Operator
	OperatorList *OperatorListX
}

type ArithmeticalExpressionX struct {
	SumOperation           *SumOperation
	Term                   *Term
	ArithmeticalExpression *ArithmeticalExpressionX
}

type TermX struct {
	MulOperation *MulOperation
	Factor       *Factor
	Term         *TermX
}

type Identifier struct {
	Value string
}

type Constant struct {
	Value string
}
