package user

import (
	"context"
	"encoding/json"
	"net/http"

	db "github.con/Ibukun-tech/trust/Db"
	model "github.con/Ibukun-tech/trust/Internals/Models"
)

type User struct {
	Db db.StoreWithTx
}

func NewUser(dbConnect db.StoreWithTx) *User {
	return &User{
		Db: dbConnect,
	}
}

func (u *User) CreateUser(w http.ResponseWriter, r *http.Request) error {
	if r.Method == http.MethodPost {
		
		var 
		var user model.User
		if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
			return err
		}
		if err := user.ValidateUser(); err != nil {
			return err
		}
		if err := u.Db.ExecuteTrx(context.Background(), func(dc *db.DbConnect) error {
			userDb, err := dc.CreateUser(context.Background(), &user)

			// I am about to create the accoun simultaneously at the same time
			if err != nil {
				return err
			}
			return nil
		}); err != nil {
			return err
		}
		// u.Db.ExecuteTrx(context.Background(), func(dc *db.DbConnect) error {
		// 	dc.CreateAccount(r.Context())
		// 	return nil
		// })
	}
	return nil
}
func (u *User) ListUser(w http.ResponseWriter, r *http.Request) error {
	if r.Method == http.MethodGet {
	}
	return nil
}

func (u *User) DeleteUser(w http.ResponseWriter, r *http.Request) error {
	if r.Method == http.MethodDelete {
	}
	return nil
}
