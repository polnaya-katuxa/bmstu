package generate

import (
	"fmt"
	"git.parasha.space/go/libs/faker"
	"time"
)

const (
	maxParkingSpots = 20
)

type ComputerClub struct {
	ID                int          `csv:"id"`
	Address           string       `csv:"address"`
	OpenTime          postgresTime `csv:"open_time"`
	CloseTime         postgresTime `csv:"close_time"`
	EstablishmentDate postgresDate `csv:"establishment_date"`
	ParkingSpots      int          `csv:"parking_spot_num"`
	RoundTheClock     bool         `csv:"is_round_the_clock"`
}

func computerClub() ComputerClub {
	var club ComputerClub
	club.RoundTheClock = faker.Bool()
	if !club.RoundTheClock {
		club.OpenTime = postgresTime(faker.Time(faker.NewTimeLimitHour(7, 12)).Round(time.Hour))
		club.CloseTime = postgresTime(faker.Time(faker.NewTimeLimitHour(19, 23)).Round(time.Hour))
	} else {
		now := time.Now()

		club.OpenTime = postgresTime(time.Date(now.Year(), now.Month(), now.Day(), 0, 0, 0, 0, time.UTC))
		club.CloseTime = postgresTime(time.Date(now.Year(), now.Month(), now.Day(), 0, 0, 0, 0, time.UTC))
	}

	club.ParkingSpots = faker.IntWithLimits(0, maxParkingSpots+1)
	club.Address, _ = faker.FullAddress("ru")
	club.EstablishmentDate = postgresDate(faker.Time(faker.NewTimeLimitYear(2017, 2021)))

	return club
}

func ComputerClubs(n int) ([]ComputerClub, error) {
	clubs := make([]ComputerClub, n)

	addresses := make(map[string]struct{})

	for i := 0; i < n; i++ {
		clubs[i] = computerClub()
		clubs[i].ID = i + 1

		if _, ok := addresses[clubs[i].Address]; ok {
			i--
			continue
		} else {
			addresses[clubs[i].Address] = struct{}{}
		}

		fmt.Printf("clubs: %d/%d\r", i, n)
	}

	return clubs, nil
}
