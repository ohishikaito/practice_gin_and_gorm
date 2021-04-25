package service

import (
	"go_myapp/config"
	"log"
	"os"

	"github.com/go-playground/validator"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

var db *gorm.DB
var err error
var validate *validator.Validate

func init() {
	// DB接続
	db, err = gorm.Open(config.SQLDriver, config.DatabaseURL)
	if err != nil {
		panic(err.Error())
	}

	// バリデーション
	validate = validator.New()

	// SQLログ出力先
	file, err := os.OpenFile("./db/log.txt", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		panic(err.Error())
	}
	log.SetOutput(file)
	db.LogMode(true)
	db.SetLogger(log.New(file, "", 0))
}
