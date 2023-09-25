package main

/*
#include <pthread.h>
#include <time.h>
#include <stdio.h>

static long long getCPUNs(){
	struct timespec time;
	if (clock_gettime(CLOCK_PROCESS_CPUTIME_ID, &time)) {
		perror("can't measure time");
		return 0;
	}
	return time.tv_sec * 1000000000LL + time.tv_nsec;
}
*/
import "C"
import (
	"flag"
	"fmt"
	"lab_02/internal/algorithms"
	"lab_02/internal/matrix"
	"lab_02/internal/memory"
	"log"
	"os"
	"testing"
)

const (
	NumTests = 100
)

type algorithm struct {
	name     string
	function func(matrix.Matrix, matrix.Matrix) (matrix.Matrix, error)
}

var algs = []algorithm{
	{"Usual", matrix.Mul},
	{"Winograd", algorithms.WinogradMulMatrix},
	{"Winograd Improved", algorithms.WinogradBetterMulMatrix},
}

func RunBenchmark(m1, m2 matrix.Matrix, alg algorithm) int {
	t := 0

	start := C.getCPUNs()
	for i := 0; i < NumTests; i++ {
		alg.function(m1, m2)
	}
	end := C.getCPUNs()
	t = int(end-start) / NumTests

	return t
}

func NewBenchmark(m1, m2 matrix.Matrix, alg algorithm) func(*testing.B) {
	return func(b *testing.B) {
		for j := 0; j < b.N; j++ {
			alg.function(m1, m2)
		}
	}
}

func main() {
	testing.Init()
	flag.Parse()

	lenghts := []int{2, 10, 50, 100, 200, 300, 500}

	matrEven := make([]matrix.Matrix, 0, len(lenghts))
	matrOdd := make([]matrix.Matrix, 0, len(lenghts))

	for _, val := range lenghts {
		m1 := matrix.CreateEmpty(val, val)
		matrix.FillRandom(m1)
		matrEven = append(matrEven, m1)

		m2 := matrix.CreateEmpty(val+1, val+1)
		matrix.FillRandom(m2)
		matrOdd = append(matrOdd, m2)
	}

	//timeFiles := make([]*os.File, len(algs))
	timeFilesBest := make([]*os.File, len(algs))
	timeFilesWorst := make([]*os.File, len(algs))
	memFiles := make([]*os.File, len(algs))

	for i := range algs {
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

	for i := range matrEven {
		for j, alg := range algs {
			res := RunBenchmark(matrEven[i], matrEven[i], alg)
			fmt.Fprintf(timeFilesBest[j], "%d %d\n", lenghts[i], res)
			fmt.Fprintf(memFiles[j], "%d %d\n", lenghts[i], memory.MemoryInfo.Max())

			res = RunBenchmark(matrOdd[i], matrOdd[i], alg)
			fmt.Fprintf(timeFilesWorst[j], "%d %d\n", lenghts[i], res)
			//fmt.Fprintf(memFiles[j], "%d %d\n", lenghts[i], memory.MemoryInfo.Max())
		}
	}
}
