package main

import (
	"fmt"
	"lab_01/internal/table"
	"lab_01/internal/tasks"
	"os"
)

func main() {
	var task int
	fmt.Print("task > ")
	if _, err := fmt.Scan(&task); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	switch task {
	case 1:
		t := table.Generate(tasks.NewTask01(), "y", "x", 1.5, 1e-6)
		t.Print(1000)
		if err := t.Plot("1.png"); err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
	case 2:
		t := table.Generate(tasks.NewTask02(), "y", "x", 1.5, 1e-6)
		t.Print(1000)
		if err := t.Plot("2.png"); err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
	case 3:
		t := table.Generate(tasks.NewTask03(), "x", "y", 2, 1e-8)
		t.Print(1000)
		//if err := t.Plot("3.png"); err != nil {
		//	fmt.Println(err)
		//	os.Exit(1)
		//}
	default:
		fmt.Println("invalid variant")
		os.Exit(1)
	}
}
