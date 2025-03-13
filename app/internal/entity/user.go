package entity

type User struct {
	UUID         string `db:"uuid"`
	Email        string `db:"email"`
	Phone        string `db:"phone"`
	PasswordHash string `db:"password"`
}

func NewUser(email, phone string) *User {
	return &User{
		Email: email,
		Phone: phone,
	}
}

type PasswordEncoder interface {
	Encode(password string) (string, error)
}

type PasswordComparator interface {
	Compare(password string, hash string) (ok bool, err error)
}

func (r *User) AddPassword(password string, passwordEncoder PasswordEncoder) error {
	encoded, err := passwordEncoder.Encode(password)

	if err != nil {
		return err
	}

	r.PasswordHash = encoded

	return nil
}

func (r *User) CheckPassword(password string, comparator PasswordComparator) (ok bool, err error) {
	return comparator.Compare(r.PasswordHash, password)
}
