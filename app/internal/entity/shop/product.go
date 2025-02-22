package entityshop

import "github.com/google/uuid"

type ProductStatus string

const (
	ProductStatusAvailable ProductStatus = "available"
	ProductStatusSold      ProductStatus = "sold"
)

type Product struct {
	Uuid   uuid.UUID     `db:"uuid"`
	Name   string        `db:"name"`
	Status ProductStatus `db:"status"`
	Price  string        `db:"price"`
}

func NewProduct(
	name string,
	status ProductStatus,
	price string,
) *Product {
	return &Product{
		Uuid:   uuid.New(),
		Name:   name,
		Status: status,
		Price:  price,
	}
}
