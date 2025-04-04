package token

import (
	"fmt"
	"os"
	"strconv"
	"strings"
	"unicode"
)

type TokenType int

const (
	EOF TokenType = iota
	LeftParen
	RightParen
	Less
	LessOrEqual
	Greater
	GreaterOrEqual
	Equal
	NotEqual
	Plus
	Minus
	Multiply
	Divide
	Constant
	Identifier
)

var tokenPriority = map[TokenType]int{
	Less:           0,
	LessOrEqual:    0,
	Greater:        0,
	GreaterOrEqual: 0,
	Equal:          0,
	NotEqual:       0,
	Plus:           1,
	Minus:          1,
	Multiply:       2,
	Divide:         2,
	Constant:       3,
	Identifier:     3,
	LeftParen:      4,
	RightParen:     4,
}

var typeMapping = map[string]TokenType{
	"(":  LeftParen,
	")":  RightParen,
	"<":  Less,
	"<=": LessOrEqual,
	">":  Greater,
	">=": GreaterOrEqual,
	"==": Equal,
	"<>": NotEqual,
	"+":  Plus,
	"-":  Minus,
	"*":  Multiply,
	"/":  Divide,
}

func getTokenType(s string) TokenType {
	t, ok := typeMapping[s]
	if ok {
		return t
	}

	_, err := strconv.ParseFloat(s, 64)
	if err == nil {
		return Constant
	}

	return Identifier
}

type Token struct {
	Type  TokenType
	Value string
}

func ReadTokens(path string) ([]Token, error) {
	content, err := os.ReadFile(path)
	if err != nil {
		return nil, fmt.Errorf("read file: %w", err)
	}

	var tokens []Token

	splitted := strings.Fields(string(content))
	for _, s := range splitted {
		buf := &strings.Builder{}
		previousSymbolKind := 0
		for i, symbol := range s {
			currentSymbolKind := getSymbolKind(symbol)

			if i == 0 {
				previousSymbolKind = currentSymbolKind
				buf.WriteRune(symbol)
				continue
			}

			if currentSymbolKind != previousSymbolKind {
				previousSymbolKind = currentSymbolKind
				tokens = append(tokens, Token{
					Value: buf.String(),
					Type:  getTokenType(buf.String()),
				})
				buf.Reset()
				buf.WriteRune(symbol)
				continue
			}

			buf.WriteRune(symbol)
		}

		if buf.Len() > 0 {
			tokens = append(tokens, Token{
				Value: buf.String(),
				Type:  getTokenType(buf.String()),
			})
		}
	}

	return tokens, nil
}

func getSymbolKind(symbol rune) int {
	if unicode.IsLetter(symbol) || unicode.IsNumber(symbol) || symbol == '.' {
		return 0
	}

	if symbol == '(' || symbol == ')' {
		return 1
	}

	return 2
}
