package main

import (
	"fmt"
	"lab_05/internal/document"
	"lab_05/internal/pipeline"
	"lab_05/internal/rule"
	"log"
	"os"
	"testing"
)

func NewBenchmarkParallel(docs []document.Document, rules []rule.Rule) func(*testing.B) {
	return func(b *testing.B) {
		for j := 0; j < b.N; j++ {
			pipeline.LaunchParallel(docs, rules, 0)
		}
	}
}

func NewBenchmarkLinear(docs []document.Document, rules []rule.Rule) func(*testing.B) {
	return func(b *testing.B) {
		for j := 0; j < b.N; j++ {
			pipeline.LaunchLinear(docs, rules, 0)
		}
	}
}

func main() {
	dirName := "data/measure"
	configPath := "cfg/config.json"

	docNum := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	for _, v := range docNum {
		docs, err := document.InputDocs(dirName, v)
		if err != nil {
			log.Fatalln(err)
		}

		configFile, err := os.Open(configPath)
		if err != nil {
			log.Fatalln(err)
		}

		rules, err := rule.ReadRules(configFile)
		if err != nil {
			log.Fatalln(err)
		}

		resParallel := testing.Benchmark(NewBenchmarkParallel(docs, rules))
		resLinear := testing.Benchmark(NewBenchmarkLinear(docs, rules))
		fmt.Printf("n = %d, parallel = %d, linear = %d\n", v, resParallel.NsPerOp()/1000, resLinear.NsPerOp()/1000)
		//fmt.Fprintf(timeFiles[j], "%d %d\n", len(strings1[i]), res.NsPerOp())
	}

}
