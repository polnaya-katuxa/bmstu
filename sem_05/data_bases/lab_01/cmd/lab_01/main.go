package main

import (
	"db_gen/internal/generate"
	"github.com/gocarina/gocsv"
	"log"
	"os"
)

func writeToCSV(filename string, data interface{}) error {
	file, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	err = gocsv.MarshalFile(data, file)
	if err != nil {
		return err
	}

	return nil
}

func main() {
	games, err := generate.Games(1000)
	if err != nil {
		log.Fatalln(err)
	}

	err = writeToCSV("data/games.csv", games)

	log.Printf("Sucessfully generated %d games", len(games))

	clubs, err := generate.ComputerClubs(1000)
	if err != nil {
		log.Fatalln(err)
	}

	err = writeToCSV("data/clubs.csv", clubs)

	log.Printf("Sucessfully generated %d computer clubs", len(clubs))

	programs := generate.LoaltyPrograms()

	err = writeToCSV("data/loyalty_programs.csv", programs)

	log.Printf("Sucessfully generated %d loyalty programs", len(programs))

	clients, err := generate.Clients(1000)
	if err != nil {
		log.Fatalln(err)
	}

	err = writeToCSV("data/clients.csv", clients)

	log.Printf("Sucessfully generated %d clients", len(clients))

	machines, err := generate.Machines(100)
	if err != nil {
		log.Fatalln(err)
	}

	err = writeToCSV("data/machines.csv", machines)

	log.Printf("Sucessfully generated %d machines", len(machines))

	staff, err := generate.Staffs(10000, clubs)
	if err != nil {
		log.Fatalln(err)
	}

	err = writeToCSV("data/staff.csv", staff)

	log.Printf("Sucessfully generated %d staff", len(staff))

	attendances, err := generate.Attendances(100000, clubs, clients)
	if err != nil {
		log.Fatalln(err)
	}

	err = writeToCSV("data/attendances.csv", attendances)

	log.Printf("Sucessfully generated %d attendances", len(attendances))

	cards := generate.Cards(2000, clients, programs)

	err = writeToCSV("data/cards.csv", cards)

	log.Printf("Sucessfully generated %d cards", len(cards))

	machinesInClubs := generate.MachinesInClubs(20000, machines, clubs)

	err = writeToCSV("data/machines_in_clubs.csv", machinesInClubs)

	log.Printf("Sucessfully generated %d machines in clubs", len(machinesInClubs))

	gamesOnMachines := generate.GamesOnMachines(200000, machinesInClubs, games)

	err = writeToCSV("data/games_on_machines.csv", gamesOnMachines)

	log.Printf("Sucessfully generated %d games on machines", len(gamesOnMachines))
}
