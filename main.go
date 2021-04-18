package main

import (
	"fmt"
	"time"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

func init() {
	fmt.Println("init????????")
}

func main() {
	db := sqlConnect()
	defer db.Close()

	// openWebPage()
}

func sqlConnect() (database *gorm.DB) {
	DBMS := "mysql"
	USER := "root"
	PASS := "root"
	// protocolはtcp(db:)で良いのか？？サービス名とコンテナ名にしろっていう記述が乱立してて草
	// PROTOCOL := "tcp(db:3308)"
	// つながらないからコンテナ名にした
	// PROTOCOL := "tcp(go_myapp_db_1:3308)"
	PROTOCOL := "tcp(go_myapp_db_1)"
	DBNAME := "go_database"

	CONNECT := USER + ":" + PASS + "@" + PROTOCOL + "/" + DBNAME + "?charset=utf8&parseTime=true&loc=Asia%2FTokyo"
	fmt.Println(CONNECT)

	count := 0
	db, err := gorm.Open(DBMS, CONNECT)
	if err != nil {
		for {
			if err == nil {
				fmt.Println("")
				break
			}
			fmt.Println(".")
			time.Sleep(time.Second)
			count++

			if count > 5 {
				fmt.Println("")
				fmt.Println("fail!!!!!!!!!!!!!!!!!!!!!!!!!")
				panic(err)
			}
			db, err = gorm.Open(DBMS, CONNECT)
		}
	}
	fmt.Println("ok")

	return db
}

func openWebPage() {
	router := gin.Default()
	router.LoadHTMLGlob("templates/*.html")

	router.GET("/", func(ctx *gin.Context) {
		ctx.HTML(200, "index.html", gin.H{})
	})

	router.Run()
}
