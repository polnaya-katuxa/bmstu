package main

import (
	"fmt"
	"lab_01/internal/table"
	"lab_01/internal/tasks"
	"os"
)

func main() {
	for {
		var task int
		fmt.Print("Enter task (0 to exit): ")
		if _, err := fmt.Scan(&task); err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		switch task {
		case 0:
			os.Exit(0)
		case 1:
			t := table.Generate(tasks.NewTask01(), "Y", "X", 0.95, 1e-5)
			t.Print(1000)
			if err := t.Plot("task_01.png"); err != nil {
				fmt.Println(err)
				os.Exit(1)
			}
		case 2:
			t := table.Generate(tasks.NewTask02(), "Y", "X", 1, 1e-5)
			t.Print(1000)
			if err := t.Plot("task_02.png"); err != nil {
				fmt.Println(err)
				os.Exit(1)
			}
		case 3:
			t := table.Generate(tasks.NewTask03(), "X", "Y", 2.006, 1e-10)
			t.Print(1000)
			if err := t.Plot("task_03.png"); err != nil {
				fmt.Println(err)
				os.Exit(1)
			}
		default:
			fmt.Println("Invalid variant")
			os.Exit(1)
		}
	}
}
