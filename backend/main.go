package main

import (
	"app/app/controller"
	"app/db"
	"log"
	"net/http"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
)

func init() {
	db.Migrate()
}

func main() {
	// AuthTest()
	Route()

	NotUsedGin()
	// QiitaNoYatsu()
	// openWebPage()
}

func Route() {
	engine := gin.Default()
	CORSHandler := cors.New(cors.Config{
		AllowAllOrigins:  true,
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "HEAD"},
		AllowHeaders:     []string{"Origin", "Content-Length", "Content-Type"},
		AllowCredentials: false,
		MaxAge:           12 * time.Hour,
	})
	// config.AllowOrigins = []string{"http://localhost:3000"}
	// config.AllowOrigins = []string{"http://127.0.0.1:3000"}
	// engine.Use(middleware.Cors())
	engine.Use(CORSHandler)
	// engine.Use(middleware.TestMiddleware())
	bookEngine := engine.Group("/books")
	{
		bookEngine.GET("/", controller.BookIndex)
		bookEngine.POST("/", controller.BookCreate)
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
	engine.GET("/users/:id/comment_books", controller.UserCommentBooks)
	// engine.Use(cors.New(cors.Config{
	// 	// 許可したいHTTPメソッドの一覧
	// 	AllowMethods: []string{
	// 		"POST",
	// 		"GET",
	// 		"OPTIONS",
	// 		"PUT",
	// 		"DELETE",
	// 	},
	// 	// 許可したいHTTPリクエストヘッダの一覧
	// 	AllowHeaders: []string{
	// 		"Access-Control-Allow-Headers",
	// 		"Content-Type",
	// 		"Content-Length",
	// 		"Accept-Encoding",
	// 		"X-CSRF-Token",
	// 		"Authorization",
	// 	},
	// 	// 許可したいアクセス元の一覧
	// 	AllowOrigins: []string{
	// 		"http://localhost:3000",
	// 		"http://localhost:3000/",
	// 		"*",
	// 	},
	// 	// 自分で許可するしないの処理を書きたい場合は、以下のように書くこともできる
	// 	// AllowOriginFunc: func(origin string) bool {
	// 	//  return origin == "https://www.example.com:8080"
	// 	// },
	// 	// preflight requestで許可した後の接続可能時間
	// 	// https://godoc.org/github.com/gin-contrib/cors#Config の中のコメントに詳細あり
	// 	MaxAge: 24 * time.Hour,
	// }))
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
