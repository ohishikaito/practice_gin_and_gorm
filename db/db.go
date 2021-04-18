package db

import (
	"fmt"
	"log"
	"os"

	"go_myapp/app/models"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

func init() {
	migrate()
}

func migrate() {
	fmt.Println("------------ migrate database... ------------")
	db, err := SqlConnect()
	if err != nil {
		panic(err.Error())
	}
	// start migration
	fmt.Println("migration...")
	db.AutoMigrate(&models.Book{})
	// end
	fmt.Println("------------ finish migrate! ------------")
}

func SqlConnect() (database *gorm.DB, err error) {
	DBMS := "mysql"
	USER := "root"
	PASS := "root"
	// dockerのネットワークを経由してアクセスするから、コンテナ名にしろ
	PROTOCOL := "tcp(go_myapp_db_1)"
	DBNAME := "go_database"

	CONNECT := USER + ":" + PASS + "@" + PROTOCOL + "/" + DBNAME + "?charset=utf8&parseTime=true&loc=Asia%2FTokyo"

	db, err := gorm.Open(DBMS, CONNECT)
	if err != nil {
		panic(err.Error())
	}

	// 出力先
	file, err := os.OpenFile("./db/log", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		panic(err)
	}
	log.SetOutput(file)

	// ★ログ設定
	db.LogMode(true)
	db.SetLogger(log.New(file, "", 0))

	return db, err
}
