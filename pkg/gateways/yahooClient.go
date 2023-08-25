package gateways

import (
	"context"
	"net/http"

	"github.com/RafaelTorres20/backend-challenge-mid-level/pkg/domain/assets"
)

type yahooClient struct {
	client  *http.Client
	baseURL string
}

// GetAssetByName implements assets.AssetService.
func (*yahooClient) GetAssetByName(ctx context.Context, name string) (assets.Asset, error) {
	//TODO: implement http request

	return assets.Asset{}, nil
}

func NewYahooClient(client *http.Client, baseURL string) assets.AssetService {
	if client == nil {
		client = http.DefaultClient
	}
	return &yahooClient{
		client:  client,
		baseURL: baseURL,
	}
}
