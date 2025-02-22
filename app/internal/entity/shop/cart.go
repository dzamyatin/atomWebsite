package entityshop

import (
	"time"

	"github.com/dzamyatin/atomWebsite/internal/entity"
	"github.com/google/uuid"
)

type CartUuid uuid.UUID

type Cart struct {
	Uuid       CartUuid        `db:"uuid"`
	UserUuid   entity.UserUuid `db:"user_uuid"`
	created_at time.Time       `db:"created_at"`
	updated_at time.Time       `db:"updated_at"`
}

func NewCart(
	userUuid entity.UserUuid,
	created_at time.Time,
	updated_at time.Time,
) *Cart {
	return &Cart{
		Uuid:       CartUuid(uuid.New()),
		UserUuid:   userUuid,
		created_at: created_at,
		updated_at: updated_at,
	}
}
