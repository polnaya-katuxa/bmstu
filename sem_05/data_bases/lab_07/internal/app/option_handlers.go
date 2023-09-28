package app

import (
	"bufio"
	"fmt"
	"lab_07/internal/db"
	"os"
	"strings"
	"time"
)

var feedbacks []string

type optionHandler struct {
	name string
	f    func() error
}

func printSlice[T any](s []T) {
	fmt.Println("RESULT:")
	for i, e := range s {
		fmt.Printf("%d: %+v\n", i, e)
	}
}

func (a *App) getAllLoyalties() error {
	result, err := a.database.GetAllLoyalties()
	if err != nil {
		return err
	}

	printSlice(result)

	return nil
}

func (a *App) getSortedAttendances() error {
	result, err := a.database.GetSortedAttendances()
	if err != nil {
		return err
	}

	printSlice(result)

	return nil
}

func (a *App) getOldClients() error {
	result, err := a.database.GetOldClients()
	if err != nil {
		return err
	}

	printSlice(result)

	return nil
}

func (a *App) getMaxPriceByRating() error {
	result, err := a.database.GetMaxPriceByRating()
	if err != nil {
		return err
	}

	printSlice(result)

	return nil
}

func (a *App) getMaxPriceByRatingP() error {
	var price float64

	fmt.Print("Введите цену-минимум: ")

	if _, err := fmt.Scan(&price); err != nil {
		return err
	}

	result, err := a.database.GetMaxPriceByRatingP(price)
	if err != nil {
		return err
	}

	printSlice(result)

	return nil
}

func (a *App) getFeedbacks() error {
	result, err := a.database.GetFeedbacks()
	if err != nil {
		return err
	}

	printSlice(result)
	feedbacks = result

	return nil
}

func (a *App) getUpdatedFeedbacks() error {
	result, err := a.database.GetUpdatedFeedbacks(feedbacks)
	if err != nil {
		return err
	}

	printSlice(result)
	feedbacks = result

	return nil
}

func (a *App) getNewFeedbacks() error {
	fmt.Print("Введите отзыв на сотрудников: ")
	stuff, err := bufio.NewReader(os.Stdin).ReadString('\n')
	if err != nil {
		return err
	}
	stuff = strings.TrimSpace(stuff)

	fmt.Print("Введите отзыв на атмосферу: ")
	atm, err := bufio.NewReader(os.Stdin).ReadString('\n')
	if err != nil {
		return err
	}
	atm = strings.TrimSpace(atm)

	fmt.Print("Введите отзыв на компьютеры: ")
	mac, err := bufio.NewReader(os.Stdin).ReadString('\n')
	if err != nil {
		return err
	}
	mac = strings.TrimSpace(mac)

	fmt.Print("Введите отзыв на парковку: ")
	park, err := bufio.NewReader(os.Stdin).ReadString('\n')
	if err != nil {
		return err
	}
	park = strings.TrimSpace(park)

	j := db.FeedbackJSON{
		Staff:      stuff,
		Machines:   mac,
		Atmosphere: atm,
		Parking:    park,
	}

	result, err := a.database.GetNewFeedbacks(feedbacks, j)
	if err != nil {
		return err
	}

	printSlice(result)
	feedbacks = result

	return nil
}

func (a *App) getAllLoyalties3() error {
	result, err := a.database.GetAllLoyalties3()
	if err != nil {
		return err
	}

	printSlice(result)

	return nil
}

func (a *App) getJoin3() error {
	result, err := a.database.GetJoin3()
	if err != nil {
		return err
	}

	printSlice(result)

	return nil
}

func (a *App) getInsert3() error {
	var name string
	fmt.Print("Введите имя: ")
	if _, err := fmt.Scan(&name); err != nil {
		return err
	}
	name = strings.TrimSpace(name)

	var surname string
	fmt.Print("Введите фамилию: ")
	if _, err := fmt.Scan(&surname); err != nil {
		return err
	}
	surname = strings.TrimSpace(surname)

	var patronymic string
	fmt.Print("Введите отчество: ")
	if _, err := fmt.Scan(&patronymic); err != nil {
		return err
	}
	patronymic = strings.TrimSpace(patronymic)

	var sex string
	fmt.Print("Введите пол: ")
	if _, err := fmt.Scan(&sex); err != nil {
		return err
	}
	sex = strings.TrimSpace(sex)

	var birth string
	fmt.Print("Введите дату рождения как 2002-09-23: ")
	if _, err := fmt.Scan(&birth); err != nil {
		return err
	}
	birthDate, err := time.Parse("2006-01-02", birth)
	if err != nil {
		return err
	}

	var phoneNum string
	fmt.Print("Введите номер телефона как +7 930 768-89-09: ")
	phoneNum, err = bufio.NewReader(os.Stdin).ReadString('\n')
	if err != nil {
		return err
	}
	phoneNum = strings.TrimSpace(phoneNum)

	var login string
	fmt.Print("Введите логин: ")
	if _, err := fmt.Scan(&login); err != nil {
		return err
	}
	login = strings.TrimSpace(login)

	id, err := a.database.GetCount()
	if err != nil {
		return err
	}

	c := db.Clients{
		ID:          int(id) + 1,
		Name:        name,
		Surname:     surname,
		Patronymic:  patronymic,
		BirthDate:   birthDate,
		Sex:         sex,
		PhoneNumber: phoneNum,
		Login:       login,
	}

	err = a.database.GetInsert(c)
	if err != nil {
		return err
	}

	fmt.Printf("Вставлен: %v\n", c)

	return nil
}

