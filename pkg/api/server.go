package api

import (
	"fmt"
	"net/http"

	"github.com/RafaelTorres20/backend-challenge-mid-level/pkg/domain/assets"
	"github.com/RafaelTorres20/backend-challenge-mid-level/pkg/domain/users"
	"github.com/go-chi/chi/v5"
)

type server struct {
	router         *chi.Mux
	assetsServices assets.Service
	usersService   users.Service
}

func NewServer(assetsServices assets.Service, usersService users.Service) *server {
	r := chi.NewRouter()
	server := &server{
		router:         r,
		assetsServices: assetsServices,
		usersService:   usersService,
	}

	server.router.Post("/users/", server.Create)
	server.router.Post("/users/login", server.Login)

	server.router.With(ValidateJWTMiddleware).Post("/assets/users/{id}", server.AddUserAssets)
	server.router.With(ValidateJWTMiddleware).Post("/assets/", server.AddAssets)
	server.router.With(ValidateJWTMiddleware).Get("/assets/users/{id}", server.GetAssetsByUserID)
	server.router.With(ValidateJWTMiddleware).Post("/assets/users/{id}/order", server.OrderUserAssets)
	server.router.With(ValidateJWTMiddleware).Post("/assets/prices", server.GetAssetsPrices)
	server.router.With(ValidateJWTMiddleware).Delete("/users/{id}", server.DeleteByID)
	server.router.With(ValidateJWTMiddleware).Get("/users/{id}", server.GetByID)
	server.router.With(ValidateJWTMiddleware).Get("/users/email/{email}", server.GetByEmail)
	server.router.With(ValidateJWTMiddleware).Put("/users/{id}", server.UpdateByID)

	return server
}

func (s *server) Serve(port int) error {
	return http.ListenAndServe(fmt.Sprintf(":%d", port), s.router)
}
