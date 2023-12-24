package router

import (
	"blog-platform/handler"

	_ "blog-platform/docs"

	"github.com/labstack/echo/v4"

	echoSwagger "github.com/swaggo/echo-swagger"
)

// @title Blog Platform API
// @version 1.0
// @description This is a blog posts server
// @termsOfService http://swagger.io/terms/

// @host localhost:8082
// @BasePath /
// @schemes http
func RegisterRoutes(e *echo.Echo) {
	// For SwaggerUI
	e.GET("/swagger/*", echoSwagger.WrapHandler)

	// For health check of application and mongodb
	e.GET("/healthcheck", handler.HealthCheckHandler)

	e.GET("/posts", handler.RetrieveAllPosts)

	e.POST("/posts", handler.CreateNewPost)

	e.GET("/posts/:id", handler.RetrievePostByID)
}
