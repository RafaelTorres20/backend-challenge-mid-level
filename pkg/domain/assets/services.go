package assets

import (
	"context"
	"sort"
)

type logic struct {
	assetRepo    Repository
	assetService AssetService
}

// AddUserAssets implements Logic.
func (r *logic) AddUserAssets(ctx context.Context, id string, asset Asset) error {
	if id == "" {
		return ErrUserIDNotFound
	}
	//TODO: validate asset
	if asset.Symbol == "" {
		return ErrInvalidAsset
	}

	assets, err := r.GetAssetsByUserID(ctx, id)
	if err != nil {
		return err
	}

	enrollment := AssetUserEnrollment{
		UserID:      id,
		AssetSymbol: asset.Symbol,
		Position:    len(assets),
	}
	return r.assetRepo.UpsertUserAssets(ctx, id, []AssetUserEnrollment{enrollment})
}

// GetAssetsByUserID implements Logic.
func (r *logic) GetAssetsByUserID(ctx context.Context, id string) ([]Asset, error) {
	if id == "" {
		return nil, ErrUserIDNotFound
	}
	return r.assetRepo.GetAssetsByUserID(ctx, id)
}

// GetAssetsPrices implements Logic.
func (r *logic) GetAssetsPrices(ctx context.Context, assets []Asset) ([]Asset, error) {
	for i, asset := range assets {
		if asset.Symbol == "" {
			return nil, ErrInvalidAsset
		}
		res, err := r.assetService.GetAssetBySymbol(ctx, asset.Symbol)
		if err != nil {
			return nil, err
		}
		assets[i] = res
	}
	return assets, nil
}

// OrderUserAssets implements Logic.
func (r *logic) OrderUserAssets(ctx context.Context, id string, assets []Asset, order Order) ([]Asset, error) {
	if id == "" {
		return nil, ErrUserIDNotFound
	}

	switch order {
	case Alpha:
		sort.SliceStable(assets, func(i int, j int) bool {
			return assets[i].Symbol < assets[j].Symbol
		})
	case Price:
		sort.SliceStable(assets, func(i int, j int) bool {
			return assets[i].Price.LessThan(assets[j].Price)
		})
	}
	assetUserEnrollments := make([]AssetUserEnrollment, len(assets))
	for i, asset := range assets {
		assetUserEnrollments[i] = AssetUserEnrollment{
			UserID:      id,
			AssetSymbol: asset.Symbol,
			Position:    i,
		}
	}
	return assets, r.assetRepo.UpsertUserAssets(ctx, id, assetUserEnrollments)
}

func NewAssetLogic(repo Repository, svc AssetService) Logic {
	return &logic{
		assetRepo:    repo,
		assetService: svc,
	}
}
