package model

import (
	"log"
	"time"

	"github.com/jinzhu/gorm"
	// get mysql dialect
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

// DB ref
var DB *gorm.DB
var Err error

type Activity struct {
	SrNo      string    `gorm:"primary_key" json: "SrNo"`
	Activity  string    `gorm:"activity" json: "Activity"`
	StartDate time.Time `gorm:"start_date" json: "StartDate"`
	EndDate   time.Time `gorm:"end_date" json: "EndDate"`
}

func InitMigration() {
	DB.AutoMigrate(&Activity{})
	log.Println("Models AutoMigrate")
}

func (m *Activity) InsertActivity() error {
	result := DB.Create(&m)
	if result.Error != nil {
		return result.Error
	}
	log.Println("Data Added")
	return nil
}

func GetAllActivities() ([]*Activity, error) {
	var data []*Activity
	result := DB.Find(&data)
	if result.Error != nil {
		return data, result.Error
	}
	log.Println("Data Found")
	return data, nil
}
