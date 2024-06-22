package handler

import (
	"time"

	"github.com/daifiyum/cat-box/subservice/database"
	auth "github.com/daifiyum/cat-box/subservice/middleware"
	"github.com/daifiyum/cat-box/subservice/models"
	"github.com/golang-jwt/jwt/v5"

	"github.com/gofiber/fiber/v2"
	"golang.org/x/crypto/bcrypt"
)

// CheckPasswordHash compare password with hash
func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

// Get user by identity
func getUserByIdentity(username string) (*models.Users, error) {
	db := database.DB
	var user models.Users
	err := db.Where("username = ?", username).First(&user).Error
	if err != nil {
		return nil, err
	}
	return &user, nil
}

// Login
func Login(c *fiber.Ctx) error {
	type Response struct {
		Username string `json:"username"`
		Token    string `json:"token"`
	}
	user := new(models.Users)

	if err := c.BodyParser(user); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"status": "error", "message": "Error on login request", "data": err})
	}

	// validate username
	userData, err := getUserByIdentity(user.Username)
	if err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"status": "error", "message": "Invalid username or password", "data": nil})
	}

	// validate password
	pass := user.Password
	if !CheckPasswordHash(pass, userData.Password) {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"status": "error", "message": "Invalid username or password", "data": nil})
	}

	// generate token
	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)
	claims["username"] = userData.Username
	claims["user_id"] = userData.ID
	claims["exp"] = time.Now().Add(time.Hour * 72).Unix()

	t, err := token.SignedString([]byte(auth.SecretKey))
	if err != nil {
		return c.SendStatus(fiber.StatusInternalServerError)
	}
	response := Response{
		Username: user.Username,
		Token:    t,
	}
	return c.JSON(fiber.Map{"status": "success", "message": "Success login", "data": response})
}
