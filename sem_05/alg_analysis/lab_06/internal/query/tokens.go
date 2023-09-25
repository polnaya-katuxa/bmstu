package query

import (
	"errors"
	"gitlab.com/opennota/morph"
	"regexp"
	"strings"
)

func init() {
	if err := morph.Init(); err != nil {
		panic(err)
	}
}

func Tokenize(text string) ([]string, error) {
	tokens := make([]string, 0)

	re := regexp.MustCompile(`([^0-9а-яА-ЯЁё-])`)
	splitted := re.Split(text, -1)

	for i := range splitted {
		if splitted[i] != "" && splitted[i] != "-" {
			tokens = append(tokens, strings.ToLower(splitted[i]))
		}
	}

	if len(tokens) == 0 {
		return []string{}, errors.New("невалидный запрос")
	}

	return tokens, nil
}

func Normalize(tokens []string) ([]string, error) {
	terms := make([]string, 0)

	for _, v := range tokens {
		_, normals, _ := morph.Parse(strings.ToLower(v))
		if len(normals) == 0 {
			return nil, errors.New("незнакомое слово")
		}

		terms = append(terms, normals[0])
	}

	return terms, nil
}
