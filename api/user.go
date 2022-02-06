package api

import (
	"github.com/feamo/esrest/storage"
	//"github.com/feamo/esrest/models"
)

func New(es *storage.Conn) *User {
	return &User{es: es}
}

type User struct {
	es *storage.Conn
}
