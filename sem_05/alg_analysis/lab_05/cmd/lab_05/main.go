package main

import (
	"flag"
	"fmt"
	"lab_05/internal/document"
	"lab_05/internal/pipeline"
	"lab_05/internal/rule"
	"log"
	"os"
)

func main() {
	configPath := flag.String("config", "config.json", "config file path")
	flag.Parse()

	var dirName string
	fmt.Println("\nВведите абсолютный или относительный путь\nк папке с документами: ")
	if _, err := fmt.Scanf("%s", &dirName); err != nil {
		log.Fatalln(err)
	}

	var docNum int
	fmt.Println("\nВведите количество документов: ")
	if _, err := fmt.Scanf("%d", &docNum); err != nil || docNum <= 0 {
		log.Fatalln(err)
	}

	docs, err := document.InputDocs(dirName, docNum)
	if err != nil {
		log.Fatalln(err)
	}

	configFile, err := os.Open(*configPath)
	if err != nil {
		log.Fatalln(err)
	}

	rules, err := rule.ReadRules(configFile)
	if err != nil {
		log.Fatalln(err)
	}

	var num int
	for num != 5 {
		fmt.Print("\nMenu:\n" +
			"1. Linear pipeline log\n" +
			"2. Parallel pipeline log\n" +
			"3. Linear pipeline res\n" +
			"4. Parallel pipeline res\n" +
			"5. Exit\n" +
			"\nEnter the chosen number: ")
		fmt.Scan(&num)
		if num < 1 || num > 5 {
			fmt.Println("\nTry again.")
			continue
		}

		if num == 1 {
			pipeline.LaunchLinear(docs, rules, 1)
		} else if num == 2 {
			pipeline.LaunchParallel(docs, rules, 1)
		} else if num == 3 {
			res := pipeline.LaunchParallel(docs, rules, 0)
			for _, v := range res {
				v.PrintTokens()
			}
		} else if num == 4 {
			res := pipeline.LaunchParallel(docs, rules, 0)
			for _, v := range res {
				v.PrintTokens()
			}
		} else if num == 5 {
			break
		}
	}
}
