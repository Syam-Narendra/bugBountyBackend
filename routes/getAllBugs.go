package routes

import (
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func GetAllRoutes(res *fiber.Ctx, db *gorm.DB) error {
	var bugs []Bug
	db.Find(&bugs)
	return res.Status(200).JSON(bugs)
}
