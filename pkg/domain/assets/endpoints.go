package assets

import (
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/render"
)

type Endpoints struct {
	logic  Logic
	router *chi.Mux
}

func NewEndpoints(logic Logic) Endpoints {
	endpoints := Endpoints{
		logic: logic,
	}
	router := chi.NewRouter()
	router.Post("/users/{id}", endpoints.AddUserAssets)
	router.Post("/", endpoints.AddAssets)
	router.Get("/users/{id}", endpoints.GetAssetsByUserID)
	router.Post("/users/{id}/order", endpoints.OrderUserAssets)
	router.Post("/prices", endpoints.GetAssetsPrices)
	endpoints.router = router
	return endpoints
}

func (e *Endpoints) Router() *chi.Mux {
	return e.router
}

// AddUserAssetsRequest represents the request body for AddUserAssets.
type AddUserAssetsRequest struct {
	Asset
}

func (*AddUserAssetsRequest) Bind(r *http.Request) error {
	return nil
}

// AddUserAssets implements Endpoints.
func (e *Endpoints) AddUserAssets(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")

	asset := new(AddUserAssetsRequest)
	if err := render.Bind(r, asset); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	err := e.logic.AddUserAssets(r.Context(), id, asset.Asset)
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
}

// GetAssetsByUserID implements Endpoints.
func (e *Endpoints) GetAssetsByUserID(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")

	assets, err := e.logic.GetAssetsByUserID(r.Context(), id)
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	render.JSON(w, r, assets)
}

type GetAssetPricesRequest struct {
	Assets []Asset `json:"assets"`
}

func (g *GetAssetPricesRequest) Bind(r *http.Request) error {
	if len(g.Assets) == 0 {
		return ErrBadRequest
	}
	return nil
}

// GetAssetsPrices implements Endpoints.
func (e *Endpoints) GetAssetsPrices(w http.ResponseWriter, r *http.Request) {

	assets := new(GetAssetPricesRequest)
	if err := render.Bind(r, assets); err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	res, err := e.logic.GetAssetsPrices(r.Context(), assets.Assets)
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	render.JSON(w, r, res)
}

type OrderUserAssetsRequest struct {
	Assets []Asset `json:"assets"`
	Order  Order   `json:"order"`
}

func (o *OrderUserAssetsRequest) Bind(r *http.Request) error {
	if len(o.Assets) == 0 {
		return ErrBadRequest
	}
	return nil
}

// OrderUserAssets implements Endpoints.
func (e *Endpoints) OrderUserAssets(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")

	assets := new(OrderUserAssetsRequest)
	if err := render.Bind(r, assets); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	res, err := e.logic.OrderUserAssets(r.Context(), id, assets.Assets, assets.Order)
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	render.JSON(w, r, res)
}

// AddAssetsRequest represents the request body for AddAssets.
type AddAssetsRequest struct {
	Asset
}

func (*AddAssetsRequest) Bind(r *http.Request) error {
	return nil
}

// AddAssets implements Endpoints.
func (e *Endpoints) AddAssets(w http.ResponseWriter, r *http.Request) {
	asset := new(AddAssetsRequest)
	if err := render.Bind(r, asset); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	err := e.logic.AddAssets(r.Context(), asset.Asset)
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
}
