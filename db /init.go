package db

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var db *gorm.DB

func New() {
	dsn := "host=os user=gorm password=gorm dbname=gorm port=9920 sslmode=disable"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	// Migrate the schema
	db.AutoMigrate()
}

func Repository(value interface{}) *gorm.DB {
	return db.Model(value)
}
