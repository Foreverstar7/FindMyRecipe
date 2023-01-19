package handler

import (
	db "main/server/database"
	m "main/server/models"
	"strconv"

	"github.com/gofiber/fiber/v2"
	jwt "github.com/golang-jwt/jwt/v4"
)

func SetupRecipesRoutes(router fiber.Router) {
	recipesRoutes := router.Group("/recipes")

  recipesRoutes.Get("/", func(c *fiber.Ctx) error {
    queryPage := c.Query("page")
    var ret []m.Recipe
    if (queryPage == "") {
      ret = db.GetRecipesPage(1)
    } else {
      i, err := strconv.Atoi(queryPage)
      if err != nil {
        return fiber.ErrBadRequest
      }
      ret = db.GetRecipesPage(i)
    }
    return c.JSON(fiber.Map{"status": "success", "message": "List of Recipes", "data": ret})
  })

  recipesRoutes.Get("id/:id", func(c *fiber.Ctx) error {
    parseId, err := strconv.Atoi(c.Params("id"))
    if err != nil {
      return c.JSON(fiber.Map{"status": "failed", "message": "Unable to process ID"})
    }
    ret, qerr := db.GetRecipeById(parseId)
    if (qerr != nil) {
      return c.JSON(fiber.Map{"status": "failed", "message": "Recipe not found"})
    }
    return c.JSON(fiber.Map{"status": "success", "message": "Recipe by ID", "data": ret})
  })

  recipesRoutes.Get("author/:id", func(c *fiber.Ctx) error {
    parseId, err := strconv.Atoi(c.Params("id"))
    if err != nil {
      return c.JSON(fiber.Map{"status": "failed", "message": "Unable to process ID"})
    }
    ret := db.GetRecipesByAuthor(parseId)
    if (len(ret) == 0) {
      return c.JSON(fiber.Map{"status": "failed", "message": "Recipes not found"})
    }
    return c.JSON(fiber.Map{"status": "success", "message": "Recipes by author", "data": ret})
  })

  recipesRoutes.Get("/weekly", func(c *fiber.Ctx) error {
    ret := db.WeeklyFavorites()
    if ret == nil {
      return c.JSON(fiber.Map{"status": "failed", "message": "No Weekly Favorites", "data" : "[]"})
    }
    return c.JSON(fiber.Map{"status": "success", "message": "Weekly Favorite Recipes", "data" : ret})
  })

  recipesRoutes.Post("/tags", func(c *fiber.Ctx) error {
    payload := struct {
      TagIds []int `json:"tagIds"`
    }{}

    if err := c.BodyParser(&payload); err != nil {
      return err
    }

    var data []m.Recipe
    if len(payload.TagIds) == 1 {
      data = db.GetRecipesByTags(payload.TagIds[0])
    } else {
      data = db.GetRecipesByManyTags(payload.TagIds)
    }

    if data == nil {
      return c.JSON(fiber.Map{"status":"failed", "message": "List of Recipes by Tags", "data": "[]"})
    }
    return c.JSON(fiber.Map{"status": "success", "message": "List of Recipes by Tags", "data": data })
  })

  recipesRoutes.Post("/match", func(c *fiber.Ctx) error {
    payload := struct {
      IngredientIds []int `json:"ingredientIds"`
      TagIds []int `json:"tagIds"`
    }{}
    
    if err := c.BodyParser(&payload); err != nil {
      return err
    }

    var data []m.Recipe
    if (len(payload.TagIds) == 0) {
      // match by ingredientIds
      data = db.GetMatchingRecipesByIngredients(payload.IngredientIds)
    } else if (len(payload.IngredientIds) == 0) {
      // match by tagIds
      data = db.GetRecipesByManyTags(payload.TagIds)
    } else {
      // match by ingredientIds and tagIds
      data = db.GetMatchingRecipes(payload.IngredientIds, payload.TagIds)
    }

    if data == nil {
      return c.JSON(fiber.Map{"status":"failed", "message": "No Matching Recipes", "data": "[]"})
    }
    return c.JSON(fiber.Map{"status": "success", "message": "Matching Recipes", "data": data, "len": len(data) })
  })

  recipesRoutes.Post("/recommend", func(c *fiber.Ctx) error {
    payload := struct {
      UserId int `json:"userId"`
    }{}

    tknStr := c.Cookies("token")
    if tknStr == "" {
      return fiber.ErrUnauthorized
    }

    claims := &Claims{}

    tkn, err := jwt.ParseWithClaims(tknStr, claims, func(token *jwt.Token) (interface{}, error) {
      return jwtKey, nil
    })

    if err != nil {
      return fiber.ErrBadRequest
    }
    if !tkn.Valid {
      return fiber.ErrUnauthorized
    }
    if err := c.BodyParser(&payload); err != nil {
      return err
    }

    ret := db.GetRecommendation(payload.UserId)
    if len(ret) == 0 {
      return c.JSON(fiber.Map{"status":"failed", "message": "No Recommendation for User", "data": "[]"})
    }
    return c.JSON(fiber.Map{"status": "success", "message": "Recipe Recommendation", "data": ret })
  })

  recipesRoutes.Post("/add", func(c *fiber.Ctx) error {
    tknStr := c.Cookies("token")
    if tknStr == "" {
      return fiber.ErrUnauthorized
    }

    claims := &Claims{}

    tkn, err := jwt.ParseWithClaims(tknStr, claims, func(token *jwt.Token) (interface{}, error) {
      return jwtKey, nil
    })

    if err != nil {
      return fiber.ErrBadRequest
    }
    if !tkn.Valid {
      return fiber.ErrUnauthorized
    }
    var payload m.RecipeRequest
    if err := c.BodyParser(&payload); err != nil {
      return err
    }
    ret := db.AddNewRecipe(payload)
    if ret != true {
      return c.JSON(fiber.Map{"status":"failed", "message": "Failed to add recipe", "data": "[]"})
    }
    return c.JSON(fiber.Map{"status": "success", "message": "Recipe added", "data": ret })
  })

  recipesRoutes.Post("/similar", func(c *fiber.Ctx) error {
    payload := struct {
      IngredientIds []int `json:"ingredientIds"`
    }{}
    
    if err := c.BodyParser(&payload); err != nil {
      return err
    }

    data := db.GetSimilarRecipes(payload.IngredientIds)
    if data == nil {
      return c.JSON(fiber.Map{"status":"failed", "message": "No Similar Recipes", "data": "[]"})
    }
    return c.JSON(fiber.Map{"status": "success", "message": "Similar Recipes by Ingredients", "data": data })
  })
}

