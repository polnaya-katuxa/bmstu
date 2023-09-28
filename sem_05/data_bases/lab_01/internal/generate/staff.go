package generate

import (
	"fmt"
	"git.parasha.space/go/libs/faker"
	"time"
)

var positions = []string{
	"Администратор",
	"Менеджер",
	"Системный администратор",
	"Уборщик",
	"Секретарь",
	"Бухгалтер",
	"Бармен",
	"Консультант",
	"Старший администратор",
	"Стажер",
}

type Staff struct {
	ID             int          `csv:"id"`
	ClubID         int          `csv:"id_club"`
	Name           string       `csv:"name"`
	Surname        string       `csv:"surname"`
	Patronymic     string       `csv:"patronymic"`
	BirthDate      postgresDate `csv:"birth_date"`
	Sex            string       `csv:"sex"`
	Phone          string       `csv:"phone_number"`
	EmploymentDate postgresDate `csv:"employment_date"`
	Position       string       `csv:"position"`
}

func staff(clubs []ComputerClub) (Staff, error) {
	info := new(faker.PersonInfo)
	var err error
	for info.Surname == "" || info.Name == "" || info.Patronymic == "" {
		info, err = faker.Person("ru")
		if err != nil {
			return Staff{}, err
		}
	}

	i := faker.IntWithLimits(0, len(clubs))

	return Staff{
		ClubID:         i + 1,
		Name:           info.Name,
		Surname:        info.Surname,
		Patronymic:     info.Patronymic,
		BirthDate:      postgresDate(faker.Time(faker.NewTimeLimitYear(1990, 2005))),
		Sex:            string(info.Gender),
		Phone:          faker.RussianMobilePhone(),
		EmploymentDate: postgresDate(faker.Time(faker.NewTimeLimitTime(time.Time(clubs[i].EstablishmentDate), time.Now()))),
		Position:       faker.ArrayElement(positions),
	}, nil
}

func Staffs(n int, clubs []ComputerClub) ([]Staff, error) {
	staffs := make([]Staff, n)
	var err error

	for i := 0; i < n; i++ {
		staffs[i], err = staff(clubs)
		if err != nil {
			return nil, err
		}
		staffs[i].ID = i + 1
		fmt.Printf("staff: %d/%d\r", i, n)
	}

	return staffs, nil
}
