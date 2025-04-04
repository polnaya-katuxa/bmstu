package ast

import (
	"fmt"

	"github.com/polnaya-katuxa/bmstu/sem_02_mag/compilers/lab_03/internal/token"
)

type parser struct {
	tokens  []token.Token
	pos     int
	current token.Token
}

func newParser(tokens []token.Token) *parser {
	return &parser{
		tokens:  tokens,
		pos:     0,
		current: tokens[0],
	}
}

func (p *parser) next() {
	p.pos++
	if p.pos < len(p.tokens) {
		p.current = p.tokens[p.pos]
	} else {
		p.current = token.Token{Type: token.EOF, Value: ""}
	}
}

func (p *parser) parseProgram() (*Program, error) {
	block, err := p.parseBlock()
	if err != nil {
		return nil, fmt.Errorf("invalid block: %w", err)
	}

	return &Program{
		Block: block,
	}, nil
}

func (p *parser) parseBlock() (*Block, error) {
	if p.current.Type != token.StartBlock {
		return nil, fmt.Errorf("no start block")
	}

	p.next()

	operatorList, err := p.parseOperatorList()
	if err != nil {
		return nil, fmt.Errorf("invalid operator list: %w", err)
	}

	if p.current.Type != token.EndBlock {
		return nil, fmt.Errorf("no end block")
	}

	p.next()

	return &Block{
		OperatorList: operatorList,
	}, nil
}

func (p *parser) parseOperatorList() (*OperatorList, error) {
	operator, err := p.parseOperator()
	if err != nil {
		return nil, fmt.Errorf("invalid operator: %w", err)
	}

	ol := &OperatorList{
		Operator: operator,
	}

	if p.current.Type == token.EndBlock {
		return ol, nil
	}

	operatorListX, err := p.parseOperatorListX()
	if err != nil {
		return nil, fmt.Errorf("invalid operator list x: %w", err)
	}

	ol.OperatorList = operatorListX

	return ol, nil
}

func (p *parser) parseOperator() (*Operator, error) {
	identifier, err := p.parseIdentifier()
	if err != nil {
		return nil, fmt.Errorf("invalid identifier: %w", err)
	}

	if p.current.Type != token.Assignment {
		return nil, fmt.Errorf("no assign")
	}

	p.next()

	expression, err := p.parseExpression()
	if err != nil {
		return nil, fmt.Errorf("invalid expression: %w", err)
	}

	return &Operator{
		Identifier: identifier,
		Expression: expression,
	}, nil
}

func (p *parser) parseIdentifier() (*Identifier, error) {
	if p.current.Type != token.Identifier {
		return nil, fmt.Errorf("invalid identifier")
	}

	value := p.current.Value
	p.next()

	return &Identifier{
		Value: value,
	}, nil
}

func (p *parser) parseExpression() (*Expression, error) {
	left, err := p.parseArithmeticalExpression()
	if err != nil {
		return nil, fmt.Errorf("invalid left arithmetical expression: %w", err)
	}

	expr := &Expression{
		Left: left,
	}

	relation, err := p.parseRelationOperation()
	if err != nil {
		return expr, nil
	}

	expr.RelationOperation = relation

	right, err := p.parseArithmeticalExpression()
	if err != nil {
		return nil, fmt.Errorf("invalid right arithmetical expression: %w", err)
	}

	expr.Right = right

	return expr, nil
}

func (p *parser) parseRelationOperation() (*RelationOperation, error) {
	switch p.current.Type {
	case token.Less, token.LessOrEqual, token.Greater, token.GreaterOrEqual, token.Equal, token.NotEqual:
		value := p.current.Value
		p.next()
		return &RelationOperation{
			Value: value,
		}, nil
	default:
		return nil, fmt.Errorf("invalid relation operation")
	}
}

func (p *parser) parseArithmeticalExpression() (*ArithmeticalExpression, error) {
	term, err := p.parseTerm()
	if err != nil {
		return nil, fmt.Errorf("invalid term: %w", err)
	}

	ae := &ArithmeticalExpression{
		Term: term,
	}

	if p.current.Type != token.Plus && p.current.Type != token.Minus {
		return ae, nil
	}

	aex, err := p.parseArithmeticalExpressionX()
	if err != nil {
		return nil, fmt.Errorf("invalid aex: %w", err)
	}

	ae.ArithmeticalExpression = aex

	return ae, nil
}

