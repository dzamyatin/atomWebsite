package entity

type UserEntity struct {
	Email    string
	Phone    string
	password string
}

func NewUserEntity(email, phone string) UserEntity {
	return UserEntity{
		Email: email,
		Phone: phone,
	}
}

type PasswordEncoder interface {
	Encode(password string) (string, error)
}

type PasswordComparator interface {
	Compare(password string, hash string) error
}

func (r UserEntity) AddPassword(password string, passwordEncoder PasswordEncoder) error {
	encoded, err := passwordEncoder.Encode(password)

	if err != nil {
		return err
	}

	r.password = encoded

	return nil
}

func (r UserEntity) CheckPassword(password string, comparator PasswordComparator) error {
	return comparator.Compare(r.password, password)
}
