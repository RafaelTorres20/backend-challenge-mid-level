package api

import (
	"fmt"
	"net/http"

	"github.com/RafaelTorres20/backend-challenge-mid-level/pkg/domain/assets"
	"github.com/RafaelTorres20/backend-challenge-mid-level/pkg/domain/users"
	"github.com/go-chi/chi/v5"
)

type server struct {
	assetsEndpoints assets.Endpoints
	usersEndpoints  users.Endpoints
	router          *chi.Mux
}

func NewServer(assetsEndpoints assets.Endpoints, usersEndpoints users.Endpoints) *server {
	server := &server{
		assetsEndpoints: assetsEndpoints,
		router:          chi.NewRouter(),
	}
	server.router.Mount("/assets", assetsEndpoints.Router())
	server.router.Mount("/users", usersEndpoints.Router())

	return server
}

func (s *server) Serve(port int) error {
	return http.ListenAndServe(fmt.Sprintf(":%d", port), s.router)
}
