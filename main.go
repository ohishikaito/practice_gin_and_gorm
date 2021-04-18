package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"go_myapp/app/controllers"
	"go_myapp/db"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

// 別パッケージのファイルが読み込めん。どっかーが悪い？

func main() {
	controllers.Sample()

	// engine := gin.Default()
	// ua := ""
	// engine.Use(func(c *gin.Context) {
	// 	ua = c.GetHeader("User-Agent")
	// 	c.Next()
	// })
	// engine.GET("/", func(c *gin.Context) {
	// 	c.JSON(http.StatusOK, gin.H{
	// 		"message":    "hello world addddddddddddda",
	// 		"User-Agent": ua,
	// 	})
	// })
	// engine.Static("/static", "./static")

	// // 下の処理はいじるな。理由は分からんけど起動しない
	// port := os.Getenv("PORT")
	// if len(port) == 0 {
	// 	port = "8080"
	// }
	// engine.Run(":" + port)

	// NotUsedGin()
	QiitaNoYatsu()
	// openWebPage()
}

func openWebPage() {
	router := gin.Default()
	router.LoadHTMLGlob("templates/*.html")

	router.GET("/test", func(ctx *gin.Context) {
		ctx.HTML(200, "test.html", gin.H{})
	})

	router.Run()
}

type User struct {
	gorm.Model
	Name  string
	Email string
}

func QiitaNoYatsu() {
	// // Qiitaでやったやつら！簡単な方
	db := db.SqlConnect()
	db.AutoMigrate(&User{})
	// defer db.Close()

	router := gin.Default()
	router.LoadHTMLGlob("templates/*.html")

	router.GET("/", func(ctx *gin.Context) {
		// db := db.SqlConnect()
		var users []User
		db.Order("created_at asc").Find(&users)
		// defer db.Close()

		ctx.HTML(200, "index.html", gin.H{
			"users": users,
		})
	})

	router.POST("/create", func(ctx *gin.Context) {
		// db := db.SqlConnect()
		name := ctx.PostForm("name")
		email := ctx.PostForm("email")
		fmt.Printf("create uesr name = %s, email = %s \n", name, email)
		db.Create(&User{
			Name:  name,
			Email: email,
		})
		// defer db.Close()

		ctx.Redirect(302, "/")
	})

	router.POST("/delete/:id", func(ctx *gin.Context) {
		// db := db.SqlConnect()
		n := ctx.Param("id")
		id, err := strconv.Atoi(n)
		if err != nil {
			return
		}
		var user User
		db.First(&user, id)
		db.Delete(&user)
		// defer db.Close()

		ctx.Redirect(302, "/")
	})

	// router.Runしないと認識しないからね！
	router.Run()
}

func NotUsedGin() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", func(w http.ResponseWriter, q *http.Request) {
		message := map[string]string{
			"message": "hello world",
		}
		jsonMessage, err := json.Marshal(message)
		if err != nil {
			panic(err.Error())
		}
		w.Write(jsonMessage)
	})
	http.ListenAndServe(":8080", mux)
}
