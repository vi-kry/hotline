package token

import "hotline/internal/models"

type Token struct {
	tokenConfig models.Token
}

func NewToken(tokenConfig models.Token) *Token {
	return &Token{tokenConfig}
}

func (t *Token) GetTokenParams() models.Token {
	return t.tokenConfig
}
