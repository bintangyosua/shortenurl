package routes

import (
	"github.com/bintangyosua/shortenurl/controllers"

	"github.com/labstack/echo/v4"
)

func RegisterRoutes(e *echo.Echo) {
	e.GET("/", controllers.Index)
	e.POST("/shorten", controllers.ShortenURL)
	e.GET("/:code", controllers.Redirect)
}
