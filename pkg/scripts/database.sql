-- Criação da tabela Asset
CREATE TABLE Asset (
    id SERIAL PRIMARY KEY,
    symbol VARCHAR(255) NOT NULL,
    price NUMERIC(15, 2) NOT NULL,
    currency VARCHAR(3) NOT NULL
);

-- Criação da tabela User
CREATE TABLE User(
    id VARCHAR(36) PRIMARY KEY,
    email VARCHAR(100) NOT NULL,
    password VARCHAR(255) NOT NULL
);

-- Criação da tabela AssetUserEnrollment
CREATE TABLE AssetUserEnrollment (
    user_id VARCHAR(36) REFERENCES User (id),
    asset_symbol VARCHAR(255) REFERENCES Asset (symbol),
    position INT NOT NULL,
    PRIMARY KEY (user_id, asset_symbol)
);
