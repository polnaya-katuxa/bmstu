package generate

import (
	"fmt"
	"git.parasha.space/go/libs/faker"
)

type Machine struct {
	ID          int    `csv:"id"`
	Brand       string `csv:"brand"`
	Model       string `csv:"model"`
	Country     string `csv:"country"`
	ReleaseYear int    `csv:"release_year"`
	Type        string `csv:"type"`
}

func machine() (Machine, error) {
	info := faker.Machine() //("ru")

	c, err := faker.Country("ru")
	if err != nil {
		return Machine{}, err
	}

	return Machine{
		Brand:       info.Brand,
		Model:       info.Model,
		Country:     c,
		ReleaseYear: faker.PostgresYear(faker.NewTimeLimitYear(2010, 2021)),
		Type:        info.Type,
	}, nil
}

func Machines(n int) ([]Machine, error) {
	machines := make([]Machine, n)
	var err error

	for i := 0; i < n; i++ {
		machines[i], err = machine()
		if err != nil {
			return nil, nil
		}
		machines[i].ID = i + 1
		fmt.Printf("machines: %d/%d\r", i, n)
	}

	return machines, nil
}
