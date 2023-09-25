package graph

import (
	"fmt"
	"os"
)

type Graph struct {
	Size       int
	Connection [][]int
}

func (g *Graph) ReadGraphFromFile(fileName string) error {
	file, err := os.Open(fileName)
	if err != nil {
		return err
	}

	var size int
	n, err := fmt.Fscan(file, &size)
	if err != nil || n != 1 {
		return err
	}

	conn := make([][]int, size)
	for i := 0; i < size; i++ {
		connRow := make([]int, size)
		conn[i] = connRow
	}

	var weight int
	for i := 0; i < size; i++ {
		for j := 0; j < size; j++ {
			n, err = fmt.Fscan(file, &weight)
			if err != nil || n != 1 {
				return err
			}
			conn[i][j] = weight
		}
	}

	g.Size = size
	g.Connection = conn

	return nil
}

func (g *Graph) Print() {
	fmt.Println("\nGRAPH:")
	for i := 0; i < g.Size; i++ {
		for j := 0; j < g.Size; j++ {
			fmt.Printf("%d ", g.Connection[i][j])
		}
		fmt.Println()
	}
}

func (g *Graph) IsOKRoute(route []int) bool {
	if g.Size == 1 {
		return true
	}

	for i := range route {
		if g.Connection[route[i]][route[(i+1)%g.Size]] == 0 {
			return false
		}
	}

	return true
}

func (g *Graph) RouteTotalTax(path []int) int {
	cost := 0

	for i := range path {
		cost += g.Connection[path[i]][path[(i+1)%g.Size]]
	}

	return cost
}
