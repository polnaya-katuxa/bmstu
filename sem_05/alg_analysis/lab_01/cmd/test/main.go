package main

import (
	"aa/internal/algorithms"
	"fmt"
	"log"
	"math/rand"
	"os"
	"testing"
)

type algorithm struct {
	name     string
	function func(string, string) int
}

var algs = []algorithm{
	{"Levenshtein", algorithms.Levenshtein},
	{"Damerau-Levenshtein", algorithms.DamerauLevenshtein},
	{"Recursive Damerau-Levenshtein", algorithms.RecursiveDamerauLevenshtein},
	{"Recursive Damerau-Levenshtein with cache", algorithms.RecursiveDamerauLevenshteinCached},
}

func generateString(size int) string {
	var symbols = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

	var str = ""
	for j := 0; j < size; j++ {
		str += string(symbols[rand.Int()%len(symbols)])
	}

	return str
}

func NewBenchmark(s1, s2 string, alg algorithm) func(*testing.B) {
	return func(b *testing.B) {
		for j := 0; j < b.N; j++ {
			alg.function(s1, s2)
		}
	}
}

func main() {
	lenghts := []int{1, 2, 3, 5, 6, 7, 10, 15, 20, 30, 50}
	strings1 := make([]string, 0, len(lenghts))
	strings2 := make([]string, 0, len(lenghts))

	for _, val := range lenghts {
		strings1 = append(strings1, generateString(val))
		strings2 = append(strings2, generateString(val))
	}

	timeFiles := make([]*os.File, len(algs))
	memFiles := make([]*os.File, len(algs))

	for i := range algs {
		timeFile, err := os.Create(fmt.Sprintf("data/time%d.txt", i))
		if err != nil {
			log.Fatalln(err)
		}
		timeFiles[i] = timeFile
		defer timeFiles[i].Close()

		memFile, err := os.Create(fmt.Sprintf("data/mem%d.txt", i))
		if err != nil {
			log.Fatalln(err)
		}
		memFiles[i] = memFile
		defer memFiles[i].Close()
	}

	for i := range strings1 {
		for j, alg := range algs {
			if j == 2 && len(strings1[i]) > 10 {
				continue
			}

			res := testing.Benchmark(NewBenchmark(strings1[i], strings2[i], alg))
			fmt.Fprintf(timeFiles[j], "%d %d\n", len(strings1[i]), res.NsPerOp())
			fmt.Fprintf(memFiles[j], "%d %d\n", len(strings1[i]), algorithms.MemoryInfo.Max())
		}
	}
}
