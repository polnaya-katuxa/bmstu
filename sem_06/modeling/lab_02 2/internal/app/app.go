package app

import (
	"fmt"
	"lab_02/internal/app/task1"
	"lab_02/internal/app/task2"
	"lab_02/internal/app/task3"
)

func Run() error {
	var task int
	fmt.Print("Enter task number (1-3): ")
	if _, err := fmt.Scan(&task); err != nil {
		return fmt.Errorf("read task number: %w", err)
	}

	switch task {
	case 1:
		return task1.Run()
	case 2:
		return task2.Run()
	case 3:
		return task3.Run()
	}

	return nil
}
