package api

func New(es *storage.Conn) *User {
	return &User{es: es}
}

type User struct {
	es *storage.Conn
}
