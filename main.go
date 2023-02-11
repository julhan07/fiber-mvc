package main

import (
	"app-test-fiber/entity"
	"app-test-fiber/infra"
	"app-test-fiber/routes"
	"log"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
)

func main() {

	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	app := fiber.New()

	if os.Getenv("APP_ENV") == "DEV" {
		db := infra.GetDB()
		db.Debug().AutoMigrate(entity.UserEntity{})
	}

	routes.Router(app)
	app.Listen(":3000")
}
