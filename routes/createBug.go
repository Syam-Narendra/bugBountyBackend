package routes

import (
	"github.com/gofiber/fiber/v2"
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

func CreateBug(res *fiber.Ctx, db *gorm.DB) error {
	var data Bug
	if err := res.BodyParser(&data); err != nil {
		return res.JSON(fiber.Map{"error": "Failed to parse request body"})
	}

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
	if result.Error != nil {
		return res.JSON(fiber.Map{"error": "Failed to create bug"})
	}

	return res.Status(200).JSON(fiber.Map{"message": "Bug created successfully"})
}
