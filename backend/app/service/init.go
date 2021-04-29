package service

import (
	"app/auth"
	"app/config"
	"log"
	"os"

	firebase "firebase.google.com/go"
	"github.com/go-playground/validator"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

var (
	// DB接続用をグローバルに
	db *gorm.DB
	// バリデーションを定義
	validate    *validator.Validate
	firebaseApp *firebase.App
)

func init() {
	DatabaseInit()
	FirebaseInit()
}

func DatabaseInit() {
	var err error
	db, err = gorm.Open(config.SQLDriver, config.DatabaseURL)
	if err != nil {
		panic(err.Error())
	}

	// バリデーション
	validate = validator.New()

	// SQLログ出力先
	file, err := os.OpenFile("./db/sql_log", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		panic(err.Error())
	}
	log.SetOutput(file)
	db.LogMode(true)
	db.SetLogger(log.New(file, "", 0))
}

func FirebaseInit() {
	firebaseApp = auth.NewFirebaseApp()
}
