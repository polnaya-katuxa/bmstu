package fa

import (
	"fmt"
	"log/slog"
	"slices"

	"github.com/polnaya-katuxa/bmstu/sem_02_mag/compilers/lab_01/internal/ast"
	"golang.org/x/exp/maps"
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

	slog.Info("marking if last", slog.String("state", first.String()))

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

	slog.Info("starting unmarked states cycle")

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

		slog.Info("found unmarked", slog.String("unmarked", s.String()))

		slices.Sort(s.State)
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

			slog.Info("got follow pos union, marking if last", slog.String("union", state.String()))

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
				slog.Info("adding new state", slog.String("state", state.String()))
				states = append(states, state)
			}

			tr, ok := tran[s.String()]
			if !ok {
				tr = map[rune]*State{a: state}
			} else {
				tr[a] = state
			}
			tran[s.String()] = tr

			slog.Info("adding new tran", slog.String("from", s.String()), slog.String("by", string(a)), slog.String("to", state.String()))
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

func (d *DFA) GetLast() *State {
	for _, s := range d.States {
		if s.Last {
			return s
		}
	}

	return nil
}

func (d *DFA) Model(in string) bool {
	fmt.Println(commentC, "START MODELING FOR", in, endC)

	curState := d.GetStart()

	way := &way{
		steps: []*step{
			{
				symbol: "",
				dst:    curState.String(),
				border: true,
			},
		},
	}

	for _, s := range in {
		slog.Info("current state", slog.String("state", curState.String()))

		tr, ok := d.Tran[curState.String()]
		if !ok {
			slog.Error("can't find transition starting with such state", slog.Any("valid options", maps.Keys(d.Tran)))
			way.Show()
			return false
		}

		slog.Info("found transition with such start state", slog.String("from", curState.String()))

		next, ok := tr[s]
		if !ok {
			slog.Error("can't find transition by such symbol", slog.String("symbol", string(s)),
				slog.Any("valid options", maps.Keys(tr)))
			way.Show()
			return false
		}

		slog.Info("found transition by such symbol", slog.String("from", curState.String()),
			slog.String("by", string(s)), slog.String("to", next.String()))

		way.steps = append(way.steps, &step{
			symbol: string(s),
			dst:    next.String(),
		})

		curState = next
	}

	if !curState.Last {
		slog.Error("end state isn't last", slog.String("state", curState.String()),
			slog.String("last", d.GetLast().String()))
		way.Show()
		return false
	}

	way.steps[len(way.steps)-1].border = true
	way.Show()

	return true
}
