package grammar

const (
	Empty = "ε"
)

type Rule struct {
	NonTerminal  string
	Combinations [][]string
}

type Grammar struct {
	NonTerminals []string
	Terminals    []string
	Start        string

	Rules []Rule
}
