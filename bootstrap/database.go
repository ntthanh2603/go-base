package bootstrap

import (
	env "gin-base/utils"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var db *gorm.DB

func DatabaseConnect() {
	dsn := env.Load("POSTGRES_DSN")

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
