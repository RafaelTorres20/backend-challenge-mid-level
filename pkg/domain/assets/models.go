package assets

import (
	"strings"

	"github.com/shopspring/decimal"
)

type Order string

var (
	Alpha  Order = "alpha"
	Price  Order = "price"
	Custom Order = "custom"
)

func NewOrder(o string) Order {
	switch strings.ToLower(o) {
	case "alpha":
		return Alpha
	case "price":
		return Price
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

type User struct {
	ID     string  `json:"id"`
	Assets []Asset `json:"assets"`
}

type AssetUserEnrollment struct {
	UserID    string `json:"user_id"`
	AssetName string `json:"asset_name"`
	Position  int    `json:"position"`
}
