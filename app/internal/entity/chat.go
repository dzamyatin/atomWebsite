package entity

type Chat struct {
	Messenger string `db:"messenger"`
	Phone     string `db:"phone"`
	ChatID    string `db:"chat_id"`
	State     string `db:"state"`
}
