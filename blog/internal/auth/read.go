package auth

import (
	"blog/models"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
	"net/http"
)

func GetUser(c echo.Context) error {
	id := c.Param("id")
	db := c.Get("db").(*gorm.DB)

	var user models.User
	if err := db.Where("id = ?", id).First(&user).Error; err != nil {
		return c.JSON(http.StatusNotFound, map[string]string{"message": "User not found"})
	}
	return c.JSON(http.StatusOK, user)
}

func GetUsers(c echo.Context) error {
	db := c.Get("db").(*gorm.DB)

	var users []models.User
	if err := db.Find(&users).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": "Failed to get users"})
	}
	return c.JSON(http.StatusOK, users)
}
