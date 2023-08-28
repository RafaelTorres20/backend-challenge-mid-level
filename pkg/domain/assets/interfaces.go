package assets

import "context"

type Service interface {
	GetAssetsByUserID(ctx context.Context, id string) ([]Asset, error)
	AddUserAssets(ctx context.Context, id string, asset Asset) error
	AddAssets(ctx context.Context, asset Asset) error
	GetAssetsPrices(ctx context.Context, assets []Asset) ([]Asset, error)
	OrderUserAssets(ctx context.Context, id string, assets []Asset, order Order) ([]Asset, error)
}

type Repository interface {
	GetAssetsByUserID(ctx context.Context, id string) ([]Asset, error)
	AddAssets(ctx context.Context, asset Asset) error
	UpsertUserAssets(ctx context.Context, id string, assets []AssetUserEnrollment) error
}

type YahooService interface {
	GetAssetBySymbol(ctx context.Context, symbol string) (Asset, error)
}
