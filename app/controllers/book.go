package controllers

import (
	"go_myapp/app/models"
	"go_myapp/service"

	"github.com/gin-gonic/gin"
)

func BookIndex(c *gin.Context) {
	BookService := service.BookService{}
	BookLists := BookService.GetBookList()
	c.JSON(200, gin.H{
		"data": BookLists,
		"unko": "tintin",
	})
}

func BookCreate(c *gin.Context) {
	// 反応しねえええええええ
	// pry.Pry()

	book := models.Book{}
	err := c.Bind(&book)
	// if err != nil {
	// 	fmt.Println("hoge")
	// 	return
	// }
	bookService := service.BookService{}
	err = bookService.CreateBook(&book)
	// err := bookService.CreateBook(&models.Book{
	// 	Title:   c.PostForm("title"),
	// 	Content: c.PostForm("content"),
	// })
	if err != nil {
		c.JSON(422, gin.H{
			"status": 422,
			"data":   err,
		})
	}
	c.JSON(201, gin.H{
		"status": "created",
	})
}

// router.POST("/create", func(ctx *gin.Context) {
// 	name := ctx.PostForm("name")
// 	email := ctx.PostForm("email")
// 	fmt.Printf("create uesr name = %s, email = %s \n", name, email)
// 	db.Create(&User{
// 		Name:  name,
// 		Email: email,
// 	})
// 	// defer db.Close()

// 	ctx.Redirect(302, "/")
// })

// func BookAdd(c *gin.Context) {
//      book := model.Book{}
//      err := c.Bind(&book)
//      if err != nil{
//          c.String(http.StatusBadRequest, "Bad request")
//          return
//      }
//     bookService :=service.BookService{}
//     err = bookService.SetBook(&book)
//     if err != nil{
//         c.String(http.StatusInternalServerError, "Server Error")
//         return
//     }
//     c.JSON(http.StatusCreated, gin.H{
//         "status": "ok",
//     })
// }
