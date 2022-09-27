package controllers

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func Admin() echo.HandlerFunc {
	return func(c echo.Context) error{
		return c.String(http.StatusOK, "Selamat Kamu mendapat akses!")
	}
}