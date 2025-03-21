package regexp

import (
	"fmt"
	"io"
	"os"
	"unicode"

	"github.com/polnaya-katuxa/bmstu/sem_02_mag/compilers/lab_01/internal/common"
)

func NewRegexp(regexp string) error {
	if err := ValidateRegexp(regexp); err != nil {
		return common.ErrInvalidRegexp
	}

	f, err := os.Create("temp/regexp.txt")
	defer f.Close()
	if err != nil {
		return fmt.Errorf("open file: %w", err)
	}

	regexp = regexp + "#"
	_, err = f.WriteString(regexp)
	if err != nil {
		return fmt.Errorf("write to file: %w", err)
	}

	return nil
}

func CheckRegexp() (*Regexp, error) {
	f, err := os.OpenFile("temp/regexp.txt", os.O_RDONLY, 666)
	defer f.Close()
	if err != nil {
		return nil, fmt.Errorf("open file: %w", err)
	}

	data, err := io.ReadAll(f)
	if err != nil {
		return nil, fmt.Errorf("read from file: %w", err)
	}

	regexp := string(data)
	if err := ValidateRegexp(regexp); err != nil {
		return nil, common.ErrInvalidRegexp
	}

	return &Regexp{
		Initial: regexp,
	}, nil
}

func ValidateRegexp(regexp string) error {
	return nil
}

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
