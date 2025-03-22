package regexp

import (
	"fmt"
	"unicode"

	"github.com/polnaya-katuxa/bmstu/sem_02_mag/compilers/lab_01/internal/common"
)

type Regexp struct {
	Initial string
}

func (r *Regexp) GetInitial() string {
	return r.Initial
}

func (r *Regexp) Tokenize() ([]common.Token, error) {
	tokens := make([]common.Token, 0, len(r.Initial))

	for _, cur := range r.Initial {
		switch {
		case cur == '(':
			tokens = append(tokens, common.Token{common.LParen, cur})
		case cur == ')':
			tokens = append(tokens, common.Token{common.RParen, cur})
		case cur == '*':
			tokens = append(tokens, common.Token{common.KleeneStar, cur})
		case cur == '|':
			tokens = append(tokens, common.Token{common.Pipe, cur})
		case cur == '#':
			tokens = append(tokens, common.Token{common.Symbol, cur})
		case unicode.IsLetter(cur) || unicode.IsDigit(cur):
			tokens = append(tokens, common.Token{common.Symbol, cur})
		default:
			return nil, fmt.Errorf("invalid sequence: %s", string(cur))
		}
	}

	tokens = append(tokens, common.Token{common.EOF, 0})

	return tokens, nil
}
