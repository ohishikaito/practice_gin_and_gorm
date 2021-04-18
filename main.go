package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

type User struct {
	gorm.Model
	Name  string
	Email string
}

func main() {
	// mux := http.NewServeMux()
	// mux.HandleFunc("/", func(w http.ResponseWriter, q *http.Request) {
	// 	message := map[string]string{
	// 		"message": "hello world",
	// 	}
	// 	jsonMessage, err := json.Marshal(message)
	// 	if err != nil {
	// 		panic(err.Error())
	// 	}
	// 	w.Write(jsonMessage)
	// })
	// http.ListenAndServe(":8080", mux)

	fmt.Println("hoge")
	fmt.Println("fuga")
	engine := gin.Default()
	engine.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "hello world!!!!??????????????/!!!!",
		})
	})
	engine.Run()

	// QiitaNoYatsu()
	// openWebPage()
}

func sqlConnect() (database *gorm.DB) {
	DBMS := "mysql"
	USER := "root"
	PASS := "root"
	// dockerのネットワークを経由してアクセスするから、コンテナ名にしろ
	PROTOCOL := "tcp(go_myapp_db_1)"
	DBNAME := "go_database"

	CONNECT := USER + ":" + PASS + "@" + PROTOCOL + "/" + DBNAME + "?charset=utf8&parseTime=true&loc=Asia%2FTokyo"

	db, err := gorm.Open(DBMS, CONNECT)
	if err != nil {
		panic(err)
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

func QiitaNoYatsu() {

	// // Qiitaでやったやつら！簡単な方
	db := sqlConnect()
	db.AutoMigrate(&User{})
	defer db.Close()

	router := gin.Default()
	router.LoadHTMLGlob("templates/*.html")

	router.GET("/", func(ctx *gin.Context) {
		db := sqlConnect()
		var users []User
		db.Order("created_at asc").Find(&users)
		defer db.Close()

		ctx.HTML(200, "index.html", gin.H{
			"users": users,
		})
	})

	router.POST("/create", func(ctx *gin.Context) {
		db := sqlConnect()
		name := ctx.PostForm("name")
		email := ctx.PostForm("email")
		fmt.Printf("create uesr name = %s, email = %s \n", name, email)
		db.Create(&User{
			Name:  name,
			Email: email,
		})
		defer db.Close()

		ctx.Redirect(302, "/")
	})

	router.POST("/delete/:id", func(ctx *gin.Context) {
		db := sqlConnect()
		n := ctx.Param("id")
		id, err := strconv.Atoi(n)
		if err != nil {
			return
		}
		var user User
		db.First(&user, id)
		db.Delete(&user)
		defer db.Close()

		ctx.Redirect(302, "/")
	})

	// router.Runしないと認識しないからね！
	router.Run()
}
