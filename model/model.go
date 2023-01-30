package model

import (
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"time"
)

const DSN = "host=postgres user=admin password=1234 dbname=bank port=5432 sslmode=disable TimeZone=Asia/Tehran"

var DB *gorm.DB

type Customer struct {
	gorm.Model

	NationalId int     `json:"national_id" gorm:"primaryKey"`
	Address    string  `json:"address"`
	FirstName  string  `json:"first_name" gorm:"not null"`
	LastName   string  `json:"last_name" gorm:"not null"`
	Sex        string  `json:"sex" gorm:"check:sex in ('male', 'female');not null"`
	Age        int     `json:"age" gorm:"not null"`
	Credit     float64 `json:"credit"`
}

type BusinessPlan struct {
	gorm.Model

	Description     string `json:"description"`
	DiscountPercent int    `gorm:"check:discount_percent <= 100;not null" json:"discount_percent"`
}

type BusinessCustomer struct {
	gorm.Model

	NationalId   int          `json:"nationalId" gorm:"primaryKey"`
	BPId         int          `json:"bpid"`
	Customer     Customer     `gorm:"foreignKey:NationalId;references:NationalId"`
	BusinessPlan BusinessPlan `gorm:"foreignKey:BPId;references:ID"`
}

type Report struct {
	gorm.Model

	OccurrenceTime time.Time `gorm:"not null"`
}

type Employee struct {
	gorm.Model

	Salary      float64 `gorm:"not null"`
	EmpPassword string  `gorm:"not null"`
	FirstName   string  `json:"first_name" gorm:"not null"`
	LastName    string  `json:"last_name" gorm:"not null"`
	Sex         string  `json:"sex" gorm:"check:sex in ('male', 'female');not null"`
	Age         int     `json:"age" gorm:"not null"`
}

type StorageHallEmployee struct {
	SHEId    int      `json:"she_id" gorm:"primaryKey"`
	Employee Employee `gorm:"foreignKey:SHEId;references:ID"`
}

type StorageHall struct {
	gorm.Model

	CameraNumbers       int                 `json:"camera_numbers" gorm:"default:0"`
	WallQuality         int                 `json:"wall_quality" gorm:"check:wall_quality in(1, 2, 3);not null"`
	SHEId               int                 `json:"she_id"`
	StorageHallEmployee StorageHallEmployee `gorm:"foreignKey:SHEId;references:SHEId"`
}

type Account struct {
	gorm.Model

	Balance float64 `json:"balance" gorm:"default:0"`
	CUId    int     `json:"cu-id"`
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
