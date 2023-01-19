package handler

import (
	"fmt"
	db "main/server/database"
	"os"
	"time"

	"github.com/gofiber/fiber/v2"
	jwt "github.com/golang-jwt/jwt/v4"
)

var jwtKey = []byte(os.Getenv("JWT_SECRET"))

type Claims struct {
	Username string `json:"username"`
  UserId int `json:"userid"`
	jwt.RegisteredClaims
}

func SetupUserRoutes(router fiber.Router) {
  router.Post("/login", func (c *fiber.Ctx) error {
    req := struct {
      Username string `json:"username"`
      Password string `json:"password"`
    }{}

    if err := c.BodyParser(&req); err != nil {
      return err
    }
    q, err := db.UserLogin(req.Username, req.Password)
    if err != nil {
      return c.JSON(fiber.Map{"status":"failed", "message":"failed to login"})
    }

    expirationTime := time.Now().Add(time.Hour)

    claims := &Claims{
      Username: req.Username,
      UserId: q.Id,
      RegisteredClaims: jwt.RegisteredClaims{
        ExpiresAt: jwt.NewNumericDate(expirationTime),
      },
    }

    token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

    tokenString, err := token.SignedString(jwtKey)
    if err != nil {
      return c.JSON(fiber.Map{"status":"failed", "message":"Failed to create JWT"})
    }

    c.Cookie(&fiber.Cookie{
      Name: "token",
      Value: tokenString,
      Expires: expirationTime,
    })

    return c.JSON(fiber.Map{"status":"success", "token":tokenString})
  })

  router.Post("/signup", func (c *fiber.Ctx) error {
    req := struct {
      Username string `json:"username"`
      Password string `json:"password"`
    }{}

    if err := c.BodyParser(&req); err != nil {
      return err
    }
    q := db.UserSignUp(req.Username, req.Password)
    if q != true {
      return c.JSON(fiber.Map{"status":"failed", "message":"username already exists"})
    }
    return c.JSON(fiber.Map{"status":"success"})
  })

  router.Post("/refresh", func (c *fiber.Ctx) error {
    tknStr := c.Cookies("token")
    if tknStr == "" {
      return fiber.ErrUnauthorized
    }
      // Username: req.Username,
      // UserId: q.Id,
      // RegisteredClaims: jwt.RegisteredClaims{
      //   ExpiresAt: jwt.NewNumericDate(expirationTime),
      // },
    req := struct {
      Username string `json:"username"`
      UserId int `json:"userid"`
      Exp int64 `json:"exp"`
    }{}

    if err := c.BodyParser(&req); err != nil {
      return err
    }

    expirationTime := time.Unix(req.Exp, 0)
    claims := &Claims{
      Username: req.Username,
      UserId: req.UserId,
      RegisteredClaims: jwt.RegisteredClaims{
        ExpiresAt: jwt.NewNumericDate(expirationTime),
      },
    }
    tkn, err := jwt.ParseWithClaims(tknStr, claims, func(token *jwt.Token) (interface{}, error) {
      return jwtKey, nil
    })
    if err != nil {
      return fiber.ErrBadRequest
    }
    if !tkn.Valid {
      return fiber.ErrUnauthorized
    }

    if time.Until(claims.ExpiresAt.Time) > 5*time.Minute {
      fmt.Println(tknStr)
      return fiber.ErrBadRequest
    }

    expirationTime = time.Now().Add(time.Hour)
    claims.ExpiresAt = jwt.NewNumericDate(expirationTime)
    token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
    tokenString, err := token.SignedString(jwtKey)
    if err != nil {
      return fiber.ErrInternalServerError
    }

    c.Cookie(&fiber.Cookie{
      Name: "token",
      Value: tokenString,
      Expires: expirationTime,
    })

    return c.JSON(fiber.Map{"token":tokenString})
  })

  router.Post("/logout", func (c *fiber.Ctx) error {
    c.Cookie(&fiber.Cookie{
      Name: "token",
      Expires: time.Now(),
    })
    return c.JSON(fiber.Map{"status":"success"})
  })

  router.Post("/favorite/toggle", func (c *fiber.Ctx) error {// not secure
    tknStr := c.Cookies("token")
    if tknStr == "" {
      return fiber.ErrUnauthorized
    }

    req := struct {
      UserId int `json:"userId"`
      RecipeId int `json:"recipeId"`
    }{}

    if err := c.BodyParser(&req); err != nil {
      return err
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

    q := db.ToggleUserFavorite(req.UserId, req.RecipeId)
    if q != true {
      return c.JSON(fiber.Map{"status":"failed", "message":"failed to toggle favorite"})
    }
    return c.JSON(fiber.Map{"status":"success", "modifiedBy":claims.Username})
  })

  // router.Post("/favorite/test", func (c *fiber.Ctx) error {// not secure
  //   return c.JSON(fiber.Map{"status":"failed", "message":"Removed, only for testing"})
  //   req := struct {
  //     UserId int `json:"userId"`
  //     RecipeIds []int `json:"recipeIds"`
  //   }{}
  //   if err := c.BodyParser(&req); err != nil {
  //     return c.JSON(fiber.Map{"status": "failed", "message": "Failed to load user favorites"})
  //   }
  //
  //   q := db.TestFavorites(req.UserId, req.RecipeIds)
  //   if q != true {
  //     return c.JSON(fiber.Map{"status":"failed", "message":"failed to toggle favorite"})
  //   }
  //   return c.JSON(fiber.Map{"status":"success"})
  // })

  router.Post("/favorite/user", func (c *fiber.Ctx) error {
    tknStr := c.Cookies("token")
    if tknStr == "" {
      return fiber.ErrUnauthorized
    }

    req := struct {
      UserId int `json:"userId"`
    }{}

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

    if err := c.BodyParser(&req); err != nil {
      return c.JSON(fiber.Map{"status": "failed", "message": "Failed to load user favorites"})
    }
    ret := db.GetUserFavorites(req.UserId)
    if (len(ret) == 0) {
      return c.JSON(fiber.Map{"status": "failed", "message": "No user favorites", "data": "[]"})
    }
    return c.JSON(fiber.Map{"status": "success", "message": "User favorite IDs", "data": ret})
  })

}