func (a *App) getUpdate3() error {
	var patronymic1 string
	fmt.Print("Введите отчество 1: ")
	if _, err := fmt.Scan(&patronymic1); err != nil {
		return err
	}
	patronymic1 = strings.TrimSpace(patronymic1)

	var patronymic2 string
	fmt.Print("Введите отчество 2: ")
	if _, err := fmt.Scan(&patronymic2); err != nil {
		return err
	}
	patronymic2 = strings.TrimSpace(patronymic2)

	err := a.database.GetUpdate(patronymic1, patronymic2)
	if err != nil {
		return err
	}

	return nil
}

func (a *App) getDelete3() error {
	var login string
	fmt.Print("Введите логин: ")
	if _, err := fmt.Scan(&login); err != nil {
		return err
	}

	err := a.database.GetDelete(login)
	if err != nil {
		return err
	}

	return nil
}

func (a *App) getPuzzleUp3() error {
	var up float64
	fmt.Print("Введите повышение цены: ")
	if _, err := fmt.Scan(&up); err != nil {
		return err
	}

	err := a.database.GetPuzzleUp(up)
	if err != nil {
		return err
	}

	return nil
}

//func (a *app.App) maxSumPriceForName() error {
//	result, err := a.database.MaxSumPriceForName()
//	if err != nil {
//		return err
//	}
//
//	printSlice(result)
//
//	return nil
//}
//
//func (a *app.App) attendData() error {
//	result, err := a.database.AttendData()
//	if err != nil {
//		return err
//	}
//
//	printSlice(result)
//
//	return nil
//}
//
//func (a *app.App) tables() error {
//	fmt.Print("Введите название таблицы: ")
//
//	name, err := bufio.NewReader(os.Stdin).ReadString('\n')
//
//	if err != nil {
//		return err
//	}
//
//	name = strings.TrimSpace(name)
//	result, err := a.database.Tables(name)
//	if err != nil {
//		return err
//	}
//
//	printSlice(result)
//
//	return nil
//}
//
//func (a *app.App) avgPriceYear() error {
//	var year int
//
//	fmt.Print("Введите год: ")
//
//	if _, err := fmt.Scan(&year); err != nil {
//		return err
//	}
//
//	result, err := a.database.AvgPriceYear(year)
//	if err != nil {
//		return err
//	}
//
//	fmt.Println(result)
//
//	return nil
//}
//
//func (a *app.App) hatePerson() error {
//	var id int
//
//	fmt.Print("Введите id нехорошего человека: ")
//
//	if _, err := fmt.Scan(&id); err != nil {
//		return err
//	}
//
//	result, err := a.database.HatePerson(id)
//	if err != nil {
//		return err
//	}
//
//	printSlice(result)
//
//	return nil
//}
//
//func (a *app.App) startPriceUp() error {
//	var discount float64
//
//	fmt.Print("Введите насколько повысить цену (вещ. число): ")
//
//	if _, err := fmt.Scan(&discount); err != nil {
//		return err
//	}
//
//	err := a.database.StartPriceUp(discount)
//	if err != nil {
//		return err
//	}
//
//	fmt.Println("Цена повышена.")
//
//	return nil
//}
//
//func (a *app.App) postgresDBName() error {
//	result, err := a.database.PostgresDBName()
//	if err != nil {
//		return err
//	}
//
//	fmt.Println(result)
//
//	return nil
//}
//
//func (a *app.App) createPetsTable() error {
//	err := a.database.CreatePetsTable()
//	if err != nil {
//		return err
//	}
//
//	fmt.Println("Таблица создана.")
//
//	return nil
//}
//
//func (a *app.App) insertPet() error {
//	var id int
//	var clubID int
//
//	fmt.Print("Введите id: ")
//
//	if _, err := fmt.Scan(&id); err != nil {
//		return err
//	}
//
//	fmt.Print("Введите id клуба: ")
//
//	if _, err := fmt.Scan(&clubID); err != nil {
//		return err
//	}
//
//	fmt.Print("Введите вид животного: ")
//
//	typePet, err := bufio.NewReader(os.Stdin).ReadString('\n')
//
//	if err != nil {
//		return err
//	}
//
//	typePet = strings.TrimSpace(typePet)
//
//	fmt.Print("Введите имя животного: ")
//
//	name, err := bufio.NewReader(os.Stdin).ReadString('\n')
//
//	if err != nil {
//		return err
//	}
//
//	name = strings.TrimSpace(name)
//
//	err = a.database.InsertPet(db.Pet{
//		ID:      id,
//		Type:    typePet,
//		Name:    name,
//		Club_ID: clubID,
//	})
//
//	if err != nil {
//		return err
//	}
//
//	fmt.Println("Запись вставлена.")
//
//	return nil
//}
