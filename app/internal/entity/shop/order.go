package entityshop

import "github.com/google/uuid"

type OrderStatus string

const (
	OrderStatusCanceled   OrderStatus = "canceled"
	OrderStatusProcessing OrderStatus = "processing"
	OrderStatusDone       OrderStatus = "done"
)

type OrderUuid uuid.UUID

type Order struct {
	Uuid   OrderUuid   `db:"uuid"`
	Total  uint64      `db:"total"`
	Status OrderStatus `db:"status"`
}
