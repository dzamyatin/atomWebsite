package entityshop

import "github.com/google/uuid"

type OrderStatus string

const (
	OrderStatusNew            OrderStatus = "new"
	OrderStatusWaitingPayment OrderStatus = "waiting_payment"
	OrderStatusDone           OrderStatus = "done"
)

type OrderUuid uuid.UUID

type Order struct {
	Uuid   OrderUuid   `db:"uuid"`
	Total  uint64      `db:"total"`
	Status OrderStatus `db:"status"`
}
