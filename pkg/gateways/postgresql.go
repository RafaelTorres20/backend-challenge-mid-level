package gateways

// "select a.name, a.price, a.currency from assets a inner join user_assets_enrollment ua on a.name = ua.asset_name where ua.user_id = $1 order by ua.position asc"
