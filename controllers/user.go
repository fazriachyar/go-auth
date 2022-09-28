package controllers

import (
	"github.com/fazriachyar/go-auth/auth"
	"github.com/fazriachyar/go-auth/user"

	"html/template"
	"net/http"
	"path"

	"github.com/labstack/echo/v4"
	"golang.org/x/crypto/bcrypt"
)

//SignInForm berfungsi untuk menghandle form dalam proses rendering form
func SignInForm() echo.HandlerFunc {
	return func(c echo.Context) error {
		fp := path.Join("templates", "signIn.html")
		tmpl, err := template.ParseFiles(fp)
		if err != nil {
			return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
		}
		if err := tmpl.Execute(c.Response().Writer, nil);
		err != nil {
			return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
		}

		return nil
	}
}

//fungsi SignIn akan dieksekusi setelah SignInForm berhasil di eksekusi
func SignIn() echo.HandlerFunc {
	return func(c echo.Context) error {
		//disini kita akan memuat "test" user yang telah kita buat tadi
		storedUser := user.LoadTestUser()
		//inisialisasi sebuah struct user yang baru
		u := new(user.User)
		//parsing data yang telah di submit di SignInForm dan isi struct User dengan
		// data dari SignInForm
		if err := c.Bind(u); err != nil {
			return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
		}
		// bandingkan password yang telah di hashed dengan password yang diterima
		if err := bcrypt.CompareHashAndPassword([]byte(storedUser.Password), []byte(u.Password)); err != nil {
			//jika password tidak sesuai return status 401
			return echo.NewHTTPError(http.StatusUnauthorized, "Password tidak sesuai")
		}// jika password sesuai, generate token dan set cookie nya
		err := auth.GenerateTokensAndSetCookies(storedUser, c)

		if err != nil {
			return echo.NewHTTPError(http.StatusUnauthorized, "Token tidak sesuai")
		}

		//userCookie, _ := c.Cookie("token")

		 return c.Redirect(http.StatusMovedPermanently, "/admin")
		//return c.String(http.StatusOK, fmt.Sprintf("%s", userCookie.Value))
	}
}