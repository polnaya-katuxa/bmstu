package document

import (
	"errors"
	"fmt"
	"lab_05/internal/rule"
	"os"
	"path"
	"regexp"
	"sort"
	"strings"
)

type Document struct {
	Name   string
	Text   string
	Tokens []string
}

func InputDocs(dirName string, num int) ([]Document, error) {
	dirEntries, err := os.ReadDir(dirName)
	if err != nil {
		return nil, err
	}

	docs := make([]Document, 0)
	for i, doc := range dirEntries {
		if i >= num {
			break
		}

		text, err := os.ReadFile(path.Join(dirName, doc.Name()))
		if err != nil {
			return nil, err
		}
		docs = append(docs, Document{
			Name:   dirEntries[i].Name(),
			Text:   string(text),
			Tokens: nil,
		})
	}

	return docs, nil
}

func (doc *Document) PrintTokens() {
	fmt.Printf("\nTokens of document \"%s\":\n", doc.Name)
	for i, v := range doc.Tokens {
		if i%5 == 0 {
			fmt.Println()
		}
		if i == len(doc.Tokens)-1 {
			fmt.Printf("%s\n", v)
		} else {
			fmt.Printf("%s, ", v)
		}
	}
	fmt.Println()
}

func (doc *Document) Tokenize() {
	doc.Tokens = nil
	re := regexp.MustCompile(`([^0-9а-яА-ЯЁё-])`) //|[+]
	splitted := re.Split(doc.Text, -1)
	for i := range splitted {
		if splitted[i] != "" && splitted[i] != "-" {
			doc.Tokens = append(doc.Tokens, strings.ToLower(splitted[i]))
		}
	}
}

func (doc *Document) ApplyRules(rules []rule.Rule) error {
	if len(doc.Tokens) == 0 {
		return errors.New("no tokens")
	}
	if len(rules) == 0 {
		return nil
	}

	ruled := make([]string, 0)

	for i := 0; i < len(doc.Tokens); i++ {
		ruleInd := -1

		for j := 0; j < len(rules) && ruleInd < 0; j++ {
			ruleInd = j

			for k := 0; k < len(rules[j].Option) && ruleInd >= 0; k++ {
				if i+k >= len(doc.Tokens) || doc.Tokens[i+k] != rules[j].Option[k] {
					ruleInd = -1
				}
			}
		}

		if ruleInd < 0 {
			ruled = append(ruled, doc.Tokens[i])
		} else {
			ruled = append(ruled, rules[ruleInd].Standard)
			i += len(rules[ruleInd].Option) - 1
		}
	}

	doc.Tokens = ruled

	return nil
}

func (doc *Document) Process(rules []rule.Rule) error {
	doc.Tokenize()

	err := doc.ApplyRules(rules)
	if err != nil {
		return err
	}

	sort.Strings(doc.Tokens)

	return nil
}
