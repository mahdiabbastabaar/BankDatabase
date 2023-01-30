package model

import (
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

const DSN = "host=postgres user=admin password=1234 dbname=bank port=5432 sslmode=disable TimeZone=Asia/Tehran"

var DB *gorm.DB

type Customer struct {
	gorm.Model

	NationalId int     `json:"national_id"`
	Address    string  `json:"address"`
	FirstName  string  `json:"first_name"`
	LastName   string  `json:"last_name"`
	Sex        string  `json:"sex"`
	Age        int     `json:"age"`
	Credit     float64 `json:"credit"`
}

type BusinessPlan struct {
	gorm.Model

	Description     string `json:"description"`
	DiscountPercent int    `gorm:"check:discount_percent <= 100" json:"discount_percent"`
}

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
