package main

import (
	"fmt"
	"lab_09/internal/app"
	"os"
)

func main() {
	dsn := "host=localhost user=polnaya_katuxa password=1234 dbname=computer_club sslmode=disable"

	a, err := app.New(dsn)
	if err != nil {
		fmt.Printf("Ошибка инициализации приложения: %s\n", err)
		os.Exit(1)
	}

	if err := a.Run(); err != nil {
		fmt.Printf("Ошибка запуска приложения: %s\n", err)
		os.Exit(1)
	}
}
