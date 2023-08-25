package domain

import (
	"context"
)

type assetLogic struct {
	// assetRepo AssetRepo
}

// AddUserAssets implements AssetLogic.
func (*assetLogic) AddUserAssets(ctx context.Context, id string, asset Asset) error {
	panic("unimplemented")
}

// GetAssetsByUserID implements AssetLogic.
func (*assetLogic) GetAssetsByUserID(ctx context.Context, id string) ([]Asset, error) {
	panic("unimplemented")
}

// GetAssetsPrices implements AssetLogic.
func (*assetLogic) GetAssetsPrices(ctx context.Context, assets []Asset) ([]Asset, error) {
	panic("unimplemented")
}

// OrderUserAssets implements AssetLogic.
func (*assetLogic) OrderUserAssets(ctx context.Context, id string, assets []Asset) ([]Asset, error) {
	panic("unimplemented")
}

func NewAssetLogic() AssetLogic {
	return &assetLogic{}
}
