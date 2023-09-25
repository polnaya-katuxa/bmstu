package main

import (
	"bufio"
	"fmt"
	"lab_03/internal/algorithms"
	"os"
	"strconv"
	"strings"
)

func inputArr() ([]int, error) {
	var arr = make([]int, 0, 100)
	var s string

	fmt.Print("\nEnter several integers > 0, separated by space:\n")
	in := bufio.NewReader(os.Stdin)
	s, err := in.ReadString('\n')
	if err != nil {
		fmt.Println("invalid input")
		return nil, err
	}
	s = strings.TrimSuffix(s, "\n")
	var elems = strings.Split(s, " ")
	if len(elems) == 0 {
		fmt.Println("invalid input")
		return nil, err
	}

	for _, e := range elems {
		num, err := strconv.Atoi(e)
		if err != nil {
			fmt.Println("invalid input")
			return nil, err
		}
		arr = append(arr, num)
	}

	return arr, nil
}

func main() {
	arr, err := inputArr()
	if err != nil {
		os.Exit(1)
	}

	var num int
	algs := []func([]int){
		algorithms.Pancakesort,
		algorithms.Quicksort,
		algorithms.Beadsort,
	}

	for num != 4 {
		fmt.Print("\nMenu:\n" +
			"1. Pancakesort\n" +
			"2. Quicksort\n" +
			"3. Beadsort\n" +
			"4. Exit\n" +
			"5. Reenter array\n" +
			"\nEnter the chosen number: ")
		fmt.Scanln(&num)
		if num < 1 || num > 5 {
			fmt.Println("\nTry again.")
			continue
		}

		if num > 0 && num < 4 {
			fmt.Println("\nArray: ", arr)
			copyA := make([]int, len(arr))
			copy(copyA, arr)
			algs[num-1](copyA)
			fmt.Println("\nSorted array: ", copyA)
		} else if num == 4 {
			break
		} else {
			copyA := make([]int, len(arr))
			copy(copyA, arr)
			arr, err = inputArr()
			if err != nil {
				fmt.Println("\nTry again.")
				copy(arr, copyA)
				continue
			}
		}
	}
}
