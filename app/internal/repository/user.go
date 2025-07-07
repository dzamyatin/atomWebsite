package repository

import (
	"context"
	"database/sql"
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
	UpdateUser(ctx context.Context, user entity.User) error
}

type UserRepository struct {
	db db.IDatabase
}

func NewUserRepository(db db.IDatabase) *UserRepository {
	return &UserRepository{db: db}
}

func (r *UserRepository) GetUserByEmail(ctx context.Context, email string) (*entity.User, error) {
	sb := sqlbuilder.Select(r.getInsertCols()...)

	sb.From("users")
	sb.Where(sb.ILike("email", email))

	q, args := sb.Build()

	q = r.db.Rebind(sb.String())

	user := entity.User{}
	err := r.db.Get(ctx, &user, q, args...)

	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, ErrUserNotFound
		}
		return nil, errors.Wrap(err, "get user by email")
	}

	return &user, nil
}

func (r *UserRepository) GetUserByPhone(ctx context.Context, phone string) (*entity.User, error) {
	return &entity.User{}, ErrUserNotFound
}

func (r *UserRepository) AddUser(ctx context.Context, user entity.User) error {
	sb := sqlbuilder.InsertInto("users")

	sb.Cols(r.getInsertCols()...)

	user.GenerateUUID()

	sb.Values(r.getInsertVal(user)...)

	sql, args := sb.Build()
	sql = r.db.Rebind(sql)

	_, err := r.db.Exec(ctx, sql, args...)

	return errors.Wrap(err, "error adding user")
}

func (r *UserRepository) getInsertVal(user entity.User) []any {
	return []any{
		user.UUID,
		user.Email,
		user.PasswordHash,
		user.Phone,
	}
}

func (r *UserRepository) getInsertCols() []string {
	return []string{
		"uuid",
		"email",
		"password",
		"phone",
	}
}

func (r *UserRepository) UpdateUser(ctx context.Context, user entity.User) error {
	ub := Builder.NewUpdateBuilder()

	ub.Update("users")

	cols := r.getInsertCols()
	vals := r.getInsertVal(user)

	for i, col := range cols {
		ub.Set(ub.Assign(col, vals[i]))
	}

	sql, args := ub.Build()
	sql = r.db.Rebind(sql)

	_, err := r.db.Exec(ctx, sql, args...)

	return errors.Wrap(err, "error adding user")
}
