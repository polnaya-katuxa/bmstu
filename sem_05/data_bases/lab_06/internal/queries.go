package internal

import "database/sql"

type DB struct {
	db *sql.DB
}

func NewDB(db *sql.DB) *DB {
	return &DB{
		db: db,
	}
}

// Выполнить скалярный запрос
// Найти среднюю цену на игры младше заданного года
func (db *DB) AvgOlderGamesPrice(year int) (float64, error) {
	var result float64

	if err := db.db.QueryRow("select avg(price) as a from games where release_year >= $1", year).Scan(&result); err != nil {
		return 0, err
	}

	return result, nil
}

type AttendInfo struct {
	Address string
	Login   string
}

// Выполнить запрос с несколькими соединениями (JOIN)
// Адрес клуба и логин посетителя, где рейтинг = 1
func (db *DB) AttendData() ([]AttendInfo, error) {
	query := `
select j2.address, j2.login
from (
    (computer_clubs join attendance a on computer_clubs.id = a.id_club) as j1 join clients as c
    on j1.id_client = c.id) as j2
where j2.rating = 1
`

	rows, err := db.db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	result := make([]AttendInfo, 0)
	for rows.Next() {
		var p AttendInfo
		if err := rows.Scan(&p.Address, &p.Login); err != nil {
			return nil, err
		}

		result = append(result, p)
	}

	return result, nil
}

type SumInfo struct {
	Name string
	Sum  float64
}

// Выполнить запрос с ОТВ(CTE) и оконными функциями
// Вывести максимальную сумму трат для всех клиентов с данным именем
func (db *DB) MaxSumPriceForName() ([]SumInfo, error) {
	query := `
with client_sums (id, sum, name) as (
    select id_client, sum(price) as sum,
           (select name
            from clients
            where id = id_client)
    from attendance
    group by id_client
)
select name, max(sum) over(partition by c.name)
from client_sums as c
`

	rows, err := db.db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	result := make([]SumInfo, 0)
	for rows.Next() {
		var s SumInfo
		if err := rows.Scan(&s.Name, &s.Sum); err != nil {
			return nil, err
		}

		result = append(result, s)
	}

	return result, nil
}

// Выполнить запрос к метаданным
// Вывести текстовые атрибуты заданной таблицы
func (db *DB) Tables(table string) ([]string, error) {
	query := `
select column_name from information_schema.columns
where table_name = $1 and data_type = 'text'
`

	rows, err := db.db.Query(query, table)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	result := make([]string, 0)
	for rows.Next() {
		var s string
		if err := rows.Scan(&s); err != nil {
			return nil, err
		}

		result = append(result, s)
	}

	return result, nil
}

// Вызвать скалярную функцию
// Найти среднюю цену на игры младше заданного года
func (db *DB) AvgPriceYear(year int) (float64, error) {
	var result float64

	if err := db.db.QueryRow("select avg_price_arter_year($1)", year).Scan(&result); err != nil {
		return 0, err
	}

	return result, nil
}

type ClientInfo struct {
	Login string
	Phone string
}

// Вызвать многооператорную или табличную функцию
// Повысить рейтинги всех посещений на 1 и деактивировать все карты заданного клиента по id,
// вывести его login и телефон
func (db *DB) HatePerson(ind int) ([]ClientInfo, error) {
	rows, err := db.db.Query("select * from hate_person($1)", ind)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	result := make([]ClientInfo, 0)
	for rows.Next() {
		var s ClientInfo
		if err := rows.Scan(&s.Login, &s.Phone); err != nil {
			return nil, err
		}

		result = append(result, s)
	}

	return result, nil
}

// Вызвать хранимую процедуру
// Повысить цену на указанное число единиц на все игры категории Паззлы
func (db *DB) StartPriceUp(price float64) error {
	_, err := db.db.Exec("call puzzle_up($1)", price)
	return err
}

// Вызвать системную функцию или процедуру
// Имя текущей базы данных
func (db *DB) PostgresDBName() (string, error) {
	var result string

	if err := db.db.QueryRow("select current_database()").Scan(&result); err != nil {
		return "", err
	}

	return result, nil
}

// Создать таблицу в базе данных, соответствующую тематике БД
// Животные-талисманы клуба
func (db *DB) CreatePetsTable() error {
	query := `
drop table if exists club_pets;
create table club_pets(
    id serial primary key,
    type text,
    name text,
    club_id int references computer_clubs on delete cascade
);
`

	_, err := db.db.Exec(query)
	return err
}

type Pet struct {
	ID      int
	Type    string
	Name    string
	Club_ID int
}

// Выполнить вставку данных в созданную таблицу с использованием
// инструкции INSERT или COPY
// Вставить животное
func (db *DB) InsertPet(p Pet) error {
	_, err := db.db.Exec("insert into club_pets(id, type, name, club_id) values ($1, $2, $3, $4) on conflict do nothing", p.ID, p.Type, p.Name, p.Club_ID)
	return err
}
