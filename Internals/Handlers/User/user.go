package user

import (
	"net/http"

	db "github.con/Ibukun-tech/trust/Db"
)

type User struct {
	Db db.DB
}

func NewUser(dbConnect db.DB) *User {
	return &User{
		Db: dbConnect,
	}
}

func (u *User) CreateUser(w http.ResponseWriter, r *http.Request) error {
	return nil
}
