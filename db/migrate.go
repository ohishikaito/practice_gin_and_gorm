// packageをmainにしないとDBを読み込めないっていううんこ仕様
package db

import (
	"fmt"
	"go_myapp/config"
	"go_myapp/models"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

func main() {
	fmt.Println("------------ migrate database... ------------")
	db, err := gorm.Open(config.SQLDriver, config.DatabaseURL)
	if err != nil {
		panic(err.Error())
	}
	db.CreateTable(&models.User{})
	fmt.Println("------------ finish migrate! ----------------")
}
