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

	Salary      float64 `json:"salary" gorm:"not null"`
	EmpPassword string  `json:"emp_password" gorm:"not null"`
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

	Balance  float64  `json:"balance" gorm:"default:0"`
	CUId     int      `json:"cu_id"`
	Customer Customer `gorm:"foreignKey:CUId;references:NationalId"`
}

type SafeBox struct {
	gorm.model

	MaximumValue float64     `json:"maximum_value" gorm:"not null"`
	CUId         int         `json:"cu_id"`
	SHId         int         `josn:"sh_id"`
	priceClass   int         `json:"price_class" gorm:"not null"`
	Customer     Customer    `gorm:"foreignKey:CUId;references:NationalId"`
	StorageHall  StorageHall `gorm:"foreignKey:SHId;references:ID"`
}

type Rent struct {
	gorm.Model

	REId     int      `json:"re_id" gorm:"unique;not null"`
	CUId     string   `json:"cu_id" gorm:"primaryKey"`
	SBId     int      `json:"sb_id" gorm:"primaryKey"`
	Customer Customer `gorm:"foreignKey:CUId;references:NationalId"`
	SafeBox  SafeBox  `gorm:"foreignKey:SBId;references:ID`
}

type Contract struct {
	gorm Model

	baseAmount      float64   `json:"base_amount" gorm:"not null"`
	fromTime        time.Time `json:"from_time" gorm:"not null"`
	toTime	        time.Time `json:"to_time" gorm:"not null"`
	duration        time.Time `json:"duration" gorm:"not null"`
	discountPercent	int       `json:"discount_percent" gorm:"check:discountPercent <= 100;default:0"`
	REId            int       `json:"re_id"`
	Rent            Rent      `gorm:"foreignKey:REId;references:REId"`
}

type Services struct {
	gorm.Model

	COId        int      `json:"co_id" gorm:"primaryKey"`
	serviceName string   `json:"service_name" gorm:"primaryKey;not null"`
	serviceCost float64  `json:"service_cost" gorm:"not null"`
	Contract    Contract `gorm:"foreignKey:COId;references:COId"`
}

type DamageReport {
	gorm.Model

	RId                 int                 `json:"r_id"`
	damageMeasure       float64             `json:"damage_measure" gorm:"not null"`
	damageDescription   string              `json:"damage_description"`
	SHEId               int                 `json:"she_id"`
	REId                int                 `json:"re_id"`
	Report              Report              `gorm:"foreignKey:RId;references:ID"`
	StorageHallEmployee StorageHallEmployee `gorm:"foreignKey:SHEId;references:SHEId"`
	Rent                Rent                `gorm:"foreignKey:REId;references:REId"`
}

type EvacuationReport struct {
	gorm.Model

	RId                 int                 `json:"r_id"`
	checkAmount         float64             `json:"check_amount" gorm:"not null"`
	checkStatus         string              `json:"check_status" gorm:"check:check_status in ('Done', 'Not Done');not null"`
	SHEId               int                 `json:"she_id"`
	REId                int                 `json:"re_id"`
	Report              Report              `gorm:"foreignKey:RId;references:ID"`
	StorageHallEmployee StorageHallEmployee `gorm:"foreignKey:SHEId;references:SHEId"`
	Rent                Rent                `gorm:"foreignKey:REId;references:REId"`
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
