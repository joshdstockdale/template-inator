package main

import (
	"template-inator/handler"
	"template-inator/middleware"

	"github.com/labstack/echo/v4"
	// "template-inator/handler"
)

func main() {
	app := echo.New()
	app.Static("/public", "public")
	pageHandler := handler.PageHandler{}
	app.Use(middleware.WithUser)
	app.GET("/", pageHandler.HandleDashShow)

	app.Start(":3000")
}
