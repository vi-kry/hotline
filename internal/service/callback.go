package service

import "hotline/internal/transport"

type CallbackService struct {
	transport *transport.Transport
}

func NewCallbackService(transport *transport.Transport) *CallbackService {
	return &CallbackService{transport}
}

func (c *CallbackService) GetCallback(code, clientSecret, redirectURI string) (string, error) {
	t, err := c.transport.InterfaceISSO.PostCodeGetTokenISSO(code, clientSecret, redirectURI)
	if err != nil {
		return "", err
	}
	user, err := c.GetUserInfoFromTokenISSO(t.AccessToken)
	if err != nil {
		return "", err
	}
	tokenParams := c.transport.InterfaceToken.GetTokenParams()
	return c.GetToken(user, tokenParams)
}
