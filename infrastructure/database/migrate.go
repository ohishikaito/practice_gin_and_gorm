package database

import (
	"fmt"
	"go_myapp/app/models"
	"go_myapp/db"
)

func init() {
	Migrate()
}

func Migrate() {
	fmt.Println("------------ migrate database... ------------")
	db, err := db.SqlConnect()
	if err != nil {
		panic(err.Error())
	}
	// start migration
	fmt.Println("migration...")
	db.AutoMigrate(&models.Book{})
	// end
	fmt.Println("------------ finish migrate! ------------")
}
