package aco

import (
	"errors"
	"lab_07/internal/graph"
	"math"
	"math/rand"
	"time"
)

const (
	EPS = 1e-8
)

type Ant struct {
	IsElite bool
	Route   []int
	Tax     int
	Pos     int
}

func (a *Ant) Move(tax int, pos int) {
	a.Route = append(a.Route, pos)
	a.Tax += tax
	a.Pos = pos
}

func (a *Ant) Go(toGo []int, ph [][]float64, vis [][]float64, alpha, beta float64) (int, error) {
	probs := make([]float64, 0)
	sumProbs := 0.0

	for _, v := range toGo {
		if vis[a.Pos][v] != -1 {
			greed := math.Pow(vis[a.Pos][v], beta)
			herd := math.Pow(ph[a.Pos][v], alpha)
			prob := greed * herd

			probs = append(probs, prob)
			sumProbs += prob
		}
	}

	if sumProbs <= 0 {
		return -1, errors.New("probability error")
	}

	maxProb := 0.0
	greedChoice := 0
	for i, v := range probs {
		v /= sumProbs
		if v > maxProb {
			maxProb = v
			greedChoice = toGo[i]
		}
	}

	if a.IsElite {
		return greedChoice, nil
	}

	choice := 0
	curSum := 0.0
	rand.Seed(time.Now().UnixNano())
	randPoint := rand.Float64() * sumProbs
	for curSum < randPoint {
		curSum += probs[choice]
		choice++
	}

	return toGo[choice-1], nil
}

func (a *Ant) FindRoute(g graph.Graph, ph [][]float64, vis [][]float64, alpha, beta float64) error {
	toGo := make([]int, 0)

	for i := range g.Connection {
		if i != a.Pos {
			toGo = append(toGo, i)
		}
	}

	cycle := false
	for len(toGo) != 0 {
		next, err := a.Go(toGo, ph, vis, alpha, beta)
		if err != nil {
			return err
		}

		a.Move(g.Connection[a.Pos][next], next)
		toGo = removeElement(toGo, next)

		if len(toGo) == 0 && !cycle {
			toGo = append(toGo, a.Route[0])
			cycle = true
		}
	}

	return nil
}

func removeElement(arr []int, el int) []int {
	idx := -1

	for i, v := range arr {
		if v == el {
			idx = i
			break
		}
	}

	if idx == -1 {
		return arr
	}

	return append(arr[:idx], arr[idx+1:]...)
}
