package auth

import (
	"blog/models"
	"fmt"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
	"net/http"
)

type CreateUserRequest struct {
	UserName string `json:"user_name" validate:"required"`
	Email    string `json:"email" validate:"required"`
	Password string `json:"password" validate:"required"`
}

func CreateUser(c echo.Context) error {
	req := new(CreateUserRequest)
	if err := c.Bind(req); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"message": "Invalid Request"})
	}

	if err := c.Validate(req); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"message": "Validation failed", "error": err.Error()})
	}

	db := c.Get("db").(*gorm.DB)
	var existingUser models.User
	if err := db.Where("email = ?", req.Email).First(&existingUser).Error; err == nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"message": "User already exists"})
	}

	fmt.Printf("Registering user: \n username: %s \n email: %s \n", req.UserName, req.Email)

	user := models.User{
		UserName: req.UserName,
		Email:    req.Email,
	}

	if err := user.SetPassword(req.Password); err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": "Failed to set password"})
	}

	if err := db.Create(&user).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": "Failed to create user"})
	}

	return c.JSON(http.StatusCreated, map[string]string{"message": "User created successfully"})
}
