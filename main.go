package main

import (
	"fmt"
	"log"
	"main/server/router"
  "main/server/database"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
  // "github.com/gofiber/fiber/v2/middleware/encryptcookie"
	"github.com/gofiber/template/html"
	"github.com/joho/godotenv"
)

func main() {
	//template render engine
	engine := html.New("./dist", ".html")

  database.ConnectDB()

	app := fiber.New(fiber.Config{
		Views: engine, //set as render engine
	})

	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	// env := os.Getenv("APP_ENV")
	port := os.Getenv("PORT")

	fmt.Println(port)

	if port == "" {
		port = "8080"
	}

	app.Use(logger.New())
	app.Use(cors.New(cors.Config{
		AllowOrigins:     "*",
		AllowHeaders:     "Origin, Content-Type, Accept, Authorization",
		AllowMethods:     "GET, HEAD, PUT, PATCH, POST, DELETE",
		AllowCredentials: true,
	}))
//  app.Use(encryptcookie.New(encryptcookie.Config{
//    Key: os.Getenv("COOKIE_SECRET"),
//  }))

	router.SetupRoutes(app)

	app.Static("/", "./dist")
	app.Get("/", mainPage)

	app.Listen(":" + port)
}

func mainPage(c *fiber.Ctx) error {
	return c.Render("index", nil)
}
