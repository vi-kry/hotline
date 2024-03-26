package route

import (
	"encoding/json"
	"hotline/internal/handler/api"
	"hotline/internal/service"
	"net/http"
)

type Route struct {
	service *service.Service
}

func NewRoute(service *service.Service) *Route {
	return &Route{service}
}

func (c *Route) GetRedirectURL(w http.ResponseWriter, r *http.Request) {
	redirectURL := c.service.InterfaceRedirectURL.GetRedirectURL(r.Referer())
	http.Redirect(w, r, redirectURL, http.StatusSeeOther)
}

func (c *Route) GetISSOURL(w http.ResponseWriter, r *http.Request) {
	redirectURL := c.service.InterfaceRedirectURL.GetRedirectURL(r.Referer())

	w.WriteHeader(http.StatusOK)
	_ = json.NewEncoder(w).Encode(api.SISSO{Url: &redirectURL})
}

func (c *Route) GetCallback(w http.ResponseWriter, r *http.Request, params api.GetCallbackParams) {
	token, err := c.service.InterfaceCallback.GetCallback(params.Code, params.ClientSecret, r.Referer())
	if err != nil {
		newErrorResponse(w, http.StatusBadRequest, err.Error(), err.Error())
		return
	}

	cookie := &http.Cookie{}
	cookie.Name = "accessToken"
	cookie.Value = token
	cookie.HttpOnly = true
	http.SetCookie(w, cookie)

	bearer := api.Bearer
	w.WriteHeader(http.StatusOK)
	_ = json.NewEncoder(w).Encode(api.SToken{AccessToken: &token, TokenType: &bearer})
}
