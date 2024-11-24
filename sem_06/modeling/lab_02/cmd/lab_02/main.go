package main

import (
	"fmt"
	"lab_02/internal/app"
	"os"
)

func main() {
	if err := app.Run(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
