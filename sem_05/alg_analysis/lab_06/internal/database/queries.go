package database

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"lab_06_01/internal/models"
)

type DB struct {
	db *gorm.DB
}

func NewDB(dsn string) (*DB, error) {
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}
	return &DB{
		db: db,
	}, nil
}

func (db *DB) GetDBData() ([]models.Cat, error) {
	var cats []models.Cat
	if result := db.db.Table("cat").Order("fluffiness").Find(&cats); result.Error != nil {
		return nil, result.Error
	}

	return cats, nil
}
