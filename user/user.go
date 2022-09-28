package user

import (
	"golang.org/x/crypto/bcrypt"
)

type User struct {
	Password string `json:"password" form:"password"`
	Name string `json:"name" form:"name"`
}

func LoadTestUser() *User {
	//hanya untuk pengetesan, dalam project nyata kita harus mengambil
	//data user dari database langsung.
	//dalam kasus ini kita akan menggunakan mock data dengan
	//password yang telah di encrypt password := test
	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte("test"), 8)
	return &User{
		Password: string(hashedPassword),
		Name: "Test user",
	}
}