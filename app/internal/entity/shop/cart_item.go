package entityshop

import "github.com/google/uuid"

type CartItemUuid uuid.UUID

type CartItem struct {
	Uuid        CartItemUuid `db:"uuid"`
	CartUuid    CartUuid     `db:"cart_uuid"`
	ProductUuid string       `db:"product_uuid"`
}

func NewCartItem(
	cartUuid CartUuid,
	productUuid string,
) *CartItem {
	return &CartItem{
		Uuid:        CartItemUuid(uuid.New()),
		CartUuid:    cartUuid,
		ProductUuid: productUuid,
	}
}
