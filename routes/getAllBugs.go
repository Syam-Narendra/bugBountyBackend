package routes

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func GetAllRoutes(res *gin.Context, db *gorm.DB) {

	var bugs []Bug
	db.Find(&bugs)
	res.JSON(200, bugs)
	res.JSON(200, gin.H{
		"message": "Hello Gin baby !!",
	})
}
