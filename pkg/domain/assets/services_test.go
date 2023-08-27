package assets

// import (
// 	"context"
// 	"testing"

// 	"github.com/shopspring/decimal"
// 	"github.com/stretchr/testify/assert"
// )

// var (
// 	correctAsset  = Asset{Symbol: "GOGL", Price: decimal.NewFromFloat(1000), Currency: "BRL"}
// 	notExistAsset = Asset{Symbol: "NOTEXIST", Price: decimal.NewFromFloat(1000), Currency: "BRL"}
// )

// func TestGetAssetsByUserID(t *testing.T) {
// 	var tests = []struct {
// 		symbol         string
// 		expectedError  error
// 		given          string
// 		expectedAssets []Asset
// 	}{
// 		{symbol: "Should return error", expectedError: ErrUserNotFound, given: "randomUserID"},
// 		{symbol: "Should return assets by given ID", expectedError: nil, expectedAssets: []Asset{correctAsset}, given: "idtest"},
// 		{symbol: "Should return no content", expectedError: nil, given: "idNoContent"},
// 	}
// 	for _, tt := range tests {
// 		tt := tt
// 		t.Run(tt.symbol, func(t *testing.T) {
// 			assetLogic := NewAssetLogic()
// 			assets, err := assetLogic.GetAssetsByUserID(context.Background(), tt.given)
// 			if err != nil {
// 				assert.Equal(t, tt.expectedError, err)
// 			}
// 			assert.Equal(t, tt.expectedAssets, assets)

// 		})
// 	}
// }

// func TestAddUserAssets(t *testing.T) {
// 	var tests = []struct {
// 		symbol        string
// 		expectedError error
// 		givenAsset    Asset
// 		givenUserID   string
// 	}{
// 		{symbol: "Should return error when user not exist", expectedError: ErrUserNotFound, givenUserID: "randomUserID"},
// 		{symbol: "Should return invalid asset when asset is not valid", expectedError: ErrInvalidAsset, givenAsset: notExistAsset, givenUserID: "randomUserID"},
// 		{symbol: "Should return internal server error when a random error appears", expectedError: ErrInternalServerError, givenUserID: "randomUserID", givenAsset: correctAsset},
// 		{symbol: "Should return without error", expectedError: nil, givenUserID: "randomUserID", givenAsset: correctAsset},
// 	}
// 	for _, tt := range tests {
// 		tt := tt
// 		t.Run(tt.symbol, func(t *testing.T) {
// 			assetLogic := NewAssetLogic()
// 			err := assetLogic.AddUserAssets(context.Background(), tt.givenUserID, tt.givenAsset)
// 			if err != nil {
// 				assert.Equal(t, tt.expectedError, err)
// 				return
// 			}
// 			assert.Nil(t, err)
// 		})
// 	}
// }

// func TestGetAssetsPrices(t *testing.T) {
// 	var tests = []struct {
// 		symbol         string
// 		expectedError  error
// 		expectedAssets []Asset
// 		givenAssets    []Asset
// 	}{
// 		{symbol: "Should return error when assets is invalid", expectedError: ErrInvalidAsset, expectedAssets: []Asset{correctAsset}, givenAssets: []Asset{correctAsset}},
// 		{symbol: "Should return internal server error when a random error appears", expectedError: ErrInternalServerError, expectedAssets: []Asset{correctAsset}, givenAssets: []Asset{notExistAsset}},
// 		{symbol: "Should return assets with prices when assets exists", expectedError: nil, expectedAssets: []Asset{correctAsset}, givenAssets: []Asset{notExistAsset}},
// 		{symbol: "Should return error when assets is invalid", expectedError: ErrInvalidAsset, expectedAssets: []Asset{correctAsset}, givenAssets: []Asset{notExistAsset}},
// 	}
// 	for _, tt := range tests {
// 		tt := tt
// 		t.Run(tt.symbol, func(t *testing.T) {
// 			assetLogic := NewAssetLogic()
// 			assets, err := assetLogic.GetAssetsPrices(context.Background(), tt.givenAssets)
// 			if err != nil {
// 				assert.Equal(t, tt.expectedError, err)
// 				return
// 			}
// 			assert.Equal(t, tt.expectedAssets, assets)
// 		})
// 	}
// }
