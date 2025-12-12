package entityshop

import "github.com/google/uuid"

type ProductStatus string

const (
	ProductStatusAvailable ProductStatus = "available"
	ProductStatusSold      ProductStatus = "sold"
)

type ProductType string

const (
	ProductTypeDefault      ProductType = "default"
	ProductTypeProxyPlanOne ProductType = "proxy_plan_one"
)

type ProductUuid uuid.UUID

type Product struct {
	Uuid   ProductUuid   `db:"uuid"`
	Name   string        `db:"name"`
	Status ProductStatus `db:"status"`
	Price  string        `db:"price"`
	Type   ProductType   `db:"type"`
}

func NewProduct(
	name string,
	status ProductStatus,
	price string,
	productType ProductType,
) *Product {
	return &Product{
		Uuid:   ProductUuid(uuid.New()),
		Name:   name,
		Status: status,
		Price:  price,
		Type:   productType,
	}
}
