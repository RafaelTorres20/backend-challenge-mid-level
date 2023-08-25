package domain

import "github.com/shopspring/decimal"

type Asset struct {
	Name     string          `json:"name"`
	Price    decimal.Decimal `json:"price"`
	Currency string          `json:"currency"`
}

type User struct {
	ID     string  `json:"id"`
	Assets []Asset `json:"assets"`
}
