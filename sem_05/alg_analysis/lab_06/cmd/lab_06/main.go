package main

import (
	"fmt"
	"lab_06_01/internal/database"
	"lab_06_01/internal/dictionary"
	"lab_06_01/internal/query"
	"lab_06_01/internal/server"
	"log"
	"os"
	"os/signal"
)

func main() {
	dsn := "host=localhost user=user password=password dbname=cats sslmode=disable port=5434"

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)

	go func() {
		db, err := database.NewDB(dsn)
		if err != nil {
			log.Fatal(err)
		}

		dict, err := dictionary.InitDictionary(db)
		if err != nil {
			log.Fatal(err)
		}

		manager := query.New(dict)

		if err := server.Run(manager, 8080); err != nil {
			log.Fatalln(err)
		}
	}()

	<-c
	fmt.Println("\nЗавершение работы...")
}
