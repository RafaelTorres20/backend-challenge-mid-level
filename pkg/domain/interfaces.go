package domain

import "context"

type AssetLogic interface {
	GetAssetsByUserID(ctx context.Context, id string) ([]Asset, error)
	AddUserAssets(ctx context.Context, id string, asset Asset) error
	GetAssetsPrices(ctx context.Context, assets []Asset) ([]Asset, error)
	OrderUserAssets(ctx context.Context, id string, assets []Asset) ([]Asset, error)
}
