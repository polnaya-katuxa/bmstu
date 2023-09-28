package main

import (
	"fmt"
	_ "github.com/lib/pq"
	"lab_06/internal"
	"os"
)

func main() {
	dsn := "user=polnaya_katuxa password=1234 dbname=computer_club sslmode=disable"

	a, err := internal.New(dsn)
	if err != nil {
		fmt.Printf("Ошибка инициализации приложения: %s\n", err)
		os.Exit(1)
	}

	if err := a.Run(); err != nil {
		fmt.Printf("Ошибка запуска приложения: %s\n", err)
		os.Exit(1)
	}
}
