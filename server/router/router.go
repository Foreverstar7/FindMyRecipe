package router

import (
	route "main/server/handler"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

func SetupRoutes(app *fiber.App) {
	api := app.Group("/api", logger.New())

	// Setup the Node Routes
	route.SetupIngredientsRoutes(api)
  route.SetupRecipesRoutes(api)
  route.SetupUserRoutes(api)
  route.SetupTagsRoutes(api)
}
