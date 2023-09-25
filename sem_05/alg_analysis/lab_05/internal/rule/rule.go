package rule

import (
	"encoding/json"
	"errors"
	"os"
	"regexp"
)

type Rule struct {
	Option   []string
	Standard string
}

func isEqual(st string, opt []string) bool {
	re := regexp.MustCompile(`([^а-яА-ЯЁё_-]|[+])`)
	splitted := re.Split(st, -1)
	i := 0
	for _, v := range splitted {
		if v == "" {
			continue
		}
		if v != opt[i] {
			return false
		}
		i++
	}

	return true
}

func ReadRules(configFile *os.File) ([]Rule, error) {
	rules := make([]Rule, 0)
	err := json.NewDecoder(configFile).Decode(&rules)
	if err != nil {
		return nil, err
	}

	length := len(rules)
	for i := 0; i < length; i++ {
		isOption := false
		for j := 0; j < length; j++ {
			if isEqual(rules[i].Standard, rules[j].Option) {
				if rules[i].Standard != rules[j].Standard {
					return nil, errors.New("not turned to itself")
				}
				isOption = true
			}
		}

		if !isOption {
			return nil, errors.New("not an option")
		}
	}

	return rules, nil
}
