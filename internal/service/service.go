package service

import (
	"hotline/internal/models"
	"hotline/internal/transport"
)

type InterfaceCallback interface {
	GetCallback(code, clientSecret, host string) (string, error)
}

type InterfaceRedirectURL interface {
	GetRedirectURL(host string) string
}

type InterfaceToken interface {
	GetTokenParams() models.Token
}

type Service struct {
	InterfaceCallback
	InterfaceRedirectURL
	InterfaceToken
}

func NewService(t *transport.Transport) *Service {
	return &Service{
		InterfaceCallback:    NewCallbackService(t),
		InterfaceRedirectURL: NewRedirectURLService(t),
	}
}
