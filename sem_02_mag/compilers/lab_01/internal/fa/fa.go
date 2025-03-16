package fa

import (
	"fmt"

	"github.com/polnaya-katuxa/bmstu/sem_02_mag/compilers/lab_01/internal/ast"
	"github.com/polnaya-katuxa/bmstu/sem_02_mag/compilers/lab_01/internal/regexp"
)

const debugMode = "debug"

func BuildFaByRegexp(in *regexp.Regexp, mode string) error {
	ast, err := ast.NewAST(in)
	if err != nil {
		return fmt.Errorf("ast: %w", err)
	}
	err = ast.Show()
	if err != nil {
		return fmt.Errorf("show ast: %w", err)
	}

	dfa := NewDFA(ast)
	err = dfa.Show("temp/dfa")
	if err != nil {
		return fmt.Errorf("show dfa: %w", err)
	}

	rDFA := NewNFA(dfa)
	if mode == debugMode {
		err = rDFA.Show("temp/nfa")
		if err != nil {
			return fmt.Errorf("show NFA: %w", err)
		}
	}

	dNFA := rDFA.Determine(ast)
	if mode == debugMode {
		err = dNFA.Show("temp/dnfa")
		if err != nil {
			return fmt.Errorf("show dNFA: %w", err)
		}
	}

	rdNFA := NewNFA(dNFA)
	if mode == debugMode {
		err = rDFA.Show("temp/rdNFA")
		if err != nil {
			return fmt.Errorf("show rdNFA: %w", err)
		}
	}

	drdNFA := rdNFA.Determine(ast)
	err = drdNFA.Show("temp/minimized")
	if err != nil {
		return fmt.Errorf("show drdNFA: %w", err)
	}

	err = drdNFA.Save()
	if err != nil {
		return fmt.Errorf("save drdNFA: %w", err)
	}

	return nil
}
