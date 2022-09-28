package auth

import (
	"net/http"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/fazriachyar/go-auth/user"
	"github.com/labstack/echo/v4"
)

const (
	accessTokenCookieName = "access-token"

	//pada kasus real, jwt secret key harus didapatkan dari file .env atau json dengan viper
	jwtSecretKey = "some-secret-key"
)

func GetJWTSecret() string{
	return jwtSecretKey
}

//buat struct yang akan mengencode jwt
//jwt.StandardClaims adalah type yang di embedd untuk memberikan field waktu kadaluarsa
type Claims struct {
	Name string `json:"name"`
	jwt.StandardClaims
}

//fungsi GenerateTokensAndSetCookies akan mengenerate jwt token dan akan disimpan di http-only cookie
func GenerateTokensAndSetCookies(user *user.User, c echo.Context) error{
	accessToken, exp, err := generateAccessToken(user)
	if err != nil {
		return err
	}

	setTokenCookie(accessTokenCookieName, accessToken, exp, c)
	setUserCookie(user, exp, c)

	return nil
}

func generateAccessToken(user *user.User) (string, time.Time, error) {
	//deklarasikan waktu kadaluarsa, dalam kasus ini token akan kadaluarsa dalam waktu 1 jam
	expirationTime := time.Now().Add(1 * time.Hour)

	return generateToken(user, expirationTime, []byte(GetJWTSecret()))
}

//fungsi utama yang harus diperhatikan karena
// fungsi ini yang akan menampung logic dalam pembuatan logic jwt nya
func generateToken(user *user.User, expirationTime time.Time, secret []byte) (string, time.Time, error) {
	//buatlah variabel Claims untuk JWT, yang berisi username dan waktu kadaluarsa
	claims := &Claims{
		Name: user.Name,
		StandardClaims: jwt.StandardClaims{
			//dalam jwt, waktu kadaluarsa di ekpresikan dengan menggunakan unix milisecond
			ExpiresAt: expirationTime.Unix(),
		},
	}

	//Deklarasikan token dengan algoritma H256 digunakan untuk signin dan mengklaim token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	//buat string JWT
	tokenString, err := token.SignedString(secret)
	if err != nil {
		return "", time.Now(), err
	}

	return tokenString, expirationTime, nil
}

//disini kita membuat cookie baru, yang akan digunakan untuk menyimpan token jwt yang valdi
func setTokenCookie(name, token string, expiration time.Time, c echo.Context) {
	cookie := new(http.Cookie)
	cookie.Name = name
	cookie.Value = token
	cookie.Expires = expiration
	cookie.Path = "/"
	//http-only akan membantu kita untuk
	//mengurangi resiko pada sisi client dalam pengaksesan cookie menggunakan script
	cookie.HttpOnly = true

	c.SetCookie(cookie)
}

//disini kita membuat cookie baru untuk menyimpan nama dari user
func setUserCookie(user *user.User,expiration time.Time,c echo.Context) {
	cookie := new(http.Cookie)
	cookie.Name = "user"
	cookie.Value = user.Name
	cookie.Expires = expiration
	cookie.Path = "/"
	c.SetCookie(cookie)
}

//JWTErrorChecker akan dieksekusi apabila user mencoba untuk mengakses path yang sudah dilindungi
func JWTErrorChecker(err error, c echo.Context) error{
	//redirect ke form sign in
	return c.Redirect(http.StatusMovedPermanently, c.Echo().Reverse("userSignInForm"))
}



