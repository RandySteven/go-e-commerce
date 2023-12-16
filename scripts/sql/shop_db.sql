CREATE DATABASE e_commerce_db;

CREATE SEQUENCE IF NOT EXISTS my_sequence

CREATE TABLE IF NOT EXISTS users (
    id BIGSERIAL PRIMARY KEY NOT NULL,
    name VARCHAR NOT NULL,
    email VARCHAR UNIQUE NOT NULL,
    password VARCHAR NOT NULL,
    birthday TIMESTAMP NOT NULL,
    phone_number VARCHAR NOT NULL UNIQUE,
    created_at TIMESTAMP NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMP NOT NULL DEFAULT NOW(),
    deleted_at TIMESTAMP
);

CREATE TABLE IF NOT EXISTS shops (
    id BIGSERIAL PRIMARY KEY NOT NULL,
    name VARCHAR UNIQUE NOT NULL,
    email VARCHAR UNIQUE NOT NULL,
    password VARCHAR NOT NULL,
    phone_number VARCHAR NOT NULL UNIQUE,
    created_at TIMESTAMP NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMP NOT NULL DEFAULT NOW(),
    deleted_at TIMESTAMP
);

CREATE TABLE IF NOT EXISTS provinces (
    id BIGSERIAL PRIMARY KEY NOT NULL,
    province VARCHAR UNIQUE NOT NULL
);

CREATE TABLE IF NOT EXISTS cities (
    id BIGSERIAL PRIMARY KEY NOT NULL,
    city VARCHAR UNIQUE NOT NULL,
    province_id BIGINT NOT NULL,
    CONSTRAINT fk_province_id FOREIGN KEY (province_id) REFERENCES provinces (id)
);

CREATE TABLE IF NOT EXISTS postal_codes (
    id BIGSERIAL PRIMARY KEY NOT NULL,
    postal_code VARCHAR NOT NULL UNIQUE,
    city_id BIGINT NOT NULL,
    CONSTRAINT fk_city_id FOREIGN KEY (city_id) REFERENCES cities(id)
);

CREATE TABLE IF NOT EXISTS addresses(
    id BIGSERIAL PRIMARY KEY NOT NULL,
    province_id BIGINT NOT NULL,
    city_id BIGINT NOT NULL,
    postal_code_id BIGINT NOT NULL,
    address VARCHAR NOT NULL,
    created_at TIMESTAMP NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMP NOT NULL DEFAULT NOW(),
    deleted_at TIMESTAMP NULL,
    CONSTRAINT fk_province_id FOREIGN KEY (province_id) REFERENCES provinces (id),
    CONSTRAINT fk_city_id FOREIGN KEY (city_id) REFERENCES cities (id),
    CONSTRAINT fk_postal_code_id FOREIGN KEY (postal_code_id) REFERENCES postal_codes (id)
);

CREATE TABLE IF NOT EXISTS user_addresses (
    id BIGSERIAL PRIMARY KEY NOT NULL,
    user_id BIGINT NOT NULL,
    address_id BIGINT NOT NULL,
    created_at TIMESTAMP NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMP NOT NULL DEFAULT NOW(),
    deleted_at TIMESTAMP,
    CONSTRAINT fk_user_id FOREIGN KEY (user_id) REFERENCES users(id),
    CONSTRAINT fk_address_id FOREIGN KEY (address_id) REFERENCES addresses(id)
);

CREATE TABLE IF NOT EXISTS shop_address (
    id BIGSERIAL PRIMARY KEY NOT NULL,
    shop_id BIGINT NOT NULL UNIQUE,
    address_id BIGINT NOT NULL UNIQUE,
    created_at TIMESTAMP NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMP NOT NULL DEFAULT NOW(),
    deleted_at TIMESTAMP,
    CONSTRAINT fk_shop_id FOREIGN KEY (shop_id) REFERENCES shops(id),
    CONSTRAINT fk_address_id FOREIGN KEY (address_id) REFERENCES addresses(id)
);

CREATE TABLE IF NOT EXISTS categories (
    id BIGSERIAL PRIMARY KEY NOT NULL,
    category VARCHAR NOT NULL UNIQUE,
    created_at TIMESTAMP NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMP NOT NULL DEFAULT NOW(),
    deleted_at TIMESTAMP
);

CREATE TABLE IF NOT EXISTS products (
    id BIGSERIAL PRIMARY KEY NOT NULL,
    name VARCHAR NOT NULL,
    price NUMERIC NOT NULL DEFAULT 0,
    stock INT NOT NULL DEFAULT 0,
    description VARCHAR,
    shop_id BIGINT NOT NULL,
    category_id BIGINT NOT NULL,
    created_at TIMESTAMP NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMP NOT NULL DEFAULT NOW(),
    deleted_at TIMESTAMP,
    CONSTRAINT fk_shop_id FOREIGN KEY (shop_id) REFERENCES shops (id),
    CONSTRAINT fk_category_id FOREIGN KEY (category_id) REFERENCES categories (id)
);

CREATE TABLE IF NOT EXISTS carts (
    id BIGSERIAL PRIMARY KEY NOT NULL,
    shop_id BIGINT NOT NULL,
    user_id BIGINT NOT NULL,
    product_id BIGINT NOT NULL,
    quantity INT NOT NULL DEFAULT 0,
    created_at TIMESTAMP NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMP NOT NULL DEFAULT NOW(),
    deleted_at TIMESTAMP,
    CONSTRAINT fk_shop_id FOREIGN KEY (shop_id) REFERENCES shops (id),
    CONSTRAINT fk_product_id FOREIGN KEY (product_id) REFERENCES products (id),
    CONSTRAINT fk_user_id FOREIGN KEY (user_id) REFERENCES users (id)
);

CREATE TABLE IF NOT EXISTS transactions (
    id BIGSERIAL PRIMARY KEY NOT NULL,
    shop_id BIGINT NOT NULL,
    user_id BIGINT NOT NULL,
    transaction_code VARCHAR NOT NULL UNIQUE,
    transaction_date TIMESTAMP NOT NULL DEFAULT NOW(),
    created_at TIMESTAMP NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMP NOT NULL DEFAULT NOW(),
    deleted_at TIMESTAMP,
    CONSTRAINT fk_shop_id FOREIGN KEY (shop_id) REFERENCES shops (id),
    CONSTRAINT fk_user_id FOREIGN KEY (user_id) REFERENCES users (id)
);

CREATE TABLE IF NOT EXISTS transaction_details (
    id BIGSERIAL PRIMARY KEY NOT NULL,
    transaction_id BIGINT NOT NULL,
    product_id BIGINT NOT NULL,
    quantity INT NOT NULL DEFAULT 0,
    created_at TIMESTAMP NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMP NOT NULL DEFAULT NOW(),
    deleted_at TIMESTAMP,
    CONSTRAINT fk_product_id FOREIGN KEY (product_id) REFERENCES products (id),
    CONSTRAINT fk_transaction_id FOREIGN KEY (transaction_id) REFERENCES transactions (id)
);