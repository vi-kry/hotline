package service

import (
	"fmt"
	"hotline/internal/transport"
)

const (
	path = "/auth/realms/mts/protocol/openid-connect/auth"
)

type RedirectURLService struct {
	transport *transport.Transport
}

func NewRedirectURLService(transport *transport.Transport) *RedirectURLService {
	return &RedirectURLService{transport}
}

func (r *RedirectURLService) GetRedirectURL(host string) string {
	i := r.transport.InterfaceISSO.GetISSOParams()
	fmt.Println(i.URL + path +
		"?" + "response_type=code" +
		"&" + "client_id" + i.ClientID +
		"&" + "redirect_uri=" + host)
	return i.URL + path +
		"?" + "response_type=code" +
		"&" + "client_id" + i.ClientID +
		"&" + "redirect_uri=" + host
}
