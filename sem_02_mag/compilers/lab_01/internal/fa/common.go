package fa

import (
	"fmt"
	"slices"
)

type State struct {
	State  []int
	Marked bool
	Last   bool
	Start  bool
}

func (s *State) Equal(in *State) bool {
	if s == nil || in == nil {
		return true
	} // todo

	slices.Sort(s.State)
	slices.Sort(in.State)

	if slices.Compare(s.State, in.State) != 0 {
		return false
	}

	return true
}

func MultiEpsClosure(s []*State) []*State {
	return slices.Clone(s)
}

func (s *State) String() string {
	slices.Sort(s.State)
	return fmt.Sprintf("%v", s.State)
}

func Union(s1 []*State, s2 []*State) []*State {
	u := make([]*State, 0, len(s1)+len(s2))
	u = append(u, s1...)

	m := make(map[string]*State, len(s1)+len(s2))
	for _, s := range s1 {
		m[s.String()] = s
	}

	for _, s := range s2 {
		if _, ok := m[s.String()]; !ok {
			u = append(u, s)
			m[s.String()] = s
		}
	}

	return u
}
