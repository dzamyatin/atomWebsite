package entityshop

import "github.com/google/uuid"

type OrderItemUuid uuid.UUID

type OrderItem struct {
	Uuid  OrderItemUuid `db:"uuid"`
	Price uint64        `db:"price"`
	Name  string        `db:"name"`
}
