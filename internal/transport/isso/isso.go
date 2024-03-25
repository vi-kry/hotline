package isso

import (
	"encoding/json"
	"fmt"
	"hotline/internal/models"
	"net/http"
	"net/url"
)

const (
	path = "/auth/realms/mts/protocol/openid-connect/token"
)

type ISSO struct {
	cfgISSO models.ISSO
}

func NewISSO(cfgISSO models.ISSO) *ISSO {
	return &ISSO{cfgISSO}
}

func (i *ISSO) PostCodeGetTokenISSO(code, clientSecret, redirectURI string) (models.TokenISSO, error) {
	var res models.TokenISSO

	data := url.Values{}
	data.Set("code", code)
	data.Set("client_id", i.cfgISSO.ClientID)
	data.Set("client_secret", clientSecret)
	data.Set("redirect_uri", redirectURI)
	data.Set("grant_type", i.cfgISSO.GrantType)

	resp, err := http.PostForm(i.cfgISSO.URL+path, data)
	if err != nil {
		return res, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		var resErr models.TokenISSOError
		err = json.NewDecoder(resp.Body).Decode(&resErr)
		if err != nil {
			return res, err
		}
		textErr := fmt.Errorf(`status_code: %d; error: %s; error_description: %s`,
			resp.StatusCode, resErr.Error, resErr.ErrorDescription)

		return res, textErr
	}

	err = json.NewDecoder(resp.Body).Decode(&res)
	if err != nil {
		return res, err
	}

	return res, nil
}

func (i *ISSO) GetISSOParams() models.ISSO {
	return i.cfgISSO
}
