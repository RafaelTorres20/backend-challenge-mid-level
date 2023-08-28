package api

import (
	"log"
	"net/http"

	"github.com/RafaelTorres20/backend-challenge-mid-level/pkg/domain/assets"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/render"
)

// AddUserAssetsRequest represents the request body for AddUserAssets.
type AddUserAssetsRequest struct {
	assets.Asset
}

func (*AddUserAssetsRequest) Bind(r *http.Request) error {
	return nil
}

// AddUserAssets implements server.
func (e *server) AddUserAssets(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")

	asset := new(AddUserAssetsRequest)
	if err := render.Bind(r, asset); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	err := e.assetsServices.AddUserAssets(r.Context(), id, asset.Asset)
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
}

// GetAssetsByUserID implements server.
func (e *server) GetAssetsByUserID(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")

	assets, err := e.assetsServices.GetAssetsByUserID(r.Context(), id)
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	render.JSON(w, r, assets)
}

type GetAssetPricesRequest struct {
	Assets []assets.Asset `json:"assets"`
}

func (g *GetAssetPricesRequest) Bind(r *http.Request) error {
	if len(g.Assets) == 0 {
		return assets.ErrBadRequest
	}
	return nil
}

// GetAssetsPrices implements server.
func (e *server) GetAssetsPrices(w http.ResponseWriter, r *http.Request) {

	assets := new(GetAssetPricesRequest)
	if err := render.Bind(r, assets); err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	res, err := e.assetsServices.GetAssetsPrices(r.Context(), assets.Assets)
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	render.JSON(w, r, res)
}

type OrderUserAssetsRequest struct {
	Assets []assets.Asset `json:"assets"`
	Order  assets.Order   `json:"order"`
}

func (o *OrderUserAssetsRequest) Bind(r *http.Request) error {
	if len(o.Assets) == 0 {
		return assets.ErrBadRequest
	}
	return nil
}

// OrderUserAssets implements server.
func (e *server) OrderUserAssets(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")

	assets := new(OrderUserAssetsRequest)
	if err := render.Bind(r, assets); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	res, err := e.assetsServices.OrderUserAssets(r.Context(), id, assets.Assets, assets.Order)
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	render.JSON(w, r, res)
}

// AddAssetsRequest represents the request body for AddAssets.
type AddAssetsRequest struct {
	assets.Asset
}

func (*AddAssetsRequest) Bind(r *http.Request) error {
	return nil
}

// AddAssets implements server.
func (e *server) AddAssets(w http.ResponseWriter, r *http.Request) {
	asset := new(AddAssetsRequest)
	if err := render.Bind(r, asset); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	err := e.assetsServices.AddAssets(r.Context(), asset.Asset)
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
}
