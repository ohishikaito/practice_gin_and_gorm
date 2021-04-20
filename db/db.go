package db

import (
	"go_myapp/config"
	"log"
	"os"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

func SqlConnect() (database *gorm.DB, err error) {
	db, err := gorm.Open(config.SQLDriver, config.DatabaseURL)
	if err != nil {
		panic(err.Error())
	}

	// 出力先
	file, err := os.OpenFile("./db/log", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		panic(err.Error())
	}
	log.SetOutput(file)

	// ★ログ設定
	db.LogMode(true)
	db.SetLogger(log.New(file, "", 0))

	return db, err
}
