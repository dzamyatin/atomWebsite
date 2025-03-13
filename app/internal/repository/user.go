package repository

import (
	"context"
	"github.com/dzamyatin/atomWebsite/internal/entity"
	"github.com/dzamyatin/atomWebsite/internal/service/db"
	"github.com/huandu/go-sqlbuilder"
	"github.com/pkg/errors"
)

var (
	ErrUserNotFound = errors.New("user not found")
)

type IUserRepository interface {
	GetUserByEmail(ctx context.Context, email string) (*entity.User, error)
	GetUserByPhone(ctx context.Context, phone string) (*entity.User, error)
	AddUser(ctx context.Context, user entity.User) error
}

type UserRepository struct {
	db db.IDatabase
}

func NewUserRepository(db db.IDatabase) *UserRepository {
	return &UserRepository{db: db}
}

func (u *UserRepository) GetUserByEmail(ctx context.Context, email string) (*entity.User, error) {
	return &entity.User{}, ErrUserNotFound
}

func (u *UserRepository) GetUserByPhone(ctx context.Context, phone string) (*entity.User, error) {
	return &entity.User{}, ErrUserNotFound
}

func (u *UserRepository) AddUser(ctx context.Context, user entity.User) error {
	sb := sqlbuilder.InsertInto("users")

	sb.Cols(
		"email",
		"password",
		"phone",
	)
	sb.Values(
		user.Email,
		user.PasswordHash,
		user.Phone,
	)

	sql, args := sb.Build()
	sql = u.db.Rebind(sql)

	_, err := u.db.Exec(ctx, sql, args...)

	return errors.Wrap(err, "error adding user")
}
