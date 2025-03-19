package fa

import (
	"fmt"
	"log/slog"

	"github.com/polnaya-katuxa/bmstu/sem_02_mag/compilers/lab_01/internal/ast"
	"github.com/polnaya-katuxa/bmstu/sem_02_mag/compilers/lab_01/internal/regexp"
)

const debugMode = "debug"

func BuildFaByRegexp(in *regexp.Regexp, mode string) error {
	slog.Info("start building AST by regexp", slog.String("regexp", in.Initial))

	ast, err := ast.NewAST(in)
	if err != nil {
		return fmt.Errorf("ast: %w", err)
	}
	err = ast.Show()
	if err != nil {
		return fmt.Errorf("show ast: %w", err)
	}

	slog.Info("start building DFA by AST", slog.String("symdols", string(ast.GetSymbols())))

	dfa := NewDFA(ast)
	err = dfa.Show("temp/dfa")
	if err != nil {
		return fmt.Errorf("show dfa: %w", err)
	}

	slog.Info("start reversing DFA")

	rDFA := NewNFA(dfa)
	if mode == debugMode {
		err = rDFA.Show("temp/nfa")
		if err != nil {
			return fmt.Errorf("show NFA: %w", err)
		}
	}

	slog.Info("start determining NFA")

	dNFA := rDFA.Determine(ast)
	if mode == debugMode {
		err = dNFA.Show("temp/dnfa")
		if err != nil {
			return fmt.Errorf("show dNFA: %w", err)
		}
	}

	slog.Info("start reversing DFA")

	rdNFA := NewNFA(dNFA)
	if mode == debugMode {
		err = rDFA.Show("temp/rdNFA")
		if err != nil {
			return fmt.Errorf("show rdNFA: %w", err)
		}
	}

	slog.Info("start determining NFA")

	drdNFA := rdNFA.Determine(ast)
	err = drdNFA.Show("temp/minimized")
	if err != nil {
		return fmt.Errorf("show drdNFA: %w", err)
	}

	err = drdNFA.Save()
	if err != nil {
		return fmt.Errorf("save drdNFA: %w", err)
	}

	slog.Info("finished minimizing dfa by brzhzovsky algorythm")

	return nil
}
