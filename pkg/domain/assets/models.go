package assets

import (
	"strings"

	"github.com/shopspring/decimal"
)

type Order string

var (
	Alpha        Order = "alpha"
	LessPrice    Order = "lessPrice"
	GreaterPrice Order = "greaterPrice"
	Custom       Order = "custom"
)

func NewOrder(o string) Order {
	switch strings.ToLower(o) {
	case "alpha":
		return Alpha
	case "lessPrice":
		return LessPrice
	case "greaterPrice":
		return GreaterPrice
	case "custom":
		return Custom
	default:
		return Alpha
	}
}

type Asset struct {
	Symbol   string          `json:"symbol"`
	Price    decimal.Decimal `json:"price"`
	Currency string          `json:"currency"`
}

type AssetUserEnrollment struct {
	UserID      string `json:"user_id"`
	AssetSymbol string `json:"asset_symbol"`
	Position    int    `json:"position"`
}
