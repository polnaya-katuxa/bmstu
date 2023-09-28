package generate

import (
	"fmt"
	"git.parasha.space/go/libs/faker"
	"math/rand"
	"time"
)

type cardState int

const (
	a cardState = iota
	na
)

type Card struct {
	Number           string       `csv:"number"`
	LPID             int          `csv:"loyalty_program_id"`
	ClientID         int          `csv:"id_client"`
	RegistrationDate postgresDate `csv:"registration_date"`
	State            string       `csv:"state"`
}

var numbersMap = make(map[string]struct{})

func generateCardNumber() string {
	var num string

	for {
		num = fmt.Sprintf("%04d %04d %04d %04d", faker.IntWithLimits(0, 10000),
			faker.IntWithLimits(0, 10000), faker.IntWithLimits(0, 10000),
			faker.IntWithLimits(0, 10000))

		if _, ok := numbersMap[num]; !ok {
			break
		}
	}

	return num
}

func card(clients []Client, lps []LoyaltyProgram) Card {
	i := faker.IntWithLimits(0, len(lps))
	j := faker.IntWithLimits(0, len(clients))

	c := Card{
		LPID:     i + 1,
		ClientID: j + 1,
	}

	t := cardState(rand.Int() % 2)
	if t == a {
		c.State = "activated"
	} else {
		c.State = "non-activated"
	}

	clientBirth := time.Time(clients[j].BirthDate)
	c.RegistrationDate = postgresDate(faker.Time(faker.NewTimeLimitTime(clientBirth, time.Now())))

	c.Number = generateCardNumber()

	return c
}

func Cards(n int, clients []Client, lps []LoyaltyProgram) []Card {
	cards := make([]Card, n)

	for i := range cards {
		cards[i] = card(clients, lps)
		fmt.Printf("cards: %d/%d\r", i, n)
	}

	return cards
}
