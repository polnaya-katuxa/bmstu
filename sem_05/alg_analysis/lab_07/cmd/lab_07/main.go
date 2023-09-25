package main

import (
	"flag"
	"fmt"
	"lab_07/internal/aco"
	"lab_07/internal/bf"
	"lab_07/internal/graph"
	"log"
	"strconv"
)

var mode = flag.Bool("m", false, "mode: true - param, false - usual")
var alpha = flag.Float64("a", 0.0, "ph coeff")

// var beta = flag.Float64("b", 0.3, "tax coeff")
var k = flag.Float64("k", 0.0, "vaporize coeff")
var time = flag.Int("t", 0, "days num")

var n = flag.Int("f", 1, "file num in data/measure")

func main() {
	flag.Parse()
	nStr := strconv.Itoa(*n)
	dirName := "data/measure/"
	g := graph.Graph{}

	err := g.ReadGraphFromFile(dirName + nStr + ".txt")
	if err != nil {
		log.Fatal(err)
	}

	route, idealTax := bf.TravellingSalesmanBF(g)

	route, tax, err := aco.TravellingSalesmanACO(g, *alpha, 1-*alpha, *k, *time)
	if err != nil {
		log.Fatal(err)
	}

	if !(*mode) {
		fmt.Println(route, idealTax)
		fmt.Println(route, tax)
	} else {
		fmt.Printf("%.1f & %.1f & %.1f & %d & %d & %d \\\\ \\hline \n", *alpha, 1-*alpha, *k, *time, idealTax, tax-idealTax)
	}
}
