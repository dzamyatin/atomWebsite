package repository

import (
	"database/sql"
	"errors"
	"github.com/dzamyatin/atomWebsite/internal/entity"
)

var (
	ErrUserNotFound = errors.New("user not found")
)

type IUserRepository interface {
	GetUserByEmail(email string) (entity.UserEntity, error)
	GetUserByPhone(phone string) (entity.UserEntity, error)
	AddUser(user entity.UserEntity) error
}

type UserRepository struct {
	db *sql.DB
}

func NewUserRepository(db *sql.DB) *UserRepository {
	return &UserRepository{db: db}
}

func (u *UserRepository) GetUserByEmail(email string) (entity.UserEntity, error) {
	panic("implement me")
}

func (u *UserRepository) GetUserByPhone(phone string) (entity.UserEntity, error) {
	panic("implement me")
}

func (u *UserRepository) AddUser(user entity.UserEntity) error {
	panic("implement me")
}
