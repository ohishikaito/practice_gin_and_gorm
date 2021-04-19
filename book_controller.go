package main

// import (
// 	"go_myapp/models"
// 	"net/http"

// 	"github.com/gin-gonic/gin"
// )

// func BookIndex(c *gin.Context) {
// 	BookLists := models.GetBookList()
// 	c.JSON(200, gin.H{x
// 		"data": BookLists,
// 		"unko": "tintin",
// 	})
// }

// func BookCreate(c *gin.Context) {
// 	// form-dataにkeyを指定したら取り出せた。
// 	// でもフロントからjsonで送りたいんだよなあ
// 	// fmt.Println(c.PostForm("Title"), "Title")

// 	// ついでにサービスクラスが分かりにくいからモデルに移行する
// 	// 移行したらDBのmigrateがうんこになったから、なんとかしたい。サービスクラスは必要だった説w
// 	// っかjsonでレスポンス受け取れねえええええ

// 	book := models.Book{}
// 	err := c.Bind(&book)
// 	// if err != nil {
// 	// 	fmt.Println("hoge")
// 	// 	return
// 	// }

// 	if err := c.ShouldBindJSON(&models.Book); err != nil {
// 		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
// 		return
// 	}

// 	err = models.CreateBook(&book)
// 	// err := bookService.CreateBook(&models.Book{
// 	// 	Title:   c.PostForm("title"),
// 	// 	Content: c.PostForm("content"),
// 	// })
// 	if err != nil {
// 		c.JSON(422, gin.H{
// 			"status": 422,
// 			"data":   err,
// 		})
// 	}
// 	c.JSON(201, gin.H{
// 		"status": "created",
// 	})
// }

// // router.POST("/create", func(ctx *gin.Context) {
// // 	name := ctx.PostForm("name")
// // 	email := ctx.PostForm("email")
// // 	fmt.Printf("create uesr name = %s, email = %s \n", name, email)
// // 	db.Create(&User{
// // 		Name:  name,
// // 		Email: email,
// // 	})
// // 	// defer db.Close()

// // 	ctx.Redirect(302, "/")
// // })

// // func BookAdd(c *gin.Context) {
// //      book := model.Book{}
// //      err := c.Bind(&book)
// //      if err != nil{
// //          c.String(http.StatusBadRequest, "Bad request")
// //          return
// //      }
// //     bookService :=service.BookService{}
// //     err = bookService.SetBook(&book)
// //     if err != nil{
// //         c.String(http.StatusInternalServerError, "Server Error")
// //         return
// //     }
// //     c.JSON(http.StatusCreated, gin.H{
// //         "status": "ok",
// //     })
// // }
