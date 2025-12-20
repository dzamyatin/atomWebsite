-- +goose Up
DROP TYPE IF EXISTS product_status;
CREATE TYPE product_status AS ENUM (
    'available',
    'sold'
);
DROP TYPE IF EXISTS product_type;
CREATE TYPE product_type AS ENUM (
    'default',
    'proxy_plan_one'
);
CREATE TABLE IF NOT EXISTS product (
  uuid TEXT primary key,
  name TEXT NOT NULL,
  type product_type NOT NULL default 'default',
  status product_status NOT NULL default 'available',
  price BIGINT NOT NULL
);
CREATE TABLE IF NOT EXISTS cart (
    uuid TEXT primary key,
    user_uuid TEXT REFERENCES users (uuid),
    created_at timestamp without time zone NOT NULL,
    updated_at timestamp without time zone NOT NULL
);
CREATE TABLE IF NOT EXISTS cart_item (
    uuid TEXT primary key,
    cart_uuid TEXT REFERENCES cart (uuid)  NOT NULL,
    product_uuid TEXT REFERENCES product (uuid) ON DELETE CASCADE NOT NULL
);
DROP TYPE IF EXISTS order_status;
CREATE TYPE order_status AS ENUM (
    'canceled',
    'processing',
    'done'
);
CREATE TABLE IF NOT EXISTS orders (
    uuid TEXT primary key,
    total BIGINT NOT NULL,
    status order_status NOT NULL DEFAULT 'processing'
);
CREATE TABLE IF NOT EXISTS order_item (
    uuid TEXT primary key,
    price BIGINT NOT NULL,

    product_name TEXT NOT NULL,
    product_uuid TEXT NOT NULL
);

CREATE TABLE IF NOT EXISTS documents (
    uuid TEXT primary key,
    order_item TEXT REFERENCES order_item (uuid) NOT NULL,
    content BIGINT NOT NULL,
    name TEXT NOT NULL
);
DROP TYPE IF EXISTS payment_status;
CREATE TYPE payment_status AS ENUM (
    'new',
    'authorized',
    'canceled',
    'payed',
    'refunded',
    'partial_refund'
);
CREATE TABLE IF NOT EXISTS payment (
    uuid TEXT primary key,
    order_uuid TEXT REFERENCES order_item (uuid) NOT NULL,
    amount BIGINT NOT NULL,
    refund BIGINT NOT NULL,
    status payment_status NOT NULL
);
-- +goose Down

DROP TABLE IF EXISTS cart_item;
DROP TABLE IF EXISTS cart;
DROP TABLE IF EXISTS documents;
DROP TABLE IF EXISTS payment;
DROP TABLE IF EXISTS order_item;
DROP TABLE IF EXISTS orders;
DROP TABLE IF EXISTS product;
DROP TYPE IF EXISTS payment_status;
DROP TYPE IF EXISTS product_status;
DROP TYPE IF EXISTS order_status;