package main

import (
	"encoding/json"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

func init() {
	// db.Migrate()
}

type P struct {
	Name      string `json:"-"`
	Age       int
	Nicknames []string
}

func main() {
	b := []byte(`{"name":"mikeeeeeee"}`)
	var p P
	if err := json.Unmarshal(b, &p); err != nil {
		fmt.Println(err)
	}
	fmt.Println(p)
	fmt.Println(p.Name)

	v, _ := json.Marshal(p)
	fmt.Println(string(v))
}

// func main() {
// engine := gin.Default()
// engine.Use(middleware.TestMiddleware())
// bookEngine := engine.Group("/book")
// {
// 	bookEngine.GET("/index", controller.BookIndex)
// 	bookEngine.POST("/create", controller.BookCreate)
// 	bookEngine.GET("/:id", controller.BookShow)
// 	bookEngine.PUT("/:id/update", controller.BookUpdate)
// 	bookEngine.DELETE("/:id/delete", controller.BookDelete)
// }
// engine.Run()

// engine.Static("/static", "./static") // http://localhost:8080/static/s.png で画像が見れる
// NotUsedGin()
// QiitaNoYatsu()
// openWebPage()
// }

// func openWebPage() {
// 	router := gin.Default()
// 	router.LoadHTMLGlob("templates/*.html")

// 	router.GET("/test", func(ctx *gin.Context) {
// 		ctx.HTML(200, "templates/test.html", gin.H{})
// 	})

// 	router.Run()
// }

// type User struct {
// 	gorm.Model
// 	Name  string
// 	Email string
// }

// // // Qiitaでやったやつら！簡単な方
// // リクエストの後にdeferするとDBの接続が切れるからコメントアウトしてる〜
// func QiitaNoYatsu() {
// 	db, err := db.SqlConnect()
// 	if err != nil {
// 		panic(err.Error())
// 	}
// 	defer db.Close()

// 	router := gin.Default()
// 	router.LoadHTMLGlob("templates/*.html")

// 	router.GET("/", func(ctx *gin.Context) {
// 		var users []User
// 		db.Order("created_at asc").Find(&users)
// 		// defer db.Close()

// 		ctx.HTML(200, "index.html", gin.H{
// 			"users": users,
// 		})
// 	})

// 	router.POST("/create", func(ctx *gin.Context) {
// 		name := ctx.PostForm("name")
// 		email := ctx.PostForm("email")
// 		fmt.Printf("create uesr name = %s, email = %s \n", name, email)
// 		db.Create(&User{
// 			Name:  name,
// 			Email: email,
// 		})
// 		// defer db.Close()

// 		ctx.Redirect(302, "/")
// 	})

// 	router.POST("/delete/:id", func(ctx *gin.Context) {
// 		n := ctx.Param("id")
// 		id, err := strconv.Atoi(n)
// 		if err != nil {
// 			return
// 		}
// 		var user User
// 		db.First(&user, id)
// 		db.Delete(&user)
// 		// defer db.Close()

// 		ctx.Redirect(302, "/")
// 	})

// 	// router.Runしないと認識しないからね！
// 	router.Run()
// }

// func NotUsedGin() {
// 	mux := http.NewServeMux()
// 	mux.HandleFunc("/", func(w http.ResponseWriter, q *http.Request) {
// 		message := map[string]string{
// 			"message": "hello world",
// 		}
// 		jsonMessage, err := json.Marshal(message)
// 		if err != nil {
// 			panic(err.Error())
// 		}
// 		w.Write(jsonMessage)
// 	})
// 	http.ListenAndServe(":8080", mux)
// }
