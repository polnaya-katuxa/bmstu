package recursion

import (
	"slices"

	"github.com/polnaya-katuxa/bmstu/sem_02_mag/compilers/lab_02/internal/grammar"
)

func RemoveLeftRecursionV0(g *grammar.Grammar) *grammar.Grammar {
	n := len(g.Rules)
	for i := 0; i < n; i++ {
		for j := 0; j < i; j++ {
			// Замена Ai -> Aj
			newCombinations := make([][]string, 0, len(g.Rules[i].Combinations))
			for _, iCombination := range g.Rules[i].Combinations {
				if iCombination[0] != g.Rules[j].NonTerminal {
					newCombinations = append(newCombinations, iCombination)
				} else {
					for _, jCombination := range g.Rules[j].Combinations {
						newCombinations = append(newCombinations, append(slices.Clone(jCombination), iCombination[1:]...))
					}
				}
			}

			g.Rules[i].Combinations = newCombinations
		}

		// Устранение непосредственно левой рекурсии
		recursiveCombinations := make([][]string, 0, len(g.Rules[i].Combinations))
		nonRecursiveCombinations := make([][]string, 0, len(g.Rules[i].Combinations))
		for _, iCombination := range g.Rules[i].Combinations {
			if iCombination[0] == g.Rules[i].NonTerminal {
				recursiveCombinations = append(recursiveCombinations, iCombination)
			} else {
				nonRecursiveCombinations = append(nonRecursiveCombinations, iCombination)
			}
		}

		if len(recursiveCombinations) == 0 {
			continue
		}

		newNonTerminal := g.Rules[i].NonTerminal + "'"

		g.Rules[i].Combinations = make([][]string, 0, len(nonRecursiveCombinations))
		for _, nonRecursiveCombination := range nonRecursiveCombinations {
			g.Rules[i].Combinations = append(g.Rules[i].Combinations, append(nonRecursiveCombination, newNonTerminal))
		}

		newRule := grammar.Rule{
			NonTerminal:  newNonTerminal,
			Combinations: make([][]string, 0, len(recursiveCombinations)),
		}
		for _, recursiveCombination := range recursiveCombinations {
			newRule.Combinations = append(newRule.Combinations, append(recursiveCombination[1:], newNonTerminal))
		}
		newRule.Combinations = append(newRule.Combinations, []string{grammar.Empty})

		g.Rules = append(g.Rules, newRule)
		g.NonTerminals = append(g.NonTerminals, newNonTerminal)
	}

	return g
}

func RemoveLeftRecursionV1(g *grammar.Grammar) *grammar.Grammar {
	n := len(g.Rules)
	for i := 0; i < n; i++ {
		for j := 0; j < i; j++ {
			// Замена Ai -> Aj
			newCombinations := make([][]string, 0, len(g.Rules[i].Combinations))
			for _, iCombination := range g.Rules[i].Combinations {
				if iCombination[0] != g.Rules[j].NonTerminal {
					newCombinations = append(newCombinations, iCombination)
				} else {
					for _, jCombination := range g.Rules[j].Combinations {
						newCombinations = append(newCombinations, append(slices.Clone(jCombination), iCombination[1:]...))
					}
				}
			}

			g.Rules[i].Combinations = newCombinations
		}

		// Устранение непосредственно левой рекурсии
		recursiveCombinations := make([][]string, 0, len(g.Rules[i].Combinations))
		nonRecursiveCombinations := make([][]string, 0, len(g.Rules[i].Combinations))
		for _, iCombination := range g.Rules[i].Combinations {
			if iCombination[0] == g.Rules[i].NonTerminal {
				recursiveCombinations = append(recursiveCombinations, iCombination)
			} else {
				nonRecursiveCombinations = append(nonRecursiveCombinations, iCombination)
			}
		}

		if len(recursiveCombinations) == 0 {
			continue
		}

		newNonTerminal := g.Rules[i].NonTerminal + "'"

		g.Rules[i].Combinations = make([][]string, 0, len(nonRecursiveCombinations))
		for _, nonRecursiveCombination := range nonRecursiveCombinations {
			g.Rules[i].Combinations = append(g.Rules[i].Combinations, slices.Clone(nonRecursiveCombination))
			g.Rules[i].Combinations = append(g.Rules[i].Combinations, append(nonRecursiveCombination, newNonTerminal))
		}

		newRule := grammar.Rule{
			NonTerminal:  newNonTerminal,
			Combinations: make([][]string, 0, len(recursiveCombinations)),
		}
		for _, recursiveCombination := range recursiveCombinations {
			newRule.Combinations = append(newRule.Combinations, slices.Clone(recursiveCombination[1:]))
			newRule.Combinations = append(newRule.Combinations, append(recursiveCombination[1:], newNonTerminal))
		}

		g.Rules = append(g.Rules, newRule)
		g.NonTerminals = append(g.NonTerminals, newNonTerminal)
	}

	return g
}
