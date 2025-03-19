package ast

import (
	"fmt"
	"log/slog"

	"github.com/polnaya-katuxa/bmstu/sem_02_mag/compilers/lab_01/internal/common"
	"github.com/polnaya-katuxa/bmstu/sem_02_mag/compilers/lab_01/internal/regexp"
	"k8s.io/utils/set"
)

// <alternation>     ::=     <concatenation> "|" <alternation> | <concatenation>
// <concatenation>   ::=     <quantifier> <concatenation> | <quantifier>
// <quantifier>      ::=     <symbol-or-group> "*" | <symbol-or-group> "+" | <symbol-or-group>
// <symbol-or-group> ::=     "(" <alternation> ")" | <symbol>

type AST struct {
	root       Node
	symbols    []rune
	symbolsMap map[int]*Symbol
}

func NewAST(regexp *regexp.Regexp) (*AST, error) {
	slog.Info("start tokenize")

	tokens, err := regexp.Tokenize()
	if err != nil {
		return nil, fmt.Errorf("tokenize: %w", err)
	}

	symbols := getSymbolsList(tokens)

	slog.Info("start building AST by tokens")

	parser := newParser(tokens)
	root, err := parser.parseAlternation()
	if err != nil {
		return nil, fmt.Errorf("parse: %w", err)
	}

	slog.Info("counting follow pos")

	parser.countFollowPos(root)
	symbolsMap := parser.GetLeavesMap()

	slog.Info("built AST")

	return &AST{
		root:       root,
		symbols:    symbols,
		symbolsMap: symbolsMap,
	}, nil
}

func getSymbolsList(tokens []common.Token) []rune {
	symbols := make([]rune, 0, len(tokens))
	for _, t := range tokens {
		if t.Type == common.Symbol {
			symbols = append(symbols, t.Value)
		}
	}

	return symbols
}

func (a *AST) GetRoot() Node {
	return a.root
}

func (a *AST) GetSymbols() []rune {
	return a.symbols
}

func (a *AST) GetSymbolsMap() map[int]*Symbol {
	return a.symbolsMap
}

func (a *AST) Print() {
	printAST(a.root, "", true)
}

type Node interface {
	String() string
	IsNullable() bool
	GetFirstPos() []int
	GetLastPos() []int
}

type Concatenation struct {
	Nullable *bool
	FirstPos []int
	LastPos  []int

	Left  Node
	Right Node
}

func (c *Concatenation) IsNullable() bool {
	if c.Nullable != nil {
		return *c.Nullable
	}

	nullable := c.Left.IsNullable() && c.Right.IsNullable()
	c.Nullable = &nullable

	return nullable
}

func (c *Concatenation) GetFirstPos() []int {
	if c.FirstPos != nil {
		return c.FirstPos
	}

	if c.Left.IsNullable() {
		c.FirstPos = set.New(c.Left.GetFirstPos()...).Union(set.New(c.Right.GetFirstPos()...)).SortedList()
	} else {
		c.FirstPos = set.New(c.Left.GetFirstPos()...).SortedList()
	}

	return c.FirstPos
}

func (c *Concatenation) GetLastPos() []int {
	if c.LastPos != nil {
		return c.LastPos
	}

	if c.Right.IsNullable() {
		c.LastPos = set.New(c.Right.GetLastPos()...).Union(set.New(c.Left.GetLastPos()...)).SortedList()
	} else {
		c.LastPos = set.New(c.Right.GetLastPos()...).SortedList()
	}

	return c.LastPos
}

func (c *Concatenation) String() string {
	return "●"
}

type Alternation struct {
	Nullable *bool
	FirstPos []int
	LastPos  []int

	Left  Node
	Right Node
}

func (a *Alternation) IsNullable() bool {
	if a.Nullable != nil {
		return *a.Nullable
	}

	nullable := a.Left.IsNullable() || a.Right.IsNullable()
	a.Nullable = &nullable

	return nullable
}

func (a *Alternation) GetFirstPos() []int {
	if a.FirstPos != nil {
		return a.FirstPos
	}

	a.FirstPos = set.New(a.Left.GetFirstPos()...).Union(set.New(a.Right.GetFirstPos()...)).SortedList()

	return a.FirstPos
}

func (a *Alternation) GetLastPos() []int {
	if a.LastPos != nil {
		return a.LastPos
	}

	a.LastPos = set.New(a.Left.GetLastPos()...).Union(set.New(a.Right.GetLastPos()...)).SortedList()

	return a.LastPos
}

func (a *Alternation) String() string {
	return "|"
}

type KleeneStar struct {
	FirstPos []int
	LastPos  []int

	Child Node
}

func (k *KleeneStar) IsNullable() bool { return true }

func (k *KleeneStar) GetFirstPos() []int {
	if k.FirstPos != nil {
		return k.FirstPos
	}

	k.FirstPos = k.Child.GetFirstPos()

	return k.FirstPos
}

func (k *KleeneStar) GetLastPos() []int {
	if k.LastPos != nil {
		return k.LastPos
	}

	k.LastPos = k.Child.GetLastPos()

	return k.LastPos
}

func (k *KleeneStar) String() string {
	return "*"
}

type Symbol struct {
	FirstPos  []int
	LastPos   []int
	FollowPos []int

	Index int

	Value rune
}

func (s *Symbol) IsNullable() bool { return false }

func (s *Symbol) GetFirstPos() []int {
	if s.FirstPos != nil {
		return s.FirstPos
	}

	s.FirstPos = []int{s.Index}

	return s.FirstPos
}

func (s *Symbol) GetLastPos() []int {
	if s.LastPos != nil {
		return s.LastPos
	}

	s.LastPos = []int{s.Index}

	return s.LastPos
}

func (s *Symbol) GetFollowPos() []int {
	return s.FollowPos
}

func (s *Symbol) String() string {
	return string(s.Value)
}
