package model

import (
	"fmt"
	"time"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var DB *gorm.DB
var Err error

type Masterplan struct {
	SrNo      string    `gorm:"primary_key" json: "SrNo"`
	Activity  string    `gorm:"activity" json: "Activity"`
	StartDate time.Time `gorm:"start_date" json: "StartDate"`
	EndDate   time.Time `gorm:"end_date" json: "EndDate"`
}

func InitMigration() {
	// DB, err := gorm.Open("mysql", "user1:user1@(localhost)/masterplan?charset=utf8&parseTime=True&loc=Local")
	// if err != nil {
	// 	fmt.Println(err.Error())
	// 	panic("Failed to open database")
	// }
	// defer DB.Close()

	DB.AutoMigrate(&Masterplan{})

	fmt.Println("This is model file")
}

func (m *Masterplan) InsertMasterplan() error {
	result := DB.Create(&m)
	if result.Error != nil {
		return result.Error
	}
	fmt.Println(m, " Added")
	return nil
}
