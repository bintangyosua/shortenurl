package main

import (
	"github.com/bintangyosua/shortenurl/config"
	"github.com/bintangyosua/shortenurl/routes"

	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()

	config.ConnectDB()
	routes.RegisterRoutes(e)

	e.Static("/static", "static")
	e.Logger.Fatal(e.Start(":8080"))
}
