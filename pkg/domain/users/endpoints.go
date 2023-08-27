package users

import (
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/render"
)

type Endpoints struct {
	router       *chi.Mux
	usersService UserService
}

func NewEndpoints(svc UserService) Endpoints {
	endpoints := Endpoints{
		usersService: svc,
	}
	router := chi.NewRouter()
	router.Post("/", endpoints.Create)
	router.Delete("/{id}", endpoints.DeleteByID)
	router.Get("/{id}", endpoints.GetByID)
	router.Get("/email/{email}", endpoints.GetByEmail)
	router.Put("/{id}", endpoints.UpdateByID)
	endpoints.router = router
	return endpoints
}

func (e *Endpoints) Router() *chi.Mux {
	return e.router
}

// CreateRequest represents the request body for Create.
type CreateRequest struct {
	User
}

func (*CreateRequest) Bind(r *http.Request) error {
	return nil
}

// Create implements Endpoints.
func (e *Endpoints) Create(w http.ResponseWriter, r *http.Request) {
	user := new(CreateRequest)
	if err := render.Bind(r, user); err != nil {
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

// DeleteByID implements Endpoints.
func (e *Endpoints) DeleteByID(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")

	err := e.usersService.DeleteByID(r.Context(), id)
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}

// GetByEmail implements Endpoints.
func (e *Endpoints) GetByEmail(w http.ResponseWriter, r *http.Request) {
	email := chi.URLParam(r, "email")

	user, err := e.usersService.GetByEmail(r.Context(), email)
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

	render.JSON(w, r, user)
}

// GetByID implements Endpoints.
func (e *Endpoints) GetByID(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")

	user, err := e.usersService.GetByID(r.Context(), id)
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	render.JSON(w, r, user)
}

// UpdateByID implements Endpoints.
func (e *Endpoints) UpdateByID(w http.ResponseWriter, r *http.Request) {
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
