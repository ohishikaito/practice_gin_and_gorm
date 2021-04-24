package db

import (
	"fmt"
	"go_myapp/app/model"
	"go_myapp/config"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

func Migrate() {
	fmt.Println("------------ migrate database... ------------")
	db, err := gorm.Open(config.SQLDriver, config.DatabaseURL)
	if err != nil {
		panic(err.Error())
	}
	fmt.Println("migrate...")
	db.AutoMigrate(&model.Book{})
	db.AutoMigrate(&model.User{})
	db.AutoMigrate(&model.BookDetail{})
	fmt.Println("------------ finish migrate! ----------------")
}
