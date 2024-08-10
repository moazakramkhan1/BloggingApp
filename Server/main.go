package main

import (
	"log"
	database "server/Database"
	routes "server/Routes"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/joho/godotenv"
)

func init() {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file")
	}
	database.Connect()
}

func main() {

	sqlDb, err := database.DB.DB()
	if err != nil {
		panic("Error in database connection")
	}
	defer sqlDb.Close()
	app := fiber.New()
	app.Use(cors.New(cors.Config{
		AllowOrigins:     "http://localhost:3000",
		AllowMethods:     "GET,POST,HEAD,PUT,DELETE,PATCH,OPTIONS",
		AllowHeaders:     "Origin,Content-Type,Accept,Authorization,",
		AllowCredentials: true,
	}))
	app.Static("/static", "./static")
	routes.SetUp(app)
	app.Listen(":8000")

}
