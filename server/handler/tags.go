package handler

import (
	db "main/server/database"
	"strconv"
	"net/url"

	"github.com/gofiber/fiber/v2"
)

func SetupTagsRoutes(router fiber.Router) {
	tagsRoutes := router.Group("/tags")

  tagsRoutes.Get("/", func(c *fiber.Ctx) error {
    ret := db.GetAllTags()
    return c.JSON(fiber.Map{"status": "success", "message": "List of Tags", "data": ret})
  })

  tagsRoutes.Get("recipeid/:recipeid", func(c *fiber.Ctx) error {
    parseId, err := strconv.Atoi(c.Params("recipeid"))
    if err != nil {
      return c.JSON(fiber.Map{"status": "failed", "message": "Unable to process ID"})
    }
    ret := db.GetTagsOfRecipe(parseId)
    return c.JSON(fiber.Map{"status": "success", "message": "Tags by Recipe ID", "data": ret})
  })

  tagsRoutes.Get("id/:id", func(c *fiber.Ctx) error {
    parseId, err := strconv.Atoi(c.Params("id"))
    if err != nil {
      return c.JSON(fiber.Map{"status": "failed", "message": "Unable to process ID"})
    }
    ret := db.GetTagById(parseId)
    return c.JSON(fiber.Map{"status": "success", "message": "Tag by ID", "data": ret})
  })

  tagsRoutes.Get("name/:name", func(c *fiber.Ctx) error {
    tagName, err := url.QueryUnescape(c.Params("name"))
    if err != nil {
      return c.JSON(fiber.Map{"status": "failed", "message": "Unable to process name"})
    }
    ret := db.GetTagByName(tagName)
    return c.JSON(fiber.Map{"status": "success", "message": "Tag by name", "data": ret})
  })

  tagsRoutes.Post("/autocomplete", func(c *fiber.Ctx) error {
    payload := struct {
      SubString string `json:"substring"`
    }{}

    if err := c.BodyParser(&payload); err != nil {
      return err
    }
    ret := db.GetAutocompleteTag(payload.SubString)
    return c.JSON(fiber.Map{"status": "success", "message": "List of Similar Tags", "data": ret, "len": len(ret)})
  })
}

