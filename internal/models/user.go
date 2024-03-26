package models

import (
	"fmt"
	"github.com/golang-jwt/jwt"
)

type UserInfo struct {
	Email    string `json:"email"`
	Username string `json:"username"`
	Fullname string `json:"fullname"`
	Name     string `json:"name"`
	Domain   string `json:"domain"`
}

func (u *UserInfo) GetUserInfo(claims jwt.MapClaims) {
	domain := claims["domain"]
	name := claims["name"]
	fullname := claims["fullname"]
	username := claims["username"]
	email := claims["email"]

	if name != nil {
		u.Name = fmt.Sprint(name)
	}
	if domain != nil {
		u.Domain = fmt.Sprint(domain)
	}
	if fullname != nil {
		u.Fullname = fmt.Sprint(fullname)
	}
	if username != nil {
		u.Username = fmt.Sprint(username)
	}
	if email != nil {
		u.Email = fmt.Sprint(email)
	}
}
