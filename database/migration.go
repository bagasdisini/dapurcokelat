package database

import (
	"app/models"
	mysql "app/pkg"
	"fmt"
)

func RunMigration() {
	err := mysql.DB.AutoMigrate(&models.Data{})

	if err != nil {
		fmt.Println(err)
		panic("Migration Failed!")
	}

	fmt.Println("Migration Successful!")
}
