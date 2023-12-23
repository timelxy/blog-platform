package router

import (
	"ocr/handler"

	"github.com/labstack/echo/v4"
)

// Register api routes
func RegisterRoutes(e *echo.Echo) {
	e.GET("/posts", handler.RetrieveAllPosts)

	e.POST("/posts", handler.CreateNewPost)

	e.GET("/posts/:id", handler.RetrievePostByID)
}
