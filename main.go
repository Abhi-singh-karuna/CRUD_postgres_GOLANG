package main

import (
	"log"
	"path/database"
	"path/router"

	"github.com/joho/godotenv"


	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"

	_ "github.com/jinzhu/gorm/dialects/postgres"
)

func main() {

	app := fiber.New()
	app.Use(logger.New())

	godotenv.Load()


	database.ConnectDB()

	router.SetupRoutes(app)
	log.Fatal(app.Listen(":4000"))

	defer database.DB.Close()
}
