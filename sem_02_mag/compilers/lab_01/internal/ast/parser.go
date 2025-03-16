package ast

import (
	"fmt"

	"github.com/polnaya-katuxa/bmstu/sem_02_mag/compilers/lab_01/internal/common"
	"k8s.io/utils/set"
)

type parser struct {
	index     int
	leavesMap map[int]*Symbol
	tokens    []common.Token
	pos       int
	current   common.Token
}

func newParser(tokens []common.Token) *parser {
	return &parser{
		index:     0,
		leavesMap: make(map[int]*Symbol, len(tokens)),
		tokens:    tokens,
		pos:       0,
		current:   tokens[0],
	}
}

func (p *parser) GetLeavesMap() map[int]*Symbol {
	return p.leavesMap
}

func (p *parser) next() {
	p.pos++
	if p.pos < len(p.tokens) {
		p.current = p.tokens[p.pos]
	} else {
		p.current = common.Token{common.EOF, 0}
	}
}

func (p *parser) countFollowPos(root Node) {
	switch node := root.(type) {
	case *Concatenation:
		for _, v := range node.Left.GetLastPos() {
			leaf := p.leavesMap[v]
			leaf.FollowPos = set.New(leaf.FollowPos...).Union(set.New(node.Right.GetFirstPos()...)).SortedList()
		}
		p.countFollowPos(node.Left)
		p.countFollowPos(node.Right)
	case *KleeneStar:
		for _, v := range node.GetLastPos() {
			leaf := p.leavesMap[v]
			leaf.FollowPos = set.New(leaf.FollowPos...).Union(set.New(node.GetFirstPos()...)).SortedList()
		}
		p.countFollowPos(node.Child)
	case *Alternation:
		p.countFollowPos(node.Left)
		p.countFollowPos(node.Right)
	default:
		return
	}
}

func (p *parser) parseAlternation() (Node, error) {
	node, err := p.parseConcatenation()
	if err != nil {
		return nil, fmt.Errorf("invalid concat: %w", err)
	}

	for p.current.Type == common.Pipe {
		p.next()
		right, err := p.parseConcatenation()
		if err != nil {
			return nil, fmt.Errorf("invalid right concat: %w", err)
		}

		node = &Alternation{
			Left:  node,
			Right: right,
		}
	}

	return node, nil
}

func (p *parser) parseConcatenation() (Node, error) {
	node, err := p.parseQuantifier()
	if err != nil {
		return nil, fmt.Errorf("invalid term: %w", err)
	}

	for {
		switch p.current.Type {
		case common.Symbol, common.LParen:
			right, err := p.parseQuantifier()
			if err != nil {
				return nil, fmt.Errorf("invalid right term: %w", err)
			}

			node = &Concatenation{
				Left:  node,
				Right: right,
			}
		default:
			return node, nil
		}
	}
}

func (p *parser) parseQuantifier() (Node, error) {
	node, err := p.parseSymbolOrGroup()
	if err != nil {
		return nil, fmt.Errorf("invalid factor: %w", err)
	}

	for {
		switch p.current.Type {
		case common.KleeneStar:
			p.next()
			node = &KleeneStar{
				Child: node,
			}
		default:
			return node, nil
		}
	}
}

func (p *parser) parseSymbolOrGroup() (Node, error) {
	switch p.current.Type {
	case common.Symbol:
		sym := &Symbol{
			Index: p.index,
			Value: p.current.Value,
		}
		p.leavesMap[p.index] = sym
		p.index++
		p.next()
		return sym, nil
	case common.LParen:
		p.next()
		node, err := p.parseAlternation()
		if err != nil {
			return nil, fmt.Errorf("invalid alt: %w", err)
		}

		if p.current.Type != common.RParen {
			return nil, fmt.Errorf("invalid token: %w", common.ErrUnclosedParen)
		}
		p.next()
		return node, nil
	default:
		return nil, fmt.Errorf("invalid token: %s", string(p.current.Value))
	}
}
