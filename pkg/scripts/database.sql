-- Criação da tabela Asset
CREATE TABLE IF NOT EXISTS Assets (
    symbol VARCHAR(10) NOT NULL PRIMARY KEY,
    currency VARCHAR(3) NOT NULL
);

-- Criação da tabela User
CREATE TABLE IF NOT EXISTS Users(
    id VARCHAR(36) PRIMARY KEY,
    email VARCHAR(100) NOT NULL,
    password VARCHAR(255) NOT NULL
);

-- Criação da tabela AssetUserEnrollment
CREATE TABLE IF NOT EXISTS AssetUserEnrollments (
    user_id VARCHAR(36) REFERENCES Users (id),
    asset_symbol VARCHAR(255) REFERENCES Assets (symbol),
    position INT NOT NULL,
    PRIMARY KEY (user_id, asset_symbol)
);
