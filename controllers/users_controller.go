package controllers

import (
	"go_myapp/config"
	"go_myapp/models"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

var db *gorm.DB

func init() {
	db, _ = gorm.Open(config.SQLDriver, config.DatabaseURL)
}

func UserIndex(c *gin.Context) {
	Users := []models.User{}
	if err := db.Find(&Users).Error; err != nil {
		panic(err.Error())
	}
	c.JSON(200, gin.H{
		"data": Users,
	})
}

// func BookIndex(c *gin.Context) {
// 	BookLists := models.GetBookList()
// 	c.JSON(200, gin.H{
// 		"data": BookLists,
// 		"unko": "tintin",
// 	})
// }

// func UserNew(c web.C, w http.ResponseWriter, r *http.Request) {
// }

// func UserCreate(c web.C, w http.ResponseWriter, r *http.Request) {
// }

// func UserEdit(c web.C, w http.ResponseWriter, r *http.Request) {
// }

// func UserUpdate(c web.C, w http.ResponseWriter, r *http.Request) {
// }

// func UserDelete(c web.C, w http.ResponseWriter, r *http.Request) {
// }
