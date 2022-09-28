package main

import (
	"github.com/fazriachyar/go-auth/auth"
	"github.com/fazriachyar/go-auth/controllers"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	 e := echo.New()
	 e.GET("/user/signin", controllers.SignInForm()).Name = "userSignInForm"
	 e.POST("/user/signin", controllers.SignIn())

	 adminGroup := e.Group("/admin")

	 adminGroup.Use(middleware.JWTWithConfig(middleware.JWTConfig{
		Claims:                  &auth.Claims{},
		SigningKey:              []byte(auth.GetJWTSecret()),
		TokenLookup:             "cookie:access-token",
		ErrorHandlerWithContext: auth.JWTErrorChecker,
	}))
	 adminGroup.GET("", controllers.Admin())

	//  e.GET("/user/signin", controllers.SignInForm()).Name = "userSignInForm"
	//  e.POST("/user/signin", controllers.SignIn())
	 e.Logger.Fatal(e.Start(":1337"))
}