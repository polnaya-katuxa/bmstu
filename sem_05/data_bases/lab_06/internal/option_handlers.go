package internal

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type optionHandler struct {
	name string
	f    func() error
}

func printSlice[T any](s []T) {
	for _, e := range s {
		fmt.Printf("%+v\n", e)
	}
}

func (a *App) avgOlderGamesPrice() error {
	var year int

	fmt.Print("Введите год: ")

	if _, err := fmt.Scan(&year); err != nil {
		return err
	}

	result, err := a.database.AvgOlderGamesPrice(year)
	if err != nil {
		return err
	}

	fmt.Println(result)

	return nil
}

func (a *App) attendData() error {
	result, err := a.database.AttendData()
	if err != nil {
		return err
	}

	printSlice(result)

	return nil
}

func (a *App) maxSumPriceForName() error {
	result, err := a.database.MaxSumPriceForName()
	if err != nil {
		return err
	}

	printSlice(result)

	return nil
}

func (a *App) tables() error {
	fmt.Print("Введите название таблицы: ")

	name, err := bufio.NewReader(os.Stdin).ReadString('\n')

	if err != nil {
		return err
	}

	name = strings.TrimSpace(name)
	result, err := a.database.Tables(name)
	if err != nil {
		return err
	}

	printSlice(result)

	return nil
}

func (a *App) avgPriceYear() error {
	var year int

	fmt.Print("Введите год: ")

	if _, err := fmt.Scan(&year); err != nil {
		return err
	}

	result, err := a.database.AvgPriceYear(year)
	if err != nil {
		return err
	}

	fmt.Println(result)

	return nil
}

func (a *App) hatePerson() error {
	var id int

	fmt.Print("Введите id нехорошего человека: ")

	if _, err := fmt.Scan(&id); err != nil {
		return err
	}

	result, err := a.database.HatePerson(id)
	if err != nil {
		return err
	}

	printSlice(result)

	return nil
}

func (a *App) startPriceUp() error {
	var discount float64

	fmt.Print("Введите насколько повысить цену (вещ. число): ")

	if _, err := fmt.Scan(&discount); err != nil {
		return err
	}

	err := a.database.StartPriceUp(discount)
	if err != nil {
		return err
	}

	fmt.Println("Цена повышена.")

	return nil
}

func (a *App) postgresDBName() error {
	result, err := a.database.PostgresDBName()
	if err != nil {
		return err
	}

	fmt.Println(result)

	return nil
}

func (a *App) createPetsTable() error {
	err := a.database.CreatePetsTable()
	if err != nil {
		return err
	}

	fmt.Println("Таблица создана.")

	return nil
}

func (a *App) insertPet() error {
	var id int
	var clubID int

	fmt.Print("Введите id: ")

	if _, err := fmt.Scan(&id); err != nil {
		return err
	}

	fmt.Print("Введите id клуба: ")

	if _, err := fmt.Scan(&clubID); err != nil {
		return err
	}

	fmt.Print("Введите вид животного: ")

	typePet, err := bufio.NewReader(os.Stdin).ReadString('\n')

	if err != nil {
		return err
	}

	typePet = strings.TrimSpace(typePet)

	fmt.Print("Введите имя животного: ")

	name, err := bufio.NewReader(os.Stdin).ReadString('\n')

	if err != nil {
		return err
	}

	name = strings.TrimSpace(name)

	err = a.database.InsertPet(Pet{
		ID:      id,
		Type:    typePet,
		Name:    name,
		Club_ID: clubID,
	})

	if err != nil {
		return err
	}

	fmt.Println("Запись вставлена.")

	return nil
}
