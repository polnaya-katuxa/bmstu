package app

import (
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"lab_07/internal/db"
	"os"
)

type App struct {
	database       *db.DB
	optionHandlers []optionHandler
}

func New(dsn string) (*App, error) {
	pureDB, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	a := &App{
		database: db.NewDB(pureDB),
	}

	a.optionHandlers = []optionHandler{
		{
			name: "Вывести все программы лояльности",
			f:    a.getAllLoyalties,
		},
		{
			name: "Вывести клиентов, у которых дата рождения в мае 1991",
			f:    a.getOldClients,
		},
		{
			name: "Вывести посещения с рейтингом больше 3 и ценой выше 8500, сортировать по цене",
			f:    a.getSortedAttendances,
		},
		{
			name: "Вывести максимальные цены для рейтингов",
			f:    a.getMaxPriceByRating,
		},
		{
			name: "Вывести средние цены для рейтингов для клиента с введённой ценой-минимум",
			f:    a.getMaxPriceByRatingP,
		},
		{
			name: "Прочитать данные об отзывах из таблицы посещений в формате json",
			f:    a.getFeedbacks,
		},
		{
			name: "Добавить данные о парковке",
			f:    a.getUpdatedFeedbacks,
		},
		{
			name: "Добавить новый отзыв по вводу",
			f:    a.getNewFeedbacks,
		},
		{
			name: "Вывести информацию о всех программах лояльности (классы сущностей)",
			f:    a.getAllLoyalties3,
		},
		{
			name: "Вывести 10 записей о логине посетившего, времени посещения и рейтинге (классы сущностей)",
			f:    a.getJoin3,
		},
		{
			name: "Вставить клиента",
			f:    a.getInsert3,
		},
		{
			name: "Обновить отчества клиентов",
			f:    a.getUpdate3,
		},
		{
			name: "Удалить клиента по логину",
			f:    a.getDelete3,
		},
		{
			name: "Вызвать процедуру увеличения цен на паззлы",
			f:    a.getPuzzleUp3,
		},
		{
			name: "Выход",
			f: func() error {
				os.Exit(0)
				return nil
			},
		},
	}

	return a, nil
}

func (a *App) printMenu() {
	fmt.Println("\nMenu:")
	for i, r := range a.optionHandlers {
		fmt.Printf("%02d - %s\n", i, r.name)
	}
}

func (a *App) Run() error {
	for {
		a.printMenu()

		fmt.Println()

		var option int
		fmt.Print("Введите номер пункта меню: ")
		if _, err := fmt.Scan(&option); err != nil {
			fmt.Println(err)
			continue
		}

		if option < 0 || option >= len(a.optionHandlers) {
			fmt.Printf("Error: invalid menu option\n")
			continue
		}

		if err := a.optionHandlers[option].f(); err != nil {
			fmt.Printf("Error: %s\n", err)
			continue
		}

		fmt.Println()
	}
}
