package transport

import (
	"hotline/internal/models"
	"hotline/internal/transport/isso"
	"hotline/internal/transport/token"
)

type InterfaceISSO interface {
	PostCodeGetTokenISSO(code, clientSecret, redirectURI string) (models.TokenISSO, error)
	GetISSOParams() models.ISSO
}

type InterfaceToken interface {
	GetTokenParams() models.Token
}

type Transport struct {
	InterfaceISSO
	InterfaceToken
}

func NewTransport(cfg *models.Config) *Transport {
	return &Transport{
		InterfaceISSO:  isso.NewISSO(cfg.ISSO),
		InterfaceToken: token.NewToken(cfg.Token),
	}
}
