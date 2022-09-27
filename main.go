package main

import (
	"github.com/fazriachyar/controllers"
	"github.com/labstack/echo/v4"
)

func main() {
	 e := echo.New()

	 adminGroup := e.Group("/admin")

	 adminGroup.GET("", controllers.Admin())
	 e.Logger.Fatal(e.Start(":1337"))
}