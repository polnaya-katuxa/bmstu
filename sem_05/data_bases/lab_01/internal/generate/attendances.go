package generate

import (
	"fmt"
	"git.parasha.space/go/libs/faker"
	"time"
)

const (
	pricePerSecond = 0.1
)

type Attendance struct {
	Start    postgresTs `csv:"time_start"`
	ClubID   int        `csv:"id_club"`
	ClientID int        `csv:"id_client"`
	End      postgresTs `csv:"time_end"`
	Rating   int        `csv:"rating"`
	Price    float32    `csv:"price"`
}

func attendance(clubs []ComputerClub, clients []Client) (Attendance, error) {
	i := faker.IntWithLimits(0, len(clubs))
	j := faker.IntWithLimits(0, len(clients))

	a := Attendance{
		ClubID:   i + 1,
		ClientID: j + 1,
		Rating:   faker.IntWithLimits(1, 6),
	}

	latest := time.Time(clients[j].BirthDate)
	if time.Time(clubs[i].EstablishmentDate).After(latest) {
		latest = time.Time(clubs[i].EstablishmentDate)
	}

	clubOpenTime := time.Time(clubs[i].OpenTime)
	clubCloseTime := time.Time(clubs[i].CloseTime)

	var startTime time.Time
	var endTime time.Time

	randomTime := faker.Time(faker.NewTimeLimitTime(latest, time.Now().Add(-24*time.Hour).UTC()))

	if clubs[i].RoundTheClock {
		startTime = randomTime

		earliest := startTime.Add(24 * time.Hour)
		if time.Now().UTC().Before(earliest) {
			earliest = time.Now().UTC()
		}

		endTime = faker.Time(faker.NewTimeLimitTime(startTime.Add(time.Second), earliest))
	} else {
		openTime := time.Date(randomTime.Year(), randomTime.Month(), randomTime.Day(), clubOpenTime.Hour(), clubOpenTime.Minute(), 0, 0, time.UTC)
		closeTime := time.Date(randomTime.Year(), randomTime.Month(), randomTime.Day(), clubCloseTime.Hour(), clubCloseTime.Minute(), 0, 0, time.UTC)

		startTime = faker.Time(faker.NewTimeLimitTime(openTime, closeTime))
		endTime = faker.Time(faker.NewTimeLimitTime(startTime.Add(time.Second), closeTime))
	}

	a.Start = postgresTs(startTime)
	a.End = postgresTs(endTime)

	a.Price = float32(endTime.Unix()-startTime.Unix()) * pricePerSecond

	return a, nil
}

func Attendances(n int, clubs []ComputerClub, clients []Client) ([]Attendance, error) {
	attendances := make([]Attendance, n)
	var err error

	for i := range attendances {
		attendances[i], err = attendance(clubs, clients)
		if err != nil {
			return nil, err
		}
		fmt.Printf("attendances: %d/%d\r", i, n)
	}

	return attendances, nil
}
