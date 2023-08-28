package api

import (
	"fmt"
	"log"
	"net/http"

	"github.com/RafaelTorres20/backend-challenge-mid-level/pkg/domain/users"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/render"
)

// CreateRequest represents the request body for Create.
type CreateRequest struct {
	users.User
}

func (*CreateRequest) Bind(r *http.Request) error {
	return nil
}

// Create implements server.
func (e *server) Create(w http.ResponseWriter, r *http.Request) {
	user := new(CreateRequest)
	if err := render.Bind(r, user); err != nil {
		fmt.Println(err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	err := e.usersService.Create(r.Context(), &user.User)
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
}

// DeleteByID implements server.
func (e *server) DeleteByID(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")

	err := e.usersService.DeleteByID(r.Context(), id)
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}

// GetByEmail implements server.
func (e *server) GetByEmail(w http.ResponseWriter, r *http.Request) {
	email := chi.URLParam(r, "email")

	user, err := e.usersService.GetByEmail(r.Context(), email)
	if user == nil {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("user not foundd"))
		return
	}
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	render.JSON(w, r, user)
}

// GetByID implements server.
func (e *server) GetByID(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")

	user, err := e.usersService.GetByID(r.Context(), id)
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	render.JSON(w, r, user)
}

// UpdateByID implements server.
func (e *server) UpdateByID(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")

	user := new(CreateRequest)
	if err := render.Bind(r, user); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	err := e.usersService.UpdateByID(r.Context(), id, &user.User)
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}

func (e *server) Login(w http.ResponseWriter, r *http.Request) {
	createRequest := new(CreateRequest)
	if err := render.Bind(r, createRequest); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	email := createRequest.Email
	pass := createRequest.Password
	user, token, err := e.usersService.Login(r.Context(), email, pass)
	if user == nil {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("user not found"))
		return
	}

	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	render.JSON(w, r, map[string]string{
		"token": token,
		"email": user.Email,
		"id":    user.ID,
	})
}
