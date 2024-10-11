package main

import (
	"bugbounty/backend/routes"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {
	url := os.Getenv("CONNECTION_STRING")
	db, err := gorm.Open(mysql.Open(url), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	router := gin.Default()
	router.GET("/", func(res *gin.Context) {
		res.JSON(http.StatusOK, gin.H{
			"message": "Hello Buggdddddddhhhhhhy !!",
		})
	})
	router.POST("/create-bug", func(ctx *gin.Context) {
		routes.CreateBug(ctx, db)
	})
	router.GET("/get-all-bugs", func(ctx *gin.Context) {
		routes.GetAllRoutes(ctx, db)
	})
	router.Run(":3000")

}
