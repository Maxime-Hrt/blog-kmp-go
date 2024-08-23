package api

import (
	"blog/internal/auth"
	"github.com/labstack/echo/v4"
)

func AuthController(e *echo.Echo) {
	authGroup := e.Group("/users")

	authGroup.POST("", auth.CreateUser)
	authGroup.GET("", auth.GetUsers)
	authGroup.GET("/:id", auth.GetUser)
	authGroup.PUT("/:id", auth.UpdateUser)
	authGroup.DELETE("/:id", auth.DeleteUser)
	authGroup.DELETE("/email/:email", auth.DeleteUserEmail)
}
