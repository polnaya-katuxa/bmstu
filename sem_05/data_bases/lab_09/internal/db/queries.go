package db

import (
	"lab_09/internal/models"
	"math/rand"
	"time"
)

// Вывести все программы лояльности
func (db *DB) GetExpGames() ([]models.Games, error) {
	var l []models.Games
	if result := db.db.Table("games").Select("name, price").Order("price desc").Limit(10).Find(&l); result.Error != nil {
		return nil, result.Error
	}

	return l, nil
}

// Вывести все программы лояльности
func (db *DB) GetExpGamesRedis() ([]models.Games, error) {
	var l []models.Games
	if result := db.db.Table("games").Select("name, price").Order("price desc").Limit(10).Find(&l); result.Error != nil {
		return nil, result.Error
	}

	return l, nil
}

func (db *DB) GetUpdate() error {
	var n int
	db.db.Table("games").Model(&models.Games{}).Select("max(id)").Find(&n)
	id := rand.Int63n(int64(n))

	if result := db.db.Table("games").Where("id = ?", id).Update("price", 4500.0+float64(rand.Intn(1000))); result.Error != nil {
		return result.Error
	}

	return nil
}

const letterBytes = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

func randString(n int) string {
	rand.Seed(time.Now().UnixNano())

	b := make([]byte, n)
	for i := range b {
		b[i] = letterBytes[rand.Intn(len(letterBytes))]
	}
	return string(b)
}

func (db *DB) GetInsert() error {
	var id int
	db.db.Table("games").Model(&models.Games{}).Select("max(id)").Find(&id)

	g := models.Games{
		ID:              id + 1,
		Name:            randString(8),
		Genre:           "Shooter",
		ReleaseYear:     2012,
		Company:         "WorstCompany",
		AgeRating:       "T",
		Price:           4500.0 + float64(rand.Intn(1000)),
		MultiplayerMode: false,
	}

	if result := db.db.Table("games").Create(&g); result.Error != nil {
		return result.Error
	}

	return nil
}

func (db *DB) GetDelete() error {
	var n int64
	db.db.Count(&n)

	if result := db.db.Table("games").Where("id = ?", n).Delete(&models.Games{}); result.Error != nil {
		return result.Error
	}

	return nil
}
