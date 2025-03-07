package handler

// import (
// 	"errors"
// 	"main/server/config"
// 	"main/server/database"
// 	"main/server/model"
// 	"time"
// 
// 	"gorm.io/gorm"
// 
// 	"github.com/gofiber/fiber/v2"
// 	"github.com/golang-jwt/jwt"
// 	"golang.org/x/crypto/bcrypt"
// )
// 
// // CheckPasswordHash compare password with hash
// func CheckPasswordHash(password, hash string) bool {
// 	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
// 	return err == nil
// }
// 
// func getUserByEmail(e string) (*model.MemberInfo, error) {
// 	db := database.DB
// 	var user model.MemberInfo
// 	if err := db.Where(&model.MemberInfo{MiUserid: e}).Find(&user).Error; err != nil {
// 		if errors.Is(err, gorm.ErrRecordNotFound) {
// 			return nil, nil
// 		}
// 		return nil, err
// 	}
// 	return &user, nil
// }
// 
// func getUserByUsername(u string) (*model.MemberInfo, error) {
// 	db := database.DB
// 	var user model.MemberInfo
// 	if err := db.Where(&model.MemberInfo{MiUserid: u}).Find(&user).Error; err != nil {
// 		if errors.Is(err, gorm.ErrRecordNotFound) {
// 			return nil, nil
// 		}
// 		return nil, err
// 	}
// 	return &user, nil
// }
// 
// // Login get user and password
// func Login(c *fiber.Ctx) error {
// 	type LoginInput struct {
// 		Identity string `json:"identity"`
// 		Password string `json:"password"`
// 	}
// 	type UserData struct {
// 		ID       uint   `json:"id"`
// 		Username string `json:"username"`
// 		Email    string `json:"email"`
// 		Password string `json:"password"`
// 	}
// 	var input LoginInput
// 	var ud UserData
// 
// 	if err := c.BodyParser(&input); err != nil {
// 		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"status": "error", "message": "Error on login request", "data": err})
// 	}
// 	identity := input.Identity
// 	pass := input.Password
// 
// 	email, err := getUserByEmail(identity)
// 	if err != nil {
// 		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"status": "error", "message": "Error on email", "data": err})
// 	}
// 
// 	user, err := getUserByUsername(identity)
// 	if err != nil {
// 		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"status": "error", "message": "Error on username", "data": err})
// 	}
// 
// 	if email == nil && user == nil {
// 		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"status": "error", "message": "User not found", "data": err})
// 	}
// 
// 	if email == nil {
// 		ud = UserData{
// 			ID:       user.ID,
// 			Username: user.MiUsernm,
// 			Email:    user.MiUserid,
// 			Password: user.MiPasswd,
// 		}
// 	} else {
// 		ud = UserData{
// 			ID:       email.ID,
// 			Username: email.MiUsernm,
// 			Email:    email.MiUserid,
// 			Password: email.MiPasswd,
// 		}
// 	}
// 
// 	if !CheckPasswordHash(pass, ud.Password) {
// 		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"status": "error", "message": "Invalid password", "data": nil})
// 	}
// 
// 	token := jwt.New(jwt.SigningMethodHS256)
// 
// 	claims := token.Claims.(jwt.MapClaims)
// 	claims["username"] = ud.Username
// 	claims["user_id"] = ud.ID
// 	claims["exp"] = time.Now().Add(time.Hour * 72).Unix()
// 
// 	t, err := token.SignedString([]byte(config.Config("SECRET")))
// 	if err != nil {
// 		return c.SendStatus(fiber.StatusInternalServerError)
// 	}
// 
// 	return c.JSON(fiber.Map{"status": "success", "message": "Success login", "data": t})
// }
