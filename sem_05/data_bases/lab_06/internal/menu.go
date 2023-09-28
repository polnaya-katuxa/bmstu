package internal

import (
	"database/sql"
	"fmt"
	"os"
)

type App struct {
	database       *DB
	optionHandlers []optionHandler
}

func New(dsn string) (*App, error) {
	pureDB, err := sql.Open("postgres", dsn)
	if err != nil {
		return nil, err
	}

	a := &App{
		database: NewDB(pureDB),
	}

	a.optionHandlers = []optionHandler{
		{
			name: "Найти среднюю цену на игры младше заданного года",
			f:    a.avgOlderGamesPrice,
		},
		{
			name: "Вывести адреса клубов и логины посетителей, где рейтинг = 1",
			f:    a.attendData,
		},
		{
			name: "Вывести максимальную сумму трат для всех клиентов с данным именем",
			f:    a.maxSumPriceForName,
		},
		{
			name: "Вывести текстовые атрибуты заданной таблицы",
			f:    a.tables,
		},
		{
			name: "Найти среднюю цену на игры младше заданного года (вызов функции)",
			f:    a.avgPriceYear,
		},
		{
			name: "Повысить рейтинги всех посещений на 1 и деактивировать все карты \n заданного клиента по id, вывести его login и телефон",
			f:    a.hatePerson,
		},
		{
			name: "Повысить цену на указанное число единиц на все игры категории Паззлы",
			f:    a.startPriceUp,
		},
		{
			name: "Получить имя БД",
			f:    a.postgresDBName,
		},
		{
			name: "Создать таблицу с животными клубов",
			f:    a.createPetsTable,
		},
		{
			name: "Добавить животное в таблицу",
			f:    a.insertPet,
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
