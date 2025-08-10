package entity

type Chat struct {
	Messenger string `db:"messenger"`
	ChatID    string `db:"chat_id"`
	State     string `db:"state"`
}
