package routes

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type Bug struct {
	ID             uint   `gorm:"primaryKey"`
	Rolename       string `json:"roleName"`
	BugTitle       string `json:"bugTitle"`
	Description    string `json:"description"`
	SeverityLevel  string `json:"severityLevel"`
	BrowserEnv     string `json:"browserEnv"`
	TechnologyUsed string `json:"technologyUsed"`
	Email          string `json:"email"`
	Name           string `json:"name"`
}

func CreateBug(res *gin.Context, db *gorm.DB) {

	db.AutoMigrate(&Bug{})
	var data Bug

	if err := res.ShouldBindJSON(&data); err != nil {
		res.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	fmt.Println(data.Rolename)
	newBug := Bug{
		Rolename:       data.Rolename,
		BugTitle:       data.BugTitle,
		Description:    data.Description,
		SeverityLevel:  data.SeverityLevel,
		BrowserEnv:     data.BrowserEnv,
		TechnologyUsed: data.TechnologyUsed,
		Email:          data.Email,
		Name:           data.Name,
	}
	result := db.Create(&newBug)
	println(result)
	if result.Error != nil {
		res.JSON(http.StatusBadRequest, gin.H{
			"error": result.Error,
		})
	}

	res.JSON(http.StatusAccepted, gin.H{
		"status": "Bug Created",
	})
}
