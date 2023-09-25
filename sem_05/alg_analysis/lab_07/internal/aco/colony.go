package aco

import (
	"errors"
	"fmt"
	"lab_07/internal/graph"
	"math"
	"math/rand"
	"time"
)

const (
	tau0 = 0.1
)

type Colony struct {
	Size    int
	Members []Ant
}

func createPheromoneMatrix(sizeGraph int) [][]float64 {
	phMatrix := make([][]float64, sizeGraph)
	for i := 0; i < sizeGraph; i++ {
		phRow := make([]float64, sizeGraph)
		phMatrix[i] = phRow
	}

	for i := 0; i < sizeGraph; i++ {
		for j := 0; j < sizeGraph; j++ {
			phMatrix[i][j] = tau0
		}
	}

	return phMatrix
}

func createVisionMatrix(g graph.Graph) [][]float64 {
	visMatrix := make([][]float64, g.Size)
	for i := 0; i < g.Size; i++ {
		visRow := make([]float64, g.Size)
		visMatrix[i] = visRow
	}

	for i := 0; i < g.Size; i++ {
		for j := 0; j < g.Size; j++ {
			if g.Connection[i][j] > 0 {
				visMatrix[i][j] = 1 / float64(g.Connection[i][j])
			} else {
				visMatrix[i][j] = -1
			}
		}
	}

	return visMatrix
}

func createColony(sizeColony int) Colony {
	members := make([]Ant, sizeColony)

	for i := 0; i < sizeColony; i++ {
		route := make([]int, 0)
		route = append(route, i)

		rand.Seed(time.Now().UnixNano())
		r := rand.Float64()
		isElite := r < 0.2

		members[i] = Ant{
			IsElite: isElite,
			Route:   route,
			Tax:     0,
			Pos:     i,
		}
	}

	return Colony{
		Size:    sizeColony,
		Members: members,
	}
}

func vaporize(p [][]float64, k float64) {
	for i := range p {
		for j := range p[i] {
			p[i][j] = (1 - k) * p[i][j]
		}
	}
}

func increase(p [][]float64, c Colony, q float64) {
	for _, v := range c.Members {
		for j := 0; j < len(v.Route)-1; j++ {
			v1 := v.Route[j]
			v2 := v.Route[j+1]
			p[v1][v2] += q / float64(v.Tax)
		}
	}
}

func correction(p [][]float64) {
	for i := range p {
		for j := range p[i] {
			if p[i][j] < tau0 {
				p[i][j] = tau0
			}
		}
	}
}

func getQ(g graph.Graph) float64 {
	q := 0.0

	for i := 0; i < g.Size; i++ {
		for j := 0; j < g.Size; j++ {
			q += float64(g.Connection[i][j])
		}
	}

	q /= float64(g.Size)

	return q
}

func TravellingSalesmanACO(g graph.Graph, alpha, beta, k float64, time int) ([]int, int, error) { //q float64
	pheromone := createPheromoneMatrix(g.Size)
	vision := createVisionMatrix(g)
	minRoute := make([]int, g.Size)
	minTax := math.MaxInt
	q := getQ(g)

	for t := 0; t < time; t++ {
		c := createColony(g.Size)

		for _, v := range c.Members {
			err := v.FindRoute(g, pheromone, vision, alpha, beta)
			if err != nil {
				return nil, -1, fmt.Errorf("route error: %w", err)
			}

			if v.Tax < minTax && (len(v.Route) == g.Size+1 || g.Size == 1) {
				minTax = v.Tax
				minRoute = v.Route
			}
		}

		vaporize(pheromone, k)
		increase(pheromone, c, q)
		correction(pheromone)
	}

	if len(minRoute) == 0 {
		return nil, -1, errors.New("min route error")
	}

	//fmt.Println(minRoute[:len(minRoute)-1], minTax)

	return minRoute[:len(minRoute)-1], minTax, nil
}
