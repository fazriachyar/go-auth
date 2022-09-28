package controllers

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
)

func Admin() echo.HandlerFunc {
	return func(c echo.Context) error{
		//ambil user cookie dan tampilkan nama dia biar keliatan keren aowkwkk
		tokenCookie, _ := c.Cookie("access-token")

		return c.String(http.StatusOK, fmt.Sprintf("%s", tokenCookie.Value))
	}
}

