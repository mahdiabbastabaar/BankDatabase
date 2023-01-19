package model

import (
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

const DSN = "host=postgres user=admin password=1234 dbname=bank port=5432 sslmode=disable TimeZone=Asia/Tehran"

var DB *gorm.DB

func InitialMigration() {
	db, err := gorm.Open(postgres.Open(DSN), &gorm.Config{})
	DB = db
	if err != nil {
		fmt.Println(err.Error())
		panic("failed to connect database")
	}

	// Migrate the schema
	err = db.AutoMigrate()
	if err != nil {
		return
	}
}
