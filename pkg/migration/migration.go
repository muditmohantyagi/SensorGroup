package migration

import (
	"fmt"

	"test.com/config"
	"test.com/models"
)

func Migrate() {
	db := config.GoConnect()
	err := db.AutoMigrate(&models.Group{}, &models.Sensor{}, &models.Species{})
	if err != nil {
		fmt.Println("Migration error:", err.Error())
	} else {
		fmt.Println("Migration successfull")
	}

}
