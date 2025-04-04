package grammar

import (
	"bufio"
	"fmt"
	"os"
	"slices"
	"strconv"
	"strings"
)

func NewFromFile(path string) (*Grammar, error) {
	f, err := os.Open(path)
	if err != nil {
		return nil, fmt.Errorf("open grammar file: %w", err)
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)

	nonTerminals, err := readSymbols(scanner)
	if err != nil {
		return nil, fmt.Errorf("read non-terminal symbols: %w", err)
	}

	terminals, err := readSymbols(scanner)
	if err != nil {
		return nil, fmt.Errorf("read non-terminal symbols: %w", err)
	}

	rules, err := readRules(scanner)
	if err != nil {
		return nil, fmt.Errorf("read rules: %w", err)
	}

	if !scanner.Scan() {
		return nil, fmt.Errorf("read start symbol: unexpected EOF")
	}

	g := &Grammar{
		NonTerminals: nonTerminals,
		Terminals:    terminals,
		Start:        strings.TrimSpace(scanner.Text()),
		Rules:        rules,
	}

	if !g.Valid() {
		return nil, fmt.Errorf("invalid grammar")
	}

	return g, nil
}

func (g *Grammar) Valid() bool {
	if !slices.Contains(g.NonTerminals, g.Start) {
		return false
	}

	allSymbols := append(slices.Clone(g.NonTerminals), g.Terminals...)

	for _, rule := range g.Rules {
		if !slices.Contains(g.NonTerminals, rule.NonTerminal) {
			return false
		}

		for _, combination := range rule.Combinations {
			for _, symbol := range combination {
				if !slices.Contains(allSymbols, symbol) && symbol != Empty {
					return false
				}
			}
		}
	}

	return true
}

func (g *Grammar) String() string {
	result := ""

	result += strconv.Itoa(len(g.NonTerminals)) + "\n"
	for _, nonTerminal := range g.NonTerminals {
		result += string(nonTerminal) + " "
	}
	result += "\n"

	result += strconv.Itoa(len(g.Terminals)) + "\n"
	for _, terminal := range g.Terminals {
		result += string(terminal) + " "
	}
	result += "\n"

	rulesRaw := make([]string, 0, len(g.Rules))
	for _, rule := range g.Rules {
		for _, combination := range rule.Combinations {
			rulesRaw = append(rulesRaw, string(rule.NonTerminal)+" -> "+strings.Join(combination, ""))
		}
	}

	result += strconv.Itoa(len(rulesRaw)) + "\n"
	for _, ruleRaw := range rulesRaw {
		result += string(ruleRaw) + "\n"
	}

	result += string(g.Start)
	result += "\n"

	return result
}

func readSymbols(scanner *bufio.Scanner) ([]string, error) {
	if !scanner.Scan() {
		return nil, fmt.Errorf("read symbols count: unexpected EOF")
	}

	n, err := strconv.Atoi(scanner.Text())
	if err != nil {
		return nil, fmt.Errorf("read symbols count: %w", err)
	}

	if !scanner.Scan() {
		return nil, fmt.Errorf("read symbols: unexpected EOF")
	}

	symbols := strings.Split(strings.TrimSpace(scanner.Text()), " ")

	if len(symbols) != n {
		return nil, fmt.Errorf("unexpected number of symbols: %d != %d", len(symbols), n)
	}

	return symbols, nil
}

func readRules(scanner *bufio.Scanner) ([]Rule, error) {
	if !scanner.Scan() {
		return nil, fmt.Errorf("read rules count: unexpected EOF")
	}

	n, err := strconv.Atoi(scanner.Text())
	if err != nil {
		return nil, fmt.Errorf("read rules count: %w", err)
	}

	rulesMap := make(map[string]int)
	rules := make([]Rule, 0, n)

	for i := range n {
		if !scanner.Scan() {
			return nil, fmt.Errorf("read rule %d: unexpected EOF", i)
		}

		splitted := strings.Split(strings.ReplaceAll(scanner.Text(), " ", ""), "->")
		if len(splitted) != 2 {
			return nil, fmt.Errorf("read rule %d: unexpected number of rule parts: %d", i, len(splitted))
		}

		nonTerminal := splitted[0]
		combination := splitSymbols(splitted[1])

		index, ok := rulesMap[nonTerminal]
		if ok {
			rules[index].Combinations = append(rules[index].Combinations, combination)
		} else {
			rules = append(rules, Rule{
				NonTerminal:  nonTerminal,
				Combinations: [][]string{combination},
			})
			rulesMap[nonTerminal] = len(rules) - 1
		}
	}

	return rules, nil
}

func splitSymbols(s string) []string {
	result := make([]string, 0, len(s))

	for _, symbol := range s {
		if symbol == '\'' {
			result[len(result)-1] += string(symbol)
		} else {
			result = append(result, string(symbol))
		}
	}

	return result
}
