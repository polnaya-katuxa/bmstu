package generate

import (
	"fmt"
	"git.parasha.space/go/libs/faker"
)

type Client struct {
	ID         int          `csv:"id"`
	Name       string       `csv:"name"`
	Surname    string       `csv:"surname"`
	Patronymic string       `csv:"patronymic"`
	BirthDate  postgresDate `csv:"birth_date"`
	Sex        string       `csv:"sex"`
	Phone      string       `csv:"phone_number"`
	Login      string       `csv:"login"`
}

func client() (Client, error) {
	info := new(faker.PersonInfo)
	var err error
	for info.Surname == "" || info.Name == "" || info.Patronymic == "" {
		info, err = faker.Person("ru")
		if err != nil {
			return Client{}, err
		}
	}

	return Client{
		Name:       info.Name,
		Surname:    info.Surname,
		Patronymic: info.Patronymic,
		BirthDate:  postgresDate(faker.Time(faker.NewTimeLimitYear(1990, 2005))),
		Sex:        string(info.Gender),
		Phone:      faker.RussianMobilePhone(),
		Login:      faker.Login(),
	}, nil
}

func Clients(n int) ([]Client, error) {
	clients := make([]Client, n)
	var err error

	loginsMap := make(map[string]struct{})

	for i := 0; i < n; i++ {
		clients[i], err = client()
		if err != nil {
			return nil, err
		}

		// Check login uniq.
		if _, ok := loginsMap[clients[i].Login]; ok {
			i--
			continue
		} else {
			loginsMap[clients[i].Login] = struct{}{}
		}

		clients[i].ID = i + 1
		fmt.Printf("clients: %d/%d\r", i, n)
	}

	return clients, nil
}
