package common

// Типы токенов
const (
	LParen TokenType = iota
	RParen
	KleeneStar
	Pipe
	Symbol
	Hash
	EOF
)

type TokenType int

type Token struct {
	Type  TokenType
	Value rune
}
