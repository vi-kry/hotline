package handler

import (
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"
	"hotline/internal/handler/api"
	"hotline/internal/handler/route"
	"hotline/internal/service"
	"net/http"
	"path"
)

type Handler struct {
	service *service.Service
}

func NewHandler(service *service.Service) *Handler {
	return &Handler{service}
}

func (h *Handler) InitRoutes(basePath, env string) http.Handler {
	routes := route.NewRoute(h.service)
	router := chi.NewRouter()

	router.Use(middleware.Logger)
	router.Use(cors.New(cors.Options{
		AllowedOrigins:   []string{"*"},
		AllowedMethods:   []string{http.MethodGet, http.MethodPost, http.MethodPut, http.MethodDelete, http.MethodOptions},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: false,
		MaxAge:           300,
	}).Handler)

	router.Use(middleware.SetHeader("Content-Type", "application/json"))
	router.Use(middleware.SetHeader("Accept", "application/json"))

	router.Route(path.Join("/", basePath), func(router chi.Router) {
		router.Use(middleware.NoCache)
		router.Mount("/", api.Handler(routes))
	})
	return router
}
