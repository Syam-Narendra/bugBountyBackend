package main

import (
	"bugbounty/backend/routes"
	"fmt"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {
	godotenv.Load()
	url := os.Getenv("CONNECTION_STRING")
	fmt.Println(url)
	db, err := gorm.Open(mysql.Open(url), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	router := gin.Default()
	router.GET("/", func(res *gin.Context) {
		res.JSON(http.StatusOK, gin.H{
			"message": "Hello dellllllllll !!",
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
