package database

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var Db *gorm.DB

func init() {
	dsn := "host=localhost port=5432 user=postgres password=postgres dbname=fake_admagic sslmode=disable"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	Db = db
	if err != nil {
		panic(err)
	}
}