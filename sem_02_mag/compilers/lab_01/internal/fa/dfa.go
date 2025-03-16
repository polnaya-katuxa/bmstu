package fa

import (
	"github.com/polnaya-katuxa/bmstu/sem_02_mag/compilers/lab_01/internal/ast"
	"k8s.io/utils/set"
)

type DFA struct {
	States []*State
	Tran   map[string]map[rune]*State
}

func NewDFA(ast *ast.AST) *DFA {
	return buildDFA(ast)
}

func buildDFA(ast *ast.AST) *DFA {
	root := ast.GetRoot()
	first := &State{State: root.GetFirstPos(), Start: true}
	states := []*State{first}
	tran := make(map[string]map[rune]*State, 0)
	symbolsList := ast.GetSymbols()
	symbolsMap := ast.GetSymbolsMap()

	for _, st := range first.State {
		v, ok := symbolsMap[st]
		if !ok {
			panic("symbols incorrect")
		}

		if v.Value == '#' {
			first.Last = true
			break
		}
	}

	for {
		var s *State
		for i := range states {
			if !states[i].Marked {
				s = states[i]
				break
			}
		}

		if s == nil {
			break
		}

		s.Marked = true
		for _, a := range symbolsList {
			u := set.New[int]()
			for _, p := range s.State {
				v, ok := symbolsMap[p]
				if !ok {
					panic("symbols incorrect")
				}

				if v.Value != a {
					continue
				}

				u = u.Union(set.New(v.GetFollowPos()...))
			}

			if u.Len() == 0 {
				continue
			}

			state := &State{State: u.SortedList()}
			for _, st := range state.State {
				v, ok := symbolsMap[st]
				if !ok {
					panic("symbols incorrect")
				}

				if v.Value == '#' {
					state.Last = true
					break
				}
			}

			contains := false
			for _, v := range states {
				if v.Equal(state) {
					contains = true
					break
				}
			}
			if !contains {
				states = append(states, state)
			}

			tr, ok := tran[s.String()]
			if !ok {
				tr = map[rune]*State{a: state}
			} else {
				tr[a] = state
			}
			tran[s.String()] = tr
		}
	}

	// fmt.Println("NORM")
	// for i := range states {
	// 	fmt.Printf("%#v\n", states[i])
	// }
	// fmt.Println()
	// for k, v := range tran {
	// 	for k1, v1 := range v {
	// 		fmt.Printf("%s %c %#v\n", k, k1, v1)
	// 	}
	// }

	return &DFA{
		States: states,
		Tran:   tran,
	}
}

func (d *DFA) GetStart() *State {
	for _, s := range d.States {
		if s.Start {
			return s
		}
	}

	return nil
}

func (d *DFA) Model(in string) bool {
	curState := d.GetStart()
	for _, s := range in {
		tr, ok := d.Tran[curState.String()]
		if !ok {
			return false
		}

		next, ok := tr[s]
		if !ok {
			return false
		}

		curState = next
	}

	if !curState.Last {
		return false
	}

	return true
}
