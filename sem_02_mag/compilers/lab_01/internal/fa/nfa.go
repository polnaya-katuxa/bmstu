package fa

import (
	"fmt"
	"log/slog"
	"slices"

	"github.com/polnaya-katuxa/bmstu/sem_02_mag/compilers/lab_01/internal/ast"
)

type NFA struct {
	States []*State
	Tran   map[string]map[rune][]*State
}

func (n *NFA) GetStartState() *State {
	for _, s := range n.States {
		if s.Start {
			return s
		}
	}

	return nil
}

func (n *NFA) GetByIndexes(s *State) []*State {
	states := make([]*State, 0, len(s.State))
	for _, v := range s.State {
		states = append(states, n.States[v])
	}

	return states
}

func (n *NFA) Move(states []*State, a rune) []*State {
	u := make([]*State, 0, len(states))
	for _, in := range states {
		tr, ok := n.Tran[in.String()]
		if !ok {
			continue
		}

		end, ok := tr[a]
		if !ok {
			continue
		}

		u = Union(u, end)
	}

	return u
}

func (n *NFA) EpsClosure(s []*State) []int {
	closure := make([]int, 0, len(s))
	for _, in := range s {
		for j, st := range n.States {
			if in.Equal(st) {
				closure = append(closure, j)
				break
			}
		}
	}

	return closure
}

func NewNFA(d *DFA) *NFA {
	nfa := &NFA{
		States: make([]*State, 0, len(d.States)),
		Tran:   make(map[string]map[rune][]*State, len(d.Tran)),
	}

	slog.Info("start reversing DFA")

	for _, s := range d.States {
		last := false
		start := false

		if s.Last {
			start = true
		}
		if s.Start {
			last = true
		}
		state := &State{
			State:  slices.Clone(s.State),
			Marked: s.Marked,
			Last:   last,
			Start:  start,
		}

		nfa.States = append(nfa.States, state)
	}

	slog.Info("got NFA states", slog.Int("count", len(nfa.States)))

	for stateStart, tran := range d.Tran {
		var state *State
		for _, s := range nfa.States {
			if stateStart == s.String() {
				state = &State{
					State:  slices.Clone(s.State),
					Marked: s.Marked,
					Last:   s.Last,
					Start:  s.Start,
				}
				break
			}
		}

		if state == nil {
			continue
		}

		for sym, stateEnd := range tran {
			tr, ok := nfa.Tran[stateEnd.String()]
			if !ok {
				tr = map[rune][]*State{sym: []*State{state}}
			} else {
				tr[sym] = append(tr[sym], state)
			}
			nfa.Tran[stateEnd.String()] = tr
		}
	}

	slog.Info("got NFA transitions", slog.Int("count", len(nfa.Tran)))

	// fmt.Println("REV")
	// for i := range nfa.States {
	// 	fmt.Printf("%#v\n", nfa.States[i])
	// }
	// fmt.Println()
	// for k, v := range nfa.Tran {
	// 	for k1, v1 := range v {
	// 		fmt.Printf("%s %c %#v\n", k, k1, v1)
	// 	}
	// }

	return nfa
}

func (n *NFA) Determine(ast *ast.AST) *DFA {
	symbolsList := ast.GetSymbols()
	dfa := &DFA{
		States: make([]*State, 0, len(n.States)),
		Tran:   make(map[string]map[rune]*State, len(n.Tran)),
	}

	fmt.Println("NFA")
	for i := range n.States {
		fmt.Printf("%#v\n", n.States[i])
	}
	fmt.Println()
	for k, v := range n.Tran {
		for k1, v1 := range v {
			fmt.Printf("%s %c %#v\n", k, k1, v1)
		}
	}

	slog.Info("start determine NFA")

	s0EpcClosure := n.EpsClosure([]*State{n.GetStartState()})
	first := &State{
		State:  s0EpcClosure,
		Marked: false,
		Last:   false,
		Start:  true,
	}
	dfa.States = append(dfa.States, first)

	slog.Info("got first state (indexes of NFA states)", slog.String("state", first.String()))

	nStates := n.GetByIndexes(first)
	for _, s := range nStates {
		if s.Last {
			first.Last = true
			break
		}
	}

	slog.Info("starting unmarked states cycle")

	for {
		var s *State
		for i := range dfa.States {
			if !dfa.States[i].Marked {
				s = dfa.States[i]
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
			u := n.EpsClosure(n.Move(n.GetByIndexes(s), a))
			if len(u) == 0 {
				continue
			}

			slog.Info("calculated union eps-closure", slog.String("ec", fmt.Sprintf("%v", u)))

			contains := false
			for _, v := range dfa.States {
				slices.Sort(v.State)
				slices.Sort(u)
				if slices.Equal(v.State, u) {
					contains = true
					break
				}
			}
			uState := &State{
				State:  u,
				Marked: false,
				Last:   false,
				Start:  false,
			}

			nStates := n.GetByIndexes(uState)
			for _, s := range nStates {
				if s.Last {
					uState.Last = true
					break
				}
			}

			if !contains {
				slog.Info("adding new state", slog.String("state", uState.String()))
				dfa.States = append(dfa.States, uState)
			}

			tr, ok := dfa.Tran[s.String()]
			if !ok {
				tr = map[rune]*State{a: uState}
			} else {
				tr[a] = uState
			}
			dfa.Tran[s.String()] = tr

			slog.Info("adding new tran", slog.String("from", s.String()), slog.String("by", string(a)), slog.String("to", uState.String()))
		}
	}

	return dfa
}
