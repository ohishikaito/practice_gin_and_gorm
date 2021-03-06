package main

import (
	"app/app/controller"
	"app/middleware"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
)

func init() {
	// db.Migrate()
}

func main() {
	// AuthTest()
	Route()

	// NotUsedGin()
	// QiitaNoYatsu()
	// openWebPage()
}

func Route() {
	engine := gin.Default()
	engine.Use(middleware.Cors())
	bookEngine := engine.Group("books")
	{
		// indexとかなにかしらつけないと何故か動かないので、一旦このまま
		bookEngine.GET("/index", controller.BookIndex)
		bookEngine.POST("/create", controller.BookCreate)
		bookEngine.GET("/:id", controller.BookShow)
		bookEngine.PUT("/:id", controller.BookUpdate)
		bookEngine.DELETE("/:id", controller.BookDelete)

		bookDetailEngine := bookEngine.Group("/:id/book_detail")
		{
			bookDetailEngine.POST("/", controller.BookDetailCreate)
		}

		commentEngine := bookEngine.Group("/:id/comments")
		{
			commentEngine.GET("/", controller.CommentIndex)
			commentEngine.POST("/", controller.CommentCreate)
		}
	}
	// ginでuser/createしないといけないから
	engine.POST("/users/create", controller.UserCreate)
	engine.GET("/users/:id/comment_books", controller.UserCommentBooks)
	engine.Run(":8080")
}

// type post struct {
// 	Title string `json:"title"`
// 	Tag   string `json:"tag"`
// 	URL   string `json:"url"`
// }

// var public = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
// 	post := &post{
// 		Title: "title",
// 		Tag:   "tag",
// 		URL:   "url",
// 	}
// 	json.NewEncoder(w).Encode(post)
// })
// var private = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
// 	post := &post{
// 		Title: "VGolangとGoogle Cloud Vision APIで画像から文字認識するCLIを速攻でつくる",
// 		Tag:   "Go",
// 		URL:   "https://qiita.com/po3rin/items/bf439424e38757c1e69b",
// 	}
// 	json.NewEncoder(w).Encode(post)
// })

// func AuthTest() {

// 	r := mux.NewRouter()
// 	r.Handle("/private", auth.JwtMiddleware.Handler(private))
// 	r.Handle("/auth", auth.GetTokenHandler)
// 	r.Handle("/public", public)
// 	if err := http.ListenAndServe(":8080", r); err != nil {
// 		log.Fatal("ListenAndServe:", err)
// 	}
// }

func NotUsedGin() {
	app := controller.NewApp()
	r := controller.HttpRoute(app)
	http.Handle("/", r)
	log.Fatal(http.ListenAndServe(":8080", nil))
}

// func openWebPage() {
// 	engine.Static("/static", "./static") // http://localhost:8080/static/s.png で画像が見れる
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
