package main

import (
	"bufio"
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
	"os"
	"strings"
)

type DBSQL struct {
	db *sql.DB
}

func NewDBSQL(dsn string) (*DBSQL, error) {
	db, err := sql.Open("postgres", dsn)
	if err != nil {
		return nil, err
	}
	return &DBSQL{
		db: db,
	}, nil
}

type DBG struct {
	db *gorm.DB
}

func NewDBG(dsn string) (*DBG, error) {
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}
	return &DBG{
		db: db,
	}, nil
}

func printSlice[T any](s []T) {
	fmt.Println("RESULT:")
	for i, e := range s {
		fmt.Printf("%d: %+v\n", i, e)
	}
}

func (db *DBSQL) GetDepsMore10() error {
	query := `
select dep
from employee
group by dep
having count(id) > 10;
`

	rows, err := db.db.Query(query)
	if err != nil {
		return err
	}
	defer rows.Close()

	fmt.Println("Вывести отделы, где более 10 сотрудников")
	for rows.Next() {
		var name string
		if err := rows.Scan(&name); err != nil {
			return err
		}

		fmt.Println(name)
	}

	return nil
}

func (db *DBSQL) GetEmployeesStay() error {
	query := `
select id, fio
from employee
where id in(
    select eid
    from kpp
    group by eid, kdate, ktype
    having ktype=2 and min(ktime) >= '18:00'
);
`

	rows, err := db.db.Query(query)
	if err != nil {
		return err
	}
	defer rows.Close()

	fmt.Println("Вывести сотрудников, которые не выходят с места весь рабочий день (пусть рабочий день c 9:00 до 18:00)")
	for rows.Next() {
		var id int
		var name string
		if err := rows.Scan(&id, &name); err != nil {
			return err
		}

		fmt.Printf("%d: %s\n", id, name)
	}

	return nil
}

func (db *DBSQL) GetDepLaters(data string) error {
	query := `
select distinct dep
from employee
where id in
      (
          select eid
          from kpp
          where ktype = 1 and kdate = $1
          group by eid
          having min(ktime) > '9:00'
      );
`

	rows, err := db.db.Query(query, data)
	if err != nil {
		return err
	}
	defer rows.Close()

	fmt.Println("Вывести отделы, где есть сотрудники, опоздавшие в данный день")
	for rows.Next() {
		var name string
		if err := rows.Scan(&name); err != nil {
			return err
		}

		fmt.Println(name)
	}

	return nil
}

func (db *DBG) GetDepsMore10App() error {
	var a []map[string]interface{}

	if result := db.db.Table("employee").Select("dep").Group("dep").Having("count(id) > 10").Find(&a); result.Error != nil {
		return result.Error
	}

	printSlice(a)

	return nil
}

func (db *DBG) GetEmployeesStayApp() error {
	var a []int
	var b []map[string]interface{}
	var result1 *gorm.DB
	var result2 *gorm.DB

	if result1 = db.db.Table("kpp").Select("eid").Group("eid, kdate, ktype").Having("ktype=2 and min(ktime) >= '18:00'").Find(&a); result1.Error != nil {
		return result1.Error
	}

	if result2 = db.db.Table("employee").Select("id, fio").Where("id in ?", a).Find(&b); result2.Error != nil {
		return result2.Error
	}

	printSlice(b)

	return nil
}

func (db *DBG) GetDepLatersApp(data string) error {
	var a []int
	var b []string
	var result1 *gorm.DB
	var result2 *gorm.DB

	if result1 = db.db.Table("kpp").Select("eid").Where("ktype = 1 and kdate = ?", data).Group("eid").Having("min(ktime) > '9:00'").Find(&a); result1.Error != nil {
		return result1.Error
	}

	if result2 = db.db.Table("employee").Select("dep").Distinct("dep").Where("id in ?", a).Find(&b); result2.Error != nil {
		return result2.Error
	}

	printSlice(b)

	return nil
}

func main() {
	dsn1 := "host=localhost user=polnaya_katuxa password=1234 dbname=rk3 sslmode=disable"

	DBG1, err := NewDBG(dsn1)
	if err != nil {
		log.Fatal(err)
	}

	dsn2 := "user=polnaya_katuxa password=1234 dbname=rk3 sslmode=disable"

	DBSQL1, err := NewDBSQL(dsn2)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Print("\nВведите дату в формате 2006-12-01 для задания №3: ")
	data, err := bufio.NewReader(os.Stdin).ReadString('\n')
	if err != nil {
		log.Fatal(err)
	}
	data = strings.TrimSpace(data)

	fmt.Println("\nBD-LEVEL:")
	err = DBSQL1.GetDepsMore10()
	if err != nil {
		log.Fatal(err)
	}

	err = DBSQL1.GetEmployeesStay()
	if err != nil {
		log.Fatal(err)
	}

	err = DBSQL1.GetDepLaters(data)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("\nAPP-LEVEL:")
	err = DBG1.GetDepsMore10App()
	if err != nil {
		log.Fatal(err)
	}

	err = DBG1.GetEmployeesStayApp()
	if err != nil {
		log.Fatal(err)
	}

	err = DBG1.GetDepLatersApp(data)
	if err != nil {
		log.Fatal(err)
	}
}
