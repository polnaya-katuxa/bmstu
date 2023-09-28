package db

import "encoding/json"

// Вывести все программы лояльности
func (db *DB) GetAllLoyalties() ([]map[string]interface{}, error) {
	var l []map[string]interface{}
	if result := db.db.Table("loyalty_programs").Find(&l); result.Error != nil {
		return nil, result.Error
	}

	return l, nil
}

// Вывести клиентов, у которых дата рождения в мае 1991
func (db *DB) GetOldClients() ([]map[string]interface{}, error) {
	var c []map[string]interface{}
	if result := db.db.Table("clients").Find(&c, "birth_date >= '1991-05-01' and birth_date <= '1991-05-31'"); result.Error != nil {
		return nil, result.Error
	}

	return c, nil
}

// Вывести посещения с рейтингом больше 3 и ценой выше 8500, сортировать по цене
func (db *DB) GetSortedAttendances() ([]map[string]interface{}, error) {
	var a []map[string]interface{}
	if result := db.db.Table("attendance").Order("price desc").Find(&a, "rating > ? and price > ?", 3, 8500); result.Error != nil {
		return nil, result.Error
	}

	return a, nil
}

func (db *DB) GetMaxPriceByRating() ([]map[string]interface{}, error) {
	var a []map[string]interface{}
	if result := db.db.Table("attendance").Select("rating, max(price) as price").Group("rating").Find(&a); result.Error != nil {
		return nil, result.Error
	}

	return a, nil
}

func (db *DB) GetMaxPriceByRatingP(price float64) ([]map[string]interface{}, error) {
	var a []map[string]interface{}

	if result := db.db.Table("attendance").Select("rating, avg(price) as price").Group("rating").Having("avg(price) > ?", price).Find(&a); result.Error != nil {
		return nil, result.Error
	}

	return a, nil
}

func (db *DB) GetFeedbacks() ([]string, error) {
	var a []string

	if result := db.db.Table("attendance").Select("feedback").Distinct("feedback").Find(&a); result.Error != nil {
		return nil, result.Error
	}

	return a, nil
}

func (db *DB) GetUpdatedFeedbacks(f []string) ([]string, error) {
	for i, v := range f {
		var j FeedbackJSON
		if err := json.Unmarshal([]byte(v), &j); err != nil {
			return nil, err
		}

		j.Parking = "so many cars i've even stole one, excellent!"

		fStr, err := json.Marshal(j)
		if err != nil {
			return nil, err
		}

		f[i] = string(fStr)
	}

	return f, nil
}

func (db *DB) GetNewFeedbacks(f []string, j FeedbackJSON) ([]string, error) {
	fJSON, err := json.Marshal(j)
	if err != nil {
		return nil, err
	}

	f = append(f, string(fJSON))

	return f, nil
}

func (db *DB) GetAllLoyalties3() ([]LoyaltyPrograms, error) {
	var l []LoyaltyPrograms
	if result := db.db.Table("loyalty_programs").Find(&l); result.Error != nil {
		return nil, result.Error
	}

	return l, nil
}

func (db *DB) GetJoin3() ([]Joined, error) {
	var l []Joined

	if result := db.db.Table("clients").Select("login,time_start,time_end,rating").Joins("join attendance on attendance.id_client = clients.id").Limit(10).Scan(&l); result.Error != nil {
		return nil, result.Error
	}

	return l, nil
}

func (db *DB) GetInsert(c Clients) error {
	if result := db.db.Table("clients").Create(&c); result.Error != nil {
		return result.Error
	}

	return nil
}

func (db *DB) GetUpdate(p1 string, p2 string) error {
	if result := db.db.Table("clients").Where("patronymic = ?", p1).Update("patronymic", p2); result.Error != nil {
		return result.Error
	}

	return nil
}

func (db *DB) GetDelete(l string) error {
	if result := db.db.Table("clients").Where("login = ?", l).Delete(&Clients{}); result.Error != nil {
		return result.Error
	}

	return nil
}

func (db *DB) GetCount() (int, error) {
	var l int

	result := db.db.Table("clients").Select("max(id) as i").Find(&l)
	if result.Error != nil {
		return 0, result.Error
	}

	return l, nil
}

func (db *DB) GetPuzzleUp(up float64) error {
	result := db.db.Raw("call puzzle_up(?)", up).Scan(nil)
	if result.Error != nil {
		return result.Error
	}

	return nil
}
