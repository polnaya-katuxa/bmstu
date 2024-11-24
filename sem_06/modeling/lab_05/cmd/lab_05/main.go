package main

import (
	"fmt"
	"lab_05/internal/app"
	"os"
)

func main() {
	if err := app.Run(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
