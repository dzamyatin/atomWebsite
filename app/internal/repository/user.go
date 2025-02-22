package repository

import (
	"context"
	"database/sql"
	"github.com/dzamyatin/atomWebsite/internal/entity"
	"github.com/dzamyatin/atomWebsite/internal/service/db"
	"github.com/google/uuid"
	"github.com/huandu/go-sqlbuilder"
	"github.com/pkg/errors"
	"go.uber.org/zap"
)

var (
	ErrUserNotFound = errors.New("user not found")
)

type IUserRepository interface {
	GetUserByEmail(ctx context.Context, email string) (entity.User, error)
	GetUserByPhone(ctx context.Context, phone string) (entity.User, error)
	AddUser(ctx context.Context, user entity.User) error
	UpdateUser(ctx context.Context, user entity.User) error
	GetByUUID(ctx context.Context, uuid uuid.UUID) (entity.User, error)
}

type UserRepository struct {
	logger *zap.Logger
	db     db.IDatabase
}

func NewUserRepository(
	logger *zap.Logger,
	db db.IDatabase,
) *UserRepository {
	return &UserRepository{
		logger: logger,
		db:     db,
	}
}

func (r *UserRepository) GetByUUID(ctx context.Context, uuid uuid.UUID) (entity.User, error) {
	sb := sqlbuilder.PostgreSQL.NewSelectBuilder()
	sb.Select(r.getInsertCols()...)
	sb.From("users")
	sb.Where(sb.Equal("uuid", uuid))
	sb.Limit(1)

	q, args := sb.Build()

	var user entity.User
	err := r.db.Get(ctx, &user, q, args...)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return entity.User{}, ErrUserNotFound
		}

		r.logger.Error("GetUserByUUID get user by uuid error", zap.Error(err))
		return entity.User{}, err
	}

	return user, nil
}

func (r *UserRepository) GetUserByEmail(ctx context.Context, email string) (entity.User, error) {
	sb := sqlbuilder.Select(r.getInsertCols()...)

	sb.From("users")
	sb.Where(sb.ILike("email", email))

	q, args := sb.Build()

	q = r.db.Rebind(sb.String())

	user := entity.User{}
	err := r.db.Get(ctx, &user, q, args...)

	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return entity.User{}, ErrUserNotFound
		}
		return entity.User{}, errors.Wrap(err, "get user by email")
	}

	return user, nil
}

func (r *UserRepository) GetUserByPhone(ctx context.Context, phone string) (entity.User, error) {
	sb := sqlbuilder.Select(r.getInsertCols()...)

	sb.From("users")
	sb.Where(sb.ILike("phone", phone))

	q, args := sb.Build()

	q = r.db.Rebind(sb.String())

	user := entity.User{}
	err := r.db.Get(ctx, &user, q, args...)

	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return entity.User{}, ErrUserNotFound
		}
		return entity.User{}, errors.Wrap(err, "get user by phone")
	}

	return user, nil
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
		user.ConfirmedEmail,
		user.ConfirmedPhone,
	}
}

func (r *UserRepository) getInsertCols() []string {
	return []string{
		"uuid",
		"email",
		"password",
		"phone",
		"confirmed_email",
		"confirmed_phone",
	}
}

func (r *UserRepository) UpdateUser(ctx context.Context, user entity.User) error {
	ub := Builder.NewUpdateBuilder()

	ub.Update("users")
	ub.Where(ub.Equal("uuid", user.UUID))

	cols := r.getInsertCols()
	vals := r.getInsertVal(user)

	for i, col := range cols {
		ub.SetMore(ub.Assign(col, vals[i]))
	}

	sql, args := ub.Build()
	sql = r.db.Rebind(sql)

	_, err := r.db.Exec(ctx, sql, args...)

	return errors.Wrap(err, "error update user")
}
