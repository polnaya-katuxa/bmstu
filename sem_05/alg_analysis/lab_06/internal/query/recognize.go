package query

import (
	"errors"
	"lab_06_01/internal/dictionary"
	"lab_06_01/internal/models"
	"strings"
)

type Manager struct {
	dict dictionary.Dictionary
}

type Interval struct {
	Names []string
	Start int
	End   int
}

var Starters = []string{"показать", "показывать", "дать", "вывести", "выводить", "какой", "выдать", "хотеть", "найти"}
var Objects = []string{"кошка", "кот", "котёнок", "кошечка", "котик", "котёночек"}
var Intervals = []Interval{{
	Names: []string{"лысый", "не пушистый"},
	Start: 0,
	End:   417,
}, {
	Names: []string{"плешивый", "не особо пушистый", "не пушистый", "не лысый"},
	Start: 418,
	End:   1000,
}, {
	Names: []string{"волосатый", "не особо пушистый", "не пушистый", "не очень пушистый", "не лысый"},
	Start: 1001,
	End:   7500,
}, {
	Names: []string{"пушистый", "не лысый", "немного пушистый"},
	Start: 7501,
	End:   11667,
}, {
	Names: []string{"очень пушистый", "не лысый", "сильно пушистый"},
	Start: 11668,
	End:   17500,
}, {
	Names: []string{"невероятно пушистый", "не лысый"},
	Start: 17501,
	End:   20000,
}}

func New(dict dictionary.Dictionary) *Manager {
	return &Manager{dict}
}

func IsCorrectQuery(tokens []string) (int, int, error) {
	start := 0

	for _, v := range Starters {
		if tokens[0] == v {
			start = 1
			break
		}
	}

	subquery := strings.Join(tokens[start:], " ")

	start = -1
	length := 0
	for i, v := range Objects {
		if strings.Contains(subquery, v) {
			if len(v) > length {
				length = len(v)
				start = i
			}
		}
	}

	if start == -1 {
		return -1, -1, errors.New("неверный объект запроса")
	} else {
		subquery = strings.Replace(subquery, Objects[start], "", 1)
		subquery = strings.TrimSpace(subquery)
	}

	start = -1
	end := -1
	for _, v := range Intervals {
		for _, e := range v.Names {
			if subquery == e {
				if start == -1 {
					start = v.Start
				}
				if v.End > end {
					end = v.End
				}
			}
		}
	}

	if start == -1 || end == -1 {
		return -1, -1, errors.New("некорректный формат запроса")
	}

	return start, end, nil
}

func (m *Manager) Recognize(query string) ([]models.Cat, error) {
	tokens, err := Tokenize(query)
	if err != nil {
		return nil, err
	}

	normalized, err := Normalize(tokens)
	if err != nil {
		return nil, err
	}

	start, end, err := IsCorrectQuery(normalized)
	if err != nil {
		return nil, err
	}

	var cats []models.Cat
	cats = m.dict.Search(start, end)

	return cats, nil
}
