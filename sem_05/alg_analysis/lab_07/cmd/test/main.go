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
	"fmt"
	"lab_07/internal/aco"
	"lab_07/internal/bf"
	"lab_07/internal/graph"
	"log"
	"strconv"
)

const NumTests = 100

func RunBenchmarkACO(g graph.Graph, alpha, beta, k float64, time int) int {
	t := 0

	start := C.getCPUNs()
	for i := 0; i < NumTests; i++ {
		aco.TravellingSalesmanACO(g, alpha, beta, k, time)
	}
	end := C.getCPUNs()
	t = int(end-start) / NumTests

	return t
}

func RunBenchmarkBF(g graph.Graph) int {
	t := 0

	start := C.getCPUNs()
	for i := 0; i < NumTests; i++ {
		bf.TravellingSalesmanBF(g)
	}
	end := C.getCPUNs()
	t = int(end-start) / NumTests

	return t
}

func main() {
	dirName := "data/measure/"

	docNum := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	for _, v := range docNum {
		g := graph.Graph{}

		fileName := dirName + strconv.Itoa(v) + ".txt"
		err := g.ReadGraphFromFile(fileName)
		if err != nil {
			log.Fatal(err)
		}

		resBF := RunBenchmarkBF(g)
		resACO := RunBenchmarkACO(g, 0.5, 0.5, 0.5, 10)
		fmt.Printf("n = %d, BF = %d, ACO = %d\n", v, resBF/1000, resACO/1000)
	}
}
