package assets

import (
	"context"
	"testing"

	"github.com/shopspring/decimal"
	"github.com/stretchr/testify/assert"
)

var (
	correctAsset  = Asset{Name: "GOGL", Price: decimal.NewFromFloat(1000), Currency: "BRL"}
	notExistAsset = Asset{Name: "NOTEXIST", Price: decimal.NewFromFloat(1000), Currency: "BRL"}
)

func TestGetAssetsByUserID(t *testing.T) {
	var tests = []struct {
		name           string
		expectedError  error
		given          string
		expectedAssets []Asset
	}{
		{name: "Should return error", expectedError: ErrUserNotFound, given: "randomUserID"},
		{name: "Should return assets by given ID", expectedError: nil, expectedAssets: []Asset{correctAsset}, given: "idtest"},
		{name: "Should return no content", expectedError: nil, given: "idNoContent"},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			assetLogic := NewAssetLogic()
			assets, err := assetLogic.GetAssetsByUserID(context.Background(), tt.given)
			if err != nil {
				assert.Equal(t, tt.expectedError, err)
			}
			assert.Equal(t, tt.expectedAssets, assets)

		})
	}
}

func TestAddUserAssets(t *testing.T) {
	var tests = []struct {
		name          string
		expectedError error
		givenAsset    Asset
		givenUserID   string
	}{
		{name: "Should return error when user not exist", expectedError: ErrUserNotFound, givenUserID: "randomUserID"},
		{name: "Should return invalid asset when asset is not valid", expectedError: ErrInvalidAsset, givenAsset: notExistAsset, givenUserID: "randomUserID"},
		{name: "Should return internal server error when a random error appears", expectedError: ErrInternalServerError, givenUserID: "randomUserID", givenAsset: correctAsset},
		{name: "Should return without error", expectedError: nil, givenUserID: "randomUserID", givenAsset: correctAsset},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			assetLogic := NewAssetLogic()
			err := assetLogic.AddUserAssets(context.Background(), tt.givenUserID, tt.givenAsset)
			if err != nil {
				assert.Equal(t, tt.expectedError, err)
				return
			}
			assert.Nil(t, err)
		})
	}
}

func TestGetAssetsPrices(t *testing.T) {
	var tests = []struct {
		name           string
		expectedError  error
		expectedAssets []Asset
		givenAssets    []Asset
	}{
		{name: "Should return error when assets is invalid", expectedError: ErrInvalidAsset, expectedAssets: []Asset{correctAsset}, givenAssets: []Asset{correctAsset}},
		{name: "Should return internal server error when a random error appears", expectedError: ErrInternalServerError, expectedAssets: []Asset{correctAsset}, givenAssets: []Asset{notExistAsset}},
		{name: "Should return assets with prices when assets exists", expectedError: nil, expectedAssets: []Asset{correctAsset}, givenAssets: []Asset{notExistAsset}},
		{name: "Should return error when assets is invalid", expectedError: ErrInvalidAsset, expectedAssets: []Asset{correctAsset}, givenAssets: []Asset{notExistAsset}},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			assetLogic := NewAssetLogic()
			assets, err := assetLogic.GetAssetsPrices(context.Background(), tt.givenAssets)
			if err != nil {
				assert.Equal(t, tt.expectedError, err)
				return
			}
			assert.Equal(t, tt.expectedAssets, assets)
		})
	}
}
