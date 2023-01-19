package handler

import (
	db "main/server/database"
	"strconv"
  "net/url"
	"github.com/gofiber/fiber/v2"
)

func SetupIngredientsRoutes(router fiber.Router) {
	ingredientsRoutes := router.Group("/ingredients")

  ingredientsRoutes.Get("/", func(c *fiber.Ctx) error {
    ret := db.GetIngredients()
    return c.JSON(fiber.Map{"status": "success", "message": "List of Ingredients", "data": ret})
  })

  ingredientsRoutes.Get("id/:id", func(c *fiber.Ctx) error {
    parseId, err := strconv.Atoi(c.Params("id"))
    if err != nil {
      return c.JSON(fiber.Map{"status": "failed", "message": "Unable to process ID"})
    }
    ret := db.GetIngredientById(parseId)
    return c.JSON(fiber.Map{"status": "success", "message": "Ingredient by ID", "data": ret})
  })

  ingredientsRoutes.Get("byRecipeId/:id", func(c *fiber.Ctx) error {
    parseId, err := strconv.Atoi(c.Params("id"))
    if err != nil {
      return c.JSON(fiber.Map{"status": "failed", "message": "Unable to process Recipe ID"})
    }
    ret := db.GetIngredientsByRecipeId(parseId)
    if len(ret) == 0 {
      return c.JSON(fiber.Map{"status": "failed", "message": "Ingredients not found"})
    }
    return c.JSON(fiber.Map{"status": "success", "message": "Ingredients by Recipe ID", "data": ret})
  })
  
  ingredientsRoutes.Get("name/:name", func(c *fiber.Ctx) error {
    parseName, err := url.QueryUnescape(c.Params("name"))
    ret := db.GetIngredientByName(parseName)
    if (err != nil) {
      return c.JSON(fiber.Map{"status": "failed", "message": "Ingredient not found"})
    }
    return c.JSON(fiber.Map{"status": "success", "message": "Ingredient by name", "data": ret})
  })

  ingredientsRoutes.Post("/autocomplete", func(c *fiber.Ctx) error {
    payload := struct {
      SubString string `json:"substring"`
    }{}

    if err := c.BodyParser(&payload); err != nil {
      return err
    }
    ret := db.GetAutocompleteIngredient(payload.SubString)
    return c.JSON(fiber.Map{"status": "success", "message": "List of Similar Ingredients", "data": ret})
  })
}