func (p *parser) parseTerm() (*Term, error) {
	factor, err := p.parseFactor()
	if err != nil {
		return nil, fmt.Errorf("invalid factor: %w", err)
	}

	term := &Term{
		Factor: factor,
	}

	if p.current.Type != token.Multiply && p.current.Type != token.Divide {
		return term, nil
	}

	termX, err := p.parseTermX()
	if err != nil {
		return nil, fmt.Errorf("invalid term x: %w", err)
	}

	term.Term = termX

	return term, nil
}

func (p *parser) parseFactor() (*Factor, error) {
	f := &Factor{}

	switch p.current.Type {
	case token.Identifier:
		identifier, err := p.parseIdentifier()
		if err != nil {
			return nil, fmt.Errorf("invalid identifier: %w", err)
		}

		f.Identifier = identifier
	case token.Constant:
		constant, err := p.parseConstant()
		if err != nil {
			return nil, fmt.Errorf("invalid constant: %w", err)
		}

		f.Constant = constant
	case token.LeftParen:
		p.next()

		ae, err := p.parseArithmeticalExpression()
		if err != nil {
			return nil, fmt.Errorf("invalid arithmetical expression: %w", err)
		}

		if p.current.Type != token.RightParen {
			return nil, fmt.Errorf("no right paren")
		}

		p.next()

		f.ArithmeticalExpression = ae
	default:
		return nil, fmt.Errorf("invalid factor content")
	}

	return f, nil
}

func (p *parser) parseConstant() (*Constant, error) {
	if p.current.Type != token.Constant {
		return nil, fmt.Errorf("invalid constant")
	}

	value := p.current.Value
	p.next()

	return &Constant{
		Value: value,
	}, nil
}

func (p *parser) parseOperatorListX() (*OperatorListX, error) {
	if p.current.Type != token.Semicolon {
		return nil, fmt.Errorf("no semicolon")
	}

	p.next()

	operator, err := p.parseOperator()
	if err != nil {
		return nil, fmt.Errorf("invalid operator: %w", err)
	}

	olx := &OperatorListX{
		Operator: operator,
	}

	if p.current.Type != token.Semicolon {
		return olx, nil
	}

	operatorListX, err := p.parseOperatorListX()
	if err != nil {
		return nil, fmt.Errorf("invalid operator list x: %w", err)
	}

	olx.OperatorList = operatorListX

	return olx, nil
}

func (p *parser) parseArithmeticalExpressionX() (*ArithmeticalExpressionX, error) {
	aex := &ArithmeticalExpressionX{}

	sum, err := p.parseSumOperation()
	if err != nil {
		return nil, fmt.Errorf("invalid sum operation: %w", err)
	}

	aex.SumOperation = sum

	term, err := p.parseTerm()
	if err != nil {
		return nil, fmt.Errorf("invalid term: %w", err)
	}

	aex.Term = term

	if p.current.Type != token.Plus && p.current.Type != token.Minus {
		return aex, nil
	}

	aexIn, err := p.parseArithmeticalExpressionX()
	if err != nil {
		return nil, fmt.Errorf("invalid aex: %w", err)
	}

	aex.ArithmeticalExpression = aexIn

	return aex, nil
}

func (p *parser) parseSumOperation() (*SumOperation, error) {
	so := &SumOperation{}

	switch p.current.Type {
	case token.Plus:
	case token.Minus:
	default:
		return nil, fmt.Errorf("unknown sum operation")
	}

	so.Value = p.current.Value

	p.next()

	return so, nil
}

func (p *parser) parseTermX() (*TermX, error) {
	termX := &TermX{}

	mul, err := p.parseMulOperation()
	if err != nil {
		return nil, fmt.Errorf("invalid mul operation: %w", err)
	}

	termX.MulOperation = mul

	factor, err := p.parseFactor()
	if err != nil {
		return nil, fmt.Errorf("invalid tefactorrm: %w", err)
	}

	termX.Factor = factor

	if p.current.Type != token.Multiply && p.current.Type != token.Divide {
		return termX, nil
	}

	termIn, err := p.parseTermX()
	if err != nil {
		return nil, fmt.Errorf("invalid term x: %w", err)
	}

	termX.Term = termIn

	return termX, nil
}

func (p *parser) parseMulOperation() (*MulOperation, error) {
	mo := &MulOperation{}

	switch p.current.Type {
	case token.Multiply:
	case token.Divide:
	default:
		return nil, fmt.Errorf("unknown multiply operation")
	}

	mo.Value = p.current.Value

	p.next()

	return mo, nil
}
