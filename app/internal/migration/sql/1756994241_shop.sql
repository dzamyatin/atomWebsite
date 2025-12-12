-- +goose Up
DROP TYPE IF EXISTS product_status;
CREATE TYPE product_status AS ENUM (
    'available',
    'sold'
);
CREATE TABLE IF NOT EXISTS product (
  uuid TEXT primary key,
  name TEXT NOT NULL,
  status product_status NOT NULL default 'available',
  price BIGINT NOT NULL
);
CREATE TABLE IF NOT EXISTS cart (
    uuid TEXT primary key,
    user_uuid TEXT REFERENCES users (uuid),
    created_at timestamp without time zone NOT NULL ,
    updated_at timestamp without time zone NOT NULL
);
CREATE TABLE IF NOT EXISTS cart_item (
    uuid TEXT primary key,
    cart_uuid TEXT REFERENCES cart (uuid)  NOT NULL,
    product_uuid TEXT REFERENCES product (uuid) ON DELETE CASCADE NOT NULL
);
DROP TYPE IF EXISTS order_status;
CREATE TYPE order_status AS ENUM (
    'new',
    'waiting_payment',
    'done'
);
CREATE TABLE IF NOT EXISTS orders (
    uuid TEXT primary key,
    total BIGINT NOT NULL,
    status order_status NOT NULL DEFAULT 'new'
);
CREATE TABLE IF NOT EXISTS order_item (
    uuid TEXT primary key,
    total BIGINT NOT NULL,
    name TEXT NOT NULL
);
CREATE TABLE IF NOT EXISTS documents (
    uuid TEXT primary key,
    order_item TEXT REFERENCES order_item (uuid) NOT NULL,
    content BIGINT NOT NULL,
    name TEXT NOT NULL
);
DROP TYPE IF EXISTS invoice_status;
CREATE TYPE invoice_status AS ENUM (
    'issued',
    'payed',
    'refund'
);
CREATE TABLE IF NOT EXISTS invoice (
    uuid TEXT primary key,
    order_uuid TEXT REFERENCES order_item (uuid) NOT NULL,
    total BIGINT NOT NULL,
    status invoice_status NOT NULL
);
-- +goose Down
DROP TYPE IF EXISTS product_status;
DROP TABLE IF EXISTS product;
DROP TABLE IF EXISTS cart;
DROP TABLE IF EXISTS cart_item;
DROP TYPE IF EXISTS order_status;
DROP TABLE IF EXISTS orders;
DROP TABLE IF EXISTS order_item;
DROP TABLE IF EXISTS documents;
DROP TYPE IF EXISTS invoice_status;
DROP TABLE IF EXISTS invoice;
