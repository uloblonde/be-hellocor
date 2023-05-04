package database

import (
	"fmt"
	"halocorona/models"
	"halocorona/pkg/mysql"
)

func RunMigrations() {
	err := mysql.DB.AutoMigrate(
		&models.User{},
		&models.Article{},
		&models.Consulting{},
		&models.Response{},
	)

	if err != nil {
		fmt.Println(err)
		panic("Migration error")
	}

	fmt.Println("Migration succes")
}
