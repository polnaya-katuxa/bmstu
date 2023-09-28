package main

import (
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"os"
	"time"
)

var n = 1

const (
	table = "club_pets"
)

type ClubPet struct {
	Type   string `json:"type"`
	Name   string `json:"name"`
	ClubID int    `json:"club_id"`
}

var names = []string{"Василий", "Стас", "Гоша", "Гоге", "Барсик", "Борис", "Александр", "Иван", "Карло"}
var types = []string{"Котик", "Осел", "Лошадка", "Собачка", "Барсук", "Енот", "Слоник", "Черепашка", "Лиса", "Шиншилла"}

func genClubPet() ClubPet {
	return ClubPet{
		Type:   types[rand.Intn(len(types))],
		Name:   names[rand.Intn(len(names))],
		ClubID: rand.Intn(50) + 1, // [1; 50]
	}
}

func genClubPets(n int) []ClubPet {
	res := make([]ClubPet, n)

	for i := range res {
		res[i] = genClubPet()
	}

	return res
}

func genClubPetsToFile(n int, filename string) error {
	file, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	pets := genClubPets(n)

	err = json.NewEncoder(file).Encode(pets)
	if err != nil {
		return err
	}

	return nil
}

func genFileName() string {
	defer func() {
		n++
	}()

	return fmt.Sprintf("%04d_%s_%s.json", n, table, time.Now().Format(time.RFC3339))
}

func main() {
	ticker := time.NewTicker(10 * time.Second)

	for {
		select {
		case <-ticker.C:
			log.Println("start generating file")
			err := genClubPetsToFile(5, fmt.Sprintf("../deployments/nifi/in_file/%s", genFileName()))
			if err != nil {
				log.Println(err)
			}
		}
	}
}
