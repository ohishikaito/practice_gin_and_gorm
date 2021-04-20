package service

import (
	"errors"
	"go_myapp/config"
	"log"
	"os"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

var db *gorm.DB

func init() {
	err := errors.New("")
	db, err = gorm.Open(config.SQLDriver, config.DatabaseURL)
	if err != nil {
		panic(err.Error())
	}

	// ログ出力先
	file, err := os.OpenFile("./db/log", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		panic(err.Error())
	}
	log.SetOutput(file)

	// ★ログ設定
	db.LogMode(true)
	db.SetLogger(log.New(file, "", 0))
}
