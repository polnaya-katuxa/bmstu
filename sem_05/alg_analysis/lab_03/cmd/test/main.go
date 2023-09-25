package main

import (
	"flag"
	"fmt"
	"lab_03/internal/algorithms"
	"log"
	"math/rand"
	"os"
	"testing"
	"time"
)

type algorithm struct {
	name     string
	function func([]int)
}

var algs = []algorithm{
	{"Pancakesort", algorithms.Pancakesort},
	{"Quicksort", algorithms.Quicksort},
	{"Beadsort", algorithms.Beadsort},
}

func generateArray(size int) []int {
	rand.Seed(time.Now().UnixNano())
	arr := make([]int, 0, size)
	for j := 0; j < size; j++ {
		arr = append(arr, rand.Intn(size)+1)
	}

	return arr
}

func generateAscArray(size int) []int {
	arr := make([]int, 0, size)
	for j := 0; j < size; j++ {
		arr = append(arr, j+1)
	}

	return arr
}

func generateFilledArray(size int, val int) []int {
	arr := make([]int, 0, size)
	for j := 0; j < size; j++ {
		arr = append(arr, val)
	}

	return arr
}

func flip(arr []int) {
	n := len(arr)

	for left := 0; left < n; left++ {
		arr[left], arr[n-1] = arr[n-1], arr[left]
		n--
	}
}

func generateDiffArrayRev(size int) []int {
	arr := make([]int, size)

	for j := range arr {
		arr[j] = j + 1
	}

	for n := 2; n <= size; n++ {
		flip(arr[:n])
		flip(arr[:n-1])
	}

	return arr
}

func NewBenchmark(arr []int, alg algorithm) func(*testing.B) {
	return func(b *testing.B) {
		copyA := make([]int, len(arr))

		b.ResetTimer()
		for j := 0; j < b.N; j++ {
			b.StopTimer()
			copy(copyA, arr)
			b.StartTimer()
			alg.function(copyA)
		}
	}
}

func main() {
	testing.Init()
	flag.Parse()

	lenghts := []int{1, 10, 50, 100, 200, 300, 500, 600, 800, 1000}

	arrsRandom := make([][]int, 0, len(lenghts))
	arrsAsc := make([][]int, 0, len(lenghts))
	arrsMax := make([][]int, 0, len(lenghts))
	arrsDiff := make([][]int, 0, len(lenghts))
	arrsOnes := make([][]int, 0, len(lenghts))

	for _, val := range lenghts {
		arrsRandom = append(arrsRandom, generateArray(val))
		arrsAsc = append(arrsAsc, generateAscArray(val))
		arrsMax = append(arrsMax, generateFilledArray(val, val))
		arrsDiff = append(arrsDiff, generateDiffArrayRev(val))
		arrsOnes = append(arrsOnes, generateFilledArray(val, 1))
	}

	timeFiles := make([]*os.File, len(algs))
	timeFilesBest := make([]*os.File, len(algs))
	timeFilesWorst := make([]*os.File, len(algs))
	memFiles := make([]*os.File, len(algs))

	for i := range algs {
		timeFile, err := os.Create(fmt.Sprintf("data/time%d.txt", i))
		if err != nil {
			log.Fatalln(err)
		}
		timeFiles[i] = timeFile
		defer timeFiles[i].Close()

		timeFileBest, err := os.Create(fmt.Sprintf("data/time_best%d.txt", i))
		if err != nil {
			log.Fatalln(err)
		}
		timeFilesBest[i] = timeFileBest
		defer timeFilesBest[i].Close()

		timeFileWorst, err := os.Create(fmt.Sprintf("data/time_worst%d.txt", i))
		if err != nil {
			log.Fatalln(err)
		}
		timeFilesWorst[i] = timeFileWorst
		defer timeFilesWorst[i].Close()

		memFile, err := os.Create(fmt.Sprintf("data/mem%d.txt", i))
		if err != nil {
			log.Fatalln(err)
		}
		memFiles[i] = memFile
		defer memFiles[i].Close()
	}

	for i := range arrsRandom {
		for j, alg := range algs {
			res := testing.Benchmark(NewBenchmark(arrsRandom[i], alg))
			fmt.Fprintf(timeFiles[j], "%d %d\n", lenghts[i], res.NsPerOp())
			fmt.Fprintf(memFiles[j], "%d %d\n", lenghts[i], algorithms.MemoryInfo.Max())

			if alg.name == "Pancakesort" {
				res = testing.Benchmark(NewBenchmark(arrsDiff[i], alg))
				fmt.Fprintf(timeFilesWorst[j], "%d %d\n", lenghts[i], res.NsPerOp())

				res = testing.Benchmark(NewBenchmark(arrsAsc[i], alg))
				fmt.Fprintf(timeFilesBest[j], "%d %d\n", lenghts[i], res.NsPerOp())
			} else if alg.name == "Quicksort" {
				res = testing.Benchmark(NewBenchmark(arrsMax[i], alg))
				fmt.Fprintf(timeFilesWorst[j], "%d %d\n", lenghts[i], res.NsPerOp())

				res = testing.Benchmark(NewBenchmark(arrsAsc[i], alg))
				fmt.Fprintf(timeFilesBest[j], "%d %d\n", lenghts[i], res.NsPerOp())
			} else {
				res = testing.Benchmark(NewBenchmark(arrsMax[i], alg))
				fmt.Fprintf(timeFilesWorst[j], "%d %d\n", lenghts[i], res.NsPerOp())

				res = testing.Benchmark(NewBenchmark(arrsOnes[i], alg))
				fmt.Fprintf(timeFilesBest[j], "%d %d\n", lenghts[i], res.NsPerOp())
			}
		}
	}
}
