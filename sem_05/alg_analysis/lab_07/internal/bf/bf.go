package bf

import (
	"lab_07/internal/graph"
	"math"
)

func permutations(arr []int) [][]int {
	var mixer func([]int, int)
	var mix [][]int

	mixer = func(arr []int, n int) {
		if n == 1 {
			tmp := make([]int, len(arr))
			copy(tmp, arr)
			mix = append(mix, tmp)
		} else {
			for i := 0; i < n; i++ {
				mixer(arr, n-1)
				if n%2 == 1 {
					tmp := arr[i]
					arr[i] = arr[n-1]
					arr[n-1] = tmp
				} else {
					tmp := arr[0]
					arr[0] = arr[n-1]
					arr[n-1] = tmp
				}
			}
		}
	}

	mixer(arr, len(arr))

	return mix
}

func TravellingSalesmanBF(g graph.Graph) ([]int, int) {
	var minRoute []int

	if g.Size == 1 {
		return []int{}, 0
	}

	verts := make([]int, g.Size)
	for i := range verts {
		verts[i] = i
	}

	permuts := permutations(verts)
	minTax := math.MaxInt

	for i := range permuts {
		if g.IsOKRoute(permuts[i]) {
			cost := g.RouteTotalTax(permuts[i])

			if cost < minTax {
				minTax = cost
				minRoute = permuts[i]
			}
		}
	}

	//fmt.Println(minRoute, minTax)

	return minRoute, minTax
}
