package gateways

import (
	"context"
	"database/sql"

	"github.com/RafaelTorres20/backend-challenge-mid-level/pkg/domain/assets"
	_ "github.com/lib/pq"
)

// "select a.name, a.price, a.currency from assets a inner join user_assets_enrollment ua on a.name = ua.asset_name where ua.user_id = $1 order by ua.position asc"

type PostgresRepository struct {
	db *sql.DB
}

// AddAssets implements assets.Repository.
func (r *PostgresRepository) AddAssets(ctx context.Context, asset assets.Asset) error {
	stmt, err := r.db.Prepare("insert into assets (symbol, currency) values ($1, $2)")
	if err != nil {
		return err
	}

	_, err = stmt.Exec(asset.Symbol, asset.Currency)
	if err != nil {
		return err
	}

	return nil
}

// GetAssetsByUserID implements assets.Repository.
func (r *PostgresRepository) GetAssetsByUserID(ctx context.Context, id string) ([]assets.Asset, error) {
	rows, err := r.db.Query(`select a.symbol, a.currency from assets a inner join AssetUserEnrollments ua on a.symbol = ua.asset_symbol where ua.user_id = $1 order by ua.position asc`, id)
	if err != nil {
		return nil, err
	}
	assetList := make([]assets.Asset, 0)
	for rows.Next() {
		asset := new(assets.Asset)
		err := rows.Scan(&asset.Symbol, &asset.Currency)
		if err != nil {
			return nil, err
		}

		assetList = append(assetList, *asset)
	}
	return assetList, nil
}

var queryUpsertUserAssets string = `INSERT INTO AssetUserEnrollments (user_id, asset_symbol, position)
									VALUES ($1, $2, $3)
									ON CONFLICT (user_id, asset_symbol)
									DO UPDATE SET position = EXCLUDED.position;`

// UpsertUserAssets implements assets.Repository.
func (r *PostgresRepository) UpsertUserAssets(ctx context.Context, id string, assets []assets.AssetUserEnrollment) error {
	stmt, err := r.db.Prepare(queryUpsertUserAssets)
	if err != nil {
		return err
	}

	tx, err := r.db.Begin()
	if err != nil {
		return err
	}

	for _, asset := range assets {
		_, err := tx.Stmt(stmt).Exec(id, asset.AssetSymbol, asset.Position)
		if err != nil {
			return err
		}
	}

	err = tx.Commit()
	if err != nil {
		return err
	}

	return nil

}

func NewPostgresRepository(db *sql.DB) assets.Repository {
	return &PostgresRepository{db: db}
}
