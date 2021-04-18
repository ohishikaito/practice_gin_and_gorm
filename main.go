package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

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
	mux := http.NewServeMux()
	mux.HandleFunc("/", func(w http.ResponseWriter, q *http.Request) {
		message := map[string]string{
			"message": "hello world",
		}
		jsonMessage, err := json.Marshal(message)
	})

	// engine := gin.Default()
	// engine.GET("/", func(c *gin.Context) {
	// 	c.JSON(http.StatusOK, gin.H{
	// 		"message": "hello world",
	// 	})
	// })
	// engine.Run()

	// // Qiitaでやったやつら！簡単な方
	// db := sqlConnect()
	// db.AutoMigrate(&User{})
	// defer db.Close()

	// router := gin.Default()
	// router.LoadHTMLGlob("templates/*.html")

	// router.GET("/", func(ctx *gin.Context) {
	// 	db := sqlConnect()
	// 	var users []User
	// 	res := db.Order("created_at asc").Find(&users)
	// 	fmt.Println(res)
	// 	defer db.Close()

	// 	ctx.HTML(200, "index.html", gin.H{
	// 		"users": users,
	// 	})
	// })

	// router.POST("/create", func(ctx *gin.Context) {
	// 	db := sqlConnect()
	// 	name := ctx.PostForm("name")
	// 	email := ctx.PostForm("email")
	// 	fmt.Printf("create uesr name = %s, email = %s \n", name, email)
	// 	db.Create(&User{
	// 		Name:  name,
	// 		Email: email,
	// 	})
	// 	defer db.Close()

	// 	ctx.Redirect(302, "/")
	// })

	// router.DELETE("/delete/:id", func(ctx *gin.Context) {
	// 	db := sqlConnect()
	// 	n := ctx.Param("id")
	// 	id, err := strconv.Atoi(n)
	// 	if err != nil {
	// 		return
	// 	}
	// 	var user User
	// 	db.First(&user, id)
	// 	db.Delete(&user)
	// 	defer db.Close()

	// 	ctx.Redirect(302, "/")
	// })

	// // router.Runしないと認識しないからね！
	// router.Run()

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
	// fmt.Println(CONNECT)

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
