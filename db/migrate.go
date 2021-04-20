// packageをmainにしないとDBを読み込めないっていううんこ仕様
package db

import (
	"fmt"
	"go_myapp/config"
	"go_myapp/model"

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
	fmt.Println("------------ finish migrate! ----------------")
}
