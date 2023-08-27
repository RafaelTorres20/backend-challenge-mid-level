package gateways

import (
	"context"
	"encoding/json"
	"io"
	"net/http"

	"github.com/RafaelTorres20/backend-challenge-mid-level/pkg/domain/assets"
	"github.com/shopspring/decimal"
)

type yahooClient struct {
	client  *http.Client
	baseURL string
	apiKey  string
}

type Result struct {
	Price    string `json:"regularMarketPrice"`
	Symbol   string `json:"symbol"`
	Currency string `json:"currency"`
}

type QuoteResponse struct {
	Result []Result `json:"result"`
}

type YahooResponseDTO struct {
	QuoteResponse QuoteResponse `json:"quoteResponse"`
}

// GetAssetByName implements assets.AssetService.
func (c *yahooClient) GetAssetBySymbol(ctx context.Context, symbol string) (assets.Asset, error) {
	//TODO: implement http request
	req, err := http.NewRequest("GET", "https://yfapi.net/v6/finance/quote", nil)
	if err != nil {
		return assets.Asset{}, err
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("X-API-KEY", c.apiKey)
	query := req.URL.Query()
	query.Add("symbols", symbol)
	query.Add("region", "US")
	query.Add("lang", "en")

	req.URL.RawQuery = query.Encode()

	res, err := c.client.Do(req)
	if err != nil {
		return assets.Asset{}, err
	}
	defer res.Body.Close()

	responseDTO := new(YahooResponseDTO)
	b, err := io.ReadAll(res.Body)
	if err != nil {
		return assets.Asset{}, err
	}
	err = json.Unmarshal(b, responseDTO)
	if err != nil {
		return assets.Asset{}, err
	}

	if len(responseDTO.QuoteResponse.Result) == 0 {
		return assets.Asset{}, assets.ErrAssetNotFound
	}

	price, err := decimal.NewFromString(responseDTO.QuoteResponse.Result[0].Price)
	if err != nil {
		return assets.Asset{}, err
	}

	asset := assets.Asset{
		Symbol:   responseDTO.QuoteResponse.Result[0].Symbol,
		Currency: responseDTO.QuoteResponse.Result[0].Currency,
		Price:    price,
	}

	return asset, nil
}

func NewYahooClient(client *http.Client, baseURL string, apiKey string) assets.AssetService {
	if client == nil {
		client = http.DefaultClient
	}
	return &yahooClient{
		client:  client,
		baseURL: baseURL,
		apiKey:  apiKey,
	}
}
