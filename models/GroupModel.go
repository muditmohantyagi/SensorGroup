package models

import (
	"errors"
	"time"

	"gorm.io/gorm"
	"test.com/config"
)

/*
type Sensor struct {
	ID        uint
	GroupName string `gorm:"type:varchar(500)"`
	Code      int
	CodeName  string
	Species   []Species `gorm:"foreignKey:SensorID;references:ID"`
	Created   time.Time `gorm:"autoCreateTime"`
	Updated   int       `gorm:"autoUpdateTime"`
}
type Species struct {
	ID           uint `gorm:"primaryKey;type:int(10)"`
	SensorID     uint `gorm:"foreignKey"`
	X            int  `gorm:"default:0"`
	Y            int  `gorm:"default:0"`
	Z            int  `gorm:"default:0"`
	Temperature  float32
	Transparency int
	Name         string `gorm:"type:varchar(500)"`
	Count        int
	Created      time.Time `gorm:"autoCreateTime"`
	Updated      int       `gorm:"autoUpdateTime"`
}
*/

type Group struct {
	ID        uint
	GroupName string `gorm:"type:varchar(500)"`
	Code      int
	CodeName  string
	Created   time.Time `gorm:"autoCreateTime"`
	Updated   int       `gorm:"autoUpdateTime"`
	Sensor    []Sensor  `gorm:"foreignKey:GroupID;references:ID"`
	CommanVar float64
}

type Sensor struct {
	ID           uint
	GroupID      uint `gorm:"foreignKey"`
	X            int  `gorm:"default:0"`
	Y            int  `gorm:"default:0"`
	Z            int  `gorm:"default:0"`
	Temperature  float32
	Transparency int
	Created      time.Time `gorm:"autoCreateTime"`
	Updated      int       `gorm:"autoUpdateTime"`
	Species      []Species `gorm:"foreignKey:SensorID;references:ID"`
}
type Species struct {
	ID       uint   `gorm:"primaryKey;type:int(10)"`
	SensorID uint   `gorm:"foreignKey"`
	Name     string `gorm:"type:varchar(500)"`
	Count    int
	Created  time.Time `gorm:"autoCreateTime"`
	Updated  int       `gorm:"autoUpdateTime"`
}

var db = config.GoConnect()

func MaxSensor(group_name string) (Group, bool) {
	var sensor_max Group

	if result := db.Where("group_name=?", group_name).Order("code desc").Take(&sensor_max); result.Error != nil {

		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return sensor_max, false
		}
		return sensor_max, false

	}
	return sensor_max, true
}

/////
func SensorByGroup(group_name string) ([]Group, error) {
	var sensors []Group

	if result := db.Where("group_name=?", group_name).Order("code").Find(&sensors); result.Error != nil {

		return nil, result.Error

	}
	return sensors, nil
}
