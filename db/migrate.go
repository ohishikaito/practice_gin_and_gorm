// packageをmainにしないとDBを読み込めないっていううんこ仕様
package main

import (
	"fmt"
	"go_myapp/config"
	"go_myapp/models"
	"log"
	"os"

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

	// 出力先
	file, err := os.OpenFile("./db/log", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		panic(err.Error())
	}
	log.SetOutput(file)

	// ★ログ設定
	db.LogMode(true)
	db.SetLogger(log.New(file, "", 0))
}
