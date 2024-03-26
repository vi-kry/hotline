package service

import (
	"github.com/golang-jwt/jwt"
	"hotline/internal/models"
	"time"
)

type tokenClaims struct {
	jwt.StandardClaims
	models.UserInfo
}

func (c *CallbackService) GetUserInfoFromTokenISSO(tokenString string) (models.UserInfo, error) {
	var item models.UserInfo

	token, _, err := new(jwt.Parser).ParseUnverified(tokenString, jwt.MapClaims{})
	if err != nil {
		return item, err
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok {
		item.GetUserInfo(claims)
	}

	return item, nil
}

func (c *CallbackService) GetToken(user models.UserInfo, params models.Token) (string, error) {
	tokenTTL := time.Duration(params.ExpiresAt) * time.Minute
	createToken := jwt.NewWithClaims(jwt.SigningMethodHS256, &tokenClaims{
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(tokenTTL).Unix(),
			IssuedAt:  time.Now().Unix(),
			Subject:   user.Username,
		},
		models.UserInfo{
			Email:    user.Email,
			Username: user.Username,
			Fullname: user.Fullname,
			Name:     user.Name,
			Domain:   user.Domain,
		},
	})
	return createToken.SignedString([]byte(params.Secret))
}
