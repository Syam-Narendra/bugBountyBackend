package main

import (
	"bugbounty/backend/routes"
	"log"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {
	app := fiber.New()
	app.Use(cors.New())
	godotenv.Load()
	url := os.Getenv("CONNECTION_STRING")
	db, err := gorm.Open(mysql.Open(url), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	app.Get("/", func(res *fiber.Ctx) error {
		return res.JSON(fiber.Map{"hello": "gfij"})
	})

	app.Post("/api/create-bug", func(res *fiber.Ctx) error {
		return routes.CreateBug(res, db)
	})

	app.Get("/api/get-all-bugs", func(res *fiber.Ctx) error {
		return routes.GetAllRoutes(res, db)
	})

	log.Fatal(app.Listen(":3000"))
}
