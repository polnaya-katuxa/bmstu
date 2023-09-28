package app

import (
	"fmt"
	"lab_09/internal/models"
	"log"
	"os"
	"time"
)

var feedbacks []string

const (
	gKey           = "expencive_games_top"
	getDuration    = 500 * time.Millisecond
	changeDuration = 1000 * time.Millisecond
	limit          = 200
)

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

func (a *App) getExpGamesRedis() error {
	result, err := a.getExpGamesRedisB()
	if err != nil {
		return err
	}

	printSlice(result)

	return nil
}

func (a *App) getExpGamesRedisB() ([]models.Games, error) {
	result, err := a.redis.Get(gKey)
	if err == nil {
		return result, nil
	}

	result, err = a.database.GetExpGames()
	if err != nil {
		return nil, err
	}

	err = a.redis.Set(gKey, result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (a *App) getExpGames() error {
	result, err := a.database.GetExpGames()
	if err != nil {
		return err
	}

	printSlice(result)

	return nil
}

func (a *App) Delete() error {
	err := a.database.GetDelete()
	if err != nil {
		return err
	}

	return a.redis.Delete(gKey)
}

func (a *App) Update() error {
	err := a.database.GetUpdate()
	if err != nil {
		return err
	}

	return a.redis.Delete(gKey)
}

func (a *App) Insert() error {
	err := a.database.GetInsert()
	if err != nil {
		return err
	}

	return a.redis.Delete(gKey)
}

func (a *App) bench(change func() error, n int) func() error {
	return func() error {
		tickerDB := time.NewTicker(getDuration)
		tickerRedis := time.NewTicker(getDuration)
		tickerChange := time.NewTicker(changeDuration)

		fRedis, err := os.Create(fmt.Sprintf("data/%02d_redis.txt", n))
		if err != nil {
			return err
		}
		defer fRedis.Close()

		fDB, err := os.Create(fmt.Sprintf("data/%02d_db.txt", n))
		if err != nil {
			return err
		}
		defer fDB.Close()

		i := 0
		begin := time.Now()

		for i < limit*2 {
			select {
			case <-tickerDB.C:
				i++
				log.Println("bench db")

				start := time.Now()
				a.database.GetExpGames()
				end := time.Now()

				fmt.Fprintln(fDB, end.Sub(begin).Milliseconds(), end.Sub(start).Microseconds())
			case <-tickerRedis.C:
				i++
				log.Println("bench redis")

				start := time.Now()
				a.getExpGamesRedisB()
				end := time.Now()

				fmt.Fprintln(fRedis, end.Sub(begin).Milliseconds(), end.Sub(start).Microseconds())
			case <-tickerChange.C:
				if change != nil {
					log.Println("change data")
					change()
				}
			}
		}

		return nil
	}
}
